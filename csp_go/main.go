package main

import (
	"bufio"
	"bytes"
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
	Position position `json:"position"`
	Player   player   `json:"player"`
}

type player struct {
	PositionAt   position `json:"positionAt"`
	PlayerName   string   `json:"playerName"`
	PlayerNumber int      `json:"playerNumber"`
	Alive        bool     `json:"alive"`
	NumTurns     int      `json:"numTurns"`
	Power        int      `json:"power"`
}

type position struct {
	Column string `json:"column"`
	Row    int    `json:"row"`
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
var playerMoveResolved = make(chan playerMove, 0)
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
		if b.PlayerName == player {
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
			PositionAt:   position{Row: playingBoard.rows[randRow], Column: playingBoard.columns[randColumn]},
			PlayerName:   newPlayerName,
			PlayerNumber: i,
			Alive:        true,
			NumTurns:     0,
			Power:        rand.Intn(numberOfPlayers),
		}
		players[i] = newPlayer
	}
	fmt.Println(fmt.Sprintf("Welcome to the %v deathmatch! \nLet's greet our players", time.Now()))
	for _, i := range players {
		fmt.Println(fmt.Sprintf("Player %s welcome!", i.PlayerName))
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
				if playingBoard.columns[i] == p.PositionAt.Column && playingBoard.rows[j] == p.PositionAt.Row {
					playerThere = true
					label = p.PlayerName
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
	if p2.PlayerName != p1.PlayerName && p2.PositionAt.Row == p1.PositionAt.Row && p2.PositionAt.Column == p1.PositionAt.Column {
		return true
	}
	return false
}

func playerResolver() {
	var pm playerMove
	var p player
	for {
		pm = <-writeToBoard
		p = pm.Player
		numOfPlayersAlive := numberOfPlayers
		for i, p2 := range players {
			if p2.PlayerName == p.PlayerName {
				players[i] = p
			}
			if playersInBattle(p, p2) {
				if p.Power > p2.Power {
					p2.Alive = false
					numOfPlayersAlive--
				} else {
					p.Alive = false
					numOfPlayersAlive--
				}
			}
		}
		printBoardState()
	}
}

func sendData(conn net.Conn, in <-chan string) {
	defer conn.Close()
	for {
		message := <-in
		log.Print(message)
		io.Copy(conn, bytes.NewBufferString(message))
	}
}

func decreaseColumn(Column string) string {
	if Column == "D" {
		return "C"
	} else if Column == "C" {
		return "B"
	} else if Column == "B" {
		return "A"
	}
	return "Column"
}

func increaseColumn(Column string) string {
	if Column == "A" {
		return "B"
	} else if Column == "B" {
		return "C"
	} else if Column == "C" {
		return "D"
	}
	return Column
}

func handleConnection(conn net.Conn, out chan string) {
	connbuf := bufio.NewReader(conn)
	var playerForConnection player
	for {
		bstr, err := connbuf.ReadBytes('\n')
		if len(bstr) > 0 {
			//fmt.Println(fmt.Sprintf("%v", bstr))
			str := strings.Trim(strings.Trim(string(bstr), "\n"), "\r")
			//fmt.Println(fmt.Sprintf("%v", str))
			if str == "exit" {
				break
			} else if fmt.Sprintf("%v", str) == "connect" {
				message := ""
				for _, p := range players {
					message += p.PlayerName + "\n"
				}
				out <- message
				playerForConnection = players[rand.Intn(len(players))]
				out <- playerForConnection.PlayerName
			} else {
				command := str
				validCommand := false
				if command == "up" {
					if playerForConnection.PositionAt.Row > 0 {
						playerForConnection.PositionAt.Row--
					}
					validCommand = true
				} else if command == "down" {
					if playerForConnection.PositionAt.Row < 3 {
						playerForConnection.PositionAt.Row++
					}
					validCommand = true
				} else if command == "left" {
					if playerForConnection.PositionAt.Column != "D" {
						playerForConnection.PositionAt.Column = increaseColumn(playerForConnection.PositionAt.Column)
					}
					validCommand = true
				} else if command == "right" {
					if playerForConnection.PositionAt.Column != "A" {
						playerForConnection.PositionAt.Column = decreaseColumn(playerForConnection.PositionAt.Column)
					}
					validCommand = true
				}
				if !validCommand {
					message := "unsupported command"
					out <- message
				} else {
					message := playerForConnection.PositionAt.Column + string(playerForConnection.PositionAt.Row)
					out <- message
					pm := playerMove{
						Position: playerForConnection.PositionAt,
						Player:   playerForConnection,
					}
					writeToBoard <- pm
				}
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
		channel := make(chan string)
		go handleConnection(conn, channel)
		go sendData(conn, channel)
	}
}
