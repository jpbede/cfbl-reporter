package cfbl

import "github.com/go-mail/mail"

type ComposerOption func(msg *mail.Message)

func WithFrom(name, email string) ComposerOption {
	return func(msg *mail.Message) {
		msg.SetAddressHeader("From", email, name)
	}
}
