// +build linux

package hostid

import "io/ioutil"

// getID returns unique string based on ...
func getID() (string, error) {
	id, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/product_uuid")
	if err != nil {
		if id, err = ioutil.ReadFile("/etc/machine-id"); err != nil {
			if id, err = ioutil.ReadFile("/var/lib/dbus/machine-id"); err != nil {
				return "", err
			}
		}
	}

	return string(id), nil
}
