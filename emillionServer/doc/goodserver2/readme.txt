客户端使用
net.DialTCP("tcp", localaddr, tcpAddr)
函数可以采用本地不同的ip地址和服务器通信，每个ip地址最多６５５３５个端口

所以在本地可以建立多个虚拟的ＩＰ地址，可建立百万级的客户端链接，做测试
