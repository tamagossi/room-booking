package models

type TemplateData struct {
	CSRFToken string
	Data      map[string]interface{}
	Error     string
	Flash     string
	FLoatMap  map[string]float64
	IntMap    map[string]int
	StringMap map[string]string
	Warning   string
}
