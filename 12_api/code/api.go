package main

import (
	//reminder io/ioutil parses the body, while encoding/json parses the json
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

// structs for our types
// the backticks are to map that part of the response object to each part of our struct
type Planet struct {
	Name       string `json:"name"`
	Terrain    string `json:"terrain"`
	Population string `json:"population"`
}
type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}
type AllPeople struct {
	People []Person `json:"results"`
}

// adding a method to the Person Struct called getHomeworld() to run another API request
func (p *Person) getHomeworld() {
	//http.Get runs a GET request to a specific URL
	//in this example, we are getting it from our person struct since that's what we get from the initial API call
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Print("Error fetching homeworld", err)
	}

	//since we are using the integrated if block syntax, we need to declare bytes outside of it and not include the : in the assignment
	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Error reading response body", err)
	}
	//json.Unmarshal parses the bytes we get from the res.Body and maps it to the pointer in which we want to add it to
	json.Unmarshal(bytes, &p.Homeworld)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Getting people") //shows up on the browser

	//the inital GET request to the URL
	res, err := http.Get(BaseURL + "people/")

	//non-integrated error handling for the GET request
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	}

	// fmt.Println(res) //<-- shows up as a request

	//ioutil.ReadAll() parses the response body into bytes
	bytes, err := ioutil.ReadAll(res.Body)

	//error handling
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse request body")
	}

	//creating our variable in which to map the json parsing to
	var people AllPeople
	fmt.Println(string(bytes)) //transforming the bytes into a readable string
	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("error parsing json", err)
	}
	fmt.Println(people)

	//looping through our array of people to map the homeworld from the homeworldURL using the method we created
	for _, pers := range people.People {
		pers.getHomeworld()
		fmt.Println(pers)
	}
}

func main() {
	fmt.Println(BaseURL)

	http.HandleFunc("/people/", getPeople)
	fmt.Println("SERVING ON PORT :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
