package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"restapi.com/controllers"
	"restapi.com/models"
	"restapi.com/pkg"
)

const (
	PORT         int    = 3000
	DATABASE_URL string = "root@tcp(127.0.0.1:3306)/abc"
)

func main() {

	dbConn, err := pkg.ConnectToDatabase(DATABASE_URL)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer dbConn.Close()
	// set to global variable so that it can be used elsewhere

	pkg.DbConn = dbConn

	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			// do something
			emps, err := controllers.GetEmployees()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			rawJson, err := json.Marshal(emps)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Error while unmarshalling response, %s", err.Error())))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(rawJson)
			// end of GET
		case http.MethodPost:
			// do something
			var emp models.Employee
			err := json.NewDecoder(r.Body).Decode(&emp)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("Error while decoding JSON, %s", err.Error())))
				return
			}
			id, err := controllers.CreateEmployee(dbConn, emp)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("There is an error inserting into database, %s", err.Error())))
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf("Employee created with ID %d", id)))

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}

	})

	http.HandleFunc("/employee/{Id}", func(w http.ResponseWriter, r *http.Request) {
		// do function for delete

		err := controllers.DeleteEmployee(dbConn, Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error while deleting employee data, %s", err.Error())))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Employee with id %d is deleted", Id)))

	})
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
