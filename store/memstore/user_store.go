// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemUserStore struct{}

func (us *MemUserStore) ClearCaches() {}

func (us *MemUserStore) InvalidateProfileCacheForUser(userId string) {}

func newMemUserStore() store.UserStore {
	return &MemUserStore{}
}

func (us *MemUserStore) Save(user *model.User) (*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) DeactivateGuests() ([]string, error) {
	panic("not implemented")
}

func (us *MemUserStore) Update(user *model.User, trustedUpdateData bool) (*model.UserUpdate, error) {
	panic("not implemented")
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
	panic("not implemented")
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
	panic("not implemented")
}

func (us *MemUserStore) GetAll() ([]*model.User, error) {
	panic("not implemented")
}

func (us *MemUserStore) GetAllAfter(limit int, afterId string) ([]*model.User, error) {
	panic("not implemented")
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
	panic("not implemented")
}

func (us *MemUserStore) GetProfilesNotInChannel(teamId string, channelId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, error) {
	panic("not implemented")
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
	panic("not implemented")
}

func (us *MemUserStore) PermanentDelete(userId string) error {
	panic("not implemented")
}

func (us *MemUserStore) Count(options model.UserCountOptions) (int64, error) {
	panic("not implemented")
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
	panic("not implemented")
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
	panic("not implemented")
}
