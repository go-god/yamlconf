package yamlconf

import (
	"log"
	"testing"
	"time"
)

type dataConf struct {
	RedisConf redisConf
	Ip        []string
}

// redisConf redis连接信息,仅作为单元测试使用
type redisConf struct {
	Host           string
	Port           int
	Password       string
	Database       int
	MaxIdle        int // 空闲pool个数
	MaxActive      int // 最大激活数量
	ConnectTimeout int // 连接超时，单位s
	ReadTimeout    int // 读取超时
	WriteTimeout   int // 写入超时

	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout int // 空闲连接超时,单位s

	// Close connections older than this duration. If the value is zero, then
	// the pool does not close connections based on age.
	MaxConnLifetime int // 连接最大生命周期,单位s，默认1800s
}

func TestYaml(t *testing.T) {
	conf := NewConf()
	err := conf.LoadConf("test.yaml")
	log.Println(conf.GetData(), err)

	data := conf.GetData()

	var graceful time.Duration
	conf.Get("GracefulWait", &graceful)
	log.Println("graceful: ", graceful)

	log.Println("RedisCommon: ", data["RedisCommon"])

	// 读取数据到结构体中
	var redisConf = &dataConf{}
	conf.GetStruct("RedisCommon", redisConf)
	log.Println(redisConf)
	log.Println("Ip:", redisConf.Ip)
	log.Println(redisConf.RedisConf.Password == "")

}

/**
go test -v
=== RUN   TestYaml
2021/05/01 21:38:39 map[] <nil>
2021/05/01 21:38:39 graceful:  5s
2021/05/01 21:38:39 RedisCommon:  <nil>
2021/05/01 21:38:39 &{{127.0.0.1 6379  0 3 10 0 0 0 120 0} [11.12.1.1 11.12.1.2 11.12.1.3]}
2021/05/01 21:38:39 Ip: [11.12.1.1 11.12.1.2 11.12.1.3]
2021/05/01 21:38:39 true
--- PASS: TestYaml (0.00s)
PASS
ok  	github.com/go-god/yamlconf	0.018s
*/
