package models

var address = []string{"https://ngcareers.com/cfeed/ict-jobs.xml"}

type Item struct {
	Title string `xml:"title"`
	Traffic string `xml:"ht:approx_traffic"`
	Link string `xml:"link"`
	Date string `xml:"pubDate"`
}

type Channel struct {
	 Title string `xml:"title"`
   Link  string `xml:"link"`
   Desc  string `xml:"description"`
   Items []Item `xml:"item"`
 }

 type Rss struct {
	 Channel Channel `xml:"channel"`
 }
