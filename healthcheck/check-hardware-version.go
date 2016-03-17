package healthcheck

import (
	"fmt"
	types "github.com/skyscape-cloud-services/vmware-govcd/types/v56"
)

func HardwareVersion(s types.QueryResultVMRecordType, vm *types.VM) (string, error) {

	HWVersion := fmt.Sprintf("%d", s.HardwareVersion)
	if s.HardwareVersion != 9 {
		return HWVersion, fmt.Errorf("VM H/W is: %s", HWVersion)
	}

	return HWVersion, nil
}
