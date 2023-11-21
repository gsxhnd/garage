use clap::ArgMatches;
use sea_orm::{ConnectOptions, Database, DatabaseConnection};
use std::time::Duration;

use std::path::PathBuf;

pub struct DbSyncBf {
    csv_path: String,
    db_path: String,
    db_conn: DatabaseConnection,
}

impl DbSyncBf {
    pub async fn new(args: &ArgMatches) -> Self {
        let csv_path = args
            .get_one::<PathBuf>("from_csv")
            .expect("csv path not set")
            .to_str()
            .unwrap()
            .to_owned();
        let db_path = args
            .get_one::<PathBuf>("db_path")
            .expect("db path not set")
            .to_str()
            .unwrap()
            .to_owned();

        let mut opt = ConnectOptions::new("sqlite://data/test.db?mode=rw");
        opt.max_connections(100)
            .min_connections(5)
            .connect_timeout(Duration::from_secs(8))
            .idle_timeout(Duration::from_secs(8))
            .sqlx_logging(true);
        let db_conn = Database::connect(opt)
            .await
            .expect("Database connection failed");

        DbSyncBf {
            csv_path,
            db_path,
            db_conn,
        }
    }

    pub fn sync(&self) {
        println!("jav_sync_db_bf, csv: {:?}", self.csv_path);
        println!("jav_sync_db_bf, db: {:?}", self.db_path);
    }
}
