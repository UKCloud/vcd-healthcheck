package healthcheck

import (
	"fmt"
	types "github.com/vmware/govcloudair/types/v56"
)

// NetworkDevice checks that the device type is VMXNET3
func NetworkDevice(s types.QueryResultVMRecordType, vm *types.VM) (string, error) {

	deviceType := "unknown"

	for _, v := range vm.VirtualHardwareSection.Item {
		if v.ResourceType == 10 {
			deviceType = v.ResourceSubType
		}
	}
	if deviceType != "VMXNET3" {
		return deviceType, fmt.Errorf("VMs NIC is: %s", deviceType)
	}

	return deviceType, nil
}
