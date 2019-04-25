package main

import (
	"errors"
	"fmt"
	"time"
)

type Person struct {
	Name 	string
	Age		int
}

func (person Person) olderThan(age int) bool {
	return person.Age > age
}
func (person *Person) IncreaseAge() error {
	if person.Age == 0 {
		return errors.New("person age not defined") // No capitalization, no punctuation
		// Yes, I know a new-born baby can be 0 - ignore that for this example
	}

	person.Age++

	return nil
}

func main() {
	var person Person
	person.Name = "Timothy"

	currentTime := time.Now()
	person.Age = (currentTime.Year()-1990)
	//person.Age = 0

	fmt.Println(person)
	fmt.Println("Name:", person.Name)
	fmt.Printf("Age: %d \n", person.Age)

	err := person.IncreaseAge()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Age: %d", person.Age)
	}

}
