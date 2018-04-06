package web

import (
    "os"
    "net"
    "net/http"
    "net/url"
    "io"
	"io/ioutil"
	"fmt"
)

type DownloadFileParams struct {
    LocalFile string
    RemoteURL string
    UserName string
    PassWord string
}

type DownloadURLParams struct {
    RemoteURL string
    UserName string
    PassWord string
}

type HTTPError struct {
	StatusCode int
	Status     string
}

func ( hError HTTPError) Error() string {
	return fmt.Sprintf("HTTP Error: %s", hError.Status)
}

func DownloadFile( dlFileParams DownloadFileParams ) ( errMsg error ) {
    
    // The following is groundwork for accepting various schemas
    parsedURL, errMsg := url.Parse( dlFileParams.RemoteURL )
    if errMsg != nil {
        return errMsg
    }

	_, _ = splitURL( parsedURL.Host )

    // Create the local file
    localFH, errMsg := os.Create( dlFileParams.LocalFile )
    if errMsg != nil  {
        return errMsg
    }
    defer localFH.Close()

    // Get the data
    httpResp, errMsg := http.Get( dlFileParams.RemoteURL )
    if errMsg != nil {
        return errMsg
    }
    defer httpResp.Body.Close()

    // Writer the body to file
    _, errMsg = io.Copy( localFH, httpResp.Body )
    if errMsg != nil  {
        return errMsg
    }

    return nil
}
func splitURL( fullURL string ) ( hName string, hPort string ) {
	hostName, hostPort, errMsg := net.SplitHostPort( fullURL )
	if errMsg != nil {
		hostName = fullURL
		hostPort = ""
	}
	fmt.Printf( hostName, hostPort )

	return hName, hPort
}


func DownloadURL( dlURLParams DownloadURLParams ) ( urlContent string, errMsg error ) {

    // The following is groundwork for accepting various schemas
    parsedURL, errMsg := url.Parse( dlURLParams.RemoteURL )
    if errMsg != nil {
        return "", errMsg
    }
	_, _ = splitURL( parsedURL.Host )

    // Get the data
    httpResp, errMsg := http.Get( dlURLParams.RemoteURL )
    if errMsg != nil {
        return "", errMsg
    }
	if httpResp.StatusCode < 200 || httpResp.StatusCode > 300 {
		return "", HTTPError{ StatusCode: httpResp.StatusCode, Status: httpResp.Status }
	}

	defer httpResp.Body.Close()
	remoteContent, errMsg := ioutil.ReadAll( httpResp.Body )

//	contentLength := bytes.IndexByte( remoteContent, 0 )
    return string( remoteContent[ : ] ), nil
}