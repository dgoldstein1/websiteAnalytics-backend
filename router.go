package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "io"
  "io/ioutil"
)


////////////
// ROUTER //
////////////

func NewRouter(loggerOn bool) *mux.Router {

  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    var handler http.Handler

    handler = route.HandlerFunc

    if loggerOn {handler = Logger(handler, route.Name)}

    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)
  }

  return router
}

////////////
// ROUTES //
////////////

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    "show all site visits",
    "GET",
    "/",
    getAllVisits,
  },
  Route{
    "show all site visits",
    "GET",
    "/visits",
    getAllVisits,
  },
  Route{
    "post a new visit",
    "POST",
    "/visits",
    addVisit,
  },
  Route{
    "show visits by IP address",
    "GET",
    "/visits/{ip}",
    showByIp,
  },
}

//////////////
// HANDLERS //
//////////////

/**
 * writes list of all site visits
 * TODO
 **/
func getAllVisits(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("GET /visits"))
  w.WriteHeader(http.StatusOK)
}

/**
 * Adds a visit to the postgres DB
 * TODO
 **/
func addVisit(w http.ResponseWriter, r *http.Request) {
  body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  s := string(body[:])
  w.Write([]byte("POST /visits : " + s))
  w.WriteHeader(http.StatusOK)
}

/**
 * writes list of all site visits
 * TODO
 **/
func showByIp(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  w.Write([]byte("GET /visits/{ip}" + vars["ip"]))
  w.WriteHeader(http.StatusOK)
}


