mod cmd;
mod ffmpeg;
mod jav;
use ruspider::Ruspider;

#[tokio::main]
async fn main() {
    let cmd = cmd::new_cmd();
    let _ = Ruspider::new();

    match cmd.get_matches().subcommand() {
        Some(("jav", sub_m)) => {
            println!("jav:: sync_db_bf");
            let sub_cmd = sub_m.subcommand().unwrap();
            jav::jav_sub_cmd(sub_cmd.0, sub_cmd.1).await;
        }
        Some(("crawl", _sub_m)) => todo!(),
        Some(("tenhou", _sub_m)) => todo!(),
        Some(("ffmpeg-batch", sub_m)) => {
            let sub_cmd = sub_m.subcommand().unwrap();
            ffmpeg::sub_ffmpeg_cmd(sub_cmd.0, sub_cmd.1)
        }
        _ => {}
    }
}
