use clap::{ArgMatches, Command};

pub fn spider_cmd() -> Command {
    Command::new("spider").about("jav data crawl and process")
}

pub async fn parse_spider_cmd(sub_cmd: &str, args: &ArgMatches) {}
