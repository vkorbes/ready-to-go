# Ready to Go: Uma introdução prática à linguagem Go

# Variáveis, Valores & Tipos

- Go Playground
- Hello World
- Variáveis:
    - ":=" é utilizado pra criar novas variáveis, dentro de code blocks
    - "var" é utilizado onde quiser, usa-se pra package-level scope
    - "=" atribui valores a variáveis já existentes
- Tipos:
	- "Type is life" — Bill Kennedy
	- Tipos em Go são estáticos.
    - O tipo pode ser deduzido pelo compilador:
        - x := 10
        - var y = "a tia do batima"
    - Ou pode ser declarado especificamente:
        - var w string = "isso é uma string"
        - var z int = 15
    - Tipos de dados primitivos: int, string, bool, float64.
    - Tipos de dados compostos: slice, array, struct, map.
    - Zero value. Eles são:
        - ints: 0
        - floats: 0.0
        - booleans: false
        - strings: ""
        - pointers, functions, interfaces, slices, channels, maps: nil
    - Criando seu próprio tipo: type hotdog int → var b hotdog
    - Conversão: hotdog(x) ↔ int(y)
- Ponteiros:
	- &variável mostra o endereço de uma variável
	- *variável faz de-reference, mostra o valor que consta nesse endereço
	- YOLO: *&var
	- Exemplo: a := 0; b := &a; *b++
- Visualizando tudo isso:
	- package fmt
	- fmt.Print(), fmt.Println(), fmt.Printf()
	- Formatting verbs: %v, %T, %d, %b, etc.
	



# Fluxo de Controle

- Fluxo iterativo:
	- for: for init, condition, post { code }
	- "while": for condition { code }
	- for... ever: for { code }
	- break & continue
- Condicional:
	- Switch:
	- Switch, case, default. Não há fallthrough por padrão.
	- Pode-se usar ou não um switch expression (default==true).
- Expressão de inicialização.
- Type switches.
- If:
- if, else if, else if, else if, else.
- Expressão de inicialização.
- Operadores lógicos condicionais: &&, ||, !, etc.




# Agrupamentos de Dados

- Arrays: Não.
- Slices:
	- Declaração: slice := []type{a, b, c}
	- Acesso: slice[i]
	- len(slice)
	- for i, v := range slice { code }
	- append(slice, a, b, c)
	- append(slice, outraslice...)
	- slice[:], slice[a:], slice[:b], slice[a:b]
		- "a" é incluso, "b" não é.
	- É "fatiando" que se deleta um item de uma slice. E.g.:
		- x := append(x[:i], x[:i]...)
	- A saber, armadilhas:
		- out of range / make([]type, len, cap)
		- x := []int{...números}; y := append(x[:i], x[:i]...)
			- O append que declara "y" modifica o valor de "x"!
	- Multi-dimensionais: [][]type
- Maps:
	- key:value
	- Declaração: map[key]value{ key: value }
	- Acesso: map[key]
	- Key inexistente retorna 0. Para evitar confusão:
		- v, ok := map[key]
		- Na prática: if v, ok := map[key]; ok { code }
	- Range: for k, v := range map { code } ← Ordem aleatória!
	- delete(map, key)
- Structs:
	- Declaração: 
- Tipo: type batima struct { número int; texto string }
- Var: joker := batima{10, "vixe!"}
	- Acesso: struct.field → joker.texto
	- Pode-se ter structs dentro de struct. 
- Ou structs com slices. 
- Ou slices de structs. 
- Etc.
	- Structs anônimos: x := struct{ nome tipo }{ nome: valor }



# Funções (funcs, métodos, interfaces)
	
- Escreve na testa: func (receiver) identifier(parameters) (returns) { code }
- Em Go é tudo pass by value.
- Na prática:
	- Função simples
	- Função que aceita um argumento
	- Função com retorno
	- Função com múltiplos argumentos e múltiplos retornos
	- Função variádica
	- Passando um slice como argumento para uma função variádica
	- Função anônima
	- Função como expressão: f := func(p params){ code }
	- Retornando uma função
		- func f() func() int { [...]; return func() int{ return [int] } }
	- Callbacks: passando funções como argumentos
	- Closure
