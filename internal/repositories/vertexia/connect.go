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

	var strBuilder strings.Builder
	for _, line := range strings.Split(dataStr, "\n") {
		if strings.HasPrefix(line, "data: ") {

			var dat map[string]interface{}
			if err := json.Unmarshal([]byte(strings.TrimPrefix(line, "data: ")), &dat); err != nil {
				panic(err)
			}

			byt := []byte(strings.TrimPrefix(line, "data: "))

			if err := json.Unmarshal(byt, &dat); err != nil {
				panic(err)
			}

			if candidates, ok := dat["candidates"].([]interface{}); ok {
				if len(candidates) > 0 {
					if content, ok := candidates[0].(map[string]interface{})["content"].(map[string]interface{}); ok {
						if parts, ok := content["parts"].([]interface{}); ok {
							if len(parts) > 0 {
								if text, ok := parts[0].(map[string]interface{})["text"].(string); ok {
									strBuilder.WriteString(text)
								} else {
									fmt.Println("Campo 'text' no encontrado o no es una cadena")
								}
							} else {
								fmt.Println("Lista 'parts' vacía o no es una lista")
							}
						} else {
							fmt.Println("Campo 'parts' no encontrado o no es una lista")
						}
					} else {
						fmt.Println("Campo 'content' no encontrado o no es un mapa")
					}
				} else {
					fmt.Println("Lista 'candidates' vacía o no es una lista")
				}
			} else {
				fmt.Println("Campo 'candidates' no encontrado o no es una lista")
			}

		}
	}

	return strBuilder.String()
}

func MakeRequests() error {
	token := "ya29.a0AeDClZBhS75h5CKsgms2qJb5Wb91JRLrKzBADr04dikTW2zlklh_UFxCPDVi6wOpxNydA8i30jVnUkM6RlSsJzq1UG8-oXneJdDUiKRK8hqz4Yj47srxH8QtngesLnVssIfLphK__FsDjq8tuzmcpeVV0Mc9_Zwvc8LxSe1rs_N2UFwaCgYKATwSARESFQHGX2MixQwq7s0zjLaU9uBluvvweQ0182"
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

	response := formatJSON(body)
	fmt.Println("Respuesta formateada:", response)
	return nil
}
