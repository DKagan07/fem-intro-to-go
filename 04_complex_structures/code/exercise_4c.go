package main

import "fmt"

//part1
// func averageArr(arr [5]float64) float64 {
// 	total := 0.0

// 	for _, val := range arr {
// 		total += val
// 	}

// 	return total / float64(len(arr))
// }

//part2
// func confirmPetNames(name string, daMap map[string]string) bool {
// 	_, ok := daMap[name];
// 	return ok;
// }

//part3
func addGroceries (list []string, adds ...string) []string{

	for _, val := range adds {
		list = append(list, val);
	}

	return list;
}


func main() {
	// Part 1
	// scores := [5]float64 { 2.1, 3.2, 5.4, 6.7, 9.1 };

	// fmt.Println(averageArr(scores))


	//Part 2
	// petNames := make(map[string] string);

	// petNames["Fido"] = "Dog";
	// petNames["Chester"] = "Cat"


	// fmt.Println(confirmPetNames("Fido", petNames));
	// fmt.Println(confirmPetNames("Spot", petNames));



	//Part 3
	// groceries := make([]string, 5);
	groceries := []string {"bread", "veggies"}

	groceries = addGroceries(groceries, "eggs", "protein", "milk")
	
	fmt.Println(groceries)
}
