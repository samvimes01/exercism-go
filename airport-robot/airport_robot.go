package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
  LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
  return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct {}

func (i Italian) Greet(name string) string {
  return "Ciao " + name + "!"
}
func (i Italian) LanguageName() string {
  return "Italian"
}

type Portuguese struct {}

func (i Portuguese) Greet(name string) string {
  return "Olá " + name + "!"
}
func (i Portuguese) LanguageName() string {
  return "Portuguese"
}