package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type addonOptions struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Namespace  string `json:"namespace"`
	LogLevel   string `json:"log_level"`
	ServerCert string `json:"server_cert"`
	ServerKey  string `json:"server_key"`
}

type telemetryConfig struct {
	Host                   string              `json:"host"`
	Port                   int                 `json:"port"`
	LogLevel               string              `json:"log_level"`
	JSONLogEnable          bool                `json:"json_log_enable"`
	Namespace              string              `json:"namespace"`
	TransmitDecodedRecords bool                `json:"transmit_decoded_records"`
	Records                map[string][]string `json:"records"`
	TLS                    telemetryTLSConfig  `json:"tls"`
}

type telemetryTLSConfig struct {
	ServerCert string `json:"server_cert"`
	ServerKey  string `json:"server_key"`
}

func main() {
	options, err := loadOptions("/data/options.json")
	if err != nil {
		log.Fatalf("load add-on options: %v", err)
	}

	if err := ensureFile(options.ServerCert); err != nil {
		log.Fatal(err)
	}
	if err := ensureFile(options.ServerKey); err != nil {
		log.Fatal(err)
	}

	config := telemetryConfig{
		Host:                   options.Host,
		Port:                   options.Port,
		LogLevel:               options.LogLevel,
		JSONLogEnable:          true,
		Namespace:              options.Namespace,
		TransmitDecodedRecords: true,
		Records: map[string][]string{
			"alerts":       {"logger"},
			"errors":       {"logger"},
			"connectivity": {"logger"},
			"V":            {"logger"},
		},
		TLS: telemetryTLSConfig{
			ServerCert: options.ServerCert,
			ServerKey:  options.ServerKey,
		},
	}

	configPath := "/tmp/fleet-telemetry-config.json"
	if err := writeConfig(configPath, config); err != nil {
		log.Fatalf("write config: %v", err)
	}

	log.Printf("Starting fleet-telemetry for %s:%d", options.Host, options.Port)
	cmd := exec.Command("/fleet-telemetry", "-config="+configPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		log.Fatalf("fleet-telemetry exited: %v", err)
	}
}

func loadOptions(path string) (addonOptions, error) {
	var options addonOptions
	data, err := os.ReadFile(path)
	if err != nil {
		return options, err
	}
	if err := json.Unmarshal(data, &options); err != nil {
		return options, err
	}
	return options, nil
}

func ensureFile(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("required file missing %s: %w", path, err)
	}
	if info.IsDir() {
		return fmt.Errorf("required file is a directory: %s", path)
	}
	return nil
}

func writeConfig(path string, config telemetryConfig) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}
