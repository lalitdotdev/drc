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
		scanner := bufio.NewScanner(os.Stdin) // create new scanner to read input from user (domain name) 

	fmt.Printf("domain,hasMX,hasSPF,spfRecords,hasDMARC,dmarcRecord\n") // prompt user to enter domain name

	for scanner.Scan() { // loop through each line of input

		checkDomain(scanner.Text()) // call checkDomain function to check domain name
	}
	if err := scanner.Err(); err != nil { // check for errors

		log.Fatal("Couldn't read from input: %v\n",err) // log error

}
}



func checkDomain(domain string) {

		var hasMX,hasSPF,hasDMARC bool // create variables to store boolean values

		var spfRecords,dmarcRecord string // create variables to store string values


		mxRecords, err := net.LookupMX(domain) // lookup MX records for domain name

		if err != nil { // check for errors
		log.Printf("Couldn't lookup MX records for %s: %v\n",domain,err) // log error
		
		}
		if len(mxRecords) > 0 { // check if MX records exist
			hasMX = true // set hasMX to true
		}

		txtRecords, err := net.LookupTXT(domain) // lookup SPF records for domain name
		if err != nil { // check for errors
			log.Printf("Couldn't lookup SPF records for %s: %v\n",domain,err) // log error
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
		if err != nil { // check for errors
			log.Printf("Couldn't lookup DMARC records for %s: %v\n",domain,err) // log error
		}

		// loop through each DMARC record
	  for _ , record := range dmarcRecords { 
			if strings.HasPrefix(record, "v=DMARC1") { // check if DMARC record exists
		}
			hasDMARC = true // set hasDMARC to true
			dmarcRecord = record // set dmarcRecord to record
			break

		}

		fmt.Printf("%s,%t,%t,%s,%t,%s\n",domain,hasMX,hasSPF,spfRecords,hasDMARC,dmarcRecord) // print results to console



}