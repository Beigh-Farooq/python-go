//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
//"James", "Bond", "Shaken, not stirred" "Miss", "Moneypenny", "Helloooooo, James."
//Range over the records, then range over the data in each record. 
package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
