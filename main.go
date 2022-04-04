package main

import (
	"fmt"
	// _ "github.com/milkywoosh/gomod_slice"
	// _ "github.com/milkywoosh/gomod_struct_modif"
)

type Humandroid struct {
	id        int
	name      string
	component map[string]string
	ability   []string
}

func main() {
	prototype1 := Humandroid{}
	prototype1.id = 1
	prototype1.name = "franky"
	prototype1.component = map[string]string{"tes": "tes1"}
	prototype1.ability = []string{"speak", "count", "fly"}

	fmt.Println(prototype1)
}
