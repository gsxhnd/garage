use tracing_appender::rolling;
use tracing_subscriber::{self, filter, fmt, layer::SubscriberExt, Layer};

pub struct Logger {
    log_path: Option<String>,
    show_file: bool,
    show_thread_names: bool,
    log_level: filter::LevelFilter,
}

impl Logger {
    pub fn new() -> Self {
        Logger {
            log_path: None,
            show_file: false,
            show_thread_names: false,
            log_level: filter::LevelFilter::TRACE,
        }
    }

    pub fn set_log_path(mut self, log_path: String) -> Self {
        self.log_path = Some(log_path);
        self
    }

    pub fn set_show_file(mut self, show: bool) -> Self {
        self.show_file = show;
        self
    }

    pub fn set_show_thread_name(mut self, show: bool) -> Self {
        self.show_thread_names = show;
        self
    }

    pub fn set_log_level(mut self, lvl: filter::LevelFilter) -> Self {
        self.log_level = lvl;
        self
    }

    pub fn init(self) {
        match self.log_path {
            Some(path) => {
                let debug_file = rolling::minutely(path, "debug");
                let a = fmt::Layer::new()
                    .json()
                    .with_file(self.show_file)
                    .with_thread_names(self.show_thread_names)
                    .with_writer(debug_file)
                    .with_filter(self.log_level);
                let subscriber = tracing_subscriber::Registry::default().with(a);
                tracing::subscriber::set_global_default(subscriber)
                    .expect("unable to set global subscriber");
            }
            None => {
                let a = fmt::Layer::new()
                    .with_file(self.show_file)
                    .with_thread_names(self.show_thread_names)
                    .with_writer(std::io::stdout)
                    .with_filter(self.log_level);
                let subscriber = tracing_subscriber::Registry::default().with(a);
                tracing::subscriber::set_global_default(subscriber)
                    .expect("unable to set global subscriber");
            }
        }
    }
}
