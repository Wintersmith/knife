package web

import (
    "os"
    "net"
    "net/http"
    "net/url"
    "io"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"golang.org/x/net/html"
	"regexp"
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

func RespondWithError( httpWriter http.ResponseWriter, retCode int, errMessage string) {
	RespondWithJSON( httpWriter, retCode, map[string]string{ "error": errMessage } )
}
func RespondWithJSON( httpWriter http.ResponseWriter, retCode int, httpPayload interface{}) {
	httpResp, _ := json.Marshal( httpPayload )
	httpWriter.Header().Set("Content-Type", "application/json")
	httpWriter.WriteHeader( retCode )
	httpWriter.Write( httpResp )
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

func GetGoVersion( ) ( goVersion string ) {
	var versionNo = ""

	httpResp, errMsg := http.Get( "https://www.golang.org" )
	if errMsg != nil || httpResp.StatusCode < 200 || httpResp.StatusCode > 300 {
		return "Got Error, Or Non 200 Code"
	}
	defer httpResp.Body.Close()

	htmlRoot:= html.NewTokenizer( httpResp.Body )
	if errMsg != nil {
		return " Failed To Get Tokenizer"
	}
	reMatchScript, _ := regexp.Compile( "goVersion = \"(.*)\";")
	inTags := false
	loop:
	for {
		htmlToken := htmlRoot.Next()
		switch htmlToken {
			case html.ErrorToken:
				break loop
			case html.TextToken:
				if inTags {
					groupMatch := reMatchScript.FindAllStringSubmatch( string( htmlRoot.Text()[:] ), -1)
					if len(groupMatch) > 0 {
						versionNo = groupMatch[ 0 ][ 1 ]
						break loop
					}
				}
			case html.StartTagToken, html.EndTagToken:
				if htmlRoot.Token().String() == "<script>" {
					inTags = true
				} else if htmlRoot.Token().String() == "</script>" && inTags {
					inTags = false
				}
			}
		}

	return versionNo
}