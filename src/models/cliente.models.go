package models

import "api-eco360/config"

type Cliente struct {
	Codigo  int
	Nome    string
	CpfCnpj string
	RgIe    string
}

func BuscaTodosClientes() []Cliente {
	db, err := config.Conexao()

	if err != nil {
		panic(err)
	}

	selecionaTodosClientes, err := db.Query(`
		select
    		codigo,
    		nome,
    		cpfcnpj,
    		rgie
		from trecclientegeral
		order by codigo`)
	if err != nil {
		panic(err.Error())
	}

	c := Cliente{}
	clientes := []Cliente{}

	for selecionaTodosClientes.Next() {
		var codigo int
		var nome, cpfcnpj, rgie string

		err = selecionaTodosClientes.Scan(&codigo, &nome, &cpfcnpj, &rgie)

		if err != nil {
			panic(err.Error())
		}

		c.Codigo = codigo
		c.Nome = nome
		c.CpfCnpj = cpfcnpj
		c.RgIe = rgie

		clientes = append(clientes, c)
	}

	defer db.Close()

	return clientes
}

func ClienteCodigo(codigo string) []Cliente {
	db, err := config.Conexao()

	if err != nil {
		panic(err)
	}

	selecionaTodosClientes, err := db.Query(`
		select
    		codigo,
    		nome,
    		cpfcnpj,
    		rgie
		from trecclientegeral
		where codigo = ?`, codigo)
	if err != nil {
		panic(err.Error())
	}

	c := Cliente{}
	clientes := []Cliente{}

	for selecionaTodosClientes.Next() {
		var codigo int
		var nome, cpfcnpj, rgie string

		err = selecionaTodosClientes.Scan(&codigo, &nome, &cpfcnpj, &rgie)

		if err != nil {
			panic(err.Error())
		}

		c.Codigo = codigo
		c.Nome = nome
		c.CpfCnpj = cpfcnpj
		c.RgIe = rgie

		clientes = append(clientes, c)
	}

	defer db.Close()

	return clientes
}
