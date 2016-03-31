package db

import (
	"log"

	"strconv"

	"github.com/DavidSkeppstedt/recitation/models"
)

func (this *Database) AddCourse(course *models.Course) {
	var id int
	err := this.conn.QueryRow("INSERT INTO "+
		"recitation.course(name,numtracks) VALUES ($1,$2) returning id;",
		course.Name, course.NumTracks).Scan(&id)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func (this *Database) ReadCourse(id int) (course models.Course) {
	this.conn.QueryRow("SELECT * from recitation.course "+
		"WHERE id = $1", id).Scan(&course.Id, &course.Name)
	return
}

func (this *Database) GetCourses() (courses []models.Course) {
	rows, err := this.conn.Query("SELECT * FROM recitation.course")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var tmp models.Course
		rows.Scan(&tmp.Id, &tmp.Name, &tmp.NumTracks)
		courses = append(courses, tmp)
	}
	return
}

func (this *Database) GetRecitations(courseId int) (rec []models.Recitation) {
	rows, err := this.conn.Query("SELECT name FROM "+
		"recitation.recitation where cid = $1", courseId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var tmp models.Recitation
		rows.Scan(&tmp.Name)
		rec = append(rec, tmp)
	}
	return
}

func (this *Database) AddRecitation(rec *models.RecitationSub) {
	transaction, transErr := this.conn.Begin()
	if transErr != nil {
		log.Println("Could not begin transaction whyy!!")
		panic(transErr)
	}
	//Create a recitation row
	var recitationName string
	err := transaction.QueryRow("INSERT INTO "+
		"recitation.recitation(cid,name) VALUES($1,$2) RETURNING name;",
		rec.CourseId, rec.Name).Scan(&recitationName)

	if err != nil {
		log.Println("halla", err)
		transaction.Rollback()
		panic(err)
	}

	for _, value := range rec.Problems {
		var problem string

		//Create main problem rows
		err := transaction.QueryRow(
			"INSERT INTO "+
				"recitation.problem(cid,recitation,problem,compulsory) VALUES($1,$2,$3,$4) RETURNING problem;",
			rec.CourseId, recitationName, value.Id, value.Com).Scan(&problem)
		if err != nil {
			log.Println("Hej", err)
			transaction.Rollback()
			panic(err)
		}
		//Create sub problem row
		subtask, convErr := strconv.Atoi(value.Task)
		if convErr != nil {
			log.Println("Can not convert", value.Task, "to string")
			err := transaction.Rollback()
			log.Println(err)
			panic(convErr)
			return
		}
		for i := 0; i < subtask; i++ {
			subLetter := string(i + 65)
			_, err = transaction.Exec("INSERT INTO "+
				"recitation.subproblem(cid,recitation,problem,letter) VALUES($1,$2,$3,$4);",
				rec.CourseId,
				recitationName,
				problem,
				subLetter)

			if err != nil {
				log.Println("Subproblem", err)
				transaction.Rollback()
				panic(err)
			}
		}
	}

	//and finally we are done..
	log.Println("We are done...")
	transaction.Commit()
}
