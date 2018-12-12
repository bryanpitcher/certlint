package all

import (
	// Import all default checks
	_ "github.com/bryanpitcher/certlint/checks/certificate/aiaissuers"
	_ "github.com/bryanpitcher/certlint/checks/certificate/basicconstraints"
	_ "github.com/bryanpitcher/certlint/checks/certificate/extensions"
	_ "github.com/bryanpitcher/certlint/checks/certificate/extkeyusage"
	_ "github.com/bryanpitcher/certlint/checks/certificate/internal"
	_ "github.com/bryanpitcher/certlint/checks/certificate/issuerdn"
	_ "github.com/bryanpitcher/certlint/checks/certificate/keyusage"
	_ "github.com/bryanpitcher/certlint/checks/certificate/publickey"
	_ "github.com/bryanpitcher/certlint/checks/certificate/publicsuffix"
	_ "github.com/bryanpitcher/certlint/checks/certificate/revocation"
	_ "github.com/bryanpitcher/certlint/checks/certificate/serialnumber"
	_ "github.com/bryanpitcher/certlint/checks/certificate/signaturealgorithm"
	_ "github.com/bryanpitcher/certlint/checks/certificate/subject"
	_ "github.com/bryanpitcher/certlint/checks/certificate/subjectaltname"
	_ "github.com/bryanpitcher/certlint/checks/certificate/validity"
	_ "github.com/bryanpitcher/certlint/checks/certificate/version"
	_ "github.com/bryanpitcher/certlint/checks/certificate/wildcard"
)
