package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func main() {
	s := Student{"Tom", 59}
	s.Score = 60
	students := []Student{s}

	students[0].Score = 70

	fmt.Printf("s: %+v\n", s)                     // s: {Name:Tom Score:60}
	fmt.Printf("students[0]: %+v\n", students[0]) //

	students2 := []Student{{"Tom", 58}, {"Lucy", 59}}
	for _, s2 := range students2 {
		s2.Score = 60
	}

	fmt.Println("students:", students2) // students: [{Tom 58} {Lucy 59}]

	for i, _ := range students2 {
		students2[i].Score = 1
	}

	fmt.Println("students:", students2) // students: [{Tom 58} {Lucy 59}]

	for i := range students2 {
		fmt.Println("students:", i) // students: [{Tom 58} {Lucy 59}]

	}

}
