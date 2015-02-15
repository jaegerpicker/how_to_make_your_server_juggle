##  Wait why?

This happens because your code doesn't, can't, and in no way should control the OS's under lying scheduler. There are far too
many things that go into deciding which code should run next.

### So how do we safely communicate in a concurrent application?

There are many different models but they all revolve around locking the data down so that it's only accessed by one piece
of code at a time.
