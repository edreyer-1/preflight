package podmanexec

import (
	"github.com/komish/preflight/certification/errors"
	"github.com/komish/preflight/certification/internal/policy"
)

func HasUniqueTag() *policy.Definition {
	return &policy.Definition{
		ValidatorFunc: hasUniqueTagValidatorFunc,
		Metadata:      hasUniqueTagTagPolicyMeta,
		HelpText:      hasUniqueTagTagPolicyHelp,
	}
}

var hasUniqueTagValidatorFunc = func(image string) (bool, error) {
	return false, errors.ErrFeatureNotImplemented
}

var hasUniqueTagTagPolicyMeta = policy.Metadata{
	Description:      "Checking if container has a tag other than 'latest'.",
	Level:            "best",
	KnowledgeBaseURL: "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
	PolicyURL:        "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
}

var hasUniqueTagTagPolicyHelp = policy.HelpText{
	Message:    "Containers should have a tag other than latest, so that the image can be uniquely identfied.",
	Suggestion: "Add a tag to your image. Consider using Semantic Versioning. https://semver.org/",
}
