# learnToWork
试图让它运行起来的一个golang服务器框架.<br>
主要包括gate和gs两部分.其中gate可以作为网关,gs可以作为逻辑服务使用.gate和gs之间使用[grpc](https://github.com/grpc/grpc-go)通信.<br>
gate和gs均可以水平扩展,同时运行多个.

##### 初始版本,文件结构可能会变化.

快速启动
--------
##### 1.安装[mongodb](https://www.mongodb.com/download-center?jmp=nav#atlas)并启动.
##### 2.设置config.yaml中gate和gs的db为启动的mongodb ip和端口.
##### 3.sh startup.sh
