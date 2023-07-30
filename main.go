package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"google-trends/models"
)


func main() {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=KR")
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return
	}
	defer resp.Body.Close()

	rss := models.Rss{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
		return
	}

	for _, item := range rss.Channel.Items {
		data := item
		fmt.Println(data)
		// store data to database
	}
}

