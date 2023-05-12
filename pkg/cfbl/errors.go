package cfbl

import "errors"

var (
	ErrMissingCFBLAddressHeader = errors.New("no CFBL-Address header field found")

	ErrDKIMVerificationFailed = errors.New("DKIM verification failed")

	ErrCFBLDomainNotAligned = errors.New("CFBL Address domain not aligned with DKIM signature")

	ErrDKIMSigMissingHeader = errors.New("DKIM signature doesn't contain CFBL-Address header")
)
