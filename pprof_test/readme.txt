web界面:
block:查看导致阻塞同步的堆栈跟踪
goroutine:查看当前所有运行的 goroutines 堆栈跟踪
heap:查看活动对象的内存分配情况
mutex:查看导致互斥锁的竞争持有者的堆栈跟踪
threadcreate:查看创建新OS线程的堆栈跟踪

终端交互：
help查看所有命令说明
CPU：go tool pprof http://localhost:10000/debug/pprof/profile
heap: go tool pprof http://localhost:10000/debug/pprof/heap
block: go tool pprof http://localhost:10000/debug/pprof/block
mutex: go tool pprof http://localhost:10000/debug/pprof/mutex

flat：给定函数上运行耗时
flat%：同上的 CPU 运行耗时总比例
sum%：给定函数累积使用 CPU 总比例
cum：当前函数加上它之上的调用运行总耗时
cum%：同上的 CPU 运行耗时总比例

inuse_space:分析应用程序的常驻内存占用情况
alloc_objects：分析应用程序的内存临时分配情况


测试用例:
go test -h
go test -bench=. -cpuprofile=cpu.prof
go test -bench=. -memprofile=mem.prof
生成文件 cpu.prof mem.prof

启动可视化界面
go tool pprof -http=:10000 cpu.prof

http://localhost:10000/top
http://localhost:10000/peek
http://localhost:10000/source
