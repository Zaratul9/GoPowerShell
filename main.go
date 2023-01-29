package main

import (
	"fmt"
	"log"
	"os/exec"
) 

func PowerShell (command string) string {
	cmd := exec.Command("powershell", "-nologo", "-noprofile")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, command)
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func main() {
	out := PowerShell("Get-WmiObject Win32_Processor | Select-Object -ExpandProperty Expand -Property LoadPercentage, L2CacheSize, L3CacheSize, NumberOfCores, NumberOfLogicalProcessors")
	//outTable := strings.Split(out, ":")
	fmt.Println(out)
}
