package main

import (
	"fmt"
	"scene/aoi_list"
)

func main() {
	scene := aoi_list.NewSceneAOIManagerByLink(1, 45, 31)
	// player := NewPlayer(1, 5)
	// player2 := NewPlayer(2, 3)
	// player3 := NewPlayer(3, 9)
	// player.Enter(1, 2, 2)
	// player2.Enter(1, 5, 3)
	// player3.Enter(1, 42, 3)
	// player2.Leave(1)
	scene.Travel()
	fmt.Println("--------------------")
	player := aoi_list.NewPlayer(1, 1)
	player.Enter(1, 2, 2)
	scene.Travel()
	fmt.Println("--------------------")
	player.Move(1, 1, 1)
	// player3.Move(1, 5, 3)
	scene.Travel()
	fmt.Println("--------------------")
	player.Leave(1)
	scene.Travel()
}
