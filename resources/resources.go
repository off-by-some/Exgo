package resources

import (
  http "net/http"
  mux "github.com/gorilla/mux"
  sessionResource "Exgo/resources/session"
)

var Router = mux.NewRouter()

type HandlerFunc func(res http.ResponseWriter, req *http.Request)
var RoutesMap = map[string]HandlerFunc {
  "/session": sessionResource.HandleLogin,
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
