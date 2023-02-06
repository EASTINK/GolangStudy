package config

import (
	"encoding/json"
	"io"
	"log"
	"openwechat"
	"os"
	"sync"
)

// Configuration 项目配置
type Configuration struct {
	// gtp apikey
	ApiKey string          `json:"api_key"`
	Npush  map[string]bool `json:"no_morning"`
}

//开启一个推送群
func (s *Configuration) Open_push(gpname string) string {
	for k, v := range s.Npush {
		if k == gpname {
			if v == false {
				return "本群已开启了消息订阅,你有可能是在找 #T : 关闭本群订阅"
			}
		}
	}
	s.Npush[gpname] = false //no-message:flase
	return "设置完成"
}

//关闭一个推送群
func (t *Configuration) Close_push(gpname string) string {
	for k, v := range t.Npush {
		if k == gpname {
			if v == true {
				return "本群已关闭了消息订阅,你有可能是在找 #S : 开启本群订阅"
			}
		}
	}
	t.Npush[gpname] = true //no-message:true
	return "设置完成"
}

// LoadConfig 加载配置
func Load(target string) *Configuration {
	var once sync.Once
	var config *Configuration
	once.Do(func() {
		// 从文件中读取
		config = &Configuration{}
		f, err := os.Open(target)
		if err != nil {
			log.Fatalf("open config err: %v", err)
			return
		}
		defer f.Close()
		encoder := json.NewDecoder(f)
		err = encoder.Decode(config)
		if err != nil {
			log.Fatalf("decode config err: %v", err)
			return
		}
	})
	return config
}

//写入config.json
func (config *Configuration) WriteConfig(target string) {
	var once sync.Once
	once.Do(func() {
		var (
			f       *os.File
			err     error
			strpush []byte
		)

		f, err = os.OpenFile(target, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
		if err != nil {
			log.Fatalln("打开文件错误", err)
			return
		}

		defer f.Close()

		strpush, err = json.Marshal(config)
		if err != nil {
			log.Fatalf("encode config err: %v", err)
		}
		_, err = io.WriteString(f, string(strpush))
		if err != nil {
			log.Fatalln("写入错误", err)
		}

	})
}

//登录时将所有群组序列存到本地
func ReLocalGroup(s *openwechat.Self) {
	var once sync.Once
	groups, _ := s.Groups()
	Npush := map[string]bool{}
	//记录所有目标群到本地
	once.Do(func() {
		var (
			f     *os.File
			err   error
			spush []byte
		)

		for x := range groups {
			Npush[groups[x].UserName] = false
		}

		f, err = os.OpenFile("group.json", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
		if err != nil {
			log.Fatalln("打开文件错误", err)
			return
		}

		defer f.Close()
		spush, err = json.Marshal(Npush)
		if err != nil {
			log.Fatalf("encode config err: %v", err)
		}
		_, err = io.WriteString(f, string(spush))
		if err != nil {
			log.Fatalln("写入错误", err)
		}
	})

	//go plug.Cron(groups, s)
}

func InitLocalGroups(s *openwechat.Self) {
	//保存群名称到目录下 group.json
	//ReLocalGroup(s)
}
