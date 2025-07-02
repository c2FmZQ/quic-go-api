package main

import (
	"fmt"
	"github.com/c2FmZQ/quic-go-api"
	"reflect"
	"strings"
)

var types = []struct {
	name string
	typ  any
}{
	{"Transport", &quic.Transport{}},
	{"Listener", &quic.Listener{}},
	{"EarlyListener", &quic.EarlyListener{}},
	{"Conn", &quic.Conn{}},
	{"SendStream", &quic.SendStream{}},
	{"ReceiveStream", &quic.ReceiveStream{}},
	{"Stream", &quic.Stream{}},
	{"Path", &quic.Path{}},
	{"Dial", quic.Dial},
	{"DialEarly", quic.DialEarly},
	{"DialAddr", quic.DialAddr},
	{"DialAddrEarly", quic.DialAddrEarly},
	{"Listen", quic.Listen},
	{"ListenEarly", quic.ListenEarly},
	{"ListenAddr", quic.ListenAddr},
	{"ListenAddrEarly", quic.ListenAddrEarly},
}

func main() {
	typeToName := make(map[string]string)

	for _, t := range types {
		v := reflect.ValueOf(t.typ)
		typeToName[v.Type().String()] = t.name
	}
	for _, t := range types {
		v := reflect.ValueOf(t.typ)
		if v.Kind() != reflect.Pointer {
			continue
		}
		fmt.Printf("// %s is an auto-generated interface for [quic.%s]\n", t.name, t.name)
		fmt.Printf("type %s interface {\n", t.name)

		vt := v.Type()

		for i := 0; i < vt.NumMethod(); i++ {
			m := vt.Method(i)
			var inArgs []string
			for j := 1; j < m.Type.NumIn(); j++ {
				name := m.Type.In(j).String()
				if n, ok := typeToName[name]; ok {
					name = n
				}
				inArgs = append(inArgs, name)
			}
			var outArgs []string
			for j := 0; j < m.Type.NumOut(); j++ {
				name := m.Type.Out(j).String()
				if n, ok := typeToName[name]; ok {
					name = n
				}
				outArgs = append(outArgs, name)
			}
			fmt.Printf("\t%s(%s)", m.Name, strings.Join(inArgs, ", "))
			if len(outArgs) > 0 {
				if len(outArgs) > 1 {
					fmt.Printf(" (%s)", strings.Join(outArgs, ", "))
				} else {
					fmt.Printf(" %s", outArgs[0])
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("}\n\n")
	}

	writeFunc := func(receiver, funcName string, rt reflect.Type) {
		var inArgs []string
		for i := 0; i < rt.NumIn(); i++ {
			if i == 0 && receiver != "" {
				continue
			}
			name := rt.In(i).String()
			if n, ok := typeToName[name]; ok {
				name = n
			}
			inArgs = append(inArgs, fmt.Sprintf("a%d %s", i, name))
		}
		var outArgs []string
		var outArgs2 []string
		var retNames []string
		var wrappers []string
		var vars []string
		for i := 0; i < rt.NumOut(); i++ {
			name := rt.Out(i).String()
			tmpName := fmt.Sprintf("t%d", i)
			retName := fmt.Sprintf("r%d", i)
			if n, ok := typeToName[name]; ok {
				vars = append(vars, fmt.Sprintf("var %s %s", tmpName, name))
				name = n
				wrappers = append(wrappers, fmt.Sprintf("if %s != nil {\n\t\t%s = &%sWrapper{Base:%s}\n\t}", tmpName, retName, name, tmpName))
				retName = tmpName
			}
			outArgs = append(outArgs, fmt.Sprintf("r%d %s", i, name))
			outArgs2 = append(outArgs2, name)
			retNames = append(retNames, retName)
		}
		if receiver != "" {
			fmt.Printf("func (w *%s) %s(%s)", receiver, funcName, strings.Join(inArgs, ", "))
		} else {
			fmt.Printf("func %s(%s)", funcName, strings.Join(inArgs, ", "))
		}
		if len(outArgs) > 0 {
			if len(wrappers) > 0 {
				fmt.Printf(" (%s)", strings.Join(outArgs, ", "))
			} else if len(outArgs2) > 1 {
				fmt.Printf(" (%s)", strings.Join(outArgs2, ", "))
			} else {
				fmt.Printf(" %s", outArgs2[0])
			}
		}
		fmt.Printf(" {\n")
		var argNames []string
		for i := 0; i < rt.NumIn(); i++ {
			if i == 0 && receiver != "" {
				continue
			}
			name := rt.In(i).String()
			if n, ok := typeToName[name]; ok {
				argNames = append(argNames, fmt.Sprintf("a%d.(*%sWrapper).Base", i, n))
				continue
			}
			argNames = append(argNames, fmt.Sprintf("a%d", i))
		}
		prefix := "w.Base"
		if receiver == "" {
			prefix = "quic"
		}
		if len(outArgs) > 0 {
			if len(wrappers) > 0 {
				for _, v := range vars {
					fmt.Printf("\t%s\n", v)
				}
				fmt.Printf("\t%s = %s.%s(%s)\n", strings.Join(retNames, ", "), prefix, funcName, strings.Join(argNames, ", "))
				for _, w := range wrappers {
					fmt.Printf("\t%s\n", w)
				}
				fmt.Printf("\treturn\n")
			} else {
				fmt.Printf("\treturn %s.%s(%s)\n", prefix, funcName, strings.Join(argNames, ", "))
			}
		} else {
			fmt.Printf("\t%s.%s(%s)\n", prefix, funcName, strings.Join(argNames, ", "))
		}
		fmt.Printf("}\n\n")
	}

	for _, t := range types {
		v := reflect.ValueOf(t.typ)
		if v.Kind() != reflect.Pointer {
			continue
		}
		structName := t.name + "Wrapper"
		fmt.Printf("var _ %s = (*%s)(nil)\n\n", t.name, structName)
		fmt.Printf("// %s is an auto-generated wrapper for [quic.%s]\n", structName, t.name)
		fmt.Printf("type %s struct {\n", structName)

		vt := v.Type()

		fmt.Printf("\tBase %s\n", vt)
		fmt.Printf("}\n\n")

		for i := 0; i < vt.NumMethod(); i++ {
			m := vt.Method(i)
			writeFunc(structName, m.Name, m.Type)
		}
	}

	for _, t := range types {
		v := reflect.ValueOf(t.typ)
		if v.Kind() != reflect.Func {
			continue
		}
		fmt.Printf("// %s is an auto-generated wrapper for [quic.%s]\n", t.name, t.name)
		writeFunc("", t.name, v.Type())
	}
}
