package services

import (
	"beer/core/model"
	"log"
	"sync"
	"time"

	"github.com/google/gopacket/pcap"
)

func GetConnectionsToChannel(connectionsChan chan model.Connections) {
	for {
		time.Sleep(time.Duration(10) * time.Second)
		connections := GetConnections()
		connectionsChan <- connections
	}
}

func GetConnections() model.Connections {
	timestamp := time.Now()
	connections := model.Connections{
		Connections: []model.Connection{
			{
				SourceApp: model.App{
					Name:  "web-1",
					Host:  "localhost",
					Ports: []int{4000},
				},
				TargetApp: model.App{
					Name:  "db-1",
					Host:  "localhost",
					Ports: []int{5432},
				},
				Protocol:       "TCP",
				TargetPort:     5432,
				Status:         "ESTABLISHED",
				StartTimestamp: timestamp.Format(time.RFC3339),
				EndTimestamp:   timestamp.Add(time.Duration(10) * time.Second).Format(time.RFC3339),
				Duration:       "10s",
			},
		},
	}
	return connections
}

func getHostActiveNetworkInterfaceNames() []string {
	var result = make([]string, 0)
	// get the interfaces list
	allInterfaces, _ := pcap.FindAllDevs()

	var wg sync.WaitGroup

	for _, itf := range allInterfaces {
		log.Printf("Device: %s Flag: %d Desc: %s  \n", itf.Name, itf.Flags, itf.Description)
		wg.Add(1)
		go func(itf pcap.Interface) {
			handle, err := pcap.OpenLive(itf.Name, 1600, true, pcap.BlockForever)
			defer wg.Done()
			defer handle.Close()

			if err == nil {
				result = append(result, itf.Name)
			}
		}(itf)

	}
	wg.Wait()

	log.Printf("interfaces: %s", result)

	return result
}

//TODO: define BPF filter (tcp, udp, ports...)

//TODO: after get the final list of interfaces, iterate that list, openning a goroutine for each, publishing
// data in a channel to be used in a SSE endpoint
// handle, err := pcap.OpenLive("lo0", 1600, true, pcap.BlockForever)
// if err != nil {
// 	log.Fatal(err)
// }
// defer handle.Close()

// var filter string = "tcp"
// err = handle.SetBPFFilter(filter)
// if err != nil {
// 	log.Fatal(err)
// }

// // Use the handle as a packet source to process all packets
// packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
// for packet := range packetSource.Packets() {
// 	// Process packet here
// 	// Check for the TCP layer
// 	tcpLayer := packet.Layer(layers.LayerTypeTCP)
// 	if tcpLayer != nil {
// 		tcp, _ := tcpLayer.(*layers.TCP)

// 		// Print TCP information
// 		log.Printf("From src port: %d to dst port: %d\n", tcp.SrcPort, tcp.DstPort)
// 		log.Printf("Sequence number: %d\n", tcp.Seq)

// 		// If there's payload, print it as a string
// 		if len(tcp.Payload) > 0 {
// 			log.Printf("Payload: %s\n", string(tcp.Payload))
// 		}
// 	}
// }
