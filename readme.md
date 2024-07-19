# 一、 docker postgreSql 
docker run --name rjpostgres -e POSTGRES_PASSWORD=123456  -e ALLOW_IP_RANGE=0.0.0.0/0 -e POSTGRES_HOST_AUTH_METHOD=md5 -d -p 5432:5432 postgres

docker run --name rjpostgres -e POSTGRES_PASSWORD=123456  -e ALLOW_IP_RANGE=0.0.0.0/0 -e POSTGRES_HOST_AUTH_METHOD=md5 -v d:/docker_data/pg:/var/lib/postgresql/data -d -p 5432:5432 postgres



1、 安装环境
    包括数据库 mysql redis 等。

    drop database aidb;
    create database aidb;

2、 修改 env 文件
    主要是数据库链接 地址 端口 密码等。

3、 执行命令导入表结构
    `go run main.go migrate up`

4、 执行命令 seed 预定义数据
    `go run main.go seed`

5、 导入数据
    ``

# 二、 windowns部署相关
`go build -ldflags "-H=windowsgui" -o goexcel.exe`

1、完全杀死 nginx 命令 
`taskkill /f /t /im nginx.exe`

2、完全杀死 goexcel.exe 命令
`taskkill /f /t /im goexcel.exe`

`CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o goexcel.exe`

go build -ldflags "-H=windowsgui" -o goexcel.exe
docker run -d --name RJmysql -p 3306:3306 -v /d/dataserver/data/mysql/data:/var/lib/mysql/ -v /d/dataserver/data/mysql/conf/my.cnf:/etc/mysql/my.cnf  -v /d/dataserver/data/mysql/logs:/logs -e MYSQL_ROOT_PASSWORD=123456 mysql

docker run --name mysql -v D:/docker/mysql/conf/my.cnf:/etc/mysql/my.cnf -v D:/docker/mysql/logs:/logs -v D:/docker/mysql/data:/var/lib/mysql -v  D:/docker/mysql/conf/conf.d:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=123456 -d -i -p 3306:3306 mysql:5.7

nginx -t

docker run -d --name RJredis -p 6379:6379 -v /d/dataserver/data/redis/data:/data redis --appendonly yes

 kill -9 $(lsof -ti:3000)
 
重新启动

1、双击 nginx.exe 看到一个小黑窗一闪而过，启动成功

2、窗口中 输入cmd 回车 
再弹出的命令行窗口中输入 : 
`goexcel.exe -d true`
 后回车



解决错误: 函数 uuid_generate_v4() 不存在
CREATE EXTENSION pgcrypto;
create extension "uuid-ossp"

语言参考表
http://www.lingoes.net/zh/translator/langcode.htm

# 4 Linux 部署

## 4.1 前置条件：

### 4.1.1 硬件条件
CPU：
GPU：
内存：
硬盘：


### 4.1.2 软件条件

操作系统： linux(ubuntu, centos)
依赖库： cifs-utils、cuda, cudnn
软件： docker、

## 4.2 实际步骤：

以下所有步骤均认为已满足前置条件.

### 4.2.1 复制文件
假设

源文件目录： D:\share\go
目标文件目录：/home/datawiz/datawiz-ai

使用任意办法将源文件目录下所有文件复制到目标文件目录下。
完成后的目录结果如下所示
```
├── build
│   └── i18n
│       └── en.json
│       └── zh-CN.json
│       └── zh-TW.json
├── build
│   └── images
│       └── starwiz_ai_go.tar
├── deployments
│   ├── deploy.sh
│   ├── docker-compose
│   │   └── docker-compose.yaml
├── .env
```

### 4.3.2 执行脚本
chmod -R 777 .
cd deplyments
chmod +x deploy.sh
./deploy.sh

