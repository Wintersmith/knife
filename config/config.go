package config

import (
    "knife/files"
    "bufio"
    "os"
	"github.com/golang/glog"
	"regexp"
	"strings"
)

type Config struct {
	sections []Sections
}
type Sections struct {
    sectionName string
    sectionElements []Elements
}

type Elements struct {
    elementName string
    elementValue string
}


func CP() *Config {
	configObj := Config{}

	return &configObj
}

func ( configFile *Config ) Load( fileName string ) {
    if doesExist := files.OSPathExists( fileName ); doesExist == true {
        fileConn, errMsg := os.Open( fileName )
        if errMsg != nil {
            glog.Fatal( errMsg )
        }
        defer fileConn.Close()

        reSectionHeading, _ := regexp.Compile( `\[(.*)\]` )
        reKeyValue, _ := regexp.Compile( `(.+)=(.+)` )
        reComment, _ := regexp.Compile( `^#|//.*` )
        sectionTitle := ""
        curSectionTitle := ""
        var newElements Elements
        var newSection Sections

        fileReader := bufio.NewScanner( fileConn )
        for fileReader.Scan() {
        	if reComment.MatchString( fileReader.Text() ) != true {
				vValueResult := reKeyValue.FindStringSubmatch( fileReader.Text() )
				if len( vValueResult ) > 0 {
					newElements.elementName = vValueResult[ 1 ]
					newElements.elementValue = vValueResult[ 2 ]
					newSection.sectionElements = append( newSection.sectionElements, newElements )
					newElements = Elements{}
					continue
				}
				if reSectionHeading.MatchString( fileReader.Text() ) {
					sectionTitle = strings.ToLower( fileReader.Text()[ 1:len( fileReader.Text() ) - 1 ] )
					if curSectionTitle != sectionTitle && curSectionTitle != "" {
						configFile.sections = append( configFile.sections, newSection )
					}
					newSection = Sections{}
					newSection.sectionName = sectionTitle
					newElements = Elements{}
					curSectionTitle = sectionTitle
				}
			}
        }
		if curSectionTitle != "" {
			configFile.sections = append( configFile.sections, newSection )
		}

	} else {
		glog.Error( "Settings File Doesn't Exist, Please Verify: ", fileName )
	}
}
func ( configFile *Config ) Get( sectionName string, elementName string ) {

}
