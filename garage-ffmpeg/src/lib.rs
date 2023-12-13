mod option;
use std::fmt::format;
use std::path::Path;
use std::path::PathBuf;
use std::process::Command;
use std::process::Stdio;
use tracing::info;
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

    pub fn convert(&self) {
        let video_list = self.get_video_list(
            self.option.input_path.clone(),
            self.option.input_format.clone(),
        );
        let mut args: Vec<Vec<String>> = Vec::new();

        for v in video_list {
            let mut output = PathBuf::new();
            let mut input = PathBuf::new();

            input.push(self.option.input_path.clone());
            input.push(v.clone());

            let file_name = Path::new(&v)
                .file_name()
                .unwrap_or_default()
                .to_str()
                .unwrap_or_default();
            output.push(self.option.output_path.clone());
            output.push(format!(
                "{}.{}",
                file_name,
                self.option.output_format.clone()
            ));

            let arg = vec![
                "-i".to_string(),
                input.to_str().unwrap().to_owned(),
                output.to_str().unwrap().to_owned(),
            ];
            args.push(arg);
        }
        if self.option.exec {
            for arg in args {
                let mut child = Command::new("ffmpeg")
                    .args(arg)
                    .stdout(Stdio::inherit())
                    .spawn()
                    .expect("fail to execute");
                child.wait().unwrap();
            }
        } else {
            for arg in args {
                info!("ffmpeg {} {} {}", arg[0], arg[1], arg[2])
            }
        }
    }

    pub fn add_sub(&self) {}

    pub fn add_fonts(&self) {
        let video_list = self.get_video_list(
            self.option.input_path.clone(),
            self.option.input_format.clone(),
        );
        let font_list = self.get_fonts_params(self.option.font_path.clone());
        for f in font_list {
            println!("font: {}", f);
        }
    }

    pub fn get_fonts_params(&self, font_path: PathBuf) -> Vec<String> {
        let mut file_list: Vec<String> = Vec::new();
        let font_format_list = vec!["oft", "ttf"];
        for entry in WalkDir::new(font_path) {
            match entry {
                Err(err) => {
                    println!("error: {:?}", err);
                }
                Ok(dir_entry) => {
                    if dir_entry.file_type().is_file() {
                        match Path::new(dir_entry.file_name()).extension() {
                            Some(format) => {
                                let file_name = dir_entry.file_name().to_str().unwrap();
                                let font_format = format.to_str().unwrap();
                                if font_format_list.contains(&font_format) {
                                    file_list.push(format!("{}.{}", file_name, font_format))
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
}
