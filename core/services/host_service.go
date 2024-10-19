package services

import (
	"beer/core/entities"

	"github.com/elastic/go-sysinfo"
	"github.com/inhies/go-bytesize"
)

func GetHost() entities.Host {
	host, err := sysinfo.Host()
	if err != nil {
		panic(err)
	}

	memo, _ := host.Memory()
	info := host.Info()

	return entities.Host{
		Host: info.Hostname,
		OS:   info.OS.Type + " " + info.OS.Version,
		Arch: info.NativeArchitecture,
		Memo: entities.Memo{
			Total:     bytesize.New(float64(memo.Total)).String(),
			Free:      bytesize.New(float64(memo.Free)).String(),
			Used:      bytesize.New(float64(memo.Used)).String(),
			Available: bytesize.New(float64(memo.Available)).String(),
		},
	}

}
