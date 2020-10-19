package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/PuerkitoBio/goquery"
)

// CheckUserOnline - Request the HTML page aka  user profile
// and if online load the site in the browser
func CheckUserOnline() {

	// res, err := http.Get("https://www.chess.com/member/davit_tiraturyan")
	// res, err := http.Get("https://www.chess.com/member/danielnaroditsky")

	res, err := http.Get("https://www.chess.com/member/erichansen")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".profile-info-item-value").Each(func(i int, s *goquery.Selection) {
		// it is the second class that contains the info
		if i == 1 {
			player := s.Text()
			if player == "In Live" || player == "Online Now" {
				fmt.Println("YES!!!")
				url := "https://chess.com/live"
				exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
			}
		}
	})
}

func main() {
	CheckUserOnline()
}
