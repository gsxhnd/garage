# Garage

Help you sync local jav's info.

## 命令行

### 爬虫命令

爬取指定视频对应的信息,封面。

爬取数据暂时不提供磁链链接。

```shell
NAME:
   garage.exe crawl - crawl jav data.

USAGE:
   crawl --site [javbus/javlibrary] -s XXX-001

DESCRIPTION:
   crawl jav data, support javbus and javlibrary site.

OPTIONS:
   --search value, -s value
   --site value              选择爬取数据的网站 (default: "javbus")
   --base value, -b value
   --star value, -t value
   --help, -h                show help (default: false)
```



### API命令

`API`命令启动接口服务。服务提供多个接口服务：

- 后台数据爬取
- `jav`和`star`数据信息
- 封面数据

### UI

`ui`部分提供`Windows`,`macOS`, 未来会支持 `linux` ,`iOS`,`Android`

仓库地址 https://github.com/gsxhnd/garage_app



