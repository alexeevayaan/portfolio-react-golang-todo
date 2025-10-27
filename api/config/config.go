package config


type Config struct{

}

func New() (Config, error){
	var config Config
	return config, nil
}