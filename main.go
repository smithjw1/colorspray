package main

import (
	//"fmt"
  "log"
  "github.com/smithjw1/colorspray/dice" // started from: https://github.com/narqo/go-dice
  "encoding/json"
  "github.com/gorilla/mux"
  "net/http"
)

type Roll struct {
  Notation string
	Total int
  Natural int
	Dicevals []int
}

func main() {

  r := mux.NewRouter()
  r.HandleFunc("/roll/{notation}", RollHandler)
  r.HandleFunc("/roll", RollHandler)
  if err := http.ListenAndServe("localhost:3000", nil); err != nil {
    log.Fatal("ListenAndServe: ", err)
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
    var theroll Roll
    theroll.Notation = notation
    theroll.Total, theroll.Natural, theroll.Dicevals = dc.Roll()
    js, err := json.Marshal(theroll)
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
