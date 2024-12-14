package sysinfo

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Meta struct {
	Version   string    `yaml:"version"`
	Timestamp time.Time `yaml:"timestamp"`
}

func (si *SysInfo) getMetaInfo() {
	si.Meta.Version = GetVersionMeta()
	si.Meta.Timestamp = time.Now()
}

func GetMetaInfo() {
	var si SysInfo
	si.GetSysInfo()
	data, err := json.MarshalIndent(&si.Meta, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
