use clap::ArgMatches;
use clap::{value_parser, Arg, ArgAction, Command, Parser, Subcommand};
use garage_jav::{Crawl, DbSyncBf};
use std::path::PathBuf;

pub fn jav_cmd() -> Command {
    Command::new("jav")
        .about("jav data crawl and process")
        .subcommand(crawl_movie_code())
        .subcommand(crawl_star_code())
        .subcommand(crawl_movie_prefix_code())
        .subcommand(crawl_movie_code_from_dir())
        .subcommand(sync_db_bf_cmd())
}

pub fn crawl_movie_code() -> Command {
    Command::new("crawl_movie_code").about("crawl jav movie info by movie code")
}

pub fn crawl_star_code() -> Command {
    Command::new("crawl_star_code").about("crawl jav movie info by star code")
}

pub fn crawl_movie_prefix_code() -> Command {
    Command::new("crawl_movie_prefix_code").about("crawl jav movie info by movie prefix code")
}

pub fn crawl_movie_code_from_dir() -> Command {
    Command::new("crawl_movie_code_from_dir").about("crawl jav movie info by directory movie ext")
}

pub fn sync_db_bf_cmd() -> Command {
    Command::new("sync_db_bf")
        .about("jav info data sync to bf database")
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

pub async fn parse_jav_cmd(sub_cmd: &str, args: &ArgMatches) {
    match sub_cmd {
        "sync_db_bf" => {
            let _csv_path = args
                .get_one::<PathBuf>("from_csv")
                .expect("csv path not set")
                .to_str()
                .unwrap()
                .to_owned();
            let _db_path = args
                .get_one::<PathBuf>("db_path")
                .expect("db path not set")
                .to_str()
                .unwrap()
                .to_owned();
            let action = DbSyncBf::new().await;
            // action.sync();
        }
        "crawl_movie_code" => {
            let c = Crawl::new();
            c.start_jav_code().await;
        }
        "crawl_star_code" => {
            let c = Crawl::new();
            c.start_jav_star_code().await;
        }
        _ => println!("no complete sub command"),
    }
}
