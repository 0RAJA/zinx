## 读写分离

就是与客户端进修数据交互的Goroutine由一个变成两个，一个专门负责从客户端读取数据，一个专门负责向客户端写数据

这么设计有什么好处，当然是目的就是高内聚，模块的功能单一，对于我们今后扩展功能更加方便。

Server依然是处理客户端的响应，主要关键的几个方法是Listen、Accept等。当建立与客户端的套接字后，那么就会开启两个Goroutine分别处理读数据业务和写数据业务，读写数据之间的消息通过一个Channel传递。

![img.png](D:\Projects\GoFile\zinx\img\读写分离.png)

## 消息队列和多任务机制

我们需要添加消息队列和多worker机制,通过worker的数量限定业务的固定goroutine数量.太多的goroutine会带来不必要的环境切换成本.可以使用消息队列来缓冲worker工作的数据

![img.png](D:\Projects\GoFile\zinx\img\工作池和消息队列.png)
