package split

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// Definimos la estructura para los posts
type Post struct {
	Text      string `json:"text"`
	Code      string `json:"code"`
	CreatedAt int64  `json:"created_at"`
}

// Definimos la estructura para la data
type Data struct {
	Posts []Post `json:"posts"`
}

func Split() {
	// Leemos el archivo JSON
	file, err := os.Open("posts.json")
	if err != nil {
		log.Fatalf("Error abriendo el archivo: %v", err)
	}
	defer file.Close()

	// Leemos el contenido del archivo
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error leyendo el archivo: %v", err)
	}

	// Parseamos el JSON
	var result interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Fatalf("Error parseando el JSON: %v", err)
	}

	var data Data

	timestampStr := "1722510804"
	// timestampStr := "1583929121"

	// Convertir la cadena a un entero
	limitTime, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		fmt.Println("Error converting timestamp:", err)
		return
	}

	// Extraemos la información
	extractText(result, &data, limitTime, "")

	fmt.Printf("Se extrajeron %d posts\n", len(data.Posts))

	// Guardamos los datos extraídos en un nuevo archivo JSON
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error creando el JSON de salida: %v", err)
	}

	if err := ioutil.WriteFile("extracted_posts.json", output, 0644); err != nil {
		log.Fatalf("Error guardando el archivo: %v", err)
	}

	fmt.Println("Datos extraídos y guardados en extracted_posts.json")
}

// Función recursiva para extraer valores de texto
func extractText(value interface{}, data *Data, limitTime int64, code string) {
	switch v := value.(type) {
	case map[string]interface{}:
		// Extraer el code si está presente
		if id, ok := v["code"].(string); ok {
			code = id // Actualizamos la variable code
			fmt.Println(code)
		}
		if text, ok := v["text"].(string); ok {
			// Verificar si created_at está presente y es mayor o igual a limitTime
			if createdAt, ok := v["created_at"].(float64); ok {
				if int64(createdAt) >= limitTime {
					// Agregar el post a la lista
					data.Posts = append(data.Posts, Post{Text: text, Code: code, CreatedAt: int64(createdAt)})
				}
			}
		}
		// Continuar explorando los valores del mapa
		for _, value := range v {
			extractText(value, data, limitTime, code)
		}
	case []interface{}:
		// Si es un array, iterar sobre los elementos
		for _, item := range v {
			// Continuar la extracción recursiva
			extractText(item, data, limitTime, code)
		}
	}
}
