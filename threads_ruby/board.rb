class Board
  attr_accessor :rows
  attr_accessor :columns
  @rows = [*1..5]
  @columns = [*"A".."Z"]
  
    def initialize(board_size)
      if board_size < 26
        @rows = [*1..board_size]
        alpha = [*"A".."Z"]
        @columns = alpha[0,board_size]
      else
        return "Invalid board_size"
      end
    end
end
