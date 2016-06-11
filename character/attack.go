package character

import (
  "github.com/smithjw1/colorspray/dice"
)

type Attack struct {
	Base int
  Ability int
  SizeMod int
  RangePen int
  Damage string
  MinCrit int
  CritMulti int
  OppAC int
}

type AttackResult struct {
  Hit bool
  AttackRoll dice.RollData
  Crit bool
  CritRoll dice.RollData
  Damage int
  DamageRolls []dice.RollData
}

func (a Attack) Make() (*AttackResult) {

  hit := false
  crit := false
  var cR dice.RollData
  var dR []dice.RollData
  var dam int

  d := &dice.Dice{
    Num: 1,
    Modifier: a.Base + a.Ability + a.SizeMod - a.RangePen,
		Die: dice.Die{Faces: 20},
  }

  aR := d.Roll()

  if aR.Total > a.OppAC {
    hit = true

    dD, _ := dice.Parse(a.Damage)
    dR = append(dR, dD.Roll())
    dam += dR[0].Total
    a.CritMulti -= 1

    if aR.Total >= a.MinCrit {
      cR = d.Roll()
      if cR.Total > a.OppAC {
        crit = true
        for i := 0; i < a.CritMulti; i++ {
          mR := dD.Roll()
      		dR = append(dR, mR)
          dam += mR.Total
      	}
      }
    }

  }

  theAttack := &AttackResult {
    Hit: hit,
    AttackRoll: aR,
    Crit: crit,
    CritRoll: cR,
    Damage: dam,
    DamageRolls: dR,
  }

  return theAttack
}
