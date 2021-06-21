package podmanexec

import (
	"github.com/komish/preflight/certification/errors"
	"github.com/komish/preflight/certification/internal/policy"
)

func HasMinimalVulnerabilities() *policy.Definition {
	return &policy.Definition{
		ValidatorFunc: hasMinimalVulnerabilitiesValidatorFunc,
		Metadata:      hasMinimalVulnerabilitiesPolicyMeta,
		HelpText:      hasMinimalVulnerabilitiesPolicyHelp,
	}
}

var hasMinimalVulnerabilitiesValidatorFunc = func(image string) (bool, error) {
	return false, errors.ErrFeatureNotImplemented
}

var hasMinimalVulnerabilitiesPolicyMeta = policy.Metadata{
	Description:      "Checking for critical or important security vulnerabilites.",
	Level:            "good",
	KnowledgeBaseURL: "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
	PolicyURL:        "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
}

var hasMinimalVulnerabilitiesPolicyHelp = policy.HelpText{
	Message:    "Components in the container image cannot contain any critical or important vulnerabilities, as defined at https://access.redhat.com/security/updates/classification",
	Suggestion: "Update your UBI image to the latest version or update the packages in your image to the latest versions distrubuted by Red Hat.",
}
