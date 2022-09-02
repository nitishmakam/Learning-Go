package main

import "fmt"

type animalInt interface {
	Eat()
	Move()
	Speak()
}

type animal struct {
	food       string
	locomotion string
	noise      string
}

func (ani animal) Eat() {
	fmt.Println(ani.food)
}

func (ani animal) Move() {
	fmt.Println(ani.locomotion)
}

func (ani animal) Speak() {
	fmt.Println(ani.noise)
}

func main() {
	animalMap := make(map[string]animal)
	animalMap["cow"] = animal{"grass", "walk", "moo"}
	animalMap["bird"] = animal{"worms", "fly", "peep"}
	animalMap["snake"] = animal{"mice", "slither", "hsss"}

	animalToNameMap := make(map[string]animal)

	var genericAnimal animalInt
	for {
		var op, ani, reqType string
		fmt.Print(">")
		fmt.Scan(&op, &ani, &reqType)
		if op == "newanimal" {
			animalToNameMap[ani] = animalMap[reqType]
			fmt.Println("Created it!")
		}
		if op == "query" {
			genericAnimal = animalToNameMap[ani]
			switch reqType {
			case "eat":
				genericAnimal.Eat()
			case "move":
				genericAnimal.Move()
			case "speak":
				genericAnimal.Speak()
			}
		}
	}
}
