# network-cli
Simple command line interface

## usage

Pretty simple and straighforward yet powerful cli
almost ready for use.

Run simply:

``` go

go run cli.go command --host <host_name>

i.e

go run cli.go ns --host github.com
go run cli.go ip --host github.com
go run cli.go cname --host github.com
go run cli.go mx --host github.com


```

Which will return ns for github.com:

```
dns1.p08.nsone.net.
dns3.p08.nsone.net.
ns-421.awsdns-52.com.
ns-520.awsdns-01.net.
ns-1283.awsdns-32.org.
dns2.p08.nsone.net.
dns4.p08.nsone.net.
ns-1707.awsdns-21.co.uk.

```
IP address of the host:
```
140.82.121.3
```
CNAME of the host:

```
github.com.
```
and MX record:

```
aspmx.l.google.com. 1
alt1.aspmx.l.google.com. 5
alt2.aspmx.l.google.com. 5
alt4.aspmx.l.google.com. 10
alt3.aspmx.l.google.com. 10
```
List of commands can be found by issuing:

``` go

go run cli.go --help

```
