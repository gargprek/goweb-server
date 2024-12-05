package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"restapi.com/database"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	data, err := database.GetEmployees(nil)
	if err != nil {
		log.Println("error while getting employee data from database", "err", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	rawJson, err := json.Marshal(data)
	if err != nil {
		log.Println("error while marshalling employee data", "err", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(rawJson)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	empId := r.PathValue("id")

	id, err := strconv.ParseUint(empId, 10, 8)
	if err != nil {
		log.Println("error while parsing id from path", "err", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect id in path params."))
		return
	}
	uintId := uint8(id)

	data, err := database.GetEmployees(&uintId)
	if err != nil {
		log.Println("error while getting employee data from database", "id", uintId, "err", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if len(data) == 0 {
		log.Println("requested employee data not found", "id", uintId)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("employee not found"))
	}

	rawJson, err := json.Marshal(data[0])
	if err != nil {
		log.Println("error while marshalling employee data", "err", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(rawJson)

}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	/*
	* Generate a new unique ID
	* read data from body
	* make an entry in database
	* return 201 (status created)

	 */
	w.WriteHeader(http.StatusNotImplemented)
}

// PATCH only
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	/*
		Logic :
		* read id path parameter
		* validate it
		* read email id from body
		* validate email id against email-id regex (optional)
		* find its entry in database
		* if entry is present, then update only email
		* handle errors wherever necessary

	*/
	w.WriteHeader(http.StatusNotImplemented)

}
