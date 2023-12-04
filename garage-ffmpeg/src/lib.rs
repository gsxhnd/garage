mod option;
use std::io;
use std::path::Path;
use std::path::PathBuf;
use std::process::Command;
use std::process::Stdio;
use walkdir::{DirEntry, WalkDir};

pub use option::BatchffmpegOptions;

pub struct Batchffmpeg {
    option: option::BatchffmpegOptions,
}

impl Batchffmpeg {
    pub fn new(opt: option::BatchffmpegOptions) -> Self {
        Batchffmpeg { option: opt }
    }

    pub fn create_dest_dir() {}

    pub fn get_video_list(&self, input_path: PathBuf, input_format: String) -> Vec<String> {
        let mut file_list: Vec<String> = Vec::new();
        for entry in WalkDir::new(input_path) {
            match entry {
                Err(err) => {
                    println!("error: {:?}", err);
                }
                Ok(dir_entry) => {
                    if dir_entry.file_type().is_file() {
                        match Path::new(dir_entry.file_name()).extension() {
                            Some(format) => {
                                if format.to_str() == Some(input_format.trim()) {
                                    println!("format: {:?}", format);
                                    file_list
                                        .push(dir_entry.file_name().to_str().unwrap().to_string())
                                }
                            }
                            None => {}
                        }
                    }
                }
            }
        }
        file_list
    }

    pub fn get_fonts_params(&self) {}

    pub fn convert(&self) {
        let video_list = self.get_video_list(
            self.option.input_path.clone(),
            self.option.input_format.clone(),
        );
        println!("video_list: {:?}", video_list);

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
