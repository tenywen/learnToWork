# learnToWork
试图让它运行起来的一个golang 服务器框架。主要包括gate和gs两部分。
gate可以作为网关。
gs可以作为逻辑服务使用。
gate和gs使用grpc-go链接。gate和gs均可以水平扩展，同时运行多个。
初始版本，文件结构可能变化较大。

# start
1.
install [govendor](https://github.com/kardianos/govendor)
go get -u github.com/kardianos/govendor
install dependencies

2.
install [mongodb](https://www.mongodb.com/download-center?jmp=nav#atlas)
start mongodb

3.
sh startup.sh
