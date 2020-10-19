package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/PuerkitoBio/goquery"
)

// CheckUserOnline - Request the HTML page aka  user profile
// and if online load the site in the browser
func CheckUserOnline(user string) int {

	res, err := http.Get("https://www.chess.com/member/" + user)

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

	// Find the user profile
	doc.Find(".profile-info-item-value").Each(func(i int, s *goquery.Selection) {
		// it is the second class that contains the info that's why we check i == 1
		if i == 1 {
			player := s.Text()
			if player == "In Live" || player == "Online Now" {
				fmt.Println("YES!!!")
				url := "https://chess.com/live"
				exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
				os.Exit(0)
			}
		}
	})
	return 0
}

func main() {

	// if any of the users is online open the site, although
	// the data in user's profile is not always consistent(currently online user is written as 1 hr ago)
	user := []string{"wonderfultime", "erichansen", "danielnaroditsky", "davit_tiraturyan", "maghalashvili"}
	for i := 0; i < len(user); i++ {
		CheckUserOnline(user[i])
	}
}
