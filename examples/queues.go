package examples

import (
	"fmt"
	"sorting-algorithms/structures"
)

func Queues() {
	user1 := [2]string{"Bob", "join"}
	user2 := [2]string{"Taylor", "join"}
	user3 := [2]string{"Nouman", "join"}
	user4 := [2]string{"James", "join"}
	usersQueue := new(structures.Queue[string])
	fmt.Println(structures.Matchmake(usersQueue, user1))
	fmt.Println(structures.Matchmake(usersQueue, user2))
	fmt.Println(structures.Matchmake(usersQueue, user3))
	fmt.Println(structures.Matchmake(usersQueue, user4))
	user5 := [2]string{"James", "leave"}
	fmt.Println(structures.Matchmake(usersQueue, user5))
}
