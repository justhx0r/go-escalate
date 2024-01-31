package escalate

import (
	"errors"
	"log"
)

// Escalate TODO: implement
//garble:controlflow junk_jumps=4  flatten_hardening=xor
func Escalate(path string) error {
	log.Println("Path for bypass: (", path, ")")
	return errors.New("Not implemented yet")
}
