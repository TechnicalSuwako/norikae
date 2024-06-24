package main

import (
  "fmt"
  "os"
  "time"

  "076/norikae/src"
)

var sofname = "norikae"
var version = "1.1.0"
var avalopt = "ABEfFjmnrStX"

func usage() {
  fmt.Printf("%s-%s\nusage: %s [-%s] [string]\n", sofname, version, sofname, avalopt)
}

func main() {
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
    usage()
    return
  }

  var foundf, foundt bool

  for _, v := range os.Args {
    if (v == "-f") { foundf = true }
    if (v == "-t") { foundt = true }
  }

  if !foundf || !foundt {
    usage()
    return
  }

  for i := 1; i < len(os.Args); i++ {
    if os.Args[i] == "-f" { opts.From  = os.Args[i+1] }
    if os.Args[i] == "-t" { opts.To    = os.Args[i+1] }
    if os.Args[i] == "-n" { opts.Date  = os.Args[i+1] }
    if os.Args[i] == "-j" { opts.Time  = os.Args[i+1] }
    if os.Args[i] == "-m" { opts.Mode  = os.Args[i+1] }
    if os.Args[i] == "-r" { opts.Route = os.Args[i+1] }

    if os.Args[i] == "-A" {
      opts.NoAirplane = false
    }
    if os.Args[i] == "-S" {
      opts.NoShinkansen = false
    }
    if os.Args[i] == "-E" {
      opts.NoExpress = false
    }
    if os.Args[i] == "-X" {
      opts.NoExpressBus = false
    }
    if os.Args[i] == "-B" {
      opts.NoBus = false
    }
    if os.Args[i] == "-F" {
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
