package src

import (
  "strings"

  "github.com/gocolly/colly"
)

func getSummary(i int, e *colly.HTMLElement) Route {
  r := Route{}
  if e.Attr("class") == "icnPriTime"  { r.Badges = append(r.Badges, 1) }
  if e.Attr("class") == "icnPriFare"  { r.Badges = append(r.Badges, 2) }
  if e.Attr("class") == "icnPriTrans" { r.Badges = append(r.Badges, 3) }

  return r
}

func handleFare(el *colly.HTMLElement, f Fare, s Stop) Fare {
  fixTrain := strings.ReplaceAll(
    el.ChildText("li.transport div"), "[train]", "【電車】",
  )
  fixBus := strings.ReplaceAll(fixTrain, "[bus]", "【バス】")
  fixAir := strings.ReplaceAll(fixBus, "[air]", "【空路】")
  fixEki := strings.ReplaceAll(fixAir, "当駅始発", "【当駅始発】")

  f.Train = fixEki
  f.Platform = el.ChildText("li.platform")
  f.Color = strings.ReplaceAll(el.ChildAttr("span", "style"), "border-color:#", "")
  el.ForEach("li.stop ul", func (js int, els *colly.HTMLElement) {
    s.Time = els.ChildText("li dl dt")
    s.Name = strings.ReplaceAll(els.ChildText("li dl dd"), "○", "")
    f.Stops = append(f.Stops, s)
  })

  return f
}

func handleWalk(el *colly.HTMLElement, f Fare) Fare {
  f.Train = strings.ReplaceAll(
    el.ChildText("li.transport"), "[line][walk]", "",
  )
  f.Platform = ""
  f.Color = "a8a8a8"

  return f
}

func getRouteDetail(e *colly.HTMLElement) Route {
  r := Route{}
  onDivs := "div.routeSummary div ul.priority li span"
  e.ForEach(onDivs, func(i int, el *colly.HTMLElement) {
    summary := getSummary(i, el)
    r.Badges = append(r.Badges, summary.Badges...)
  })

  base := e.ChildText("ul.summary li.time span")
  time := strings.ReplaceAll(base, e.ChildText("ul.summary li.time span.small"), "")
  time2 := strings.Split(time, "着")
  r.Time = time2[0] + "着"
  durabase := e.ChildText("ul.summary li.time")
  durasi := strings.Index(durabase, "着") + len("着")
  duraei := strings.Index(durabase[durasi:], "分") + len("分") + durasi

  r.Duration = durabase[durasi:duraei]
  r.TransitCunt = strings.ReplaceAll(
    e.ChildText("ul.summary li.transfer"), "乗換：", "",
  )
  r.Fare = strings.ReplaceAll(
    e.ChildText("ul.summary li.fare"), "[priic]IC優先：", "",
  )

  onDivs = "div.routeDetail div.station"
  e.ForEach(onDivs, func (j int, el *colly.HTMLElement) {
    station := Station{}
    station.Time = el.ChildText("ul.time li")
    if el.ChildText("p.icon span") == "[dep]" { station.Time += "発" }
    if el.ChildText("p.icon span") == "[arr]" { station.Time += "着" }
    station.Name = el.ChildText("dl dt a")

    fares := []Fare{}
    onDivs = "div.routeDetail div.fareSection div.access"
    e.ForEach(onDivs, func (jf int, elf *colly.HTMLElement) {
      fare := Fare{}
      fare = handleFare(elf, fare, Stop{})
      fares = append(fares, fare)
    })

    onDivs = "div.routeDetail div.walk ul.info"
    e.ForEach(onDivs, func (jw int, elw *colly.HTMLElement) {
      fare := Fare{}
      fare = handleWalk(elw, fare)
      fares = append(fares, fare)
    })

    station.Fares = fares
    r.Stations = append(r.Stations, station)
  })

  return r
}
