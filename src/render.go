package src

import (
  "fmt"

  "github.com/gookit/color"
)

func Render (route []Route) {
  var r, g, b uint8
  col := "fcfcfc"
  fmt.Sscanf("ff7e56", "%2x%2x%2x", &r, &g, &b)
  b1 := color.RGB(uint8(r), uint8(g), uint8(b)).Sprint("早")
  fmt.Sscanf("60bddb", "%2x%2x%2x", &r, &g, &b)
  b2 := color.RGB(uint8(r), uint8(g), uint8(b)).Sprint("楽")
  fmt.Sscanf("fab60a", "%2x%2x%2x", &r, &g, &b)
  b3 := color.RGB(uint8(r), uint8(g), uint8(b)).Sprint("安")
  fmt.Sscanf(col, "%2x%2x%2x", &r, &g, &b)
  c := color.RGB(uint8(r), uint8(g), uint8(b)).Sprint("")

  for key, value := range route {
    color.Style{color.FgBlack, color.BgMagenta, color.OpBold}.Println("# ルート" + fmt.Sprintf("%d", key+1))
    badges := ""
    for _, badge := range value.Badges {
      if badge == 1 {
        badges += "〈" + b1 + "〉"
      }
      if badge == 2 {
        badges += "〈" + b2 + "〉"
      }
      if badge == 3 {
        badges += "〈" + b3 + "〉"
      }
    }
    fmt.Println(value.Time + " (" + value.Duration + "), " + value.Fare + ", 乗換数：" + value.TransitCunt + " " + badges)
    for k, v := range value.Stations {
      fmt.Println(v.Time + " " + v.Name)
      for i, fare := range v.Fares {
        if k == i {
          col = fare.Color
          fmt.Sscanf(col, "%2x%2x%2x", &r, &g, &b)
          text := fare.Train
          if fare.Platform != "" { text += "\n" + fare.Platform }
          c = color.RGB(uint8(r), uint8(g), uint8(b)).Sprint(text)
      	  fmt.Println(c)
        }
      }
    }
    fmt.Println("")
  }
}
