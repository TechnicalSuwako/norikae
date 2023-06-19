package main

import (
  "log"
  "net/http"

  "github.com/gocolly/colly"
)

type (
  Badge struct {
    Text, Color string
  }
  Route struct {
    Id, From, To, Duration, RideDuration, TransitCunt, Km string
    Price, Sum int
    Badges []Badge
    Trains []Train
  }
  Stops struct {
    Time, Name string
  }
  Train struct {
    From, To, Name, LineColor, Station, LineName, FromTrack, ToTrack string
    Stop []Stops
  }
  ScrapeData struct {
    Err string
  }
)

func scrape (gurl string) *ScrapeData {
  var scrapeData *ScrapeData
  var badge1, badge2, badge3 Badge
  badge1.Text = "早"
  badge1.Color = "#ff7e56"
  badge2.Text = "楽"
  badge2.Color = "#60bddb"
  badge3.Text = "安"
  badge3.Color = "#fab60a"
  log.Println(gurl)

  resp, err := http.Get(gurl)
  if err != nil {
    log.Fatal(err)
  }
  if resp.StatusCode == 404 {
    log.Fatal("見つけられません。")
  }

  ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
  sc := colly.NewCollector(colly.AllowedDomains("transit.yahoo.co.jp"), colly.UserAgent(ua))

  sc.OnHTML("div.elmRouteDetail", func (e *colly.HTMLElement) {
    log.Println(e.Attr("class"))
    log.Println(e.Attr("id"))
    if e.Attr("class") == "pos-im" {
      return
    }

    e.ForEach("div", func (i int, el *colly.HTMLElement) {
      log.Println("cunny")
      //el.ClassText()
    })
  })

  return scrapeData
}
