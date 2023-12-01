mod option;
use std::io;
use std::path;
use std::path::PathBuf;
use std::process::Command;
use std::process::Stdio;
use walkdir::WalkDir;

pub use option::BatchffmpegOptions;

pub struct Batchffmpeg {
    option: option::BatchffmpegOptions,
}

impl Batchffmpeg {
    pub fn new(opt: option::BatchffmpegOptions) -> Self {
        Batchffmpeg { option: opt }
    }

    pub fn create_dest_dir() {}

    pub fn get_video_list(&self, input_path: PathBuf) -> Result<(), io::Error> {
        for entry in WalkDir::new(input_path) {
            println!("{}", entry?.path().display());
        }

        Ok(())
    }

    pub fn get_fonts_params(&self) {}

    pub fn convert(&self) {
        let _ = self.get_video_list(self.option.input_path.clone());

        let mut cmd = Command::new("ping")
            .arg("www.baidu.com")
            .stdout(Stdio::inherit())
            .spawn()
            .expect("fail to execute");
        cmd.wait().unwrap();
    }

    pub fn add_sub(&self) {}

    pub fn add_fonts(&self) {}
}
