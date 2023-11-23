# EdgeSys

EdgeSys企业级边缘平台快速开发框架

基于Go 1.19前后端分离架构，前端EdgeSysUi使用 Vue3.0 + TypeScript + vite3 + Element-plus技术


## 🌈框架简介

* 对前后端进行了大部分功能的封装，后端自封装go-restful，使用起来更加简洁，功能逻辑清晰，能快速上手学习，并用在生产中。
* 前端采用VUE3.0+ TypeScript + vite3 + Element-plus：EdgeSysUi，适配手机、平板、pc 内置多种ui功能减少开发量
* 高效率的开发，使用代码生成器可以一键生成前后端代码，可在线预览代码，减少代码开发量。。
* 完善的权限认证系统：完善的权限认证系统，包含，菜单按钮权限，api权限，组织权限。
* 多数据库：项目同时支持MySQL，PostgreSql等数据库根据自身需求更改。


## ⚡ 内置功能

- **`用户管理`** - _用户是系统操作者，该功能主要完成系统用户配置。._
- **`组织管理`** - _配置系统组织机构（公司、组织、小组），树结构展现支持数据权限。_
- **`岗位管理`** - _配置系统用户所属担任职务。_
- **`菜单管理`** - _配置系统菜单，操作权限，按钮权限标识等。_
- **`角色管理`** - _角色菜单,API权限分配、设置角色按机构进行数据范围权限划分。_
- **`字典管理`** - _对系统中经常使用的一些较为固定的数据进行维护。_
- **`参数管理`** - _对系统动态配置常用参数。_
-  **`通知公告`** - _系统通知公告信息发布维护_
- **`日志系统`** - _记录日志，更直观浏览_
- **`系统接口`** - _根据业务代码自动生成相关的api接口文档。_
- **`服务监控`** - _监视当前系统CPU、内存、磁盘、堆栈等相关信息。_
- **`代码生成`** - _可直接通过框架生成前后端基础业务代码（go、vue），减少开发时间。_



## 后端工程结构

|     目录     | 功能                                   |
|:----------:|:-------------------------------------|
|   `apps`   | 基本功能,所有功能模块全在这里面                     |
|  `deploy`  | 部署文件，本项目部署是利用`docker-compose`进行部署的，因此里面的文档为部署文档 |
|   `pkg`    | 所有开发过程中的全局通用代码。                      |
| `resource` | 项目启动或生成的资源文件存放目录。                    |
| `uploads`  | 存储上传的文件的地方                           |

更多功能请查看源码。

---
版权说明
---

* 结合Go-Admin及Panda低代码开发采用AGPL-3.0技术协议

#### 💌 其他说明

暂无