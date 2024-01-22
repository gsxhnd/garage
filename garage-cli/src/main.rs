mod ffmpeg;
mod jav;
mod spider;
mod tenhou;
mod utils;
use crate::utils::Logger;

use clap::Command;
use tracing::info;

#[tokio::main]
async fn main() {
    Logger::new().init();

    let cmd = Command::new("garage")
        .bin_name("garage")
        .version("v1")
        .about("about")
        .subcommand_required(true)
        // .subcommand(jav::jav_cmd())
        // .subcommand(ffmpeg::ffmpeg_cmd())
        .subcommand(spider::spider_cmd())
        .subcommand(tenhou::tenhou_cmd());

    match cmd.get_matches().subcommand() {
        Some(("jav", sub_m)) => {
            let sub_cmd = sub_m.subcommand().unwrap();
            // jav::parse_jav_cmd(sub_cmd.0, sub_cmd.1).await;
        }
        Some(("spider", _sub_m)) => {
            info!("todo");
        }
        Some(("tenhou", _sub_m)) => {
            info!("todo");
        }
        Some(("ffmpeg-batch", sub_m)) => {
            info!("ffmpeg-batch starting...");
            let sub_cmd = sub_m.subcommand().unwrap();
            // ffmpeg::parse_ffmpeg_cmd(sub_cmd.0, sub_cmd.1)
        }
        _ => {}
    }
}
