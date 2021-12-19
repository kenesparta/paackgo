package config

// AwsSecretsManager Get configuration from AWS Secrets Manager
type AwsSecretsManager struct{}

func (a AwsSecretsManager) Get(_ string) string {
	// Not implemented jet
	return ""
}
