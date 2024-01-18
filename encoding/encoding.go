package encoding

import (
	"encoding/json"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	jsnFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsnFile, &j.DockerCompose)
	if err != nil {
		return err
	}

	ymlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		return err
	}

	ymlFile, err := os.Create(j.FileOutput)
	if err != nil {
		return err
	}

	defer ymlFile.Close()

	_, err = ymlFile.Write(ymlData)
	if err != nil {
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	ymlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(ymlFile, &y.DockerCompose)
	if err != nil {
		return err
	}

	jsnData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		return err
	}

	jsnFile, err := os.Create(y.FileOutput)
	if err != nil {
		return err
	}

	defer jsnFile.Close()

	_, err = jsnFile.Write(jsnData)
	if err != nil {
		return err
	}

	return nil
}
