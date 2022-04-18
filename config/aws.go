package config

type AWSConfig struct {
	Lambda LambdaConfig `json:"lambda"`
}

type LambdaConfig struct {
	RoleName string `json:"roleName"`
}
