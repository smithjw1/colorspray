package dice

type dice struct {
	num int
	die die
}

func (d dice) Roll() int {
	r := 0
	for i := 0; i < d.num; i++ {
		val := d.die.Roll()
		r += val
	}
	return r;
}

func Parse(notation string) (d *dice, err error) {
  d = &dice{
		num: 1,
		die: die{faces: 20},
	}
	return d, nil
}
