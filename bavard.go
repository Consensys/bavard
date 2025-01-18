// Copyright 2020-2024 Consensys Software Inc.
// Licensed under the Apache License, Version 2.0. See the LICENSE file for details.

// Package bavard contains helper functions to generate consistent code from text/template templates.
// it is used by github.com/consensys/gnark && github.com/consensys/gnark-crypto
package bavard

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
	"time"

	"rsc.io/tmplfunc"
)

const (
	envVar = "BAVARD" // environment variable to filter generation
)

// Bavard root object to configure the code generation from text/template
type Bavard struct {
	verbose     bool
	fmt         bool
	imports     bool
	docFile     bool
	packageName string
	license     string
	generated   string
	buildTag    string
	funcs       template.FuncMap
}

// BatchGenerator enables more efficient and clean multiple file generation
type BatchGenerator struct {
	defaultOpts []func(*Bavard) error
}

// NewBatchGenerator returns a new BatchGenerator
func NewBatchGenerator(copyrightHolder string, copyrightYear int, generatedBy string) *BatchGenerator {
	return &BatchGenerator{
		defaultOpts: []func(*Bavard) error{
			Apache2(copyrightHolder, copyrightYear),
			GeneratedBy(generatedBy),
			Format(false),
			Import(false),
			Verbose(true),
		},
	}
}

// Entry to be used in batch generation of files
type Entry struct {
	File      string
	Templates []string
	BuildTag  string
}

func shouldGenerate(output string) bool {
	envFilter := os.Getenv(envVar)
	if envFilter == "" {
		return true
	}

	return strings.Contains(output, envFilter)
}

// GenerateFromString will concatenate templates and create output file from executing the resulting text/template
// see other package functions to add options (package name, licensing, build tags, ...)
func GenerateFromString(output string, templates []string, data interface{}, options ...func(*Bavard) error) error {
	if !shouldGenerate(output) {
		fmt.Printf("skipping generation of %s\n", output)
		return nil // skip generation
	}
	var b Bavard

	var buf bytes.Buffer

	if err := b.config(&buf, output, options...); err != nil {
		return err
	}

	fnHelpers := helpers()
	for k, v := range b.funcs {
		fnHelpers[k] = v
	}

	tmpl := template.New("").Funcs(fnHelpers)

	if err := tmplfunc.Parse(tmpl, aggregate(templates)); err != nil {
		return err
	}

	// execute template
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	return b.create(output, &buf)
}

// GenerateFromFiles will concatenate templates and create output file from executing the resulting text/template
// see other package functions to add options (package name, licensing, build tags, ...)
func GenerateFromFiles(output string, templateF []string, data interface{}, options ...func(*Bavard) error) error {
	if !shouldGenerate(output) {
		fmt.Printf("skipping generation of %s\n", output)
		return nil // skip generation
	}
	var b Bavard
	var buf bytes.Buffer

	b.config(&buf, output, options...)

	// parse templates
	fnHelpers := helpers()
	for k, v := range b.funcs {
		fnHelpers[k] = v
	}

	if len(templateF) == 0 {
		return errors.New("missing templates")
	}
	tName := path.Base(templateF[0])

	tmpl := template.New(tName).Funcs(fnHelpers)

	if err := tmplfunc.ParseFiles(tmpl, templateF...); err != nil {
		return err
	}

	// execute template
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}
	return b.create(output, &buf)
}

func (b *Bavard) config(buf *bytes.Buffer, output string, options ...func(*Bavard) error) error {
	// default settings
	b.imports = false
	b.fmt = false
	b.verbose = true
	b.generated = "bavard"
	b.docFile = strings.HasSuffix(output, "doc.go")

	// handle options
	for _, option := range options {
		if err := option(b); err != nil {
			return err
		}
	}

	if b.buildTag != "" {
		if _, err := buf.WriteString("//go:build  " + b.buildTag + "\n"); err != nil {
			return err
		}
	}

	if b.license != "" {
		if _, err := buf.WriteString(b.license + "\n"); err != nil {
			return err
		}
	}
	if _, err := buf.WriteString(fmt.Sprintf("// Code generated by %s DO NOT EDIT\n\n", b.generated)); err != nil {
		return err
	}

	if !b.docFile && b.packageName != "" {
		if _, err := buf.WriteString("package " + b.packageName + "\n\n"); err != nil {
			return err
		}
	}
	return nil
}

