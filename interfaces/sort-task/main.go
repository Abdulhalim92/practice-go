package main

import (
	"fmt"
	"sort"
)

type User struct {
	firstname     string
	lastname      string
	totalTurnover float64
}

type Users []User

func (users Users) Len() int {
	return len(users)
}

func (users Users) Less(i, j int) bool {
	return users[i].totalTurnover < users[j].totalTurnover
}

func (users Users) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

func main() {
	user0 := User{
		firstname:     "John",
		lastname:      "Doe",
		totalTurnover: 1000,
	}
	user1 := User{
		firstname:     "Dany",
		lastname:      "Boyu",
		totalTurnover: 20000,
	}
	user2 := User{
		firstname:     "Elissa",
		lastname:      "Smith Brown",
		totalTurnover: 70,
	}

	users := make(Users, 3)
	users[0] = user0
	users[1] = user1
	users[2] = user2

	sort.Sort(users)
	fmt.Println(users)
}
