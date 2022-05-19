package main

import "fmt"

type adress struct {
	name   string
	number string
	city   string
}

type people struct {
	name     string
	cpf      string
	adresses []*adress
}

func main() {
	user := newPeople("Jhon", "999.999.999-99")
	user.adresses = append(user.adresses, &adress{name: "centro", number: "34"})
	fmt.Println(user.adresses[0])

}

func newPeople(pName, pCPf string) *people {
	return &people{
		name:     pName,
		cpf:      pCPf,
		adresses: make([]*adress, 0),
	}
}
