package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	DawnSpecFile = "dawn.json"
	EnumsFile    = "enums.go"
	TypesFile    = "types.go"
	ObjectsFile  = "lib.go"
	Package      = "wgpu"
)

type DawnSpec map[string]Entry

type Entry struct {
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
	Methods  []Method `json:"methods"`
	Members  []Member `json:"members"`
	Values   []Value  `json:"values"`
	Args     []Member `json:"args"`
	Returns  string   `json:"returns"`

	Name   string
	GoName string
}

type Method struct {
	Name    string   `json:"name"`
	Args    []Member `json:"args"`
	Returns any      `json:"returns"`
	Tags    []string `json:"tags"`

	GoName    string
	GoReturns string
}

type Member struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Optional bool   `json:"optional"`

	GoName string
	GoType string
}

type Value struct {
	Name  string `json:"name"`
	Value any    `json:"value"` // Values can be ints or strings in dawn.json
}

var SkipStructs = map[string]bool{
	"string view": true,
}

func main() {
	spec := loadSpec(DawnSpecFile)

	objects := extractObjects(spec)
	structs := extractStructs(spec, objects)
	funcs := extractFunctions(spec, objects)

	writeTypes(TypesFile, structs, funcs)
	writeEnums(EnumsFile, spec)
	writeObjects(ObjectsFile, objects)
	// writeConstants(spec)
	// objs := writeObjects(spec)
	// writeTypes(spec, objs)
}

func loadSpec(path string) DawnSpec {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var rawSpec map[string]json.RawMessage

	err = json.Unmarshal(bytes, &rawSpec)
	if err != nil {
		panic(err)
	}

	filtered := make(DawnSpec)
	for k, v := range rawSpec {
		var entry Entry
		if err := json.Unmarshal(v, &entry); err == nil {
			if entry.Category != "" {
				filtered[k] = entry
			}
		} else {

			if k == "device" {
				fmt.Println(err)
			}
		}
	}
	return filtered
}

func extractObjects(spec DawnSpec) map[string]Entry {
	entries := make(map[string]Entry)

	for name, entry := range spec {
		// Skip Dawn Objects

		if len(entry.Tags) > 0 {
			continue
		}

		if entry.Category != "object" {
			continue
		}

		entry.Name = name
		entry.GoName = toGoName(name)

		var j int
		for i := 0; i < len(entry.Methods); i++ {
			if len(entry.Methods[i].Tags) == 0 {
				m := entry.Methods[i]
				m.GoName = toGoName(m.Name)

				if ret, ok := m.Returns.(string); ok {
					m.GoReturns = toGoType(ret)
				}

				for k := 0; k < len(m.Args); k++ {
					m.Args[k].GoName = toGoArgName(m.Args[k].Name)
					m.Args[k].GoType = toGoType(m.Args[k].Type)
				}
				entry.Methods[j] = m
				j++
			}
		}
		entry.Methods = entry.Methods[:j]

		entries[name] = entry
	}

	return entries
}

func extractStructs(spec DawnSpec, objects map[string]Entry) map[string]Entry {
	structs := make(map[string]Entry)

	for name, entry := range spec {
		// Skip dawn types
		if len(entry.Tags) > 0 {
			continue
		}

		// Skip non-structs
		if entry.Category != "structure" && entry.Category != "callback info" {
			continue
		}

		entry.GoName = toGoName(name)

		for i := 0; i < len(entry.Members); i++ {
			entry.Members[i].GoName = toGoName(entry.Members[i].Name)
			entry.Members[i].GoType = toGoType(entry.Members[i].Type)

			if _, ok := objects[entry.Members[i].Type]; ok {
				entry.Members[i].GoType = "*" + entry.Members[i].GoType
			}
		}

		structs[name] = entry
	}

	return structs
}

