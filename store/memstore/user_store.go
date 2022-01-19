// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemUserStore struct {
	MemStore *MemStore
	users    []*model.User
	mutex    sync.RWMutex
}

func (us *MemUserStore) ClearCaches() {}

func (us *MemUserStore) InvalidateProfileCacheForUser(userId string) {}

func newMemUserStore(memStore *MemStore) store.UserStore {
	return &MemUserStore{MemStore: memStore}
}

func (us *MemUserStore) Save(user *model.User) (*model.User, error) {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	if user.Id != "" && !user.IsRemote() {
		return nil, store.NewErrInvalidInput("User", "id", user.Id)
	}

	user.PreSave()
	if err := user.IsValid(); err != nil {
		return nil, err
	}

	for _, u := range us.users {
		if u.Email == user.Email {
			return nil, store.NewErrInvalidInput("User", "email", user.Email)
		}
		if u.Username == user.Username {
			return nil, store.NewErrInvalidInput("User", "username", user.Username)
		}
	}

	us.users = append(us.users, user)

	return user, nil
}

func (us *MemUserStore) DeactivateGuests() ([]string, error) {
	deletedUsers := []string{}
	curTime := model.GetMillis()
	for _, u := range us.users {
		if u.Roles == "system_guest" && u.DeleteAt != 0 {
			u.UpdateAt = curTime
			u.DeleteAt = curTime
			deletedUsers = append(deletedUsers, u.Id)
		}
	}
	return deletedUsers, nil
}

func (us *MemUserStore) Update(user *model.User, trustedUpdateData bool) (*model.UserUpdate, error) {
	user.PreUpdate()

	if err := user.IsValid(); err != nil {
		return nil, err
	}

	oldUserResult, err := us.Get(context.Background(), user.Id)
	if err != nil {
		return nil, store.NewErrInvalidInput("User", "id", user.Id)
	}

	oldUser := oldUserResult
	user.CreateAt = oldUser.CreateAt
	user.AuthData = oldUser.AuthData
	user.AuthService = oldUser.AuthService
	user.Password = oldUser.Password
	user.LastPasswordUpdate = oldUser.LastPasswordUpdate
	user.LastPictureUpdate = oldUser.LastPictureUpdate
	user.EmailVerified = oldUser.EmailVerified
	user.FailedAttempts = oldUser.FailedAttempts
	user.MfaSecret = oldUser.MfaSecret
	user.MfaActive = oldUser.MfaActive

	if !trustedUpdateData {
		user.Roles = oldUser.Roles
		user.DeleteAt = oldUser.DeleteAt
	}

	if user.IsOAuthUser() {
		if !trustedUpdateData {
			user.Email = oldUser.Email
		}
	} else if user.IsLDAPUser() && !trustedUpdateData {
		if user.Username != oldUser.Username || user.Email != oldUser.Email {
			return nil, store.NewErrInvalidInput("User", "id", user.Id)
		}
	} else if user.Email != oldUser.Email {
		user.EmailVerified = false
	}

	if user.Username != oldUser.Username {
		user.UpdateMentionKeysFromUsername(oldUser.Username)
	}

	var foundUser *model.User
	for _, u := range us.users {
		if u.Id != user.Id && u.Email == user.Email {
			return nil, store.NewErrConflict("Email", err, user.Email)
		}
		if u.Id != user.Id && u.Username == user.Username {
			return nil, store.NewErrConflict("Username", err, user.Username)
		}
		if u.Id == user.Id {
			if foundUser != nil {
				return nil, fmt.Errorf("multiple users were update: userId=%s, count=%d", user.Id, 2)
			}
			if foundUser == nil {
				foundUser = u
			}
		}
	}

	if foundUser != nil {
		*foundUser = *user
	}

	user.Sanitize(map[string]bool{})
	oldUser.Sanitize(map[string]bool{})
	return &model.UserUpdate{New: user, Old: oldUser}, nil
}

func (us *MemUserStore) UpdateNotifyProps(userID string, props map[string]string) error {
	panic("not implemented")
}

func (us *MemUserStore) UpdateLastPictureUpdate(userId string) error {
	panic("not implemented")
}

func (us *MemUserStore) ResetLastPictureUpdate(userId string) error {
	panic("not implemented")
}

func (us *MemUserStore) UpdateUpdateAt(userId string) (int64, error) {
	curTime := model.GetMillis()
	for _, u := range us.users {
		if u.Id == userId {
			u.UpdateAt = curTime
			return curTime, nil
		}
	}
	return curTime, nil
}

