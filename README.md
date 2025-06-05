para cada `TOKEN` em `TOKENS`:
    se `TOKEN` é numero:
        adiciona em `queue`

    senão:
        se a pilha esta vazia:
            adiciona na pilha
        
        senão:
            verifica se o operador atual tem maior precedencia que o operador do topo da pilha 
            remove o operador da pilha, adiciona na queue, e faz o push do atual na stack

            ao chegar no fim da equação faz o pop dos operadores restantes e adiciona na queue