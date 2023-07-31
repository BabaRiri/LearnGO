package main

import "fmt"

func main() {
	//creating a map using make function and assigning values to the map of type string with values of type slice
	myMap := make(map[string][]string)
	myMap["Go"] = []string{"Cloud Computing", "Networking", "System Administration"}
	myMap["Python"] = []string{"Data Science", "Machine Learning", "Web Development", "Finance"}
	myMap["JavaScript"] = []string{"SaaS", "FinTech", "Cloud Tech", "Web Development"}
	myMap["Rust"] = []string{"Cloud Computing", "Networking", "System Administration"}
	myMap["R"] = []string{"System Programming", "Web Development", "Networking"}
	myMap["Kotlin"] = []string{"Information Technology and Services", "Computer Software", "Internet"}
	myMap["C++"] = []string{"Operating Systems and Systems Programming", "Databases", "Machine Learning Tools", "Fintech and Banking"}

	fmt.Println("--------------------------------")
	fmt.Println(myMap)
	fmt.Println("--------------------------------")
	for k, v := range myMap {
		fmt.Printf("%v\n", k)
		for i, v2 := range v {
			fmt.Printf("%v - %v\n", i+1, v2)
		}
		fmt.Printf("\n")
	}
	fmt.Println("--------------------------------")
}
