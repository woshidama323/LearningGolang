#### note 
```shell
调度实验
设计：
利用环境变量看看 当前的golang的调度一些细节


scheddetail: setting schedtrace=X and scheddetail=1 causes the scheduler to emit
detailed multiline info every X milliseconds, describing state of the scheduler,
processors, threads and goroutines.

GODEBUG="scheddetail=1,schedtrace=1000" ./LearningGolang runtime >>log 2>&1


## Advanced Note:
Think of a processor in this context as a logical processor and not a physical processor. The scheduler
runs goroutines on these logical processors which are bound to a physical processor via the operating
system thread that is attached. The operating system will schedule the thread against any physical
processor that is available.

In many cases, goroutines are not moved back to the global run queue prior to being terminated. This program has created a special situation because the for loop is performing logic that runs for more than 10ms and is not calling into any functions. 10ms is the scheduling quant in the scheduler. After 10ms of execution, the scheduler tries to preempt goroutines. These goroutines can’t be preempted because they do not call into any functions. In this case, once the goroutines reach the wg.Done call, the goroutines are instantly preempted and moved to the global run queue for termination.

## 输出的结果

SCHED 10062ms: ## Time in milliseconds since the program started.  This is the trace for the 1 second mark.
gomaxprocs=12  ## Number of processors configured. 12 processor is configured for this program.
idleprocs=12   

threads=5  ## Number of threads that the runtime is managing.Three threads exist. One for the processor and 2 others used by the runtime.
spinningthreads=0  ### suspended thread 
idlethreads=3 
runqueue=0 ## Number of goroutines in the global run queue
gcwaiting=0 
nmidlelocked=0 
stopwait=0 
sysmonwait=0 



### runtime 的几个状态
status: http://golang.org/src/runtime/
Gidle,            // 0
Grunnable,        // 1 runnable and on a run queue
Grunning,         // 2 running
Gsyscall,         // 3 performing a syscall
Gwaiting,         // 4 waiting for the runtime
Gmoribund_unused, // 5 currently unused, but hardcoded in gdb scripts
Gdead,            // 6 goroutine is dead
Genqueue,         // 7 only the Gscanenqueue is used
Gcopystack,       // 8 in this state when newstack is moving the stack



  P0: status=0 schedtick=2 syscalltick=1 m=-1 runqsize=0 gfreecnt=0 timerslen=1
  P1: status=0 schedtick=5 syscalltick=11 m=-1 runqsize=0 gfreecnt=0 timerslen=2
  P2: status=0 schedtick=0 syscalltick=1 m=-1 runqsize=0 gfreecnt=0 timerslen=1
  P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0 timerslen=0
  P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0 timerslen=0
  P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0 timerslen=0
  P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0 timerslen=0
  P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0 timerslen=0
  P8: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0 timerslen=0
  P9: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0 timerslen=0
  P10: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0 timerslen=0
  P11: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0 timerslen=0
  M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=false lockedg=-1
  M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=2 dying=0 spinning=false blocked=false lockedg=-1
  M0: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  G1: status=4(semacquire) m=-1 lockedm=-1
  G2: status=4(force gc (idle)) m=-1 lockedm=-1
  G3: status=4(GC sweep wait) m=-1 lockedm=-1
  G4: status=4(GC scavenge wait) m=-1 lockedm=-1
  G17: status=4(finalizer wait) m=-1 lockedm=-1
  G33: status=4(sleep) m=-1 lockedm=-1
  G34: status=4(sleep) m=-1 lockedm=-1
  G35: status=4(sleep) m=-1 lockedm=-1
  G36: status=4(sleep) m=-1 lockedm=-1
```

[runtime](https://pkg.go.dev/runtime#hdr-Environment_Variables)


### data race detector的功能
```shell
Data races are among the most common and hardest to debug types of bugs in concurrent systems. A data race occurs when two goroutines access the same variable concurrently and at least one of the accesses is a write. See the The Go Memory Model for details.
```

[The Go Memory Model](https://golang.org/ref/mem/)



### [golang debug 技巧](https://www.ardanlabs.com/blog/2015/02/scheduler-tracing-in-go.html)
[字段描述](https://programming.vip/docs/looking-at-scheduling-tracking-with-godebug.html)
```shell


M
P: Which P belongs to?
curg: Which G is currently in use?
runqsize: The number of G in the running queue.
gfreecnt: Available G (in Gdead state).
mallocing: Is memory being allocated?
Throwing: throwing an exception.
preemptoff: If it's not an empty string, keep curg running on this m.

P
status: The running state of P.
Schedule tick: The number of schedules in P.
syscalltick: The number of system calls in P.
M: Which M belongs to?
runqsize: The number of G in the running queue.
gfreecnt: Available G (in Gdead state).

G
status: The running state of G.
M: Which M belongs to?
lockedm: Is there a lock M?

```

  
  G2: status=4(force gc (idle)) m=-1 lockedm=-1
  G3: status=4(GC sweep wait) m=-1 lockedm=-1
  G4: status=4(GC scavenge wait) m=-1 lockedm=-1

[G1: status=4(semacquire) m=-1 lockedm=-1]
[G3: status=4(GC sweep wait) m=-1 lockedm=-1]
[G4: status=4(GC scavenge wait) m=-1 lockedm=-1]((https://golang.org/src/runtime/mgcscavenge.go))
[G17: status=4(finalizer wait) m=-1 lockedm=-1](https://medium.com/a-journey-with-go/go-finalizers-786df8e17687)