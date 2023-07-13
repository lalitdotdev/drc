package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("domain,hasMX,hasSPF,spfRecords,hasDMARC,dmarcRecord\n")

		for scanner.Scan() { 
		checkDomain(scanner.Text()) // call checkDomain function to check domain name
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Couldn't read from input: %v\n",err) 
}
}



func checkDomain(domain string) {

		var hasMX,hasSPF,hasDMARC bool 
		var spfRecords,dmarcRecord string 


		mxRecords, err := net.LookupMX(domain) // lookup MX records for domain name

		if err != nil { // check for errors
		log.Printf("Couldn't lookup MX records for %s: %v\n",domain,err) 
		
		}
		if len(mxRecords) > 0 { 
			hasMX = true
		}

		txtRecords, err := net.LookupTXT(domain) // lookup SPF records for domain name
		if err != nil { 
			log.Printf("Couldn't lookup SPF records for %s: %v\n",domain,err)
		}
    // loop through each SPF record
		for _, record := range txtRecords { 
				if strings.HasPrefix(record, "v=spf1") { // check if SPF record exists
				hasSPF = true // set hasSPF to true
				spfRecords = record // set spfRecords to record
				break

		 }
		}
		// lookup DMARC record for domain name 
		dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
		if err != nil { 
			log.Printf("Couldn't lookup DMARC records for %s: %v\n",domain,err) 
		}

		// loop through each DMARC record
	  for _ , record := range dmarcRecords { 
			if strings.HasPrefix(record, "v=DMARC1") { // check if DMARC record exists
		}
			hasDMARC = true 
			dmarcRecord = record 
			break

		}

		fmt.Printf("%s,%t,%t,%s,%t,%s\n",domain,hasMX,hasSPF,spfRecords,hasDMARC,dmarcRecord) // print results to console

}