package config

import(
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var(
	Port string
	User string
	Password string
	BaseName string
	
	config *Configuration
)

type Configuration struct {
	Port string `json:"Port"`
	User string `json:"User"`
	Password string `json:"Password"`
	BaseName string `json:"BaseName"`
}

func LoadConfiguration() error {
	log.Println("Reading from config file...")
	
	configFile, err := ioutil.ReadFile("./config.json")

	
	if err != nil {
		log.Println(err.Error())
		
		configFile, err := os.Create("./config.json")
		
		if err != nil {
			log.Println(err.Error())
			return err
		}
		
		_, err = configFile.WriteString(
	`{
	"Port": "",
	"User": "",
	"Password": "",
	"BaseName": ""
}`)
		
		if err != nil {
			log.Println(err.Error())
			return err
		}
		
		err = configFile.Sync()
		
		if err != nil {
			log.Println(err.Error())
			return err
		}
		
		defer configFile.Close()
		
		log.Println("Done")
		
		return nil
	}
	
	log.Println(string(configFile))
	
	err = json.Unmarshal(configFile, &config)
	
	if err != nil {
		log.Println(err.Error())
		return err
	}
	
	Port = config.Port
	User = config.User
	Password = config.Password
	BaseName = config.BaseName
	
	return nil
}