#! /usr/bin/ruby

class player
    def initialize(position_at, player_name, player_number, alive, num_turns, power)
        @position_at = position_at
        @player_name = player_name
        @player_number = player_number
        @alive = alive
        @num_turns = num_turns
        @power = power
    end
end

class playerMove
    def initialize(position, player)
        @position = position
        @player = player
    end
end

class board
    def initialize(rows, columns)
        @rows = rows
        @columns = columns
    end
end

class position
    def initialize(row, column)
        @row = row
        @column = column
    end
end

@board_x = 5
@board_y = 5
@playing_board = board.new([*1..5], [*"A".."E"])


def setup

end
