# CHANGELOG

## Version 0.0.8

- REFACTOR: `ffmpeg`子命令添加自动执行参数
- DOC: 文档更新
- REFACTOR: 目录变更
  <!-- - ADD: 爬虫添加前缀爬取 -->
  <!-- - ADD: 爬虫添加前缀爬取 -->

## Version 0.0.7

- FIX: 爬虫部分没有提前创建下载数据

## Version 0.0.6

- CHORE: 编译去除路径
- REFACTOR: `ffmpeg`部分子命令不再自动执行命令而是生成脚本文件
- REFACTOR: 修改单个字体文件，使用字体文件夹代替，自动识别 TTF,OTF 文件并添加
- REFACTOR: 修改添加字幕子命令 `video_import_subtitle`
- REFACTOR: 修改导出字幕子命令 `video_export_subtitle`

## Version 0.0.5

- FIX: 参数字符处理
- FIX: 批量命令参数缺少导致异常

## Version 0.0.4

- FIX: 修改视频输出文件未知错误
- FEAT: 视频命令行功能参数添加
- FEAT: 视频命令行功能日志输出内容增加

## Version 0.0.3

- FEAT: 添加了视频处理的相关命令

## Version 0.0.2

- FEAT: 添加根据演员`ID`爬取数据
- FEAT: 添加根据番号前缀爬取数据
- REFACTOR: 爬指定番号的命令删除了`code`参数

## Version 0.0.1

- FEAT: 爬取指定番号信息数据和封面信息并保存在当前文件夹下
- FEAT: 全局代理设置
