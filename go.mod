module go.bnck.me/cfbl-reporter

go 1.16

require (
	github.com/go-mail/mail v2.3.1+incompatible
	github.com/stretchr/testify v1.8.2
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208
	github.com/urfave/cli/v2 v2.25.3
)

replace (
	github.com/go-mail/mail v2.3.1+incompatible => github.com/jpbede/mail v0.0.0-20210708152751-3f6da5e71cfa
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208 => github.com/jpbede/go-dkim v0.0.0-20210706184804-5069ad4f6b15
)
