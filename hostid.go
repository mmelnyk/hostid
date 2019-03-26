package hostid

import (
	"bytes"
	"crypto/sha256"

	"go.melnyk.org/hufid"
)

var (
	iscached bool
	cached   string
)

// ID returns persistent unique host id based
func ID() (string, error) {
	if !iscached {
		rawid, err := getID()

		if err != nil {
			return "", err
		}

		hashedid := sha256.Sum256([]byte(rawid))
		cached = string(hufid.NewID(3, bytes.NewReader(hashedid[:])))
		iscached = true
	}

	return cached, nil
}
