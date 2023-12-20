#[derive(Debug, Clone)]
pub struct Element {
    pub link: String,
    pub html: String,
}

pub struct ElementBuilder {
    link: Option<String>,
    html: Option<String>,
}

impl Element {
    pub fn new() -> ElementBuilder {
        ElementBuilder {
            link: None,
            html: None,
        }
    }
}

impl ElementBuilder {
    pub fn set_link(mut self, link: String) -> Self {
        self.link = Some(link);
        self
    }

    pub fn parse(mut self, content: String) -> Self {
        self.html = Some(content);
        self
    }

    pub fn build(self) -> Element {
        let link = self.link.expect("error no link");
        let html = self.html.unwrap();
        Element { link, html }
    }
}
