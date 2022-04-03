package main

import (
	"github.com/milkywoosh/gomod_struct_modif"
)

func main() {
	data := gomod_struct_modif.TestStruct{}
	data.SetName("ben")
	data.SetWealth(100000)
	data.ShowData()

}
