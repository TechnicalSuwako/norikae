package main

import (
  "fmt"
  "os"
  "time"
  "076/norikae/src"
)

var sofname = "norikae"
var version = "1.1.0"

func help () {
  fmt.Printf("０７６ 乗換 - %s-%s - CLIでの路線情報\n", sofname, version)
  fmt.Println("https://076.moe/ | https://gitler.moe/suwako/norikae")
  fmt.Println("\n使い方：")
  fmt.Println("-v\n\tバージョンを表示")
  fmt.Println("-h\n\tヘルプを表示")
  fmt.Println("\n【必須のオプション】")
  fmt.Println("-f [出発駅]\n\t例：秋葉原、渋谷、大手町（東京）")
  fmt.Println("-t [到着駅]\n\t例：秋葉原、渋谷、大手町（東京）")
  fmt.Println("\n【任意のオプション】")
  fmt.Println("-n [YYYY-MM-DD]\n\t例：2023-05-02（デフォルト：今）")
  fmt.Println("-j [HH:MM]\n\t例：18:45（デフォルト：今）")
  fmt.Println("-m [0〜4]\n\t" +
    "0 = 出発、1 = 指定なし、2 = 終電、3 = 始発、4 = 到着（デフォルト：0）")
  fmt.Println("-r [0〜2]\n\t" +
    "0 = 到着が早い順、1 = 料金が高い順、2 = 乗り換え回数順（デフォルト：0）")
  fmt.Println("\n--no-airplane\n\t空路を省く")
  fmt.Println("--no-shinkansen\n\t新幹線を省く")
  fmt.Println("--no-express\n\t有料特急を省く")
  fmt.Println("--no-expressbus\n\t高速バスを省く")
  fmt.Println("--no-bus\n\t路線/連絡バスを省く")
  fmt.Println("--no-ferry\n\tフェリーを省く")
  fmt.Println("\n例： " + sofname + " -f 秋葉原 -t 渋谷 -j 16:23 -m 4 --no-bus")
}

func main () {
  var opts src.Opts
  // デフォルトな値
  t := time.Now()
  opts.Date         = t.Format("2006-01-02")
  opts.Time         = t.Format("15:04")
  opts.Mode         = "0"
  opts.Route        = "0"
  opts.NoAirplane   = true
  opts.NoShinkansen = true
  opts.NoExpress    = true
  opts.NoExpressBus = true
  opts.NoBus        = true
  opts.NoFairy      = true

  if len(os.Args) == 1 {
    help()
    return
  }

  var foundf, foundt bool

  for _, v := range os.Args {
    if (v == "-f") { foundf = true }
    if (v == "-t") { foundt = true }
    if (v == "-v") {
      fmt.Printf("%s-%s\n", sofname, version)
      return
    }
    if (v == "-h") {
      help()
      return
    }
  }

  if !foundf || !foundt {
    help()
    return
  }

  for i := 1; i < len(os.Args); i++ {
    if os.Args[i] == "-f" { opts.From  = os.Args[i+1] }
    if os.Args[i] == "-t" { opts.To    = os.Args[i+1] }
    if os.Args[i] == "-n" { opts.Date  = os.Args[i+1] }
    if os.Args[i] == "-j" { opts.Time  = os.Args[i+1] }
    if os.Args[i] == "-m" { opts.Mode  = os.Args[i+1] }
    if os.Args[i] == "-r" { opts.Route = os.Args[i+1] }

    if os.Args[i] == "--no-airplane" {
      opts.NoAirplane = false
    }
    if os.Args[i] == "--no-shinkansen" {
      opts.NoShinkansen = false
    }
    if os.Args[i] == "--no-express" {
      opts.NoExpress = false
    }
    if os.Args[i] == "--no-expressbus" {
      opts.NoExpressBus = false
    }
    if os.Args[i] == "--no-bus" {
      opts.NoBus = false
    }
    if os.Args[i] == "--no-ferry" {
      opts.NoFairy = false
    }
  }

  gurl, err := src.GetUrl(opts)
  if err != nil {
    fmt.Println(err)
    return
  }

  route := src.Scrape(gurl)
  src.Render(route)
}
