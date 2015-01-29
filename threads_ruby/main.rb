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

class position
    def initialize(row, column)
        @row = row
        @column = column
    end
end

def setup

end
