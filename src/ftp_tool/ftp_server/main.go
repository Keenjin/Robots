package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
	"github.com/keenjin/gomini/kfile"
	"github.com/keenjin/gomini/klog"
	"os"
)

type ftpConfig struct {
	Server serverInfo `toml:"server"`
	Auth   authInfo   `toml:"auth"`
}

type serverInfo struct {
	IP   string
	Port int
	Root string
}

type authInfo struct {
	User   string `toml:"user"`
	Passwd string `toml:"passwd"`
}

type kLogger struct{}

func (logger *kLogger) Print(sessionId string, message interface{}) {
	klog.Info("%s  %s", sessionId, message)
}

func (logger *kLogger) Printf(sessionId string, format string, v ...interface{}) {
	logger.Print(sessionId, fmt.Sprintf(format, v...))
}

func (logger *kLogger) PrintCommand(sessionId string, command string, params string) {
	if command == "PASS" {
		klog.Info("%s > PASS ****", sessionId)
	} else {
		klog.Info("%s > %s %s", sessionId, command, params)
	}
}

func (logger *kLogger) PrintResponse(sessionId string, code int, message string) {
	klog.Info("%s < %d %s", sessionId, code, message)
}

func main() {
	var config ftpConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		klog.Error("cfg: %s can't decoded. err: %v\n", "config.toml", err)
		return
	}

	if !kfile.IsPathExist(config.Server.Root) {
		os.Mkdir(config.Server.Root, os.ModePerm)
	}

	opt := &server.ServerOpts{
		Name: "Keen FtpServer",
		Factory: &filedriver.FileDriverFactory{
			RootPath: config.Server.Root,
			Perm:     server.NewSimplePerm("root", "root"),
		},
		Port: config.Server.Port,
		Auth: &server.SimpleAuth{
			Name:     config.Auth.User,
			Password: config.Auth.Passwd,
		},
		Logger: &kLogger{},
	}

	ftpServer := server.NewServer(opt)
	err := ftpServer.ListenAndServe()
	if err != nil {
		klog.Error("error starting server:", err)
	}
}
