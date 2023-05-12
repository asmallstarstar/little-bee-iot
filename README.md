# little-bee开源物联网平台

#### 介绍
little-bee是一个部署简单，运维容易的物联网平台，后端采用golang开发，编译后即可以以绿色软件的形式运行。在工作中，作者发现很多项目需要数据采集功能，不同的是前端的展示不一样，同时网上这方面的开源系统都很庞大，不利于中小型集成商的运维，因此开发了这个部署运维简单的物联网平台。虽然简单，但是同样具有可伸缩，容错，健壮和高效。

#### 软件架构
1. 技术栈

   后端：golang gin grpc gorm mqtt casbin

   前端：vue3 element plus pinia vue-router vite

   配置库：mysql

   被监控设备驱动: c
   
2. 系统架构图

   ![系统架构图](https://gitee.com/asmallstar/little-bee/raw/master/little-bee/agent/agent/architecture/architecture.png)

#### 项目特点

1. 访问控制RBAC，基于API的细粒度访问权限控制
2. 网关集成，智能设备驱动的参数使用元数据形式，易于扩充
3. 绿色软件的形式部署，易于使用

#### 编译教程（windows）

1. 后端安装好golang开发环境，配置好protobuf环境变量

2. 在项目当前目录下运行下面脚本编译后端可执行程序:

   .\message\compile.bat

   .\bin\service\build.bat

   .\bin\agent\build-realdata.bat

   .\bin\agent\build-fsuserver.bat

   .\bin\agent\build-driverloader.bat

3. 前端

   npm install

   npm run dev 

#### 使用说明

1. 安装mysql 

2. 运行dao\resources目录下数据库脚本文件

   little_bee_mysql.ddl

   basic_data_mysql_zh-CN.ddl

   menu_actions.ddl 

3. 安装mqtt server，推荐mosquitto

4. 启动编译好的可执行程序

   在bin\service目录下运行service.exe

   在bin\agent目录下运行realdata.exe fsuserver.exe driverloader.exe(这三个进程由agent.exe管理，agent没有开发完成，目前直接启动)
   
#### 前端界面截图

![登录](https://gitee.com/asmallstat/little-bee/raw/master/frontend/console/public/screenshot/login.png)

![实时数据](https://gitee.com/asmallstat/little-bee/raw/master/frontend/console/public/screenshot/realdata.png)

![实时告警](https://gitee.com/asmallstat/little-bee/raw/master/frontend/console/public/screenshot/alarm.png)

![远程控制](https://gitee.com/asmallstat/little-bee/raw/master/frontend/console/public/screenshot/control.png)

![监控量配置](https://gitee.com/asmallstat/little-bee/raw/master/frontend/console/public/screenshot/config.png)

![系统设置](https://gitee.com/asmallstat/little-bee/raw/master/frontend/console/public/screenshot/setting.png)

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_feature 分支
3.  提交代码
4.  新建 Pull Request
4.  欢迎加入qq群：539046464


#### 计划

1.  本项目是一个物联框架，后期将围绕这一平台增加门禁系统，视频监控系统，信息发布（大屏）系统
