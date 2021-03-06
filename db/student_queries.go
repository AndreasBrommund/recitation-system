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

func (this *Database) ReadCourseStudent(id int) (courses []models.Course) {
	rows, err := this.conn.Query("SELECT name,id "+
		"from recitation.course JOIN recitation.takes on "+
		"id = cid WHERE sid =$1;", id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var tmp models.Course
		rows.Scan(&tmp.Name, &tmp.Id)
		courses = append(courses, tmp)
	}
	return
}

func (this *Database) GetCoursesNotEnrolled(id int) (courses []models.Course) {
	rows, err := this.conn.Query("(select id,name from recitation.course) "+
		"EXCEPT (SELECT id,name FROM recitation.course "+
		"JOIN recitation.takes on id = cid where sid = $1)", id)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var tmp models.Course

		rows.Scan(&tmp.Id, &tmp.Name)
		courses = append(courses, tmp)
	}
	return
}

func (this *Database) EnrollStudent(enrollment *models.Enrollment) {
	for _, value := range enrollment.Courses {
		_, err := this.conn.Exec("INSERT INTO "+
			"recitation.takes(cid,sid) VALUES($1,$2);", value, enrollment.Student)
		if err != nil {
			log.Println("something wrong enrolling student")
			panic(err)
		}
	}
}

func (this *Database) RegisterSolved(sid int, solved *models.Solved) {

	for k, v := range solved.Problems {

		for _, letter := range v {
			_, err := this.conn.Exec("INSERT INTO "+
				"recitation.solved(sid,cid,recitation,problem,letter) VALUES($1,$2,$3,$4,$5);",
				sid,
				solved.Course,
				solved.RecitationName,
				k,
				letter)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (this *Database) RegisterTrack(sid int, solved *models.Solved) {
	_, err := this.conn.Exec("INSERT INTO "+
		"recitation.track(cid,name,sid,track) VALUES($1,$2,$3,$4);",
		solved.Course, solved.RecitationName, sid, solved.Track)
	if err != nil {
		panic(err)
	}
}

func (this *Database) PointsForRecitation(sid int, cid int, recitaitons []models.Recitation) {

	for i, v := range recitaitons {
		points := 0
		this.conn.QueryRow("SELECT points FROM recitation.points WHERE cid=$1 AND recitation =$2 and sid = $3",
			cid, v.Name, sid).Scan(&points)
		recitaitons[i].Points = points
	}
}
