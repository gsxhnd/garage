use sea_orm::{ConnectOptions, Database, DatabaseConnection};
use std::time::Duration;

pub struct DbSyncBf {
    csv_path: String,
    db_path: String,
    db_conn: DatabaseConnection,
}

impl DbSyncBf {
    pub async fn new() -> Self {
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
            csv_path: "".to_string(),
            db_path: "".to_string(),
            db_conn,
        }
    }

    pub fn sync_star_tag_from_csv(&self) {
        println!("jav_sync_db_bf, csv: {:?}", self.csv_path);
        println!("jav_sync_db_bf, db: {:?}", self.db_path);
    }

    pub fn sync_star_tag_from_db(&self) {
        println!("jav_sync_db_bf, csv: {:?}", self.csv_path);
        println!("jav_sync_db_bf, db: {:?}", self.db_path);
    }
}
