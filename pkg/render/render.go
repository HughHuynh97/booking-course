package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, fileName string) {
	filePath := "./templates/" + fileName
	parseTemplate, err := template.ParseFiles(filePath)

	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	err = parseTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		fmt.Printf("Error executing template: %v\n", err)
		return
	}
}
