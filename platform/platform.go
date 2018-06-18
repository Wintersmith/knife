package platform

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"runtime"
)
func InstalledVersion() string {
	goVersion := ""
	goPath, OK := os.LookupEnv( "GOPATH" )
	if OK {
		versionFilePath := filepath.Join(goPath, "VERSION")
		if _, errMsg := os.Stat( versionFilePath ); os.IsExist( errMsg ) {
			bytesInFile, errMsg := ioutil.ReadFile( versionFilePath )
			if errMsg == nil {
				goVersion = string( bytesInFile )
			}
		}
	}

	return goVersion
}

func CompiledVersion() string {
	return runtime.Version()
}