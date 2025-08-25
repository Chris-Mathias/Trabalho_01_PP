all: compile
    
compile:
	go build paralelismo_quebra_senha.go
	go build sequencial_quebra_senha.go

clean:
	rm -rf sequencial_quebra_senha
	rm -rf paralelismo_quebra_senha
	rm -rf resultado_paralelo.csv
	rm -rf resultado_sequencial.csv

limpa:
	clear


