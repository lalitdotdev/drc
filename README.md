# Domain Record Checker

The Domain Record Checker is a simple command-line tool written in Go that allows you to check the existence of MX, SPF, and DMARC records for a given domain. It provides a summary of the domain's record status, including whether it has MX records, SPF records, DMARC records, and the specific records found, if any.

## Installation

To use the Domain Record Checker, make sure you have Go installed on your system. Then, follow these steps:

1. Clone the repository:

```shell
git clone https://github.com/your-username/domain-record-checker.git
```

2. Change to the project directory:

```shell
cd domain-record-checker
```

3. Build the executable:

```shell
go build
```

4. Run the program:

```shell
go run main.go
```

## Usage

Once you have the Domain Record Checker running, you can enter domain names to check their record status. Simply type the domain name and press Enter. The program will display the results in a tabular format, indicating whether the domain has MX records, SPF records, and DMARC records. If any records are found, the program will also display the specific records.

Here's an example of how to use the Domain Record Checker:

```shell
Enter domain name: example.com
```

```shell
Domain, Has MX, Has SPF, SPF Records, Has DMARC, DMARC Record
example.com, true, true, v=spf1 mx ~all, true, v=DMARC1; p=quarantine; sp=quarantine; pct=100; rua=mailto:dmarc@example.com
```

---

## About Domain Records

### MX Records

MX records (Mail Exchanger records) are DNS records that specify the mail servers responsible for receiving email messages for a particular domain. When someone sends an email to a specific domain, the MX records associated with that domain determine where the email should be delivered. MX records are essential for proper email routing and delivery.

### SPF Records

SPF records (Sender Policy Framework records) are DNS records that indicate which mail servers are authorized to send email on behalf of a domain. These records help prevent email spoofing and unauthorized use of a domain for sending spam or malicious emails. SPF records specify the allowed sending servers by listing their IP addresses or domain names.

### DMARC Records

DMARC records (Domain-based Message Authentication, Reporting, and Conformance records) are DNS records that provide instructions to email receivers on how to handle emails from a particular domain. DMARC records help combat email fraud and phishing attacks. They allow domain owners to specify policies for email authentication and provide guidelines for email receivers to follow when dealing with emails that fail authentication checks.

---

## License

The Domain Record Checker is released under the [MIT License](LICENSE). Feel free to use, modify, and distribute this tool as per the terms of the license.

## Acknowledgements

This project was inspired by the need to quickly check the record status of domains for email authentication and security purposes. We would like to thank the Go community for providing the necessary tools and libraries to create this simple yet useful tool.
