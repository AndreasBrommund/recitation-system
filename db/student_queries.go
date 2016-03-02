package db

import (
	"log"

	"github.com/DavidSkeppstedt/recitation/models"
)

func (this *Database) CreateStudent(student *models.Student) (id int) {
	err := this.conn.QueryRow("INSERT INTO "+
		"recitation.student(name,password) VALUES($1,$2) returning id",
		student.Name, student.Password).Scan(&id)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return id
}

func (this *Database) CheckPassword(name, password string) (ok bool) {
	err := this.conn.QueryRow("SELECT EXISTS(SELECT * FROM recitation.student "+
		"where name=$1 AND password=$2 LIMIT 1)", name, password).Scan(&ok)
	if err != nil {
		log.Println("buuu")
		panic(err)
	}
	return
}

func (this *Database) ReadStudent(name, password string) (student models.Student) {
	err := this.conn.QueryRow("SELECT * from recitation.student "+
		"WHERE name=$1 AND password=$2", name, password).Scan(&student.Id,
		&student.Name, &student.Password)
	if err != nil {
		log.Println("Could not query a student")
		panic(err)
	}
	return
}
