+ 实现的功能
    + [x] 根据entity生成实体属性常量
    + [x] 根据entity生成model
    + [x] 根据entity生成queryType
    + [ ] 根据entity生成rootQueryType
    + [ ] 根据entity生成数据库建表语句及docker配置
    + [ ] 根据entity自动生成对应的微服务项目及docker配置
  
+ 目录介绍
    + 公共代码: 各个模块都会使用跟具体业务无关的代码
    + entity: 模型配置目录，存放各个模型的json文件
    + gen: 生成公共代码的脚本
    + graphql: gql查询项目，所有gql查询都使用该项目统一处理，从这里发送请求到具体的业务模块
    + gateway: 网关
    + devenv: docker-compose管理模块