package factory

// IRuleConfigParser IRuleConfigParser
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParser struct {
}

// Parse Parse
func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParser yamlRuleConfigParser
type yamlRuleConfigParser struct {
}

// Parse Parse
func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type IRuleConfigParseFactory interface {
	CreateParser() IRuleConfigParser
}

type jsonRuleConfigParseFactory struct{}

func (y jsonRuleConfigParseFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

type yamlRuleConfigParseFactory struct{}

func (y yamlRuleConfigParseFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

func NewIRuleConfigParseFactory(t string) IRuleConfigParseFactory {
	switch t {
	case "json":
		return jsonRuleConfigParseFactory{}
	case "yaml":
		return yamlRuleConfigParseFactory{}
	}
	return nil
}
