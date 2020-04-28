// Package i2c provides low level control over the linux i2c bus.
//
// Before usage you should load the i2c-dev kernel module
//
//      sudo modprobe i2c-dev
//
// Each i2c bus can address 127 independent i2c devices, and most
// linux systems contain several buses.
package chioti2c

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

// New opens a connection to an i2c device.
func (i2c *ChiotI2C) Connect() error {
	//f, err := os.OpenFile(fmt.Sprintf("/dev/i2c-%d", i2c.I2CBus), os.O_RDWR, 0600)
	f, err := os.OpenFile("/dev/i2c-1", os.O_RDWR, 0600)
	if err != nil {
		log.Printf("ERROR1: %v     %v", err, i2c.I2CBus)
		return err
	}

	//_, _, err = syscall.Syscall6(syscall.SYS_IOCTL, f.Fd(), uintptr(i2c.I2CSlave), uintptr(i2c.I2CAddress), 0, 0, 0)

	if err := ioctl(f.Fd(), uintptr(i2c.I2CSlave), uintptr(i2c.I2CAddress)); err != nil {
		//if err != nil {
		log.Printf("ERROR2: %v       %v           %v", err, i2c.I2CSlave, i2c.I2CAddress)
		return err
	}
	log.Print("Successfully connected")
	i2c.Rc = f
	log.Print("Rc set successfully")
	return nil
}

func ioctl(fd, cmd, arg uintptr) error {
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, cmd, arg, 0, 0, 0)
	if err != 0 {
		return err
	}
	return nil
}

// Find any modules and return the addresses of the detected
// func (i *i2c) ScanRange() {

// 	for n := uint8(0); n < i.deviceAddressRange; n++ {

// 		a := i.deviceAddress + n

// 		if v, err := i.ReadRegU16BE(i.vendorIDCmd); err != nil || v != i.vendorIDValue {
// 			continue
// 		}
// 		if v, err := i.ReadRegU16BE(i.deviceIDCmd); err != nil || v != i.deviceIDValue {
// 			continue
// 		}

// 	}
// }

// Write sends buf to the remote i2c device. The interpretation of
// the message is implementation dependant.
func (i2c *ChiotI2C) Write(buf []byte) (int, error) {
	return i2c.Rc.Write(buf)
}

// func (i *i2c) WriteByte(b byte) (int, error) {
// 	var buf [1]byte
// 	buf[0] = b
// 	return i.Rc.Write(buf[:])
// }

func (i2c *ChiotI2C) Read(p []byte) (int, error) {
	return i2c.Rc.Read(p)
}

func (i2c *ChiotI2C) Close() error {
	return i2c.Rc.Close()
}

// SMBus (System Management Bus) protocol over i2c.
// Read byte from i2c device register specified in reg.
// func (i *i2c) ReadRegU8(reg byte) (byte, error) {
// 	_, err := i.Write([]byte{reg})
// 	if err != nil {
// 		return 0, err
// 	}
// 	buf := make([]byte, 1)
// 	_, err = i.Read(buf)
// 	if err != nil {
// 		return 0, err
// 	}
// 	log.Printf("Read U8 %d from reg 0x%0X", buf[0], reg)
// 	return buf[0], nil
// }

// SMBus (System Management Bus) protocol over i2c.
// Write byte to i2c device register specified in reg.
// func (i *i2c) WriteRegU8(reg byte, value byte) error {
// 	buf := []byte{reg, value}
// 	_, err := i.Write(buf)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Write U8 %d to reg 0x%0X", value, reg)
// 	return nil
// }

// SMBus (System Management Bus) protocol over i2c.
// Read unsigned big endian word (16 bits) from i2c device
// starting from address specified in reg.
func (i2c *ChiotI2C) ReadRegU16BE(reg byte) (uint16, error) {
	_, err := i2c.Write([]byte{reg})
	if err != nil {
		return 0, fmt.Errorf("Error during Write    %s", err)
	}
	buf := make([]byte, 2)
	_, err = i2c.Read(buf)
	if err != nil {
		return 0, fmt.Errorf("Error during Read    %s", err)
	}
	w := uint16(buf[0])<<8 + uint16(buf[1])
	log.Printf("Read U16 %d from reg 0x%0X", w, reg)
	return w, nil
}

// SMBus (System Management Bus) protocol over i2c.
// Read unsigned little endian word (16 bits) from i2c device
// starting from address specified in reg.
// func (i *i2c) ReadRegU16LE(reg byte) (uint16, error) {
// 	w, err := i.ReadRegU16BE(reg)
// 	if err != nil {
// 		return 0, err
// 	}
// 	// exchange bytes
// 	w = (w&0xFF)<<8 + w>>8
// 	return w, nil
// }

// SMBus (System Management Bus) protocol over i2c.
// Read signed big endian word (16 bits) from i2c device
// starting from address specified in reg.
// func (i *i2c) ReadRegS16BE(reg byte) (int16, error) {
// 	_, err := i.Write([]byte{reg})
// 	if err != nil {
// 		return 0, err
// 	}
// 	buf := make([]byte, 2)
// 	_, err = i.Read(buf)
// 	if err != nil {
// 		return 0, err
// 	}
// 	w := int16(buf[0])<<8 + int16(buf[1])
// 	log.Printf("Read S16 %d from reg 0x%0X", w, reg)
// 	return w, nil
// }

// SMBus (System Management Bus) protocol over i2c.
// Read unsigned little endian word (16 bits) from i2c device
// starting from address specified in reg.
// func (i *i2c) ReadRegS16LE(reg byte) (int16, error) {
// 	w, err := i.ReadRegS16BE(reg)
// 	if err != nil {
// 		return 0, err
// 	}
// 	// exchange bytes
// 	w = (w&0xFF)<<8 + w>>8
// 	return w, nil

// }

// SMBus (System Management Bus) protocol over i2c.
// Write unsigned big endian word (16 bits) value to i2c device
// starting from address specified in reg.
// func (i *i2c) WriteRegU16BE(reg byte, value uint16) error {
// 	buf := []byte{reg, byte((value & 0xFF00) >> 8), byte(value & 0xFF)}
// 	_, err := i.Write(buf)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Write U16 %d to reg 0x%0X", value, reg)
// 	return nil
// }

// SMBus (System Management Bus) protocol over i2c.
// Write unsigned big endian word (16 bits) value to i2c device
// starting from address specified in reg.
// func (i *i2c) WriteRegU16LE(reg byte, value uint16) error {
// 	w := (value*0xFF00)>>8 + value<<8
// 	return i.WriteRegU16BE(reg, w)
// }

// SMBus (System Management Bus) protocol over i2c.
// Write signed big endian word (16 bits) value to i2c device
// starting from address specified in reg.
// func (i *i2c) WriteRegS16BE(reg byte, value int16) error {
// 	buf := []byte{reg, byte((uint16(value) & 0xFF00) >> 8), byte(value & 0xFF)}
// 	_, err := i.Write(buf)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Write S16 %d to reg 0x%0X", value, reg)
// 	return nil
// }

// SMBus (System Management Bus) protocol over i2c.
// Write signed big endian word (16 bits) value to i2c device
// starting from address specified in reg.
// func (i *i2c) WriteRegS16LE(reg byte, value int16) error {
// 	w := int16((uint16(value)*0xFF00)>>8) + value<<8
// 	return i.WriteRegS16BE(reg, w)
// }
