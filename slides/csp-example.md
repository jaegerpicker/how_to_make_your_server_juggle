##  CSP Example

```
} else if fmt.Sprintf("%v", str) == "connect" {
    message := ""
    for _, p := range players {
        message += p.PlayerName + "\n"
    }
    out <- message //write a message to send to the client
    playerForConnection = players[rand.Intn(len(players))]
    out <- playerForConnection.PlayerName

    ....

    func sendData(conn net.Conn, in <-chan string) {
        defer conn.Close()
        for {
            message := <-in // wait on a message to arrive
            log.Print(message)
            io.Copy(conn, bytes.NewBufferString(message))
        }
    }
```
