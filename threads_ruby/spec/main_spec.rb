require './board.rb'
require './main.rb'

RSpec.describe Board do
  describe "#board.initialize" do
    it "returns a board with the number of rows and columns I pass in" do
         test_board = Board.new(3)
         expect(test_board.columns[0]).to eq("A")
         expect(test_board.rows[2]).to eq(3)
         expect(test_board.rows.length).to eq(3)
         expect(test_board.columns.length).to eq(3)
    end
  end
end

RSpec.describe Game do
  describe "#main.setup" do
    it "creates a array of random playnames eq to number_of_players" do
      m = Game.new
      expect(m.players.length).to eq(m.number_of_players)
      expect(m.players.empty?).to eq(false)
    end
  end
end
         
