package cfbl

import (
	"bytes"
	"github.com/toorop/go-dkim"
	"net/mail"
	"net/textproto"
	"strings"
)

var cfblHeaderName = "CFBL-Address"

func CheckRequirements(mailBytes []byte) (bool, error) {
	msg, parseErr := mail.ReadMessage(bytes.NewReader(mailBytes))
	if parseErr != nil {
		return false, parseErr
	}

	// Is Complaint-FBL-Address header there ?
	if len(msg.Header[textproto.CanonicalMIMEHeaderKey(cfblHeaderName)]) == 0 {
		return false, ErrMissingCFBLAddressHeader
	}

	addrParts := strings.Split(msg.Header.Get(cfblHeaderName), ";")
	addr, err := mail.ParseAddress(addrParts[0])
	if err != nil {
		return false, err
	}

	var cfbldomain string
	t := strings.SplitAfter(addr.Address, "@")
	if len(t) > 1 {
		cfbldomain = strings.ToLower(t[1])
	}

	// Verify DKIM signature for CFBL domain
	if _, err := dkim.VerifyByDomain(&mailBytes, cfbldomain); err != nil {
		return false, err
	}

	// Has this signature the required tags
	return hasHeaderCoverageByDKIM(&mailBytes, &cfbldomain)
}

func hasHeaderCoverageByDKIM(mailBytes *[]byte, cfblDomain *string) (bool, error) {
	dkHeader, err := dkim.GetHeaderForDomain(mailBytes, *cfblDomain)
	if err != nil {
		return false, err
	}

	found := false
	for _, header := range dkHeader.Headers {
		if strings.ToLower(header) == strings.ToLower(cfblHeaderName) {
			found = true
			break
		}
	}
	if !found {
		return false, ErrDKIMSigMissingHeader
	}
	return true, nil
}
