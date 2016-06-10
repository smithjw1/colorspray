package character

import (
  "github.com/smithjw1/colorspray/dice"
)

type Attack struct {
	Base int
  Ability int
  Damage string
  MinCrit int
  OppAC int
}

func ParseAttack() (Attack) {

  a := Attack{
    Base: 2,
    Ability: 1,
    Damage: "2d10",
    MinCrit: 19,
    OppAC: 15,
  }
  return a
}

func (a Attack) Make() (bool, int, bool, int, int) {

  hit := false
  crit := false
  cR := 0
  dR := 0

  d := &dice.Dice{
    Num: 1,
    Modifier: a.Base + a.Ability,
		Die: dice.Die{Faces: 20},
  }

  var aR int
  aR, _, _ = d.Roll()

  if aR > a.OppAC {
    hit = true

    dD, _ := dice.Parse(a.Damage)

    dD.Modifier = a.Base + a.Ability

    dR, _, _ = dD.Roll()

    if aR > a.MinCrit {
      cR, _, _ = d.Roll()
      if cR > a.OppAC {
        crit = true
        dR = dR * 2
      }
    }

  }

  return hit, aR, crit, cR, dR
}
