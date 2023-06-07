<a name="readme-top"></a>

# ScanBot
## Your network scanning assistant

<div align="center">
    <p align="center">
        <a href="mailto:ahmetimalf2@gmail.com">Report Bug</a>
    </p>
</div>

[![ForTheBadge made-with-python](http://ForTheBadge.com/images/badges/made-with-python.svg)](https://www.python.org/)
[![ForTheBadge built-with-love](http://ForTheBadge.com/images/badges/built-with-love.svg)](https://GitHub.com/ahmetburaki/)


This is a command-line tool written in Go that performs various network security scanning tasks. It utilizes external tools like Nmap, Metasploit Framework, Hydra, and nbtscan to conduct scans and gather information about a target IP address.

## Features

- **Port Scanner**: Uses Nmap to perform a simple network scan and identify open ports on the target IP address.
- **EternalBlue Exploit**: Utilizes the Metasploit Framework to launch the ms17-010 EternalBlue exploit against the target IP address.
- **SSH Brute-Force**: Uses Hydra to perform a brute-force attack on the SSH service of the target IP address using a list of passwords.
- **nbtscan**: Executes the nbtscan tool to perform a scan and retrieve NetBIOS information of the target IP address.

## Prerequisites

Before running this tool, ensure that the following dependencies are installed on your system:

- Nmap: Used for port scanning.
- Metasploit Framework: Required for the EternalBlue exploit.
- Hydra: Used for the SSH brute-force attack.
- nbtscan: Utilized for the nbtscan scan.

## Usage

1. Update the `targetIP` variable in the code with the desired target IP address.
2. Ensure that the required dependencies (Nmap, Metasploit Framework, Hydra, and nbtscan) are installed on your system.
3. Compile and run the code using the Go compiler (`go run main.go`).

The tool will perform the following scans:

- Nmap Scan: Conducts a simple network scan using Nmap and saves the output to the "otobot.txt" file.
- EternalBlue Exploit: Launches the ms17-010 EternalBlue exploit using the Metasploit Framework and saves the output to the "otobot.txt" file.
- SSH BruteForce: Performs a brute-force attack on the SSH service using Hydra and saves the output to the "otobot.txt" file.
- nbtscan: Executes the nbtscan tool to scan and retrieve NetBIOS information of the target IP address, saving the output to the "otobot.txt" file.

Please note that running some of the scans, such as the EternalBlue exploit or SSH brute-force attack, may be illegal or unethical without proper authorization. Ensure that you have appropriate permissions and use the tool responsibly.

## Output

The tool saves the output of each scan to the "otobot.txt" file in the current directory. Each scan is labeled with a heading for easy identification.

```bash
---- Nmap Scan ----
(Nmap scan output)

---- EternalBlue Exploit ----
(Metasploit Framework output)

---- SSH BruteForce ----
(Hydra output)

---- nbtscan ----
(nbtscan output)
```


Please refer to the "otobot.txt" file for the detailed output of each scan.

## Disclaimer

This tool is provided for educational and informational purposes only. The author and the tool are not responsible for any misuse or damage caused by the tool or its usage. Use it responsibly and at your own risk.
