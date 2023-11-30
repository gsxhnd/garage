mod args;
use std::process::Command;
use std::process::Stdio;

pub struct Batchffmpeg {}

impl Batchffmpeg {
    pub fn new() -> Self {
        Batchffmpeg {}
    }

    pub fn set_dest_path(&self) {}

    pub fn convert(&self) {
        let mut cmd = Command::new("wget")
            .arg("https://www.baidu.com")
            .stdout(Stdio::inherit())
            .spawn()
            .expect("fail to execute");
        cmd.wait().unwrap();
    }
}
