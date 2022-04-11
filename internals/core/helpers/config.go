package helpers

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
)

type ConfigStruct struct {
	AppName            string `mapstructure:"app_name"`
	ServiceAddress     string `mapstructure:"service_address"`
	ServicePort        string `mapstructure:"service_port"`
	ServiceMode        string `mapstructure:"service_mode"`
	DBType             string `mapstructure:"db_type"`
	DBHost             string `mapstructure:"mongo_db_host"`
	DBName             string `mapstructure:"mongo_db_name"`
	MongoUsername      string `mapstructure:"mongo_db_username"`
	MongoDBPassword    string `mapstructure:"mongo_db_password"`
	MongoDBPort        string `mapstructure:"mongo_db_port"`
	MongoDBAuth        string `mapstructure:"mongo_db_auth_db"`
	ServiceName        string `mapstructure:"service_name"`
	SandBoxURL         string `mapstructure:"sandbox_url"`
	PageLimit          string `mapstructure:"page_limit"`
	LogDir             string `mapstructure:"log_dir"`
	LogFile            string `mapstructure:"log_file"`
	ExternalConfigPath string `mapstructure:"external_config_path"`
}

var (
	service_address    string
	service_port       string
	mongodb_port       string
	dbtype             string
	service_mode       string
	mongo_DBHost       string
	dbName             string
	externalConfigPath string
)

func LoadConfig() (string, string, string, string, string, string, string, string) {
	flag.StringVar(&service_address, "service_address", Config.ServiceAddress, "local host")
	flag.StringVar(&service_port, "service_port", Config.ServicePort, "application ports")
	flag.StringVar(&dbtype, "dbtype", Config.DBType, "application db type")
	flag.StringVar(&mongodb_port, "mongodb_port", Config.MongoDBPort, "application ports")
	flag.StringVar(&service_mode, "service_mode", Config.ServiceMode, "application mode, either dev or production")
	flag.StringVar(&mongo_DBHost, "mongo_DBhost", Config.DBHost, "database host")
	flag.StringVar(&dbName, "dbname", Config.DBName, "database name")
	flag.StringVar(&externalConfigPath, "external_config_path", Config.DBName, "external config path")

	flag.Parse()
	for i, value := range flag.Args() {
		os.Args[i] = value
	}
	return service_address, service_port, dbtype, mongodb_port, service_mode, mongo_DBHost, dbName, externalConfigPath
}

func LoadEnv(path string) (config ConfigStruct, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("opay-wallet-engine")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return ConfigStruct{}, err
	}
	err = viper.Unmarshal(&config)
	return
}

func ReturnConfig() ConfigStruct {
	config, err := LoadEnv(".")
	if err != nil {
		log.Println(err)
	}
	if config.ExternalConfigPath != "" {
		viper.Reset()
		config, err = LoadEnv(config.ExternalConfigPath)
		if err != nil {
			log.Println(err)
		}

	}
	return config
}

var Config = ReturnConfig()
