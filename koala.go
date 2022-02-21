package main

import (
	"fmt"
    "net/http"
    "io"
    "log"
    "github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
  // This request gives back a json file with all holdings data
  // curl 'https://www.ishares.com/uk/individual/en/products/280510/ishares-sp-500-information-technology-sector-ucits-etf/1506575576011.ajax?tab=all&fileType=json' -J -O -L

  // Request the HTML page.
  res, err := http.Get("https://www.ishares.com/uk/individual/en/products/280510/ishares-sp-500-information-technology-sector-ucits-etf?switchLocale=y&siteEntryPassthrough=true")
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }
  body, err := io.ReadAll(res.Body)
  if err != nil {
      log.Fatal(err)
  }
  log.Println(string(body))

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  // Find the review items
  doc.Find("#allHoldingsTable tbody tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
        log.Println(s)
		title := s.Find("td").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}

func main() {
  ExampleScrape()
}
