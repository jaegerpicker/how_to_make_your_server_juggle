##  Actor Example

```
server_pid = spawn(Chat.Server, :loop,  [room])

spawn fn() ->
Chat.Client.join server_pid
Chat.Client.say server_pid, "Hi!"
Chat.Client.leave server_pid
end

spawn fn() ->
Chat.Client.join server_pid
Chat.Client.say server_pid, "Hi from another process"
Chat.Client.leave server_pid
end

....

receive do
{pid, :join} ->
IO.puts "Joined"
notify_all room, self, "User with pid #{inspect pid} joined"
send pid, {self, :ok }
updated_room = put_in room.clients, List.flatten(room.clients, [pid])

```
