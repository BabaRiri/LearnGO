package main

import "fmt"

func main() {
	zimProvinces := make([]string, 0, 10)
	zimProvinces = append(zimProvinces, "Harare",
										"Bulawayo",
										"Mashonaland Central",
										"Mashonaland East",
										"Mashonaland West",
										"Manicaland",
										"Midlands",
										"Masvingo",
										"Matabeleland North",
										"Matabeleland South",)
						
    fmt.Printf("LENGTH: %v\n", len(zimProvinces))
    fmt.Printf("CAPACITY: %v\n", cap(zimProvinces))

	for i := 0; i < len(zimProvinces); i++ {
		fmt.Printf("INDEX: %v\t VALUE: %v\n", i, zimProvinces[i])
	}
}
