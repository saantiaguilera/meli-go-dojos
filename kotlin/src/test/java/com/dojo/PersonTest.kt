package com.dojo

import org.junit.jupiter.api.Assertions
import org.junit.jupiter.api.Test

import com.dojo.Person

class PersonTest {

    @Test
    fun `Test complete name includes name and surname`() {
        val person = Person("santiago", "aguilera")

        val completeName = person.completeName()

        Assertions.assertEquals("santiago aguilera", completeName)
    }

}