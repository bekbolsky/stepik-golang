package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Group - группа студентов
type Group struct {
	// ID       int
	// Number   string
	// Year     int
	Students []struct {
		// LastName   string
		// FirstName  string
		// MiddleName string
		// Birthday   string
		// Address    string
		// Phone      string
		Rating []int
	}
}

type averRating struct {
	Average float64
}

func main() {
	// file, _ := os.Open("./text.json")
	input, _ := ioutil.ReadAll(os.Stdin)
	// fmt.Println(input)

	newGroup := Group{}
	err := json.Unmarshal(input, &newGroup)
	if err != nil {
		fmt.Println(err)
	}

	students := 0
	marks := 0
	for i := 0; i < len(newGroup.Students); i++ {
		// fmt.Println(newGroup.Students[i])
		students++
		for m := 0; m < len(newGroup.Students[i].Rating); m++ {
			marks++
			// fmt.Println(newGroup.Students[i].Rating[m])
		}
	}
	newAvg := averRating{float64(marks) / float64(students)}
	avg, _ := json.MarshalIndent(newAvg, "", "    ")
	os.Stdout.Write(avg)
	// fmt.Printf("Students: %v\n Marks: %v", students, marks)
}
