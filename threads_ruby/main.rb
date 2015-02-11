#! /usr/bin/ruby
require './board.rb'
require './player.rb'
require 'thread'
require 'socket'

class Position
  attr_accessor :row
  attr_accessor :column
  def initialize(row, column)
    @row = row
    @column = column
  end
end

class Game
  attr_accessor :board_x
  attr_accessor :board_y
  attr_accessor :playing_board
  attr_accessor :random_positions
  attr_accessor :number_of_players
  attr_accessor :players
  attr_accessor :players_name_length
  attr_accessor :player_move

  def print_board_state 
    player_there = false
    label = ""
    puts "++++++++++++++++++++++++++++++++++++++++++++++++++++++"
    @playing_board.columns.each { |column|
      @playing_board.rows.each { |row|
        @players.each { |p|
          #puts "\n" + p.position_at.column + p.position_at.row.to_s + " " + p.player_name + " " + column + row.to_s 
          if p.position_at.column == column and p.position_at.row == row 
            label = p.player_name
            player_there = true
          elsif player_there == false
            label = column + row.to_s
          end
        }
        print " | %5.5s | " % label
        player_there = false
      }
      print "\n"
    }
    puts "+++++++++++++++++++++++++++++++++++++++++++++++++++++++"
  end

  def increase_column(column)
    if @playing_board.columns.include? column.next
      return column.next
    else
      return column
    end
  end

  def decrease_column(column)
    hash = Hash[@playing_board.columns.map.with_index.to_a]
    if hash[column] == 0
      return column
    else
      return @playing_board.columns[hash[column] - 1]
    end
  end

  def players_in_battle(p1, p2)
    if p1.player_name == p2.player_name and p1.position_at.row == p2.position_at.row and p1.position_at.column == p2.position_at.column
      return true
    else
      return false
    end
  end

  def player_resolver
    loop do
      player_move = queue.pop
      if player_move != nil
        p = player_move.player
        number_of_players_alive = @number_of_players
        @players.each do |p2|
          if players_in_battle(p, p2)
            if p.power > p2.power
              p2.alive = false
              number_of_players_alive -= 1
            else
              p.alive = false
              number_of_players_alive -= 1
            end
          end
        end
        print_board_state()
      end
    end
  end

  def initialize
    @board_x = 5
    @board_y = 5
    @playing_board = Board.new(5)
    @random_positions = true
    @number_of_players = 4
    @players = [@number_of_players]
    @players_name_length = 5
    @player_move = {"player" => Player.new, "position" => [], "new_move" => false}
    letters = [*"A".."Z"]                    
    (0..(@board_y -1)).each do |i|
      @playing_board.columns[i] = letters[i]
    end
    (0..(@board_x - 1)).each do |i|
      @playing_board.rows[i] = i
    end
    (0..(@number_of_players -1)).each do |i|
      new_player_name = letters.sample(@players_name_length).join
      loop do
        break if not players.include?(new_player_name)
        new_player_name = letters.sample(@players_name_length).join
      end
      pos = Position.new(@playing_board.rows.sample, @playing_board.columns.sample)
      power = rand(@number_of_players)
      p = Player.new(pos, new_player_name, i, true, 0, power)
      @players[i] = p
    end
    print_board_state()
  end

  def run_game
    semaphore = Mutex.new
    resolver = Thread.new 
      loop do
        if @player_move['new_move']
          semaphore.synchronize 
            @players.each do |p|
              if p.player_name == @player_move['player'].player_name
                p.position = @player_move['position']
            end
            player_resolver()
            @player_move['new_move'] = false
          end
      end
    end

    sockets = Thread.new 
          server = TCPServer.new 6543
          player_for_connection = @players[rand(@number_of_players)]
          loop do
            client = server.accept
            input = server.gets
            semaphore.synchronize
              case input
              when "connect"
                @players.each do |p|
                  client.puts p.player_name
                end
                client.puts "Your player is " + player_for_connection.player_name 
              when "up"
                if player_for_connection.position.row != 0
                  player_for_connection.position.row -= 1
                  client.puts player_for_connection.position
                end
              when "down"
                if player_for_connection.position.row != (@board_y - 1)
                  player_for_connection.position.row += 1
                  client.puts player_for_connection.position
                end
              when "left"
                if player_for_connection.position.column != @playing_board.columns[0]
                  player_for_connection.position.column = decrease_column()
                  client.puts player_for_connection.position
                end
              when "right"
                if player_for_connection.position.column != @playing_board.columns[-1]
                  player_for_connection.position.column = increase_column()
                  client.puts player_for_connection.position
                end
              when "exit"
                break
              else
                client.puts "Invalid command"
              end
            end
          end
    end
    
  end

end

m = Game.new
m.run_game()
