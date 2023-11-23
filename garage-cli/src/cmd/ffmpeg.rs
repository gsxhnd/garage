use clap::{value_parser, Arg, ArgAction, Command, Parser, Subcommand};
use std::path::PathBuf;

pub fn ffmpeg_cmd() -> Command {
    Command::new("ffmpeg-batch").about("ffmpeg batch")
}
