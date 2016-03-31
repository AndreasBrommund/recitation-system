package models

type Course struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	NumTracks int    `json:"tracks"`
}

func (this *Course) Validate() bool {
	return this.Name != "" && this.NumTracks > 0
}

type Recitation struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Track string `json:"track"`
}

type RecitationSub struct {
	CourseId   int       `json:"course_id"`
	Name       string    `json:"name"`
	NrProblems string    `json:"nr_problems"`
	Problems   []Problem `json:"problems"`
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
	Problem     string
	Compulsory  int
	Subproblems []Subproblem
}

type Subproblem struct {
	Letter string
}

type Solved struct {
	Problems       map[string][]string `json:"problems"`
	RecitationName string              `json:"recitation_name"`
	Course         int                 `json:"course_id"`
}
