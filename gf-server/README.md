

## 1. 项目目录结构

```shell 
├── app
│   ├── api
│   │   ├── request
│   │   ├── response
│   │   └── v1
│   ├── model
│   │   └── user
│   └── service
│       ├── auth
│       └── user
├── boot
├── config
├── docker
├── document
├── global
├── i18n
├── library
│   └── response
├── logs
│   ├── log
│   ├── server
│   └── sql
├── middleware
├── packed
├── public
│   ├── html
│   ├── plugin
│   └── resource
│       ├── css
│       ├── image
│       ├── js
│       ├── page
│       │   ├── css
│       │   ├── img
│       │   └── js
│       └── template
│           ├── fe
│           └── te
├── router
│   ├── extend
│   └── system
└── template

├── Dockerfile
├── go.mod
└── main.go
```
| 目录/文件名称 | 说明       | 描述                                                         |
| :------------ | :--------- | :----------------------------------------------------------- |
| `app`         | 业务逻辑层 | 所有的业务逻辑存放目录。                                     |
| - `api`       | 业务接口   | 接收/解析用户输入参数的入口/接口层。                         |
| --`request`   | 入参结构体 | 接收前端发送到后端的数据。                                   |
| --`response`  | 出参结构体 | 返回给前端的数据结构体                                       |
| --`v1`        | api版本    | api的v1版本                                                  |
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

