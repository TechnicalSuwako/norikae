package src

import (
  "strings"
  "net/url"
)

func b2s (val bool) string {
  if val { return "1" }
  return "0"
}

func GetUrl (f Opts) (string, error) {
  date := strings.Split(f.Date, "-")
  year := date[0]
  month := date[1]
  day := date[2]
  hour := strings.Split(f.Time, ":")[0]
  minute := strings.Split(f.Time, ":")[1]
  m1 := string(minute[0])
  m2 := string(minute[1])

  curl, err := url.Parse(
    "https://transit.yahoo.co.jp/search/result" +
    "?from=" + url.QueryEscape(f.From) +
    "&to=" + url.QueryEscape(f.To) +
    "&y=" + year +
    "&m=" + month +
    "&d=" + day +
    "&hh=" + hour +
    "&m1=" + m1 +
    "&m2=" + m2 +
    "&type=" + f.Mode +
    "&ticket=ic&expkind=1&userpass=1&ws=" + f.Route +
    "&al=" + b2s(f.NoAirplane) +
    "&shin=" + b2s(f.NoShinkansen) +
    "&ex=" + b2s(f.NoExpress) +
    "&hb=" + b2s(f.NoExpressBus) +
    "&lb=" + b2s(f.NoBus) +
    "&sr=" + b2s(f.NoFairy))
  if err != nil {
    return "", nil
  }

  return curl.String(), nil
}
