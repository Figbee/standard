package global

import "go.uber.org/zap/zapcore"

type Configuration struct {
	System SystemConfiguration `json:"system"  mapstructure:"system"`
	Logs   LogsConfiguration   `json:"logs"  mapstructure:"logs"`
	Mysql  MysqlConfiguration  `json:"mysql" mapstructure:"mysql"`
}

type SystemConfiguration struct {
	Host        string `json:"host" mapstructure:"host"`
	Port        int    `json:"port" mapstructure:"port"`
	IdentityKey string `json:"identityKey" mapstructure:"identity-key"`
}

type LogsConfiguration struct {
	Level      zapcore.Level `json:"level" mapstructure:"level" mapstructure:"level"`
	Path       string        `json:"path" mapstructure:"path" mapstructure:"path"`
	MaxSize    int           `json:"max_size" mapstructure:"max-size" mapstructure:"max-size"`
	MaxBackups int           `json:"maxBackups" mapstructure:"max-backups"`
	MaxAge     int           `json:"maxAge" mapstructure:"max-age"`
	Compress   bool          `json:"compress" mapstructure:"compress"`
}

type MysqlConfiguration struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Query    string `json:"query" mapstructure:"query"`
	LogMode  bool   `json:"logMode" mapstructure:"log-mode"`
}
