package main

import (
    "log"
	"strconv"
	"io/ioutil"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Pastel struct {
	Id int64 
    Sabor string 
    Quantidade int32
    Valor float32 
}

var Pasteis []Pastel

func criarPastel(w http.ResponseWriter, r *http.Request) {
	var pastel Pastel 
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.WriteHeader(http.StatusCreated)

    json.Unmarshal(reqBody, &pastel)

	Pasteis = append(Pasteis, pastel)

	json.NewEncoder(w).Encode(pastel)
}

func listarPasteis(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Pasteis)
}

func buscarPastelPorId(w http.ResponseWriter, r *http.Request){
	var id = pegarId(r)

    for _, pastel := range Pasteis {
        if pastel.Id == id {
            json.NewEncoder(w).Encode(pastel)
        }
    }
}

func atualizarPastel(w http.ResponseWriter, r *http.Request) {
	var pastelAtualizado Pastel 
	reqBody, _ := ioutil.ReadAll(r.Body)

	var id = pegarId(r)

	json.Unmarshal(reqBody, &pastelAtualizado)

	for index, pastel := range Pasteis {
        if pastel.Id == id {
			Pasteis[index].Sabor = pastelAtualizado.Sabor
			Pasteis[index].Quantidade = pastelAtualizado.Quantidade
			Pasteis[index].Valor = pastelAtualizado.Valor
        }
    }
}

func apagarPastel(w http.ResponseWriter, r *http.Request) {
	var id = pegarId(r)
	w.WriteHeader(http.StatusNoContent)

    for index, pastel := range Pasteis {
        if pastel.Id == id {
            Pasteis = append(Pasteis[:index], Pasteis[index+1:]...)
        }
    }
}

func handleRequests() {
	appRouter := mux.NewRouter().StrictSlash(true)

	appRouter.HandleFunc("/pasteis", criarPastel).Methods("POST")
	appRouter.HandleFunc("/pasteis", listarPasteis).Methods("GET")
	appRouter.HandleFunc("/pasteis/{id}", buscarPastelPorId).Methods("GET")
	appRouter.HandleFunc("/pasteis/{id}", atualizarPastel).Methods("PUT")
	appRouter.HandleFunc("/pasteis/{id}", apagarPastel).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", appRouter))
}

func pegarId(r *http.Request) int64 {
    vars := mux.Vars(r)
    key := vars["id"]
	id, err := strconv.ParseInt(key, 10, 64)

	if err != nil {
		panic(err)
	}

	return id
}

func main() {
	Pasteis = []Pastel{
        Pastel{Id: 1, Sabor: "Carne", Quantidade: 10, Valor: 7.5},
        Pastel{Id: 2 ,Sabor: "Queijo", Quantidade: 15, Valor: 5.5},
    }
    handleRequests()
}
