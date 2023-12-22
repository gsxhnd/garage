use std::sync::Arc;
use tokio::sync::mpsc::{self, Receiver, Sender};
use tokio::sync::Mutex;
use tokio::time::{sleep, Duration};

use ruspider::{Element, Queue, Ruspider};
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
        let (tx, mut rx) = mpsc::channel(10);
        let r = Ruspider::new().set_sender(tx).build();
        tokio::spawn(async move {
            r.visit("https://www.baidu.com".to_string()).await;
        });
        while let Some(message) = rx.recv().await {
            println!("GOT = {:?}", message);
        }
    }

    pub async fn start_jav_star_code(&self) {
        let (tx, mut rx) = mpsc::channel(10);
        let mut queue = Queue::new(Ruspider::new().set_sender(tx).build());

        tokio::spawn(async move {
            queue.add_url("https://www.baidu.com".to_string());
            queue.add_url("https://www.baidu.com".to_string());
            queue.add_url("https://www.baidu.com".to_string());
            queue.add_url("https://www.baidu.com".to_string());
            queue.add_url("https://www.baidu.com".to_string());
            queue.add_url("https://www.baidu.com".to_string());
            queue.add_url("https://www.baidu.com".to_string());
            queue.run().await;
        });

        while let Some(message) = rx.recv().await {
            println!("GOT = {:?}", message);
        }

        println!("infos: {:?}", self.infos)
    }

    pub fn start_jav_code_form_dir() {}

    pub fn start_jav_prefiex_code() {}
}
