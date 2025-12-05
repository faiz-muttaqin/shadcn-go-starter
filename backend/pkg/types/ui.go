package types

import (
	"strings"
)

type UIOptions struct {
	Type         string
	Visible      bool
	Visibility   bool // toggle visibility / visibility control
	Creatable    bool
	Editable     bool
	Filterable   bool
	Sortable     bool
	Passwordable bool
	Selection    string
}

func ParseUIOptions(tag string) UIOptions {
	opts := UIOptions{
		Visible:    false,
		Visibility: false,
		Sortable:   false,
		Filterable: false,
		Editable:   false,
	}

	if tag == "" {
		return opts
	}

	parts := strings.Split(tag, ";")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		// Case: key:value
		if strings.Contains(p, ":") {
			kv := strings.SplitN(p, ":", 2)
			key := strings.TrimSpace(kv[0])
			val := strings.TrimSpace(kv[1])

			switch key {
			case "type":
				opts.Type = val
			case "selection":
				opts.Selection = val
			}
			continue
		}

		// Case: flag-only
		switch p {
		case "visible":
			opts.Visible = true
		case "hidden":
			opts.Visible = false
		case "visibility":
			opts.Visibility = true
		case "creatable":
			opts.Creatable = true
		case "editable":
			opts.Editable = true
		case "filterable":
			opts.Filterable = true
		case "sortable":
			opts.Sortable = true
		case "passwordable":
			opts.Passwordable = true
		}
	}

	return opts
}
