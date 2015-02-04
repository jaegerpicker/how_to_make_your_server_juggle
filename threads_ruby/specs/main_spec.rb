require './board.rb'

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
         
