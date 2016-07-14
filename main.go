package main

import (
	"fmt"
  "github.com/smithjw1/colorspray/dice" // started from: https://github.com/narqo/go-dice
  "github.com/smithjw1/colorspray/character"
	"github.com/smithjw1/colorspray/csdb"
  "encoding/json"
  "github.com/gorilla/mux"
  "net/http"
  "os"
	"strconv"
	"log"

)

type CSError struct {
	Message string
}


func main() {
	fmt.Printf("%v\n", "clashing colors")

	f, err := os.OpenFile("colorspray.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	db := csdb.Open()
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS people (id UUID PRIMARY KEY, data jsonb)")
	err != nil {
		log.Printf("Error creating database table: %q", err)
	}

  r := mux.NewRouter()
  r.HandleFunc("/roll/{notation}", RollHandler)
  r.HandleFunc("/roll", RollHandler)
  r.HandleFunc("/attack", AttackHandler)
	r.HandleFunc("/attack/{attacker}/{weapon}/{defender}", AttackHandler)
	r.HandleFunc("/attack/{attacker}/{weapon}/{defender}/{range}", AttackHandler)
	r.HandleFunc("/character/create", CreatePerson)

  if err := http.ListenAndServe(":"+os.Getenv("PORT"), r); err != nil {
    panic(err)
  }
}

func RollHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  notation, ok := vars["notation"]
  if !ok {
    notation = "1d20"
  }

  dc, err := dice.Parse(notation);
  if err == nil {
    theRoll := dc.Roll()
    js, err := json.Marshal(theRoll)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Server", "colorspray")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(js)
    return
  }
  http.Error(w, err.Error(), http.StatusInternalServerError)
}

func AttackHandler(w http.ResponseWriter, r *http.Request) {

	var aT character.Attack
	w.Header().Set("Server", "colorspray")
  w.Header().Set("Content-Type", "application/json")


	vars := mux.Vars(r)
  a, ok := vars["attacker"]
	we,ok := vars["weapon"]
	d,ok := vars["defender"]
  if ok {
		tW := character.GetWeapon(we)

		if tW.Loaded == 0 {
			w.WriteHeader(200)
			resp := CSError {
				Message: "item is not loaded",
			}
			js, _ := json.Marshal(resp)
		  w.Write(js)
			return
		}

		if tW.Ammo == 0 {
			w.WriteHeader(200)
			resp := CSError {
				Message: "out of ammo",
			}
			js, _ := json.Marshal(resp)
		  w.Write(js)
			return
		}

		a = "44888e26-d4d9-4fad-b807-e0ccb954bea1";

		tA := character.GetPerson(a)
		tD := character.GetPerson(d)
		aT.Base = tA.Base
	  aT.SizeMod = tA.SizeMod

		rGi := 5
		rG,rok := vars["range"]
		if rok {
			rGi, _ = strconv.Atoi(rG)
			aT.Ability = tA.Abilities.Dex
		} else {
			aT.Ability = tA.Abilities.Str
		}

		p := 0
		o := rGi/tW.Wrange
		if o > 1 {
			p = o * 2
		}

	  aT.RangePen = p
	  aT.Damage  = tW.Damage
	  aT.MinCrit = tW.MinCrit
	  aT.CritMulti = tW.CritMulti
	  aT.OppAC = tD.Ac
  } else {
		decoder := json.NewDecoder(r.Body)
	  err := decoder.Decode(&aT)
	  if err != nil {
	    panic(err)
	  }
	}


  theAttack := aT.Make();

  js, err := json.Marshal(theAttack)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

	w.WriteHeader(200)
  w.Write(js)

  return
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p character.Person
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	id := p.Create()
	w.Header().Set("Server", "colorspray")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
  w.Write([]byte(id))

	return
}
