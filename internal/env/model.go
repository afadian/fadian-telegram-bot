package env

type system struct {
	Listen    string
	Debug     bool
	TracerDSN string
	Secret    string
	CertFile  string
	KeyFile   string
	PublicURL string
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
	URL           string // default is https://api.telegram.org
	Token         string
	Updates       int
	PollerTimeout int // in seconds
	Offline       bool
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
