package db

import (
	"log"

	"github.com/DavidSkeppstedt/recitation/models"
)

func (this *Database) ReadProblems(rid int) (problems []models.DisplayProblem) {
	//SELECT * from recitation.problem join recitation.have on pid = id where rid = 34

	rows, err := this.conn.Query("SELECT id,problemid,compulsory from "+
		"recitation.problem JOIN recitation.have on pid = id where rid = $1", rid)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		//do something for every problem...
		var tmp models.DisplayProblem
		rows.Scan(&tmp.Id, &tmp.ProblemNr, &tmp.Compulsory)
		//SELECT letter from recitation.subproblem join recitation.belongs on id=spid where pid=37

		//inner query
		r, err := this.conn.Query("SELECT letter from "+
			"recitation.subproblem join recitation.belongs on id=spid where pid=$1", tmp.Id)

		if err != nil {
			log.Println("Subproblems fail")
			panic(err)
		}

		for r.Next() {
			var t models.Subproblem
			r.Scan(&t.Letter)
			tmp.Subproblems = append(tmp.Subproblems, t)
		}
		problems = append(problems, tmp)
	}
	return
}
