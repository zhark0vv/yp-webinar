package structtags

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type HappyNewYear struct {
	Company string `json:"company" yaml:"company" toml:"company" xml:"company"`
	Number  int    `json:"number" yaml:"number" toml:"number" xml:"number"`
	Text    string `json:"text" yaml:"text" toml:"text" xml:"text"`
}

func FromJSON(j string) HappyNewYear {
	var hn HappyNewYear
	err := json.Unmarshal([]byte(j), &hn)
	if err != nil {
		log.Fatal(err)
	}
	return hn
}

// FromYAML deserializes YAML string to HappyNewYear struct
func FromYAML(y string) HappyNewYear {
	var hn HappyNewYear
	err := yaml.Unmarshal([]byte(y), &hn)
	if err != nil {
		log.Fatal(err)
	}
	return hn
}

// FromTOML deserializes TOML string to HappyNewYear struct
func FromTOML(t string) HappyNewYear {
	var hn HappyNewYear
	err := toml.Unmarshal([]byte(t), &hn)
	if err != nil {
		log.Fatal(err)
	}
	return hn
}

func FromXML(x string) HappyNewYear {
	var hn HappyNewYear
	err := xml.Unmarshal([]byte(x), &hn)
	if err != nil {
		log.Fatal(err)
	}
	return hn
}

func (hn HappyNewYear) Congratulations() string {
	message := fmt.Sprintf("Дорогая когорта %s %d,\n", hn.Company, hn.Number)
	message += "Поздравляю вас с наступающим Новым годом!\n"
	message += fmt.Sprintf("Напутствие для вас: %s\n", hn.Text)
	return message
}
