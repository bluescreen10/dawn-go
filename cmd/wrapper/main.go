package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
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

	Name      string
	GoName    string
	CName     string
	GoReturns string
}

type Method struct {
	Name    string   `json:"name"`
	Args    []Member `json:"args"`
	Returns any      `json:"returns"`
	Tags    []string `json:"tags"`

	GoName    string
	CName     string
	GoReturns string
}

type Member struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Optional   bool   `json:"optional"`
	Annotation string `json:"annotation"`

	GoName string
	CName  string
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
	enums := extractEnums(spec)

	writeTypes(TypesFile, structs, funcs)
	writeEnums(EnumsFile, enums)
	writeObjects(ObjectsFile, objects, funcs, structs, enums)
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
		entry.GoName = GoName(name)

		var j int
		for i := 0; i < len(entry.Methods); i++ {
			if len(entry.Methods[i].Tags) == 0 {
				m := entry.Methods[i]
				m.GoName = GoName(m.Name)
				m.CName = CFuncName(name, m.Name)

				if ret, ok := m.Returns.(string); ok {
					m.GoReturns = GoType(ret)
				} else {
					if m.Returns != nil {
						m.GoReturns = "(*Buffer, error)"
					}
				}

				for k := 0; k < len(m.Args); k++ {
					m.Args[k].GoName = GoArgName(m.Args[k].Name)
					m.Args[k].GoType = GoType(m.Args[k].Type)
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

		entry.Name = name
		entry.GoName = GoName(name)

		for i := 0; i < len(entry.Members); i++ {
			entry.Members[i].GoName = GoName(entry.Members[i].Name)
			entry.Members[i].GoType = GoType(entry.Members[i].Type)
			entry.Members[i].CName = CMemberName(entry.Members[i].Name)

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
		if entry.Category != "callback function" && entry.Category != "function" {
			continue
		}

		entry.GoName = GoName(name)
		entry.GoReturns = GoType(entry.Returns)
		entry.CName = CFuncName("", name)

		for i := 0; i < len(entry.Args); i++ {
			entry.Args[i].GoName = GoArgName(entry.Args[i].Name)
			entry.Args[i].GoType = GoType(entry.Args[i].Type)

			if _, ok := objects[entry.Args[i].Type]; ok {
				entry.Args[i].GoType = "*" + entry.Args[i].GoType
			}
		}

		funcs[name] = entry
	}

	return funcs
}

func extractEnums(spec DawnSpec) map[string]Entry {
	enums := make(map[string]Entry)

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
		entry.GoName = GoName(name)
		entry.Name = name
		enums[name] = entry
	}

	return enums
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

	// Write Func Types
	for _, f := range funcs {

		if f.Category != "callback function" {
			continue
		}

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

func writeEnums(path string, enums map[string]Entry) {
	w, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	writePreable(w)

	for _, entry := range enums {
		fmt.Fprintf(w, "type %s int\n\n", entry.GoName)

		fmt.Fprintf(w, "const (\n")
		for i, v := range entry.Values {
			if i == 0 {
				fmt.Fprintf(w, "\t%s%s %s =  %v\n", entry.GoName, GoName(v.Name), entry.GoName, v.Value)
			} else {
				fmt.Fprintf(w, "\t%s%s = %v\n", entry.GoName, GoName(v.Name), v.Value)
			}
		}
		fmt.Fprintf(w, ")\n\n")
	}
}

func writeObjects(path string, objects, funcs, types, enums map[string]Entry) {
	w := new(bytes.Buffer)
	writePreable(w)
	writeCGoPreamble(w)
	writeCUtils(w)

	for _, object := range objects {
		fmt.Fprintf(w, "type %s struct{\n", object.GoName)
		fmt.Fprintf(w, "ref uintptr\n")
		fmt.Fprintf(w, "}\n\n")

		for _, m := range object.Methods {
			writeMethodBody(w, object, m, objects, types, enums)
		}
		fmt.Fprintf(w, "\n\n")
	}

	for _, f := range funcs {
		if f.Category != "function" {
			continue
		}

		writeFuncBody(w, f, objects, types, enums)
	}

	src, err := format.Source(w.Bytes())
	if err != nil {
		fmt.Println(w.String())
		panic(err)
	}

	err = os.WriteFile(path, src, 0o644)
	if err != nil {
		panic(err)
	}
}

func writeMethodBody(w io.Writer, object Entry, method Method, objects, types, enums map[string]Entry) {
	receiver := strings.ToLower(string(object.GoName[0]))

	var funcSignature []string
	for _, a := range method.Args {
		typ := a.GoType
		if a.Optional {
			typ = "*" + typ
		}
		funcSignature = append(funcSignature, a.GoName+" "+typ)
	}

	// Write Method Signature
	fmt.Fprintf(w, "func (%s *%s) %s(%s) %s {\n",
		receiver,
		object.GoName,
		method.GoName,
		strings.Join(funcSignature, ", "),
		method.GoReturns,
	)

	cReciever := fmt.Sprintf("c%s", object.GoName)
	fmt.Fprintf(w, "%s := C.WGPU%s(unsafe.Pointer(%s.ref))\n", cReciever, object.GoName, receiver)

	cCallArgs := []string{cReciever}

	for _, a := range method.Args {
		cVar := "c" + capitalize(a.GoName)
		if a.Optional {
			cVar = "p" + capitalize(a.GoName)
		}
		if typ, ok := types[a.Type]; ok {
			writeStructConvert(w, a.GoName, typ, a.Optional)
		} else if obj, ok := objects[a.Type]; ok {
			fmt.Fprintf(w, "%s := C.WGPU%s(unsafe.Pointer(%s.ref))\n", cVar, obj.GoName, a.GoName)
		} else if enum, ok := enums[a.Type]; ok {
			fmt.Fprintf(w, "%s := C.WGPU%s(%s)\n", cVar, enum.GoName, a.GoName)
		} else if a.GoType == "[]byte" {
			fmt.Fprintf(w, "%s := unsafe.Pointer(&%s[0])\n", cVar, a.GoName)
			a.Annotation = ""
		} else {
			fmt.Fprintf(w, "%s := C.%s(%s)\n", cVar, a.Type, a.GoName)
		}

		if (a.Annotation == "const*" || a.Annotation == "*") && !a.Optional {
			cVar = "&" + cVar
		}

		cCallArgs = append(cCallArgs, cVar)
	}

	cCall := fmt.Sprintf("C.%s(%s)", method.CName, strings.Join(cCallArgs, ", "))

	// 3. Handle Returns
	if method.GoReturns == "" {
		fmt.Fprintf(w, "%s\n", cCall)
	} else if strings.HasPrefix(method.GoReturns, "*") {
		// It's a WebGPU Object (Handle). We wrap the C pointer in our Go struct.
		resType := strings.TrimPrefix(method.GoReturns, "*")
		fmt.Fprintf(w, "return &%s{ref: uintptr(%s)}\n", resType, cCall)
	} else if method.GoReturns == "(*Buffer, error)" {
		// Special case for async or complex returns identified in extractObjects
		fmt.Fprintf(w, "_ = %s // TODO: Implement async/error logic\n", cCall)
		fmt.Fprintf(w, "return nil, nil\n")
	} else if method.Returns == "future" {
		fmt.Fprintf(w, "return %s{Id: uint64(%s.id)}\n", method.GoReturns, cCall)
	} else if method.Returns == "bool" {
		fmt.Fprintf(w, "return %s(%s != 0)\n", method.GoReturns, cCall)
	} else if _, ok := objects[method.Returns.(string)]; ok {
		fmt.Fprintf(w, "return %s{ref: uintptr(unsafe.Pointer(%s))}\n", method.GoReturns, cCall)
	} else {
		// It's a basic type (int, bool, etc.)
		fmt.Fprintf(w, "return %s(%s)\n", method.GoReturns, cCall)
	}

	// End
	fmt.Fprintf(w, "}\n\n")
}

func writeFuncBody(w io.Writer, method Entry, objects, types, enums map[string]Entry) {

	var funcSignature []string
	for _, a := range method.Args {
		typ := a.GoType
		if a.Optional {
			typ = "*" + typ
		}
		funcSignature = append(funcSignature, a.GoName+" "+typ)
	}

	// Write Method Signature
	fmt.Fprintf(w, "func %s(%s) %s {\n",
		method.GoName,
		strings.Join(funcSignature, ", "),
		method.GoReturns,
	)

	cCallArgs := []string{}

	for _, a := range method.Args {
		cVar := "c" + capitalize(a.GoName)
		if a.Optional {
			cVar = "p" + capitalize(a.GoName)
		}
		if typ, ok := types[a.Type]; ok {
			writeStructConvert(w, a.GoName, typ, a.Optional)
		} else if obj, ok := objects[a.Type]; ok {
			fmt.Fprintf(w, "%s := C.WGPU%s(unsafe.Pointer(%s.ref))\n", cVar, obj.GoName, a.GoName)
		} else if enum, ok := enums[a.Type]; ok {
			fmt.Fprintf(w, "%s := C.WGPU%s(%s)\n", cVar, enum.GoName, a.GoName)
		} else {
			fmt.Fprintf(w, "%s := C.%s(%s)\n", cVar, a.Type, a.GoName)
		}

		if (a.Annotation == "const*" || a.Annotation == "*") && !a.Optional {
			cVar = "&" + cVar
		}

		cCallArgs = append(cCallArgs, cVar)
	}

	cCall := fmt.Sprintf("C.%s(%s)", method.CName, strings.Join(cCallArgs, ", "))

	// 3. Handle Returns
	if method.GoReturns == "" {
		fmt.Fprintf(w, "%s\n", cCall)
	} else if strings.HasPrefix(method.GoReturns, "*") {
		// It's a WebGPU Object (Handle). We wrap the C pointer in our Go struct.
		resType := strings.TrimPrefix(method.GoReturns, "*")
		fmt.Fprintf(w, "return &%s{ref: uintptr(%s)}\n", resType, cCall)
	} else if method.GoReturns == "(*Buffer, error)" {
		// Special case for async or complex returns identified in extractObjects
		fmt.Fprintf(w, "_ = %s // TODO: Implement async/error logic\n", cCall)
		fmt.Fprintf(w, "return nil, nil\n")
	} else if method.Returns == "future" {
		fmt.Fprintf(w, "return %s{id: uint64(%s.id)}\n", method.GoReturns, cCall)
	} else if method.Returns == "bool" {
		fmt.Fprintf(w, "return %s(%s != 0)\n", method.GoReturns, cCall)
	} else if _, ok := objects[method.Returns]; ok {
		fmt.Fprintf(w, "return %s{ref: uintptr(unsafe.Pointer(%s))}\n", method.GoReturns, cCall)
	} else {
		// It's a basic type (int, bool, etc.)
		fmt.Fprintf(w, "return %s(%s)\n", method.GoReturns, cCall)
	}

	// End
	fmt.Fprintf(w, "}\n\n")
}

func writeStructConvert(w io.Writer, argName string, entry Entry, optional bool) {
	cTypeName := "C.WGPU" + entry.GoName
	cVarName := "c" + capitalize(argName)

	fmt.Fprintf(w, "\t// Convert %s to %s\n", argName, cTypeName)

	if optional {
		// If optional, we define a pointer to the C struct that can stay nil
		fmt.Fprintf(w, "\tvar p%s *%s\n", capitalize(argName), cTypeName)
		fmt.Fprintf(w, "\tif %s != nil {\n", argName)

		// Internal scope for the actual struct value
		writeStructFields(w, "\t\t", argName, cVarName, entry)

		fmt.Fprintf(w, "\t\tp%s = &%s\n", capitalize(argName), cVarName)
		fmt.Fprintf(w, "\t}\n")
	} else {
		// If not optional, we define the struct on the stack
		if entry.Name == "string view" {
			tempVar := cVarName + "Str"
			fmt.Fprintf(w, "\tvar %s C.WGPUStringView\n", cVarName)
			fmt.Fprintf(w, "\t%s := C.CString(%s)\n", tempVar, argName)
			fmt.Fprintf(w, "\t%s.data = %s\n", cVarName, tempVar)
			fmt.Fprintf(w, "\t%s.length = C.size_t(len(%s))\n", cVarName, argName)
			fmt.Fprintf(w, "\tdefer C.free(unsafe.Pointer(%s))\n", tempVar)
		} else {
			writeStructFields(w, "\t", argName, cVarName, entry)
		}

	}
}

// Internal helper to handle the member-by-member assignment
func writeStructFields(w io.Writer, indent, goVar, cVar string, entry Entry) {
	fmt.Fprintf(w, "%svar %s %s\n", indent, cVar, "C.WGPU"+entry.GoName)

	for i, m := range entry.Members {
		goField := goVar + "." + m.GoName
		cField := cVar + "." + m.CName

		switch {
		case m.Type == "string view":
			// Note: In a real production scenario, you would need to defer C.free
			// or use a specialized memory tracker here.
			tempVar := fmt.Sprintf("c%sStr%d", m.GoName, i)
			fmt.Fprintf(w, "%s%s := C.CString(%s)\n", indent, tempVar, goField)
			fmt.Fprintf(w, "%s%s.data = %s\n", indent, cField, tempVar)
			fmt.Fprintf(w, "%s%s.length = C.size_t(len(%s))\n", indent, cField, goField)
			fmt.Fprintf(w, "%sdefer C.free(unsafe.Pointer(%s))\n", indent, tempVar)

		case strings.HasPrefix(m.GoType, "*"):
			// It's a WebGPU Object (Handle). Cast the uintptr ref.
			typeName := strings.TrimPrefix(m.GoType, "*")
			fmt.Fprintf(w, "%sif %s != nil {\n", indent, goField)
			fmt.Fprintf(w, "%s\t%s = C.WGPU%s(unsafe.Pointer(%s.ref))\n", indent, cField, typeName, goField)
			fmt.Fprintf(w, "%s}\n", indent)

		case m.Type == "bool":
			// Cgo often handles bool, but sometimes C expects C.WGPUBool
			fmt.Fprintf(w, "%s%s = boolToWGPUBool(%s)\n", indent, cField, goField)
		case m.Type == "size_t":
			// Cgo often handles bool, but sometimes C expects C.WGPUBool
			fmt.Fprintf(w, "%s%s = C.size_t(%s)\n", indent, cField, goField)

		default:
			// Basic numeric types (uint32, etc)
			// We cast to the C type to be safe
			//fmt.Fprintf(w, "%s%s = C.WGPU%s(%s)\n", indent, cField, m.Type, goField)
		}
	}
}

func writePreable(w io.Writer) {
	fmt.Fprintf(w, "// CODE GENERATED. DO NOT EDIT\n")
	fmt.Fprintf(w, "//go:generate go run ./cmd/wrapper/. \n")
	fmt.Fprintf(w, "package wgpu\n\n")
}

func writeCGoPreamble(w io.Writer) {
	fmt.Fprintf(w, "/*\n")
	fmt.Fprintf(w, "#include <stdio.h>\n")
	fmt.Fprintf(w, "#include <stdlib.h>\n\n")
	fmt.Fprintf(w, "#cgo CFLAGS: -I./include\n")
	fmt.Fprintf(w, "#cgo LDFLAGS: -L./lib -lwebgpu_dawn -framework Metal -framework IOKit -framework QuartzCore -framework Foundation -framework IOSurface -lc++\n")
	fmt.Fprintf(w, "#include <webgpu/webgpu.h>\n")
	fmt.Fprintf(w, "*/\n")

	fmt.Fprintf(w, "import \"C\"\n\n")
	fmt.Fprintf(w, "import (\n")
	fmt.Fprintf(w, "\"unsafe\"\n")
	fmt.Fprintf(w, ")\n\n")
}

func writeCUtils(w io.Writer) {
	fmt.Fprintf(w, "func boolToWGPUBool(in bool) C.WGPUBool {\n")
	fmt.Fprintf(w, "if in {\n")
	fmt.Fprintf(w, "return 1\n")
	fmt.Fprintf(w, "} else {\n")
	fmt.Fprintf(w, "return 0\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "}\n")
}

func GoName(in string) string {
	parts := strings.Split(in, " ")
	var out strings.Builder

	for _, p := range parts {
		out.WriteString(capitalize(p))
	}

	return out.String()
}

func CFuncName(class, method string) string {
	return "wgpu" + GoName(class+" "+method)
}

func GoArgName(in string) string {
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

func CMemberName(in string) string {
	return GoArgName(in)
}

func GoType(cType string) string {
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
	case "proc":
		return "uintptr"
	default:
		return GoName(cType)
	}
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
