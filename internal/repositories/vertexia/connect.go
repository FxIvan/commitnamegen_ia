package vertexai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func formatJSON(data []byte) string {
	// Convertir `data` a string para manejar líneas individualmente
	dataStr := string(data)

	// Filtrar solo las líneas que contienen JSON (eliminando el prefijo "data: ")
	var jsonBlocks []string
	for _, line := range strings.Split(dataStr, "\n") {
		if strings.HasPrefix(line, "data: ") {
			// Agregar cada bloque JSON como un elemento en el array
			jsonBlocks = append(jsonBlocks, strings.TrimPrefix(line, "data: "))
		}
	}

	// Combinar los bloques en un solo array JSON
	cleanedData := "[" + strings.Join(jsonBlocks, ",") + "]"

	// Formatear el JSON
	var out bytes.Buffer
	err := json.Indent(&out, []byte(cleanedData), "", "  ")
	if err != nil {
		fmt.Println("Error al formatear JSON:", err)
		return cleanedData // Retornar JSON sin formato en caso de error
	}

	return out.String()
}

func MakeRequests() error {
	token := "ya29.a0AeDClZAh1-QgW0c8H76VJr5fJZTMWh0L1JgyUmvJ94jTmJyJOi0kiJBWOIki1F9Zxk5Ie3MThYgaKEQCNUi08XKz8pY0752vcQoAQ_IeWShAH3PGPcFU0t8QTOJTc1nKcNOlEGsYuro-HPPLpb_a6_HbjJ6E4WE_Ut90EODQjTHWs0kaCgYKAfwSARESFQHGX2Mio-OtfBbfdQ2MgBmtQEWjCg0182"
	url := "https://us-central1-aiplatform.googleapis.com/v1/projects/proyectia-440422/locations/us-central1/publishers/google/models/gemini-1.0-pro:streamGenerateContent?alt=sse"

	jsonData := `{
		"contents": {
			"role": "user",
			"parts": {
				"text": "Me das las buenas practicas para armar un commit?"
			}
		},
		"safety_settings": {
			"category": "HARM_CATEGORY_SEXUALLY_EXPLICIT",
			"threshold": "BLOCK_LOW_AND_ABOVE"
		},
		"generation_config": {
			"temperature": 0.2,
			"topP": 0.8,
			"topK": 40
		}
	}`

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return fmt.Errorf("error creando la solicitud: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error ejecutando la solicitud: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error leyendo el cuerpo de la respuesta: %v", err)
	}

	formattedData := formatJSON(body)
	fmt.Println(formattedData)
	return nil
}
