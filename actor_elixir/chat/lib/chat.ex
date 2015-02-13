defmodule Chat do
  use Application

  # See http://elixir-lang.org/docs/stable/Application.Behaviour.html
  # for more information on OTP Applications
  def start(_type, _args) do
    Chat.Supervisor.start_link
    room = %Chat.Room{ clients: [self] }

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

    Enum.each 1..6, fn(x) -> 
      Chat.Client.flush server_pid
    end
  end
end
