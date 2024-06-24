package src

import (
  "fmt"
)

var r, g, b uint8
var col string

func getFares(v Station, k int) {
	for i, fare := range v.Fares {
		if k != i {
			continue
		}

		col = fare.Color
		fmt.Sscanf(col, "%2x%2x%2x", &r, &g, &b)
		text := fare.Train

		if fare.Platform != "" {
			text += "\n" + fare.Platform
		}

		c := fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, text)
		fmt.Println(c)
	}
}

func Render (route []Route) {
  col = "fcfcfc"

	fmt.Sscanf("ff7e56", "%2x%2x%2x", &r, &g, &b)
	b1 := fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, "早")

	fmt.Sscanf("60bddb", "%2x%2x%2x", &r, &g, &b)
	b2 := fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, "楽")

	fmt.Sscanf("fab60a", "%2x%2x%2x", &r, &g, &b)
	b3 := fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, "安")

  fmt.Sscanf(col, "%2x%2x%2x", &r, &g, &b)

  for key, value := range route {
		fmt.Printf("\x1b[1;35m# ルート%d\x1b[0m\n", key+1)

    badges := ""
    for _, badge := range value.Badges {
      if badge == 1 { badges += "〈" + b1 + "〉" }
      if badge == 2 { badges += "〈" + b2 + "〉" }
      if badge == 3 { badges += "〈" + b3 + "〉" }
    }

    fmt.Println(
      value.Time + " (" + value.Duration + "), " +
      value.Fare + ", 乗換数：" + value.TransitCunt + " " + badges,
    )

    for k, v := range value.Stations {
      fmt.Println(v.Time + " " + v.Name)
      getFares(v, k)
    }

    fmt.Println("")
  }
}
