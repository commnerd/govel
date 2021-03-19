package config

func (c *Config) GetString(label string) string {
	return c.Viper.GetString(label)
}
