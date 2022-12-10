package httpHandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mactsouk/restdb"
)

// GetAllHandler is for getting all data from the user database
func GetAllHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("GetAllHandler Serving:", r.URL.Path, "from", r.Host)
	d, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(d) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println("No input!")
		return
	}

	var user = restdb.User{}
	err = json.Unmarshal(d, &user)
	fmt.Println(&user)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if !restdb.IsUserAdmin(user) {
		log.Println("User", user, "is not an admin!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = SliceToJSON(restdb.ListAllUsers(), rw)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}
