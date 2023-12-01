use clap::ArgMatches;
use garage_ffmpeg::{Batchffmpeg, BatchffmpegOptions};
use std::path::PathBuf;

pub fn sub_ffmpeg_cmd(sub_cmd: &str, args: &ArgMatches) {
    match sub_cmd {
        "convert" => {
            let input_path = args
                .get_one::<PathBuf>("from_csv")
                .expect("csv path not set")
                .to_owned();

            let opt = BatchffmpegOptions::new().input_path(input_path);
            let f = Batchffmpeg::new(opt);
            f.convert();
        }
        "add_sub" => {}
        "add_fonts" => {}
        _ => {
            todo!()
        }
    }
}
