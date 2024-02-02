package update

import (
	"fmt"
	"log"
	"main/calculatehash"
	"main/httpclient"
	"strings"
)

func getNeedUpdate() bool {
	fileName := getFileName()
	selfCheckSum := calculatehash.CalculateSha256Checksum(fileName)
	log.Print(selfCheckSum)
	remoteChecksum := getRemoteChecksum()

	return remoteChecksum != selfCheckSum
}



func getRemoteChecksum() string {
	uri := getRemotePathChecksum()
	byteChecksum, statusCode := httpclient.GetRequest(uri, "")
	if statusCode != 200 {
		return ""
	}
	remoteChecksum := cleanRawChecksum(string(byteChecksum))
	
	return remoteChecksum
}

func cleanRawChecksum(checksumRaw string) string {
	splitedChecksum := strings.Split(checksumRaw, "\n")
	if len(splitedChecksum) == 0 {
		return ""
	}
	return splitedChecksum[0]
}

func getRemotePathChecksum() string {
	fileName := getFileName()
	return fmt.Sprintf("%s/checksum_%s.txt", URI_DOWNLOAD, fileName)
}