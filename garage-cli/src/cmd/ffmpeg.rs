use clap::{value_parser, Arg, ArgAction, Command, Parser, Subcommand};
use std::path::PathBuf;

pub fn ffmpeg_cmd() -> Command {
    Command::new("ffmpeg-batch")
        .about("ffmpeg batch")
        .subcommand(convert())
        .subcommand(add_sub())
        .subcommand(add_fonts())
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
