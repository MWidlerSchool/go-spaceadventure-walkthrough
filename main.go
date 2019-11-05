package main

import "fmt"

func main() {
	printWelcome()
	printGreeting(queryUser("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	travel()
}

func printWelcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 8 planets to explore.")
}

func queryUser(question string) string {
	var response string
	fmt.Println(question)
	fmt.Scan(&response)
	return response
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func travel() {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = queryUser("Shall I randomly choose a planet for you to visit? (Y or N)")
		if choice == "Y" {
			travelToRandomPlanet()
		} else if choice == "N" {
			planetName := queryUser("Name the planet you would like to visit.")
			travelToPlanet(planetName)
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
}

func travelToRandomPlanet() {
	fmt.Println("Traveling to Jupiter...")
	fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
}

func travelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}
