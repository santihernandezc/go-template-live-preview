//go:build js && wasm

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"syscall/js"
)

func parse(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return js.ValueOf(newValidationError(fmt.Sprintf("incorrect number of arguments: %d, want 2", len(args))))
	}

	tmplStr := args[0].String()
	jsonStr := args[1].String()

	fmt.Printf("Template: %s, Type: %v\n", tmplStr, args[0].Type())
	fmt.Printf("JSON: %s, Type: %v\n", jsonStr, args[1].Type())

	inputMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonStr), &inputMap); err != nil {
		return js.ValueOf(newJsonError(fmt.Sprintf("error unmarshaling input: %v", err)))
	}

	tmpl := template.New("").Option("missingkey=error")
	tmpl, err := tmpl.Parse(tmplStr)
	if err != nil {
		return js.ValueOf(newTemplateError(fmt.Sprintf("error parsing template: %v", err)))
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, inputMap); err != nil {
		return js.ValueOf(newTemplateError(fmt.Sprintf("error executing template: %v", err)))
	}

	return js.ValueOf(map[string]interface{}{
		"parsed": buf.String(),
	})
}

func registerCallbacks() {
	fmt.Println("registering callbacks...")
	js.Global().Set("parse", js.FuncOf(parse))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func newValidationError(msg string) map[string]interface{} {
	return map[string]interface{}{
		"error": msg,
	}
}

func newTemplateError(msg string) map[string]interface{} {
	return map[string]interface{}{
		"templateError": msg,
	}
}

func newJsonError(msg string) map[string]interface{} {
	return map[string]interface{}{
		"jsonError": msg,
	}
}
