package main

import (
	"fmt"
	"net/http"
)

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento.....")
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site:", site, "está com problemas. Status Code:", resp.StatusCode)

	}
}
