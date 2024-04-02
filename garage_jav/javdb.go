package garage_jav

const JAVDB_URL = "https://www.javdb.com"

type JavdbCrawl interface {
	GetJavbusMovieByHomePage() ([]JavMovie, error) // 通过首页爬取对应的电影信息
	GetJavbusMovie() ([]JavMovie, error)           // 通过番号爬取对应的电影信息
	GetJavbusMovieByPrefix() ([]JavMovie, error)   // 通过番号前缀爬取对应的电影信息
	GetJavbusMovieByStar() ([]JavMovie, error)     // 通过演员ID爬取对应的电影信息
	GetJavbusMovieByFilepath() ([]JavMovie, error) // 访问文件夹下的视频列表爬取电影信息
	SaveLocal(infos []JavMovie) error
}
