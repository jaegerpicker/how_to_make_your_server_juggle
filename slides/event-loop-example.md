##  Event Loop Example

```
socket.on('data', function(data) {
    switch(data) {
        case 'up':
        if(player_for_connection.position.row > 0) {
            player_for_connection.position.row--;
        }
        break;
    }
}        
        ```
