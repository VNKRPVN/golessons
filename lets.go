package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("chela zovut: %s. Emy %d i y nego stolko deneg: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_pAge(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
	// bob.setNewName("Lexa")
	// fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_pAge.html")
	tmpl.Execute(w, bob)
}

func next_pAge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "one more pAge")
}

func handleRequest() {
	http.HandleFunc("/", home_pAge)
	http.HandleFunc("/nextPAge/", next_pAge)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
