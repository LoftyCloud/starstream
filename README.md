## 项目介绍：

StarStream是一个个人博客类型的多媒体分享网站，站主和授权用户拥有上传视频的权限，注册用户可以自由浏览和下载这些内容，并对内容进行收藏，类似于视频版的个人博客。普通用户可以浏览视频，但是上传、编辑和删除视频需要站主为用户授予权限。

## 现状：

图片上传功能还没有实现，视频上传基本完成。有可能的话之后还想让用户能够对视频进行评论。
目前项目已经部署，但是部署的时候遇到了些请求问题，现在还没有完全部署成功。

11.25更新：基本部署成功了，项目本身还有待完善：http://117.72.104.1

测试管理员账号：`admin` （也可以自己注册）

密码：`admin`

有时间会先解决部署遇到的问题，并且会继续完善项目的功能。

## 前端
前端基于vue3和vite。

## 后端
项目后端使用Go语言编写，使用Gin框架。

技术栈：Gin, Gorm, MySQL, Redis, FFmpeg(视频处理)

项目架构：

1、路由

项目的路由文件都在router/Router.go文件中，静态资源、上传文件的规则以及Token的校验和JWT的中间件都在Router内配置。

2、控制器

项目的控制器全都在controllers的包内，每一个Controller都有一个对应的Services和Models以及Router，Controller进行逻辑处理反馈接收的数据。

3、业务层

项目的业务层都在services的包内，进行业务的处理，每一个service对应着一个controller以及models，业务层是进行业务的逻辑和数据处理（在这里将数据写入数据库），可复用性高。

4、缓冲层

项目采用Redis与Mysql进行数据的读写分离，在该项目内caches的包内是处理Redis逻辑业务的处理，缓冲热门数据以及业务的频繁读取的数据，从而达到业务的优化和提升用户的使用体验。

5、数据模型

项目的所有的结构体都在models的包内，每一个model都对应着一个controller和service，基于业务建立数据模型，而不是基于数据表。

6、配置文件

项目的配置文件在configs/config.yaml中，包含了监听端口号、数据库连接信息、Redis连接信息、文件路径，加密盐值等。

## 项目部署
配置 Nginx 进行反向代理和负载均衡。
