package main

import "fmt"

func main() {
	arr := [...]string{"Harare", "Bulawayo", "Mutare", "Gweru",
		"Chitungwiza", "Epworth", "Kwekwe", "Kadoma", "Masvingo",
		"Chinhoyi", "Marondera", "Norton", "Chegutu", "Bindura",
		"Zvishavane", "Victoria Falls", "Hwange", "Redcliff", "Rusape",
		"Chiredzi", "Beitbridge", "Kariba", "Karoi", "Gokwe", "Shurugwi",
		"Gwanda", "Shamva", "Chivhu", "Banket", "Mvurwi", "Lupane", "Mazowe",
		"Glendale", "Binga", "Chakari", "Centenary", "Concession", "Esigodini",
		"Filabusi", "Inyati", "Lalapanzi", "Macheke", "Maphisa", "Nyazura",
		"Raffingora", "Sanyati", "Tsholotsho", "Umguza", "Mberengwa"}

	fmt.Println(arr)
	fmt.Printf("LENGTH: %v\n", len(arr))
	fmt.Printf("%T", arr)
}
