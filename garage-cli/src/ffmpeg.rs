use clap::ArgMatches;
use garage_ffmpeg::Batchffmpeg;

pub fn sub_ffmpeg_cmd(sub_cmd: &str, args: &ArgMatches) {
    match sub_cmd {
        "convert" => {
            let f = Batchffmpeg::new();
            f.convert();
        }
        "add_sub" => {}
        "add_fonts" => {}
        _ => {
            todo!()
        }
    }
}
