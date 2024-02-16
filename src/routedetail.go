package src

import (
  "strings"

  "github.com/gocolly/colly"
)

func getRouteDetail(e *colly.HTMLElement) Route {
  r := Route{}
  e.ForEach("div.routeSummary div ul.priority li span", func (j int, el *colly.HTMLElement) {
    if el.Attr("class") == "icnPriTime" {
      r.Badges = append(r.Badges, 1)
    }
    if el.Attr("class") == "icnPriFare" {
      r.Badges = append(r.Badges, 2)
    }
    if el.Attr("class") == "icnPriTrans" {
      r.Badges = append(r.Badges, 3)
    }
  })
  base := e.ChildText("ul.summary li.time span")
  time := strings.ReplaceAll(base, e.ChildText("ul.summary li.time span.small"), "")
  time2 := strings.Split(time, "着")
  r.Time = time2[0] + "着"
  durabase := e.ChildText("ul.summary li.time")
  durasi := strings.Index(durabase, "着") + len("着")
  duraei := strings.Index(durabase[durasi:], "分") + len("分") + durasi

  r.Duration = durabase[durasi:duraei]
  r.TransitCunt = strings.ReplaceAll(e.ChildText("ul.summary li.transfer"), "乗換：", "")
  r.Fare = strings.ReplaceAll(e.ChildText("ul.summary li.fare"), "[priic]IC優先：", "")
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
        Fares.Train = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(elf.ChildText("li.transport div"), "[train]", "【電車】"), "[bus]", "【バス】"), "[air]", "【空路】"), "当駅始発", "【当駅始発】") + "　"

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
    r.Stations = append(r.Stations, Stations)
  })

  return r
}
