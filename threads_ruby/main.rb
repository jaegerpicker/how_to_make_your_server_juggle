#! /usr/bin/ruby

class board
    def initialize(board_size)
      if board_size < 26 do
        @rows = [*1..board_size]
        alpha = [*"A".."Z"]
        @columns = alpha[0,board_size]
      else
        return "Invalid board_size"
      end
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
