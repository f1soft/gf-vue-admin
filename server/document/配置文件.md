
## 2. 配置文件详解(config.toml)

### 2.1 [server]

| 配置分类 |      配置名       |     配置类型      | 英文描述                                                     | 中文描述                                                     |                 官方文档链接                 |
| -------- | :---------------: | :---------------: | :----------------------------------------------------------- | :----------------------------------------------------------- | :------------------------------------------: |
|          |                   |                   |                                                              |                                                              |                                              |
| Basic    |      Address      |      string       | Address specifies the server listening address like "port" or ":port",multiple addresses joined using ','. | 地址指定服务器的侦听地址，例如“ port”或“：port”，使用'，'连接多个地址 |                                              |
| Basic    |     HTTPSAddr     |      string       | HTTPSAddr specifies the HTTPS addresses, multiple addresses joined using char ','. | HTTPSAddr指定HTTPS地址，即使用','连接的多个地址              |                                              |
| Basic    |   HTTPSCertPath   |      string       | HTTPSCertPath specifies certification file path for HTTPS service. | HTTPSCertPath指定HTTPS服务的认证文件路径                     |                                              |
| Basic    |   HTTPSKeyPath    |      string       | HTTPSKeyPath specifies the key file path for HTTPS service.  | HTTPSKeyPath指定HTTPS服务的密钥文件路径                      |                                              |
| Basic    |    ServerAgent    |      string       | specifies the server agent information, which is wrote to Http response header as "Server" | ServerAgent指定服务器代理信息，该信息将写入，HTTP响应标头为“服务器” |                                              |
| Basic    |     TLSConfig     |    *tls.Config    | TLSConfig optionally provides a TLS configuration for use by ServeTLS and ListenAndServeTLS. Note that this value is cloned by ServeTLS and ListenAndServeTLS, so it's not  possible to modify the configuration with methods like tls.Config.SetSessionTicketKeys. To use  SetSessionTicketKeys, use Server.Serve with a TLS Listener instead. | （可选）提供由ServeTLS和ListenAndServeTLS使用的TLS配置请注意，此值是由ServeTLS和ListenAndServeTLS克隆的，因此无法使用tls.Config.SetSessionTicketKeys之类的方法修改配置。若要使用SetSessionTicketKeys，请改为将Server.Serve与TLS侦听器一起使用 |                                              |
| Basic    |      Handler      |   http.Handler    | Handler the handler for HTTP request.                        | 处理程序HTTP请求的处理程序                                   |                                              |
| Basic    |    ReadTimeout    |   time.Duration   | ReadTimeout is the maximum duration for reading the entire  request, including the body. Because ReadTimeout does not let Handlers make per-request  decisions on each request body's acceptable deadline or  upload rate, most users will prefer to use  ReadHeaderTimeout. It is valid to use them both. | ReadTimeout是读取整个请求（包括正文）的最大持续时间。由于ReadTimeout不允许处理程序根据每个请求正文的可接受的截止日期或上载速率做出每个请求的决定，因此大多数用户将更喜欢使用ReadHeaderTimeout。同时使用它们是有效的 |                                              |
| Basic    |   WriteTimeout    |   time.Duration   | WriteTimeout is the maximum duration before timing out writes of the response. It is reset whenever a new request's header is read. Like ReadTimeout, it does not let Handlers make decisions on a per-request basis. | WriteTimeout是超时写入响应之前的最大持续时间。每当读取新请求的标头时，都会将其重置。与ReadTimeout一样，它也不允许处理程序根据每个请求做出决策 |                                              |
| Basic    |    IdleTimeout    |   time.Duration   | IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled. If IdleTimeout  is zero, the value of ReadTimeout is used. If both are zero, there is no timeout. | IdleTimeout是启用保持活动状态后等待下一个请求的最长时间。如果IdleTimeout为零，则使用ReadTimeout的值。如果两者均为零，则没有超时 |                                              |
| Basic    |  MaxHeaderBytes   |        int        | MaxHeaderBytes controls the maximum number of bytes the server will read parsing the request header's keys and	 values, including the request line. It does not limit the size of the request body. It can be configured in configuration file using string like: 1m, 10m, 500kb etc. It's 1024 bytes in default. | MaxHeaderBytes控制服务器读取的最大字节数，以解析请求标头的键和值（包括请求行）。它不限制请求主体的大小。可以使用以下字符串在配置文件中对其进行配置：1m，10m，500kb等。默认值为1024字节 |                                              |
| Basic    |    ServerAgent    |      string       | ServerAgent specifies the server agent information, which is wrote to HTTP response header as "Server" | ServerAgent指定服务器代理信息，该信息作为“服务器”写入HTTP响应标头 |                                              |
| Basic    |       View        |    *gview.View    | View specifies the default template view object for the server. | 视图指定服务器的默认模板视图对象                             |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
| Static   |     Rewrites      | map[string]string | Rewrites specifies the URI rewrite rules map.                | 重写指定URI重写规则映射                                      |                                              |
| Static   |    IndexFiles     |     []string      | IndexFiles specifies the index files for static folder.      | 指定静态文件夹的索引文件                                     | [点我](https://goframe.org/net/ghttp/static) |
| Static   |    IndexFolder    |       bool        | IndexFolder specifies if listing sub-files when requesting folder.The server responses HTTP status code 403 if it is false. | 指定在请求文件夹时是否列出子文件。服务器返回HTTP状态码403（如果为false） |                                              |
| Static   |    ServerRoot     |      string       | specifies the root directory for static service              | 指定静态服务的根目录,用来设置Server的主目录（类似nginx中的root配置，默认为空） |                                              |
| Static   |    SearchPaths    |     []string      | SearchPaths specifies additional searching directories for static service. | 指定用于静态服务的其他搜索目录                               |                                              |
| Static   |    StaticPaths    | []staticPathItem  | StaticPaths specifies URI to directory mapping array.        | 指定URI到目录的映射数组                                      |                                              |
| Static   | FileServerEnabled |       bool        | FileServerEnabled is the global switch for static service.It is automatically set enabled if any static path is set. | 静态服务的全局开关。如果设置了任何静态路径，它将自动设置为启用 |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
| Cookie   |   CookieMaxAge    |   time.Duration   | CookieMaxAge specifies the max TTL for cookie items.         | 指定Cookie项目的最大TTL                                      |                                              |
| Cookie   |    CookiePath     |      string       | CookiePath specifies cookie path.It also affects the default storage for session id. | 指定cookie路径，它也会影响会话ID的默认存储                   |                                              |
| Cookie   |   CookieDomain    |      string       | CookieDomain specifies cookie domain.It also affects the default storage for session id. | 指定cookie域，还会影响会话ID的默认存储                       |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
| Session  |   SessionMaxAge   |   time.Duration   | specifies max TTL for session items.                         | 指定会话项目的最大TTL                                        |                                              |
| Session  |   SessionIdName   |      string       | specifies the session id name.                               | 指定会话ID名称,SessionId的识别名称                           |                                              |
| Session  |    SessionPath    |      string       | specifies the session storage directory path for storing session files,It only makes sense if the session storage is type of file storage. | 指定用于存储会话文件的会话存储目录路径。仅当会话存储为文件存储类型时才有意义 |                                              |
| Session  |  SessionStorage   | gsession.Storage  | SessionStorage specifies the session storage                 | SessionStorage指定会话存储                                   |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
| Logging  |      Logger       |   *glog.Logger    | Logger specifies the logger for server.                      | 记录器指定服务器的记录器                                     |                                              |
| Logging  |      LogPath      |      string       | LogPath specifies the directory for storing logging files.   | 指定用于存储日志文件的目录                                   |                                              |
| Logging  |     LogStdout     |       bool        | LogStdout specifies whether printing logging content to stdout. | LogStdout指定是否将日志记录内容打印到stdout                  |                                              |
| Logging  |    ErrorStack     |       bool        | ErrorStack specifies whether logging stack information when error. | ErrorStack指定是否在发生错误时记录堆栈信息                   |                                              |
| Logging  |  ErrorLogEnabled  |       bool        | enables error logging content to files.                      | 启用错误日志记录文件                                         |  [点我](https://goframe.org/net/ghttp/logs)  |
| Logging  | AccessLogEnabled  |       bool        | enables access logging content to files.                     | 启用访问日志记录内容到文件                                   |  [点我](https://goframe.org/net/ghttp/logs)  |
| Logging  | AccessLogPattern  |      string       | AccessLogPattern specifies the error log file pattern like: access-{Ymd}.log | AccessLogPattern指定错误日志文件模式，例如：access- {Ymd} .log |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
| PProf    |   PProfEnabled    |       bool        | enables PProf feature.                                       | 启用PProf功能                                                |  [点我](https://goframe.org/net/ghttp/logs)  |
| PProf    |   PProfPattern    |      string       | PProfPattern specifies the PProf service pattern for router. | 指定路由器的PProf服务模式                                    |                                              |
|          |                   |                   |                                                              |                                                              |                                              |
| Other    | ClientMaxBodySize |       int64       | ClientMaxBodySize specifies the max body size limit in bytes for client request.  It can be configured in configuration file using string like: 1m, 10m, 500kb etc.  It's 8MB in default. | ClientMaxBodySize指定客户端请求的最大主体大小限制（以字节为单位）。可以使用以下字符串在配置文件中进行配置：1m，10m，500kb等。默认为8MB。 |                                              |
| Other    | FormParsingMemory |       int64       | FormParsingMemory specifies max memory buffer size in bytes which can be used for parsing multimedia form. It can be configured in configuration file using string like: 1m, 10m, 500kb etc. It's 1MB in default. | FormParsingMemory以字节为单位指定最大内存缓冲区大小，可用于解析多媒体格式。可以使用以下字符串在配置文件中对其进行配置：1m，10m，500kb等。默认为1MB。 |                                              |
| Other    |   NameToUriType   |        int        | NameToUriType specifies the type for converting struct method name to URI when registering routes. | NameToUriType指定在注册路由时用于将结构方法名称转换为URI的类型 |                                              |
| Other    |  RouteOverWrite   |       bool        | RouteOverWrite allows overwrite the route if duplicated.     | 如果复制，RouteOverWrite允许覆盖路由                         |                                              |
| Other    |   DumpRouterMap   |       bool        | specifies whether automatically dumps router map when server starts. | 指定服务器启动时是否自动转储路由器映射。                     |                                              |
| Other    |     Graceful      |       bool        | Graceful enables graceful reload feature for all servers of the process. | 为进程的所有服务器启用优美的重新加载功能                     |                                              |

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

