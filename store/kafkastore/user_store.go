package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaUserStore struct {
	baseStore store.UserStore
	root      *KafkaStore
}

func (s KafkaUserStore) Save(user *model.User) store.StoreChannel {
	return s.baseStore.Save(user)
}

func (s KafkaUserStore) Update(user *model.User, allowRoleUpdate bool) store.StoreChannel {
	return s.baseStore.Update(user, allowRoleUpdate)
}

func (s KafkaUserStore) UpdateLastPictureUpdate(userId string) store.StoreChannel {
	return s.baseStore.UpdateLastPictureUpdate(userId)
}

func (s KafkaUserStore) ResetLastPictureUpdate(userId string) store.StoreChannel {
	return s.baseStore.ResetLastPictureUpdate(userId)
}

func (s KafkaUserStore) UpdateUpdateAt(userId string) store.StoreChannel {
	return s.baseStore.UpdateUpdateAt(userId)
}

func (s KafkaUserStore) UpdatePassword(userId string, newPassword string) store.StoreChannel {
	return s.baseStore.UpdatePassword(userId, newPassword)
}

func (s KafkaUserStore) UpdateAuthData(userId string, service string, authData *string, email string, resetMfa bool) store.StoreChannel {
	return s.baseStore.UpdateAuthData(userId, service, authData, email, resetMfa)
}

func (s KafkaUserStore) UpdateMfaSecret(userId string, secret string) store.StoreChannel {
	return s.baseStore.UpdateMfaSecret(userId, secret)
}

func (s KafkaUserStore) UpdateMfaActive(userId string, active bool) store.StoreChannel {
	return s.baseStore.UpdateMfaActive(userId, active)
}

func (s KafkaUserStore) Get(id string) (*model.User, *model.AppError) {
	return s.baseStore.Get(id)
}

func (s KafkaUserStore) GetAll() store.StoreChannel {
	return s.baseStore.GetAll()
}

func (s KafkaUserStore) ClearCaches() {
	s.baseStore.ClearCaches()
}

func (s KafkaUserStore) InvalidateProfilesInChannelCacheByUser(userId string) {
	s.baseStore.InvalidateProfilesInChannelCacheByUser(userId)
}

func (s KafkaUserStore) InvalidateProfilesInChannelCache(channelId string) {
	s.baseStore.InvalidateProfilesInChannelCache(channelId)
}

func (s KafkaUserStore) GetProfilesInChannel(channelId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetProfilesInChannel(channelId, offset, limit)
}

func (s KafkaUserStore) GetProfilesInChannelByStatus(channelId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetProfilesInChannelByStatus(channelId, offset, limit)
}

func (s KafkaUserStore) GetAllProfilesInChannel(channelId string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetAllProfilesInChannel(channelId, allowFromCache)
}

func (s KafkaUserStore) GetProfilesNotInChannel(teamId string, channelId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetProfilesNotInChannel(teamId, channelId, offset, limit, viewRestrictions)
}

func (s KafkaUserStore) GetProfilesWithoutTeam(offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetProfilesWithoutTeam(offset, limit, viewRestrictions)
}

func (s KafkaUserStore) GetProfilesByUsernames(usernames []string, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetProfilesByUsernames(usernames, viewRestrictions)
}

func (s KafkaUserStore) GetAllProfiles(options *model.UserGetOptions) store.StoreChannel {
	return s.baseStore.GetAllProfiles(options)
}

func (s KafkaUserStore) GetProfiles(options *model.UserGetOptions) store.StoreChannel {
	return s.baseStore.GetProfiles(options)
}

func (s KafkaUserStore) GetProfileByIds(userId []string, allowFromCache bool, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetProfileByIds(userId, allowFromCache, viewRestrictions)
}

func (s KafkaUserStore) InvalidatProfileCacheForUser(userId string) {
	s.baseStore.InvalidatProfileCacheForUser(userId)
}

func (s KafkaUserStore) GetByEmail(email string) store.StoreChannel {
	return s.baseStore.GetByEmail(email)
}

func (s KafkaUserStore) GetByAuth(authData *string, authService string) store.StoreChannel {
	return s.baseStore.GetByAuth(authData, authService)
}

func (s KafkaUserStore) GetAllUsingAuthService(authService string) store.StoreChannel {
	return s.baseStore.GetAllUsingAuthService(authService)
}

func (s KafkaUserStore) GetByUsername(username string) store.StoreChannel {
	return s.baseStore.GetByUsername(username)
}

func (s KafkaUserStore) GetForLogin(loginId string, allowSignInWithUsername bool, allowSignInWithEmail bool) store.StoreChannel {
	return s.baseStore.GetForLogin(loginId, allowSignInWithUsername, allowSignInWithEmail)
}

