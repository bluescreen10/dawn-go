package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"iter"
	"maps"
	"os"
	"slices"
	"strings"
)

const (
	DawnSpecFile  = "dawn.json"
	EnumsFile     = "enums.go"
	TypesFile     = "types.go"
	ObjectsFile   = "lib.go"
	CallbacksFile = "callbacks.c"
	Package       = "wgpu"
)

type RawSpec map[string]Entry

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
	CReturns  string
}

func (e Entry) GetName() string {
	return e.Name
}

func (e Entry) GetCName() string {
	return e.CName
}

func (e Entry) GetGoName() string {
	return e.GoName
}

func (e Entry) GetArgs() []Member {
	return e.Args
}

func (e Entry) GetReturns() any {
	return e.Returns
}

func (e Entry) GetCReturns() string {
	return e.CReturns
}

func (e Entry) GetGoReturns() string {
	return e.GoReturns
}

type Method struct {
	Name    string   `json:"name"`
	Args    []Member `json:"args"`
	Returns any      `json:"returns"`
	Tags    []string `json:"tags"`

	GoName    string
	CName     string
	GoReturns string
	CReturns  string
}

func (m Method) GetName() string {
	return m.Name
}

func (m Method) GetCName() string {
	return m.CName
}

func (m Method) GetGoName() string {
	return m.GoName
}
func (m Method) GetArgs() []Member {
	return m.Args
}

func (m Method) GetReturns() any {
	return m.Returns
}

func (m Method) GetCReturns() string {
	return m.CReturns
}

func (m Method) GetGoReturns() string {
	return m.GoReturns
}

type Member struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Optional   bool   `json:"optional"`
	Annotation string `json:"annotation"`

	GoName string
	CName  string
	GoType string
	CType  string
}

type Value struct {
	Name  string `json:"name"`
	Value any    `json:"value"` // Values can be ints or strings in dawn.json
}

var SkipStructs = map[string]bool{
	"string view":        true,
	"supported features": true,
}

var SkipMethod = map[string]bool{
	"get lost future": true,
}

var Unexported = map[string]bool{
	"status":                 true,
	"callbackMode":           true,
	"request adapter status": true,
}

type SortedMap map[string]Entry

func (s SortedMap) Iter() iter.Seq2[string, Entry] {
	return func(yield func(K string, V Entry) bool) {
		var keys []string
		for k := range maps.Keys(s) {
			keys = append(keys, k)
		}

		slices.Sort(keys)
		for _, k := range keys {
			if !yield(k, s[k]) {
				break
			}
		}
	}
}

type DawnSpec struct {
	Objects   SortedMap
	Structs   SortedMap
	Enums     SortedMap
	Funcs     SortedMap
	CallBacks SortedMap
}

func (s *DawnSpec) isCallbackInfo(name string) bool {
	if s, ok := s.Structs[name]; ok {
		return s.Category == "callback info"
	}
	return false
}

func (s *DawnSpec) isStruct(name string) bool {
	if name == "string view" {
		return false
	}
	_, ok := s.Structs[name]
	return ok
}

func (s *DawnSpec) isObject(name string) bool {
	_, ok := s.Objects[name]
	return ok
}

func (s *DawnSpec) isEnum(name string) bool {
	_, ok := s.Enums[name]
	return ok
}

func (s *DawnSpec) getStruct(name string) Entry {
	return s.Structs[name]
}

func (s *DawnSpec) getCallbackInfo(name string) Entry {
	return s.Structs[name]
}

func (s *DawnSpec) getFunction(name string) Entry {
	return s.Funcs[name]
}

func (s *DawnSpec) getCallback(name string) Entry {
	return s.CallBacks[name]
}

func (s *DawnSpec) getObject(name string) Entry {
	return s.Objects[name]
}

func (s *DawnSpec) getEnum(name string) Entry {
	return s.Enums[name]
}

type Function interface {
	GetArgs() []Member
	GetName() string
	GetCName() string
	GetGoName() string
	GetReturns() any
	GetCReturns() string
	GetGoReturns() string
}

