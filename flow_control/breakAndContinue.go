package main
import ("fmt")

func main() {
	/*
		* break statement terminates the loop while
		* continue keeps the current iteration of the loop

		for initialization; condition; update {
			break
		}
	*/

	for i := 1; i <= 5; i ++ {

		// if i is equal to 3, terminate the loop
		if i == 3 {
			break
		}
		fmt.Println(i)
	} 


	//  # break nested loops

	for x := 1; x <= 3; x ++ {

		for j := 1; j <= 3; j ++ {
			if x == 2 {
				break
			}
			fmt.Println("x=", x, "j=", j)
		}
	}


	/*
		* Go Continue
	*/


	for m := 1; m <=5; m ++ {
		if m == 3 {
			continue
		}

		fmt.Println(m)
	}
	//  4 is not printed and it continues to 4
}