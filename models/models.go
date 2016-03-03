package models

type Course struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	CourseId string `json:"code"`
}

type Recitation struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Track string `json:"track"`
}

type RecitationSub struct {
	RecitationId int       `json:"course_id"`
	Name         string    `json:"name"`
	Track        string    `json:"track"`
	NrProblems   string    `json:"nr_problems"`
	Problems     []Problem `json:"problems"`
}
type Problem struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
	Com  string `json:"com"`
}

type Student struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Enrollment struct {
	Student int   `json:"student"`
	Courses []int `json:"courses"`
}

type DisplayProblem struct {
	Id          int
	ProblemNr   int
	Compulsory  int
	Subproblems []Subproblem
}

type Subproblem struct {
	Letter string
}