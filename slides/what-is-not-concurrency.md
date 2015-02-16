##  What is not Concurrency?

Concurrency is not always a way to speed up an application. In fact you will often find that a simpler single path of
execution is often faster when it has discrete inputs and outputs.

###Parallelism

This is also not Concurrency. Parallelism is two or more concerns executing at the same time. In practice this almost
always requires hardware level support like multiple cores or CPU's, GPU's, or a distrubuted network.
Concurrency is almost always used to make Parallelism easier but it is entirely possible to have a concurrent application
but not a parallel executing application.
