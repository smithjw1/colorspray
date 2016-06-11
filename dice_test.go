package main

import (
  "testing"
  "github.com/smithjw1/colorspray/dice"
)

func TestOneDTwenty(t *testing.T) {
  t.Log("1d20 (expected score: between 1 and 20)")
  notation := "1d20"
  dc, err := dice.Parse(notation);
  if err == nil {
    var r dice.RollData

    r = dc.Roll()
    if r.Total > 20 || r.Total < 1 {
      t.Errorf("Total out of range %d", r.Total)
    }
    if r.Natural > 20 || r.Natural < 1 {
      t.Errorf("Natural out of range %d", r.Natural)
    }
    if r.Natural != r.Total {
      t.Errorf("Natural is %d total is %d they should be equal", r.Natural, r.Total)
    }
    if len(r.Dicevals) > 1 {
      t.Errorf("More dice rolls than exepected %d", len(r.Dicevals))
    }

  } else {
    t.Errorf("Could not parse notation %s", notation)
  }
}


func TestNoNumber(t *testing.T) {
  t.Log("d20 (expected score: between 1 and 20)")
  notation := "d20"
  dc, err := dice.Parse(notation);
  if err == nil {
    var r dice.RollData

    r = dc.Roll()
    if r.Total > 20 || r.Total < 1 {
      t.Errorf("Total out of range %d", r.Total)
    }
    if r.Natural > 20 || r.Natural < 1 {
      t.Errorf("Natural out of range %d", r.Natural)
    }
    if r.Natural != r.Total {
      t.Errorf("Natural is %d total is %d they should be equal", r.Natural, r.Total)
    }
    if len(r.Dicevals) > 1 {
      t.Errorf("More dice rolls than exepected %d", len(r.Dicevals))
    }

  } else {
    t.Errorf("Could not parse notation %s", notation)
  }
}

func TestPositiveModifier(t *testing.T) {
  t.Log("3d1+7 (expected score: 10)")
  notation := "3d1+7"
  dc, err := dice.Parse(notation);
  if err == nil {
    var r dice.RollData
    r = dc.Roll()
    if r.Total != 10 {
      t.Errorf("Total out of range %d", r.Total)
    }
    if r.Natural != 3 {
      t.Errorf("Natural out of range %d", r.Natural)
    }
    if len(r.Dicevals) != 3 {
      t.Errorf("Different dice rolls than exepected %d", len(r.Dicevals))
    }
  } else {
    t.Errorf("Could not parse notation %s", notation)
  }
}

func TestNegativeModifier(t *testing.T) {
  t.Log("3d1+7 (expected score: 10)")
  notation := "4d1+-3"
  dc, err := dice.Parse(notation);
  if err == nil {
    var r dice.RollData
    r = dc.Roll()
    if r.Total != 1 {
      t.Errorf("Total out of range %d", r.Total)
    }
    if r.Natural != 4 {
      t.Errorf("Natural out of range %d", r.Natural)
    }
    if len(r.Dicevals) != 4 {
      t.Errorf("Different dice rolls than exepected %d", len(r.Dicevals))
    }
  } else {
    t.Errorf("Could not parse notation %s", notation)
  }
}
