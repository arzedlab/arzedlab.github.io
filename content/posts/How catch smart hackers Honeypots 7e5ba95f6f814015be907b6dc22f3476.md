---
title: "How to catch smart hackers? Honeypots"
description: "This post explains how honeypots attract and study attacker behavior by simulating vulnerable systems. It describes deployment strategies, common attacker patterns, and how the collected data improves defensive detection."
date: 2024-06-20
draft: false 
tags: [Honeypot, ssh, golang, Chameleon] 
toc: true 
---

# Honeypot: An Advanced Cybersecurity Strategy, understanding the Concept of a Honeypot

In the vast landscape of cybersecurity, the term **"honeypot"** often crops up. But what exactly is a honeypot, and why is it so vital in the world of computer security? Let’s dive into this intriguing concept, breaking it down into easily digestible pieces.

## The Lure of the Honeypot

Imagine a jar of honey left out in the open. It’s sweet, tempting, and irresistible to anyone who happens to stumble upon it. In the digital world, a honeypot works similarly. It’s a system or a network deliberately set up to attract cyber attackers, just like the honey attracts bees. However, unlike the real honey which is meant to be consumed, a honeypot is there to watch, learn, and trap malicious actors.

## The Purpose Behind the Sweet Trap

A honeypot is designed to deceive **hackers** and **cybercriminals** by appearing as a vulnerable and valuable target. It mimics a legitimate system, perhaps looking like a server with weak security or a database filled with sensitive information. But behind this façade, the honeypot is closely monitored to see who comes sniffing around.

The main goal is not just to catch intruders but to understand their tactics, techniques, and procedures. By observing how attackers interact with the honeypot, cybersecurity experts gain valuable insights into the latest threats and can bolster their defenses accordingly.

## Types of Honeypots: From Simple to Sophisticated

Honeypots can be quite simple or incredibly sophisticated, depending on their intended use.

- **Low-Interaction Honeypots:** These are basic setups that simulate limited services and capture rudimentary attack data. They’re easier to deploy and manage but don’t provide deep insights into more complex attack methods.
- **High-Interaction Honeypots:** These are more elaborate and closely resemble real systems. They allow attackers to interact with them more extensively, giving security teams a richer source of data about their behavior and techniques. However, they require more resources and careful management to ensure they don’t become a security risk themselves.

## Simple SSH Honeypot in Golang

A basic SSH honeypot crafted with Go acts as a deceptive SSH server, purposefully set up to attract and engage would-be attackers.

```go
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Define the port to listen on
	port := 23 // SSH port

	// Start the honeypot server
	startHoneypot(port)
}

func startHoneypot(port int) {
	// Create a TCP listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to start honeypot: %v", err)
	}
	defer listener.Close()

	fmt.Printf("Honeypot started on port %d...\n", port)

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle incoming connections in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Log connection details
	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("Connection from: %s\n", remoteAddr)

	// Optionally, you can log more details or implement fake responses here

	// Example: Send a fake SSH banner
	conn.Write([]byte("SSH-2.0-OpenSSH_7.9p1 Debian-10+deb10u2\r\n"))

	// Example: Read and log incoming data (if any)
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Error reading from connection: %v", err)
		return
	}
	fmt.Printf("Received %d bytes from %s: %s\n", n, remoteAddr, buf[:n])
}

```

When operational, the honeypot is compiled using the command `go build main.go`, transforming it into an executable that starts listening on port 22, the standard port for SSH traffic. This setup mimics a typical SSH server, making it a prime target for attackers looking for vulnerabilities to exploit.

![Untitled](Untitled.png)

Upon attempting to connect to our honeypot on localhost, we are greeted with a mock SSH banner, designed to mimic the welcoming message of a genuine SSH server. This pseudo-SSH banner further enhances the illusion, convincing potential attackers that they have found a legitimate entry point into the system.

![Untitled](Untitled%201.png)

And server logs clients’ ip address shows up

![Untitled](Untitled%202.png)

This example illustrates the basic concept of a simple SSH honeypot. However, more sophisticated honeypots go far beyond this simplicity, offering advanced capabilities that can surprise even the most experienced attackers with their power and intricacy. Let’s checkout 

