package all

import (
	// Import all default extensions
	_ "github.com/bryanpitcher/certlint/checks/extensions/adobetimestamp"
	_ "github.com/bryanpitcher/certlint/checks/extensions/authorityinfoaccess"
	_ "github.com/bryanpitcher/certlint/checks/extensions/authoritykeyid"
	_ "github.com/bryanpitcher/certlint/checks/extensions/basicconstraints"
	_ "github.com/bryanpitcher/certlint/checks/extensions/crldistributionpoints"
	_ "github.com/bryanpitcher/certlint/checks/extensions/ct"
	_ "github.com/bryanpitcher/certlint/checks/extensions/extkeyusage"
	_ "github.com/bryanpitcher/certlint/checks/extensions/keyusage"
	_ "github.com/bryanpitcher/certlint/checks/extensions/nameconstraints"
	_ "github.com/bryanpitcher/certlint/checks/extensions/ocspmuststaple"
	_ "github.com/bryanpitcher/certlint/checks/extensions/ocspnocheck"
	_ "github.com/bryanpitcher/certlint/checks/extensions/pdfrevocation"
	_ "github.com/bryanpitcher/certlint/checks/extensions/policyidentifiers"
	_ "github.com/bryanpitcher/certlint/checks/extensions/smimecapabilities"
	_ "github.com/bryanpitcher/certlint/checks/extensions/subjectaltname"
	_ "github.com/bryanpitcher/certlint/checks/extensions/subjectkeyid"
)
