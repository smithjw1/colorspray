package dice

import (
	"errors"
  "strconv"
	"strings"
)

type Dice struct {
	Num int
  Modifier int
	Die Die
}

type RollData struct {
	Total int
  Natural int
  Modifier int
	Dicevals []int
}

func (d Dice) Roll() (RollData) {
  var theRoll RollData
	r := 0
  diceval := make([]int, d.Num)
	for i := 0; i < d.Num; i++ {
		val := d.Die.Roll()
		r += val
    diceval[i] += val
	}
  theRoll.Total = r + d.Modifier
  theRoll.Natural = r
  theRoll.Modifier = d.Modifier
  theRoll.Dicevals = diceval
	return theRoll
}

func Parse(notation string) (d *Dice, err error) {
  var num int
  var modifier int
	var faces int

  dc := strings.Split(notation, "d")
  if len(dc) == 2 {
    s := dc[0]
		if s == "" {
			num = 1
		} else {
			num, err = strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
		}
    facepart := strings.Split(dc[1], "+")
    if len(facepart) == 2 {
      faces, err = strconv.Atoi(facepart[0])
		  modifier, err = strconv.Atoi(facepart[1])
		  if err != nil {
			     return nil, err
		  }
    } else {
      modifier = 0
      faces, err = strconv.Atoi(dc[1])
		  if err != nil {
			     return nil, err
		  }
    }
  } else {
    return nil, errors.New("bad notation \"" + notation + "\"")
  }

  d = &Dice{
		Num: num,
    Modifier: modifier,
		Die: Die{Faces: faces},
	}
	return d, nil
}
