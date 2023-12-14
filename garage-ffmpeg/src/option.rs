use std::path::PathBuf;

#[derive(Debug, Clone)]
pub struct BatchffmpegOptions {
    pub input_path: PathBuf,
    pub input_format: String,
    pub font_path: PathBuf,
    pub sub_suffix: String,
    pub sub_number: u32,
    pub output_format: String,
    pub output_path: PathBuf,
    pub advance: String,
    pub exec: bool,
}

impl BatchffmpegOptions {
    pub fn new() -> Self {
        BatchffmpegOptions {
            input_path: PathBuf::new(),
            input_format: "".to_string(),
            font_path: PathBuf::new(),
            sub_suffix: "".to_string(),
            sub_number: 0,
            output_path: PathBuf::new(),
            output_format: "".to_string(),
            advance: "".to_string(),
            exec: false,
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

    pub fn font_path(mut self, font_path: PathBuf) -> Self {
        self.font_path = font_path;
        self
    }

    pub fn sub_suffix(mut self, suffix: String) -> Self {
        self.sub_suffix = suffix;
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

    pub fn advance(mut self, advance: Option<&String>) -> Self {
        match advance {
            Some(a) => self.advance = a.to_owned(),
            None => {}
        }
        self
    }

    pub fn exec(mut self, exec: bool) -> Self {
        self.exec = exec;
        self
    }
}
