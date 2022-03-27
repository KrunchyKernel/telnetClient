package models

type Config struct {
	Port     int    `json:"port"`
	Target   string `json:"target"`
	UserName string `json:"userName"`
}

func (c Config) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}
