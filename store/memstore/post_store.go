// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemPostStore struct {
	MemStore *MemStore
	posts    []*model.Post
}

func newMemPostStore(memStore *MemStore) store.PostStore {
	return &MemPostStore{MemStore: memStore}
}

func (s *MemPostStore) ClearCaches() {}
func (s *MemPostStore) SaveMultiple(posts []*model.Post) ([]*model.Post, int, error) {
	channelNewPosts := make(map[string]int)
	channelNewRootPosts := make(map[string]int)
	maxDateNewPosts := make(map[string]int64)
	maxDateNewRootPosts := make(map[string]int64)
	rootIds := make(map[string]int)
	maxDateRootIds := make(map[string]int64)
	for idx, post := range posts {
		if post.Id != "" && !post.IsRemote() {
			return nil, idx, store.NewErrInvalidInput("Post", "id", post.Id)
		}
		post.PreSave()
		maxPostSize := s.GetMaxPostSize()
		if err := post.IsValid(maxPostSize); err != nil {
			return nil, idx, err
		}

		if currentChannelCount, ok := channelNewPosts[post.ChannelId]; !ok {
			if post.IsJoinLeaveMessage() {
				channelNewPosts[post.ChannelId] = 0
			} else {
				channelNewPosts[post.ChannelId] = 1
			}
			maxDateNewPosts[post.ChannelId] = post.CreateAt
		} else {
			if !post.IsJoinLeaveMessage() {
				channelNewPosts[post.ChannelId] = currentChannelCount + 1
			}
			if post.CreateAt > maxDateNewPosts[post.ChannelId] {
				maxDateNewPosts[post.ChannelId] = post.CreateAt
			}
		}

		if post.RootId == "" {
			if currentChannelCount, ok := channelNewRootPosts[post.ChannelId]; !ok {
				if post.IsJoinLeaveMessage() {
					channelNewRootPosts[post.ChannelId] = 0
				} else {
					channelNewRootPosts[post.ChannelId] = 1
				}
				maxDateNewRootPosts[post.ChannelId] = post.CreateAt
			} else {
				if !post.IsJoinLeaveMessage() {
					channelNewRootPosts[post.ChannelId] = currentChannelCount + 1
				}
				if post.CreateAt > maxDateNewRootPosts[post.ChannelId] {
					maxDateNewRootPosts[post.ChannelId] = post.CreateAt
				}
			}
			continue
		}

		if currentRootCount, ok := rootIds[post.RootId]; !ok {
			rootIds[post.RootId] = 1
			maxDateRootIds[post.RootId] = post.CreateAt
		} else {
			rootIds[post.RootId] = currentRootCount + 1
			if post.CreateAt > maxDateRootIds[post.RootId] {
				maxDateRootIds[post.RootId] = post.CreateAt
			}
		}
	}

	s.posts = append(s.posts, posts...)

	return posts, -1, nil
}

func (s *MemPostStore) Save(post *model.Post) (*model.Post, error) {
	posts, _, err := s.SaveMultiple([]*model.Post{post})
	if err != nil {
		return nil, err
	}
	return posts[0], nil
}

func (s *MemPostStore) populateReplyCount(posts []*model.Post) error {
	panic("not implemented")
}

func (s *MemPostStore) Update(newPost *model.Post, oldPost *model.Post) (*model.Post, error) {
	panic("not implemented")
}

func (s *MemPostStore) OverwriteMultiple(posts []*model.Post) ([]*model.Post, int, error) {
	panic("not implemented")
}

func (s *MemPostStore) Overwrite(post *model.Post) (*model.Post, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetFlaggedPosts(userId string, offset int, limit int) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetFlaggedPostsForTeam(userId, teamId string, offset int, limit int) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetFlaggedPostsForChannel(userId, channelId string, offset int, limit int) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) getFlaggedPosts(userId, channelId, teamId string, offset int, limit int) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) buildFlaggedPostTeamFilterClause(teamId string, queryParams []interface{}) (string, []interface{}) {
	panic("not implemented")
}

func (s *MemPostStore) buildFlaggedPostChannelFilterClause(channelId string, queryParams []interface{}) (string, []interface{}) {
	panic("not implemented")
}

func (s *MemPostStore) getPostWithCollapsedThreads(id, userID string, extended bool) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) Get(ctx context.Context, id string, skipFetchThreads, collapsedThreads, collapsedThreadsExtended bool, userID string) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetSingle(id string, inclDeleted bool) (*model.Post, error) {
	panic("not implemented")
}

func (s *MemPostStore) InvalidateLastPostTimeCache(channelId string) {}

func (s *MemPostStore) GetEtag(channelId string, allowFromCache, collapsedThreads bool) string {
	panic("not implemented")
}

func (s *MemPostStore) Delete(postID string, time int64, deleteByID string) error {
	panic("not implemented")
}

func (s *MemPostStore) PermanentDeleteByUser(userId string) error {
	result := []*model.Post{}
	for _, p := range s.posts {
		if p.UserId != userId {
			result = append(result, p)
		}
	}
	s.posts = result
	return nil
}

func (s *MemPostStore) PermanentDeleteByChannel(channelId string) error {
	result := []*model.Post{}
	for _, p := range s.posts {
		if p.ChannelId != channelId {
			result = append(result, p)
		}
	}
	s.posts = result
	return nil
}

