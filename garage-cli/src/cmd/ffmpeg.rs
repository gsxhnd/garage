use clap::{value_parser, Arg, ArgAction, ArgGroup, Command, Parser, Subcommand};
use std::path::PathBuf;

struct Args {
    input_path: Arg,
}

pub fn ffmpeg_cmd() -> Command {
    let args = Args {
        input_path: Arg::new("input_path")
            .long("input_path")
            .help("input directory path"),
    };

    Command::new("ffmpeg-batch")
        .about("ffmpeg batch")
        .subcommand(convert().args([args.input_path]))
        .subcommand(add_sub())
        .subcommand(add_fonts())
    // .args([Arg::new("input_path")
    //     .long("input_path")
    //     .help("input directory path")])
    // .args_conflicts_with_subcommands(true)
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
