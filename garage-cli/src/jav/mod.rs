use clap::ArgMatches;
mod db_bf;

pub async fn jav_sub_cmd(sub_cmd: &str, args: &ArgMatches) {
    match sub_cmd {
        "sync_db_bf" => {
            let action = db_bf::DbSyncBf::new(args).await;
            action.sync();
        }
        _ => println!("no complete sub command"),
    }
}
