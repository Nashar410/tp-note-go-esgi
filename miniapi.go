package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", GetHour)
	http.HandleFunc("/add", PostEntry)
	http.HandleFunc("/entries", GetEntries)
	http.ListenAndServe(":9000", nil)

}

func GetHour(w http.ResponseWriter, req *http.Request) {
	currentHour:= time.Now()
	fmt.Fprintf(w, "%dh%02d", currentHour.Hour(),currentHour.Minute())
}

func PostEntry(w http.ResponseWriter, req *http.Request) {
	if _, err := os.Stat("save.txt"); os.IsNotExist(err) {
		empty := []byte("")
		errCreation := os.WriteFile("save.txt", empty, 0644)
		fmt.Println(errCreation)
	}

	if err := req.ParseForm(); err != nil {
		fmt.Println("Something went bad")
		fmt.Fprintln(w, "Something went bad")
		return
	}
	authorEntry:= AuthorEntry{
		req.FormValue("author"),
		req.FormValue("entry")}

	f, err := os.OpenFile("save.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err, "test")
	}
	defer f.Close()
	if _, err := f.WriteString(authorEntry.entry + "\n"); err != nil {
		fmt.Println(err, "test2")
	}
	fmt.Fprintf(w, "%s: %s", authorEntry.Author, authorEntry.entry)

}

func GetEntries(w http.ResponseWriter, req *http.Request) {

	// test si save.txt existe
	if _, err := os.Stat("save.txt"); os.IsNotExist(err) {
		fmt.Fprintf(w, "Pas d'entry en mémoire")
	}
	data, _ :=  os.ReadFile("save.txt")
	fmt.Fprintf(w, "%s", string(data))
}

type AuthorEntry struct {
	Author string
	entry string
}