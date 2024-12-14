package sysinfo

import (
	"fmt"
	"github.com/earthboundkid/versioninfo/v2"
)

func GetVersion() {
	fmt.Println(versioninfo.Short())
}
