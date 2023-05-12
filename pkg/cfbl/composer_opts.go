package cfbl

import "github.com/go-mail/mail"

type ComposerOption func(msg *mail.Message, report *Report)

func WithFrom(name, email string) ComposerOption {
	return func(msg *mail.Message, _ *Report) {
		msg.SetAddressHeader("From", email, name)
	}
}

func WithFullEmailInReport() ComposerOption {
	return func(_ *mail.Message, report *Report) {
		report.sendFullReport = true
	}
}
