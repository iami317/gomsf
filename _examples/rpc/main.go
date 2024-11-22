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
	detail := exploit.OptionsDetail()
	fmt.Printf("%#v", detail)
}
