use crate::spider::Ruspider;
use reqwest::{Client, Method, Request, RequestBuilder, Response};
use std::sync::Arc;
use std::thread::sleep;
use tokio::sync::mpsc::{self, Receiver, Sender};
use tokio::time::{Duration, Sleep};

#[derive(Debug, Clone)]
pub struct Queue {
    client: Ruspider,
    thread: i32,
    urls: Vec<String>,
}

impl Queue {
    pub fn new(client: Ruspider) -> Self {
        Queue {
            client,
            thread: 1,
            urls: Vec::new(),
        }
    }

    pub fn thread(mut self, thread: i32) -> Self {
        self.thread = thread;
        self
    }

    pub fn add_url(&mut self, url: String) {
        self.urls.push(url);
    }

    pub async fn run(&mut self) {
        for u in self.urls.clone() {
            let c = self.client.clone();
            tokio::spawn(async move {
                c.visit(u).await;
            });
        }
    }
}
