package handles

import (
	"fmt"
	"net/http"
)

var CountStatus int64

func Status(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		CountStatus += 1

		msg := fmt.Sprintf("COUNT: %v\n", CountStatus)
		w.Write([]byte(msg))
	case "POST":

		w.Write([]byte("post"))
	default:
		w.Write([]byte(`"Sorry, only GET and POST methods are supported."`))
	}

}



