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
