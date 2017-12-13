package web

import (
    "os"
    "net/http"
    "io"
)

func DownloadFile( localFile string, remoteURL string ) ( errMsg error ) {

  localFH, errMsg := os.Create( localFile )
  if errMsg != nil  {
    return errMsg
  }
  defer localFH.Close()

  // Get the data
  httpResp, errMsg := http.Get( remoteURL )
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