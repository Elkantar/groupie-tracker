package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	"grab/grab"
)

func server() {
	listener, err := net.Listen("tcp", ":3010")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("server", err)
		os.Exit(1)
	}
	data := make([]byte, 1)
	if _, err := conn.Read(data); err != nil {
		log.Fatal("server", err)
	}

	conn.Close()
}

func client() {
	conn, err := net.Dial("tcp", "localhost:3010")
	if err != nil {
		log.Fatal("client", err)
	}

	// write to make the connection closed on the server side
	if _, err := conn.Write([]byte("a")); err != nil {
		log.Printf("client: %v", err)
	}

	time.Sleep(1 * time.Second)

	// write to generate an RST packet
	if _, err := conn.Write([]byte("b")); err != nil {
		log.Printf("client: %v", err)
	}

	time.Sleep(1 * time.Second)

	// write to generate the broken pipe error
	if _, err := conn.Write([]byte("c")); err != nil {
		log.Printf("client: %v", err)
		if errors.Is(err, syscall.EPIPE) {
			log.Print("This is broken pipe error")
		}
	}
}

func main() {
	/*go server()

	time.Sleep(3 * time.Second) // wait for server to run

	client()
	*/
	fmt.Println("Server launch at : http://localhost:3010/ ")
	http.HandleFunc("/", grab.Serv)
	http.HandleFunc("/Artists", grab.Artistepage)
	http.HandleFunc("/Date", grab.Date)
	http.HandleFunc("/Contact", grab.Contact)
	http.HandleFunc("/Locate", grab.Locations)
	http.HandleFunc("/Error404", grab.Erreur404)
	fs := http.FileServer(http.Dir("assetCss"))
	http.Handle("/assetCss/", http.StripPrefix("/assetCss/", fs))
	img := http.FileServer(http.Dir("img"))
	http.Handle("/img/", http.StripPrefix("/img/", img))
	err := http.ListenAndServe(":3010", nil)
	if err != nil && errors.Is(err, syscall.EPIPE) {
		log.Fatal(err)
	}
	http.ListenAndServe(":3010", nil)
}
