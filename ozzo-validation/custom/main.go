package main

import (
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mingrammer/commonregex"
)

func checkIP(value interface{}) error {
	ip, ok := value.(string)
	if !ok {
		return errors.New("ip must be string")
	}

	ipList := commonregex.IPs(ip)
	if len(ipList) != 1 || ipList[0] != ip {
		return errors.New("invalid ip format")
	}

	return nil
}

type Addr struct {
	IP   string
	Port int
}

func (a *Addr) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.IP, validation.Required, validation.By(checkIP)),
		validation.Field(&a.Port, validation.Min(1024), validation.Max(65536)))
}

func main() {
	a1 := &Addr{
		IP:   "127.0.0.1",
		Port: 6666,
	}

	a2 := &Addr{
		IP:   "xxx.yyy.zzz.hhh",
		Port: 7777,
	}

	fmt.Println("addr1:", validation.Validate(a1))
	fmt.Println("addr2:", validation.Validate(a2))
}
