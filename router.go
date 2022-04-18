package gee

import (
	"Gee/Gee"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]Gee.HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func newrouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]Gee.HandlerFunc),
	}
}
func parsepattern(pattern string) []string { //只允许'*'
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addrouter(pattern string, method string, handler Gee.HandlerFunc) {
	parts := parsepattern(pattern)
	key := method + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
}
