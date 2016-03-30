package web

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/DavidSkeppstedt/suparAppen/fel"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

//loggingHandler is a middleware that logs the time it takes to
//serve a specific endpoint handler.
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("%s: [%s] %q %v\n",
			context.Get(r, "name"),
			r.Method,
			r.URL.String(),
			t2.Sub(t1))
	})
}

//Panic recovery handler, will log the error that occurred but save
//the app from crashing
func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				fel.WriteError(w, fel.ErrInternalServer)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

//Auth handler, looking at the session data
//Much secure wow!

func authHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ps := context.Get(r, "params").(httprouter.Params)
		id := ps.ByName("id")
		session, err := store.Get(r, id)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
			return

		}
		//here we check if the cookie exists at all..
		if !session.IsNew {
			name := session.Values["Name"].(string)
			password := session.Values["Password"].(string)
			myid := session.Values["Id"].(int)

			ok := database.CheckPassword(name, password)
			if ok {
				tmp := database.ReadStudent(name, password)
				if tmp.Id != myid {
					//not auth
				} else {
					next.ServeHTTP(w, r)
					return
				}
			}
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized),
			http.StatusUnauthorized)
	}
	return http.HandlerFunc(fn)
}

// Decodes a request body into the struct passed to the middleware.
// If the request body is not JSON, it will return a 400 Bad Request error.
// Stores the decoded body into a context object.
func BodyHandler(v interface{}) func(http.Handler) http.Handler {
	t := reflect.TypeOf(v)

	m := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			val := reflect.New(t).Interface()
			err := json.NewDecoder(r.Body).Decode(val)

			if err != nil {
				log.Println(err)
				fel.WriteError(w, fel.ErrBadRequest)
				return
			}

			if next != nil {
				context.Set(r, "body", val)
				next.ServeHTTP(w, r)
			}
		}

		return http.HandlerFunc(fn)
	}

	return m
}

// Body(r *http.Request) is a function to get the decoded body from the request context
func Body(r *http.Request) interface{} {
	return context.Get(r, "body")
}
