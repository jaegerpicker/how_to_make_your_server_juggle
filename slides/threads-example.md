##  Threads Example

```
    semaphore = Mutex.new
    resolver = Thread.new
      loop do
        if @player_move['new_move']
          semaphore.synchronize
            @players.each do |p|
              if p.player_name == @player_move['player'].player_name
                p.position = @player_move['position']
            end
            player_resolver()
            @player_move['new_move'] = false
          end
      end

```
