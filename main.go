package main

import "fmt"

type product struct {
	name string
	cost int
}

func main() {
	brick := product{name: "brick", cost: 10}
	corn := product{name: "corn", cost: 6}
	gamecube := product{name: "gamecube", cost: 100}

	inventory := []product{brick, corn, gamecube}

	fmt.Println("Welcome to the store!")
	fmt.Println("Options:")
	fmt.Println("(i)nventory")
	fmt.Println("(a)dd item")
	fmt.Println("(q)uit")

	var choice string
	fmt.Scanln(&choice)

	for choice != "q" {
		switch choice {
		case "i":
			for _, p := range inventory {
				fmt.Println("name:", p.name)
				fmt.Println("cost:", p.cost)
				fmt.Println()
			}
			break
		case "a":
			var newName string
			fmt.Print("Enter a product name: ")
			fmt.Scanln(&newName)

			var newCost int
			fmt.Print("Enter a cost: ")
			fmt.Scanf("%d", &newCost)

			inventory = append(inventory, product{name: newName, cost: newCost})
			break
		}

		fmt.Println("Welcome to the store!")
		fmt.Println("Options:")
		fmt.Println("(i)nventory")
		fmt.Println("(a)dd item")
		fmt.Println("(q)uit")
		fmt.Scanln(&choice)
	}
}
