package mailer

import "embed"

const (
	FromName            = "Rat"
	maxRetries          = 3
	UserWelcomeTemplate = "user_invitation.templ"
)

var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) (int, error)
}
