package character

import (
  "testing"
  "github.com/smithjw1/colorspray/character"
)

func TestCalcs(t *testing.T) {
  t.Log("Test deterministic attack")
  a := character.Attack{
    Base: 0,
    Ability: 0,
    Damage: "1d1",
    MinCrit: 0,
    OppAC: 0,
    CritMulti: 2,
  }

  aR := a.Make();
  if !aR.Hit {
    t.Errorf("Should have registered a hit")
  }
  if aR.AttackRoll.Total > 20 || aR.AttackRoll.Total < 1 {
    t.Errorf("Attack roll out of range %d", aR.AttackRoll.Total)
  }
  if !aR.Crit {
    t.Errorf("Should have registered a critical hit")
  }
  if aR.CritRoll.Total > 20 || aR.CritRoll.Total < 1 {
    t.Errorf("Crit roll out of range %d", aR.CritRoll.Total)
  }
  if aR.Damage != 2 {
    t.Errorf("Damage total is incorrect  %d", aR.Damage)
  }
  if len(aR.DamageRolls) != 2 {
    t.Errorf("Number of damage roll is incorrect  %d", len(aR.DamageRolls))
  }
}

func TestModifiers(t *testing.T) {
  t.Log("Test modifers attack")
  a := character.Attack{
    Base: 0,
    Ability: 0,
    SizeMod: -10,
    RangePen: 10,
    Damage: "1d1",
    MinCrit: 0,
    OppAC: 0,
    CritMulti: 2,
  }

  aR := a.Make();
  if aR.Hit {
    t.Errorf("Should not have registered a hit")
  }
  if aR.AttackRoll.Total > 0 || aR.AttackRoll.Total < -20 {
    t.Errorf("Attack roll out of range %d", aR.AttackRoll.Total)
  }
  if aR.Crit {
    t.Errorf("Should not have registered a critical hit")
  }
}
