package resources

import (
    "net/http"
    "github.com/gorilla/mux"
    sessionResource "Exgo/resources/session"
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
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
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
