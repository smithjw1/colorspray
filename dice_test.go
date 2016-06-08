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
    var theroll Roll
    theroll.Notation = notation
    theroll.Total, theroll.Natural, theroll.Dicevals = dc.Roll()
    if theroll.Total > 20 || theroll.Total < 1 {
      t.Errorf("Total out of range %d", theroll.Total)
    }
    if theroll.Natural > 20 || theroll.Natural < 1 {
      t.Errorf("Natural out of range %d", theroll.Natural)
    }
    if theroll.Natural != theroll.Total {
      t.Errorf("Natural is %d total is %d they should be equal", theroll.Natural, theroll.Total)
    }
    if len(theroll.Dicevals) > 1 {
      t.Errorf("More dice rolls than exepected %d", len(theroll.Dicevals))
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
    var theroll Roll
    theroll.Notation = notation
    theroll.Total, theroll.Natural, theroll.Dicevals = dc.Roll()
    if theroll.Total > 20 || theroll.Total < 1 {
      t.Errorf("Total out of range %d", theroll.Total)
    }
    if theroll.Natural > 20 || theroll.Natural < 1 {
      t.Errorf("Natural out of range %d", theroll.Natural)
    }
    if theroll.Natural != theroll.Total {
      t.Errorf("Natural is %d total is %d they should be equal", theroll.Natural, theroll.Total)
    }
    if len(theroll.Dicevals) > 1 {
      t.Errorf("More dice rolls than exepected %d", len(theroll.Dicevals))
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
    var theroll Roll
    theroll.Notation = notation
    theroll.Total, theroll.Natural, theroll.Dicevals = dc.Roll()
    if theroll.Total != 10 {
      t.Errorf("Total out of range %d", theroll.Total)
    }
    if theroll.Natural != 3 {
      t.Errorf("Natural out of range %d", theroll.Natural)
    }
    if len(theroll.Dicevals) != 3 {
      t.Errorf("Different dice rolls than exepected %d", len(theroll.Dicevals))
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
    var theroll Roll
    theroll.Notation = notation
    theroll.Total, theroll.Natural, theroll.Dicevals = dc.Roll()
    if theroll.Total != 1 {
      t.Errorf("Total out of range %d", theroll.Total)
    }
    if theroll.Natural != 4 {
      t.Errorf("Natural out of range %d", theroll.Natural)
    }
    if len(theroll.Dicevals) != 4 {
      t.Errorf("Different dice rolls than exepected %d", len(theroll.Dicevals))
    }
  } else {
    t.Errorf("Could not parse notation %s", notation)
  }
}
