package config

import (
	"bufio"
	"os"
	"strings"
)

type Config struct {
	AmqpURI string
}

func GetConfig(fileName string) (Config, error) {
	envMap, err := loadEnvFile(fileName)
	cfg := Config{}

	if err != nil {
		return cfg, err
	}

	cfg.setAmqpURI(envMap)
	return cfg, nil
}

func (c *Config) setAmqpURI(envMap map[string]string) {
	val, exits := envMap["AMQP_URI"]
	if exits {
		c.AmqpURI = val
	}
}

func loadEnvFile(fileName string) (map[string]string, error) {
	envMap := make(map[string]string)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		value = strings.Trim(value, `"'`)

		envMap[key] = value
	}

	return envMap, scanner.Err()
}
