package database
import "github.com/api/common/config"

type DbConfig struct {
	Username     string `required:"true" split_words:"true"`
	Password string `required:"true" split_words:"true"`
	Host     string `required:"true" split_words:"true"`
	Port     string    `required:"true" split_words:"true"`
	Database     string `required:"true" split_words:"true"`
}



func NewEnvDBConfig() *DbConfig {
    return &DbConfig{
        Host:     config.Config("MYSQL_HOST"),
        Port:     config.Config("MYSQL_PORT"),
        Username: config.Config("MYSQL_USER"),
        Password:config.Config("MYSQL_PASSWORD"),
        Database: config.Config("MYSQL_DATABASE"),
    }
}


func (c *DbConfig) GetHost() string {
    return c.Host
}

func (c *DbConfig) GetPort() string {
    return c.Port
}

func (c *DbConfig) GetUsername() string {
    return c.Username
}

func (c *DbConfig) GetPassword() string {
    return c.Password
}

func (c *DbConfig) GetDatabase() string {
    return c.Database
}