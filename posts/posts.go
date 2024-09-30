package posts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Posts(hashtag string) {
	url := "https://www.instagram.com/api/v1/tags/web_info/?tag_name=" + hashtag

	// Crear la solicitud HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creando la solicitud:", err)
		return
	}

	// Agregar encabezados necesarios
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.7")
	req.Header.Set("cookie", os.Getenv("COOKIE"))
	req.Header.Set("referer", "https://www.instagram.com/explore/tags/"+hashtag+"/")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="8", "Chromium";v="126", "Brave";v="126"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "Linux")
	req.Header.Set("sec-ch-ua-platform-version", "6.10.3")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	req.Header.Set("x-asbd-id", os.Getenv("x-asbd-id"))
	req.Header.Set("x-csrftoken", os.Getenv("x-csrftoken"))
	req.Header.Set("x-ig-app-id", os.Getenv("x-ig-app-id"))
	req.Header.Set("x-ig-www-claim", os.Getenv("x-ig-www-claim"))
	req.Header.Set("x-requested-with", "XMLHttpRequest")

	// Hacer la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error haciendo la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error leyendo la respuesta:", err)
		return
	}

	// Convertir la respuesta a JSON e imprimirla
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	// Guardar el resultado en un archivo JSON
	file, err := os.Create("posts.json")
	if err != nil {
		fmt.Println("Error creando el archivo JSON:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Formato bonito
	if err := encoder.Encode(result); err != nil {
		fmt.Println("Error escribiendo en el archivo JSON:", err)
		return
	}

	fmt.Println("Datos guardados en 'posts.json'")
}
