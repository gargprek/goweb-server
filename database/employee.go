package database

import (
	"database/sql"
	"fmt"

	"restapi.com/models"
	"restapi.com/pkg"
)

func GetEmployees(id *uint8) ([]*models.Employee, error) {
	db := pkg.DbConn

	queryStr := "SELECT * FROM employeedata"

	if id != nil {
		queryStr = fmt.Sprintf("SELECT * FROM employeedata WHERE id = %d", *id)
	}
	rows, err := db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := make([]*models.Employee, 0)
	for rows.Next() {
		emp := models.Employee{}
		err := rows.Scan(&emp.Id, &emp.FirstName, &emp.LastName, &emp.Email, &emp.HiringDate)
		if err != nil {
			return nil, err
		}
		employees = append(employees, &emp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func CreateEmployee(db *sql.DB, emp models.Employee) (int64, error) {
	query := "INSERT INTO employeedata (FirstName, LastName, Email, HiringDate) VALUES (?,?,?,?)"
	result, err := db.Exec(query, emp.FirstName, emp.LastName, emp.Email, emp.HiringDate)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DeleteEmployee(db *sql.DB, id int64) error {
	query := "DELETE FROM employeedata WHERE ID = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no data found with ID %d", id)
	}

	return nil
}
