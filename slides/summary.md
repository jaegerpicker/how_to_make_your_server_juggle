##  Summary

This was a just a tour of the 4 models of concurrency presented here, each one has numerous articles and books written
about them. I hope I sparked you to investigate the different models further and to truly see what modern programming has
to offer with concurrency.

### Personal opinion
This is just my personal opinion and many would disagree with it but I feel like threads and event loops should be rarely
used and only used in a environment that is propose built to use them. Threads are really great for low level systems but
I would have a really hard time seeing a justification to use them in a server side app given the available choices. With
hard real-time and audio-visual processing being the only real exceptions. Event loops are great when you can totally sure
that only libraries meant to be used in that environment are going to be used. In general if you are reaching for concurrency
in a server side application, you should be reaching for CSP or an Actor modeled language. In general I believe that
concurrency should be used and designed for far more often. It's often a great way to vastly improve the efficiency of your server side systems.
