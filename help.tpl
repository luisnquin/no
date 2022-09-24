Usage: 
   {{$v := offset .Name 6}}{{wrap .Name 3}} [... | --flag]
{{wrap .Description 3}}{{if len .Authors}}
   {{range .VisibleFlagCategories}}
   {{if .Name}}{{.Name}}
   {{end}}{{range .Flags}}{{.}}
   {{end}}{{end}}{{else}}{{if .VisibleFlags}}

   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{wrap $option.String 6}}{{end}}{{end}}{{end}}
