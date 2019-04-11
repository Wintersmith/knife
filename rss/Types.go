package rss

import "encoding/xml"

type RSSv1 struct {
	XMLName xml.Name       `xml:"RDF"`
	Channel *RSSv1Channel `xml:"channel"`
	Items   []RSSv1Item   `xml:"item"`
}

type RSSv1Channel struct {
	XMLName     xml.Name    `xml:"channel"`
	Title       string      `xml:"title"`
	Description string      `xml:"description"`
	Link        string      `xml:"link"`
	Image       Image    	`xml:"image"`
	TimeToLive  int         `xml:"ttl"`
	SkipHours   []int       `xml:"skipHours>hour"`
	SkipDays    []string    `xml:"skipDays>day"`
}

type Image struct {
	XMLName xml.Name `xml:"image"`
	Title   string   `xml:"title"`
	URL     string   `xml:"url"`
	Height  int      `xml:"height"`
	Width   int      `xml:"width"`
}

type RSSv1Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Content     string   `xml:"encoded"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Date        string   `xml:"date"`
	DateValid   bool
	ID          string            `xml:"guid"`
	Enclosures  []Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	URL     string   `xml:"resource,attr"`
	Type    string   `xml:"type,attr"`
	Length  uint     `xml:"length,attr"`
}

type Categories []string

type RSSv2 struct {
	XMLName xml.Name       `xml:"rss"`
	Channel *RSSv2Channel  `xml:"channel"`
}

type RSSv2Channel struct {
	XMLName     xml.Name     `xml:"channel"`
	Title       string       `xml:"title"`
	Description string       `xml:"description"`
	Link        []RSSv2Link `xml:"link"`
	Image       Image    `xml:"image"`
	Items       []RSSv2Item `xml:"item"`
	TimeToLive  int          `xml:"ttl"`
	SkipHours   []int        `xml:"skipHours>hour"`
	SkipDays    []string     `xml:"skipDays>day"`
}
type RSSv2Link struct {
	Rel      string `xml:"rel,attr"`
	Href     string `xml:"href,attr"`
	Type     string `xml:"type,attr"`
	Chardata string `xml:",chardata"`
}

type RSSv2Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Encoded     string   `xml:"encoded"`
	Categories  Categories `xml:"category"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Date        string   `xml:"date"`
	DateValid   bool
	ID          string    `xml:"guid"`
	Enclosures  []Enclosure `xml:"enclosure"`
}

type Atom struct {
	XMLName     xml.Name   `xml:"feed"`
	Title       string     `xml:"title"`
	Description string     `xml:"subtitle"`
	Link        []AtomLink `xml:"link"`
	Image       Image  `xml:"image"`
	Items       []AtomItem `xml:"entry"`
	Updated     string     `xml:"updated"`
}

type AtomItem struct {
	XMLName   xml.Name   `xml:"entry"`
	Title     string     `xml:"title"`
	Summary   string     `xml:"summary"`
	Content   AtomContent `xml:"content"`
	Links     []AtomLink `xml:"link"`
	Date      string     `xml:"updated"`
	DateValid bool
	ID        string `xml:"id"`
}
type AtomContent struct {
	AtomContent string `xml:",innerxml"`
}
type AtomLink struct {
	Href   string `xml:"href,attr"`
	Rel    string `xml:"rel,attr"`
	Type   string `xml:"type,attr"`
	Length uint   `xml:"length,attr"`
}