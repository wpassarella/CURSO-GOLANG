package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func envia() {

	for i := 0; i <= 0; i++ {

		//Carrega TX2
		dat, _ := ioutil.ReadFile("C:\\Arquivos\\Go\\src\\Curso-Golang\\DemoSaaS\\nfce-pr.tx2")
		fmt.Println("TX2 --> ", dat)

		//Monta form com dados da requisição
		form := url.Values{}
		form.Add("grupo", "edoc")
		form.Add("cnpj", "08187168000160")
		form.Add("arquivo", string(dat))
		form.Add("encode", "true")
		fmt.Println("Parâmetros --> ", form)

		//Monta a requisição
		req, _ := http.NewRequest("POST", "https://managersaashom.tecnospeed.com.br:7071/ManagerAPIWeb/nfce/envia",
			strings.NewReader(form.Encode()))
		fmt.Println("Requsição --> ", req)

		//Define autenticação
		req.SetBasicAuth("admin", "123mudar")
		cli := &http.Client{}
		fmt.Println("Autenticação --> ", cli)

		//Envia a requisição
		resp, _ := cli.Do(req)
		fmt.Println("Resposta --> ", resp)

		//Trata o retorno
		fmt.Println("Status:", resp.Status)
		fmt.Println("Cabeçalho:", resp.Header)

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Mensagem:", string(body))
		ioutil.WriteFile("C:\\Arquivos\\Go\\src\\Curso-Golang\\DemoSaaS\\retorno.tx2", body, 0644)
		time.Sleep(time.Second)
	}
}

func consulta() {

	//Carrega TX2
	/*fmt.Println("Digite a chave que será consultada: ")
	var ChaveConsulta int
	fmt.Scanf("%d", &ChaveConsulta)*/
	dat := "grupo%3Dedoc%0D%0Acnpj%3D08187168000160%0D%0Acampos%3Dsituacao%0D%0Afiltro%3Dchavenota%3D41180508187168000160656320000001181000005462"
	fmt.Println("TX2 --> ", dat)

	//Monta form com dados da requisição
	form := url.Values{}
	form.Add("grupo", "edoc")
	form.Add("cnpj", "08187168000160")
	form.Add("arquivo", string(dat))
	form.Add("encode", "true")
	fmt.Println("Parâmetros --> ", form)

	//Monta a requisição
	req, _ := http.NewRequest("GET", "https://managersaashom.tecnospeed.com.br:7071/ManagerAPIWeb/nfce/consulta",
		strings.NewReader(form.Encode()))
	fmt.Println("Requsição --> ", req)

	//Define autenticação
	req.SetBasicAuth("admin", "123mudar")
	cli := &http.Client{}
	fmt.Println("Autenticação --> ", cli)

	//Envia a requisição
	resp, _ := cli.Do(req)
	fmt.Println("Resposta --> ", resp)

	//Trata o retorno
	fmt.Println("Status:", resp.Status)
	fmt.Println("Cabeçalho:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Mensagem:", string(body))
	ioutil.WriteFile("C:\\Arquivos\\Go\\src\\Curso-Golang\\DemoSaaS\\retornoConsulta.tx2", body, 0644)
}

func status() {

	dat := "grupo%3Dedoc%0D%0Acnpj%3D08187168000160"
	form := url.Values{}
	form.Add("grupo", "edoc")
	form.Add("cnpj", "08187168000160")
	form.Add("arquivo", string(dat))
	form.Add("encode", "true")
	fmt.Println("Parâmetros --> ", form)

	//Monta a requisição
	req, _ := http.NewRequest("GET", "https://managersaashom.tecnospeed.com.br:7071/ManagerAPIWeb/nfce/status",
		strings.NewReader(form.Encode()))
	fmt.Println("Requsição --> ", req)

	//Define autenticação
	req.SetBasicAuth("admin", "123mudar")
	cli := &http.Client{}
	fmt.Println("Autenticação --> ", cli)

	//Envia a requisição
	resp, _ := cli.Do(req)
	fmt.Println("Resposta --> ", resp)

	//Trata o retorno
	fmt.Println("Status:", resp.Status)
	fmt.Println("Cabeçalho:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Mensagem:", string(body))
	ioutil.WriteFile("C:\\Arquivos\\Go\\src\\Curso-Golang\\DemoSaaS\\retornoConsultaStatus.tx2", body, 0644)
}

func resolve() {

	dat := "chavenota=41180508187168000160656320000001181000005462"
	form := url.Values{}
	form.Add("grupo", "edoc")
	form.Add("cnpj", "08187168000160")
	form.Add("arquivo", string(dat))
	form.Add("encode", "true")
	fmt.Println("Parâmetros --> ", form)

	//Monta a requisição
	req, _ := http.NewRequest("POST", "https://managersaashom.tecnospeed.com.br:7071/ManagerAPIWeb/nfce/resolve",
		strings.NewReader(form.Encode()))
	fmt.Println("Requsição --> ", req)

	//Define autenticação
	req.SetBasicAuth("admin", "123mudar")
	cli := &http.Client{}
	fmt.Println("Autenticação --> ", cli)

	//Envia a requisição
	resp, _ := cli.Do(req)
	fmt.Println("Resposta --> ", resp)

	//Trata o retorno
	fmt.Println("Status:", resp.Status)
	fmt.Println("Cabeçalho:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Mensagem:", string(body))
	ioutil.WriteFile("C:\\Arquivos\\Go\\src\\Curso-Golang\\DemoSaaS\\retornoresolve.tx2", body, 0644)
}

//Função Principal
func main() {
	fmt.Print("Demo SaaS Golang!!!\n\n")
	fmt.Println("Se deseja enviar digite 1: ")
	fmt.Println("Se deseja consultar digite 2: ")
	fmt.Println("Se deseja consultar Status digite 3: ")
	fmt.Println("Se deseja resolver a nota digite 4: ")
	var menu int
	fmt.Scanf("%d", &menu)
	if menu == 1 {
		envia()
	} else if menu == 2 {
		consulta()
	} else if menu == 3 {
		status()
	} else if menu == 4 {
		resolve()
	} else {
		fmt.Println("Nenhuma opção válida selecionada")
	}

}
