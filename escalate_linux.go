package escalate

import (
	"errors"
	"log"
)

// Escalate TODO: implement
//garble:controlflow flatten_passes=max junk_jumps=max block_splits=max flatten_hardening=xor,delegate_table
func Escalate(path string) error {
	log.Println("Path for bypass: (", path, ")")
	return errors.New("Not implemented yet")
}
