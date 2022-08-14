package factory

// IRuleConfigParser IRuleConfigParser
type IRuleConfigParser interface {
	ParseRule(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParser struct {
}

// Parse Parse
func (J jsonRuleConfigParser) ParseRule(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParser yamlRuleConfigParser
type yamlRuleConfigParser struct {
}

// Parse Parse
func (Y yamlRuleConfigParser) ParseRule(data []byte) {
	panic("implement me")
}

type ISystemConfigParse interface {
	ParseSystem(data []byte)
}

type jsonSystemConfigParse struct{}

func (j jsonSystemConfigParse) ParseSystem(data []byte) {
	panic("implement me")
}

type IConfigParseFactory interface {
	CreateParseRule() jsonRuleConfigParser
	CreateParseSystem() jsonSystemConfigParse
}

type jsonConfigParseFactory struct{}

func (j jsonConfigParseFactory) CreateParseRule() jsonRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParseFactory) CreateParseSystem() jsonSystemConfigParse {
	return jsonSystemConfigParse{}
}
