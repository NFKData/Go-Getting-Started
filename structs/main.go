package main

import "fmt"

type Player struct {
	name   string
	hp     int
	damage int
}

type Enemy struct {
	name   string
	hp     int
	damage int
}

func main() {
	var player Player
	player.name = "NFKData"
	player.hp = 100
	player.damage = 20

	var lime Enemy
	lime.name = "lime"
	lime.hp = 15
	lime.damage = 2

	var epicLime Enemy
	epicLime.name = "Epic Lime"
	epicLime.hp = 60
	epicLime.damage = 10

	fmt.Println(player, lime, epicLime)

	player2 := Player{
		name:   "NFKData",
		hp:     100,
		damage: 20,
	}

	lime2 := Enemy{
		name:   "Lime",
		hp:     15,
		damage: 2,
	}

	epicLime2 := Enemy{
		name:   "Epic Lime",
		hp:     60,
		damage: 10,
	}

	fmt.Println(player2, lime2, epicLime2)
}
