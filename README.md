# Golang Práticas 🏗️
Um projeto de estudo e prática com Go!

Este repositório contém exemplos e testes para aprender e aprimorar habilidades com a linguagem **Golang**.

---

## Créditos ao tutorial:
https://larien.gitbook.io/aprenda-go-com-testes/main

---

## 📌 Configuração do Projeto
Antes de começar, inicialize o módulo do projeto:

```sh
go mod init golang-praticas
```

---

## ✅ Rodando os Testes
Execute os testes recursivamente entre as pastas do projeto:

```sh
go test -v ./...
```

Forçar a reexecução sem cache:

```sh
go test -v -count=1 ./...
```

Executar os testes e verificar a cobertura de código:

```sh
go test -cover ./...
```

Para executar o benchmark, digite no terminal:

```sh
go test -bench . ./...

go test -bench . -benchmem ./...
```

---

## 🎯 O que está sendo abordado neste projeto?
✔️ Estruturação de projetos em Go  
✔️ Escrita e execução de testes  
✔️ Organização modular de código  
✔️ Boas práticas no desenvolvimento  

---

## 🚀 Contribuições
Se quiser contribuir com ideias, melhorias ou dúvidas, fique à vontade para abrir **issues** ou enviar um **pull request**!

📬 **Contato:** _adicione seu e-mail ou redes sociais aqui_
