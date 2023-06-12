# 爬虫相关命令

帮助你爬取指定番号，演员和番号系列数据，磁力链接和封面。

支持网站： `javbus`

支持功能:

- [x] 数据基础信息
- [x] Cover 图片下载
- [ ] 磁力连接保存
- [x] 常规番号爬取
- [x] 番号前缀批量爬取
- [ ] 根据演员 ID 批量爬取

```shell
# 抓取所有影片封面和信息，保存到当前目录下的 javs 文件夹下以番号命名的子文件夹中
$ garage crawl jav-code --help
$ garage crawl jav-path-code --help
$ garage crawl jav-prefix-code --help
$ garage crawl jav-star-code --help

garage jav_code --proxy "http://127.0.0.1:7890" xxx-01
```
