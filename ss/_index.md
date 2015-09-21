First, install the shadowsock client on your local machine. You'll need the `sslocal` command.

See: [https://github.com/clowwindy/shadowsocks](https://github.com/clowwindy/shadowsocks)

## 测试线路

+ Can you ping the IP?

```
$ ping 108.61.181.71
PING 108.61.181.71 (108.61.181.71): 56 data bytes
64 bytes from 108.61.181.71: icmp_seq=0 ttl=51 time=224.783 ms
64 bytes from 108.61.181.71: icmp_seq=1 ttl=51 time=224.079 ms
64 bytes from 108.61.181.71: icmp_seq=2 ttl=51 time=231.018 ms
```

+ Can you connect to port 443?

```
$ nmap -p 443 -PS 108.61.181.71
Starting Nmap 6.47 ( http://nmap.org ) at 2015-09-21 09:56 CST
Nmap scan report for vukw (108.61.181.71)
Host is up (0.22s latency).
PORT    STATE SERVICE
443/tcp open  https
```

在 [思客学员专用梯子 （测评帖）](http://bbs.sike.io/t/topic/1225) 报告情况。

## ShadowSock Server

Save the following ShadowSock configuration JSON as `config.json` (or whatever):

```
{
    "server": "108.61.181.71",
    "server_port": 443,
    "local_address": "127.0.0.1",
    "local_port":1080,
    "password":"b908e9bc4eab665fb9da5ce79ee981c28a5d9c1b",
    "timeout":360,
    "method":"aes-256-cfb",
    "fast_open": true
}
```

You could also try port 25 (SMTP), see if that's better:

```
{
    "server": "108.61.181.71",
    "server_port": 25,
    "local_address": "127.0.0.1",
    "local_port":1080,
    "password":"b908e9bc4eab665fb9da5ce79ee981c28a5d9c1b",
    "timeout":360,
    "method":"aes-256-cfb",
    "fast_open": true
}
```

Then run the shadowsock client `sslocal -c config.json`, it will start a local proxy server listening on 1080.

Something like:

```
$ sslocal -c config.json
shadowsocks 1.3.6
2014-06-06 11:12:38 INFO     starting local at 127.0.0.1:1080
```

## ProxySwitchy Sharp

ProxySwitchySharp should configure the Sock5 proxy server as localhost:1080

![proxy](https://www.evernote.com/shard/s20/sh/89adf95a-ee50-4de0-b249-e75597a47797/151bed0c82cfb03f4b046057e44ca19c/deep/0/SwitchySharp-Options.png)

## Proxy Commandline Tools With SS

Use the `proxychain` tool.

See https://github.com/clowwindy/shadowsocks/wiki/Using-Shadowsocks-with-Command-Line-Tools