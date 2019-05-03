package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentUserStore struct {
	baseStore store.UserStore
}

func (s TransparentUserStore) Save(user *model.User) store.StoreChannel {
	return s.baseStore.Save(user)
}

func (s TransparentUserStore) Update(user *model.User, allowRoleUpdate bool) store.StoreChannel {
	return s.baseStore.Update(user, allowRoleUpdate)
}

func (s TransparentUserStore) UpdateLastPictureUpdate(userId string) store.StoreChannel {
	return s.baseStore.UpdateLastPictureUpdate(userId)
}

func (s TransparentUserStore) ResetLastPictureUpdate(userId string) store.StoreChannel {
	return s.baseStore.ResetLastPictureUpdate(userId)
}

func (s TransparentUserStore) UpdateUpdateAt(userId string) store.StoreChannel {
	return s.baseStore.UpdateUpdateAt(userId)
}

func (s TransparentUserStore) UpdatePassword(userId string, newPassword string) store.StoreChannel {
	return s.baseStore.UpdatePassword(userId, newPassword)
}

func (s TransparentUserStore) UpdateAuthData(userId string, service string, authData *string, email string, resetMfa bool) store.StoreChannel {
	return s.baseStore.UpdateAuthData(userId, service, authData, email, resetMfa)
}

func (s TransparentUserStore) UpdateMfaSecret(userId string, secret string) store.StoreChannel {
	return s.baseStore.UpdateMfaSecret(userId, secret)
}

func (s TransparentUserStore) UpdateMfaActive(userId string, active bool) store.StoreChannel {
	return s.baseStore.UpdateMfaActive(userId, active)
}

func (s TransparentUserStore) Get(id string) (*model.User, *model.AppError) {
	return s.baseStore.Get(id)
}

func (s TransparentUserStore) GetAll() store.StoreChannel {
	return s.baseStore.GetAll()
}

func (s TransparentUserStore) ClearCaches() {
	s.baseStore.ClearCaches()
}

func (s TransparentUserStore) InvalidateProfilesInChannelCacheByUser(userId string) {
	s.baseStore.InvalidateProfilesInChannelCacheByUser(userId)
}

func (s TransparentUserStore) InvalidateProfilesInChannelCache(channelId string) {
	s.baseStore.InvalidateProfilesInChannelCache(channelId)
}

func (s TransparentUserStore) GetProfilesInChannel(channelId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetProfilesInChannel(channelId, offset, limit)
}

func (s TransparentUserStore) GetProfilesInChannelByStatus(channelId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetProfilesInChannelByStatus(channelId, offset, limit)
}

func (s TransparentUserStore) GetAllProfilesInChannel(channelId string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetAllProfilesInChannel(channelId, allowFromCache)
}

func (s TransparentUserStore) GetProfilesNotInChannel(teamId string, channelId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetProfilesNotInChannel(teamId, channelId, offset, limit, viewRestrictions)
}

func (s TransparentUserStore) GetProfilesWithoutTeam(offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetProfilesWithoutTeam(offset, limit, viewRestrictions)
}

func (s TransparentUserStore) GetProfilesByUsernames(usernames []string, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetProfilesByUsernames(usernames, viewRestrictions)
}

func (s TransparentUserStore) GetAllProfiles(options *model.UserGetOptions) store.StoreChannel {
	return s.baseStore.GetAllProfiles(options)
}

func (s TransparentUserStore) GetProfiles(options *model.UserGetOptions) store.StoreChannel {
	return s.baseStore.GetProfiles(options)
}

func (s TransparentUserStore) GetProfileByIds(userId []string, allowFromCache bool, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetProfileByIds(userId, allowFromCache, viewRestrictions)
}

func (s TransparentUserStore) InvalidatProfileCacheForUser(userId string) {
	s.baseStore.InvalidatProfileCacheForUser(userId)
}

func (s TransparentUserStore) GetByEmail(email string) store.StoreChannel {
	return s.baseStore.GetByEmail(email)
}

func (s TransparentUserStore) GetByAuth(authData *string, authService string) store.StoreChannel {
	return s.baseStore.GetByAuth(authData, authService)
}

func (s TransparentUserStore) GetAllUsingAuthService(authService string) store.StoreChannel {
	return s.baseStore.GetAllUsingAuthService(authService)
}

func (s TransparentUserStore) GetByUsername(username string) store.StoreChannel {
	return s.baseStore.GetByUsername(username)
}

func (s TransparentUserStore) GetForLogin(loginId string, allowSignInWithUsername bool, allowSignInWithEmail bool) store.StoreChannel {
	return s.baseStore.GetForLogin(loginId, allowSignInWithUsername, allowSignInWithEmail)
}

