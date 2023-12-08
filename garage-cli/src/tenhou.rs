use clap::{value_parser, Arg, ArgAction, Command, Parser, Subcommand};
use std::path::PathBuf;

pub fn tenhou_cmd() -> Command {
    Command::new("tenhou").about("tenhou data")
}
