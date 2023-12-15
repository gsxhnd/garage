use std::f32::consts::E;

use ruspider::{Element, Ruspider};

pub struct Crawl {
    crawl_client: Ruspider,
    data: String,
}

impl Crawl {
    pub fn new() -> Self {
        let r = Ruspider::new();
        Crawl {
            crawl_client: r,
            data: "1`23".to_string(),
        }
    }

    pub fn start(self) {
        // let a = Data::new();
        self.crawl_client.on_html("", self.test_input())
    }

    pub fn start_b(&self) {}

    pub fn test_input(self) -> impl Fn(Element) {
        let d = self.data;
        return move |e: Element| {
            println!("{}", d);
        };
    }

    // pub fn test_input(data: String) -> Fn(Element) {
    //     return |e: Element| {
    //         println!("data")
    //     }
    // }
}

// pub struct Data {
//     title: String,
// }

// impl Data {
//     pub fn new() -> Self {
//         Data {
//             title: "".to_string(),
//         }
//     }
// }

// pub fn star(e: Element) {
//     println!("e: {:?}", e);
//     self.title = e.text
// }

// pub fn test_input(data: String) -> Fn(Element) {
//     return |e: Element| println!("data");
// }
