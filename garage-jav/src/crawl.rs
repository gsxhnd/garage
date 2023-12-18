pub struct Crawl {}

impl Crawl {
    pub fn new() -> Self {
        let req_client = reqwest::Client::builder()
            // .proxy(self.proxy.clone())
            .build()
            .unwrap();
        Crawl {}
    }

    pub fn start_jav_code() {}
    pub fn start_jav_star_code() {}
    pub fn start_jav_code_form_dir() {}
    pub fn start_jav_prefiex_code() {}
}
