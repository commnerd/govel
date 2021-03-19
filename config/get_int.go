package config

func (c *Config) GetInt(label string) int {
	return c.Viper.GetInt(label)
}
