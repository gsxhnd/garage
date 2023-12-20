use crate::element::Element;
use reqwest::{Client, Proxy, Request, Response};
use scraper::Html;
use tokio::sync::mpsc::{self, Receiver, Sender};

#[derive(Debug)]
pub struct Ruspider {
    req_client: Client,
    proxy: Option<reqwest::Proxy>,
    tx: Sender<Element>,
}

impl Ruspider {
    pub fn new(tx: Sender<Element>) -> Self {
        // let (tx, rx) = mpsc::channel(10);
        let req_client = reqwest::Client::builder()
            // .proxy(self.proxy.clone())
            .build()
            .unwrap();
        Ruspider {
            req_client,
            proxy: None,
            // document: None,
            tx,
        }
    }

    pub fn proxy(&mut self, proxy_scheme: &str) {
        self.proxy = Some(reqwest::Proxy::all(proxy_scheme).unwrap())
    }

    pub fn on_request() {}

    pub fn on_response() {}

    pub fn sub(&self) {
        // self.tx
        // self.rx
    }

    pub async fn visit(&self, url: String) {
        let req = self.req_client.request(reqwest::Method::GET, url.clone());
        let resp = req.send().await.unwrap();
        match resp.text().await {
            Ok(s) => {
                let e = Element::new().set_link(url.clone()).parse(s).build();
                self.tx.send(e).await.unwrap();
            }
            Err(_) => {}
        }
        // println!("response: {:?}", resp.text().await)
    }
    pub fn stop(&self) {}
}
