use serde::{Deserialize, Serialize};

#[derive(Debug, PartialEq, Serialize, Deserialize)]
pub struct Config {
    log_url: String,
    day_data_url: String,
    year_data_url: String,
    collect: CollectConfig,
    dest_path: DestPathConfig,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
pub struct CollectConfig {
    whole_year_start: i32,
    whole_year_end: i32,
    day_start: String,
    day_end: String,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
pub struct DestPathConfig {
    raw: String,
    catalog: String,
    log: String,
}
