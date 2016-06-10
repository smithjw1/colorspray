package dice

import (
  "time"
	"math/rand"
)

type Die struct {
	Faces int
}

func (d Die) Roll() int {
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
	return r1.Intn(d.Faces) + 1
}
