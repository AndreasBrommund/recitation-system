package web

import (
	"encoding/json"
	"net/http"

	"strconv"

	"log"

	"github.com/DavidSkeppstedt/recitation/models"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func apiVersion(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		}{"Recitation Service", "0.1"})
}

func apiRecitationAdd(w http.ResponseWriter, r *http.Request) {
	data := Body(r).(*models.RecitationSub)
	database.AddRecitation(data)
}

func apiCourseAdd(w http.ResponseWriter, r *http.Request) {
	course := Body(r).(*models.Course)
	log.Println(course)
	if !course.Validate() {
		panic("Nope")
	}
	database.AddCourse(course)
}

func apiCourseList(w http.ResponseWriter, r *http.Request) {
	data := database.GetCourses()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func apiEnrollStudent(w http.ResponseWriter, r *http.Request) {
	data := Body(r).(*models.Enrollment)
	database.EnrollStudent(data)
}

func apiCreateStudent(w http.ResponseWriter, r *http.Request) {
	student := Body(r).(*models.Student)
	database.CreateStudent(student)
	log.Println(student)
}

func apiCreateSolutions(w http.ResponseWriter, r *http.Request) {
	data := Body(r).(*models.Solved)
	log.Println("I am alive!")
	log.Println(data)
}

func studentCheckPassword(w http.ResponseWriter, r *http.Request) {
	student := Body(r).(*models.Student)
	login := database.CheckPassword(student.Name, student.Password)
	if login {

		tmp := database.ReadStudent(student.Name, student.Password)
		session, err := store.Get(r, strconv.Itoa(tmp.Id))
		if err != nil {
			log.Println("Session broken")
			panic(err)
		}
		session.Values["Name"] = tmp.Name
		session.Values["Id"] = tmp.Id
		session.Values["Password"] = tmp.Password
		session.Save(r, w)
		json.NewEncoder(w).Encode(
			struct {
				Id int `json:"id"`
			}{tmp.Id})

	} else {
		w.WriteHeader(401)
		w.Write([]byte("NOOB! NO ACCESS"))
	}
}

func studentIndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "student")
}

func studentProfile(w http.ResponseWriter, r *http.Request) {
	ps := context.Get(r, "params").(httprouter.Params)
	id := ps.ByName("id")
	session, err := store.Get(r, id)
	if err != nil {
		panic(err)
	}

	name := session.Values["Name"].(string)
	password := session.Values["Password"].(string)
	student := database.ReadStudent(name, password)
	courses := database.ReadCourseStudent(student.Id)
	renderTemplate(w, "profile", struct {
		Name      string
		StudentId int
		Data      []models.Course
	}{student.Name, student.Id, courses})
}

func studentRecitation(w http.ResponseWriter, r *http.Request) {
	ps := context.Get(r, "params").(httprouter.Params)
	studentId := ps.ByName("id")
	courseId := ps.ByName("cid")
	var recitaitons []models.Recitation
	var course models.Course
	if id, err := strconv.Atoi(courseId); err == nil {
		recitaitons = database.GetRecitations(id)
		course = database.ReadCourse(id)
	} else {
		panic(err)
	}

	renderTemplate(w, "recitations_list", struct {
		Data       []models.Recitation
		CourseName string
		CourseId   int
		StudentId  string
	}{recitaitons, course.Name, course.Id, studentId})
}

func studenSolutions(w http.ResponseWriter, r *http.Request) {
	ps := context.Get(r, "params").(httprouter.Params)
	recitationId := ps.ByName("rid")
	studentId := ps.ByName("id")
	rid, err := strconv.Atoi(recitationId)
	if err != nil {
		panic(err)
	}
	data := database.ReadProblems(rid)

	renderTemplate(w, "solutions", struct {
		Data []models.DisplayProblem
		Rid  string
		Sid  string
	}{data, recitationId, studentId})

}
func enrollStudent(w http.ResponseWriter, r *http.Request) {
	ps := context.Get(r, "params").(httprouter.Params)
	id := ps.ByName("id")
	courses := database.GetCourses()
	renderTemplate(w, "enroll", struct {
		Data []models.Course
		Id   string
	}{courses, id})
}
func adminIndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "admin")
}
func adminCourseHandler(w http.ResponseWriter, r *http.Request) {
	courses := database.GetCourses()
	renderTemplate(w, "course",
		struct {
			Test string
			Data []models.Course
		}{"A course", courses})
}
func adminRecitaionHandler(w http.ResponseWriter, r *http.Request) {
	ps := context.Get(r, "params").(httprouter.Params)
	url := r.URL.Query()
	name := url.Get("title")
	id, _ := strconv.Atoi(ps.ByName("id"))
	data := database.GetRecitations(id)
	renderTemplate(w, "recitation",
		struct {
			Title       string
			Id          int
			Recitations []models.Recitation
		}{name, id, data})
}

func renderTemplate(w http.ResponseWriter, tmpl string, data ...interface{}) {
	var err error
	if len(data) > 0 {
		err = templates.ExecuteTemplate(w, tmpl+".html", data[0])
	} else {
		err = templates.ExecuteTemplate(w, tmpl+".html", nil)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}
