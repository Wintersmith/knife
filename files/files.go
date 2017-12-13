package files

import (
	"os"
	"fmt"
	"path/filepath"
)

type osDeleteParams struct {
	fullPath string
	deleteTree bool
}

func OSPathExists( fullPath string ) bool {
	if _, errMsg := os.Stat( fullPath ); os.IsNotExist( errMsg ) {
		return false
	}
	return true
}
func IsError( errObj error ) bool {
	return ( errObj != nil )
}
func OSDelete( osDelSettings osDeleteParams ) ( string, bool ) {
	if errMsg := os.Remove( osDelSettings.fullPath ); IsError( errMsg ) {
		return errMsg.Error(), false
	}
	return "", true
}
func OSRename( fileFrom string, fileTo string ) ( string, bool ) {
	if OSPathExists( fileFrom ) {
			if errMsg := os.Rename( fileFrom, fileTo ); IsError( errMsg ) {
				return errMsg.Error(), false
			}
		return "", true
	}
	return "File Not Found", false
}
func foundFile( path string, fileInfo os.FileInfo, errMsg error ) error {
    if fileInfo.IsDir() {
        return nil
    }
    fmt.Printf( "Found %s\n", path )
    return nil
}
func FindFiles( startDir string ) {
    filepath.Walk( startDir, foundFile )
}