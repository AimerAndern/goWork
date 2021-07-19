package main

import "fmt"

var(
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana","Adriano","Aaro","Elizabeth"

	}
	distribution = make(map[string]int, len(users))
)


func dispatchCoin() int {
	for i:=0;i<len(users);i++{
		for j:=0;j<len(users[i]);j++{
			switch users[i][j] {
			case "e"
			
			}	
		}
	}
}

func main(){
	left :=dispatchCoin()
	fmt.Println("left:",left)
}

