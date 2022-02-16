package main

import (
	"fmt"
    "log"
    "github.com/gocolly/colly"
)

func main() {
    c := colly.NewCollector()

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.OnError(func(_ *colly.Response, err error) {
        log.Println("Something went wrong:", err)
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Visited", r.Request.URL)
    })

    c.OnHTML("span.header-nav-data:nth-child(2)", func(e *colly.HTMLElement) {
        log.Println("Price:", e.Text)
    })

    c.OnScraped(func(r *colly.Response) {
        fmt.Println("Finished", r.Request.URL)
    })

    c.Visit("https://www.ishares.com/uk/individual/en/products/280510/ishares-sp-500-information-technology-sector-ucits-etf?switchLocale=y&siteEntryPassthrough=true")
}