- Defer:
	- "Deixa pra última hora!"
	- A execução ocorrerá no momento em que a função da qual ela faz parte finalizar (devido a um return, ao fim do code block, ou pânico).
- Métodos:
	- Método é uma função anexada a um tipo.
	- func (receiver) identifier(parameters) (returns) { code }
	- Utilização: nome.método()
	- Exemplo: todo mundo do tipo "pessoa" tem o método oibomdia()
- Interfaces:
	- Interfaces permitem que valores tenham mais que um tipo
	- Toda interface exige x métodos
	- Todos os tipos com estes métodos implementam a interface
	- Assim temos acesso a toda a funcionalidade da interface
	- E pra que serve isso?
		- Story time: Todos os eletrodomésticos implementam a interface "plugue", e todas as hidroelétricas/termoelétricas/etc implementam a interface "tacar-a-energia-nos-fio-do-poste". Assim a gente consegue fazer com que qualquer produtora de energia se comunique, por assim dizer, com qualquer eletrodoméstico da sua casa, sem um lado ter a menor idéia de como o outro lado funciona.
		- Na prática: uma ferramenta que escreve na interface io.Writer pode gravar dados em qualquer coisa que implemente esta interface, independente do que seja. E.g. arquivos, conexões, etc.




# Aplicações práticas

- Já entendemos ponteiros, já entendemos métodos. Já temos o conhecimento necessário para começar a utilizar a standard library.
- JSON 
- marshal e unmarshal
- marshal/unmarshal vs. encoder/decoder
	- https://play.golang.org/p/_JvCOlK-H9
	- https://play.golang.org/p/l6wbuLu1NS
	- https://play.golang.org/p/Pgwr0O07aL
- Package sort
	- Utilizando os sorts pré-prontos
	    - sort.Strings: https://play.golang.org/p/Rs1NVwmg7h
    	- sort.Ints: https://play.golang.org/p/I2_vsHujZa
	- Criando nosso próprio sort
		- https://play.golang.org/p/KOIhAsE3OK
- Canais
	- Converge: https://github.com/ellenkorbes/go-aprenda-a-programar/blob/master/21_canais/07/todd/main.go
	- Diverge e converge: https://github.com/ellenkorbes/go-aprenda-a-programar/blob/master/21_canais/08/01/main.go
- bcrypt
	- Gerando hashes
	- Testando senhas contra hashes salvos




# Tour: Concorrência, Canais, io, net/http

- Concorrência
	- Fun facts:
- O primeiro CPU dual core "popular" veio em 2006
- Em 2007 o Google começou a criar a linguagem Go
- Go foi a primeira linguagem criada com multi-cores em mente
- Ou seja, Go tem uma abordagem única (e fácil!) para este tópico
	- goroutines: go func
	- waitgroups: sync.WaitGroup
		- Add(int)
		- Done
		- Wait
	runtime: runtime.NumCPU() & runtime.NumGoroutine()
- Canais
	- Tô com dor de cabeça depois eu penso.
	- Tem setinhas assim<-
	- Podem ser bidirecionais ou não
	- Canais bloqueiam!
	- Usando range em canais
	- Select: igual switch, mas diferente
- Comma ok
- net/http
	- Que sono!
	


# Resources

- Comunidade: 
    - Slack: https://invite.slack.golangbridge.org/
    - #brazil
    - #brazilian-go-studies, toda quinta às 22h!
        - gravações em: https://www.youtube.com/cesargimenes
        - exercícios em: https://github.com/crgimenes/Go-Hands-On
        - FB: https://www.facebook.com/gophers.br/
- Livros:
    - A Linguagem de Programação Go (Alan Donovan, Brian Kernighan): https://www.amazon.com.br/dp/8575225464/
    - Go em Ação (William Kennedy, Brian Ketelsen, Erik St. Martin): https://www.amazon.com.br/dp/8575225065/
    - Introdução à Linguagem Go (Caleb Doxsey): https://www.amazon.com.br/dp/8575224891/
- Em inglês: gobyexample.com, forum.golangbridge.org, stackoverflow.com/questions/tagged/go
- Documentação oficial:
    - golang.org
    - godoc.org
    - golang.org/doc/effective_go.html

