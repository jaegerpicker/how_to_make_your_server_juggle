net = require('net');
_ = require('lodash');

var clients = [];
function player(position, name, number, alive, num_turns, power) {
  return { "position": position,
           "player_name": name,
           "player_number": number, 
           "alive": alive,
           "num_turns": num_turns,
           "power": power
         }
}

function player_move(player, position) {
  return { "player": player,
           "position": position
         }
}

function position(row, column) {
  return { "column": column,
           "row": row
         }
}

function board(rows, columns) {
  return { "rows": rows,
           "columns": columns
         }
}
var board_x = 5;
var board_y = 5;
var playing_board = board([board_x], [board_y]);
var random_positions = true;
var number_of_players = 4;
var players = [number_of_players];
var player_name_length = 5;

function setup() {
  playing_board.rows = _.range(board_x);
  playing_board.columns = _.range(board_y);
  playing_board.columns = playing_board.columns.map(function(board_y, i) { String.fromCharCode(65 + i); });
  
}


net.createServer(function(socket) {
  socket.name = socket.remoteAddress + ':' + socket.remotePort;

  clients.push(socket);
  
  players.forEach(function (player) {
    socket.write(player.player_name);
  });
  var player_for_connection = players[Math.Floor((Math.random() * number_of_players) + 1)];
  socket.write(player_for_connection.player_name);

  
  socket.on('data', function(data) {
    switch(data) {
      case 'up':
        if(player_for_connection.position.row > 0) {
          player_for_connection.position.row--;
        }
        break;
      case 'down':
        if(player_for_connection.position.row > playing_board.rows[0]) {
          player_for_connection.position.row++;
        }
        break;
      case 'left':
        if(player_for_connection.position.column > playing_board.columns[0]) {
          player_for_connection.position.column = playing_board.columns[playing_board.columns.indexOf(player_for_connection.position.column) - 1];
        }
        break;
      case 'right':
        if(player_for_connection.position.column < playing_board.columns[playing_board.columns.length - 1]) {
          player_for_connection.position.column = playing_board.columns[playing_board.columns.indexOf(player_for_connection.position.column) + 1];
        }
        break;
      case 'exit':
        process.stdout.write("Client " + player_for_connection.player_name + " disconnected");
        break;
      default:
        socket.write('Invalid command');
        break;
    }
    print_game_board(player_for_connection);
  });
  
}).listen(6543);
