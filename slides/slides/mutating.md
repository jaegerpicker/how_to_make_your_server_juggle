##  Mutating

*psuedocode
```
var global_var = 7;
var t = new Thread(function(){
        global_var = 5;
        println global_var * 5;
    });
var t2 = new Thread(function(){
        println global_var * 2;
    });

t.start();
t2.start();
```
```
Output #1:
14
25
```
```
Output #2:
25
10
```
