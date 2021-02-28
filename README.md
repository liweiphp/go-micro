# go-micro
golang微服务


docker golang环境 和 consul环境 和 consul集群

创建网络：
docker network create --subnet=192.168.5.0/24 mynetwork

启动顺序：

1.centos-go镜像构建 centos-go目录下 docker build -t centos-go .

2.启动go  server go目录下 docker-compose up -d 地址映射修改成自己的 默认是：/www/study/golang/src/d5/tcp/client

3.构建consul镜像(alpine系统) consul目录下 docker build -t consul . 

4.启动consul集群 consul目录下 docker-compose up -d

tcp启动命令：docker exec -it go-tcp-server go run tcp/server

请求命令：docker exec -it go-tcp-client go run tcp/client
