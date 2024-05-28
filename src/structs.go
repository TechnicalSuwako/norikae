package src

type (
  Opts struct {
    From, To, Date, Time, Mode, Route string
    NoAirplane, NoShinkansen, NoExpress, NoExpressBus, NoBus, NoFairy bool
  }
  Route struct {
    Time, Duration, TransitCunt, Fare string
    Badges []int
    Stations []Station
  }
  Station struct {
    Time, Name string
    Fares []Fare
  }
  Fare struct {
    Train, Color, Platform string
    Stops []Stop
  }
  Stop struct {
    Time, Name string
  }
)
