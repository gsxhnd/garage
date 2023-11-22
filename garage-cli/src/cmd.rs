use clap::{value_parser, Arg, ArgAction, Command, Parser, Subcommand};
use std::path::PathBuf;

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
        .subcommand(jav_sync_db_bf_cmd())
}

pub fn jav_sync_db_bf_cmd() -> Command {
    Command::new("jav_sync_db_bf")
        .about("crawl tenhou data")
        .arg(
            Arg::new("from_csv")
                .long("from_csv")
                .value_name("FILE")
                .required(true)
                .value_parser(value_parser!(PathBuf))
                .help("info csv file path"),
        )
        .arg(
            Arg::new("db_path")
                .long("db_path")
                .value_name("FILE")
                .required(true)
                .value_parser(value_parser!(PathBuf))
                .help("billfish database path"),
        )
}
