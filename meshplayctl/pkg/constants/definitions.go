package constants

const (
	// Meshplay Repository Location
	mesheryGitHubOrg  string = "meshery"
	mesheryGitHubRepo string = "meshery"
)

// GetMeshplayGitHubOrg retrieves the name of the GitHub organization under which the Meshplay repository resides.
func GetMeshplayGitHubOrg() string {
	return mesheryGitHubOrg
}

// GetMeshplayGitHubRepo retrieves the name of the Meshplay repository
func GetMeshplayGitHubRepo() string {
	return mesheryGitHubRepo
}
