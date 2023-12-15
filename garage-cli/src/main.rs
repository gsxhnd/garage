mod ffmpeg;
mod jav;
mod spider;
mod tenhou;
use clap::Command;
use tracing::info;
use tracing_subscriber::{fmt, layer::SubscriberExt, util::SubscriberInitExt};

use garage_jav::Crawl;

#[tokio::main]
async fn main() {
    tracing_subscriber::registry().with(fmt::layer()).init();

    let cmd = Command::new("garage")
        .bin_name("garage")
        .version("v1")
        .about("about")
        .subcommand_required(true)
        .subcommand(jav::jav_cmd())
        .subcommand(spider::spider_cmd())
        .subcommand(tenhou::tenhou_cmd())
        .subcommand(ffmpeg::ffmpeg_cmd());

    match cmd.get_matches().subcommand() {
        Some(("jav", sub_m)) => {
            let sub_cmd = sub_m.subcommand().unwrap();
            jav::parse_jav_cmd(sub_cmd.0, sub_cmd.1).await;
        }
        Some(("spider", _sub_m)) => {
            let r = Crawl::new();
            r.start();
        }
        Some(("tenhou", _sub_m)) => todo!(),
        Some(("ffmpeg-batch", sub_m)) => {
            info!("ffmpeg-batch starting...");
            let sub_cmd = sub_m.subcommand().unwrap();
            ffmpeg::parse_ffmpeg_cmd(sub_cmd.0, sub_cmd.1)
        }
        _ => {}
    }
}
