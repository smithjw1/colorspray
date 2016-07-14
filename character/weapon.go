package character

type Weapon struct {
	Damage string
	Bonus int
	MinCrit int
	CritMulti int
	Wrange int
	Ammo int
	Loaded int
}

func GetWeapon(id string) (*Weapon) {

	tW := &Weapon {
    Damage: "2d4",
    Bonus: 2,
    MinCrit: 19,
		CritMulti: 2,
    Wrange: 10,
    Ammo: -1,
    Loaded: -1,
  }

  return tW
}
