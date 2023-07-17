package structs

import (
	"database/sql"
	"encoding/xml"
	"time"
)

type UserStruct struct {
	UserID    string       `db:"user_id" json:"userID"`
	Name      string       `db:"name" json:"name"`
	Email     string       `db:"email" json:"email"`
	Password  string       `db:"password" json:"password"`
	CreatedAt time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updatedAt"`
}

// Declare a struct to hold the XML data.
type XmlData struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

type Opml struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version"`
	Head    struct {
		Title string `xml:"title"`
	} `xml:"head"`
	Body struct {
		Outline struct {
			Outline []Outline `xml:"outline"`
		} `xml:"outline"`
	} `xml:"body"`
}

type Outline struct {
	Text        string `xml:"text,attr"`
	Title       string `xml:"title,attr"`
	Type        string `xml:"type,attr"`
	XMLUrl      string `xml:"xmlUrl,attr"`
	HtmlUrl     string `xml:"htmlUrl,attr"`
	Description string `xml:"description,attr"`
}

type RSSFeedDB struct {
	FeedID      string `db:"feed_id" json:"feedID"`
	XmlUrl      string `db:"xml_url" json:"xmlUrl"`
	HtmlUrl     string `db:"html_url" json:"htmlUrl"`
	Title       string `db:"title" json:"title"`
	Type        string `db:"type" json:"type"`
	Description string `db:"description" json:"description"`
	Text        string `db:"text" json:"text"`
	OpmlID      string `db:"opml_id" json:"opml_id"`
	CreatedAt   string `db:"created_at" json:"createdAt"`
}

type OpmlStoreDB struct {
	OpmlID    string `db:"opml_id" json:"opmlID"`
	Title     string `db:"title" json:"title"`
	Version   string `db:"version" json:"version"`
	Count     string `db:"entries_count" json:"count"`
	Filename  string `db:"filename" json:"fileName"`
	CreatedAt string `db:"created_at" json:"createdAt"`
}
