# todo_go

## API de la app

- _item_: Representa un item todo. Se va a implementar usando struct.
- _List_: Representa una lista de _items_. Implementado con un slice de _item_.

## Métodos de la API

- _Complete_: Marca un _item_ como completado.
- _Add_: Crea un _item_ y lo agrega a una _List_.
- _Delete_: Elimina un _item_ de un _List_.
- _Save_: Guarda la _List_ de _item_ en un archivo _JSON_.
- _Get_: Obtiene una _List_ de un archivo -JSON_.

## Command Line Tool

- Cuando se ejecute sin argumentos, se listan todos los elementos
- Al ejecutarse con dos argumentos, se concatenan y se agrega a la lista como un nuevo item

## TODO_FILENAME

Por default, `todo_go` utiliza el archivo `.todo.json` para almacenar las tareas, pero en el directorio donde se ejecuta el comando, por lo que al final vamos a tener un montón de archivos con diversas tareas, y seguramente no es eso lo que queremos.

Para tener un archivo fijo, hay que exportar la variable `TODO_FILENAME` en el archivo `.bashrc`, o `.zshrc`, dependiendo del shell que se utilice.

```sh
# bash
echo 'export TODO_FILENAME="$HOME/.todo.json"' >> ~/.bashrc
# zsh
echo 'export TODO_FILENAME="$HOME/.todo.json"' >> ~/.zshrc
```

## TODO

- [x] Agregar bandera `-d` para eliminar tareas
- [ ] Agregar bandera `-v` para modo `verbose`, mostrando fecha de creación, etc.
- [ ] Agregar bandera `-t` para modo `tiny`, para que no muestre tareas completadas
- [ ] Actualizar la función de ayuda personalizada.
- [ ] Incluir nuevas funciones de prueba para las nuevas banderas
- [ ] Actualizar las pruebas para usar la variable `TODO_FILENAME`.
- [ ] Actualizar la función `getTask` para aceptar varias líneas. Cada línea representa una tarea
