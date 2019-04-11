package rss

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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

type Entry struct {
	Title      string    `json:"title"`
	Summary    string    `json:"summary"`
	Content    string    `json:"content"`
	Link       string    `json:"link"`
	Date       time.Time `json:"date"`
	DateValid  bool
	ID         string       `json:"id"`
}


func Parse( rssString string ) {

	feedXML, errMsg := getFeed(rssString)
	if errMsg != nil {
		fmt.Println("Couldn't Access Remote Feed")
		return
	}
	matchFeed( feedXML )
}
func matchFeed( feedXML []byte ) {

	v1Feed := RSSv1{}
	xmlDecoder := xml.NewDecoder( bytes.NewReader( feedXML ) )
	xmlDecoder.CharsetReader = charset.NewReaderLabel
	xmlDecoder.Strict = false
	errMsg := xmlDecoder.Decode( &v1Feed )

	if errMsg != nil {
		v2Feed := RSSv2{}
		xmlDecoder = xml.NewDecoder( bytes.NewReader( feedXML ) )
		xmlDecoder.CharsetReader = charset.NewReaderLabel
		xmlDecoder.Strict = false
		errMsg = xmlDecoder.Decode( &v2Feed)
		if errMsg != nil {
			atomFeed := Atom{}
			xmlDecoder = xml.NewDecoder( bytes.NewReader( feedXML ) )
			xmlDecoder.CharsetReader = charset.NewReaderLabel
			xmlDecoder.Strict = false
			errMsg = xmlDecoder.Decode( &atomFeed)
			if errMsg != nil {
				fmt.Println("Failed Everything", errMsg.Error())
				return
			} else {
				fmt.Println("Returning atom")
				parseAtom(atomFeed)
			}
		} else {
			fmt.Println("Returning v2")
			parseV2(v2Feed)
		}
	} else {
		fmt.Println("Returning v1")
		parseV1(v1Feed)
	}
}
func parseAtom( aFeed Atom ) {
	for _, atomItem := range aFeed.Items {
		fmt.Println(atomItem)
	}
}
func parseV1( v1 RSSv1 ) {
	for _, v1Item := range v1.Items {
		fmt.Println( v1Item )
	}
}
func parseV2( v2 RSSv2 ) {
	for _, v2Item := range v2.Channel.Items {
		fmt.Println(v2Item)
	}
}
func getFeed( remoteURL string ) ( []byte, error ) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	httpReq, errMsg := http.NewRequest( http.MethodGet, remoteURL, nil )
	if errMsg != nil {
		fmt.Println( "Encountered Error Creating Request")
	}
	httpResp, errMsg := httpClient.Do( httpReq )
	if errMsg != nil {
		fmt.Println( "Encountered Error", errMsg.Error() )
		return nil, errMsg
	}
	if httpResp.StatusCode != 200 {
		return nil, errors.New( "Invalid URL" )
	}
	defer httpResp.Body.Close()

	feedXML, errMsg := ioutil.ReadAll( httpResp.Body )
	if errMsg != nil {
		fmt.Println( "Couldn't Read Feed XML", errMsg.Error() )
		return nil, errMsg
	}

	return feedXML, nil

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

