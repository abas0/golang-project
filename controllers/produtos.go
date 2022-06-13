package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"go_modules/models"
)

var temp = template.Must(template.ParseGlob("template/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		titulo := r.FormValue("titulo")
		genero := r.FormValue("genero")
		autor := r.FormValue("autor")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		//precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		precoConvertidaParaInt, err := strconv.Atoi(preco)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}

		models.CriaNovoProduto(titulo, genero, autor, precoConvertidaParaInt, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		titulo := r.FormValue("titulo")
		genero := r.FormValue("genero")
		autor := r.FormValue("autor")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		precoConvertidaParaInt, err := strconv.Atoi(preco)
		if err != nil {
			log.Println("Erro na convesão do preço para int:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na convesão da quantidade para int:", err)
		}

		models.AtualizaProduto(idConvertidaParaInt, titulo, autor, genero, precoConvertidaParaInt, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
