package main

import (
//	"fmt"
  "github.com/smithjw1/colorspray/dice" // started from: https://github.com/narqo/go-dice
  "github.com/smithjw1/colorspray/character"
  "encoding/json"
  "github.com/gorilla/mux"
  "net/http"
  "os"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/roll/{notation}", RollHandler)
  r.HandleFunc("/roll", RollHandler)
  r.HandleFunc("/attack", AttackHandler)
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

  decoder := json.NewDecoder(r.Body)

  var a character.Attack
  err := decoder.Decode(&a)

  if err != nil {
    panic(err)
  }

  theAttack := a.Make();

  js, err := json.Marshal(theAttack)
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
