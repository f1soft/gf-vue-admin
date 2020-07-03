## 项目目录结构

```shell 
├── app
│   ├── api
│   ├── model
│   └── service
├── boot
├── config
├── docker
├── document
├── i18n
├── library
├── packed
├── public
├── router
├── template
├── Dockerfile
├── go.mod
└── main.go
```
| 目录/文件名称 | 说明       | 描述                                                         |
| :------------ | :--------- | :----------------------------------------------------------- |
| `app`         | 业务逻辑层 | 所有的业务逻辑存放目录。                                     |
| - `api`       | 业务接口   | 接收/解析用户输入参数的入口/接口层。                         |
| - `model`     | 数据模型   | 数据管理层，仅用于操作管理数据，如数据库操作。               |
| - `service`   | 逻辑封装   | 业务逻辑封装层，实现特定的业务需求，可供不同的包调用。       |
| `boot`        | 初始化包   | 用于项目初始化参数设置，往往作为`main.go`中第一个被`import`的包。 |
| `config`      | 配置管理   | 所有的配置文件存放目录。                                     |
| `docker`      | 镜像文件   | Docker镜像相关依赖文件，脚本文件等等。                       |
| `document`    | 项目文档   | Documentation项目文档，如: 设计文档、帮助文档等等。          |
| `i18n`        | I18N国际化 | I18N国际化配置文件目录。                                     |
| `library`     | 公共库包   | 公共的功能封装包，往往不包含业务需求实现。                   |
| `packed`      | 打包目录   | 将资源文件打包的`Go`文件存放在这里，`boot`包初始化时会自动调用。 |
| `public`      | 静态目录   | 仅有该目录下的文件才能对外提供静态服务访问。                 |
| `router`      | 路由注册   | 用于路由统一的注册管理。                                     |
| `template`    | 模板文件   | MVC模板文件存放的目录。                                      |
| `Dockerfile`  | 镜像描述   | 云原生时代用于编译生成Docker镜像的描述文件。                 |
| `go.mod`      | 依赖管理   | 使用`Go Module`包管理的依赖描述文件。                        |
| `main.go`     | 入口文件   | 程序入口文件。                                               |


## 2. 配置文件详解(config.toml)

### 2.1 [server]

|                  |   配置类型    |                           英文描述                           |                           中文描述                           |                 官方文档链接                 |
| :--------------: | :-----------: | :----------------------------------------------------------: | :----------------------------------------------------------: | :------------------------------------------: |
|     Address      |    string     | Address specifies the server listening address like "port" or ":port",multiple addresses joined using ','. | //地址指定服务器的侦听地址，例如“ port”或“：port”，使用'，'连接多个地址 |                                              |
|    ServerRoot    |    string     |       specifies the root directory for static service        | 指定静态服务的根目录,用来设置Server的主目录（类似nginx中的root配置，默认为空） |                                              |
|   ServerAgent    |    string     | specifies the server agent information, which is wrote to Http response header as "Server" | ServerAgent指定服务器代理信息，该信息将写入，HTTP响应标头为“服务器” |                                              |
|    IndexFiles    |   []string    |         specifies the index files for static folder          |                   指定静态文件夹的索引文件                   | [点我](https://goframe.org/net/ghttp/static) |
| AccessLogEnabled |     bool      |           enables access logging content to files.           |                  启用访问日志记录内容到文件                  |  [点我](https://goframe.org/net/ghttp/logs)  |
| ErrorLogEnabled  |     bool      |           enables error logging content to files.            |                     启用错误日志记录文件                     |  [点我](https://goframe.org/net/ghttp/logs)  |
|   PProfEnabled   |     bool      |                    enables PProf feature.                    |                        启用PProf功能                         |  [点我](https://goframe.org/net/ghttp/logs)  |
|     LogPath      |    string     |      specifies the directory for storing logging files.      |                  指定用于存储日志文件的目录                  |                                              |
|  SessionIdName   |    string     |                specifies the session id name.                |              指定会话ID名称,SessionId的识别名称              |                                              |
|   SessionPath    |    string     | specifies the session storage directory path for storing session files,It only makes sense if the session storage is type of file storage. | 指定用于存储会话文件的会话存储目录路径。仅当会话存储为文件存储类型时才有意义。 |                                              |
|  SessionMaxAge   | time.Duration |             specifies max TTL for session items.             |                   指定会话项目的最大TTL。                    |                                              |
|  DumpRouterMap   |     bool      | specifies whether automatically dumps router map when server starts. |           指定服务器启动时是否自动转储路由器映射。           |                                              |

### 2.2 [logger]

|    配置选项     | 配置类型 |                           英文描述                           |                           中文描述                           |
| :-------------: | :------: | :----------------------------------------------------------: | :----------------------------------------------------------: |
|     LogPath     |  string  |  LogPath specifies the directory for storing logging files.  |              LogPath指定用于存储日志文件的目录               |
|    LogStdout    |   bool   | LogStdout specifies whether printing logging content to stdout. |        LogStdout指定是否将日志记录内容打印到stdout。         |
|   ErrorStack    |   bool   | ErrorStack specifies whether logging stack information when error. | ErrorStack specifies whether logging stack information when error. |
|   ErrorStack    |   bool   |   ErrorLogEnabled enables error logging content to files.    |         ErrorLogEnabled允许将内容错误记录到文件中。          |
| ErrorLogPattern |  string  | AccessLogPattern specifies the error log file pattern like: access-{Ymd}.log | AccessLogPattern指定错误日志文件模式，例如：access- {Ymd} .log |
