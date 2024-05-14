package main

type Course struct {
	nome string
	sobrenome string
	idade int 
	profisao string
	
}

func main() {
	course := Course {
		nome: "Lougans",
		sobrenome: "Meneses",
		idade: 28,
		profisao: "Desenvolvedor",
	}

	println(course.nome + course.sobrenome)

	nome := "Lougans"
	println("Seja muito bem vindo " + nome)
	println("agora vai")
	
}