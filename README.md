# Garage

`garage`是一个命令行工具，帮助你爬取指定番号，演员和番号系列数据，磁力链接和封面。

支持网站： `javbus`,

支持功能:

- [x] 数据基础信息
- [x] Cover 图片下载
- [ ] 磁力连接保存

[更新日志](./CHANGELOG.md)

## 下载命令行工具 🔧

命令行工具提供`Windows`、`macOS`、`Linux`平台。

最新下载地址: <https://github.com/gsxhnd/garage/releases>

## 命令行选项

```shell
$ ./build/garage-darwin-amd64 help
NAME:
   garage - JAV命令行工具

USAGE:
   garage [global options] command [command options] [arguments...]

COMMANDS:
   code     根据指定番号爬取数据
   star     根据演员ID爬取数据
   prefix   根据番号前缀
   version
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
```

### 全局选项

#### 代理

设置代理请求代理，非代理池。

### code 命令

```shell
$ ./build/garage-darwin-amd64 crawl  --help
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

### Star 命令

```shell
NAME:
   garage-darwin-amd64 star - 根据演员ID爬取数据

USAGE:
   garage-darwin-amd64 star [command options] [arguments...]

OPTIONS:
   --proxy value  代理配置
   --help, -h  show help (default: false)
```

### prefix 命令

```shell
NAME:
   garage-darwin-amd64 prefix - 根据番号前缀爬取数据

USAGE:
   garage-darwin-amd64 prefix [command options] [arguments...]

OPTIONS:
   --proxy value  代理配置
   --help, -h  show help (default: false)
```
