module go.bnck.me/cfbl-reporter

go 1.20

require (
	github.com/go-mail/mail v2.3.1+incompatible
	github.com/stretchr/testify v1.8.3
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208
	github.com/urfave/cli/v2 v2.25.3
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/go-mail/mail v2.3.1+incompatible => github.com/jpbede/mail v0.0.0-20210708152751-3f6da5e71cfa
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208 => github.com/jpbede/go-dkim v0.0.0-20210706184804-5069ad4f6b15
)
