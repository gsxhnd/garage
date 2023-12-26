mod ffmpeg;
mod jav;
mod spider;
mod tenhou;
use clap::Command;
use log;
use tracing::info;
use tracing_log::LogTracer;
use tracing_subscriber::{fmt, layer::SubscriberExt, util::SubscriberInitExt};

#[tokio::main]
async fn main() {
    let _ = LogTracer::init();
    log::warn!("test");
    // tracing_subscriber::registry().with(fmt::layer()).init();

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
        Some(("spider", _sub_m)) => {}
        Some(("tenhou", _sub_m)) => todo!(),
        Some(("ffmpeg-batch", sub_m)) => {
            info!("ffmpeg-batch starting...");
            let sub_cmd = sub_m.subcommand().unwrap();
            ffmpeg::parse_ffmpeg_cmd(sub_cmd.0, sub_cmd.1)
        }
        _ => {}
    }
}
