package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../entity"
	"../github.com/gorilla/mux"
	"../tiny"
)

var r_sim = [7][4]float64{
	{0, 0, 0, 0},
	{5.5, 7, 8, 14},
	{1.25, 1.4, 1.5, 1.6},
	{12.3, 12.4, 12.5, 12.6},
	{55, 57, 63, 65},
	{9, 10, 11, 12},
	{0, 0, 0, 0},
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	if key == "favicon.ico" {
		return
	}
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(key)

	art := tiny.Find(id)
	//fmt.Fprintf(w /*art.Funda+*/, "%f\nPH record:%v\nSoil Humidity record:%v\nSoil Temperature record:%v\nAir Humidity record:%v\nAir Temperature record:%v\nLight record:%v\nResilience record:%v", tiny.Rate(art, r_sim), art.PH, art.SH, art.ST, art.AH, art.AT, art.L, art.R)
	tiny.Uploadchart(w, r, art.PH, "pH")
	tiny.Uploadchart(w, r, art.L, "Light")
	tiny.Uploadchart(w, r, art.SH, "Soil Humidity")
	tiny.Uploadchart(w, r, art.ST, "Soil Temperature")
	tiny.Uploadchart(w, r, art.AH, "Air Humidity")
	tiny.Uploadchart(w, r, art.AT, "Air Temperature")
	tiny.Uploadchart(w, r, art.R, "Resilience")
}
func Read(id int) entity.Display_rate {
	art := tiny.Find(id)
	agri_rate := entity.Display_rate{Rate: tiny.Rate(art, r_sim), PH: art.PH, L: art.L, SH: art.SH, ST: art.ST, AH: art.AH, AT: art.AT, R: art.R}
	return agri_rate
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The website is successfully opened")
}
func proc(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var id entity.ID
	var rate entity.Display_rate

	decoder.Decode(&id)

	rate = Read(id.ID)

	fmt.Println(rate)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(rate); err != nil {
		panic(err)
	}
}
func trigger(w http.ResponseWriter, request *http.Request) {
	tiny.Trig()
}
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/proc", proc)
	myRouter.HandleFunc("/trial", trigger)
	myRouter.HandleFunc("/{id}", returnSingleArticle)
	myRouter.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
