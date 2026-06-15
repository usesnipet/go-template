package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: normalizeswagger <docs-directory>")
		os.Exit(1)
	}

	docsDir := os.Args[1]
	swaggerPath := filepath.Join(docsDir, "swagger.json")

	spec, err := loadSpec(swaggerPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load swagger.json: %v\n", err)
		os.Exit(1)
	}

	renames, err := normalizeDefinitions(spec)
	if err != nil {
		fmt.Fprintf(os.Stderr, "normalize definitions: %v\n", err)
		os.Exit(1)
	}

	if len(renames) == 0 {
		return
	}

	replaceRefs(spec, renames)

	if err := writeJSON(swaggerPath, spec); err != nil {
		fmt.Fprintf(os.Stderr, "write swagger.json: %v\n", err)
		os.Exit(1)
	}

	yamlPath := filepath.Join(docsDir, "swagger.yaml")
	if err := writeYAML(yamlPath, spec); err != nil {
		fmt.Fprintf(os.Stderr, "write swagger.yaml: %v\n", err)
		os.Exit(1)
	}

	docsGoPath := filepath.Join(docsDir, "docs.go")
	if err := patchDocsGo(docsGoPath, spec); err != nil {
		fmt.Fprintf(os.Stderr, "patch docs.go: %v\n", err)
		os.Exit(1)
	}
}

func loadSpec(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var spec map[string]any
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, err
	}

	return spec, nil
}

func stripPackagePrefix(name string) string {
	if i := strings.LastIndex(name, "."); i >= 0 {
		return name[i+1:]
	}
	return name
}

func normalizeDefinitions(spec map[string]any) (map[string]string, error) {
	rawDefinitions, ok := spec["definitions"].(map[string]any)
	if !ok || len(rawDefinitions) == 0 {
		return nil, nil
	}

	renames := make(map[string]string, len(rawDefinitions))
	normalized := make(map[string]any, len(rawDefinitions))

	for oldName, definition := range rawDefinitions {
		newName := stripPackagePrefix(oldName)
		if newName == oldName {
			normalized[oldName] = definition
			continue
		}

		if _, exists := normalized[newName]; exists {
			return nil, fmt.Errorf("definition name collision: %q and %q resolve to %q", oldName, findKeyByValue(renames, newName), newName)
		}

		renames[oldName] = newName
		normalized[newName] = definition
	}

	spec["definitions"] = normalized
	return renames, nil
}

func findKeyByValue(renames map[string]string, value string) string {
	for oldName, newName := range renames {
		if newName == value {
			return oldName
		}
	}
	return value
}

func replaceRefs(value any, renames map[string]string) {
	switch node := value.(type) {
	case map[string]any:
		if ref, ok := node["$ref"].(string); ok {
			if updated, ok := renameRef(ref, renames); ok {
				node["$ref"] = updated
			}
		}

		for _, child := range node {
			replaceRefs(child, renames)
		}
	case []any:
		for _, child := range node {
			replaceRefs(child, renames)
		}
	}
}

func renameRef(ref string, renames map[string]string) (string, bool) {
	const prefix = "#/definitions/"
	if !strings.HasPrefix(ref, prefix) {
		return "", false
	}

	oldName := strings.TrimPrefix(ref, prefix)
	newName, ok := renames[oldName]
	if !ok {
		return "", false
	}

	return prefix + newName, true
}

func writeJSON(path string, spec map[string]any) error {
	data, err := json.MarshalIndent(spec, "", "    ")
	if err != nil {
		return err
	}

	data = append(data, '\n')
	return os.WriteFile(path, data, 0o644)
}

func writeYAML(path string, spec map[string]any) error {
	data, err := yaml.Marshal(spec)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

func patchDocsGo(path string, spec map[string]any) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	const startMarker = "const docTemplate = `"
	const endMarker = "`"

	text := string(content)
	start := strings.Index(text, startMarker)
	if start == -1 {
		return fmt.Errorf("docTemplate start not found in %s", path)
	}

	start += len(startMarker)
	end := strings.Index(text[start:], endMarker)
	if end == -1 {
		return fmt.Errorf("docTemplate end not found in %s", path)
	}

	jsonBody, err := json.MarshalIndent(spec, "    ", "  ")
	if err != nil {
		return err
	}

	updated := text[:start] + string(jsonBody) + text[start+end:]
	return os.WriteFile(path, []byte(updated), 0o644)
}
