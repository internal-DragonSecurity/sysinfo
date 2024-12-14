package sysinfo

type SysInfo struct {
	Meta Meta `json:"sysinfo"`
	//Node    Node            `json:"node"`
	OS OS `json:"os"`
	//Kernel  Kernel          `json:"kernel"`
	//Product Product         `json:"product"`
	//Board   Board           `json:"board"`
	//Chassis Chassis         `json:"chassis"`
	//BIOS    BIOS            `json:"bios"`
	//CPU     CPU             `json:"cpu"`
	//Memory  Memory          `json:"memory"`
	//Storage []StorageDevice `json:"storage,omitempty"`
	//Network []NetworkDevice `json:"network,omitempty"`
}

// GetSysInfo gathers all available system information.
func (si *SysInfo) GetSysInfo() {
	// Meta info
	si.getMetaInfo()

	// DMI info
	// SMBIOS info
	// Node info
	// Hardware info
	// Software info
	si.getOSInfo()
}
