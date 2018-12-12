package publickey

import (
	"strings"

	"github.com/bryanpitcher/certlint/certdata"
	"github.com/bryanpitcher/certlint/checks"
	"github.com/bryanpitcher/certlint/checks/certificate/publickey/goodkey"
	"github.com/bryanpitcher/certlint/errors"
)

const checkName = "Public Key Check"

func init() {
	checks.RegisterCertificateCheck(checkName, nil, Check)
}

// Check performs a strict verification on the extension according to the standard(s)
func Check(d *certdata.Data) *errors.Errors {
	var e = errors.New(nil)

	gkp := goodkey.NewKeyPolicy()
	err := gkp.GoodKey(d.Cert.PublicKey)
	if err != nil {
		e.Err("Certificate %s", strings.ToLower(err.Error()))
		return e
	}

	return e
}
