mod cmd;
mod db_bf;

#[tokio::main]
async fn main() {
    let cmd = cmd::new_cmd();

    match cmd.get_matches().subcommand() {
        Some(("jav_sync_db_bf", sub_m)) => {
            let action = db_bf::DbSyncBf::new(sub_m).await;
            action.sync();
        }
        _ => {}
    }
}
