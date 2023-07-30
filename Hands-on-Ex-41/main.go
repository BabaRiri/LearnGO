package main

import "fmt"

func main() {
	//Cities/Towns in Zimbabwe
	slice1 := []string{"Harare", "Bulawayo", "Mutare", "Gweru",
		"Chitungwiza", "Epworth", "Kwekwe", "Kadoma", "Masvingo",
		"Chinhoyi", "Marondera", "Norton", "Chegutu", "Bindura",
		"Zvishavane", "Victoria Falls", "Hwange", "Redcliff", "Rusape",
		"Chiredzi", "Beitbridge", "Kariba", "Karoi", "Gokwe", "Shurugwi",
		"Gwanda", "Shamva", "Chivhu", "Banket", "Mvurwi", "Lupane", "Mazowe",
		"Glendale", "Binga", "Chakari", "Centenary", "Concession", "Esigodini",
		"Filabusi", "Inyati", "Lalapanzi", "Macheke", "Maphisa", "Nyazura",
		"Raffingora", "Sanyati", "Tsholotsho", "Umguza", "Mberengwa"}

	fmt.Println(slice1)
	fmt.Printf("LENGTH: %v\n", len(slice1))

	for i, city := range slice1 {
		fmt.Printf("%v - %v\n", i, city)
	}
}
