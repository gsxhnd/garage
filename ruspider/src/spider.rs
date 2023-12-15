use crate::element::Element;
use reqwest;
use scraper::Html;

pub struct Ruspider {
    proxy: Option<reqwest::Proxy>,
    document: Option<Html>,
}

impl Ruspider {
    pub fn new() -> Self {
        Ruspider {
            proxy: None,
            document: None,
        }
    }

    pub fn proxy(&mut self, proxy_scheme: &str) {
        self.proxy = Some(reqwest::Proxy::all(proxy_scheme).unwrap())
    }

    pub fn on_request() {}

    pub fn on_response() {}

    pub fn on_html(&self, query_selector: &str) -> Element {
        Element::new()
    }

    pub async fn visit(&self, url: &str) {
        let req_client = reqwest::Client::builder()
            // .proxy(self.proxy.clone())
            .build()
            .unwrap();
        let req = req_client.request(reqwest::Method::GET, url);
        let resp = req.send().await.unwrap();
        println!("response: {:?}", resp.text().await)
    }
}
