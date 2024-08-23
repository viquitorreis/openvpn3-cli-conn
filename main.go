package main

import (
	"fmt"
	"os/exec"
)

func main() {
	println("BIOFYAGRO VPN HELPER commands: [conn] | [disc] | [list]")
	var input string
	fmt.Scanln(&input)

	if input != "conn" && input != "disc" && input != "list" {
		println("wrong command")
		return
	}

	switch input {
	case "conn":
		initVPNConn()
	case "disc":
		path := readUserPathInput()
		disconnectFromVPN(path)
	case "list":
		listVPNConn()
	}
}

func initVPNConn() {
	// pwd, err := os.Getwd()
	// if err != nil {
	// 	fmt.Printf("error getting pwd: %s", err)
	// }

	// grandeoDir :=

	println("Please accept the VPN connection on your phone")

	cmdStr := fmt.Sprintf(
		"openvpn3 session-start --config /home/victorreis/Projetos/Grandeo/cerberus.conf",
		// pwd,
	)

	cmd := exec.Command(
		"sh", "-c",
		cmdStr,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error initing VPN: %v\n", err)
		return
	}

	fmt.Printf("Connected!")
	fmt.Printf("%s\n", string(out))
}

func listVPNConn() {
	println("Listing OpenVPN3 connections")

	cmdStr := fmt.Sprintf(
		"openvpn3 sessions-list",
	)

	cmd := exec.Command(
		"sh", "-c",
		cmdStr,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error listing VPN connections: %v\n", err)
		return
	}

	fmt.Printf("%s\n", string(out))
}

func readUserPathInput() string {
	var path string
	fmt.Scanln(&path)
	return path
}

func disconnectFromVPN(path string) {
	if path == "" || path == " " || path == "\n" {
		println("Please provide a valid path")
		return
	}

	cmdStr := fmt.Sprintf(
		"openvpn3 session-manage --session-path %s --disconnect",
		path,
	)

	cmd := exec.Command(
		"sh", "-c",
		cmdStr,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error disconnecting from VPN: %v\n", err)
		return
	}

	fmt.Printf("Disconnected!")
	fmt.Printf("%s\n", string(out))
}

// func discVPNConn() {
// 	cmdStr :=
// }
