package mailer

import "embed"

const (
	FromName            = "Grief"
	maxRetries          = 3
	UserWelcomeTemplate = "user_invitation.tmpl"
)

var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) (int, error)
}
