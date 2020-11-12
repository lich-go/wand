package config

import (
	"github.com/unknwon/goconfig"
	"log"
	"os"
	"strings"
)

var defaultPath = "config" // 配置文件路径
var defaultName = "config" // 默认配置文件名称
var defaultExt = ".ini"    // 配置文件扩展名

type Config struct {
	Path     string
	Filename string
	Ext      string
}

type File = goconfig.ConfigFile

// LoadConfigFile 加载配置文件
func (ts *Config) LoadConfigFile(f ...func(*File)) {
	var err error // error handle
	var file *File
	var filepath string

	if ts.Path == "" {
		ts.Path = defaultPath
	}

	if ts.Ext == "" {
		ts.Ext = defaultExt
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir = strings.Replace(dir, "\\", "/", -1) + "/" + ts.Path + "/"

	// 判断默认配置文件是否存在
	defaultFilepath := dir + defaultName + ts.Ext // 构造默认配置文件
	_, err = os.Stat(defaultFilepath)             // 判断是否存在自定义配置文件
	if err != nil {
		defaultConfig()
	}

	// 使用自定义配置文件
	if ts.Filename != "" {
		filepath = dir + "/" + ts.Filename + ts.Ext
		_, err = os.Stat(filepath) // 判断是否存在自定义配置文件
		if err == nil {
			file, err = goconfig.LoadConfigFile(defaultFilepath, filepath)
		}
	}

	// 未指定自定义配置文件或者自定义配置文件加载失败
	if ts.Filename == "" || err != nil {
		file, err = goconfig.LoadConfigFile(defaultFilepath)
	}

	// 配置文件加载失败
	if err != nil {
		log.Fatal(err)
	}

	// 执行回调
	if len(f) > 0 && f[0] != nil {
		f[0](file)
	}
}

// defaultConfig
// 默认配置文件不存在时进行的处理
func defaultConfig() {
	log.Fatal("The default config file is not exist: " + defaultName + defaultExt)
}
