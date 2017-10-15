# muxd

Multi-purpose utility for converting network traffic

Usage:
```
./muxd --input-protocol tcp --input-host 127.0.0.1 --input-port 8080 --output-protocol udp --output-host myserver.example.com --output-port 12345 --debug
```

Help:
```
NAME:
   muxd - Network multiplexer

USAGE:
   muxd [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --input-protocol value   Input protocol [tcp, udp, redis]
   --input-host value       Input host (IP)
   --input-port value       Input port (default: 0)
   --input-channel value    Input channel [redis]
   --output-protocol value  Output protocol [tcp, udp, redis]
   --output-host value      Output host (IP)
   --output-port value      Output port (default: 0)
   --output-channel value   Output channel [redis]
   --debug                  Set debug log level
   --help, -h               show help
   --version, -v            print the version
   ```