func extractFunctions(spec DawnSpec, objects map[string]Entry) map[string]Entry {
	funcs := make(map[string]Entry)

	for name, entry := range spec {
		// Skip dawn types
		if len(entry.Tags) > 0 {
			continue
		}

		// Skip non-funcs
		if entry.Category != "callback function" {
			continue
		}

		entry.GoName = toGoName(name)

		for i := 0; i < len(entry.Args); i++ {
			entry.Args[i].GoName = toGoArgName(entry.Args[i].Name)
			entry.Args[i].GoType = toGoType(entry.Args[i].Type)

			if _, ok := objects[entry.Args[i].Type]; ok {
				entry.Args[i].GoType = "*" + entry.Args[i].GoType
			}
		}

		funcs[name] = entry
	}

	return funcs
}

func writeTypes(path string, structs map[string]Entry, funcs map[string]Entry) {
	w, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	writePreable(w)

	// Write Structs
	for name, def := range structs {

		// Skip certain structs
		if _, ok := SkipStructs[name]; ok {
			continue
		}

		if def.Category == "callback function" {
			continue
		}

		// Write definition
		fmt.Fprintf(w, "type %s struct{\n", def.GoName)
		for _, m := range def.Members {
			fmt.Fprintf(w, "\t%s %s\n", m.GoName, m.GoType)
		}
		fmt.Fprintf(w, "}\n\n")
	}

	// Write Functions
	for _, f := range funcs {
		var args []string
		for _, a := range f.Args {
			typ := a.GoType
			if a.Optional {
				typ = "*" + typ
			}
			args = append(args, a.GoName+" "+a.GoType)
		}

		fmt.Fprintf(w, "type %s func (%s)\n", f.GoName, strings.Join(args, ", "))
	}
}

func writeEnums(path string, spec DawnSpec) {
	w, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	writePreable(w)

	for name, entry := range spec {
		// Skip Dawn definitions
		if len(entry.Tags) > 0 {
			continue
		}

		// Skip non-enums or non-bitmasks
		if entry.Category != "enum" && entry.Category != "bitmask" {
			continue
		}

		// Write definition
		goName := toGoName(name)
		fmt.Fprintf(w, "type %s int\n\n", goName)

		fmt.Fprintf(w, "const (\n")
		for i, v := range entry.Values {
			if i == 0 {
				fmt.Fprintf(w, "\t%s%s %s =  %v\n", goName, toGoName(v.Name), goName, v.Value)
			} else {
				fmt.Fprintf(w, "\t%s%s = %v\n", goName, toGoName(v.Name), v.Value)
			}
		}
		fmt.Fprintf(w, ")\n\n")
	}
}

func writeObjects(path string, objects map[string]Entry) {
	w, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	writePreable(w)

	for name, object := range objects {
		fmt.Fprintf(w, "type %s struct{\n", object.GoName)
		fmt.Fprintf(w, "\tref uintptr\n")
		fmt.Fprintf(w, "}\n\n")

		initial := string(name[0])

		for _, m := range object.Methods {
			var args []string
			for _, a := range m.Args {
				typ := a.GoType
				if a.Optional {
					typ = "*" + typ
				}
				args = append(args, a.GoName+" "+a.GoType)
			}

			fmt.Fprintf(w, "func (%s *%s) %s(%s) %s\n",
				initial,
				object.GoName,
				m.GoName,
				strings.Join(args, ", "),
				m.GoReturns,
			)
		}
		fmt.Fprintf(w, "\n\n")
	}

}

func writePreable(w io.Writer) {
	fmt.Fprintf(w, "// CODE GENERATED. DO NOT EDIT\n")
	fmt.Fprintf(w, "//go:generate go run ./cmd/wrapper/. \n")
	fmt.Fprintf(w, "package wgpu\n\n")
}

// func writeObjects(spec DawnSpec) map[string]struct{} {
// 	objects := make(map[string]struct{})

