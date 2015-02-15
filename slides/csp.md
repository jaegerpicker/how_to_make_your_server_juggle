##  CSP

Communicating Sequential Processes. I.E. writing to channels. Think of it in almost a pub/sub type of structure. One
co-routine* writes to the channel and all of the co-routines listening to that channel get the value. You can send complex
or simple objects on the channels and in some languages even pass the channels them selves around. Go is the current best
example of this model. I highly recommend this model for server side development. It's both flexible and easy to understand.

* Co-routines are also often called green threads or greenlets. Small and light processes that are in most cases (Python and
    Ruby being the big exceptions) multiplexed across OS level threads.