func (s *MemPostStore) GetPosts(options model.GetPostsOptions, _ bool) (*model.PostList, error) {
	if options.PerPage > 1000 {
		return nil, store.NewErrInvalidInput("Post", "<options.PerPage>", options.PerPage)
	}
	offset := options.PerPage * options.Page
	list := model.NewPostList()

	counter := 0
	posts := []*model.Post{}
	for _, p := range s.posts {
		if p.ChannelId == options.ChannelId && p.DeleteAt == 0 {
			if counter >= offset && counter < offset+options.PerPage {
				posts = append(posts, p)
			}
			counter++
		}
	}

	for _, p := range posts {
		list.AddPost(p)
		list.AddOrder(p.Id)
	}

	for _, p := range posts {
		for _, candidate := range s.posts {
			if p.RootId == candidate.Id {
				list.AddPost(candidate)
				break
			}
		}
	}

	list.MakeNonNil()
	fmt.Println(list)

	return list, nil
}

func (s *MemPostStore) GetPostsSince(options model.GetPostsSinceOptions, allowFromCache bool) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) HasAutoResponsePostByUserSince(options model.GetPostsSinceOptions, userId string) (bool, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetPostsSinceForSync(options model.GetPostsSinceForSyncOptions, cursor model.GetPostsSinceForSyncCursor, limit int) ([]*model.Post, model.GetPostsSinceForSyncCursor, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetPostsBefore(options model.GetPostsOptions) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetPostsAfter(options model.GetPostsOptions) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) getPostsAround(before bool, options model.GetPostsOptions) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetPostIdBeforeTime(channelId string, time int64, collapsedThreads bool) (string, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetPostIdAfterTime(channelId string, time int64, collapsedThreads bool) (string, error) {
	panic("not implemented")
}

func (s *MemPostStore) getPostIdAroundTime(channelId string, time int64, before bool, collapsedThreads bool) (string, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetPostAfterTime(channelId string, time int64, collapsedThreads bool) (*model.Post, error) {
	panic("not implemented")
}

func (s *MemPostStore) getRootPosts(channelId string, offset int, limit int, skipFetchThreads bool) ([]*model.Post, error) {
	panic("not implemented")
}

func (s *MemPostStore) getParentsPosts(channelId string, offset int, limit int, skipFetchThreads bool) ([]*model.Post, error) {
	panic("not implemented")
}

func (s *MemPostStore) getParentsPostsPostgreSQL(channelId string, offset int, limit int, skipFetchThreads bool) ([]*model.Post, error) {
	panic("not implemented")
}

func (s *MemPostStore) buildCreateDateFilterClause(params *model.SearchParams, queryParams map[string]interface{}, builder sq.SelectBuilder) (sq.SelectBuilder, map[string]interface{}) {
	panic("not implemented")
}

func (s *MemPostStore) buildSearchTeamFilterClause(teamId string, queryParams map[string]interface{}, builder sq.SelectBuilder) (sq.SelectBuilder, map[string]interface{}) {
	panic("not implemented")
}

func (s *MemPostStore) buildSearchChannelFilterClause(channels []string, paramPrefix string, exclusion bool, queryParams map[string]interface{}, byName bool, builder sq.SelectBuilder) (sq.SelectBuilder, map[string]interface{}) {
	panic("not implemented")
}

func (s *MemPostStore) buildSearchUserFilterClause(users []string, paramPrefix string, exclusion bool, queryParams map[string]interface{}, byUsername bool) (string, map[string]interface{}) {
	panic("not implemented")
}

func (s *MemPostStore) buildSearchPostFilterClause(fromUsers []string, excludedUsers []string, queryParams map[string]interface{}, userByUsername bool, builder sq.SelectBuilder) (sq.SelectBuilder, map[string]interface{}) {
	panic("not implemented")
}

func (s *MemPostStore) Search(teamId string, userId string, params *model.SearchParams) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) search(teamId string, userId string, params *model.SearchParams, channelsByName bool, userByUsername bool) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemPostStore) AnalyticsUserCountsWithPostsByDay(teamId string) (model.AnalyticsRows, error) {
	panic("not implemented")
}

func (s *MemPostStore) AnalyticsPostCountsByDay(options *model.AnalyticsPostCountsOptions) (model.AnalyticsRows, error) {
	panic("not implemented")
}

func (s *MemPostStore) AnalyticsPostCount(teamId string, mustHaveFile bool, mustHaveHashtag bool) (int64, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetPostsCreatedAt(channelId string, time int64) ([]*model.Post, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetPostsByIds(postIds []string) ([]*model.Post, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetPostsBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.PostForIndexing, error) {
	panic("not implemented")
}

func (s *MemPostStore) PermanentDeleteBatchForRetentionPolicies(now, globalPolicyEndTime, limit int64, cursor model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error) {
	panic("not implemented")
}

func (s *MemPostStore) DeleteOrphanedRows(limit int) (deleted int64, err error) {
	panic("not implemented")
}

func (s *MemPostStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetOldest() (*model.Post, error) {
	panic("not implemented")
}

func (s *MemPostStore) determineMaxPostSize() int {
	panic("not implemented")
}

func (s *MemPostStore) GetMaxPostSize() int {
	return 10000000
}

func (s *MemPostStore) GetParentsForExportAfter(limit int, afterId string) ([]*model.PostForExport, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetRepliesForExport(rootId string) ([]*model.ReplyForExport, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetDirectPostParentsForExportAfter(limit int, afterId string) ([]*model.DirectPostForExport, error) {
	panic("not implemented")
}

func (s *MemPostStore) SearchPostsForUser(paramsList []*model.SearchParams, userId, teamId string, page, perPage int) (*model.PostSearchResults, error) {
	panic("not implemented")
}

func (s *MemPostStore) GetOldestEntityCreationTime() (int64, error) {
	panic("not implemented")
}
