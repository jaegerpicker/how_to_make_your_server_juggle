defmodule Chat.Client do
  def join(server) do
    send_message server, :join
  end

  def say(server, message) do
    #IO.puts message
    send_message server, {:say, message}
  end

  def leave(server) do
    send_message server, :leave
  end

  def flush(server) do
    receive do 
      { ^server, {:messgage, message} } ->
          IO.puts message
    end
  end

  defp send_message(server, message) do
    send server, {self, message} 
    IO.puts message
    receive do
      { ^server, response } ->
        response
    after
      1000 ->
        IO.puts "Connection to room timed out"
        :timeout
    end
  end
end
