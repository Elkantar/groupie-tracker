package grab

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"syscall"
	"text/template"
)

var ID string

func Erreur404(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./site/Erreur.html")
}

func Serv(w http.ResponseWriter, r *http.Request) {
	// var all string
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	//--------------Traitement erreur 404 et 500--------------//
	if r.URL.Path != "/" && r.URL.Path != "/Artists" && r.URL.Path != "/Contact" && r.URL.Path != "/Date" && r.URL.Path != "/Locate" && r.URL.Path != "/Error404" {
		http.Redirect(w, r, "/Error404", http.StatusSeeOther)
	}
	_, err := template.ParseFiles("test.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500: Internal Server Error"))
		log.Println(http.StatusInternalServerError)
		return
	}
	data = gatherDataUp("https://groupietrackers.herokuapp.com/api/artists")
	Artist := data
	var crea string
	var Groupe string
	var grpmembers []string
	var locale []string
	datas := Rows{}
	switch r.Method {

	case "GET":
		for i := 1; i < 53; i++ {

			s := strconv.Itoa(i)
			ID = id(s, Artist)
			crea = Creation(ID, Artist)
			Groupe = name(ID, Artist)
			grpmembers = grp(ID, Artist)
			firstalbum := firstAlb(ID, Artist)
			imag := image(ID, Artist)
			locale = dateplace(ID, Artist)
			view := Row{
				ID:          ID,
				Imges:       imag,
				Groupes:     Groupe,
				Creations:   crea,
				Grpmem:      grpmembers,
				FirstAlbums: firstalbum,
				Locales:     locale,
			}
			datas.Items = append(datas.Items, view)
		}

		if err := tmpl.Execute(w, &datas); err != nil {
			if errors.Is(err, syscall.EPIPE) {
				// log.Fatal(err)
				return
			} else {
				// tmpl.Execute(w, &datas)
			}
		}

	case "POST":

		names := r.FormValue("TXToptions")
		if names != "" {
			ID = names
		} else {
			ID = r.FormValue("artist_page")
		}
		crea = Creation(names, Artist)
		Groupe = name(names, Artist)
		grpmembers = grp(Groupe, Artist)
		firstalbum := firstAlb(names, Artist)
		imag := image(names, Artist)
		locale = dateplace(names, Artist)

		view := Row{
			ID:          ID,
			Imges:       imag,
			Groupes:     Groupe,
			Creations:   crea,
			Grpmem:      grpmembers,
			FirstAlbums: firstalbum,
			Locales:     locale,
		}

		datas.Items = append(datas.Items, view)
		if err := tmpl.Execute(w, &datas); err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func Artistepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/Artists" && r.URL.Path != "/Contact" && r.URL.Path != "/Date" && r.URL.Path != "/Locate" {
		http.Redirect(w, r, "/Error404", http.StatusSeeOther)
	}
	_, err := template.ParseFiles("test.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500: Internal Server Error"))
		log.Println(http.StatusInternalServerError)
		return
	}
	Artist := data
	var crea string
	var Groupe string
	var grpmembers []string
	var locale []string
	datas := Rows{}
	ID = r.FormValue("artist_page")

	crea = Creation(ID, Artist)
	Groupe = name(ID, Artist)
	grpmembers = grp(ID, Artist)
	firstalbum := firstAlb(ID, Artist)
	imag := image(ID, Artist)
	locale = dateplace(ID, Artist)

	view := Row{
		ID:          ID,
		Imges:       imag,
		Groupes:     Groupe,
		Creations:   crea,
		Grpmem:      grpmembers,
		FirstAlbums: firstalbum,
		Locales:     locale,
	}

	datas.Items = append(datas.Items, view)

	if err := tmpl2.Execute(w, datas); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func Date(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/Artists" && r.URL.Path != "/Contact" && r.URL.Path != "/Date" && r.URL.Path != "/Locate" {
		http.Redirect(w, r, "/Error404", http.StatusSeeOther)
	}
	_, err := template.ParseFiles("test.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500: Internal Server Error"))
		log.Println(http.StatusInternalServerError)
		return
	}
	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "./site/Dates.html")
	case "POST":

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/Artists" && r.URL.Path != "/Contact" && r.URL.Path != "/Date" && r.URL.Path != "/Locate" {
		http.Redirect(w, r, "/Error404", http.StatusSeeOther)
	}

	_, err := template.ParseFiles("test.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500: Internal Server Error"))
		log.Println(http.StatusInternalServerError)
		return
	}
	switch r.Method {

	case "GET":
		http.ServeFile(w, r, "./site/Contact.html")
	case "POST":

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func Locations(w http.ResponseWriter, r *http.Request) {
	var UserChoice string
	var Choice string
	Artist := data
	var crea string
	var Groupe string
	var grpmembers []string
	var locale []string
	view2 := Row{}
	var result string

	if r.URL.Path != "/" && r.URL.Path != "/Artists" && r.URL.Path != "/Contact" && r.URL.Path != "/Date" && r.URL.Path != "/Locate" {
		http.Redirect(w, r, "/Error404", http.StatusSeeOther)
	}

	_, err := template.ParseFiles("site/Locations.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500: Internal Server Error"))
		log.Println(http.StatusInternalServerError)
		return
	}
	switch r.Method {

	case "GET":

		datas := Rows{}

		max := 53
		min := 0
		random := rand.Intn(max-min) + min
		s := strconv.Itoa(random)
		ID = id(s, Artist)
		crea = Creation(ID, Artist)
		Groupe = name(ID, Artist)
		grpmembers = grp(ID, Artist)
		firstalbum := firstAlb(ID, Artist)
		imag := image(ID, Artist)
		locale = dateplace(ID, Artist)

		if Choice == "Yes" {
			view2 = Row{
				ID:          ID,
				Imges:       imag,
				Groupes:     Groupe,
				Creations:   crea,
				Grpmem:      grpmembers,
				FirstAlbums: firstalbum,
				Locales:     locale,
			}
		}

		if Groupe == UserChoice && Choice == "Yes" {
			result = "You Win"
		} else {
			result = "you lose"
		}
		view := Game{
			val: result,
		}

		datas.Games = append(datas.Games, view)
		datas.Items = append(datas.Items, view2)

		if err := tmpl5.Execute(w, datas); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	case "POST":
		UserChoice = r.FormValue("Jeu")
		Choice = r.FormValue("JeuCase")
		datas := Rows{}

		max := 53
		min := 0
		random := rand.Intn(max-min) + min
		s := strconv.Itoa(random)
		ID = id(s, Artist)
		crea = Creation(ID, Artist)
		Groupe = name(ID, Artist)
		grpmembers = grp(ID, Artist)
		firstalbum := firstAlb(ID, Artist)
		imag := image(ID, Artist)
		locale = dateplace(ID, Artist)

		if Choice == "Yes" {
			view2 = Row{
				ID:          ID,
				Imges:       imag,
				Groupes:     Groupe,
				Creations:   crea,
				Grpmem:      grpmembers,
				FirstAlbums: firstalbum,
				Locales:     locale,
			}
		}
		fmt.Print("rand :", random, "\n")
		fmt.Printf("choice : %s\n", UserChoice)
		fmt.Printf("Result : %s\n", Groupe)
		if Groupe == UserChoice && Choice == "Yes" {
			result = "You Win"
		} else {
			result = "You lose"
		}
		view := Game{
			val: result,
		}

		datas.Games = append(datas.Games, view)
		datas.Items = append(datas.Items, view2)

		if err := tmpl5.Execute(w, datas); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
