package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (cow Cow) Eat() {
	fmt.Println(cow.food)
}

func (bird Bird) Eat() {
	fmt.Println(bird.food)
}

func (snake Snake) Eat() {
	fmt.Println(snake.food)
}

func (cow Cow) Move() {
	fmt.Println(cow.locomotion)
}

func (bird Bird) Move() {
	fmt.Println(bird.locomotion)
}
func (snake Snake) Move() {
	fmt.Println(snake.locomotion)
}
func (cow Cow) Speak() {
	fmt.Println(cow.noise)
}

func (bird Bird) Speak() {
	fmt.Println(bird.noise)
}

func (snake Snake) Speak() {
	fmt.Println(snake.noise)
}

func main() {
	animalToNameMap := make(map[string]Animal)

	var genericAnimal Animal
	for {
		var op, ani, reqType string
		fmt.Print(">")
		fmt.Scan(&op, &ani, &reqType)
		if op == "newanimal" {
			var animalObj Animal
			switch reqType {
			case "cow":
				animalObj = Cow{"grass", "walk", "moo"}
			case "bird":
				animalObj = Bird{"worms", "fly", "peep"}
			case "snake":
				animalObj = Snake{"mice", "slither", "hsss"}
			}
			animalToNameMap[ani] = animalObj
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
