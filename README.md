# CalVer
> small utility to handle Calendar Versioning (https://calver.org/)

# Usage
### API
```go
import (
  "fmt"
  
  "github.com/umayr/calver"
)

func main() {
  c, err := calver.New("YYYY.MM.DD", "dev")
  if err != nil {
    panic(err)
  }
  
  fmt.Println(c) // YYYY.MM.DD
  r := c.Release()
  fmt.Println(r) // 2020.12.20
  
  r = c.PreRelease()
  fmt.Println(r) // 2020.12.20-dev.1
  
  r = c.PreRelease()
  fmt.Println(r) // 2020.12.20-dev.2
  
  r = c.Release()
  fmt.Println(r) // 2020.12.20-2
  
  // different day
  
  r = c.Release()
  fmt.Println(r) // 2020.12.21
  
  p, err := calver.Parse("2020.12.20-dev.2", "YYYY.MM.DD", "dev")
  if err != nil {
    panic(err)
  }
  
  fmt.Println(p.PreRelease()) // 2020.12.20-dev.3
  fmt.Println(p.Release()) // 2020.12.20-3
}

```

Available segments:
```go
const (
  FullYear = "YYYY"
  ShortYear = "YY"
  PaddedYear = "0Y"
  ShortMonth = "MM"
  PaddedMonth = "0M"
  ShortWeek = "WW"
  PaddedWeek = "0W"
  ShortDay = "DD"
  PaddedDay = "0D"
)
```

You can build any format using these segments, one thing to note that you need to have at lease two parts `major` and `minor` (for instance `YYYY.0W`) for a format to be valid, at max you have three segments (`major`, `minor` and `micro`)

### CLI

```bash
# install it using go
λ go get github.com/umayr/calver/cmd/calver

# build it from source
λ git clone https://github.com/umayr/calver
λ cd calver
λ make

λ calver --help
calver is a small utility to handle calender versioning:

Usage of calver:
  -format string
    	format to parse the provided version (default "YYYY.MM.DD")
  -modifier string
    	modifier for prerelease versions (default "dev")
  -pre-release
    	flag to create a prerelease

For more information on Calender Versioning: https://calver.org

λ calver 2019.1.1
2020.12.20

λ calver --pre-release 2019.1.1
2020.12.20-dev

λ calver 2020.12.20-dev
2020.12.20
```
