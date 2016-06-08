package dice

import (
	"errors"
  "strconv"
	"strings"
)

type dice struct {
	num int
  modifier int
	die die
}

func (d dice) Roll() (int, int, []int) {
	r := 0
  diceval := make([]int, d.num)
	for i := 0; i < d.num; i++ {
		val := d.die.Roll()
		r += val
    diceval[i] += val
	}
  total := r + d.modifier
	return total, r, diceval;
}

func Parse(notation string) (d *dice, err error) {
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

  d = &dice{
		num: num,
    modifier: modifier,
		die: die{faces: faces},
	}
	return d, nil
}
