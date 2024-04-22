package config

const (
	openai_key         = "OPENAI_KEY"
	openai_model       = "OPENAI_MODEL"
	openai_proxy       = "OPENAI_PROXY"
	openai_base_url    = "OPENAI_BASE_URL"
	openai_temperature = "OPENAI_TEMPERATURE"
	openai_max_tokens  = "OPENAI_MAX_TOKENS"
)

type AiConfig struct {
	key         string
	model       string
	proxy       string
	baseUrl     string
	temperature float64
	maxTokens   int
}

func (c AiConfig) GetKey() string {
	return c.key
}

func (c AiConfig) GetModel() string {
	return c.model
}

func (c AiConfig) GetProxy() string {
	return c.proxy
}

func (c AiConfig) GetBaseUrl() string {
	return c.baseUrl
}

func (c AiConfig) GetTemperature() float64 {
	return c.temperature
}

func (c AiConfig) GetMaxTokens() int {
	return c.maxTokens
}
