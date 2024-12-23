package main

import (
	"log"

	"yp-webinar/internal/structtags"
)

func main() {
	jsonData := `{"company": "Yandex.Practicum", "number": 37, "text": "Удачи и процветания!"}`
	yamlData := `
company: Yandex.Practicum
number: 37
text: Крепкого здоровья!
`
	tomlData := `
company = "Yandex.Practicum"
number = 37
text = "Отличного настроения!"
`
	xmlData := `<HappyNewYear><company>Yandex.Practicum</company><number>37</number><text>Побед и свершений!</text></HappyNewYear>`

	hnFromJSON := structtags.FromJSON(jsonData)
	hnFromYAML := structtags.FromYAML(yamlData)
	hnFromTOML := structtags.FromTOML(tomlData)
	hnFromXML := structtags.FromXML(xmlData)

	log.Printf("From JSON: %+v\n", hnFromJSON.Congratulations())
	log.Printf("From YAML: %+v\n", hnFromYAML.Congratulations())
	log.Printf("From TOML: %+v\n", hnFromTOML.Congratulations())
	log.Printf("From XML: %+v\n", hnFromXML.Congratulations())
}
