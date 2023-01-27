package env

type system struct {
	Debug     bool   `json:"debug"`
	TracerDSN string `json:"tracer_dsn"`
}

type database struct {
	Type        string
	Host        string
	Port        int
	Database    string
	User        string
	Password    string
	Charset     string
	SSLMode     string
	DBFile      string
	TablePrefix string
}

type telegram struct {
	Token string
}

type redis struct {
	Network  string
	Host     string
	Port     int
	Password string
	DB       int
}

type Config struct {
	System   *system
	Database *database
	Telegram *telegram
	Redis    *redis
}
