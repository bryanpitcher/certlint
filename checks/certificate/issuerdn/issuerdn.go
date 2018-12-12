package version

import (
	"bytes"

	"github.com/bryanpitcher/certlint/certdata"
	"github.com/bryanpitcher/certlint/checks"
	"github.com/bryanpitcher/certlint/errors"
)

const checkName = "Issuer DN Check"

func init() {
	checks.RegisterCertificateCheck(checkName, nil, Check)
}

// Check performs a strict verification on the extension according to the standard(s)
func Check(d *certdata.Data) *errors.Errors {
	var e = errors.New(nil)

	if d.Issuer != nil && !bytes.Equal(d.Cert.RawIssuer, d.Issuer.RawSubject) {
		e.Err("Certificate Issuer Distinguished Name field MUST match the Subject DN of the Issuing CA")
		return e
	}

	return e
}
