package main

import (
	_ "sync"
	"fmt"
	"math/rand"
	"time"
)

type player_move struct {
	position        position
	player   		player
}

type player struct {
	position_at		position
	player_name     string
    player_number   int
    alive           bool
    num_turns       int
    power           int
}

type position struct {
	column string
	row    int
}

type board struct {
	rows	[]int
	columns	[]string
}
var board_x int = 5
var board_y int = 5
var playing_board = board{
	rows: make([]int, board_x),
	columns: make([]string, board_y),
}
var write_to_board = make(chan player_move, 0)
var random_positions bool = true
var number_of_players int = 4
var players = make([]player, number_of_players)
var player_name_length int = 5
var letters = []rune("abcdefghijklmnopqrstuvwxyz")


func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func playerInPlayerList(playerlist []player, player string) bool {
	for _, b := range playerlist {
		if b.player_name == player {
			return true
		}
	}
	return false
}

func createNewPlayer(players_list []player) string {
	new_player := randSeq(player_name_length)
	for playerInPlayerList(players_list, new_player) {
		new_player = createNewPlayer(players_list)
	}
	return new_player
}

func setup() {
	rand.Seed(time.Now().UTC().UnixNano())
	for b := 0; b < board_y; b++ {
		playing_board.columns[b] = string(letters[b])
	}
	for r := 0; r < board_x; r++ {
		playing_board.rows[r] = r
	}
	for i :=0; i < number_of_players; i++ {
		new_player_name := createNewPlayer(players)
		rand_column := rand.Intn(board_y)
		rand_row := rand.Intn(board_x)
		new_player := player {
			position_at:  position {row: playing_board.rows[rand_row], column: playing_board.columns[rand_column]},
			player_name: new_player_name,
			player_number: i,
			alive:	true,
			num_turns: 0,
			power: rand.Intn(number_of_players),
		}
		players[i] = new_player
	}
	fmt.Println(fmt.Sprintf("Welcome to the %v deathmatch! \nLet's greet our players", time.Now()))
	for _,i := range players {
		fmt.Println(fmt.Sprintf("Player %s welcome!", i.player_name))
	}
}

func printBoardState() {
	player_there := false
	label := ""
	fmt.Println("Board state:")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	for i := 0; i < board_y; i++ {
		for j:= 0; j < board_x; j++ {
			for _, p := range players {
				if playing_board.columns[i] == p.position_at.column && playing_board.rows[j] == p.position_at.row {
					player_there = true
					label = p.player_name
				}
			}
			if !player_there {
				label = playing_board.columns[i] + fmt.Sprintf("%d", playing_board.rows[j])
			}
			fmt.Printf(" | %5s | ", label)
			player_there = false
		}
		fmt.Println("")
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func playersInBattle(p1 player, p2 player) (bool) {
	if p2.player_name != p.player_name && p2.position_at.row == p.position_at.row && p2.position_at.column == p.position_at.column {
		true
	} else {
		false
	}
}

func playerResolver() {
	p <- write_to_board
	num_of_players_alive := number_of_players
	for i, p2 := range players {
		if p2.player_name == p.player_name {
			players[i] = p
		}
		if playersInBattle(p, p2) {
			if p.power > p2.power {
				p2.alive = false
				num_of_players_alive--
			} else {
				p.power = false
				num_of_players_alive--
			}
		}
	}
	printBoardState()
}

func main() {
	setup()
	printBoardState()
	go playerResolver()
}
