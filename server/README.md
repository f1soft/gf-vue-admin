## 注意问题
- Mysql的tinyint(1) 使用gf gen model 生成model文件夹里面的Entity.go文件中的字段类型是int而不是bool
- 开发环境下连接MySQL数据库尽量避免使用远程数据库,会导致运行后端服务缓慢,使用本地MySQL开发,然后使用远程数据库部署

## 需要讲解的问题
- request包与response包为什么不用model_entity的Entity进行嵌套
    - 不想破坏model_entity的完整性,model包是gf gen model进行生成
    - 需要加p与v的Tag
    
- Jwt验证方式
    - 格式"header:[Authorization], query:[token], cookie:[jwt]"
    - 注意冒号后面没有空格
    - header参数设置了就必须带有TokenHeadName,不然无法识别
    - 