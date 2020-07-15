package lolivoice

import (
	"io"
	"os/exec"
)

var (
	AllPathConstant    = "0.475"
	AdditionalHalfTone = "8.25"
)

type LoliVoiceGenerator struct {
	DictionaryDirPath string
	VoiceFilePath     string
}

func NewLoliVoiceGenerator(dicPath, voicePath string) *LoliVoiceGenerator {
	return &LoliVoiceGenerator{dicPath, voicePath}
}

func (generator *LoliVoiceGenerator) Generate(text, path string) error {
	cmd := exec.Command(
		"open_jtalk",
		"-x", generator.DictionaryDirPath,
		"-m", generator.VoiceFilePath,
		"-a", AllPathConstant,
		"-fm", AdditionalHalfTone,
		"-ow", path,
	)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	io.WriteString(stdin, text)
	stdin.Close()

	return cmd.Run()
}

