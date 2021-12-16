package main

import (
	"fmt"
)

func main() {

	// isi field struct Group
	user := Group{"Sahibul", "Sa", true}

	// memanggil method
	result := user.displayGroup()
	fmt.Println(result)
}

type Group struct {
	Name        string
	Admin       string
	IsAvailable bool
}

// method
func (group Group) displayGroup() string {
	return fmt.Sprintf("Name : %s", group.Name)
}
