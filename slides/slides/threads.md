##  Threads

### No school like the old school

Threads are the oldest and most common model. Typically it involves mutex's and semaphore's to lock data.

Typical flow:
Work gets divided into chunks and a mutex is created. A mutex is a flag signal other threads that the owner of the mutex
is the currently operating thread. While the mutex is locked no other thread can access it's data. A semaphore is used
to request access to the mutex's data.

* This process is often difficult and error prone. It also does nothing to address the problems of mutable data being changed
out from under another thread. I personally believe that if you are using threads you have likely chosen and poor environment
for anything but very low level systems work. Web apps and network servers almost always have better options.
