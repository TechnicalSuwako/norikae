package main

import (
  "log"
  "net/http"
  "time"
  "net"
  "fmt"
  "strings"

  "github.com/gocolly/colly"
)

func scrape (gurl string) []Route {
  ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
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
    r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
  })

  sc.OnError(func(_ *colly.Response, err error) {
    log.Fatal("エラー：", err)
  })

  var routeArr []Route

  for i := 1; i <= 3; i++ {
    route := fmt.Sprintf("div#route%02d", i)
    sc.OnHTML("div.elmRouteDetail " + route, func (e *colly.HTMLElement) {
      Routes := Route{}
      e.ForEach("dl.routeSummary ul.priority li span", func (j int, el *colly.HTMLElement) {
        if el.Attr("class") == "icnPriTime" {
          Routes.Badges = append(Routes.Badges, 1)
        }
        if el.Attr("class") == "icnPriFare" {
          Routes.Badges = append(Routes.Badges, 2)
        }
        if el.Attr("class") == "icnPriTrans" {
          Routes.Badges = append(Routes.Badges, 3)
        }
      })
      base := e.ChildText("dl.routeSummary li.time span")
      time := strings.ReplaceAll(base, e.ChildText("dl.routeSummary li.time span.small"), "")
      time2 := strings.Split(time, "着")
      Routes.Time = time2[0] + "着"

      Routes.Duration = e.ChildText("dl.routeSummary li.time span.small")
      Routes.TransitCunt = strings.ReplaceAll(e.ChildText("dl.routeSummary li.transfer"), "乗換：", "")
      Routes.Fare = strings.ReplaceAll(e.ChildText("dl.routeSummary li.fare"), "[priic]IC優先：", "")
      Stations := Station{}
      Fares := Fare{}
      Stops := Stop{}
      e.ForEach("div.routeDetail div.station", func (j int, el *colly.HTMLElement) {
        Stations.Time = el.ChildText("ul.time li")
        if el.ChildText("p.icon span") == "[dep]" { Stations.Time += "発" }
        if el.ChildText("p.icon span") == "[arr]" { Stations.Time += "着" }
        Stations.Name = el.ChildText("dl dt a")
        e.ForEach("div.routeDetail div.fareSection div.access", func (jf int, elf *colly.HTMLElement) {
          Fares.Stops = nil
          if jf == j {
            Fares.Train = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(elf.ChildText("li.transport div"), "[train]", "【電車】"), "[bus]", "【バス】"), "[air]", "【空路】")
            Fares.Platform = elf.ChildText("li.platform")
            Fares.Color = strings.ReplaceAll(elf.ChildAttr("span", "style"), "border-color:#", "")
            elf.ForEach("li.stop ul", func (js int, els *colly.HTMLElement) {
              Stops.Time = els.ChildText("li dl dt")
              Stops.Name = strings.ReplaceAll(els.ChildText("li dl dd"), "○", "")
              Fares.Stops = append(Fares.Stops, Stops)
            })
            Stations.Fares = append(Stations.Fares, Fares)
          }
        })
        e.ForEach("div.routeDetail div.walk ul.info", func (jw int, elw *colly.HTMLElement) {
          if jw == j {
            Fares.Train = strings.ReplaceAll(elw.ChildText("li.transport"), "[line][walk]", "")
            Fares.Platform = ""
            Fares.Color = "a8a8a8"
            Stations.Fares = append(Stations.Fares, Fares)
          }
        })
        Routes.Stations = append(Routes.Stations, Stations)
      })

      routeArr = append(routeArr, Routes)
    })
  }

  sc.Visit(gurl)
  sc.Wait()

  return routeArr
}
