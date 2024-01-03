package llm

type Provider interface {
	EmbedModel()
	Generate()
}
