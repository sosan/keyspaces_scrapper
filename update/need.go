package update

import (
	"fmt"
	"main/httpclient"
	"strings"
)

func getNeedUpdate() bool {
	selfVersion := GetVersionBuild()
	remoteVersion := getRemoteVersionBuild()

	return remoteVersion != selfVersion
}

func getRemoteVersionBuild() string {
	uri := getRemotePathBuildDate()
	versionDate, statusCode := httpclient.GetRequest(uri, "")
	if statusCode != 200 {
		return ""
	}
	remoteVersionBuild := cleanRawDate(string(versionDate))

	return remoteVersionBuild
}

func cleanRawDate(dateRaw string) string {
	splitted := strings.Split(dateRaw, "\n")
	if len(splitted) == 0 {
		return ""
	}
	return splitted[0]
}

func getRemotePathBuildDate() string {
	fileName := getFileName()
	return fmt.Sprintf("%s/date_build_%s.txt", URI_DOWNLOAD, fileName)
}
