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
