package character

import (
	"github.com/satori/go.uuid"
	"github.com/smithjw1/colorspray/csdb"
	"log"
)

type Abilities struct {
	Str int
	Dex int
	Con int
	Intl int
	Wis int
	Cha int
}

type Person struct {
	Base int
	Abilities Abilities
	SizeMod int //pg134
	Ac int
}

func GetPerson(id string) (Person) {
	var tP Person
	if id == "a" {
		tP.Base = 2
		tP.Abilities = Abilities{Str: 2,Dex:2,Con:2,Intl:2,Wis:2,Cha:2}
		tP.SizeMod = 0
		tP.Ac = 10
	} else {
		tP.Base = 3
		tP.Abilities = Abilities{Str: 2,Dex:2,Con:2,Intl:2,Wis:2,Cha:2}
		tP.SizeMod = 0
		tP.Ac = 100
	}

  return tP
}

func (p Person) Create() (id string) {

	uID := uuid.NewV4()

	db := csdb.Open()
	stmt, err := db.Prepare("INSERT INTO people(id, data) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := stmt.Exec(uID, p)
	if err2 != nil {
		log.Fatal(err2)
	}

	return uID.String()
}
