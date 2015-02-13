defmodule Chat.Server do
  def loop(room) do
    receive do
      {pid, :join} ->
        IO.puts "Joined"
        notify_all room, self, "User with pid #{inspect pid} joined"
        send pid, {self, :ok }
        updated_room = put_in room.clients, List.flatten(room.clients, [pid])

        loop updated_room
      {pid, {:say, message}} ->
        #IO.puts message
        notify_all room, pid, "#{inspect pid}: " <> message
        send pid, {self, :ok}
        loop room
      {pid, :leave} ->
        send pid, {self, :ok}
        updated_room = put_in room.clients, List.delete(room.clients, pid)
        notify_all updated_room, self, "User with pid #{inspect pid} left"
        loop updated_room
    end
  end

  defp notify_all(room, sender, message) do
    Enum.each room.clients, fn(client_pid) ->
      IO.puts message
      if client_pid != sender do
        send client_pid, {self, {:message, message}}
      end
    end
  end
end
