package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path != "" {
		fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
	} else {
		fmt.Fprint(w, "Hello World!\n")
	}
	fmt.Fprint(w, "IP Addresses:\n")
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Fprint(w, "ERROR: %s\n", err)
	} else {
		for _, i := range ifaces {
    		addrs, err := i.Addrs()
			if err != nil {
				fmt.Fprint(w, "ERROR: %s\n", err)
			} else {
    			for _, addr := range addrs {
	        		var ip net.IP
    	    		switch v := addr.(type) {
        			case *net.IPNet:
            	    	ip = v.IP
        			case *net.IPAddr:
                		ip = v.IP
	        		}
    	    		fmt.Fprint(w, "%s\n", ip)
				}
    		}
		}
	}
}
