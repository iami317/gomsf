# gomsf

> Golang based RPC client to communicate with Metasploit

https://docs.rapid7.com/metasploit/rpc-api

:warning: This is experimental and subject to breaking changes.

## Starting the RPC Server for Metasploit
```bash
msfrpcd -U user -P pass
```

## Connecting to the RPC Server
```golang
client, err := gomsf.New("0.0.0.0:55553")
if err != nil {
    panic(err)
}
if err := client.Login("user", "pass"); err != nil {
    panic(err)
}
defer client.Logout()
```
## Encode data with an encoder
```golang
encoded, err := client.Module.Encode("AAAA", "x86/shikata_ga_nai", &gomsf.EncodeOptions{
    Format: "c",
})
if err != nil {
    panic(err)
}
fmt.Printf("%s\n", encoded)
```
This will encode 'AAAA' with shikata_ga_nai, and prints the following c code:
```bash
unsigned char buf[] =
"\xbb\xc6\xee\x4d\x66\xd9\xee\xd9\x74\x24\xf4\x58\x33\xc9\xb1"
"\x02\x31\x58\x12\x83\xe8\xfc\x03\x9e\xe0\xaf\x93\x5f\xbc\x6e"
"\x1d";
```
## Get infos about a module
```golang
info, err := client.Module.Info(gomsf.ExploitType, "windows/smb/ms08_067_netapi")
if err != nil {
    panic(err)
}

fmt.Printf("Name: %s\n", info.Name)
fmt.Printf("Rank: %s\n", info.Rank)
```
This gives us the metadata of ms08_067_netapi
```bash
Name: MS08-067 Microsoft Server Service Relative Path Stack Corruption
Rank: great
```

## License
[MIT](LICENCE)


ExcellentRanking	这个漏洞永远不会使服务崩溃。这是SQL注入，CMD执行,RFI,LFI等的情况。没有典型的内存损坏漏洞应该给这个rank，除非有特殊情况
GreatRanking	    exploit有一个默认的目标和自动检测目标，或者在版本检查后使用特定于应用程序的返回地址。
GoodRanking	        该exploit具有默认目标，这是这种类型的软件（英语，桌面应用程序的Windows 7，2012的服务器等）的“常见情况”。
NormalRanking	    这个漏洞是可靠的，但取决于一个特定的版本，不能（或不）可靠地自动检测。这个
AverageRanking	    exploit通常是不可靠或者很难被利用的。
LowRanking	        对于通用平台来说，exploit几乎不可能(或者低于50％的成功率)成功。
ManualRanking	    这个exploit不稳定或难以exploit，基本上是一个DoS。当模块没有用处，除非用户特别配置（例如exploit /unix/webapp/php_eval),这个排名也被使用。