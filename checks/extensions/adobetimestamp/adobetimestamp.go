package adobetimestamp

import (
	"crypto/x509/pkix"
	"encoding/asn1"

	"github.com/bryanpitcher/certlint/certdata"
	"github.com/bryanpitcher/certlint/checks"
	"github.com/bryanpitcher/certlint/errors"
)

const checkName = "Adobe Timestamp Extension Check"

var extensionOid = asn1.ObjectIdentifier{1, 2, 840, 113583, 1, 1, 9, 1}

func init() {
	checks.RegisterExtensionCheck(checkName, extensionOid, nil, Check)
}

// Check performs a strict verification on the extension according to the standard(s)
func Check(ex pkix.Extension, d *certdata.Data) *errors.Errors {
	var e = errors.New(nil)

	if ex.Critical {
		e.Err("Adobe Timestamp extension set critical")
	}

	return e
}
