package sysinfo

// BIOS information.
type BIOS struct {
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
	Date    string `json:"date,omitempty"`
}

func (si *SysInfo) getBIOSInfo() {
	si.BIOS.Vendor = slurpFile("/sys/class/dmi/id/bios_vendor")
	si.BIOS.Version = slurpFile("/sys/class/dmi/id/bios_version")
	si.BIOS.Date = slurpFile("/sys/class/dmi/id/bios_date")
}