func (s TransparentUserStore) VerifyEmail(userId string, email string) store.StoreChannel {
	return s.baseStore.VerifyEmail(userId, email)
}

func (s TransparentUserStore) GetEtagForAllProfiles() store.StoreChannel {
	return s.baseStore.GetEtagForAllProfiles()
}

func (s TransparentUserStore) GetEtagForProfiles(teamId string) store.StoreChannel {
	return s.baseStore.GetEtagForProfiles(teamId)
}

func (s TransparentUserStore) UpdateFailedPasswordAttempts(userId string, attempts int) store.StoreChannel {
	return s.baseStore.UpdateFailedPasswordAttempts(userId, attempts)
}

func (s TransparentUserStore) GetSystemAdminProfiles() store.StoreChannel {
	return s.baseStore.GetSystemAdminProfiles()
}

func (s TransparentUserStore) PermanentDelete(userId string) store.StoreChannel {
	return s.baseStore.PermanentDelete(userId)
}

func (s TransparentUserStore) AnalyticsActiveCount(time int64) store.StoreChannel {
	return s.baseStore.AnalyticsActiveCount(time)
}

func (s TransparentUserStore) GetUnreadCount(userId string) store.StoreChannel {
	return s.baseStore.GetUnreadCount(userId)
}

func (s TransparentUserStore) GetUnreadCountForChannel(userId string, channelId string) store.StoreChannel {
	return s.baseStore.GetUnreadCountForChannel(userId, channelId)
}

func (s TransparentUserStore) GetAnyUnreadPostCountForChannel(userId string, channelId string) store.StoreChannel {
	return s.baseStore.GetAnyUnreadPostCountForChannel(userId, channelId)
}

func (s TransparentUserStore) GetRecentlyActiveUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetRecentlyActiveUsersForTeam(teamId, offset, limit, viewRestrictions)
}

func (s TransparentUserStore) GetNewUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetNewUsersForTeam(teamId, offset, limit, viewRestrictions)
}

func (s TransparentUserStore) Search(teamId string, term string, options *model.UserSearchOptions) store.StoreChannel {
	return s.baseStore.Search(teamId, term, options)
}

func (s TransparentUserStore) SearchNotInTeam(notInTeamId string, term string, options *model.UserSearchOptions) store.StoreChannel {
	return s.baseStore.SearchNotInTeam(notInTeamId, term, options)
}

func (s TransparentUserStore) SearchInChannel(channelId string, term string, options *model.UserSearchOptions) store.StoreChannel {
	return s.baseStore.SearchInChannel(channelId, term, options)
}

func (s TransparentUserStore) SearchNotInChannel(teamId string, channelId string, term string, options *model.UserSearchOptions) store.StoreChannel {
	return s.baseStore.SearchNotInChannel(teamId, channelId, term, options)
}

func (s TransparentUserStore) SearchWithoutTeam(term string, options *model.UserSearchOptions) store.StoreChannel {
	return s.baseStore.SearchWithoutTeam(term, options)
}

func (s TransparentUserStore) AnalyticsGetInactiveUsersCount() store.StoreChannel {
	return s.baseStore.AnalyticsGetInactiveUsersCount()
}

func (s TransparentUserStore) AnalyticsGetSystemAdminCount() store.StoreChannel {
	return s.baseStore.AnalyticsGetSystemAdminCount()
}

func (s TransparentUserStore) GetProfilesNotInTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetProfilesNotInTeam(teamId, offset, limit, viewRestrictions)
}

func (s TransparentUserStore) GetEtagForProfilesNotInTeam(teamId string) store.StoreChannel {
	return s.baseStore.GetEtagForProfilesNotInTeam(teamId)
}

func (s TransparentUserStore) ClearAllCustomRoleAssignments() store.StoreChannel {
	return s.baseStore.ClearAllCustomRoleAssignments()
}

func (s TransparentUserStore) InferSystemInstallDate() store.StoreChannel {
	return s.baseStore.InferSystemInstallDate()
}

func (s TransparentUserStore) GetAllAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetAllAfter(limit, afterId)
}

func (s TransparentUserStore) GetUsersBatchForIndexing(startTime int64, endTime int64, limit int) store.StoreChannel {
	return s.baseStore.GetUsersBatchForIndexing(startTime, endTime, limit)
}

func (s TransparentUserStore) Count(options model.UserCountOptions) store.StoreChannel {
	return s.baseStore.Count(options)
}

func (s TransparentUserStore) GetTeamGroupUsers(teamID string) store.StoreChannel {
	return s.baseStore.GetTeamGroupUsers(teamID)
}

func (s TransparentUserStore) GetChannelGroupUsers(channelID string) store.StoreChannel {
	return s.baseStore.GetChannelGroupUsers(channelID)
}
