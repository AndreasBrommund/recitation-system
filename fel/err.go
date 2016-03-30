package fel

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Errs struct {
	Errors []Err `json:"errors"`
}

type Err struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (this Err) Error() string {
	return this.Id + ", " + strconv.Itoa(this.Status) + ", " + this.Detail
}

func WriteError(w http.ResponseWriter, err error) {
	e := err.(Err)
	w.WriteHeader(e.Status)
	json.NewEncoder(w).Encode(Errs{[]Err{e}})
}
