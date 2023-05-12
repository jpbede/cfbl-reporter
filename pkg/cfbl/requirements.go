package cfbl

import (
	"bytes"
	"github.com/toorop/go-dkim"
	"net/mail"
	"net/textproto"
	"strings"
)

func (r *Report) CheckRequirements() (bool, error) {
	msg, parseErr := mail.ReadMessage(bytes.NewReader(*r.originalMail))
	if parseErr != nil {
		return false, parseErr
	}

	// Is CFBL-Address header there ?
	if len(msg.Header[textproto.CanonicalMIMEHeaderKey(cfblHeaderName)]) == 0 {
		return false, ErrMissingCFBLAddressHeader
	}

	cfblParts := strings.Split(msg.Header.Get(cfblHeaderName), ";")
	addr, err := mail.ParseAddress(cfblParts[0])
	if err != nil {
		return false, err
	}

	var cfblDomain string
	addrParts := strings.SplitAfter(addr.Address, "@")
	if len(addrParts) > 1 {
		cfblDomain = strings.ToLower(addrParts[1])
	}

	// Verify DKIM signature for CFBL domain
	if _, err := dkim.VerifyByDomain(r.originalMail, cfblDomain); err != nil {
		return false, err
	}

	// Has this signature the required tags
	return hasHeaderCoverageByDKIM(r.originalMail, &cfblDomain)
}

func hasHeaderCoverageByDKIM(mailBytes *[]byte, cfblDomain *string) (bool, error) {
	dkHeader, err := dkim.GetHeaderForDomain(mailBytes, *cfblDomain)
	if err != nil {
		return false, err
	}

	for _, header := range dkHeader.Headers {
		if strings.ToLower(header) == strings.ToLower(cfblHeaderName) {
			return true, nil
		}
	}
	return false, ErrDKIMSigMissingHeader
}
