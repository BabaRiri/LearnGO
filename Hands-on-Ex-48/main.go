package main

import "fmt"

func main() {
    countries := []string{"Algeria", "Angola", "Benin", "Botswana", "Burkina Faso", "Burundi", "Cameroon", "Cape Verde", "Central African Republic", "Chad", "Comoros", "Congo (Brazzaville), Republic of the", "Congo (Kinshasa), Democratic Republic of the", "Cote d'Ivoire", "Djibouti", "Egypt", "Equatorial Guinea", "Eritrea", "Eswatini (Swaziland)", "Ethiopia", "Gabon", "Gambia", "Ghana", "Guinea", "Guinea-Bissau", "Kenya", "Lesotho", "Liberia", "Libya", "Madagascar", "Malawi", "Mali", "Mauritania","Mauritius","Morocco","Mozambique","Namibia","Niger","Nigeria","Rwanda","Sao Tome and Principe","Senegal","Seychelles","Sierra Leone","Somalia","South Africa","South Sudan","Sudan","Tanzania","Togo","Tunisia","Uganda","Zambia","Zimbabwe"}
	years := []string{"1962", "1975", "1960", "1966", "1960", "1962", "1960", "1975", "1960", "1960", "1975", "1960", "1960", "1960", "1977", "1922", "1968", "1993", "1968", "1941", "1960", "1965", "1957", "1958", "1973", "1963", "1966", "1847", "1951", "1960", "1964", "1960", "1960" ,"1968" ,"1956" ,"1975" ,"1990" ,"1960" ,"1960" ,"1962" ,"1975" ,"1960" ,"1976" ,"1961" ,"1960" ,"1910" ,"2011" ,"1956" ,"1961" ,"1960" ,"1956" ,"1962" ,"1924" ,"1980"}
	slice := [][]string{countries, years}

	for i, v := range slice {
		fmt.Println(i, v)
		for j, w := range v {
            fmt.Println(j, w)
        }
		
	}
}