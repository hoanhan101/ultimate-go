// -----------------
// Grouping By State
// -----------------

// This is an example of using type hierarchies with a OOP pattern.
// This is not something we want to do in Go. Go does not have the concept of sub-typing.
// All types are their own and the concepts of base and derived types do not exist in Go.
// This pattern does not provide a good design principle in a Go program.

package main

import "fmt"

// Animal contains all the base fields for animals.
type Animal struct {
	Name     string
	IsMammal bool
}

// Speak provides generic behavior for all animals and how they speak.
// This is kind of useless because animals themselves cannot speak. This cannot apply to all
// animals.
func (a *Animal) Speak() {
	fmt.Println("UGH!",
		"My name is", a.Name,
		", it is", a.IsMammal,
		"I am a mammal")
}

// Dog contains everything an Animal is but specific attributes that only a Dog has.
type Dog struct {
	Animal
	PackFactor int
}

// Speak knows how to speak like a dog.
func (d *Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		", it is", d.IsMammal,
		"I am a mammal with a pack factor of", d.PackFactor)
}

// Cat contains everything an Animal is but specific attributes that only a Cat has.
type Cat struct {
	Animal
	ClimbFactor int
}

// Speak knows how to speak like a cat.
func (c *Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		", it is", c.IsMammal,
		"I am a mammal with a climb factor of", c.ClimbFactor)
}

func main() {
	// It's all fine until this one. This code will not compile.
	// Here, we try to group the Cat and Dog based on the fact that they are Animals. We are trying
	// to leverage sub typing in Go. However, Go doesn't have it.
	// Go doesn't say let group thing by a common DNA.
	// We need to stop designing APIs around this idea that types have a common DNA because if we
	// only focus on who we are, it is very limiting on who can we group with.
	// Sub typing doesn't promote diversity. We lock type in a very small subset that can be
	// grouped with. But when we focus on behavior, we open up entire world to us.
	animals := []Animal{
		// Create a Dog by initializing its Animal parts and then its specific Dog attributes.
		Dog{
			Animal: Animal{
				Name:     "Fido",
				IsMammal: true,
			},
			PackFactor: 5,
		},

		// Create a Cat by initializing its Animal parts and then its specific Cat attributes.
		Cat{
			Animal: Animal{
				Name:     "Milo",
				IsMammal: true,
			},
			ClimbFactor: 4,
		},
	}

	// Have the Animals speak.
	for _, animal := range animals {
		animal.Speak()
	}
}

// ----------
// Conclusion
// ----------

// This code smells bad because:
// - The Animal type is providing an abstraction layer of reusable state.
// - The program never needs to create or solely use a value of type Animal.
// - The implementation of the Speak method for the Animal type is a generalization.
// - The Speak method for the Animal type is never going to be called.
