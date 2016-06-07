package dice

import (
  "time"
	"math/rand"
)

type die struct {
	faces int
}

func (d die) Roll() int {
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
	return r1.Intn(d.faces) + 1
}
