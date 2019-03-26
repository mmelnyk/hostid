// +build windows

package hostid

import (
	"golang.org/x/sys/windows/registry"
)

// getID returns the key MachineGuid in registry `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Cryptography`.
// If there is an error running the commad an empty string is returned.
func getID() (string, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Cryptography", registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer key.Close()

	rawid, _, err := key.GetStringValue("MachineGuid")
	if err != nil {
		return "", err
	}

	return rawid, nil
}
