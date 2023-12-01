use std::path::PathBuf;

#[derive(Debug, Clone)]
pub struct BatchffmpegOptions {
    pub input_path: PathBuf,
    pub input_format: String,
    pub output_format: String,
    pub output_path: PathBuf,
}

impl BatchffmpegOptions {
    pub fn new() -> Self {
        BatchffmpegOptions {
            input_path: PathBuf::new(),
            input_format: "".to_string(),
            output_path: PathBuf::new(),
            output_format: "".to_string(),
        }
    }

    pub fn input_path(mut self, path: PathBuf) -> Self {
        self.input_path = self.input_path.join(path);
        self
    }
}
