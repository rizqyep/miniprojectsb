package controllers

import (
	"mini-project-sanbercode-rizqyep/database"
	"mini-project-sanbercode-rizqyep/repository"
	"mini-project-sanbercode-rizqyep/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPerson(c *gin.Context) {
	var (
		result gin.H
	)

	err, persons := repository.GetAllPerson(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{"result": persons}
	}

	c.JSON(http.StatusOK, result)
}

func InsertPerson(c *gin.Context) {
	var person structs.Person

	err := c.ShouldBindJSON(&person)

	if err != nil {
		panic(err)
	}

	err = repository.InsertPerson(database.DbConnection, person)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": "Success Insert Person",
	})
}

func UpdatePerson(c *gin.Context) {
	var person structs.Person

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&person)

	if err != nil {
		panic(err)
	}

	person.ID = int64(id)

	err = repository.UpdatePerson(database.DbConnection, person)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success update person",
	})
}

func DeletePerson(c *gin.Context) {
	var person structs.Person
	id, err := strconv.Atoi(c.Param("id"))

	person.ID = int64(id)

	err = repository.DeletePerson(*database.DbConnection, person)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully Delete person!",
	})
}