// 	w, err := os.Create(ObjectsFile)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Fprintf(w, "// CODE GENERATED DO NOT EDIT\n")
// 	fmt.Fprintf(w, "package wgpu\n")

// 	//extract objects
// 	for name, value := range spec {

// 		prefix := toGoName(name)

// 		cat, catOk := obj["category"]
// 		_, tagsOk := obj["tags"]

// 		if catOk && cat == "object" && !tagsOk {
// 			objects[prefix] = struct{}{}
// 		}
// 	}

// 	for name, value := range spec {
// 		obj, ok := value.(map[string]any)
// 		if !ok {
// 			continue
// 		}

// 		prefix := toGoName(name)

// 		cat, catOk := obj["category"]
// 		_, tagsOk := obj["tags"]

// 		methodsd := make([]method, 0)

// 		if catOk && cat == "object" && !tagsOk {
// 			methods, methodsOk := obj["methods"]
// 			fmt.Fprintf(w, "type %s struct{\n", prefix)
// 			if methodsOk {
// 				for _, m := range methods.([]any) {
// 					metd := m.(map[string]any)
// 					if _, ok := metd["tags"]; ok {
// 						continue
// 					}
// 					metdd := method{name: metd["name"].(string)}

// 					if retType, ok := metd["returns"]; ok {
// 						switch val := retType.(type) {
// 						case string:
// 							metdd.ret = val
// 						}
// 					}

// 					args := make([]arg, 0)
// 					if _, ok := metd["args"].([]any); ok {
// 						for _, a := range metd["args"].([]any) {
// 							arga := a.(map[string]any)
// 							args = append(args, arg{
// 								name: arga["name"].(string),
// 								typ:  arga["type"].(string),
// 							})
// 						}
// 					}

// 					metdd.args = args
// 					methodsd = append(methodsd, metdd)
// 				}
// 			}
// 			fmt.Fprintf(w, "}\n")

// 			first := strings.ToLower(string(prefix[0]))
// 			for _, m := range methodsd {
// 				methodName := toGoName(m.name)
// 				args := make([]string, 0)
// 				for _, a := range m.args {
// 					typ := toGoType(a.typ)
// 					if _, ok := objects[typ]; ok {
// 						typ = "*" + typ
// 					}
// 					args = append(args, toGoArgName(a.name)+" "+typ)
// 				}
// 				argss := strings.Join(args, ", ")

// 				ret := toGoType(m.ret)
// 				if _, ok := objects[ret]; ok {
// 					ret = "*" + ret
// 				}
// 				fmt.Fprintf(w, "func (%s *%s) %s(%s) %s{}\n", first, prefix, methodName, argss, ret)
// 			}

// 			fmt.Fprintf(w, "\n")
// 		}
// 	}

// 	return objects
// }

// func writeConstants(spec DawnSpec) {
// 	w, err := os.Create(ConstantFile)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Fprintf(w, "// CODE GENERATED DO NOT EDIT\n")
// 	fmt.Fprintf(w, "package wgpu\n")

// 	for name, value := range spec {
// 		obj, ok := value.(map[string]any)
// 		if !ok {
// 			continue
// 		}

// 		prefix := toGoName(name)

// 		cat, catOk := obj["category"]
// 		_, tagsOk := obj["tags"]
// 		values, valuesOk := obj["values"]

// 		if catOk && cat == "enum" && !tagsOk && valuesOk {
// 			fmt.Fprintf(w, "type %s int\n", prefix)
// 			fmt.Fprintf(w, "const (\n")

// 			for i, v := range values.([]any) {
// 				val := v.(map[string]any)
// 				name := toGoName(val["name"].(string))
// 				if i == 0 {
// 					fmt.Fprintf(w, "\t %s%s %s = %v\n", prefix, name, prefix, val["value"])
// 				} else {
// 					fmt.Fprintf(w, "\t %s%s = %v\n", prefix, name, val["value"])
// 				}
// 			}

// 			fmt.Fprintf(w, ")\n\n")
// 		}

