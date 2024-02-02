package update

var VersionBuild string

func GetVersionBuild() string {
	return VersionBuild
}

func SetVersionBuild(version string) {
	VersionBuild = version
}
