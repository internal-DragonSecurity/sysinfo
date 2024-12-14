package sysinfo

import (
	"fmt"
	"github.com/earthboundkid/versioninfo/v2"
)

func GetVersion() {
	fmt.Println(versioninfo.Short())
}

func GetVersionMeta() string {
	version := versioninfo.Short()
	return version
}
