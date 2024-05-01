package models

type TemplateData struct {
	CSRFToken string
	Data      map[string]interface{}
	Error     string
	Flash     string
	FloatMap  map[string]float32
	IntMap    map[string]int
	StringMap map[string]string
	Warning   string
}
