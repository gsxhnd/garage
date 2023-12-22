pub struct JavMovie {
    pub code: String,
    pub title: String,
    pub cover: String,
    pub publishdate: String,
    pub length: String,
    pub director: String,
    pub producecompany: String,
    pub publishcompany: String,
    pub series: String,
    pub stars: Vec<String>,
}

pub struct JavMovieMagnet {
    pub name: String,
    pub link: String,
    pub size: f32,
    pub subtitle: bool,
    pub hd: bool,
}
