package events

import "github.com/mattermost/mattermost-server/v6/services/systembus"

var StartUp = systembus.EventDefinition{
	ID:          "startup",
	Name:        "Startup event",
	Description: "Event executed on startup",
	Fields:      []string{},
}

var ShutDown = systembus.EventDefinition{
	ID:          "shutdown",
	Name:        "Shutdown event",
	Description: "Event executed at the beginning of the shutdown",
	Fields:      []string{},
}

var ChannelCreated = systembus.EventDefinition{
	ID:          "channel-created",
	Name:        "Channel created",
	Description: "This events gets triggered when a new channel is created.",
	Fields:      []string{"Id", "Name", "DisplayName", "CreatorId", "Type", "TeamId"},
}

var ChannelArchived = systembus.EventDefinition{
	ID:          "channel-archived",
	Name:        "Channel archived",
	Description: "This events gets triggered when a channel is archived.",
	Fields:      []string{"ChannelId", "UserId"},
}

var ChannelUnarchived = systembus.EventDefinition{
	ID:          "channel-unarchived",
	Name:        "Channel unarchived",
	Description: "This events gets triggered when a channel is unarchived.",
	Fields:      []string{"ChannelId", "UserId"},
}

var TeamCreated = systembus.EventDefinition{
	ID:          "team-created",
	Name:        "Team created",
	Description: "This events gets triggered when a new team is created.",
	Fields:      []string{"Id", "Name", "DisplayName", "CreatorId"},
}

var TeamArchived = systembus.EventDefinition{
	ID:          "channel-archived",
	Name:        "Channel archived",
	Description: "This events gets triggered when a team is archived.",
	Fields:      []string{"TeamId", "UserId"},
}

var TeamUnarchived = systembus.EventDefinition{
	ID:          "team-unarchived",
	Name:        "Team unarchived",
	Description: "This events gets triggered when a team is unarchived.",
	Fields:      []string{"TeamId", "UserId"},
}

var PostReactionAdded = systembus.EventDefinition{
	ID:          "post-reaction-added",
	Name:        "Post reaction added",
	Description: "This events gets triggered when a reaction is added to a post.",
	Fields:      []string{"PostId", "UserId", "Emoji"},
}

var PostReactionRemoved = systembus.EventDefinition{
	ID:          "post-reaction-removed",
	Name:        "Post reaction removed",
	Description: "This events gets triggered when a reaction is removed from a post.",
	Fields:      []string{"PostId", "UserId", "Emoji"},
}

var PostCreated = systembus.EventDefinition{
	ID:          "post-created",
	Name:        "Post created",
	Description: "This events gets triggered when a new post is created.",
	Fields:      []string{"Id", "Message", "UserId", "ChannelId", "RootId", "TeamId"},
}

var PostDeleted = systembus.EventDefinition{
	ID:          "post-deleted",
	Name:        "Post deleted",
	Description: "This events gets triggered when a post is deleted.",
	Fields:      []string{"Message", "PostId", "UserId", "ChannelId", "TeamId"},
}

var UserJoinChannel = systembus.EventDefinition{
	ID:          "user-join-channel",
	Name:        "User join channel",
	Description: "This events gets triggered when a user join a channel.",
	Fields:      []string{"UserId", "ChannelId"},
}

var UserLeaveChannel = systembus.EventDefinition{
	ID:          "user-leave-channel",
	Name:        "User leave channel",
	Description: "This events gets triggered when a user leave a channel.",
	Fields:      []string{"UserId", "ChannelId"},
}

var UserJoinTeam = systembus.EventDefinition{
	ID:          "user-join-team",
	Name:        "User join team",
	Description: "This events gets triggered when a user join a team.",
	Fields:      []string{"UserId", "TeamId"},
}

var UserLeaveTeam = systembus.EventDefinition{
	ID:          "user-leave-team",
	Name:        "User leave team",
	Description: "This events gets triggered when a user leave a team.",
	Fields:      []string{"UserId", "TeamId"},
}

var SlashCommandCalled = systembus.EventDefinition{
	ID:          "slash-command-called",
	Name:        "Slash command called",
	Description: "This events gets triggered when a user call a server side slash command.",
	Fields:      []string{"UserId", "TeamId", "ChannelId", "SlashCommand"},
}

var LoginSuccess = systembus.EventDefinition{
	ID:          "login-success",
	Name:        "Login success",
	Description: "This events gets triggered when a user succeed to login",
	Fields:      []string{"UserId", "UserAgent", "IPAddress"},
}

var LoginFailure = systembus.EventDefinition{
	ID:          "login-failure",
	Name:        "Login failure",
	Description: "This events gets triggered when a user failed to login.",
	Fields:      []string{"UsernameInput"},
}

var NewUser = systembus.EventDefinition{
	ID:          "new-user",
	Name:        "New User",
	Description: "This events gets triggered when a new user is created.",
	Fields:      []string{"UserId", "Source"},
}
