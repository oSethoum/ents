package ents

import (
	"strings"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
)

var (
	snake  = gen.Funcs["snake"].(func(string) string)
	_camel = gen.Funcs["camel"].(func(string) string)
	camel  = func(s string) string { return _camel(snake(s)) }
)

func init() {
	gen.Funcs["edge_field"] = edge_field
	gen.Funcs["get_name"] = get_name
	gen.Funcs["get_type"] = get_type
	gen.Funcs["go_ts"] = go_ts
	gen.Funcs["is_slice"] = is_slice
	gen.Funcs["id_type"] = id_type
	gen.Funcs["order_fields"] = order_fields
	gen.Funcs["select_fields"] = select_fields
	gen.Funcs["comparable"] = comparable
}

func get_name(f *load.Field) string {
	n := camel(f.Name)
	if strings.HasSuffix(n, "ID") {
		n = strings.TrimSuffix(n, "ID") + "Id"
	}
	return n
}

func get_type(t *field.TypeInfo) string {
	return go_ts(t.Type.String())
}

func go_ts(s string) string {
	slice := false
	if strings.HasPrefix(s, "[]") {
		slice = true
		s = strings.TrimPrefix(s, "[]")
	}
	for k, v := range gots {
		if strings.HasPrefix(s, k) {
			if slice {
				return v + "[]"
			}
			return v
		}
	}
	if slice {
		return s + "[]"
	}
	return s
}

func edge_field(e *load.Edge) bool {
	return e.Field != ""
}

func is_slice(f *load.Field) bool {
	return strings.HasPrefix(get_type(f.Info), "[]")
}

func id_type(s *load.Schema) string {
	for _, f := range s.Fields {
		if strings.ToLower(f.Name) == "id" {
			return get_type(f.Info)
		}
	}
	return "number"
}

func order_fields(s *load.Schema) string {
	fields := []string{}
	for _, f := range s.Fields {
		if orderable(f) {
			fields = append(fields, get_name(f))
		}
	}
	return "\"" + strings.Join(fields, "\" | \"") + "\""
}

func select_fields(s *load.Schema) string {
	fields := []string{}
	for _, f := range s.Fields {
		fields = append(fields, get_name(f))
	}
	return "\"" + strings.Join(fields, "\" | \"") + "\""
}

func comparable(f *load.Field) bool {
	return has_prefixes(extract_type(f), []string{
		"string",
		"int",
		"uint",
		"float",
		"time.Time",
	})
}

func orderable(f *load.Field) bool {
	return has_prefixes(extract_type(f), []string{
		"string",
		"int",
		"uint",
		"float",
		"time.Time",
		"bool",
	})
}

func extract_type(field *load.Field) string {
	if field.Info.Ident != "" {
		return field.Info.Ident
	}
	return field.Info.Type.String()
}
