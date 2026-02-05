package mensa_menu_wuerzburg_api

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	m "github.com/isathecatgirl/mensa-menu-wuerzburg"
)

type MenuData struct {
	Timestamp int64  `json:"timestamp"`
	Menu      m.Menu `json:"menu"`
}

type ResponseData struct {
	JOSEF_SCHNEIDER_STRASSE MenuData
	ROENTGENRING            MenuData
	STUDENTENHAUS           MenuData
	HUBLAND_NORD            MenuData
	HUBLAND_SUED            MenuData
}

var Mensa = ResponseData{
	JOSEF_SCHNEIDER_STRASSE: MenuData{},
	ROENTGENRING:            MenuData{},
	STUDENTENHAUS:           MenuData{},
	HUBLAND_NORD:            MenuData{},
	HUBLAND_SUED:            MenuData{},
}

var paths = []string{"/josef_schneider_strasse", "/roentgenring", "/studentenhaus", "/hubland_nord", "/hubland_sued"}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	if !slices.Contains(paths, path) {
		http.NotFound(w, r)
		return
	}

	time := time.Now().UnixMilli()

	mensa := &Mensa.HUBLAND_NORD
	mensaString := ""

	switch r.URL.Path[1:] {
	case "josef_schneider_strasse":
		mensa = &Mensa.JOSEF_SCHNEIDER_STRASSE
		mensaString = m.Mensa.JOSEF_SCHNEIDER_STRASSE

	case "roentgenring":
		mensa = &Mensa.ROENTGENRING
		mensaString = m.Mensa.ROENTGENRING

	case "studentenhaus":
		mensa = &Mensa.STUDENTENHAUS
		mensaString = m.Mensa.STUDENTENHAUS

	case "hubland_nord":
		mensa = &Mensa.HUBLAND_NORD
		mensaString = m.Mensa.HUBLAND_NORD

	case "hubland_sued":
		mensa = &Mensa.HUBLAND_SUED
		mensaString = m.Mensa.HUBLAND_SUED
	}

	if mensa.Timestamp+5*60*1000 < time {
		mensa.Timestamp = time
		mensa.Menu = m.GetMenu(mensaString)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mensa)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
