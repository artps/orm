package script_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/phogolabs/gom/script"
)

var _ = Describe("Generator", func() {
	var generator *script.Generator

	BeforeEach(func() {
		dir, err := ioutil.TempDir("", "gom_generator")
		Expect(err).To(BeNil())

		generator = &script.Generator{
			Dir: dir,
		}
	})

	Describe("Create", func() {
		It("creates a command file successfully", func() {
			path, err := generator.Create("commands", "update")
			Expect(err).To(BeNil())
			Expect(path).To(BeARegularFile())
			Expect(path).To(Equal(filepath.Join(generator.Dir, "commands.sql")))
			Expect(generator.Dir).To(BeADirectory())

			data, err := ioutil.ReadFile(path)
			Expect(err).To(BeNil())

			script := string(data)
			Expect(script).To(ContainSubstring("-- name: update"))
		})

		Context("when the file is not provided", func() {
			It("creates a command file successfully", func() {
				path, err := generator.Create("", "update")
				Expect(err).To(BeNil())
				Expect(path).To(BeARegularFile())
				Expect(generator.Dir).To(BeADirectory())

				filename := filepath.Base(path)
				ext := filepath.Ext(path)
				filename = strings.Replace(filename, ext, "", -1)

				_, err = time.Parse("20060102150405", filename)
				Expect(err).To(Succeed())

				data, err := ioutil.ReadFile(path)
				Expect(err).To(BeNil())

				script := string(data)
				Expect(script).To(ContainSubstring("-- name: update"))
			})
		})

		Context("when the file already exists", func() {
			It("adds the command to the file successfully", func() {
				path, err := generator.Create("commands", "update")
				Expect(err).To(BeNil())

				path, err = generator.Create("commands", "delete")
				Expect(err).To(BeNil())

				Expect(path).To(Equal(filepath.Join(generator.Dir, "commands.sql")))
				Expect(generator.Dir).To(BeADirectory())

				data, err := ioutil.ReadFile(path)
				Expect(err).To(BeNil())

				script := string(data)
				Expect(script).To(ContainSubstring("-- name: update"))
				Expect(script).To(ContainSubstring("-- name: delete"))
			})
		})

		Context("when the command already exists", func() {
			It("returns an error", func() {
				buffer := &bytes.Buffer{}
				fmt.Fprintln(buffer, "-- name: update")
				fmt.Fprintln(buffer, "SELECT * FROM migrations;")

				path := filepath.Join(generator.Dir, "commands.sql")
				Expect(ioutil.WriteFile(path, buffer.Bytes(), 0700)).To(Succeed())

				_, err := generator.Create("commands", "update")
				Expect(err).To(MatchError("Command 'update' already exists"))
			})
		})

		Context("when the dir is not valid", func() {
			It("returns an error", func() {
				generator.Dir = ""
				_, err := generator.Create("commands", "update")
				Expect(err).To(MatchError("mkdir : no such file or directory"))
			})
		})
	})
})