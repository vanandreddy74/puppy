package puppy

import (
	"github.com/vanandreddy74/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownup(Bark())
}

func BigBarks() string {
	return dog.WhenGrownup(Barks())
}
