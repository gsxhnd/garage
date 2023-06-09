# 视频批量处理

> 系统环境中需要预先安装 ffmpeg

使用 FFMPEG 对多个视频快速转码或其他操作

## 字幕添加

添加字幕文件替换掉已有字幕， 只支持插入单条字幕。

```shell
# 查看帮助
$ garage video_import_subtitle --help
```

```shell
## ffmpeg 对应命令
ffmpeg.exe -i INPUT -c copy  -i INPUT.ass -sub_charenc UTF-8  -c copy -map 0 -map -0:s -map 1 -metadata:s:s:0 language=chi -metadata:s:s:0 title="jp&sc" OUTPUT
## garage 对应命令
garage video_subtitle --source_root_path="queue" --source_video_type="mkv" --source_subtitle_type=".ass" --dest_path="result/" --dest_video_type=".mkv" --exec=true
```

添加字幕同时添加多个字体

```shell
ffmpeg.exe -i INPUT -c copy -attach INPUT_FONT -metadata:s:t:0 mimetype=application/x-truetype-font
```

```shell
garage video_subtitle --source_root_path="queue" --source_video_type="mkv" --source_subtitle_type=".ass" --dest_path="result/" --dest_video_type=".mkv" --advance="-attach INPUT_FONT -metadata:s:t:0 mimetype=application/x-truetype-font" --exec=true
```

## 视频转码

```shell
## ffmpeg 对应命令
ffmpeg.exe -i INPUT -c:v h264_nvenc output.mp4

## garage 对应命令
garage.exe video_convert --source_root_path="queue" --source_video_type="mkv"  --dest_path="result/" --dest_video_type=".mkv" --advance="-c:v h264_nvenc"
```

h256_10bit 转码

```shell
ffmpeg -i INPUT -c:v libx265 -crf 20 -pix_fmt yuv420p10le

## garage 对应命令
garage.exe video_convert --source_root_path="queue" --source_video_type="mkv"  --dest_path="result/" --dest_video_type=".mkv" --advance="-c:v libx265 -crf 20 -pix_fmt yuv420p10le"
```

h256_10bit 转码 nvdia 硬件加速

```shell
ffmpeg -i INPUT -c:v hevc_nvenc -pix_fmt p010le -rc vbr -cq:v 27 OUTPUT

## garage 对应命令
garage.exe video_convert --source_root_path="queue" --source_video_type="mkv"  --dest_path="result/" --dest_video_type=".mkv" --advance="-c:v hevc_nvenc -pix_fmt p010le -rc vbr -cq:v 25"
```

## 切割视频 && 导出字幕文件

推荐使用 [Lossless Cut](https://github.com/mifi/lossless-cut)软件

```shell
ffmpeg -i input.wmv -ss 00:00:30.0 -c copy -t 00:00:10.0 output.wmv
```

```shell
ffmpeg -i input.mkv -map 0:s:0 subs.ass
```