| 配置分类 |      配置名       |     配置类型      | 英文描述                                                     | 中文描述                                                     |                 官方文档链接                 |
| -------- | :---------------: | :---------------: | :----------------------------------------------------------- | :----------------------------------------------------------- | :------------------------------------------: |
|          |                   |                   |                                                              |                                                              |                                              |
| Basic    |      Address      |      string       | Address specifies the server listening address like "port" or ":port",multiple addresses joined using ','. | 地址指定服务器的侦听地址，例如“ port”或“：port”，使用'，'连接多个地址 |                                              |
| Basic    |     HTTPSAddr     |      string       | HTTPSAddr specifies the HTTPS addresses, multiple addresses joined using char ','. | HTTPSAddr指定HTTPS地址，即使用','连接的多个地址。            |                                              |
| Basic    |   HTTPSCertPath   |      string       | HTTPSCertPath specifies certification file path for HTTPS service. | HTTPSCertPath指定HTTPS服务的认证文件路径。                   |                                              |
| Basic    |   HTTPSKeyPath    |      string       | HTTPSKeyPath specifies the key file path for HTTPS service.  | HTTPSKeyPath指定HTTPS服务的密钥文件路径。                    |                                              |
| Basic    |    ServerAgent    |      string       | specifies the server agent information, which is wrote to Http response header as "Server" | ServerAgent指定服务器代理信息，该信息将写入，HTTP响应标头为“服务器” |                                              |
| Basic    |     TLSConfig     |    *tls.Config    | TLSConfig optionally provides a TLS configuration for use by ServeTLS and ListenAndServeTLS. Note that this value is cloned by ServeTLS and ListenAndServeTLS, so it's not  possible to modify the configuration with methods like tls.Config.SetSessionTicketKeys. To use  SetSessionTicketKeys, use Server.Serve with a TLS Listener instead. | （可选）提供由ServeTLS和ListenAndServeTLS使用的TLS配置。<br/>请注意，此值是由ServeTLS和ListenAndServeTLS克隆的，因此无法使用tls.Config.SetSessionTicketKeys之类的方法修改配置。<br/>若要使用SetSessionTicketKeys，请改为将Server.Serve与TLS侦听器一起使用。 |                                              |
| Basic    |      Handler      |   http.Handler    | Handler the handler for HTTP request.                        |                                                              |                                              |
| Basic    |    ReadTimeout    |   time.Duration   |                                                              |                                                              |                                              |
| Basic    |   WriteTimeout    |   time.Duration   |                                                              |                                                              |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
| Static   |     Rewrites      | map[string]string | Rewrites specifies the URI rewrite rules map.                | 重写指定URI重写规则映射。                                    |                                              |
| Static   |    IndexFiles     |     []string      | IndexFiles specifies the index files for static folder.      | 指定静态文件夹的索引文件。                                   | [点我](https://goframe.org/net/ghttp/static) |
| Static   |    IndexFolder    |       bool        | IndexFolder specifies if listing sub-files when requesting folder.The server responses HTTP status code 403 if it is false. | 指定在请求文件夹时是否列出子文件。服务器返回HTTP状态码403（如果为false）。 |                                              |
| Static   |    ServerRoot     |      string       | specifies the root directory for static service              | 指定静态服务的根目录,用来设置Server的主目录（类似nginx中的root配置，默认为空） |                                              |
| Static   |    SearchPaths    |     []string      | SearchPaths specifies additional searching directories for static service. | 指定用于静态服务的其他搜索目录                               |                                              |
| Static   |    StaticPaths    | []staticPathItem  | StaticPaths specifies URI to directory mapping array.        | 指定URI到目录的映射数组。                                    |                                              |
| Static   | FileServerEnabled |       bool        | FileServerEnabled is the global switch for static service.It is automatically set enabled if any static path is set. | 静态服务的全局开关。如果设置了任何静态路径，它将自动设置为启用。 |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
| Basic    |    IdleTimeout    |   time.Duration   |                                                              |                                                              |                                              |
| Basic    |  MaxHeaderBytes   |        int        |                                                              |                                                              |                                              |
| Basic    |     KeepAlive     |       bool        |                                                              |                                                              |                                              |
| Basic    |    ServerAgent    |      string       |                                                              |                                                              |                                              |
| Basic    |       View        |    *gview.View    |                                                              |                                                              |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
| Cookie   |   CookieMaxAge    |   time.Duration   | CookieMaxAge specifies the max TTL for cookie items.         | 指定Cookie项目的最大TTL。                                    |                                              |
| Cookie   |    CookiePath     |      string       | CookiePath specifies cookie path.It also affects the default storage for session id. | 指定cookie路径，它也会影响会话ID的默认存储。                 |                                              |
| Cookie   |   CookieDomain    |      string       | CookieDomain specifies cookie domain.It also affects the default storage for session id. | 指定cookie域，还会影响会话ID的默认存储。                     |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
| Session  |   SessionMaxAge   |   time.Duration   | specifies max TTL for session items.                         | 指定会话项目的最大TTL。                                      |                                              |
| Session  |   SessionIdName   |      string       | specifies the session id name.                               | 指定会话ID名称,SessionId的识别名称                           |                                              |
| Session  |    SessionPath    |      string       | specifies the session storage directory path for storing session files,It only makes sense if the session storage is type of file storage. | 指定用于存储会话文件的会话存储目录路径。仅当会话存储为文件存储类型时才有意义。 |                                              |
| Session  |  SessionStorage   | gsession.Storage  | SessionStorage specifies the session storage                 | SessionStorage指定会话存储。                                 |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
|          | AccessLogEnabled  |       bool        | enables access logging content to files.                     | 启用访问日志记录内容到文件                                   |  [点我](https://goframe.org/net/ghttp/logs)  |
|          |  ErrorLogEnabled  |       bool        | enables error logging content to files.                      | 启用错误日志记录文件                                         |  [点我](https://goframe.org/net/ghttp/logs)  |
|          |   PProfEnabled    |       bool        | enables PProf feature.                                       | 启用PProf功能                                                |  [点我](https://goframe.org/net/ghttp/logs)  |
|          |      LogPath      |      string       | specifies the directory for storing logging files.           | 指定用于存储日志文件的目录                                   |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
|          |   DumpRouterMap   |       bool        | specifies whether automatically dumps router map when server starts. | 指定服务器启动时是否自动转储路由器映射。                     |                                              |

### 2.2 [logger]

|    配置选项     | 配置类型 |                           英文描述                           |                           中文描述                           |
| :-------------: | :------: | :----------------------------------------------------------: | :----------------------------------------------------------: |
|     LogPath     |  string  |  LogPath specifies the directory for storing logging files.  |              LogPath指定用于存储日志文件的目录               |
|    LogStdout    |   bool   | LogStdout specifies whether printing logging content to stdout. |        LogStdout指定是否将日志记录内容打印到stdout。         |
|   ErrorStack    |   bool   | ErrorStack specifies whether logging stack information when error. |         ErrorStack指定是否在发生错误时记录堆栈信息。         |
|   ErrorStack    |   bool   |   ErrorLogEnabled enables error logging content to files.    |         ErrorLogEnabled允许将内容错误记录到文件中。          |
| ErrorLogPattern |  string  | AccessLogPattern specifies the error log file pattern like: access-{Ymd}.log | AccessLogPattern指定错误日志文件模式，例如：access- {Ymd} .log |

### 2.3 [database]

```shell
[database]
    [[database.分组名称]]
        Host         = "地址"
        Port         = "端口"
        User         = "账号"
        Pass         = "密码"
        Name         = "数据库名称"
        Type         = "数据库类型(目前支持mysql/pgsql/sqlite)"
        Role         = "(可选)数据库主从角色(master/slave)，不使用应用层的主从机制请均设置为master"
        Debug        = "(可选)开启调试模式"
        Prefix       = "(可选)表名前缀"
        DryRun       = "(可选)ORM空跑(只读不写)"
        Charset      = "(可选)数据库编码(如: utf8/gbk/gb2312)，一般设置为utf8"
        Weight       = "(可选)负载均衡权重，用于负载均衡控制，不使用应用层的负载均衡机制请置空"
        Linkinfo     = "(可选)自定义数据库链接信息，当该字段被设置值时，以上链接字段(Host,Port,User,Pass,Name)将失效，但是type必须有值"
        MaxIdle      = "(可选)连接池最大闲置的连接数"
        MaxOpen      = "(可选)连接池最大打开的连接数"
        MaxLifetime  = "(可选，单位秒)连接对象可重复使用的时间长度"
```



### 2.4 [redis]



| 配置项名称      | 是否必须 | 默认值 | 说明                                   |
| :-------------- | -------- | ------ | -------------------------------------- |
| host            | 是       | -      | 地址                                   |
| port            | 是       | -      | 端口                                   |
| db              | 否       | 0      | 数据库                                 |
| pass            | 否       | -      | 授权密码                               |
| maxIdle         | 否       | 0      | 允许限制的连接数(0表示不闲置)          |
| maxActive       | 否       | 0      | 最大连接数量限制(0表示不限制)          |
| idleTimeout     | 否       | 60     | 连接最大空闲时间(单位秒,不允许设置为0) |
| maxConnLifetime | 否       | 60     | 连接最长存活时间(单位秒,不允许设置为0) |

