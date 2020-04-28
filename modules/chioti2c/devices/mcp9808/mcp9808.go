package mcp9808

import (
	"fmt"

	"github.com/learnaddict/chiot"
)

const (
	
)

var profile = chiot.I2CProfile{
	autoDetect: true
	address: []uint8 {0x21}
	ident:  []chiot.I2CIdent {
		chiot.I2CIdent {
			name: "Manufacturer"
			address: 0x54
			value: 0x06
		},
		chiot.I2CIdent {
			name: "HardwareID"
			address: 0x07
			value: 0x0400
		},
	}
	
	BaseAddresses  = []int{24, 25, 26, 27, 28, 29, 30, 31}
	
	
	ManufacturerID = 
	HardwareID = 
	
	
	hwTempAddr  = 0x05
	hwManuAddr  = 
	hwDevIDAddr = 
	
}

var device = chiot.Device{
	DeviceName = "mcp9808"
	DeviceType = "i2c"
	I2C = i2c
}

func init() {
	setup()

}

// Read the current temperature from the MCP9808 sensor at the i2c address
func Read(address uint8) (float32, error) {

	i, err := i2c.NewI2C(address, hwBus)
	if err != nil {
		return 0, err
	}
	defer i.Close()

	t, err := i.ReadRegU16BE(hwTempAddr)
	if err != nil {
		return 0, err
	}
	tc := float32(t&0x0FFF) / float32(16)
	if int(tc)&0x1000 == 1 {
		tc -= 256.0
	}
	return tc, nil
}