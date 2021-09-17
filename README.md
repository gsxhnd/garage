# Garage

支持网站： `javbus`,

支持功能:

- [x] 数据基础信息
- [x] Cover 图片下载
- [ ] 磁力连接保存

## 下载命令行工具 🔧

命令行工具提供`Windows`、`macOS`、`Linux`平台。

### 最新下载地址

<https://github.com/gsxhnd/garage/releases>

## 命令行选项

```shell
$ ./build/garage-darwin-amd64 --help
NAME:
   garage-darwin-amd64 - garage

USAGE:
   garage-darwin-amd64 [global options] command [command options] [arguments...]

COMMANDS:
   crawl    crawl jav data.
   version
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --proxy value
   --help, -h     show help (default: false)
```

### 全局选项

#### 代理

设置代理请求代理，非代理池。

### Crawl 命令

#### 选项

```shell
$ ./build/garage-darwin-amd64 crawl  --help
NAME:
   garage-darwin-amd64 crawl - crawl jav data.

USAGE:
   crawl --site [javbus/javlibrary] -c XXX-001

DESCRIPTION:
   crawl jav data, support javbus and javlibrary site.

OPTIONS:
   --site value              选择爬取数据的网站 (default: "javbus")
   --code value, -c value    -c xxx-001
   --help, -h                show help (default: false)
```

### 事例

```shell
# 抓取所有影片封面和信息，保存到当前目录下的 javs 文件夹下以番号命名的子文件夹中
garage --proxy [proxy_host] crawl -c xxx-01
```

