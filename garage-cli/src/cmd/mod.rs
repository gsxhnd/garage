mod ffmpeg;
mod jav;
mod tenhou;
use ffmpeg::ffmpeg_cmd;
use jav::jav_cmd;
use tenhou::tenhou_cmd;

use clap::Command;

pub fn new_cmd() -> Command {
    // let config_file_flag = Arg::new("config")
    //     .short('c')
    //     .long("config")
    //     .value_name("FILE")
    //     .global(true)
    //     .value_parser(value_parser!(PathBuf))
    //     .help("Provides a config file");

    Command::new("")
        .bin_name("garage")
        .version("v1")
        .about("about")
        .subcommand_required(true)
        .subcommand(jav_cmd())
        .subcommand(tenhou_cmd())
        .subcommand(ffmpeg_cmd())
}
