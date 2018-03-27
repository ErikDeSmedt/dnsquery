package main

import (
	"fmt"
	"github.com/miekg/dns"
	"github.com/ogier/pflag"
	"os"
)

//The variable names for all the flags
var (
	domain           string
	server           string
	port             uint16
	typeName         string
	className        string
	recursionDesired bool
)

func main() {
	//parse all command line arguments. See init() for more information
	pflag.Parse()

	//Changes example.com to example.com.
	if domain == "" {
		pflag.Usage()
		return
	}
	domain = dns.Fqdn(domain)

	//gets the default server if required
	if server == "" {
		server = getDefaultServer()
	}

	//reads the type of the dns query
	dnsType, ok := dns.StringToType[typeName]
	if !ok {
		fmt.Fprintf(os.Stderr, "The dns type %v is not recognized. Please provide a valid type", typeName)
		return
	}

	//reads the class of the dns query
	dnsClass, ok := dns.StringToClass[className]
	if !ok {
		fmt.Fprintf(os.Stderr, "The class %v is not recognized.", className)
		return
	}

	//Defines the message
	m := new(dns.Msg)
	m.Question = make([]dns.Question, 1)

	q := new(dns.Question)
	q.Name = domain
	q.Qtype = dnsType
	q.Qclass = dnsClass

	m.Question[0] = *q

	m.RecursionDesired = recursionDesired

	//perform the to the secified name server
	addr := fmt.Sprintf("%v:%v", server, port)
	in, err := dns.Exchange(m, addr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not resolve %v to an ip address\n", domain)
		fmt.Fprint(os.Stderr, err)
	} else {
		fmt.Println(in)
	}
	return
}

func init() {
	pflag.StringVarP(&domain, "domain", "d", "", "is the domain to query. An example is 'www.example.com.' (required)")
	pflag.StringVarP(&server, "server", "s", "", "Specifies the server IP-address")
	pflag.Uint16VarP(&port, "port", "p", 53, "Specifies the servers port")
	pflag.StringVarP(&typeName, "type", "t", "ANY", "Specifies the type of the query")
	pflag.StringVarP(&className, "class", "c", "IN", "Specifies the type the query")
	pflag.BoolVar(&recursionDesired, "recursion", true, "true if recursion is desired")
}

func getDefaultServer() string {
	cfg, err := dns.ClientConfigFromFile("/etc/resolv/conf")
	if err != nil {
		fmt.Fprint(os.Stderr, "Could not find the default dns server")
		fmt.Fprint(os.Stderr, err)
	}

	return cfg.Servers[0]

}
