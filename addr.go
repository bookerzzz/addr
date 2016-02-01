package addr

import (
	"fmt"
	"strconv"
	"strings"
)

// Addr type for host:port addresses
type Addr struct {
	Host string
	Port uint32
}

// String implement fmt.Stringer
func (a Addr) String() string {
	if a.Port > 0 {
		return fmt.Sprintf("%s:%d", a.Host, a.Port)
	}
	return a.Host
}

// Parse an address string to the *Addr object. Host is always overridden (even when empty).
// Port is overridden only when provided.
func (a *Addr) Parse(addr string) error {
	parts := strings.Split(addr, ":")
	a.Host = parts[0]
	if len(parts) > 1 {
		p, err := strconv.ParseUint(parts[1], 10, 32)
		if err != nil {
			err = fmt.Errorf("Could not resolve port with error %s", err.Error())
			return err
		}
		a.Port = uint32(p)
	}
	return nil
}

// Make creates a new Addr object from a given host:port string
func Make(addr string) (a Addr, err error) {
	err = a.Parse(addr)
	return
}
