package constants

const (
	VersionerDir    = ".versioner"
	VersionerConfig = ".versionerconfig.json"
	VersionerIgnore = ".versionerignore"
)

var (
	IgnoreFiles = []string{VersionerIgnore, VersionerDir, ".git", ".gitignore"}
)
