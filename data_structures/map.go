package main
import "fmt"

func main () {
	/*
		* MAP
		- map data structures stores lements in key/valuepairs.
		- keys are unique identifiers associated with each value in a map
	*/

	subjectMarks := map[string]float32{"Go": 90, "JAva": 78, "Python": 89}
	fmt.Println(subjectMarks)

	// # ACEESS VALUES OF A MAP
	fmt.Println(subjectMarks["Go"])

	// # Change Vlue on a Map
	subjectMarks["Java"] = 90
	fmt.Println(subjectMarks)

	// # ADD ELEMENT TO A GO MAP
	subjectMarks["Javascript"] = 70
	fmt.Println(subjectMarks)

	// # DELETE ELEMENT OF a MAP
	// - TO DELETE we use delete() function
	delete(subjectMarks, "Go")
	fmt.Println(subjectMarks)

	//  # Loop through a map
	for marks, laguage := range subjectMarks {
		fmt.Printf("Laguage: %s score is %g\n",  marks, laguage )
	} 
}