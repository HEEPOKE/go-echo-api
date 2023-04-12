package models

type ConfigDB struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type MainConfig struct {
	PORT         string
	ENDPOINT_URL string
	SECRET_KEY   string
}
