package dtype

type AppConfig struct {
	AppName     string `koanf:"app.name"`
	Environment string `koanf:"app.env"`

	DBUsername string `koanf:"db.username"`
	DBPassword string `koanf:"db.password"`
	DBHost     string `koanf:"db.host"`
	DBPort     int    `koanf:"db.port"`
	DBName     string `koanf:"db.name"`
	DBMaxPools int    `koanf:"db.max_pools"`
	DBMaxIdle  int    `koanf:"db.max_idle"`

	YoutubeApiKeys []string `koanf:"youtube.api_keys"`
}
