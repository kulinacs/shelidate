package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"
)

//go:embed main.c.tmpl
var tmplS string

// byteArray returns a string representation of a C byte array
func byteArray(b []byte) string {
	var s strings.Builder
	fmt.Fprint(&s, "{ ")
full:
	for i := 0; i < len(b); i += 16 {
		for j := 0; j < 16; j++ {
			if i+j >= len(b)-1 {
				fmt.Fprintf(&s, "0x%x", b[i+j])
				break full
			}
			fmt.Fprintf(&s, "0x%x, ", b[i+j])
		}
		fmt.Fprint(&s, "\n")
	}
	fmt.Fprint(&s, " }")
	return s.String()
}

func add(a int, b int) int {
	return a + b
}

type State struct {
	Shellcode []byte
	RWX       bool
	Prepend   bool
}

func generate(shellcode string, rwx bool, prepend bool, out string) error {
	tmpl := template.Must(template.New("output").Funcs(template.FuncMap{
		"add":       add,
		"byteArray": byteArray,
	}).Parse(tmplS))

	shellcodeB, err := os.ReadFile(shellcode)
	if err != nil {
		return fmt.Errorf("unable to open shellcode: %v", err)
	}

	state := State{
		Shellcode: shellcodeB,
		RWX:       rwx,
		Prepend:   prepend,
	}

	output, err := os.CreateTemp("", "example-options-*.c")
	if err != nil {
		return fmt.Errorf("unable to open output file: %v", err)
	}
	defer os.Remove(output.Name())

	if err = tmpl.Execute(output, &state); err != nil {
		return fmt.Errorf("unable to template output file: %v", err)
	}

	if out, err := exec.Command("x86_64-w64-mingw32-gcc", output.Name(), "-o", out).CombinedOutput(); err != nil {
		return fmt.Errorf("unable to compile (%s): %v", out, err)
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %v <shellcode_file> <output_file>", os.Args[0])
		return
	}

	if err := generate(os.Args[1], true, true, os.Args[2]); err != nil {
		fmt.Printf("failed to generate: %v", err)
		return
	}
}
