package version

// These variables are set during build time using -ldflags
var (
	// GitCommit is the git commit hash
	GitCommit string
	// Version is the application version
	Version string
)

// GetVersion returns the application version information
func GetVersion() string {
	commit := GitCommit
	if commit == "" {
		commit = "unknown"
	}

	version := Version
	if version == "" {
		version = "development"
	}

	return "SMTP Sink v" + version + " (" + commit + ")"
}