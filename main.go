package main

import (
	"fmt"
  "github.com/smithjw1/colorspray/dice" // started from: https://github.com/narqo/go-dice
  "github.com/smithjw1/colorspray/character"
  "encoding/json"
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "github.com/satori/go.uuid"
  "io/ioutil"
)

func main() {

    u1 := uuid.NewV4()
    fmt.Printf("UUIDv4: %s\n", u1)

    client := &http.Client{}
    req, _ := http.NewRequest("GET", "https://api.airtable.com/v0/appiaQ2jqi3R7oHsA/Characters?maxRecords=1&view=Main%20View&filterByFormula=Name%3D%226922a942-6482-42d2-8bd7-2598101abf6c%22", nil)
    req.Header.Set("Authorization", "Bearer keyC8IAYh8YbF4PcG")
    res, err := client.Do(req)

    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
	  if err != nil {
        panic(err.Error())
    }
    fmt.Printf("%s", body)


    type Fields struct {
      Name string
      AC int
    }
    type Records struct {
      id string
      Fields []Fields
      createdTime string
    }
    type Characters struct {
      Records []Records
    }

    var data Characters

    json.Unmarshal(body, &data)
    fmt.Printf("Results: %v\n", data)

    return
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
