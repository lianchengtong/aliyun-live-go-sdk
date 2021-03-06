package device

import (
	"errors"
	"github.com/BPing/aliyun-live-go-sdk/aliyun"
	"github.com/BPing/aliyun-live-go-sdk/device/cdn"
	"github.com/BPing/aliyun-live-go-sdk/device/live"
)

type DevType string

const (
	CdnDevice  = DevType("cdn")
	LiveDevice = DevType("live")
)

// 初始配置项
type Config struct {
	Credentials *aliyun.Credentials
	StreamCert  *live.StreamCredentials
	DomainName  string
	AppName     string
}

// 生产实例（工厂模式）
func GetDevice(devType DevType, config Config) (instance interface{}, err error) {
	if config.Credentials == nil {
		err = errors.New("Credentials should be nil ")
		return
	}
	switch devType {
	case CdnDevice:
		instance = cdn.NewCDN(config.Credentials)
	case LiveDevice:
		if "" == config.DomainName || "" == config.AppName {
			err = errors.New("live dev: domainname and appname should not be empty ")
			return
		}
		instance = live.NewLive(config.Credentials, config.DomainName, config.AppName, config.StreamCert)
	default:
		instance = nil
	}
	return
}
