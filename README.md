# parent-child-ipc

Golang parent/child process IPC via standard in/out

### What does it do

Once `parent` executable is ran, it launches `child` executable and then it listens for messages from child process standard output.
The `child` process sends three messages through its standard output.
Those messages are received by `parent` process which prints them to console. 

### How to use it

* `make build` compiles parent and child into bin directory
* `make run` runs parent process which in turn runs child process

### Sample execution

```
$> make run
./bin/parent
2019/08/04 08:47:38 Launching executable /parent-child-ipc/bin/child
2019/08/04 08:47:38 Process started. PID 29858
2019/08/04 08:47:38 Received message: message-0
2019/08/04 08:47:39 Received message: message-1
2019/08/04 08:47:40 Received message: message-2
2019/08/04 08:47:41 Terminated
```