package main




func Hello(country string, name string) string {
	internationalisation := map[string]string {
		"en": "Hello, ",
		"es": "Hola, ",
		"fr": "Bonjour, ",
		"it": "Ciao, ",
		"de": "Hallo, ",
	}
	if internationalisation[country] == "" {
		country = "en"
	}
	if name == "" {
		name = "World"
	}
	return internationalisation[country] + name
}


func main() {
	
}