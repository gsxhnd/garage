use clap::ArgMatches;
use garage_ffmpeg::{Batchffmpeg, BatchffmpegOptions};
use std::path::PathBuf;

pub fn sub_ffmpeg_cmd(sub_cmd: &str, args: &ArgMatches) {
    match sub_cmd {
        "convert" => {
            let input_path = args
                .get_one::<PathBuf>("input_path")
                .expect("input path not set")
                .to_owned();
            let input_format = args.get_one::<String>("input_format").expect("").to_owned();
            let output_path = args.get_one::<PathBuf>("output_path").expect("").to_owned();
            let output_formart = args
                .get_one::<String>("output_format")
                .expect("")
                .to_owned();
            let advance = args.try_get_one::<String>("advance").expect("").to_owned();

            let opt = BatchffmpegOptions::new()
                .input_path(input_path)
                .input_format(input_format)
                .output_path(output_path)
                .output_format(output_formart)
                .advance(advance);
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
