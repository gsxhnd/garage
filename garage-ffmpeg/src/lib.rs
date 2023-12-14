mod option;
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

    pub fn get_video_list(&self, input_path: PathBuf, input_format: String) -> Vec<PathBuf> {
        let mut file_list: Vec<PathBuf> = Vec::new();
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
                                    file_list.push(dir_entry.path().to_path_buf())
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
                v.to_str().unwrap().to_owned(),
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

    pub fn add_sub(&self) {
        let video_list = self.get_video_list(
            self.option.input_path.clone(),
            self.option.input_format.clone(),
        );
    }

    pub fn add_fonts(&self) {
        let mut args: Vec<Vec<String>> = Vec::new();
        let mut fonts_params: Vec<String> = Vec::new();
        let video_list = self.get_video_list(
            self.option.input_path.clone(),
            self.option.input_format.clone(),
        );
        let font_list = self.get_fonts(self.option.font_path.clone());
        for (i, font) in font_list.iter().enumerate() {
            fonts_params.push("-attach".to_string());
            fonts_params.push(font.to_str().unwrap().to_string());
            fonts_params.push(format!("-metadata:s:t:{}", i));
            fonts_params.push("mimetype=application/x-truetype-font".to_string());
        }

        if !font_list.is_empty() && !video_list.is_empty() {
            for v in video_list {
                let filename = v.file_name().unwrap().to_str().unwrap().to_string();
                let mut output = PathBuf::new();
                output.push(self.option.output_path.clone());
                output.push(filename);

                let mut arg = vec![
                    "-i".to_string(),
                    v.to_str().unwrap().to_owned(),
                    "-c".to_string(),
                    "copy".to_string(),
                ];

                arg.append(&mut fonts_params.clone());
                arg.push(output.to_str().unwrap().to_owned());
                args.push(arg);
            }
        } else {
            return;
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
                info!("ffmpeg {:?}", arg);
            }
        }
    }

    pub fn get_fonts(&self, font_path: PathBuf) -> Vec<PathBuf> {
        let mut file_list: Vec<PathBuf> = Vec::new();
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
                                let font_format = format.to_str().unwrap();
                                if font_format_list.contains(&font_format) {
                                    file_list.push(dir_entry.path().to_path_buf());
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
