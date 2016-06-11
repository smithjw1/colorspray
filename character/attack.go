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

type AttackResult struct {
  Hit bool
  AttackRoll dice.RollData
  Crit bool
  CritRoll dice.RollData
  Damage dice.RollData
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

func (a Attack) Make() (*AttackResult) {

  hit := false
  crit := false
  var cR dice.RollData
  var dR dice.RollData

  d := &dice.Dice{
    Num: 1,
    Modifier: a.Base + a.Ability,
		Die: dice.Die{Faces: 20},
  }

  aR := d.Roll()

  if aR.Total > a.OppAC {
    hit = true

    dD, _ := dice.Parse(a.Damage)

    dD.Modifier = a.Base + a.Ability

    dR = dD.Roll()

    if aR.Total >= a.MinCrit {
      cR = d.Roll()
      if cR.Total > a.OppAC {
        crit = true
        dR.Total += dR.Natural
        dR.Natural = dR.Natural * 2
      }
    }

  }

  theAttack := &AttackResult {
    Hit: hit,
    AttackRoll: aR,
    Crit: crit,
    CritRoll: cR,
    Damage: dR,
  }

  return theAttack
}
