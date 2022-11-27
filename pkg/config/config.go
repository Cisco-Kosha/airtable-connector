package config

import (
	"flag"
	"os"
)

type Config struct {
	personalAccessToken string
	domainName          string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.personalAccessToken, "AccessToken", "patHrzvIUenm9HMpp.44c52aee648e36c96342f3c3d980bde9f9456fba3bad2af2e527753098b4733b", "Airtable Access token")
	flag.StringVar(&conf.domainName, "AirtableDomainName", os.Getenv("DOMAIN_NAME"), "Airtable Domain Name")

	flag.Parse()

	return conf
}

func (c *Config) GetPersonalAccessToken() string {
	return c.personalAccessToken
}

func (c *Config) GetDomainName() string {
	return c.domainName
}

func (c *Config) GetAirtableURL() string {
	return "https://api.airtable.com/v0"
}
