package com.dojo

class Person(val name: String, val surname: String) {

    fun completeName(): String {
        return "$name $surname"
    } 

}