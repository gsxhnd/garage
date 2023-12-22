use crate::element::Element;
use reqwest::{Client, Proxy, Request, Response};
use scraper::Html;
use tokio::sync::mpsc::{self, Receiver, Sender};

#[derive(Debug, Clone)]
pub struct Ruspider {
    pub(crate) req_client: Client,
    tx: Option<Sender<Element>>,
}

pub struct RuspiderBuilder {
    proxy: Option<Proxy>,
    tx: Option<Sender<Element>>,
}

impl Ruspider {
    pub fn new() -> RuspiderBuilder {
        RuspiderBuilder {
            proxy: None,
            tx: None,
        }
    }

    pub async fn visit(&self, url: String) {
        let req = self.req_client.request(reqwest::Method::GET, url.clone());
        let resp = req.send().await.unwrap();
        let element = match resp.text().await {
            Ok(s) => {
                let e = Element::new().set_link(url.clone()).set_content(s).build();
                e
            }
            Err(_) => {
                let e = Element::new().build();
                e
            }
        };

        match &self.tx {
            Some(tx) => {
                tx.send(element).await.unwrap();
            }
            None => {}
        }
    }

    pub fn stop(&self) {}
}

impl RuspiderBuilder {
    pub fn set_sender(mut self, tx: Sender<Element>) -> Self {
        self.tx = Some(tx);
        self
    }

    pub fn set_proxy(mut self, proxy_scheme: &str) -> Self {
        self.proxy = Some(reqwest::Proxy::all(proxy_scheme).unwrap());
        self
    }

    pub fn build(self) -> Ruspider {
        let req_client_builder = reqwest::Client::builder();
        let req_client = match self.proxy {
            Some(proxy) => req_client_builder.proxy(proxy).build(),
            None => req_client_builder.build(),
        }
        .unwrap();

        Ruspider {
            req_client,
            tx: self.tx,
        }
    }
}
