# e_downloader

某鹅通付费可观看视频下载

## 用前提示

+ 本软件只能用于下载付费后可观看的微信某鹅通视频
+ 下载后的资源传播与本软件无关

## 使用流程

### 1. 浏览器安装猫抓插件

+ [猫抓下载地址](https://chrome.zzzmh.cn/info?token=jfedfbgedapdagkghmgibemcoggfppbb)

### 2. 获取m3u8链接

+ 打开付费可观看的视频，使用猫抓插件获取其中带m3u8的链接
  + 一般为第一条
  + 如果视频可以切换清晰度，默认为720p，切换后猫抓会获取到另外的m3u8链接，注意区分

### 3. 下载器下载

+ [下载器项目地址](https://github.com/oopsguy/m3u8)，在此致敬oopsguy老哥
+ [下载器分享链接](https://dusklight-my.sharepoint.com/:f:/g/personal/qin_dusklight_onmicrosoft_com/EnFy4WhKYsBDia98O_XaA2wBf5_RmAwFf44_gOLl_CuiSQ?e=5xAavL)
+ 将其与e_downloader放一起即可

### 4. 写入文件

+ 请务必写成这个格式，可以自动创建文件夹以及下载链接到指定目录，更为方便美观

```shell
# 文件格式
章节标题1 	# 一、xxx			二、xxx
视频名称1	# 1.	2.		3. 
m3u8链接1
视频名称2
m3u8链接2
章节标题2
视频名称3
m3u8链接3
...
hello world 	# 最后一行务必加一个hello wolrd（其实加什么都可以，但不能不加）
```

