gf-vue-admin

## 开发计划

- [x] gf-vue-admin服务端
  - [x] 基于[gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)的server与[goframe](https://goframe.org/start/index)的官方推荐架构
  - [x] 架构具体看文档
- [x] 路由层
- [x] jwt鉴权
- [x] jwt黑名单
- [x] 管理员用户模块
- [x] 基础功模块
- [ ] menu模块
- [ ] 角色管理模块
- [x] 功能api模块
- [x] 文件上传下载功能模块
- [ ] 文件断点续传功能模块
- [ ] 工作流
- [ ] casbin权限模块
- [ ] ~~system模块~~
- [ ] 客户模块
- [ ] 自动化代码模块
- [x] 字典详情管理模块
- [x] 字典管理模块
- [x] 操作历史模块

## 测试api
- [x] 基础功模块
    - [x] 管理员登录
    - [x] 管理员注册
    - [x] 验证码获取
    - [ ] JWT刷新(功能已经写好,未测试)
- [x] 管理员用户模块    
    - [x] 管理员修改密码    
    - [x] 分页获取管理员列表    
    - [x] 设置用户权限    
    - [ ] 删除用户(ID接收问题)
  
## TODO
- [ ] 用户多角色
- [ ] 自写api批量导入权限分配列表