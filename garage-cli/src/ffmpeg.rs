use clap::{value_parser, Arg, ArgAction, ArgGroup, ArgMatches, Command, Parser, Subcommand};
use garage_ffmpeg::{Batchffmpeg, BatchffmpegOptions};
use std::path::PathBuf;

struct Args {
    input_path: Arg,
    input_format: Arg,
    font_path: Arg,
    sub_suffix: Arg,
    sub_number: Arg,
    output_path: Arg,
    output_format: Arg,
    advance: Arg,
    exec: Arg,
}

pub fn ffmpeg_cmd() -> Command {
    let args = Args {
        input_path: Arg::new("input_path")
            .long("input_path")
            .value_name("FILE")
            .value_parser(value_parser!(PathBuf))
            .required(true)
            .help("input directory path"),
        input_format: Arg::new("input_format")
            .long("input_format")
            .value_name("FILE Extension")
            .value_parser(value_parser!(String))
            .help("input directory path"),
        font_path: Arg::new("font_path")
            .long("font_path")
            .value_name("FILE")
            .value_parser(value_parser!(PathBuf))
            .required(true)
            .help("input fonts directory path"),
        sub_suffix: Arg::new("sub_suffix")
            .long("sub_suffix")
            .value_name("FILE")
            .value_parser(value_parser!(String))
            .default_value("ass")
            .help("sub suffix and extension"),
        sub_number: Arg::new("sub_number")
            .long("sub_number")
            .value_name("FILE")
            .value_parser(value_parser!(u32))
            .help("sub number"),
        output_path: Arg::new("output_path")
            .long("output_path")
            .value_name("FILE")
            .default_value("./dest")
            .value_parser(value_parser!(PathBuf))
            .help("output directory path"),
        output_format: Arg::new("output_format")
            .long("output_format")
            .value_name("FILE Extension")
            .default_value("mkv")
            .value_parser(value_parser!(String))
            .help("output directory path"),
        advance: Arg::new("advance")
            .long("advance")
            .value_name("")
            .value_parser(value_parser!(String))
            .help("advance string"),
        exec: Arg::new("exec")
            .long("exec")
            .short('e')
            .action(ArgAction::SetTrue)
            .help("exec command"),
    };

    Command::new("ffmpeg-batch")
        .about("ffmpeg batch")
        .subcommand(Command::new("convert").about("").args([
            args.input_path.clone(),
            args.input_format.clone().default_value("mp4"),
            args.output_path.clone(),
            args.output_format.clone(),
            args.advance.clone(),
            args.exec.clone(),
        ]))
        .subcommand(Command::new("add_sub").about("").args([
            args.input_path.clone(),
            args.input_format.clone().default_value("mkv"),
            args.sub_suffix.clone(),
            args.sub_number.clone().default_value("0"),
            args.output_path.clone(),
            args.exec.clone(),
        ]))
        .subcommand(Command::new("add_fonts").about("").args([
            args.input_path.clone(),
            args.input_format.clone().default_value("mkv"),
            args.font_path.clone(),
            args.output_path.clone(),
            args.output_format.clone(),
            args.advance.clone(),
            args.exec.clone(),
        ]))
}

pub fn parse_ffmpeg_cmd(sub_cmd: &str, args: &ArgMatches) {
    let input_path = args
        .get_one::<PathBuf>("input_path")
        .expect("input path not set")
        .to_owned();
    let input_format = args.get_one::<String>("input_format").expect("").to_owned();
    let output_path = args.get_one::<PathBuf>("output_path").expect("").to_owned();
    let output_format = args
        .get_one::<String>("output_format")
        .expect("")
        .to_owned();
    let advance = args.try_get_one::<String>("advance").expect("").to_owned();
    let exec = args.get_flag("exec");

    let opt = BatchffmpegOptions::new()
        .input_path(input_path)
        .input_format(input_format)
        .output_path(output_path)
        .output_format(output_format)
        .advance(advance)
        .exec(exec);

    match sub_cmd {
        "convert" => {
            let f = Batchffmpeg::new(opt);
            f.convert();
        }
        "add_sub" => {
            let opt = opt.sub_suffix("".to_string());
            let f = Batchffmpeg::new(opt);
            f.add_sub();
        }
        "add_fonts" => {
            let font_path = args.get_one::<PathBuf>("font_path").expect("").to_owned();
            let opt = opt.font_path(font_path);
            let f = Batchffmpeg::new(opt);
            f.add_fonts();
        }
        _ => {
            todo!()
        }
    }
}
