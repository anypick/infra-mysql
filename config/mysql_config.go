package config

import (
	"fmt"
	"github.com/anypick/infra/base/props/container"
	"github.com/anypick/infra/utils/common"
	"reflect"
	"time"
)

const (
	DefaultPrefix = "mysql"
)


func init() {
	container.RegisterYamContainer(&MySqlConfig{Prefix: DefaultPrefix})
}

// mysql连接配置
type MySqlConfig struct {
	Prefix          string                                 //数据库配置信息的前缀，用于获取配置信息
	DriverName      string        `yaml:"driverName"`      // 驱动名称
	IpAddr          string        `yaml:"ipAddr"`          // ip地址
	Port            string        `yaml:"port"`            // 端口
	Username        string        `yaml:"username"`        // 用户名
	Password        string        `yaml:"password"`        // 密码
	Database        string        `yaml:"database"`        // 数据库名称
	MaxOpenConn     int           `yaml:"maxOpenConn"`     // 最大连接数
	MaxIdeConn      int           `yaml:"maxIdeConn"`      // 最大等待连接
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"` // 连接最大存活时间
}

func (m *MySqlConfig) ConfigAdd(config map[interface{}]interface{}) {
	m.DriverName = config["driverName"].(string)
	m.IpAddr = config["ipAddr"].(string)
	m.Port = fmt.Sprintf("%v", config["port"])
	m.Username = config["username"].(string)
	m.Password = config["password"].(string)
	m.Database = config["database"].(string)
	m.MaxOpenConn = config["maxOpenConn"].(int)
	m.MaxIdeConn = config["maxIdeConn"].(int)
	m.ConnMaxLifetime = time.Duration(config["connMaxLifetime"].(int))
}

func (m MySqlConfig) GetStringByDefault(fieldName, defaultValue string) string {
	stringValue := reflect.ValueOf(m).FieldByName(fieldName).Interface().(string)
	if common.StrIsBlank(stringValue) {
		return defaultValue
	}
	return stringValue
}

func (m MySqlConfig) GetIntByDefault(fieldName string, defaultValue int) int {
	intValue := reflect.ValueOf(m).FieldByName(fieldName).Interface().(int)
	if intValue == 0 {
		return defaultValue
	}
	return intValue
}

func (m MySqlConfig) GetDurationDefault(fieldName string, defaultValue time.Duration) time.Duration {
	durationValue := reflect.ValueOf(m).FieldByName(fieldName).Interface().(time.Duration)
	if durationValue == 0 {
		return time.Second * defaultValue
	}
	return time.Second * durationValue
}
