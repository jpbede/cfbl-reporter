package cfbl

import (
	"bytes"
	"github.com/davecgh/go-spew/spew"
	"github.com/toorop/go-dkim"
	"net/mail"
	"net/textproto"
	"strings"
)

func CheckRequirements(mailBytes []byte) (bool, error) {
	msg, parseErr := mail.ReadMessage(bytes.NewReader(mailBytes))
	if parseErr != nil {
		return false, parseErr
	}

	// Is Complaint-FBL-Address header there ?
	if len(msg.Header[textproto.CanonicalMIMEHeaderKey("Complaint-FBL-Address")]) == 0 {
		return false, ErrMissingComplaintFBLAddressHeader
	}

	// Has CFBL Address domain a DKIM header ?
	addrParts := strings.Split(msg.Header.Get("Complaint-FBL-Address"), ";")
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

	dkHeader, err := dkim.GetHeaderForDomain(&mailBytes, cfbldomain)
	if err != nil {
		return false, err
	} else if dkHeader == nil {
		return false, ErrCFBLDomainNotAligned
	}

	spew.Dump(dkHeader)

	return true, nil
}
