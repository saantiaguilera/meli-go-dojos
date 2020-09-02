package com.dojo

/*
Hacer una heladeria. La heladeria tiene que poder servir helados de dos tipos:
- Cono (2 sabores)
- Cuarto helado (4 sabores)

La heladeria, tiene que poder inicializarse con un listado de sabores que tiene.
 */

 
 
 
 
 

/*
public publico por defalut
internal para modulos es publico, afuera privado.
private privado.
protected existe se comporta igual.
NO EXISTE package private
 */

class Person(var name: String, val surname: String?) {

    var hije: Person? = null

    open fun completeName(): String {
        
        return "$name $surname"
    } 
}
/*
public class Pajaro {

    private final String nombre;

    public Pajaro(String nombre) {
        this.nombre = nombre;
    }
}
*/