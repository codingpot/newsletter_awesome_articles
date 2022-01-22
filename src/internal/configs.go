package internal

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type EmailConfig struct {
	Title       string   `yaml:"title"`
	FooterTitle string   `yaml:"footer_title"`
	Receivers   []string `yaml:"receivers"`

	HeaderTitle           string `yaml:"header_title"`
	HeaderDescription     string `yaml:"header_description"`
	HeaderImageURL        string `yaml:"header_image_url"`
	CommunityLink         string `yaml:"community_link"`
	CommunityLinkBtnTitle string `yaml:"community_link_btn_title"`

	SectionTitle string `yaml:"section_title"`
}

// Config Type
type Config struct {
	Email EmailConfig `yaml:"email"`
}

func GetConfigs(filename string) Config {
	filename, _ = filepath.Abs(filename)
	yamlFile, _ := ioutil.ReadFile(filename)

	var config Config

	if err := yaml.NewDecoder(strings.NewReader(string(yamlFile))).Decode(&config); err != nil {
		fmt.Println(err)
	}

	return config
}
