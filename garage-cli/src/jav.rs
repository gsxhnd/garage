use clap::ArgMatches;
use garage_jav::DbSyncBf;
use std::path::PathBuf;

pub async fn jav_sub_cmd(sub_cmd: &str, args: &ArgMatches) {
    match sub_cmd {
        "sync_db_bf" => {
            let _csv_path = args
                .get_one::<PathBuf>("from_csv")
                .expect("csv path not set")
                .to_str()
                .unwrap()
                .to_owned();
            let _db_path = args
                .get_one::<PathBuf>("db_path")
                .expect("db path not set")
                .to_str()
                .unwrap()
                .to_owned();
            let action = DbSyncBf::new().await;
            action.sync();
        }
        _ => println!("no complete sub command"),
    }
}
