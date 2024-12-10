package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"restapi.com/database"
	"restapi.com/pkg"
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

/*
func AddEmployee(w http.ResponseWriter, r *http.Request) {

	//* Generate a new unique ID
	// Read data from body
	// Set the new ID to the employee
	// Make an entry in database
	// Return 201 (status created)
}
*/

// PATCH only
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	// Read id path parameter
	empId := r.PathValue("id")

	// Validate id
	id, err := strconv.ParseUint(empId, 10, 8)
	if err != nil {
		log.Println("error while parsing id from path", "err", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect id in path params."))
		return
	}
	uintId := uint8(id)

	// Read email id from body
	var updateData struct {
		Email string `json:"Email"`
	}
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		log.Println("error while decoding request body", "err", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request body"))
		return
	}

	// Validate email id against email-id regex (optional)
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(updateData.Email) {
		log.Println("email format is invalid", "email", updateData.Email)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("email format is invalid"))
		return
	}

	// Find its entry in database
	data, err := database.GetEmployees(&uintId)
	if err != nil {
		log.Println("error while getting employee data from database", "id", uintId, "err", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if len(data) == 0 {
		log.Println("requested employee's data not found", "id", uintId)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("employee not found"))
		return
	}

	// Update only email
	employee := *data[0]
	employee.Email = updateData.Email

	// Pass the database connection to the UpdateEmployee function
	err = database.UpdateEmployee(pkg.GetDB(), &employee)
	if err != nil {
		log.Println("error occurred while updating employee data in database", "id", uintId, "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Return 200 (status OK)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Employee details successfully updated"))
}
