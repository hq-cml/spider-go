package config

import (
	"github.com/Unknwon/goconfig"
	"github.com/hq-cml/spider-go/basic"
)

//解析配置文件
func ParseConfig(confPath string) (*basic.SpiderConf, error) {
	cfg, err := goconfig.LoadConfigFile(confPath)
	if err != nil {
		panic("Load conf file failed!")
	}

	c := &basic.SpiderConf{}
	if c.GrabMaxDepth, err = cfg.Int("spider", "grabMaxDepth"); err != nil {
		panic("Load conf grabDepth failed!")
	}

	if c.PluginKey, err = cfg.GetValue("spider", "pluginKey"); err != nil {
		panic("Load conf pluginKey failed!")
	}

	if c.RequestChanCapcity, err = cfg.Int("spider", "requestChanCapcity"); err != nil {
		panic("Load conf requestChanCapcity failed!")
	}

	if c.ResponseChanCapcity, err = cfg.Int("spider", "responseChanCapcity"); err != nil {
		panic("Load conf responseChanCapcity failed!")
	}

	if c.EntryChanCapcity, err = cfg.Int("spider", "entryChanCapcity"); err != nil {
		panic("Load conf entryChanCapcity failed!")
	}

	if c.ErrorChanCapcity, err = cfg.Int("spider", "errorChanCapcity"); err != nil {
		panic("Load conf errorChanCapcity failed!")
	}

	if c.DownloaderPoolSize, err = cfg.Int("spider", "downloaderPoolSize"); err != nil {
		panic("Load conf downloaderPoolSize failed!")
	}

	if c.AnalyzerPoolSize, err = cfg.Int("spider", "analyzerPoolSize"); err != nil {
		panic("Load conf analyzerPoolSize failed!")
	}

	if c.MaxIdleCount, err = cfg.Int("spider", "maxIdleCount"); err != nil {
		panic("Load conf maxIdleCount failed!")
	}

	if c.IntervalNs, err = cfg.Int("spider", "intervalNs"); err != nil {
		panic("Load conf intervalNs failed!")
	}

	if c.SummaryDetail, err = cfg.Bool("spider", "summaryDetail"); err != nil {
		panic("Load conf summaryDetail failed!" + err.Error())
	}

	if c.SummaryInterval, err = cfg.Int("spider", "summaryInterval"); err != nil {
		panic("Load conf summaryInterval failed!")
	}

	if c.LogPath, err = cfg.GetValue("log", "logPath"); err != nil {
		panic("Load conf logPath failed!")
	}

	return c, nil
}
