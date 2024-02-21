# 视频批量处理

> 系统环境中需要预先安装 ffmpeg

使用 FFMPEG 对多个视频快速转码或其他操作

- [x] convert 批量视频转码
- [ ] add_sub 批量视频添加字幕
- [ ] addfonts 批量视频添加字体

```shell
# 子命令
convert    batch video convert other format
add_sub    batch video add one suffix subtittle
add_fonts  batch video add fonts from dir 
```

## 视频转码

```shell
# 硬件加速 https://trac.ffmpeg.org/wiki/HWAccelIntro#VideoToolbox
# h264 转码
ffmpeg.exe -i input.mp4 -c:v h264_nvenc output.mkv
# h265 cpu 10bit 转码
ffmpeg -i INPUT.mp4 -c:v libx265 -crf 20 -pix_fmt yuv420p10le OUTPUT.mkv
# h265 nvida加速 10bit 转码
ffmpeg -i INPUT.mp4 -c:v hevc_nvenc -pix_fmt p010le -rc vbr -cq:v 27 OUTPUT.mkv
# h264/h265 apple silicon 硬件加速转码
# q:v 1-100 1 lowest, 100 highest
ffmpeg -i INPUT.mp4 -c:v h264_videotoolbox  -pix_fmt yuv444p -cq:v 27 OUTPUT.mkv
ffmpeg -i INPUT.mp4 -c:v hevc_videotoolbox  -pix_fmt yuv444p -cq:v 27 OUTPUT.mkv
```

```shell
$ ffmpeg-batch convert --help

--input_path <>               input directory path
--input_format <>             input directory path [default: mp4]
--output_path <>              output directory path [default: ./dest]
--output_format <>            output directory path [default: mkv]
--advance <>                  advance string
--exec                        exec command
```

## 字幕添加

添加字幕文件替换掉已有字幕， 只支持插入单条字幕。

```shell
## ffmpeg 对应命令
ffmpeg.exe -i INPUT -c copy  -i INPUT.ass -sub_charenc UTF-8  -c copy -map 0 -map -0:s -map 1 -metadata:s:s:0 language=chi -metadata:s:s:0 title="jp&sc" OUTPUT
```

```shell
## garage 对应命令
$ garage ffmpeg-batch add-sub

--input_path <FILE>               input directory path
--input_format <FILE Extension>   input directory path [default: mkv]
--sub_suffix <FILE>               sub suffix and extension [default: ass]
--sub_number <FILE>               sub number [default: 0]
--output_path <FILE>              output directory path [default: ./dest]
--output_format <FILE Extension>  output directory path [default: mkv]
--advance <>                      advance string
--exec                            exec command
```

## 添加多个字体

```shell
## ffmpeg 对应命令
ffmpeg.exe -i INPUT -c copy -attach INPUT_FONT -metadata:s:t:0 mimetype=application/x-truetype-font

## garage 对应命令
garage ffmpeg-batch add-sub \
--input-path="queue" \
--input-type=".mkv" \
--input-sub-suffix=".ass" \
--input-fonts-path="fonts" \
--output-path="result/" \
--exec=true
```

## 切割视频 && 导出字幕文件

推荐使用 [Lossless Cut](https://github.com/mifi/lossless-cut)软件

```shell
ffmpeg -i input.wmv -ss 00:00:30.0 -c copy -t 00:00:10.0 output.wmv
```

```shell
ffmpeg -i input.mkv -map 0:s:0 subs.ass
```
