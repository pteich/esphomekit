package devices

import (
	"github.com/brutella/hc/accessory"
)

type Device interface {
	Init()
	GetDeviceUrl() string
	GetAccessory() *accessory.Accessory
}

type List []Device

func (l List) GetAccessories() []*accessory.Accessory {
	accessories := make([]*accessory.Accessory, len(l))
	for i := range l {
		accessories[i] = l[i].GetAccessory()
	}

	return accessories
}