func (us *MemUserStore) UpdatePassword(userId, hashedPassword string) error {
	panic("not implemented")
}

func (us *MemUserStore) UpdateFailedPasswordAttempts(userId string, attempts int) error {
	panic("not implemented")
}

func (us *MemUserStore) UpdateAuthData(userId string, service string, authData *string, email string, resetMfa bool) (string, error) {
	panic("not implemented")
}

// ResetAuthDataToEmailForUsers resets the AuthData of users whose AuthService
// is |service| to their Email. If userIDs is non-empty, only the users whose
// IDs are in userIDs will be affected. If dryRun is true, only the number
// of users who *would* be affected is returned; otherwise, the number of
// users who actually were affected is returned.
func (us *MemUserStore) ResetAuthDataToEmailForUsers(service string, userIDs []string, includeDeleted bool, dryRun bool) (int, error) {
	panic("not implemented")
}

func (us *MemUserStore) UpdateMfaSecret(userId, secret string) error {
	panic("not implemented")
}

func (us *MemUserStore) UpdateMfaActive(userId string, active bool) error {
	panic("not implemented")
}

// GetMany returns a list of users for the provided list of ids
func (us *MemUserStore) GetMany(ctx context.Context, ids []string) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) Get(ctx context.Context, id string) (*model.User, error) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()
	fmt.Println(us.users)

	for _, u := range us.users {
		if u.Id == id {
			return u, nil
		}
	}
	return nil, store.NewErrNotFound("User", id)
}

func (us *MemUserStore) GetAll() ([]*model.User, error) {
	return us.users, nil
}

func (us *MemUserStore) GetAllAfter(limit int, afterId string) ([]*model.User, error) {
	result := []*model.User{}
	copyUsers := []*model.User{}
	for _, u := range us.users {
		copyUsers = append(copyUsers, u)
	}
	sort.Slice(copyUsers, func(i, j int) bool { return copyUsers[i].Id > copyUsers[j].Id })
	counter := 0
	for _, u := range copyUsers {
		if counter > limit {
			break
		}
		if u.Id > afterId {
			result = append(result, u)
			counter++
		}
	}
	return result, nil
}

func (us *MemUserStore) GetEtagForAllProfiles() string {
	panic("not implemented")
}

