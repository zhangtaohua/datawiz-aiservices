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

# 二、 部署相关
`go build -ldflags "-H=windowsgui" -o goexcel.exe`

1、完全杀死 nginx 命令 
`taskkill /f /t /im nginx.exe`

2、完全杀死 goexcel.exe 命令
`taskkill /f /t /im goexcel.exe`

`CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o goexcel.exe`


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

