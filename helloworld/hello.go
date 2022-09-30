package hello

const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "
const FrenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name, language string) string {

	// if name != "" && language == spanish {
	// 	return SpanishHelloPrefix + name
	// }

	prefix := EnglishHelloPrefix

	switch language {
	case french:
		prefix = FrenchHelloPrefix
	case spanish:
		prefix = SpanishHelloPrefix
	default:
	}

	if name != "" {
		return prefix + name
	}

	return prefix + "world"
}
