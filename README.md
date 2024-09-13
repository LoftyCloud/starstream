## 项目介绍：

前后端分离项目，主要是想要实现一个能够供自己上传平时遇到的好玩的视频和图片的网站，这样能够在任何设备上进行访问，不用在各个平台上查找，并且可以很方便地分享给朋友。普通用户可以浏览视频，但是上传、编辑和删除视频需要站主为用户授予权限。

## 现状：

目前图片上传功能还没有实现，视频上传基本完成，主要是前端业务不熟悉真的很难改。
而且部署的时候遇到了问题，现在还没有完全部署成功，可以先来这里看一眼：http://8.210.139.128，有时间会继续完善项目的功能，并且解决部署遇到的问题

## 前端
前端基于vue3和vite

## 后端
项目后端使用Go语言编写，使用GO+Gin+Gorm+Redis框架

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

项目的配置文件都在configs/config.yaml中，包含了监听端口号、数据库连接信息、Redis连接信息、文件路径，加密盐值等。

## 项目部署