[https://github.com/qeeqbox/honeypots](https://github.com/qeeqbox/honeypots)

This is an excellent demonstration of honeypot technology, showcasing a diverse array of 30 low to high-level honeypots conveniently bundled into a single PyPI package. Whether you're observing simple reconnaissance attempts or complex intrusion efforts, this package provides a comprehensive toolkit for understanding and defending against cyber threats.

```bash
$ pip3 install honeypots     // Install with pip3
$ python3 -m honeypots--help // Help 
```

## SSH Honeypot

We are running ssh honeypot with flag

```bash
$ sudo -E python3 -m honeypots --setup ssh --username ravshan --password ravshan
```

Low-Interaction Honeypot of SSH 

![Untitled](Untitled%203.png)

[1] The output when the password is entered correctly.
[2] The output when an incorrect password is entered.
[3] Another example of output when an incorrect password is entered.

## Elastic Honeypot

![Untitled](Untitled%204.png)

[1] One benefit of an elastic honeypot is that it provides user agents used by threat actors.

## Telnet Honeypot

![Untitled](Untitled%205.png)

[1] Telnet output of honeypot that threat actor tried to wrong username and password 

While these simple honeypots offer a good starting point for basic monitoring and initial defense, they often fall short when it comes to providing the comprehensive analytics needed for deeper threat analysis.

# Install Chameleon

QeeqBox’s "**Chameleon**" on GitHub offers an impressive suite of 19 customizable honeypots, each tailored for different network protocols and services. This versatile package is designed to monitor a broad spectrum of network traffic, detect bot activities, and capture attempts to breach systems using various username and password combinations. The Chameleon honeypots cover a wide range of services, including DNS, HTTP Proxy, HTTP, HTTPS, SSH, POP3, IMAP, SMTP, RDP, VNC, SMB, SOCKS5, Redis, TELNET, PostgreSQL, MySQL, MSSQL, Elastic, and LDAP. Each honeypot is meticulously crafted to provide deep insights into these specific areas, making it a powerful tool for enhancing cybersecurity defenses.

# Install Chameleon

```bash
$ git clone https://github.com/qeeqbox/chameleon.git
$ cd chameleon
$ sudo chmod +x ./run.sh
$ sudo ./run.sh deploy
```

After the initialization process, the Grafana interface at [http://localhost:3000](http://localhost:3000/) will open. Use the credentials **`admin`** for both username and password. If the Chameleon dashboard isn't visible initially, navigate to the left sidebar, click on the search icon, and add it.

![Untitled](Untitled%206.png)

Upon completing all the steps, we can view the visually appealing Grafana interface used for monitoring honeypots.

![Untitled](Untitled%207.png)

Nmap scan results displayed on Grafana on Opened Ports graph

```bash
$ nmap 172.19.0.4 -p-
```

![Untitled](Untitled%208.png)

---

Here is the aggregated data for all services, including login attempts and failure counts. And the number of times the threat actor attempted to log in using the entered credentials is recorded.

![Untitled](Untitled%209.png)

---

The number of times the threat actor attempted to log in using the credentials of Telnet is recorded.

![Untitled](Untitled%2010.png)

Here is all the TCP data collected.

![Untitled](Untitled%2011.png)

Here is all the proxy payload data collected.

![Untitled](Untitled%2012.png)

ICMP Payload Data Graph request and reply datas

![Untitled](Untitled%2013.png)

## Learning from the Attacks

One of the key advantages of honeypots is their ability to gather intelligence without putting real assets at risk. By analyzing the data collected from these traps, organizations can better understand where their vulnerabilities lie and how to protect against future threats.

The information gleaned from honeypots can be used to improve firewalls, develop better intrusion detection systems, and create more effective security protocols. Essentially, honeypots turn the tables on attackers, using their own curiosity and greed to strengthen defenses.

## Ethical Considerations and Risks

While honeypots are valuable tools, they come with ethical and operational considerations. It’s crucial to ensure that these traps don’t lead to unintended harm, such as exposing real user data or inadvertently becoming a launchpad for attacks against other systems.

Furthermore, managing a honeypot requires a careful balance. Security teams must ensure that while they’re gathering valuable intelligence, they’re not creating additional vulnerabilities or attracting more malicious activity than they can handle.

## The Future of Honeypots

As cyber threats evolve, so too do honeypots. They’re becoming smarter, more dynamic, and better integrated into broader cybersecurity strategies. Some modern honeypots are even equipped with artificial intelligence to automatically adapt to new attack methods in real-time.

In an era where cyberattacks are becoming more sophisticated and frequent, honeypots serve as a critical line of defense, turning the tables on attackers and helping to keep digital landscapes safe.

## In Summary

A honeypot in cybersecurity is like a honey jar left in the open: it’s designed to attract, observe, and trap. By luring in attackers, honeypots provide invaluable insights that help strengthen defenses against real threats. While they must be carefully managed and ethically deployed, their role in the evolving battlefield of cyber warfare is undeniably significant. So next time you hear about a honeypot, remember, it's more than just a trap; it's a powerful tool in the fight against cybercrime.

---

[Как поймать хитроумных хакеров? Honeypots](https://www.notion.so/Honeypots-2e30dfeda8ba48ed8a87f527c752caf1?pvs=21)
