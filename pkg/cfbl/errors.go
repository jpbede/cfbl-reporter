package cfbl

import "errors"

var (
	ErrMissingComplaintFBLAddressHeader = errors.New("no Complaint-FBL-Address header field found")

	ErrDKIMVerificationFailed = errors.New("DKIM verification failed")

	ErrCFBLDomainNotAligned = errors.New("CFBL Address domain not aligned with DKIM signature")
)
