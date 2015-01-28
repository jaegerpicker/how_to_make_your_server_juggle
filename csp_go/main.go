package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"
	_ "sync"
	"time"
)

type playerMove struct {
	position position
	player   player
}

type player struct {
	positionAt   position
	playerName   string
	playerNumber int
	alive        bool
	numTurns     int
	power        int
}

type position struct {
	column string
	row    int
}

type board struct {
	rows    []int
	columns []string
}

var boardX = 5
var boardY = 5
var playingBoard = board{
	rows:    make([]int, boardX),
	columns: make([]string, boardY),
}
var writeToBoard = make(chan playerMove, 0)
var randomPositions = true
var numberOfPlayers = 4
var players = make([]player, numberOfPlayers)
var playerNameLength = 5
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
		if b.playerName == player {
			return true
		}
	}
	return false
}

func createNewPlayer(playersList []player) string {
	newPlayer := randSeq(playerNameLength)
	for playerInPlayerList(playersList, newPlayer) {
		newPlayer = createNewPlayer(playersList)
	}
	return newPlayer
}

func setup() {
	rand.Seed(time.Now().UTC().UnixNano())
	for b := 0; b < boardY; b++ {
		playingBoard.columns[b] = string(letters[b])
	}
	for r := 0; r < boardX; r++ {
		playingBoard.rows[r] = r
	}
	for i := 0; i < numberOfPlayers; i++ {
		newPlayerName := createNewPlayer(players)
		randColumn := rand.Intn(boardY)
		randRow := rand.Intn(boardX)
		newPlayer := player{
			positionAt:   position{row: playingBoard.rows[randRow], column: playingBoard.columns[randColumn]},
			playerName:   newPlayerName,
			playerNumber: i,
			alive:        true,
			numTurns:     0,
			power:        rand.Intn(numberOfPlayers),
		}
		players[i] = newPlayer
	}
	fmt.Println(fmt.Sprintf("Welcome to the %v deathmatch! \nLet's greet our players", time.Now()))
	for _, i := range players {
		fmt.Println(fmt.Sprintf("Player %s welcome!", i.playerName))
	}
}

func printBoardState() {
	playerThere := false
	label := ""
	fmt.Println("Board state:")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	for i := 0; i < boardY; i++ {
		for j := 0; j < boardX; j++ {
			for _, p := range players {
				if playingBoard.columns[i] == p.positionAt.column && playingBoard.rows[j] == p.positionAt.row {
					playerThere = true
					label = p.playerName
				}
			}
			if !playerThere {
				label = playingBoard.columns[i] + fmt.Sprintf("%d", playingBoard.rows[j])
			}
			fmt.Printf(" | %5s | ", label)
			playerThere = false
		}
		fmt.Println("")
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func playersInBattle(p1 player, p2 player) bool {
	if p2.playerName != p1.playerName && p2.positionAt.row == p1.positionAt.row && p2.positionAt.column == p1.positionAt.column {
		return true
	}
	return false
}

func playerResolver() {
	var pm playerMove
	var p player
	pm = <-writeToBoard
	p = pm.player
	numOfPlayersAlive := numberOfPlayers
	for i, p2 := range players {
		if p2.playerName == p.playerName {
			players[i] = p
		}
		if playersInBattle(p, p2) {
			if p.power > p2.power {
				p2.alive = false
				numOfPlayersAlive--
			} else {
				p.alive = false
				numOfPlayersAlive--
			}
		}
	}
	printBoardState()
}

func handleConnection(conn net.Conn) {
	connbuf := bufio.NewReader(conn)
	for {
		str, err := connbuf.ReadString('\n')
		if len(str) > 0 {
			if strings.Trim(str, "\n") == "exit" {
				break
			}
			fmt.Println(str)
		}
		if err != nil {
			break
		}
	}
	io.Copy(conn, conn)
	conn.Close()
}

func main() {
	setup()
	printBoardState()
	go playerResolver()

	ln, err := net.Listen("tcp", ":9393")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
