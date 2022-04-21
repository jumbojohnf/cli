package config

type AWSConfig struct {
	Lambda *LambdaConfig `yaml:"lambda,omitempty"`
}

type LambdaConfig struct {
	RoleName string `yaml:"roleName"`
}