// 		if catOk && cat == "bitmask" && !tagsOk && valuesOk {
// 			fmt.Fprintf(w, "type %s int\n", prefix)
// 			fmt.Fprintf(w, "const (\n")

// 			for i, v := range values.([]any) {
// 				val := v.(map[string]any)
// 				name := toGoName(val["name"].(string))
// 				if i == 0 {
// 					fmt.Fprintf(w, "\t %s%s %s = %v\n", prefix, name, prefix, val["value"])
// 				} else {
// 					fmt.Fprintf(w, "\t %s%s = %v\n", prefix, name, val["value"])
// 				}
// 			}

// 			fmt.Fprintf(w, ")\n\n")
// 		}
// 	}
// }

// func writeTypes(spec DawnSpec, objects map[string]struct{}) {
// 	w, err := os.Create(TypesFile)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Fprintf(w, "// CODE GENERATED DO NOT EDIT\n")
// 	fmt.Fprintf(w, "package wgpu\n")

// 	for name, value := range spec {
// 		obj, ok := value.(map[string]any)
// 		if !ok {
// 			continue
// 		}

// 		cat, catOk := obj["category"]
// 		_, tagsOk := obj["tags"]

// 		prefix := toGoName(name)

// 		if catOk && (cat == "structure" || cat == "callback info") && !tagsOk {
// 			fmt.Fprintf(w, "type %s struct {\n", prefix)

// 			for _, m := range obj["members"].([]any) {
// 				mMap, ok := m.(map[string]any)
// 				if !ok {
// 					continue
// 				}

// 				name := toGoName(mMap["name"].(string))
// 				typ := toGoType(mMap["type"].(string))

// 				if _, ok := objects[typ]; ok {
// 					typ = "*" + typ
// 				}

// 				fmt.Fprintf(w, "\t %s %s\n", name, typ)
// 			}

// 			fmt.Fprintf(w, "}\n\n")
// 		}

// 		if catOk && (cat == "callback function") && !tagsOk {
// 			fmt.Fprintf(w, "type %s func(", prefix)

// 			var args []string

// 			for _, m := range obj["args"].([]any) {
// 				mMap, ok := m.(map[string]any)
// 				if !ok {
// 					continue
// 				}

// 				name := toGoArgName(mMap["name"].(string))
// 				typ := toGoType(mMap["type"].(string))

// 				if _, ok := objects[typ]; ok {
// 					typ = "*" + typ
// 				}

// 				args = append(args, name+" "+typ)
// 			}
// 			fmt.Fprintf(w, strings.Join(args, ","))
// 			fmt.Fprintf(w, ")\n\n")
// 		}
// 	}

// }

func toGoName(in string) string {
	parts := strings.Split(in, " ")
	var out strings.Builder

	for _, p := range parts {
		out.WriteString(capitalize(p))
	}

	return out.String()
}

func toGoArgName(in string) string {
	parts := strings.Split(in, " ")
	var out strings.Builder

	for i, p := range parts {
		if i == 0 {
			if p == "type" {
				p = "typ"
			}
			out.WriteString(p)
		} else {
			out.WriteString(capitalize(p))
		}
	}

	return out.String()
}

func toGoType(cType string) string {
	switch cType {
	case "uint8_t":
		return "uint8"
	case "uint16_t":
		return "uint16"
	case "uint32_t":
		return "uint32"
	case "uint64_t":
		return "uint64"
	case "int8_t":
		return "int8"
	case "int16_t":
		return "int16"
	case "int32_t":
		return "int32"
	case "int64_t":
		return "int64"
	case "bool":
		return "bool"
	case "size_t":
		return "int"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "char":
		return "byte"
	case "void *":
		return "uintptr"
	case "void const *":
		return "[]byte"
	case "void":
		return "[]byte"
	case "string view":
		return "string"
	default:
		return toGoName(cType)
	}
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
