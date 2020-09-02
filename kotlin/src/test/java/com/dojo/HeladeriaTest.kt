package com.dojo

import org.junit.jupiter.api.Assertions
import org.junit.jupiter.api.Test

import com.dojo.Heladeria
import com.dojo.Person

class HeladeriaTest {

    @Test
    fun `Testeo servir cono cuando le paso frutilla y limon`() {

        val heladeria = Heladeria("frutilla", "limon")
        val helado = heladeria.servirHelado("frutilla", "limon")

        Assertions.assertEquals("cono", helado.tipo)
        Assertions.assertEquals("frutilla", helado.gustos[0])
        Assertions.assertEquals("limon", helado.gustos[1])
    }

    @Test
    fun `Testeo servir cono cuando le paso un sabor que la heladeria no tiene`() {
        Assertions.assertThrows(Exception::class.java, {
            val heladeria = Heladeria("frutilla", "limon")
            val helado = heladeria.servirHelado("frutilla", "kinotos al whisky")
        })
    }
}