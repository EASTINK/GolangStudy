package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func linestr(key string) string {
	reg, err := http.Get(`https://apis.tianapi.com/zaoan/index?key=` + key)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reg.Body.Close()

	formData := make(map[string]interface{})
	json.NewDecoder(reg.Body).Decode(&formData)
	//for key, value := range formData {
	return formData["result"].(map[string]interface{})["content"].(string)
}
func main() {

	//}
}