func (us *MemUserStore) GetAllProfiles(options *model.UserGetOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetEtagForProfiles(teamId string) string {
	panic("not implemented")
}

func (us *MemUserStore) GetProfiles(options *model.UserGetOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) InvalidateProfilesInChannelCacheByUser(userId string) {}

func (us *MemUserStore) InvalidateProfilesInChannelCache(channelId string) {}

func (us *MemUserStore) GetProfilesInChannel(options *model.UserGetOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetProfilesInChannelByStatus(options *model.UserGetOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetAllProfilesInChannel(ctx context.Context, channelID string, allowFromCache bool) (map[string]*model.User, error) {
	result := map[string]*model.User{}
	for _, cm := range us.MemStore.Channel().(*MemChannelStore).members {
		if cm.ChannelId == channelID {
			for _, u := range us.users {
				if cm.UserId == u.Id {
					result[u.Id] = u
				}
			}
		}
	}
	return result, nil
}

func (us *MemUserStore) GetProfilesNotInChannel(teamId string, channelId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, error) {
	result := []*model.User{}
	for _, u := range us.users {
		isMember := false
		for _, cm := range us.MemStore.Channel().(*MemChannelStore).members {
			if cm.ChannelId == channelId && cm.UserId == u.Id {
				isMember = true
				break
			}
		}
		if !isMember {
			result = append(result, u)
		}
	}
	return result, nil
}

func (us *MemUserStore) GetProfilesWithoutTeam(options *model.UserGetOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetProfilesByUsernames(usernames []string, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetRecentlyActiveUsersForTeam(teamId string, offset, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetNewUsersForTeam(teamId string, offset, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetProfileByIds(ctx context.Context, userIds []string, options *store.UserGetByIdsOpts, allowFromCache bool) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetProfileByGroupChannelIdsForUser(userId string, channelIds []string) (map[string][]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetSystemAdminProfiles() (map[string]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetByEmail(email string) (*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetByAuth(authData *string, authService string) (*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetAllUsingAuthService(authService string) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetAllNotInAuthService(authServices []string) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetByUsername(username string) (*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetForLogin(loginId string, allowSignInWithUsername, allowSignInWithEmail bool) (*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) VerifyEmail(userId, email string) (string, error) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	for _, u := range us.users {
		if u.Id == userId {
			u.EmailVerified = true
			u.Email = strings.ToLower(email)
			u.UpdateAt = model.GetMillis()
			return userId, nil
		}
	}
	return "", errors.New("unable to find the user")
}

func (us *MemUserStore) PermanentDelete(userId string) error {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	result := []*model.User{}
	for _, u := range us.users {
		if u.Id != userId {
			result = append(result, u)
		}
	}
	us.users = result
	return nil
}

func (us *MemUserStore) Count(options model.UserCountOptions) (int64, error) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	var counter int64 = 0

	for _, u := range us.users {
		if !options.IncludeDeleted && u.DeleteAt == 0 {
			continue
		}
		// bot, _ := us.MemStore.Bots().Get(u.Id)
		// TODO: Add this
		var bot *model.Bot = nil
		if options.IncludeBotAccounts && bot != nil {
			continue
		}
		if options.ExcludeRegularUsers && bot == nil {
			continue
		}
		// TODO
		if options.TeamId != "" {
			// TODO: Check team memberships
		}
		if options.ChannelId != "" {
			// TODO: Check team memberships
		}
		counter++
	}
	return counter, nil
}

func (us *MemUserStore) AnalyticsActiveCount(timePeriod int64, options model.UserCountOptions) (int64, error) {
	panic("not implemented")
}

func (us *MemUserStore) AnalyticsActiveCountForPeriod(startTime int64, endTime int64, options model.UserCountOptions) (int64, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetUnreadCount(userId string) (int64, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetUnreadCountForChannel(userId string, channelId string) (int64, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetAnyUnreadPostCountForChannel(userId string, channelId string) (int64, error) {
	panic("not implemented")
}

func (us *MemUserStore) Search(teamId string, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) SearchWithoutTeam(term string, options *model.UserSearchOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) SearchNotInTeam(notInTeamId string, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) SearchNotInChannel(teamId string, channelId string, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) SearchInChannel(channelId string, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) SearchInGroup(groupID string, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) performSearch(query sq.SelectBuilder, term string, options *model.UserSearchOptions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) AnalyticsGetInactiveUsersCount() (int64, error) {
	panic("not implemented")
}

func (us *MemUserStore) AnalyticsGetExternalUsers(hostDomain string) (bool, error) {
	panic("not implemented")
}

func (us *MemUserStore) AnalyticsGetGuestCount() (int64, error) {
	panic("not implemented")
}

func (us *MemUserStore) AnalyticsGetSystemAdminCount() (int64, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetProfilesNotInTeam(teamId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetEtagForProfilesNotInTeam(teamId string) string {
	panic("not implemented")
}

func (us *MemUserStore) ClearAllCustomRoleAssignments() error {
	panic("not implemented")
}

func (us *MemUserStore) InferSystemInstallDate() (int64, error) {
	return model.GetMillis(), nil
}

func (us *MemUserStore) GetUsersBatchForIndexing(startTime, endTime int64, limit int) ([]*model.UserForIndexing, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetTeamGroupUsers(teamID string) ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetChannelGroupUsers(channelID string) ([]*model.User, error) {
	panic("not implemented")
}

func applyViewRestrictionsFilter(query sq.SelectBuilder, restrictions *model.ViewUsersRestrictions, distinct bool) sq.SelectBuilder {
	panic("not implemented")
}

func (us *MemUserStore) PromoteGuestToUser(userId string) error {
	panic("not implemented")
}

func (us *MemUserStore) DemoteUserToGuest(userID string) (*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) AutocompleteUsersInChannel(teamId, channelId, term string, options *model.UserSearchOptions) (*model.UserAutocompleteInChannel, error) {
	panic("not implemented")
}

// GetKnownUsers returns the list of user ids of users with any direct
// relationship with a user. That means any user sharing any channel, including
// direct and group channels.
func (us *MemUserStore) GetKnownUsers(userId string) ([]string, error) {
	panic("not implemented")
}

// IsEmpty returns whether or not the Users table is empty.
func (us *MemUserStore) IsEmpty(excludeBots bool) (bool, error) {
	if len(us.users) > 0 {
		return false, nil
	}
	return true, nil
}
