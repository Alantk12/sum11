package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type Prizes struct {
	Prizes []years `json:"prizes"`
}

type years struct {
	Year      string `json:"year"`
	Category  string `json:"category"`
	Laureates string `json:"laureates"`
}
type ids struct {
	ID         string `json:"id"`
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Motivation string `json:"motivation"`
	Share      string `json:"share"`
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")

	jsonFile, err := os.Open("prize.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// //fmt.Println("Successfully Opened users.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	file, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	// //fmt.Println("Successfully Opened users.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	file, _ := ioutil.ReadAll(jsonFile)

	var prizess Prizes

	json.Unmarshal(file, &prizess)

	template, _ := template.ParseFiles("hello.html")
	template.Execute(w, file)

	// Open our jsonFile
	// jsonFile, err := os.Open("prize.json")
	// if we os.Open returns an error then handle it
	// if err != nil {
	//     fmt.Println(err)
	// }

	// fmt.Println("Successfully Opened prize.json")
	// defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	// byteValue, _ := ioutil.ReadAll(jsonFile)

	//  we initialize our Users array
	// var prize prize
	// data :=prize{}

	//  we unmarshal our byteArray which contains our
	//  jsonFile's content into 'users' which we defined above
	// json.Unmarshal(byteValue, &prize)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	//// for i := 0; i < len(prizess.Prizes); i++ {
	// //	fmt.Println("Year: ", prizess.Prizes[i].Year)
	// //	fmt.Println("Category: ", prizess.Prizes[i].Category)
	// fmt.Println("Lauretes:",prizess.Prizes)
}

//}
