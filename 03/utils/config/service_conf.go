package config

type ServiceConfig struct {
	ServicePort      int  `mapstructure:"SERVICE_PORT"`
	EnableMigrations bool `mapstructure:"ENABLE_MIGRATION"`

	DBHost string `mapstructure:"POSTGRES_HOST"`
	DBPort int    `mapstructure:"POSTGRES_PORT"`
	DBUser string `mapstructure:"POSTGRES_USER"`
	DBPwd  string `mapstructure:"POSTGRES_PASSWORD"`
	DBName string `mapstructure:"POSTGRES_DB"`

	Environment string `mapstructure:"ENVIRONMENT"`
	GroqAPIKey  string `mapstructure:"GROQ_API_KEY"`
	GroqUrl     string `mapstructure:"GROQ_URL"`
}
