use std::sync::Arc;
use tokio::sync::mpsc::{self, Receiver, Sender};
use tokio::sync::Mutex;
use tokio::time::{sleep, Duration};

use ruspider::{Element, Ruspider};
use scraper::Html;

pub struct Crawl {
    infos: Option<Arc<Mutex<String>>>,
}

impl Crawl {
    pub fn new() -> Self {
        let e: Arc<Mutex<String>> = Arc::new(Mutex::new("".to_string()));
        Crawl { infos: Some(e) }
    }

    pub async fn start_jav_code(&self) {
        // let mut rx = r.sub();
        let (tx, mut rx) = mpsc::channel(10);
        let infos = self.infos.clone();
        let r = Ruspider::new(tx);

        tokio::spawn(async move {
            r.visit("https://wwww.baidu.com".to_string()).await;
            r.visit("https://wwww.baidu.com".to_string()).await;
        });

        while let Some(message) = rx.recv().await {
            println!("GOT = {:?}", message);
        }

        println!("infos: {:?}", self.infos)
    }
    pub fn start_jav_star_code() {}
    pub fn start_jav_code_form_dir() {}
    pub fn start_jav_prefiex_code() {}
}
