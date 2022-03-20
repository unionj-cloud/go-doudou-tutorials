# nacos

[EN](./README.md) 中文

本项目是go-doudou写的RESTful服务加入Nacos集群，与spring cloud写的RESTful服务相互调用的示例代码。

## 准备工作

- 需安装和配置好docker开发环境
- 现代浏览器
- 趁手的命令行工具
- 趁手的IDE，推荐goland
- 趁手的http请求工具，推荐[httpie](https://github.com/httpie/httpie)

## 启动步骤

1. 启动standalone模式的nacos server
```shell
docker-compose -f docker-compose.yml up -d
```

```shell 
➜  nacos git:(master) ✗ docker ps                                                  
CONTAINER ID   IMAGE                       COMMAND                  CREATED          STATUS                    PORTS                                                                                  NAMES
2e5400f6d085   nacos/nacos-server:v2.0.4   "bin/docker-startup.…"   18 minutes ago   Up 18 minutes (healthy)   0.0.0.0:8848->8848/tcp, :::8848->8848/tcp, 0.0.0.0:9848->9848/tcp, :::9848->9848/tcp   nacos-standalone
```

2. 切换到statsvc路径下，执行服务启动命令
```shell
go mod tidy && go run cmd/main.go
```

看到如下控制台输出，表示服务启动成功
```shell               
 _____                     _                    _
|  __ \                   | |                  | |
| |  \/  ___   ______   __| |  ___   _   _   __| |  ___   _   _
| | __  / _ \ |______| / _` | / _ \ | | | | / _` | / _ \ | | | |
| |_\ \| (_) |        | (_| || (_) || |_| || (_| || (_) || |_| |
 \____/ \___/          \__,_| \___/  \__,_| \__,_| \___/  \__,_|
INFO[2022-03-12 13:24:31] ================ Registered Routes ================ 
INFO[2022-03-12 13:24:31] +-------------+--------+-------------------------+ 
INFO[2022-03-12 13:24:31] |    NAME     | METHOD |         PATTERN         | 
INFO[2022-03-12 13:24:31] +-------------+--------+-------------------------+ 
INFO[2022-03-12 13:24:31] | Add         | POST   | /api/add                | 
INFO[2022-03-12 13:24:31] | Echo        | GET    | /api/echo               | 
INFO[2022-03-12 13:24:31] | GetDoc      | GET    | /go-doudou/doc          | 
INFO[2022-03-12 13:24:31] | GetOpenAPI  | GET    | /go-doudou/openapi.json | 
INFO[2022-03-12 13:24:31] | Prometheus  | GET    | /go-doudou/prometheus   | 
INFO[2022-03-12 13:24:31] | GetRegistry | GET    | /go-doudou/registry     | 
INFO[2022-03-12 13:24:31] | GetConfig   | GET    | /go-doudou/config       | 
INFO[2022-03-12 13:24:31] +-------------+--------+-------------------------+ 
INFO[2022-03-12 13:24:31] =================================================== 
INFO[2022-03-12 13:24:31] Started in 564.738µs                         
INFO[2022-03-12 13:24:31] Http server is listening on :6060 
```

此时，打开浏览器，导航到http://localhost:8848/nacos， 登录后点击/服务管理/服务列表，可看到statsvc服务已经注册到nacos

![nacosscreenshot.png](./nacosscreenshot.png)

3. 切换到nacos-service-j路径下，执行如下命令
```shell
mvn spring-boot:run
```

看到控制台有如下输出，即表示启动成功
```shell
2022-03-12 13:32:41.630  INFO 88684 --- [t.remote.worker] com.alibaba.nacos.common.remote.client   : [f1bf3cba-e469-43c7-a8b8-83457ee4c153] Notify connected event to listeners.
2022-03-12 13:32:41.706  INFO 88684 --- [           main] o.s.b.w.embedded.tomcat.TomcatWebServer  : Tomcat started on port(s): 8080 (http) with context path '/nacos-service-j'
2022-03-12 13:32:41.723  INFO 88684 --- [           main] c.a.c.n.registry.NacosServiceRegistry    : nacos registry, DEFAULT_GROUP nacos-service-j 192.168.43.200:8080 register finished
2022-03-12 13:32:41.727  INFO 88684 --- [           main] c.u.n.NacosServiceJApplication           : Started NacosServiceJApplication in 2.751 seconds (JVM running for 3.059)
```

此时，刷新/服务管理/服务列表页面，可看到nacos-service-j服务已经注册到nacos

![nacosscreenshot2.png](./nacosscreenshot2.png)

4. 测试statsvc服务
```shell
➜  nacos git:(master) ✗ http http://localhost:6060/api/echo\?s\=go-doudou
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 67
Content-Type: application/json; charset=UTF-8
Date: Sat, 12 Mar 2022 05:37:43 GMT
Vary: Accept-Encoding

{
    "data": "Hello Nacos Discovery go-doudou"
}


```

5. 测试nacos-service-j服务
```shell 
➜  nacos git:(master) ✗ http http://localhost:8080/nacos-service-j/add\?x\=50\&y\=20                                            
HTTP/1.1 200 
Connection: keep-alive
Content-Type: application/json
Date: Sat, 12 Mar 2022 05:38:38 GMT
Keep-Alive: timeout=60
Transfer-Encoding: chunked

{
    "data": 70
}


```