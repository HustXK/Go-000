## 学习笔记

1. keep yourself busy or do the work yourself；
2. select{}可以阻塞goroutine；
3. 如果goroutine不知道什么时候结束，则不应该调用，调用者需要感知或者控制goroutine的生命周期；
4. log.Fatal 调用了os.Exit，无法执行defer;
5. 把并行的逻辑交给调用者；
6. filepath.Walk()参考
7. 作业：https://golang.org/ref/mem
8. 先行发生，要定义读写的顺序，先写后读，写之后读之前不能有写的操作；
9. 多个goroutine访问共享变量事，必须要使用同步时间来建立先行发生这一条件来保证读操作能看到的写操作;
10. nginx内存屏障 memory barriar
11. intel 的缓存一致性

