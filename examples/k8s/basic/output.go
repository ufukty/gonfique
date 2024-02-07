package config

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ApiVersion string `yaml:"apiVersion"`
	Data       struct {
		MyKey    string `yaml:"my-key"`
		Password string `yaml:"password"`
	} `yaml:"data"`
	Kind     string `yaml:"kind"`
	Metadata struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Spec struct {
		Ports []struct {
			Port       int    `yaml:"port"`
			Protocol   string `yaml:"protocol"`
			TargetPort int    `yaml:"targetPort"`
		} `yaml:"ports"`
		Replicas int `yaml:"replicas"`
		Rules    []struct {
			Host string `yaml:"host"`
			Http struct {
				Paths []struct {
					Backend struct {
						Service struct {
							Name string `yaml:"name"`
							Port struct {
								Number int `yaml:"number"`
							} `yaml:"port"`
						} `yaml:"service"`
					} `yaml:"backend"`
					Path     string `yaml:"path"`
					PathType string `yaml:"pathType"`
				} `yaml:"paths"`
			} `yaml:"http"`
		} `yaml:"rules"`
		Selector struct {
			MatchLabels struct {
				App string `yaml:"app"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`
		Template struct {
			Metadata struct {
				Labels struct {
					App string `yaml:"app"`
				} `yaml:"labels"`
			} `yaml:"metadata"`
			Spec struct {
				Containers []struct {
					EnvFrom []struct {
						ConfigMapRef struct {
							Name string `yaml:"name"`
						} `yaml:"configMapRef"`
						SecretRef struct {
							Name string `yaml:"name"`
						} `yaml:"secretRef"`
					} `yaml:"envFrom"`
					Image string `yaml:"image"`
					Name  string `yaml:"name"`
					Ports []struct {
						ContainerPort int `yaml:"containerPort"`
					} `yaml:"ports"`
				} `yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
	Type string `yaml:"type"`
}

func ReadConfig(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("opening config file: %w", err)
	}
	cfg := Config{}
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("decoding config file: %w", err)
	}
	return cfg, nil
}
