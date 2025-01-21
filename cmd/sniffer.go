package cmd

import (
	"log"
	"sync"

	"github.com/spf13/cobra"

	"github.com/google/gopacket/pcap"
)

func init() {
	rootCmd.AddCommand(sniffCmd)
}

var sniffCmd = &cobra.Command{
	Use:   "sniff",
	Short: "BEER sniffer",
	Run: func(cmd *cobra.Command, args []string) {
		getHostActiveNetworkInterfaceNames()
	},
}

func getHostActiveNetworkInterfaceNames() []string {
	var result = make([]string, 0)
	// get the interfaces list
	allInterfaces, _ := pcap.FindAllDevs()

	var wg sync.WaitGroup

	for _, itf := range allInterfaces {
		log.Printf("Device: %s Flag: %d Addresses: %s  \n", itf.Name, itf.Flags, itf.Addresses)
		wg.Add(1)
		go func(itf *pcap.Interface) {
			handle, err := pcap.OpenLive(itf.Name, 1600, true, pcap.BlockForever)
			defer wg.Done()
			defer handle.Close()

			if err == nil && len(itf.Addresses) > 0 {
				result = append(result, itf.Name)
			}
		}(&itf)

	}
	wg.Wait()

	log.Printf("found %d interfaces: %s", len(result), result)

	return result
}
