package update

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"main/httpclient"
	"os"
	"strings"
)

func getNeedUpdate() bool {
	fileName := getFileName()
	selfCheckSum := getBinaryChecksum(fileName)
	log.Print(selfCheckSum)
	remoteChecksum := getRemoteChecksum()

	return remoteChecksum != selfCheckSum
}

func getBinaryChecksum(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	checksum := fmt.Sprintf("%x", h.Sum(nil))
	return checksum
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