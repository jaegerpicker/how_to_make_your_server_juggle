##  Threads

### No school like the old school

Threads are the oldest and most common model. Typically it involves mutex's and semaphore's to lock data.

* This process is often difficult and error-prone. It also does nothing to address the problems of mutable data being changed
out from under another thread. I personally believe that if you are using threads you have likely chosen a poor environment
for anything but very low-level systems work. Web apps and network servers almost always have better options.
