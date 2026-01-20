package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct{}
type Portuguese struct{}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
func SayHello(name string, greet Greeter) string {
	language := greet.LanguageName()
	message := greet.Greet(name)
	return fmt.Sprintf("I can speak %s: %s", language, message)
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}
func (i Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}


// extra exercises
type Shaper interface {
	Area() float64
}
type Square struct {
	SideLength float64
}

func (s Square) Area() float64 {
	return s.SideLength * s.SideLength
}

type Device interface {
	TurnOn() string
}

type TV struct{}

func (t TV) TurnOn() string {
	return "Turn on Samsung Tv"
}

func PressingPower(device Device) string {
	return fmt.Sprintf("Remote Signal sent: %s", device.TurnOn())
}
