# A tool to query DNS records

dnsquery is a simple tool to perform DNS queries in the command line which is implemented in golang.

## Installation

To install the package go to your command line and type

```
go get github.com/ErikDeSmedt/dnsquery
go install github.com/ErikDeSmedt/dnsquery
```

## usage

To retrieve all RRs for example.com you can use

```
dnsquery --domain=example.com 
```

To get a list of all available options type

```
dnsquery --help
```



