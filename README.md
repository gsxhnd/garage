![Alt](https://repobeats.axiom.co/api/embed/a600b32115176854e9ab824307882a9350d5d9de.svg "Repobeats analytics image")

[![codecov](https://codecov.io/gh/gsxhnd/garage/branch/master/graph/badge.svg?token=6EAEPP8LT7)](https://codecov.io/gh/gsxhnd/garage)

# Garage

`garage`是一个命令行工具，提供爬虫功能。

[更新日志](./CHANGELOG.md)

## 下载命令行工具 🔧

命令行工具提供`Windows`、`macOS`、`Linux`平台预编译的二进制。

最新下载地址: <https://github.com/gsxhnd/garage/releases>

```bash
garage --help # 获取帮助
```

## 命令行文档

帮助你爬取指定番号，演员和番号系列数据，磁力链接和封面。

支持网站： `javbus`

支持功能:

- [x] 常规番号爬取
- [x] 番号前缀批量爬取
- [x] 根据演员 ID 批量爬取
- [x] 遍历本地文件视频爬取对应信息
- [x] Cover 图片下载
- [x] 数据基础信息下载
- [x] 磁力连接保存

```shell
# 抓取所有影片封面和信息，保存到当前目录下的 javs 文件夹下以番号命名的子文件夹中
$ garage crawl jav-code --help
$ garage crawl jav-path-code --help
$ garage crawl jav-prefix-code --help
$ garage crawl jav-star-code --help

garage jav_code --proxy "http://127.0.0.1:7890" xxx-01
```

## Thanks

Thank for JetBrains Open Source License

[![JetBrains logo](https://resources.jetbrains.com/storage/products/company/brand/logos/jetbrains.png)](https://jb.gg/OpenSourceSupport)
