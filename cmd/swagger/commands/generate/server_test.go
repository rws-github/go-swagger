package generate_test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rws-github/go-swagger/cmd/swagger/commands/generate"
	flags "github.com/jessevdk/go-flags"
)

func TestGenerateServer(t *testing.T) {
	specs := []string{
		"billforward.discriminators.yml",
		"todolist.simplequery.yml",
	}
	log.SetOutput(ioutil.Discard)
	defer log.SetOutput(os.Stdout)

	base := filepath.FromSlash("../../../../")
	for i, spec := range specs {
		_ = t.Run(spec, func(t *testing.T) {
			path := filepath.Join(base, "fixtures/codegen", spec)
			generated, err := ioutil.TempDir(filepath.Dir(path), "generated")
			if err != nil {
				t.Fatalf("TempDir()=%s", generated)
			}
			defer func() {
				_ = os.RemoveAll(generated)
			}()
			m := &generate.Server{}
			_, _ = flags.Parse(m)
			if i == 0 {
				m.CopyrightFile = flags.Filename(filepath.Join(base, "LICENSE"))
			}
			if i == 1 {
				m.FlagStrategy = "pflag"
			}
			m.Spec = flags.Filename(path)
			m.Target = flags.Filename(generated)

			if err := m.Execute([]string{}); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGenerateServer_Checks(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	defer log.SetOutput(os.Stdout)

	m := &generate.Server{}
	_, _ = flags.Parse(m)
	m.CopyrightFile = "nowhere"
	err := m.Execute([]string{})
	assert.Error(t, err)
}
