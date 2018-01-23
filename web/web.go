package web

import (
    "os"
    "net"
    "net/http"
    "net/url"
    "io"
)

type DownloadFileParams struct {
    LocalFile string
    RemoteURL string
    UserName string
    PassWord string
}

func DownloadFile( dlFileParams DownloadFileParams ) ( errMsg error ) {
    
    // The following is groundwork for accepting various schemas
    parsedURL, errMsg := url.Parse( dlFileParams.RemoteURL )
    if errMsg != nil {
        return errMsg
    }
    hostName, hostPort, errMsg := net.SplitHostPort( parsedURL.Host )
    if errMsg != nil {
        hostName = parsedURL.Host
        hostPort = ""
    }

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