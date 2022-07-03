
## Leader Policies:

### Task Deliver Policy:
- Hay que tener en cuenta si un **Worker** se cae
- Reparto **Tasks**
- Despues las agrupo en una lista en otro lado
- El **Match** que se demore mucho, asumo que el worker se cayo, y lo vuelvo a poner en la lista de **Tasks** pendientes, sin sacarlo de donde estaba, para porsi me responden
- Cuando un **Worker** responde a un **Task** me quedo con el 1ro que respondio.
- Patron **Pipe** con los Workers

## Worker Policies:

### Election Request
- Ring Algorithm sobre el DHT ( Pag 346, 3rd Tanembaum)

### Descubrimiento de Actual Lider
- Se le pregunta el Lider al sucesor, cuando entra al anillo

### Resto de su Funcion:
- Correr Juego
- Enviar Resultados al Lider

### Pseoudocodigo de parte del flujo:
```go
    while true {
            Subscribe this Worker to --> Leader
        if error {
            count ++

            if count = limit{
                Throw Ascenso_a_Lider_Policy (and wait for it)
            }
            else {
                whait unos ms
            } 
        }
    }   
```

## Data Manager
- Falta la interfaz del DHT. y pensar un poco


## Informator Policies

### Searcher Policy
- Recopila Info y hace tablas etc.

### Announcer Policy
- Le da informacion de los torneos actuales a todo el que se lo pida