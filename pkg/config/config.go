package config

type Conf struct {
	Server     string `json:"server"`
	ServerPort int    `json:"port"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
}
