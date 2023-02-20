# 爬虫相关命令

帮助你爬取指定番号，演员和番号系列数据，磁力链接和封面。

支持网站： `javbus`

支持功能:

- [x] 数据基础信息
- [x] Cover 图片下载
- [ ] 磁力连接保存

## jav_code 命令

```shell
$ ./build/garage-darwin-amd64 jav_code  --help
NAME:
   garage-darwin-amd64 code - 根据指定番号爬取数据

USAGE:
   crawl --site [javbus/javlibrary] XXX-001

DESCRIPTION:
   crawl jav data, support javbus and javlibrary site.

OPTIONS:
   --proxy value  代理配置
   --site value  选择爬取数据的网站 (default: "javbus")
   --help, -h    show help (default: false)
```

```shell
# 抓取所有影片封面和信息，保存到当前目录下的 javs 文件夹下以番号命名的子文件夹中
garage code --proxy "http://127.0.0.1:7890" xxx-01
```

## jav_star 命令

```shell
$ ./build/garage-darwin-amd64 jav_star --help
NAME:
   garage-darwin-amd64 star - 根据演员ID爬取数据

USAGE:
   garage-darwin-amd64 star [command options] [arguments...]

OPTIONS:
   --proxy value  代理配置
   --help, -h  show help (default: false)
```

## crawl_prefix 命令

```shell
$ ./build/garage-darwin-amd64 jav_prefix --help
NAME:
   garage-darwin-amd64 prefix - 根据番号前缀爬取数据

USAGE:
   garage-darwin-amd64 prefix [command options] [arguments...]

OPTIONS:
   --proxy value  代理配置
   --help, -h  show help (default: false)
```