func main() {
	rawSpec := loadSpec(DawnSpecFile)

	spec := &DawnSpec{
		Objects:   make(SortedMap),
		Structs:   make(SortedMap),
		Enums:     make(SortedMap),
		Funcs:     make(SortedMap),
		CallBacks: make(SortedMap),
	}

	// extract object/types/enums/functions
	extractObjects(rawSpec, spec)
	extractStructs(rawSpec, spec)
	extractFunctions(rawSpec, spec)
	extractEnums(rawSpec, spec)

	// write wrapper files
	writeTypes(TypesFile, spec)
	writeEnums(EnumsFile, spec)
	writeObjects(ObjectsFile, spec)
	writeCallbacks(CallbacksFile, spec)
}

func loadSpec(path string) RawSpec {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var rawSpec map[string]json.RawMessage

	err = json.Unmarshal(bytes, &rawSpec)
	if err != nil {
		panic(err)
	}

	filtered := make(RawSpec)
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

func extractObjects(rawSpec RawSpec, spec *DawnSpec) {

	for name, entry := range rawSpec {
		// Skip Dawn Objects

		if len(entry.Tags) > 0 {
			continue
		}

		if entry.Category != "object" {
			continue
		}

		entry.Name = name
		entry.GoName = GoName(name)
		entry.CName = CType(name)

		var j int
		for i := 0; i < len(entry.Methods); i++ {
			if len(entry.Methods[i].Tags) == 0 {
				m := entry.Methods[i]
				m.GoName = GoName(m.Name)
				m.CName = CFuncName(name, m.Name)

				if ret, ok := m.Returns.(string); ok {
					m.GoReturns = GoType(ret)
					m.CReturns = ret
				} else {
					if m.Returns != nil {
						m.GoReturns = "*Buffer"
						m.CReturns = "WGPUBuffer"
					}
				}

				for k := 0; k < len(m.Args); k++ {
					m.Args[k].GoName = GoArgName(m.Args[k].Name)
					m.Args[k].GoType = GoType(m.Args[k].Type)
					m.Args[k].CType = CType(m.Args[k].Type)
				}

				entry.Methods[j] = m
				j++
			}
		}
		entry.Methods = entry.Methods[:j]

		spec.Objects[name] = entry
	}
}

func extractStructs(rawSpec RawSpec, spec *DawnSpec) {
	for name, entry := range rawSpec {
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
		entry.CName = CType(name)

		for i := 0; i < len(entry.Members); i++ {
			entry.Members[i].GoName = GoName(entry.Members[i].Name)
			entry.Members[i].CName = CMemberName(entry.Members[i].Name)
			entry.Members[i].GoType = GoType(entry.Members[i].Type)
			entry.Members[i].CType = CType(entry.Members[i].Type)

			if _, ok := spec.Objects[entry.Members[i].Type]; ok {
				entry.Members[i].GoType = "*" + entry.Members[i].GoType
			}
		}

		spec.Structs[name] = entry
	}
}

func extractFunctions(rawSpec RawSpec, spec *DawnSpec) {

	for name, entry := range rawSpec {
		// Skip dawn types
		if len(entry.Tags) > 0 {
			continue
		}

		// Skip non-funcs
		if entry.Category != "callback function" && entry.Category != "function" {
			continue
		}

		entry.Name = name
		entry.GoName = GoName(name)
		entry.CName = CFuncName("", name)
		entry.GoReturns = GoType(entry.Returns)
		entry.CReturns = entry.Returns

		for i := 0; i < len(entry.Args); i++ {
			entry.Args[i].GoName = GoArgName(entry.Args[i].Name)
			entry.Args[i].CName = CMemberName(entry.Args[i].Name)
			entry.Args[i].GoType = GoType(entry.Args[i].Type)
			entry.Args[i].CType = CType(entry.Args[i].Type)

			if spec.isObject(entry.Args[i].Type) {
				entry.Args[i].GoType = "*" + entry.Args[i].GoType
			}
		}

		if entry.Category == "function" {
			spec.Funcs[name] = entry
		} else {
			spec.CallBacks[name] = entry
		}
	}

}

func extractEnums(rawSpec RawSpec, spec *DawnSpec) {

	for name, entry := range rawSpec {
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
		entry.CName = CType(name)
		entry.Name = name
		spec.Enums[name] = entry
	}
}

func writeTypes(path string, spec *DawnSpec) {
	w, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	writePreable(w)

	// Write Structs
	for name, def := range spec.Structs.Iter() {

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
	for _, f := range spec.CallBacks.Iter() {
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

func writeEnums(path string, spec *DawnSpec) {
	w, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	writePreable(w)

	for _, entry := range spec.Enums.Iter() {
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

func writeObjects(path string, spec *DawnSpec) {
	w := new(bytes.Buffer)
	writePreable(w)
	writeCGoPreamble(w, spec.CallBacks)
	writeCUtils(w)

	// Write Objects
	for _, object := range spec.Objects.Iter() {
		fmt.Fprintf(w, "type %s struct{\n", object.GoName)
		fmt.Fprintf(w, "ref uintptr\n")
		fmt.Fprintf(w, "}\n\n")

		for _, m := range object.Methods {
			if m.CReturns == "future" {
				writeAsyncMethodBody(w, object.GoName, m, spec)
			} else {
				writeMethodBody(w, object.GoName, m, spec)
			}
		}
		fmt.Fprintf(w, "\n\n")
	}

	// Write Standalone functions
	for _, f := range spec.Funcs.Iter() {
		if f.Returns == "future" {
			writeAsyncMethodBody(w, "", f, spec)
		} else {
			writeMethodBody(w, "", f, spec)
		}
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

func writeAsyncMethodBody(w io.Writer, typ string, method Function, spec *DawnSpec) {
	// Ignore skipped methods
	if _, ok := SkipMethod[method.GetName()]; ok {
		return
	}

	writeCallbackHandler(w, method, spec)
	writeMethodBody(w, typ, method, spec)
}

func writeMethodBody(w io.Writer, typ string, method Function, spec *DawnSpec) {
	// Ignore skipped methods
	if _, ok := SkipMethod[method.GetName()]; ok {
		return
	}

	var funcSignature []string
	for _, a := range method.GetArgs() {
		typ := a.GoType
		if a.Optional || spec.isObject(a.Type) {
			typ = "*" + typ
		}
		funcSignature = append(funcSignature, a.GoName+" "+typ)
	}

	var cCallArgs []string
	var methodPrefix string
	var receiver string

	// Method
	if len(typ) > 0 {
		receiver = strings.ToLower(string(typ[0]))
		methodPrefix = fmt.Sprintf("(%s *%s)", receiver, typ)
	}

	// Signature
	fmt.Fprintf(w, "func %s %s(%s) %s {\n",
		methodPrefix,
		method.GetGoName(),
		strings.Join(funcSignature, ", "),
		method.GetGoReturns(),
	)

	if len(typ) > 0 {
		cReciever := fmt.Sprintf("c%s", typ)
		fmt.Fprintf(w, "%s := C.WGPU%s(unsafe.Pointer(%s.ref))\n\n", cReciever, typ, receiver)
		cCallArgs = append(cCallArgs, cReciever)
	}

	// Argument Conversion
	for _, a := range method.GetArgs() {
		cVar := "c" + capitalize(a.GoName)

		if a.Optional || spec.isObject(a.Type) {
			cVar = "p" + capitalize(a.GoName)
		}

		if spec.isStruct(a.Type) {
			writeConvertStruct(w, cVar, a.GoName, a, true, spec)
		} else {
			writeConvertGoToC(w, cVar, a.GoName, a, true, spec)
		}

		if (a.Annotation == "const*" || a.Annotation == "*") && !a.Optional && a.GoType != "[]byte" {
			cVar = "&" + cVar
		}

		cCallArgs = append(cCallArgs, cVar)
	}

	cCall := fmt.Sprintf("C.%s(%s)", method.GetCName(), strings.Join(cCallArgs, ", "))

	// Return
	writeReturn(w, method, cCall, spec)

	// End
	fmt.Fprintf(w, "}\n\n")
}

func writeReturn(w io.Writer, method Function, cCall string, spec *DawnSpec) {
	fmt.Fprintf(w, "// Call and return\n")
	switch {
	// no return
	case method.GetCReturns() == "":
		fmt.Fprintf(w, "%s\n", cCall)

	// pointer
	case strings.HasPrefix(method.GetGoReturns(), "*"):
		resType := strings.TrimPrefix(method.GetGoReturns(), "*")
		fmt.Fprintf(w, "return &%s{ref: uintptr(%s)}\n", resType, cCall)
	//FIXME
	case method.GetGoReturns() == "(*Buffer, error)":
		// Special case for async or complex returns identified in extractObjects
		fmt.Fprintf(w, "_ = %s // TODO: Implement async/error logic\n", cCall)
		fmt.Fprintf(w, "return nil, nil\n")

	// To be removed
	case method.GetReturns() == "future":
		fmt.Fprintf(w, "return %s{Id: uint64(%s.id)}\n", method.GetGoReturns(), cCall)

	// Bool -> Convert WGPUBool to bool
	case method.GetReturns() == "bool":
		fmt.Fprintf(w, "return %s(%s != 0)\n", method.GetGoReturns(), cCall)

	// Object
	case spec.isObject(method.GetCReturns()):
		fmt.Fprintf(w, "return %s{ref: uintptr(unsafe.Pointer(%s))}\n", method.GetGoReturns(), cCall)

	// Any other basic type
	default:
		// It's a basic type (int, bool, etc.)
		fmt.Fprintf(w, "return %s(%s)\n", method.GetGoReturns(), cCall)
	}
}

func writeConvertGoToC(w io.Writer, varName, argName string, arg Member, init bool, spec *DawnSpec) {
	switch {
	// Object
	case spec.isObject(arg.Type):
		obj := spec.getObject(arg.Type)
		if init {
			fmt.Fprintf(w, "%s := C.%s(unsafe.Pointer(%s.ref))\n", varName, obj.CName, argName)
		} else {
			fmt.Fprintf(w, "%s = C.%s(unsafe.Pointer(%s.ref))\n", varName, obj.CName, argName)
		}

	// Enum
	case spec.isEnum(arg.Type):
		enum := spec.getEnum(arg.Type)
		if init {
			fmt.Fprintf(w, "%s := C.%s(%s)\n", varName, enum.CName, argName)
		} else {
			fmt.Fprintf(w, "%s = C.%s(%s)\n", varName, enum.CName, argName)
		}
	// Struct
	case spec.isStruct(arg.Type):
		writeConvertStruct(w, varName, argName, arg, false, spec)

	// String View
	case arg.Type == "string view":
		tempVar := strings.Replace(varName, ".", "", -1) + "Str"
		fmt.Fprintf(w, "%s := C.CString(%s)\n", tempVar, argName)
		fmt.Fprintf(w, "defer C.free(unsafe.Pointer(%s))\n", tempVar)
		if init {
			fmt.Fprintf(w, "var %s C.%s\n", varName, arg.CType)
		}
		fmt.Fprintf(w, "%s.data = %s\n", varName, tempVar)
		fmt.Fprintf(w, "%s.length = C.size_t(len(%s))\n", varName, argName)

	// Slices
	case arg.GoType == "[]byte":
		if init {
			fmt.Fprintf(w, "var %s unsafe.Pointer\n", varName)
		}
		fmt.Fprintf(w, "if len(%s) >0 {\n", argName)
		fmt.Fprintf(w, "%s = unsafe.Pointer(&%s[0])\n", argName, varName)
		fmt.Fprintf(w, "}\n")
	// Bool
	case arg.CType == "bool":
		if init {
			fmt.Fprintf(w, "%s := boolToWGPUBool(%s)", varName, argName)
		} else {
			fmt.Fprintf(w, "%s = boolToWGPUBool(%s)", varName, argName)
		}

	// Anything else
	default:
		if init {
			fmt.Fprintf(w, "%s := C.%s(%s)\n", varName, arg.CType, argName)
		} else {
			fmt.Fprintf(w, "%s = C.%s(%s)\n", varName, arg.CType, argName)
		}
	}
}

func writeConvertStruct(w io.Writer, varName, argName string, arg Member, init bool, spec *DawnSpec) {
	structDef := spec.getStruct(arg.Type)
	isNullable := arg.Optional || spec.isObject(arg.Type)

	// if optional check for nil
	if isNullable {
		fmt.Fprintf(w, "if %s != nil {\n", argName)
	}

	if init {
		fmt.Fprintf(w, "var %s C.%s\n", varName, arg.CType)
	}

	for _, member := range structDef.Members {
		memberVar := fmt.Sprintf("%s.%s", varName, member.CName)
		memberArg := fmt.Sprintf("%s.%s", argName, member.GoName)
		writeConvertGoToC(w, memberVar, memberArg, member, false, spec)
	}

	if isNullable {
		fmt.Fprintf(w, "}\n")
	}
}

func writeCallbackHandler(w io.Writer, method Function, spec *DawnSpec) {

	var callback Entry
	for _, arg := range method.GetArgs() {
		if spec.isCallbackInfo(arg.Type) {
			callbackInfo := spec.getCallbackInfo(arg.Type)
			for _, member := range callbackInfo.Members {
				if member.Name == "callback" {
					callback = spec.getCallback(member.Type)
				}
			}
		}
	}

	funcName := fmt.Sprintf("go%sHandler", callback.GoName)
	fmt.Fprintf(w, "//export %s\n", funcName)

	var args []string

	for _, arg := range callback.Args {

		typ := fmt.Sprintf("C.%s", arg.CType)
		args = append(args, arg.CName+" "+typ)
	}

	args = append(args, "userData1 unsafe.Pointer", "userData2 unsafe.Pointer")

	// Signature
	fmt.Fprintf(w, "func %s(%s){\n", funcName, strings.Join(args, ","))

	// Body
	fmt.Fprintf(w, "handleID := uintptr(userData1)\n")
	fmt.Fprintf(w, "if handleID == 0{\n")
	fmt.Fprintf(w, "return\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "handle := cgo.Handle(handleID)\n")
	fmt.Fprintf(w, "defer handle.Delete()\n")
	fmt.Fprintf(w, "fn := handle.Value().(%s)\n", callback.GoName)
	fmt.Fprintf(w, "var message string\n")
	fmt.Fprintf(w, "if cMessage.data != nil && cMessage.length >0 {\n")
	fmt.Fprintf(w, "message = C.GoStringN(cMessage.data, C.int(cMessage.length))\n")
	fmt.Fprintf(w, "}\n")

	// Call callback
	fmt.Fprintf(w, "fn(\n")
	for _, arg := range callback.Args {
		if arg.Name == "message" {
			fmt.Fprintf(w, "message,\n")
		} else {
			fmt.Fprintf(w, "%s(%s),\n", arg.GoType, arg.CName)
		}
	}
	fmt.Fprintf(w, ")\n")

	// End
	fmt.Fprintf(w, "}\n")
}

func writeCallbacks(path string, spec *DawnSpec) {
	w, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "// CODE GENERATED. DO NOT EDIT\n")
	fmt.Fprint(w, "#include <webgpu/webgpu.h>\n")
	fmt.Fprintf(w, "#include \"_cgo_export.h\"\n\n")

	for _, f := range spec.CallBacks {
		if _, ok := SkipMethod[f.Name]; ok {
			continue
		}
		signature := callbackCSignature(f)
		goCallbackHandler := fmt.Sprintf("go%sHandler", f.GoName)

		var args []string
		for _, a := range f.Args {
			args = append(args, a.CName)
		}

		args = append(args, "userData1, userData2")

		fmt.Fprintf(w, "%s{\n", signature)
		fmt.Fprintf(w, "\t%s(%s);\n", goCallbackHandler, strings.Join(args, ","))
		fmt.Fprintf(w, "}\n")
	}
}

func callbackCSignature(f Entry) string {
	callbackName := fmt.Sprintf("cgo_callback_%s", f.GoName)

	var args []string
	for _, a := range f.Args {
		args = append(args, a.CType+" "+a.CName)
	}

	args = append(args, "void *userData1", "void *userData2")

	return fmt.Sprintf("void %s(%s)", callbackName, strings.Join(args, ", "))
}

func writePreable(w io.Writer) {
	fmt.Fprintf(w, "// CODE GENERATED. DO NOT EDIT\n")
	fmt.Fprintf(w, "//go:generate go run ./cmd/wrapper/. \n")
	fmt.Fprintf(w, "package wgpu\n\n")
}

func writeCGoPreamble(w io.Writer, callbacks SortedMap) {
	fmt.Fprintf(w, "/*\n")
	fmt.Fprintf(w, "#include <stdio.h>\n")
	fmt.Fprintf(w, "#include <stdlib.h>\n\n")
	fmt.Fprintf(w, "#cgo CFLAGS: -I./include\n")
	fmt.Fprintf(w, "#cgo LDFLAGS: -L./lib -lwebgpu_dawn -framework Metal -framework IOKit -framework QuartzCore -framework Foundation -framework IOSurface -lc++\n")
	fmt.Fprintf(w, "#include <webgpu/webgpu.h>\n\n")

	for _, cb := range callbacks.Iter() {
		if _, ok := SkipMethod[cb.Name]; ok {
			continue
		}
		fmt.Fprintf(w, "extern %s;\n", callbackCSignature(cb))
	}

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

func CType(in string) string {
	switch in {
	case "bool":
		return "WGPUBool"
	case "uint8_t":
		return "uint8_t"
	case "uint16_t":
		return "uint16_t"
	case "uint32_t":
		return "uint32_t"
	case "uint64_t":
		return "uint64_t"
	case "int8_t":
		return "int8_t"
	case "int16_t":
		return "int16_t"
	case "int32_t":
		return "int32_t"
	case "int64_t":
		return "int64_t"
	case "size_t":
		return "size_t"
	case "float":
		return "float"
	case "double":
		return "double"
	case "void *":
		return ""
	case "void const *":
		return ""
	default:
		return "WGPU" + GoName(in)
	}
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// func (i *Instance) RequestAdapter(options *RequestAdapterOptions, callbackInfo RequestAdapterCallbackInfo) Future {
// 	cInstance := C.WGPUInstance(unsafe.Pointer(i.ref))
// 	// Convert options to C.WGPURequestAdapterOptions
// 	var pOptions *C.WGPURequestAdapterOptions
// 	if options != nil {
// 		var cOptions C.WGPURequestAdapterOptions
// 		cOptions.forceFallbackAdapter = boolToWGPUBool(options.ForceFallbackAdapter)
// 		if options.CompatibleSurface != nil {
// 			cOptions.compatibleSurface = C.WGPUSurface(unsafe.Pointer(options.CompatibleSurface.ref))
// 		}
// 		pOptions = &cOptions
// 	}

// 	handle := cgo.NewHandle(callbackInfo.Callback)

// 	// Convert callbackInfo to C.WGPURequestAdapterCallbackInfo
// 	var cCallbackInfo C.WGPURequestAdapterCallbackInfo
// 	cCallbackInfo.mode = C.WGPUCallbackMode(callbackInfo.Mode)
// 	cCallbackInfo.callback = C.WGPURequestAdapterCallback(C.cgo_callback_RequestAdapterCallback)
// 	cCallbackInfo.userdata1 = unsafe.Pointer(handle)
// 	cCallbackInfo.userdata2 = nil
// 	return Future{Id: uint64(C.wgpuInstanceRequestAdapter(cInstance, pOptions, cCallbackInfo).id)}
// }

// //export goRequestAdapterCallbackHandler
// func goRequestAdapterCallbackHandler(status C.WGPURequestAdapterStatus, adapter C.WGPUAdapter, message C.WGPUStringView, userData1, userData2 unsafe.Pointer) {
// 	handleID := uintptr(userData1)
// 	if handleID == 0 {
// 		return
// 	}

// 	handle := cgo.Handle(handleID)
// 	defer handle.Delete()

// 	// THE FIX: Use the named type defined in your package
// 	// Instead of: .(func(RequestAdapterStatus, *Adapter, string))
// 	fn := handle.Value().(RequestAdapterCallback)

// 	// Safety check for the string (WGPUStringView is not null-terminated)
// 	// var message string
// 	// if cMessage.data != nil && cMessage.length > 0 {
// 	// 	message = C.GoStringN(cMessage.data, C.int(cMessage.length))
// 	// }

// 	// Call the function
// 	fn(
// 		RequestAdapterStatus(status),
// 		&Adapter{ref: uintptr(unsafe.Pointer(adapter))},
// 		"message",
// 	)
// }
