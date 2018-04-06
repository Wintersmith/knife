package rss

import (
	"fmt"
	"encoding/xml"
	"os"
	"io/ioutil"
)

type OPML struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    Head     `xml:"head"`
	Body    Body     `xml:"body"`
}

// Head holds some meta information about the document.
type Head struct {
	Title           string `xml:"title"`
	DateCreated     string `xml:"dateCreated,omitempty"`
	DateModified    string `xml:"dateModified,omitempty"`
	OwnerName       string `xml:"ownerName,omitempty"`
	OwnerEmail      string `xml:"ownerEmail,omitempty"`
	OwnerID         string `xml:"ownerId,omitempty"`
	Docs            string `xml:"docs,omitempty"`
	ExpansionState  string `xml:"expansionState,omitempty"`
	VertScrollState string `xml:"vertScrollState,omitempty"`
	WindowTop       string `xml:"windowTop,omitempty"`
	WindowBottom    string `xml:"windowBottom,omitempty"`
	WindowLeft      string `xml:"windowLeft,omitempty"`
	WindowRight     string `xml:"windowRight,omitempty"`
}

// Body is the parent structure of all outlines.
type Body struct {
	Outlines []Outline `xml:"outline"`
}

// Outline holds all information about an outline.
type Outline struct {
	Outlines     []Outline `xml:"outline"`
	Text         string    `xml:"text,attr"`
	Type         string    `xml:"type,attr,omitempty"`
	IsComment    string    `xml:"isComment,attr,omitempty"`
	IsBreakpoint string    `xml:"isBreakpoint,attr,omitempty"`
	Created      string    `xml:"created,attr,omitempty"`
	Category     string    `xml:"category,attr,omitempty"`
	XMLURL       string    `xml:"xmlUrl,attr,omitempty"`
	HTMLURL      string    `xml:"htmlUrl,attr,omitempty"`
	URL          string    `xml:"url,attr,omitempty"`
	Language     string    `xml:"language,attr,omitempty"`
	Title        string    `xml:"title,attr,omitempty"`
	Version      string    `xml:"version,attr,omitempty"`
	Description  string    `xml:"description,attr,omitempty"`
}

func Parse( rssString string ) {
	fmt.Printf( rssString )
}

func LoadOPML( opmlFile string ) ( Body ){
	fmt.Println( "Loading", opmlFile )
	opmlXML, errMsg := os.Open( opmlFile )
	if errMsg != nil {
		fmt.Println( "Unable To Open File:", errMsg )
	}
	defer opmlXML.Close()

	var opmlList OPML
	xmlBytes, _ := ioutil.ReadAll( opmlXML )
	xml.Unmarshal( xmlBytes, &opmlList )

	return opmlList.Body
}