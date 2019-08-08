package model

import (
	_"database/sql"
	"gocrud/koneksi"
	"fmt"
)

type Employee struct {
	Id     	 int    `json:"id"`
	Name   	 string `json:"name"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
	City     string `json:"city"`

}


type Employees struct {
	Employees []Employee `json:"employee"`
}

var (
	emps = map[int]*Employee{}
	seq   = 1
)
//var con *sql.DB
var con = koneksi.DbConn()

func GetEmployee() Employees {
	//db.CreateCon()
	sqlStatement := "SELECT * FROM employee order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}

		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Password, &employee.City)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Employees = append(result.Employees, employee)
	}
	return result

}

func SaveEmployee(name,email,password,city string){
	_, err := con.Exec("INSERT INTO employee(name, email, password, city) VALUE (?,?,?,?)",name,email,password,city)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func UpdateEmployee(id int, name,email,password,city string ){
	_, err := con.Exec("UPDATE employee SET name = ?, email = ?, password = ?, city = ? WHERE id= ?", name, email, password, city, id)
	if err != nil {
		fmt.Println(err)
	}

	defer koneksi.DbConn().Close()
}

func DeleteEmployee(id int) {
	_, err:= koneksi.DbConn().Exec("DELETE FROM employee WHERE id =?",id)
	if err != nil {
		fmt.Println(err)
	}
}

func Login(email, password string)   {

	_, err := con.Exec("SELECT * FROM employee WHERE email=? AND password=?", email, password)
	if err !=nil{
		fmt.Println(err)
	}

	defer koneksi.DbConn().Close()
	}