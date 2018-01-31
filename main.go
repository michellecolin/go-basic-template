package main

import (
	"log"
	"os"
	"text/template"
)

type restaurant struct {
	Name string
	Menu menu
}

type menu []meal

type meal struct {
	Name  string
	Itens []item
}

type item struct {
	Name  string
	Price float64
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := restaurants{
		restaurant{
			Name: "Uncle Joe's",
			Menu: []meal{
				meal{
					Name: "Breakfast",
					Itens: []item{
						item{
							Name:  "Breakfeast burrito",
							Price: 4.99,
						},
						item{
							Name:  "Gravy and biscuits",
							Price: 3.99,
						},
					},
				},
				meal{
					Name: "Dinner",
					Itens: []item{
						item{
							Name:  "Big Burrito",
							Price: 6.99,
						},
						item{
							Name:  "Fried Chicken",
							Price: 5.99,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
