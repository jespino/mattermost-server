// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package rediscachestore

import (
	"testing"

	"github.com/mattermost/mattermost-server/store/storetest"
)

func TestAuditStore(t *testing.T) {
	StoreTest(t, storetest.TestAuditStore)
}

func TestBotStore(t *testing.T) {
	StoreTest(t, storetest.TestBotStore)
}

func TestChannelMemberHistoryStore(t *testing.T) {
	StoreTest(t, storetest.TestChannelMemberHistoryStore)
}

func TestChannelStore(t *testing.T) {
	StoreTestWithSqlSupplier(t, storetest.TestChannelStore)
}

func TestClusterDiscoveryStore(t *testing.T) {
	StoreTest(t, storetest.TestClusterDiscoveryStore)
}

func TestCommandStore(t *testing.T) {
	StoreTest(t, storetest.TestCommandStore)
}

func TestCommandWebhookStore(t *testing.T) {
	StoreTest(t, storetest.TestCommandWebhookStore)
}

func TestComplianceStore(t *testing.T) {
	StoreTest(t, storetest.TestComplianceStore)
}

func TestEmojiStore(t *testing.T) {
	StoreTest(t, storetest.TestEmojiStore)
}

func TestFileInfoStore(t *testing.T) {
	StoreTest(t, storetest.TestFileInfoStore)
}

func TestGroupStore(t *testing.T) {
	StoreTest(t, storetest.TestGroupStore)
}

func TestJobStore(t *testing.T) {
	StoreTest(t, storetest.TestJobStore)
}

func TestLicenseStore(t *testing.T) {
	StoreTest(t, storetest.TestLicenseStore)
}

func TestLinkMetadataStore(t *testing.T) {
	StoreTest(t, storetest.TestLinkMetadataStore)
}

func TestNotificationRegistryStore(t *testing.T) {
	StoreTest(t, storetest.TestNotificationRegistryStore)
}

func TestOAuthStore(t *testing.T) {
	StoreTest(t, storetest.TestOAuthStore)
}

func TestPluginStore(t *testing.T) {
	StoreTest(t, storetest.TestPluginStore)
}

func TestPostStore(t *testing.T) {
	StoreTestWithSqlSupplier(t, storetest.TestPostStore)
}

func TestPreferenceStore(t *testing.T) {
	StoreTest(t, storetest.TestPreferenceStore)
}

func TestReactionStore(t *testing.T) {
	StoreTest(t, storetest.TestReactionStore)
}

func TestRoleStore(t *testing.T) {
	StoreTest(t, storetest.TestRoleStore)
}

func TestSchemeStore(t *testing.T) {
	StoreTest(t, storetest.TestSchemeStore)
}

func TestSessionStore(t *testing.T) {
	StoreTest(t, storetest.TestSessionStore)
}

func TestStatusStore(t *testing.T) {
	StoreTest(t, storetest.TestStatusStore)
}

func TestSystemStore(t *testing.T) {
	StoreTest(t, storetest.TestSystemStore)
}

func TestTeamStore(t *testing.T) {
	StoreTest(t, storetest.TestTeamStore)
}

func TestTermsOfServiceStore(t *testing.T) {
	StoreTest(t, storetest.TestTermsOfServiceStore)
}

func TestUserAccessTokenStore(t *testing.T) {
	StoreTest(t, storetest.TestUserAccessTokenStore)
}

func TestUserStore(t *testing.T) {
	StoreTest(t, storetest.TestUserStore)
}

func TestUserTermsOfServiceStore(t *testing.T) {
	StoreTest(t, storetest.TestUserTermsOfServiceStore)
}

func TestWebhookStore(t *testing.T) {
	StoreTest(t, storetest.TestWebhookStore)
}
