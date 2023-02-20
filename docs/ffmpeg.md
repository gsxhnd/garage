# 视频批量处理

> 系统环境中需要预先安装 ffmpeg

使用 FFMPEG 对多个视频快速转码或其他操作

## 字幕添加

添加字幕文件替换掉已有字幕

```shell
$ garage video_subtitle --help

NAME:
   garage video_subtitle - 视频添加字幕批处理

USAGE:
   garage video_subtitle [command options] [arguments...]

OPTIONS:
   --source_root_path value          源视频路径 (default: "./")
   --source_video_type value         源视频后缀 (default: ".mkv")
   --source_subtitle_type value      添加的字幕后缀 (default: ".ass")
   --source_subtitle_number value    添加的字幕所处流的位置 (default: 0)
   --source_subtitle_language value  添加的字幕语言缩写，其他语言请参考ffmpeg (default: "chi")
   --source_subtitle_title value     添加的字幕标题 (default: "Chinese")
   --dest_path value                 转换后文件存储位置 (default: "./result/")
   --dest_video_type value           转换后的视频后缀 (default: ".mkv")
   --advance value                   高级自定义参数
   --exec                            是否执行批处理命令，False时仅打印命令 (default: true)
   --help, -h                        show help (default: false)
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

## 导出字幕文件

```shell
ffmpeg -i input.mkv -map 0:s:0 subs.ass
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
garage.exe video_convert --source_root_path="queue" --source_video_type="mkv"  --dest_path="result/" --dest_video_type=".mkv" --advance="-c:v hevc_nvenc -pix_fmt p010le -rc vbr -cq:v 27"
```
