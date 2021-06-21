package certification

import (
	"github.com/komish/preflight/certification/internal/policy"
	"github.com/komish/preflight/certification/internal/policy/podmanexec"
)

// PolicyFunc is a function that returns a policy.
type PolicyFunc = func() Policy

// Policy as an interface containing all methods necessary
// to use and identify a given policy.
type Policy interface {
	// Validate identify whether the asset enforces
	// the policy.
	Validate(image string) (result bool, err error)
	Meta() policy.Metadata
	// Help returns the policy's help information.
	Help() policy.HelpText
}

// PolicyMap maps policy string names to policy functions.
var PolicyMap = map[string]PolicyFunc{
	"is_ubi_based":                 PolicyIsUBIBased,
	"runs_as_nonroot":              PolicyRunsAsNonRootUser,
	"has_under_40_layers":          PolicyHasLessThan40Layers,
	"has_required_labels":          PolicyImageHasRequiredLabels,
	"has_unique_tag":               PolicyHasUniqueTag,
	"no_excessive_vulnerabilities": PolicyNoExcessiveVulnerabilities,
	"no_prohibited_packages":       PolicyNoProhibitedPackages,
}

// AllPolicies returns all policies made available by this library.
func AllPolicies() []string {
	all := make([]string, len(PolicyMap))
	i := 0

	for k := range PolicyMap {
		all[i] = k
		i++
	}

	return all
}

// PolicyHasLessThan40Layers checks the container image and ensures
// it has less than 40 layers.
func PolicyHasLessThan40Layers() Policy {
	return podmanexec.UnderLayerMax()
}

// PolicyIsUBIBased checks the container image and ensures
// it is based on the Red Hat Universal Base Image.
func PolicyIsUBIBased() Policy {
	return podmanexec.BasedOnUBI()
}

// PolicyRunsAsNonRootUser checks that the container image is not
// configured to run as the root user
func PolicyRunsAsNonRootUser() Policy {
	return podmanexec.RunsAsNonRootUser()
}

// PolicyImageHasRequiredLabels checks that the container image has
// the required labels
func PolicyImageHasRequiredLabels() Policy {
	return podmanexec.HasRequiredLabels()
}

// PolicyHasUniqueTag checks that the container image's metadata
// contains any tag other than the 'latest' pseudo-tag.
func PolicyHasUniqueTag() Policy {
	return podmanexec.HasUniqueTag()
}

// PolicyNoExcessiveVulnerabilities checks that the container image does not have
// vulnerabilities above a given severity. TODO: document severity
func PolicyNoExcessiveVulnerabilities() Policy {
	return podmanexec.HasMinimalVulnerabilities()
}

// PolicyNoProhibitedPackages checks that the container image does not contain
// RHEL packages that are not redistributable due to license restrictions.
func PolicyNoProhibitedPackages() Policy {
	return podmanexec.NoProhibitedPackages()
}
