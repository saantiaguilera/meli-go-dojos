package com.dojo
/*
Hacer una heladeria. La heladeria tiene que poder servir helados de dos tipos:
- Cono (2 sabores)
- Cuarto helado (4 sabores)

La heladeria, tiene que poder inicializarse con un listado de sabores que tiene.
 */

 /*
 int -> Int
 float -> Float
 double -> Double
 boolean -> Boolean

 List<String>
 Array<String>
  */

class Helado(val tipo: String, vararg val gustos: String) {}


class Heladeria(vararg val gustos : String) {

  fun servirHelado(vararg gustosHelado: String): Helado {
      val gustosFiltrados = gustosHelado.filter { gusto ->
          gustos.contains(gusto) 
      }
      val cantidadDeGustosQueTengo = gustosFiltrados.size

      if (cantidadDeGustosQueTengo != gustosHelado.size) {
          throw Exception("no tengo alguno de tus gustos")
      }

      return when (gustos.size) { 
        2 -> Helado("cono", *gustos)
        4 -> Helado("cuarto", *gustos)
        else -> throw Exception("Debe elegir 2 o 4 gustos")
      }
  }
}

/*
public class Algo {

    private final List<extends String> lista;
}

setLista(List<String>)
*/