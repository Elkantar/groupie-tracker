package grab

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type artist struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Relation     string              `json:"relations"`
	Concerts     map[string][]string `json:"datesLocations"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Date      string   `json:"dates"`
}

type Locat struct {
	Index []Location `json:"index"`
}

type ArtistRelation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relation struct {
	Index []ArtistRelation `json:"index"`
}

type Page struct {
	Title    string
	Articles []string
}

type display struct {
	displayed bool
}

type Rows struct {
	Items []Row
	Games []Game
}

type Row struct {
	ID          string
	Imges       string
	Groupes     string
	Creations   string
	Grpmem      []string
	FirstAlbums string
	Locales     []string
}

type Game struct {
	val string
}

var (
	data  []artist
	tmpl  = template.Must(template.ParseFiles("./test.html"))
	tmpl2 = template.Must(template.ParseFiles("./site/Artists.html"))
	tmpl5 = template.Must(template.ParseFiles("./site/Locations.html"))
)

func id(names string, Artist []artist) string {
	var nam string

	date, _ := strconv.Atoi(names)
	lens := len(Artist)
	for i := 0; i < lens; i++ {
		for _, ch := range Artist[i].Members {
			if ch == names && Artist[i].ID != 13 {
				names = Artist[i].Name
			}
		}
	}
	for i := 0; i < lens; i++ {
		if Artist[i].ID == date {
			nam = Artist[i].Name
		}
	}
	return nam
}

func name(names string, Artist []artist) string {
	var nam string

	date, _ := strconv.Atoi(names)
	lens := len(Artist)
	for i := 0; i < lens; i++ {
		for _, ch := range Artist[i].Members {
			if ch == names && Artist[i].ID != 13 {
				names = Artist[i].Name
			}
		}
	}
	for i := 0; i < lens; i++ {
		if Artist[i].Name == names || Artist[i].FirstAlbum == names || Artist[i].CreationDate == date {
			nam = Artist[i].Name
		} else {
			for _, ch := range Artist[i].Members {
				if ch == names {
					nam = Artist[i].Name
				}
			}
		}
	}
	return nam
}

func image(names string, Artist []artist) string {
	var nam string

	date, _ := strconv.Atoi(names)

	lens := len(Artist)
	for i := 0; i < lens; i++ {
		for _, ch := range Artist[i].Members {
			if ch == names && Artist[i].ID != 13 {
				names = Artist[i].Name
			}
		}
	}
	for i := 0; i < lens; i++ {
		if Artist[i].Name == names || Artist[i].FirstAlbum == names || Artist[i].CreationDate == date {
			nam = Artist[i].Image
		} else {
			for _, ch := range Artist[i].Members {
				if ch == names {
					nam = Artist[i].Image
				}
			}
		}
	}
	return nam
}

func firstAlb(names string, Artist []artist) string {
	var nam string

	date, _ := strconv.Atoi(names)

	lens := len(Artist)
	for i := 0; i < lens; i++ {
		for _, ch := range Artist[i].Members {
			if ch == names && Artist[i].ID != 13 {
				names = Artist[i].Name
			}
		}
	}
	for i := 0; i < lens; i++ {
		if Artist[i].Name == names || Artist[i].FirstAlbum == names || Artist[i].CreationDate == date {
			nam = Artist[i].FirstAlbum
		} else {
			for _, ch := range Artist[i].Members {
				if ch == names {
					nam = Artist[i].FirstAlbum
				}
			}
		}
	}
	return nam
}

func Creation(names string, Artist []artist) string {
	var name int
	date, _ := strconv.Atoi(names)

	lens := len(Artist)
	for i := 0; i < lens; i++ {
		for _, ch := range Artist[i].Members {
			if ch == names && Artist[i].ID != 13 {
				names = Artist[i].Name
			}
		}
	}
	for i := 0; i < lens; i++ {
		if Artist[i].Name == names || Artist[i].FirstAlbum == names || Artist[i].CreationDate == date {
			name = Artist[i].CreationDate
		} else {
			for _, ch := range Artist[i].Members {
				if ch == names {
					name = Artist[i].CreationDate
				}
			}
		}
	}
	nam := strconv.Itoa(name)
	return nam
}

func grp(names string, Artist []artist) []string {
	var nam []string

	date, _ := strconv.Atoi(names)
	lens := len(Artist)
	for i := 0; i < lens; i++ {
		for _, ch := range Artist[i].Members {
			if ch == names && Artist[i].ID != 13 {
				names = Artist[i].Name
			}
		}
	}
	for i := 0; i < lens; i++ {
		if Artist[i].Name == names || Artist[i].FirstAlbum == names || Artist[i].CreationDate == date {
			members := Artist[i].Members
			for _, ch := range members {
				nam = append(nam, ch)
			}
		}
	}
	return nam
}

func dateplace(names string, Artist []artist) []string {
	var nam []string
	date, _ := strconv.Atoi(names)
	lens := len(Artist)
	for i := 0; i < lens; i++ {
		for _, ch := range Artist[i].Members {
			if ch == names && Artist[i].ID != 13 {
				names = Artist[i].Name
			}
		}
	}
	for i := 0; i < lens; i++ {
		if Artist[i].Name == names || Artist[i].FirstAlbum == names || Artist[i].CreationDate == date {
			key := make([]string, 0, len(Artist[i].Concerts))
			for k := range Artist[i].Concerts {
				key = append(key, k)
			}
			for _, keys := range key {
				value := Artist[i].Concerts[keys]
				keys = strings.ReplaceAll(keys, "_", " ")
				keys = keys + " :"
				nam = append(nam, keys)
				for _, strvalue := range value {
					nam = append(nam, strvalue)
				}
			}
		} else {
			for _, ch := range Artist[i].Members {
				if ch == names {
					key := make([]string, 0, len(Artist[i].Concerts))
					for k := range Artist[i].Concerts {
						key = append(key, k)
					}
					for _, keys := range key {
						value := Artist[i].Concerts[keys]
						keys = strings.ReplaceAll(keys, "_", " ")
						keys = keys + " :"
						nam = append(nam, keys)
						for _, strvalue := range value {
							nam = append(nam, strvalue)
						}
					}
				}
			}
		}
	}
	return nam
}

func gatherDataUp(link string) []artist {
	data1 := getData(link)
	Artists := []artist{}
	e := json.Unmarshal(data1, &Artists)
	if e != nil {
		log.Fatal(e)
		return nil
	}
	for i := 0; i < len(Artists); i++ {
		r := ArtistRelation{}
		json.Unmarshal(getData(Artists[i].Relation), &r)
		Artists[i].Concerts = r.DatesLocations
	}
	return Artists
}

func getData(link string) []byte {
	data1, e1 := http.Get(link)
	if e1 != nil {
		log.Fatal(e1)
		return nil
	}
	data2, e2 := ioutil.ReadAll(data1.Body)
	if e2 != nil {
		log.Fatal(e2)
		return nil
	}
	return data2
}
