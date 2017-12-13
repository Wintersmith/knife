package config

import (
    "fmt"
    "knife/files"
    "bufio"
    "log"
    "os"
    "regexp"
)

type Sections struct {
    sectionName string
    sectionElements *Elements
}

type Elements struct {
    elementName string
    elementValue string
}

func Load( fileName string ) {
    if doesExist := files.OSPathExists( fileName ); doesExist == true {
        fileConn, errMsg := os.Open( fileName )
        if errMsg != nil {
            log.Fatal( errMsg )
        }
        defer fileConn.Close()
        
        fileReader := bufio.NewScanner( fileConn )
        for fileReader.Scan() {
            fmt.Println( fileReader.Text() )
        }
    }
}