func (s KafkaUserStore) VerifyEmail(userId string, email string) store.StoreChannel {
	return s.baseStore.VerifyEmail(userId, email)
}

func (s KafkaUserStore) GetEtagForAllProfiles() store.StoreChannel {
	return s.baseStore.GetEtagForAllProfiles()
}

func (s KafkaUserStore) GetEtagForProfiles(teamId string) store.StoreChannel {
	return s.baseStore.GetEtagForProfiles(teamId)
}

func (s KafkaUserStore) UpdateFailedPasswordAttempts(userId string, attempts int) store.StoreChannel {
	return s.baseStore.UpdateFailedPasswordAttempts(userId, attempts)
}

func (s KafkaUserStore) GetSystemAdminProfiles() store.StoreChannel {
	return s.baseStore.GetSystemAdminProfiles()
}

func (s KafkaUserStore) PermanentDelete(userId string) store.StoreChannel {
	return s.baseStore.PermanentDelete(userId)
}

func (s KafkaUserStore) AnalyticsActiveCount(time int64) store.StoreChannel {
	return s.baseStore.AnalyticsActiveCount(time)
}

func (s KafkaUserStore) GetUnreadCount(userId string) store.StoreChannel {
	return s.baseStore.GetUnreadCount(userId)
}

func (s KafkaUserStore) GetUnreadCountForChannel(userId string, channelId string) store.StoreChannel {
	return s.baseStore.GetUnreadCountForChannel(userId, channelId)
}

func (s KafkaUserStore) GetAnyUnreadPostCountForChannel(userId string, channelId string) store.StoreChannel {
	return s.baseStore.GetAnyUnreadPostCountForChannel(userId, channelId)
}

func (s KafkaUserStore) GetRecentlyActiveUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetRecentlyActiveUsersForTeam(teamId, offset, limit, viewRestrictions)
}

func (s KafkaUserStore) GetNewUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetNewUsersForTeam(teamId, offset, limit, viewRestrictions)
}

func (s KafkaUserStore) Search(teamId string, term string, options *model.UserSearchOptions) store.StoreChannel {
	return s.baseStore.Search(teamId, term, options)
}

func (s KafkaUserStore) SearchNotInTeam(notInTeamId string, term string, options *model.UserSearchOptions) store.StoreChannel {
	return s.baseStore.SearchNotInTeam(notInTeamId, term, options)
}

func (s KafkaUserStore) SearchInChannel(channelId string, term string, options *model.UserSearchOptions) store.StoreChannel {
	return s.baseStore.SearchInChannel(channelId, term, options)
}

func (s KafkaUserStore) SearchNotInChannel(teamId string, channelId string, term string, options *model.UserSearchOptions) store.StoreChannel {
	return s.baseStore.SearchNotInChannel(teamId, channelId, term, options)
}

func (s KafkaUserStore) SearchWithoutTeam(term string, options *model.UserSearchOptions) store.StoreChannel {
	return s.baseStore.SearchWithoutTeam(term, options)
}

func (s KafkaUserStore) AnalyticsGetInactiveUsersCount() store.StoreChannel {
	return s.baseStore.AnalyticsGetInactiveUsersCount()
}

func (s KafkaUserStore) AnalyticsGetSystemAdminCount() store.StoreChannel {
	return s.baseStore.AnalyticsGetSystemAdminCount()
}

func (s KafkaUserStore) GetProfilesNotInTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetProfilesNotInTeam(teamId, offset, limit, viewRestrictions)
}

func (s KafkaUserStore) GetEtagForProfilesNotInTeam(teamId string) store.StoreChannel {
	return s.baseStore.GetEtagForProfilesNotInTeam(teamId)
}

func (s KafkaUserStore) ClearAllCustomRoleAssignments() store.StoreChannel {
	return s.baseStore.ClearAllCustomRoleAssignments()
}

func (s KafkaUserStore) InferSystemInstallDate() store.StoreChannel {
	return s.baseStore.InferSystemInstallDate()
}

func (s KafkaUserStore) GetAllAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetAllAfter(limit, afterId)
}

func (s KafkaUserStore) GetUsersBatchForIndexing(startTime int64, endTime int64, limit int) store.StoreChannel {
	return s.baseStore.GetUsersBatchForIndexing(startTime, endTime, limit)
}

func (s KafkaUserStore) Count(options model.UserCountOptions) store.StoreChannel {
	return s.baseStore.Count(options)
}

func (s KafkaUserStore) GetTeamGroupUsers(teamID string) store.StoreChannel {
	return s.baseStore.GetTeamGroupUsers(teamID)
}

func (s KafkaUserStore) GetChannelGroupUsers(channelID string) store.StoreChannel {
	return s.baseStore.GetChannelGroupUsers(channelID)
}
