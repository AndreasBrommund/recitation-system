package web

import (
	"github.com/DavidSkeppstedt/recitation/models"
	"github.com/gorilla/context"
	"github.com/justinas/alice"
)

func initRoutes() (routes Routes) {
	middleware := alice.New(context.ClearHandler, loggingHandler, recoverHandler)

	routes.Get("api version", "/api/1",
		middleware.ThenFunc(apiVersion))

	routes.Get("List course", "/api/1/course",
		middleware.ThenFunc(apiCourseList))
	routes.Post("Add course", "/api/1/course",
		middleware.Append(BodyHandler(models.Course{})).ThenFunc(apiCourseAdd))

	routes.Post("Student/Course enrollment", "/api/1/enroll/:id",
		middleware.Append(BodyHandler(models.Enrollment{})).Append(authHandler).ThenFunc(apiEnrollStudent))

	routes.Post("Add recitation", "/api/1/recitation",
		middleware.Append(BodyHandler(models.RecitationSub{})).ThenFunc(apiRecitationAdd))

	routes.Post("Create student", "/api/1/student",
		middleware.Append(BodyHandler(models.Student{})).ThenFunc(apiCreateStudent))
	routes.Post("Student login", "/api/1/student/login",
		middleware.Append(BodyHandler(models.Student{})).ThenFunc(studentCheckPassword))

	routes.Get("Admin site", "/admin",
		middleware.ThenFunc(adminIndexHandler))
	routes.Get("Admin course site", "/admin/course",
		middleware.ThenFunc(adminCourseHandler))
	routes.Get("Admin recitation site", "/admin/course/:id",
		middleware.ThenFunc(adminRecitaionHandler))

	routes.Get("Student site", "/student",
		middleware.ThenFunc(studentIndexHandler))

	routes.Get("Student enroll site", "/enroll/:id",
		middleware.Append(authHandler).ThenFunc(enrollStudent))
	routes.Get("Student profile site", "/student/:id",
		middleware.Append(authHandler).ThenFunc(studentProfile))

	routes.Get("Student recitation site", "/student/:id/recitations/:cid",
		middleware.Append(authHandler).ThenFunc(studentRecitation))
	return
}
