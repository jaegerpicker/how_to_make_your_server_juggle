class Player
  attr_accessor :position_at
  attr_accessor :player_name
  attr_accessor :player_number
  attr_accessor :alive
  attr_accessor :num_turns
  attr_accessor :power
  def initialize(position_at, player_name, player_number, alive, num_turns, power)
    @position_at = position_at
    @player_name = player_name
    @player_number = player_number
    @alive = alive
    @num_turns = num_turns
    @power = power
  end
end
