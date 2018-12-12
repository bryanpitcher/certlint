package keyusage

import (
	"crypto/x509/pkix"
	"encoding/asn1"

	"github.com/bryanpitcher/certlint/certdata"
	"github.com/bryanpitcher/certlint/checks"
	"github.com/bryanpitcher/certlint/errors"
)

const checkName = "KeyUsage Extension Check"

var extensionOid = asn1.ObjectIdentifier{2, 5, 29, 15}

func init() {
	checks.RegisterExtensionCheck(checkName, extensionOid, nil, Check)
}

// Check performs a strict verification on the extension according to the standard(s)
//
// https://tools.ietf.org/html/rfc5280#section-4.2.1.3
//
func Check(ex pkix.Extension, d *certdata.Data) *errors.Errors {
	var e = errors.New(nil)

	if !ex.Critical {
		e.Err("KeyUsage extension SHOULD be marked as critical when present")
	}

	return e
}