func (b *Bavard) create(output string, buf *bytes.Buffer) error {
	// create output dir if not exist
	_ = os.MkdirAll(filepath.Dir(output), os.ModePerm)

	// create output file
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	if b.verbose {
		fmt.Printf("generating %-70s\n", filepath.Clean(output))
	}
	if _, err := io.Copy(file, buf); err != nil {
		file.Close()
		return err
	}

	file.Close()

	// format generated code
	if b.fmt {
		cmd := exec.Command("gofmt", "-s", "-w", output)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	// run goimports on generated code
	if b.imports {
		cmd := exec.Command("goimports", "-w", output)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func aggregate(values []string) string {
	var sb strings.Builder
	for _, v := range values {
		sb.WriteString(v)
	}
	return sb.String()
}

// Apache2Header returns a Apache2 header string
func Apache2Header(copyrightHolder string, year int) string {
	// get current year from a timestamp
	currentYear := time.Now().Year()
	if currentYear > year {
		apache2 := `
		// Copyright %d-%d %s
		// Licensed under the Apache License, Version 2.0. See the LICENSE file for details.
		`
		return fmt.Sprintf(apache2, year, currentYear, copyrightHolder)
	} else {
		apache2 := `
		// Copyright %d %s
		// Licensed under the Apache License, Version 2.0. See the LICENSE file for details.
		`
		return fmt.Sprintf(apache2, year, copyrightHolder)
	}
}

// Apache2 returns a bavard option to be used in Generate writing an apache2 license header in the generated file
func Apache2(copyrightHolder string, year int) func(*Bavard) error {
	return func(b *Bavard) error {
		b.license = Apache2Header(copyrightHolder, year)
		return nil
	}
}

// GeneratedBy returns a bavard option to be used in Generate writing a standard
// "Code generated by 'label' DO NOT EDIT"
func GeneratedBy(label string) func(*Bavard) error {
	return func(b *Bavard) error {
		b.generated = label
		return nil
	}
}

// BuildTag returns a bavard option to be used in Generate adding build tags string on top of the generated file
func BuildTag(buildTag string) func(*Bavard) error {
	return func(b *Bavard) error {
		b.buildTag = buildTag
		return nil
	}
}

// Package returns a bavard option adding package name and optional package documentation in the generated file
func Package(name string) func(*Bavard) error {
	return func(b *Bavard) error {
		b.packageName = name
		return nil
	}
}

// Verbose returns a bavard option to be used in Generate. If set to true, will print to stdout during code generation
func Verbose(v bool) func(*Bavard) error {
	return func(b *Bavard) error {
		b.verbose = v
		return nil
	}
}

// Format returns a bavard option to be used in Generate. If set to true, will run gofmt on generated file.
// Or simple tab alignment on .s files
func Format(v bool) func(*Bavard) error {
	return func(b *Bavard) error {
		b.fmt = v
		return nil
	}
}

// Import returns a bavard option to be used in Generate. If set to true, will run goimports
func Import(v bool) func(*Bavard) error {
	return func(b *Bavard) error {
		b.imports = v
		return nil
	}
}

// Funcs returns a bavard option to be used in Generate. See text/template FuncMap for more info
func Funcs(funcs template.FuncMap) func(*Bavard) error {
	return func(b *Bavard) error {
		b.funcs = funcs
		return nil
	}
}

// Generate an entry with generator default config
func (b *BatchGenerator) Generate(data interface{}, packageName string, baseTmplDir string, entries ...Entry) error {
	return b.GenerateWithOptions(data, packageName, baseTmplDir, make([]func(*Bavard) error, 0), entries...)
}

// GenerateWithOptions allows adding extra configuration (helper functions etc.) to a batch generation
func (b *BatchGenerator) GenerateWithOptions(data interface{}, packageName string, baseTmplDir string, extraOptions []func(*Bavard) error, entries ...Entry) error {
	var firstError error
	var lock sync.RWMutex
	var wg sync.WaitGroup
	for i := 0; i < len(entries); i++ {
		wg.Add(1)
		go func(entry Entry) {
			defer wg.Done()
			opts := make([]func(*Bavard) error, len(b.defaultOpts)+len(extraOptions))
			copy(opts, b.defaultOpts)
			copy(opts[len(b.defaultOpts):], extraOptions)
			if entry.BuildTag != "" {
				opts = append(opts, BuildTag(entry.BuildTag))
			}
			opts = append(opts, Package(packageName))
			for j := 0; j < len(entry.Templates); j++ {
				entry.Templates[j] = filepath.Join(baseTmplDir, entry.Templates[j])
			}
			if err := GenerateFromFiles(entry.File, entry.Templates, data, opts...); err != nil {
				lock.Lock()
				if firstError == nil {
					firstError = err
				}
				lock.Unlock()
			}
		}(entries[i])
	}
	wg.Wait()

	return firstError
}
