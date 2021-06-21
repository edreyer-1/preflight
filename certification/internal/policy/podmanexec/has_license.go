package podmanexec

import (
	"github.com/komish/preflight/certification/errors"
	"github.com/komish/preflight/certification/internal/policy"
)

func HasLicense() *policy.Definition {
	return &policy.Definition{
		ValidatorFunc: hasLicenseValidatorFunc,
		Metadata:      hasLicensePolicyMeta,
		HelpText:      hasLicensePolicyHelp,
	}
}

var hasLicenseValidatorFunc = func(image string) (bool, error) {
	return false, errors.ErrFeatureNotImplemented
}

var hasLicensePolicyMeta = policy.Metadata{
	Description:      "Checking if terms and conditions for images are present.",
	Level:            "best",
	KnowledgeBaseURL: "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
	PolicyURL:        "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
}

var hasLicensePolicyHelp = policy.HelpText{
	Message:    "Container images must include terms and conditions applicable to the software including open source licensing information.",
	Suggestion: "Create a directory named /licenses and include all relevant licensing and/or terms and conditions as text file(s) in that directory.",
}
