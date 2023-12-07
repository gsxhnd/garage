use clap::{value_parser, Arg, ArgAction, ArgGroup, Command, Parser, Subcommand};
use std::path::PathBuf;

struct Args {
    input_path: Arg,
    input_format: Arg,
    output_path: Arg,
    output_format: Arg,
    advance: Arg,
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
            .default_value("mp4")
            .help("input directory path"),
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
    };

    Command::new("ffmpeg-batch")
        .about("ffmpeg batch")
        .subcommand(convert().args([
            args.input_path.clone(),
            args.input_format.clone(),
            args.output_path.clone(),
            args.output_format.clone(),
            args.advance.clone(),
        ]))
        .subcommand(add_sub().args([args.input_path.clone()]))
        .subcommand(add_fonts().args([args.input_path.clone()]))
}

pub fn convert() -> Command {
    Command::new("convert").about("")
}

pub fn add_sub() -> Command {
    Command::new("add_sub").about("")
}

pub fn add_fonts() -> Command {
    Command::new("add_fonts").about("")
}
