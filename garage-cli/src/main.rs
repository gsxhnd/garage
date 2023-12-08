mod ffmpeg;
mod jav;
mod tenhou;
use clap::Command;
use ruspider::Ruspider;
use tracing::info;
use tracing_subscriber::{fmt, layer::SubscriberExt, util::SubscriberInitExt};

#[tokio::main]
async fn main() {
    tracing_subscriber::registry().with(fmt::layer()).init();

    let cmd = Command::new("garage")
        .bin_name("garage")
        .version("v1")
        .about("about")
        .subcommand_required(true)
        .subcommand(jav::jav_cmd())
        .subcommand(tenhou::tenhou_cmd())
        .subcommand(ffmpeg::ffmpeg_cmd());
    let _ = Ruspider::new();

    match cmd.get_matches().subcommand() {
        Some(("jav", sub_m)) => {
            println!("jav:: sync_db_bf");
            let sub_cmd = sub_m.subcommand().unwrap();
            jav::parse_jav_cmd(sub_cmd.0, sub_cmd.1).await;
        }
        Some(("crawl", _sub_m)) => todo!(),
        Some(("tenhou", _sub_m)) => todo!(),
        Some(("ffmpeg-batch", sub_m)) => {
            info!("ffmpeg-batch starting...");
            let sub_cmd = sub_m.subcommand().unwrap();
            ffmpeg::parse_ffmpeg_cmd(sub_cmd.0, sub_cmd.1)
        }
        _ => {}
    }
}
