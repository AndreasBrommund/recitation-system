package db

import (
	"log"

	"github.com/DavidSkeppstedt/recitation/models"
)

func (this *Database) ReadProblems(rid string,cid int) (problems []models.DisplayProblem) {

	rows, err := this.conn.Query("SELECT problem,compulsory from "+
		"recitation.problem where recitation = $1 AND cid = $2", rid,cid)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		//do something for every problem...
		var tmp models.DisplayProblem
		rows.Scan(&tmp.Problem, &tmp.Compulsory)

		//inner query
		r, err := this.conn.Query("SELECT letter from "+
			"recitation.subproblem where cid = $1 and recitation = $2 and problem = $3", cid,rid,tmp.Problem)

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
