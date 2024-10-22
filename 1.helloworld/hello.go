package main




func Hello(country string, name string) string {
	internationalisation := map[string]string {
		"en": "Hello, ",
		"es": "Hola, ",
		"fr": "Bonjour, ",
		"it": "Ciao, ",
		"de": "Hallo, ",
	}
	// val of map, bool - does the val exist
	if _, ok := internationalisation[country]; !ok {
		country = "en"
	}
	if name == "" {
		name = "World"
	}
	return internationalisation[country] + name
}


func main() {
	
}