package simple_factory

import "fmt"

type API interface {
	Say(name string) string
}

type hiAPI struct {

}

type helloAPI struct {

}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

func (a *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func (a *helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}