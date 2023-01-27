package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home!")
}

// Todo is the shape of the each to-do
type Todo struct {
	Title   string
	Content string
}

var todos []Todo

// PageVariables are variables sent to the html template
type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

func getTodos(w http.ResponseWriter, r *http.Request) {

	pageVariables := PageVariables{
		PageTitle: "List of To-dos",
		PageTodos: todos,
	}

	//go doc template.ParseFiles --> parsing the HTML file
	t, err := template.ParseFiles("todos.html")

	//error handling
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error:", err)
	}

	//if all goes well aka no errors
	//writing what we have parsed (the HTML file (t)) back to the writer (w)
	err = t.Execute(w, pageVariables)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	//parses the form data
	err := r.ParseForm()

	//error handling
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Request Parsing Error: ", err)
	}

	//if everything goes well, creating our to-do
	todo := Todo{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	//appending our singular todo to our list of todos that is instantiated on 20
	todos = append(todos, todo)
	log.Print(todos)

	//redirecting to the /todos page
	http.Redirect(w, r, "/todos/", http.StatusSeeOther)

}

func main() {
	http.HandleFunc("/", home)

	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo/", addTodo)

	fmt.Println("Server is running on PORT :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
