package resources

import (
    "net/http"
    "github.com/gorilla/mux"
    sessionResource "Exgo/resources/session"
    L "Exgo/logger"
    socketResource "Exgo/resources/socket"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {

      // Set up logging for each request
      handler := L.Logger(route.HandlerFunc, route.Name)

      router.
          Methods(route.Method).
          Path(route.Pattern).
          Name(route.Name).
          Handler(handler)
    }

    return router
}

var routes = Routes{
    Route{
        "CreateUser",
        "POST",
        "/user",
        sessionResource.Create,
    },
    Route{
      "OpenSocket",
      "GET",
      "/sock",
      socketResource.Connect,
    },
}


// TODO: Continue this later
// Preform some light reflection to grab the underlying type
// func getMemberNames(recordType interface{}) []string {
//   ffType := reflect.TypeOf(recordType)
//   memberCount := ffType.NumField()
//   var fnames = make([]string, memberCount)
//
//   for i := 0; i < memberCount; i++ {
//     fnames[i] = ffType.Field(i).Name
//   }
//
//   return fnames
// }
//
// func serializeRow(row *sq.Row, schema interface{}) {
//   keys := getMemberNames(schema)
// }
