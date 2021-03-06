# aliyun-live-go-sdk
[![Build Status](https://travis-ci.org/BPing/aliyun-live-go-sdk.svg?branch=master)](https://travis-ci.org/BPing/aliyun-live-go-sdk) [![Coverage Status](https://coveralls.io/repos/github/BPing/aliyun-live-go-sdk/badge.svg?branch=master)](https://coveralls.io/github/BPing/aliyun-live-go-sdk?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/BPing/aliyun-live-go-sdk)](https://goreportcard.com/report/github.com/BPing/aliyun-live-go-sdk)
   
    阿里云直播 golang SDK

[阿里云视频直播接口文档](https://help.aliyun.com/product/29949.html?spm=5176.750001.2.12.Xa4wT1)

# 快速开始

```go
package main

import (
    "github.com/BPing/aliyun-live-go-sdk/aliyun"
    "github.com/BPing/aliyun-live-go-sdk/device/live"
    "github.com/BPing/aliyun-live-go-sdk/util"
    "time"
    "fmt"
)

const (
    AccessKeyId = "<Yours' Id>"
    AccessKeySecret = "<...>"
)

func main() {
    cert := aliyun.NewCredentials(AccessKeyId, AccessKeySecret)
    live := live.NewLive(cert, "<Yours' CDN>", "app-name",nil).SetDebug(true)
    resp := make(map[string]interface{})
    live.StreamsPublishList(time.Now().Add(-time.Hour * 12), time.Now(), &resp)
    fmt.Println(resp)
}
```
 
# 构建安装

go get:

```sh
go get github.com/BPing/aliyun-live-go-sdk
```

# 文档
* Sdk:https://godoc.org/github.com/BPing/aliyun-live-go-sdk [![GoDoc](https://godoc.org/github.com/BPing/aliyun-live-go-sdk?status.svg)](https://godoc.org/github.com/BPing/aliyun-live-go-sdk)
* CDN:https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/cdn   [![GoDoc](https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/cdn?status.svg)](https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/cdn)
* Live:https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/live [![GoDoc](https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/live?status.svg)](https://godoc.org/github.com/BPing/aliyun-live-go-sdk/device/live)

# Example
* [example](https://github.com/BPing/aliyun-live-go-sdk/tree/master/example)

## 直播(Live)

       （v0.5,v0.5-）
            方法名以"WithApp"结尾代表可以更改请求中  "应用名字（AppName）"，
            否则按默认的(初始化时设置的AppName)。
            如果为空，代表忽略参数AppName
       （v0.6+）
            移除以"WithApp"为后缀方法。

```go
    cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
    liveM := live.NewLive(cert, DomainName, AppName, nil).SetDebug(true)
```

* 获取流列表
```go
    resp := make(map[string]interface{})
    liveM.StreamsPublishList(time.Now().Add(-time.Hour * 12), time.Now(), &resp)
    fmt.Println(resp)
    // @appname 应用名 为空时，忽略此参数
    resp := make(map[string]interface{})
    liveM.StreamsPublishListWithApp(AppName,time.Now().Add(-time.Hour * 12), time.Now(), &resp)
    fmt.Println(resp)
```

* 获取黑名单
```go
    resp = make(map[string]interface{})
    err = liveM.StreamsBlockList(&resp)
    fmt.Println(err, resp)
```

* 获取流的在线人数
```go
    resp1 := live.OnlineInfoResponse{}
    err := liveM.StreamOnlineUserNum("video-name", &resp1)
    fmt.Println(err, resp1)  
   // @appname 应用名 为空时，忽略此参数   （v0.5,v0.5-）
    resp1 := live.OnlineInfoResponse{}
    err := liveM.StreamOnlineUserNumWithApp(AppName,"video-name", &resp1)
    fmt.Println(err, resp1)  
```

* 获取控制历史
```go
    resp = make(map[string]interface{})
    err = liveM.StreamsControlHistory(time.Now().Add(-time.Hour * 12), time.Now(), &resp)
    // （v0.5,v0.5-）
    //err = liveM.StreamsControlHistoryWithApp(AppName,time.Now().Add(-time.Hour * 12), time.Now(), &resp)
    fmt.Println(err, resp)
```

* 禁止
```go
    resp = make(map[string]interface{})
    err = liveM.ForbidLiveStreamWithPublisher("video-name", nil, &resp)
    fmt.Println(err, resp)
```

* 恢复
```go
    resp = make(map[string]interface{})
    err = liveM.ResumeLiveStreamWithPublisher("video-name", &resp)
    fmt.Println(err, resp)
```

### 录制（请看文档）
### 截图（2017-01-18）

* 添加截图配置
```go
	oss := live.OssInfo{
		OssBucket:       OssBucket,
		OssEndpoint:     OssEndpoint,
		OssObject:       OssObject,
		OssObjectPrefix: OssObjectPrefix,
	} 
	config:=live.SnapshotConfig{
		OssInfo:oss,
		TimeInterval       : 5,
		OverwriteOssObject  : "{AppName}/{StreamName}.jpg",
	}
  	resp := make(map[string]interface{})
  	err:=liveM.AddLiveAppSnapshotConfig(config,&resp)
  	fmt.Println(err, resp)
```

* 更新截图配置
```go
	config.SequenceOssObject="{AppName}/{StreamName}.jpg"
	resp = make(map[string]interface{})
	err=liveM.UpdateLiveAppSnapshotConfig(config,&resp)
	fmt.Println(err, resp)
```

* 查询域名截图配置
```go
	param:=live.LiveSnapshotParam{
    		PageNum:1,
    		PageSize:10,
    		Order:"asc",
    	}
    	resp = make(map[string]interface{})
    	err=liveM.LiveSnapshotConfig(param,&resp)
    	fmt.Println(err, resp)
```

* 查询截图信息
```go
	resp = make(map[string]interface{})
	err=liveM.LiveStreamSnapshotInfo("test-video-name1",time.Now().Add(-time.Hour*24*20), time.Now(),10,&resp)
	fmt.Println(err, resp)
```

* 删除截图配置
```go
	resp = make(map[string]interface{})
	err=liveM.DeleteLiveAppSnapshotConfig(&resp)
	fmt.Println(err, resp)
```


### 转码（2017-01-19）
* 添加转码配置
```go
	resp := make(map[string]interface{})
	err:=liveM.AddLiveStreamTranscode("a","no","no",&resp)
	fmt.Println(err, resp)
```

* 查询转码配置信息
```go
	resp = make(map[string]interface{})
	err=liveM.LiveStreamTranscodeInfo(&resp)
	fmt.Println(err, resp))
```

* 删除转码配置
```go
	resp = make(map[string]interface{})
	err=liveM.DeleteLiveStreamTranscode("a",&resp)
	fmt.Println(err, resp)
```

### 混流（2017-01-19）

* 开始混流操作
```go
	err=liveM.StartMixStreamsService(...)
```

* 结束混流操作
```go
	err=liveM.StopMixStreamsService(...)
```
### 拉流（2017-10-12）

* 添加拉流信息
```go
	err=liveM.AddLivePullStreamInfoConfig(...)
```

* 删除拉流信息
```go
	err=liveM.DeleteLivePullStreamInfoConfig(...)
```

* 查询域名下拉流配置信息
```go
	err=liveM.DescribeLivePullStreamConfig(...)
```

### 连麦（（2017-10-12）

* 添加连麦配置
```go
	err=liveM.AddLiveMixConfig(...)
```

* 查询连麦配置
```go
	err=liveM.DescribeLiveMixConfig(...)
```

* 删除连麦配置
```go
	err=liveM.DeleteLiveMixConfig(...)
```

* 开启多人连麦服务
```go
	err=liveM.StartMultipleStreamMixService(...)
```

* 停止多人连麦服务
```go
	err=liveM.StopMultipleStreamMixService(...)
```

* 往主流添加一路流
```go
	err=liveM.AddMultipleStreamMixService(...)
```

* 从主流移除一路流
```go
	err=liveM.RemoveMultipleStreamMixService(...)
```

* 添加连麦回调配置
```go
	err=liveM.AddLiveMixNotifyConfig(...)
```

* 查询连麦回调配置
```go
	err=liveM.DescribeLiveMixNotifyConfig(...)
```

* 删除连麦回调配置
```go
	err=liveM.DeleteLiveMixNotifyConfig(...)
```

* 更新连麦回调配置
```go
	err=liveM.UpdateLiveMixNotifyConfig(...)
```


### 状态通知（（2017-10-12）

* 设置回调链接
```go
	err=liveM.SetStreamsNotifyUrlConfig(...)
```

* 删除推流回调配置
```go
	err=liveM.DeleteLiveStreamsNotifyUrlConfig(...)
```

* 查询推流回调配置
```go
	err=liveM.StreamsNotifyUrlConfig(...)
```

### 直播转点播（（2018-05-20）

* 增加直播录制转点播配置
```go
	err=liveM.AddLiveRecordVodConfig(...)
```

* 删除直播录制转点播配置
```go
	err=liveM.DeleteLiveRecordVodConfig(...)
```

* 查询直转点配置列表
```go
	err=liveM.DescribeLiveRecordVodConfigs(...)
```


### 资源监控（（2018-05-20）

* 查询直播域名的网络带宽监控数据
```go
	err=liveM.DescribeLiveDomainBpsData(...)
```

* 查询直播域名录制时长数据
```go
	err=liveM.DescribeLiveDomainRecordData(...)
```

* 查询直播域名截图张数数据
```go
	err=liveM.DescribeLiveDomainSnapshotData(...)
```

* 查询直播域名网络流量监控数据
```go
	err=liveM.DescribeLiveDomainTrafficData(...)
```

* 查询直播域名转码时长数据
```go
	err=liveM.DescribeLiveDomainTranscodeData(...)
```

* 查询直播流历史在线人数
```go
	err=liveM.DescribeLiveStreamHistoryUserNum(...)
```

### 直播审核（（2018-05-20）

* 查询审核配置
```go
	err=liveM.DescribeLiveSnapshotDetectPornConfig(...)
```

* 添加审核配置
```go
	err=liveM.AddLiveSnapshotDetectPornConfig(...)
```

* 更新审核回调
```go
	err=liveM.UpdateLiveSnapshotDetectPornConfig(...)
```

* 删除审核回调
```go
	err=liveM.DeleteLiveSnapshotDetectPornConfig(...)
```

* 查询审核回调
```go
	err=liveM.DescribeLiveDetectNotifyConfig(...)
```

* 添加回调通知
```go
	err=liveM.AddLiveDetectNotifyConfig(...)
```


* 更新回调通知
```go
	err=liveM.UpdateLiveDetectNotifyConfig(...)
```


* 删除审核回调
```go
	err=liveM.DeleteLiveDetectNotifyConfig(...)
```



## 流(Stream)
```go
  //如果 streamCert 为空的话，则代表不开启直播流鉴权
   cert := client.NewCredentials(AccessKeyId, AccessKeySecret)
   streamCert := live.NewStreamCredentials(PrivateKey, live.DefaultStreamTimeout)
   liveM := live.NewLive(cert, DomainName, AppName, streamCert)
  // GetStream 获取直播流
  // @describe 每一次都生成新的流实例，不检查流名的唯一性，并且同一个名字会生成不同的实例的，
  //          所以，使用时候，请自行确保流名的唯一性
   stream := liveM.GetStream("video-name")
```

* 获取RTMP推流地址
```go
    // RTMP 推流地址
    // 如果开启了直播流鉴权，签名失效后，会重新生成新的有效的推流地址
    stream.RtmpPublishUrl()
```

* RTMP 直播播放地址
```go
    url:=stream.RtmpLiveUrls()
```

* HLS 直播播放地址
```go
    url:=stream.HlsLiveUrls()
```

* FLV 直播播放地址
```go
    url:=stream.HttpFlvLiveUrls()
```

* 获取在线人数
```go
    num:=stream.OnlineUserNum()
```

* 是否在线
```go
    isOnline:=stream.Online()
```

* 是否被禁止
```go
    isBlocked:=stream.Blocked()
```

* 获取直播流的帧率和码率
```go
    // type FrameRateAndBitRateInfos struct {
    //	FrameRateAndBitRateInfo []FrameRateAndBitRateInfo
    //}
    //
    //// 各直播流的帧率/码率信息
    //type FrameRateAndBitRateInfo struct {
    //	StreamUrl      string // 直播流的URL
    //	VideoFrameRate int    // 直播流的视频帧率
    //	AudioFrameRate int    // 直播流的音频帧率
    //	BitRate        int    // 直播流的码率
    //}  
    frameRateAndBitRateInfo,err:=stream.FrameRateAndBitRateData()
```

* 获取截图信息
```go
// 查询截图信息
    // type StreamSnapshotInfoResponse struct {
    //    client.Response
    //	  LiveStreamSnapshotInfoList struct {
    //		StreamSnapshotInfo []StreamSnapshotInfo `json:"StreamSnapshotInfo" xml:"StreamSnapshotInfo"`
    //	 } `json:"LiveStreamSnapshotInfoList" xml:"LiveStreamSnapshotInfoList"` //截图内容列表，没有则返回空数组
    //    NextStartTime              string                //
    // }
    streamSnapshotInfo,err:=stream.SnapshotInfo(time.Now().Add(-time.Hour * 12), time.Now(), 10)
```

# 依赖

  `github.com/BPing/go-toolkit/http-client`

# 贡献参与者
* cbping(452775680@qq.com)

# License
 采用 [Apache License, Version 2.0](https://github.com/denverdino/aliyungo/blob/master/LICENSE.txt)许可证授权原则。
 
# 参考文献
参考项目： https://github.com/denverdino/aliyungo
