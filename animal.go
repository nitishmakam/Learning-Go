package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}
func main() {
	cow := Animal{"grass", "walk", "moo"}

	bird := Animal{"worms", "fly", "peep"}

	snake := Animal{"mice", "slither", "hiss"}

	var animalType string
	var info string

	fmt.Print("Enter the animal name followed by the information required separated by a space:\n")
	for {
		fmt.Print(">")
		fmt.Scanf("%s %s", &animalType, &info)
		var animal Animal
		switch animalType {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			fmt.Println("Invalid animal type.")
			continue
		}
		switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid information request.")
			continue
		}

	}
}
