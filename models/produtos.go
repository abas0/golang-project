package models

import (
	"go_modules/db"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Titulo     string
	Genero     string
	Autor      string
	Preco      int
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}
	for selectDeTodosOsProdutos.Next() {
		var id, quantidade, preco int
		var titulo, genero, autor string

		err = selectDeTodosOsProdutos.Scan(&id, &titulo, &genero, &autor, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Titulo = titulo
		p.Genero = genero
		p.Autor = autor
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	//t.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
	return produtos
}

func CriaNovoProduto(titulo, genero string, autor string, preco int, quantidade int) {
	db := db.ConectaComBancoDeDados()
	insereDadosNoBanco, err := db.Prepare("insert into produtos(titulo, genero, autor, preco, quantidade) values(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(titulo, genero, autor, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id=?")
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
	defer db.Close()

}

func EditaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=?", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade, preco int
		var titulo, genero, autor string

		err = produtoDoBanco.Scan(&id, &titulo, &genero, &autor, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Titulo = titulo
		produtoParaAtualizar.Genero = genero
		produtoParaAtualizar.Autor = autor
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, titulo, genero, autor string, preco, quantidade int) {
	db := db.ConectaComBancoDeDados()

	AtualizaProduto, err := db.Prepare("update produtos set titulo=?, genero=?, autor=?, preco=?, quantidade=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(titulo, genero, autor, preco, quantidade, id)
	defer db.Close()
}
