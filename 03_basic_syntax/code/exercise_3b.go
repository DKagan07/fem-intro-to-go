package main;

import "fmt";

func main () {
	var str1 string
	var str2 string
	var yrs int
	var tfbool bool

	fmt.Println("Tell me you name, Hometown/State, number of years you've lived there, and if the weather is amazing (boolean).")
	nums, err := fmt.Scan(&str1, &str2, &yrs, &tfbool);

	if err != nil {
		fmt.Println(err.Error());
	}

	fmt.Println("num of args: ", nums);

	fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. People say the weather is good, which is %t.", str1, str2, yrs, tfbool);
}