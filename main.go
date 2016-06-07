package main

import (
	"fmt"
  "github.com/colorspray/dice"
)


func main() {
  dc, err := dice.Parse("1d20");
  if err == nil {
    fmt.Println(dc.Roll())
  }
}
