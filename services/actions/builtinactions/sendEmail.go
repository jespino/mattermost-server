package builtinactions

import (
	"github.com/mattermost/mattermost-server/v6/services/actions"
	"github.com/mattermost/mattermost-server/v6/services/configservice"
	"github.com/mattermost/mattermost-server/v6/shared/mail"
	"github.com/mattermost/mattermost-server/v6/utils"
)

const SendEmailID = "send-email"

func NewSendEmail(cfg configservice.ConfigService) *actions.ActionDefinition {
	handler := func(data map[string]string) (map[string]string, error) {
		emailSettings := cfg.Config().EmailSettings
		hostname := utils.GetHostnameFromSiteURL(*cfg.Config().ServiceSettings.SiteURL)

		cfg := mail.SMTPConfig{
			Hostname:                          hostname,
			ConnectionSecurity:                *emailSettings.ConnectionSecurity,
			SkipServerCertificateVerification: *emailSettings.SkipServerCertificateVerification,
			ServerName:                        *emailSettings.SMTPServer,
			Server:                            *emailSettings.SMTPServer,
			Port:                              *emailSettings.SMTPPort,
			ServerTimeout:                     *emailSettings.SMTPServerTimeout,
			Username:                          *emailSettings.SMTPUsername,
			Password:                          *emailSettings.SMTPPassword,
			EnableSMTPAuth:                    *emailSettings.EnableSMTPAuth,
			SendEmailNotifications:            *emailSettings.SendEmailNotifications,
			FeedbackName:                      *emailSettings.FeedbackName,
			FeedbackEmail:                     *emailSettings.FeedbackEmail,
			ReplyToAddress:                    *emailSettings.ReplyToAddress,
		}

		err := mail.SendMailUsingConfig(data["to"], data["subject"], data["body"], &cfg, false, "")
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return &actions.ActionDefinition{
		ID:               SendEmailID,
		Name:             "Send Email",
		Description:      "Send an email to an address",
		ConfigDefinition: map[string]string{"to": "string", "subject": "string", "body": "string"},
		Handler:          handler,
	}
}
