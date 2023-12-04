use std::path::PathBuf;

#[derive(Debug, Clone)]
pub struct BatchffmpegOptions {
    pub input_path: PathBuf,
    pub input_format: String,
    pub output_format: String,
    pub output_path: PathBuf,
    pub advance: String,
}

impl BatchffmpegOptions {
    pub fn new() -> Self {
        BatchffmpegOptions {
            input_path: PathBuf::new(),
            input_format: "".to_string(),
            output_path: PathBuf::new(),
            output_format: "".to_string(),
            advance: "".to_string(),
        }
    }

    pub fn input_path(mut self, path: PathBuf) -> Self {
        self.input_path = self.input_path.join(path);
        self
    }

    pub fn input_format(mut self, format: String) -> Self {
        self.input_format = format;
        self
    }

    pub fn output_path(mut self, path: PathBuf) -> Self {
        self.output_path = self.output_path.join(path);
        self
    }

    pub fn output_format(mut self, format: String) -> Self {
        self.output_format = format;
        self
    }

    pub fn advance(mut self, advance: String) -> Self {
        self.advance = advance;
        self
    }
}
