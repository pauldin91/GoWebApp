package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	DataMap   map[string]interface{}
	CsrfToken string
	Flash     string
	Warnings  string
	Error     string
}
