package src

import (
  "log"
  "net/http"
  "time"
  "net"
  "fmt"

  "github.com/gocolly/colly"
)

func Scrape (gurl string) []Route {
  ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
  "AppleWebKit/537.36 (KHTML, like Gecko) " +
  "Chrome/110.0.0.0 Safari/537.36"

  sc := colly.NewCollector(
    colly.AllowURLRevisit(),
    colly.Async(true),
  )

  sc.WithTransport(&http.Transport {
    Proxy: http.ProxyFromEnvironment,
    DialContext: (&net.Dialer{
      Timeout:   30 * time.Second,
      KeepAlive: 30 * time.Second,
      DualStack: true,
    }).DialContext,
    ForceAttemptHTTP2:     true,
    MaxIdleConns:          100,
    IdleConnTimeout:       90 * time.Second,
    TLSHandshakeTimeout:   10 * time.Second,
    ExpectContinueTimeout: 1 * time.Second,
  })

  sc.OnRequest(func(r *colly.Request) {
    r.Headers.Set("User-Agent", ua)
    r.Headers.Set(
      "Accept",
      "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
    )
    r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
  })

  sc.OnError(func(_ *colly.Response, err error) {
    log.Fatal("エラー：", err)
  })

  var routeArr []Route

  for i := 1; i <= 3; i++ {
    route := fmt.Sprintf("div#route%02d", i)
    sc.OnHTML("div.elmRouteDetail " + route, func (e *colly.HTMLElement) {
      Routes := getRouteDetail(e)
      routeArr = append(routeArr, Routes)
    })
  }

  sc.Visit(gurl)
  sc.Wait()

  return routeArr
}
