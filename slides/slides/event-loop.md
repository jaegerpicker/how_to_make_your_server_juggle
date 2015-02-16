##  Event Loop

An event loop is a special type of the Threads model. A event loop is a single thread that loops and should only
ever run async code. You register events that are triggered when a different set of criteria is met. This is recently made
most popular via Node.js but has been used for a long time in UI development. Typically, if you need to scale this model,
you start new processes (workers) on the same box. This makes sharing state very difficult and has given rise, in part, to
the popularity of immutable functional programming. While node is the most well known example right now, nginx is also a
very popular example of a event loop based application.
