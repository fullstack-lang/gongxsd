// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxsd/test/books/go/models"
	"github.com/fullstack-lang/gongxsd/test/books/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __A_books__dummysDeclaration__ models.A_books
var __A_books_time__dummyDeclaration time.Duration

var mutexA_books sync.Mutex

// An A_booksID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_books updateA_books deleteA_books
type A_booksID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_booksInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_books updateA_books
type A_booksInput struct {
	// The A_books to submit or modify
	// in: body
	A_books *orm.A_booksAPI
}

// GetA_bookss
//
// swagger:route GET /a_bookss a_bookss getA_bookss
//
// # Get all a_bookss
//
// Responses:
// default: genericError
//
//	200: a_booksDBResponse
func (controller *Controller) GetA_bookss(c *gin.Context) {

	// source slice
	var a_booksDBs []orm.A_booksDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_bookss", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gongxsd/test/books/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoA_books.GetDB()

	_, err := db.Find(&a_booksDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_booksAPIs := make([]orm.A_booksAPI, 0)

	// for each a_books, update fields from the database nullable fields
	for idx := range a_booksDBs {
		a_booksDB := &a_booksDBs[idx]
		_ = a_booksDB
		var a_booksAPI orm.A_booksAPI

		// insertion point for updating fields
		a_booksAPI.ID = a_booksDB.ID
		a_booksDB.CopyBasicFieldsToA_books_WOP(&a_booksAPI.A_books_WOP)
		a_booksAPI.A_booksPointersEncoding = a_booksDB.A_booksPointersEncoding
		a_booksAPIs = append(a_booksAPIs, a_booksAPI)
	}

	c.JSON(http.StatusOK, a_booksAPIs)
}

// PostA_books
//
// swagger:route POST /a_bookss a_bookss postA_books
//
// Creates a a_books
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_books(c *gin.Context) {

	mutexA_books.Lock()
	defer mutexA_books.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_bookss", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gongxsd/test/books/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoA_books.GetDB()

	// Validate input
	var input orm.A_booksAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_books
	a_booksDB := orm.A_booksDB{}
	a_booksDB.A_booksPointersEncoding = input.A_booksPointersEncoding
	a_booksDB.CopyBasicFieldsFromA_books_WOP(&input.A_books_WOP)

	_, err = db.Create(&a_booksDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_books.CheckoutPhaseOneInstance(&a_booksDB)
	a_books := backRepo.BackRepoA_books.Map_A_booksDBID_A_booksPtr[a_booksDB.ID]

	if a_books != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_books)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_booksDB)
}

// GetA_books
//
// swagger:route GET /a_bookss/{ID} a_bookss getA_books
//
// Gets the details for a a_books.
//
// Responses:
// default: genericError
//
//	200: a_booksDBResponse
func (controller *Controller) GetA_books(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_books", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gongxsd/test/books/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoA_books.GetDB()

	// Get a_booksDB in DB
	var a_booksDB orm.A_booksDB
	if _, err := db.First(&a_booksDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_booksAPI orm.A_booksAPI
	a_booksAPI.ID = a_booksDB.ID
	a_booksAPI.A_booksPointersEncoding = a_booksDB.A_booksPointersEncoding
	a_booksDB.CopyBasicFieldsToA_books_WOP(&a_booksAPI.A_books_WOP)

	c.JSON(http.StatusOK, a_booksAPI)
}

// UpdateA_books
//
// swagger:route PATCH /a_bookss/{ID} a_bookss updateA_books
//
// # Update a a_books
//
// Responses:
// default: genericError
//
//	200: a_booksDBResponse
func (controller *Controller) UpdateA_books(c *gin.Context) {

	mutexA_books.Lock()
	defer mutexA_books.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_books", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gongxsd/test/books/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoA_books.GetDB()

	// Validate input
	var input orm.A_booksAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_booksDB orm.A_booksDB

	// fetch the a_books
	_, err := db.First(&a_booksDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_booksDB.CopyBasicFieldsFromA_books_WOP(&input.A_books_WOP)
	a_booksDB.A_booksPointersEncoding = input.A_booksPointersEncoding

	db, _ = db.Model(&a_booksDB)
	_, err = db.Updates(&a_booksDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_booksNew := new(models.A_books)
	a_booksDB.CopyBasicFieldsToA_books(a_booksNew)

	// redeem pointers
	a_booksDB.DecodePointers(backRepo, a_booksNew)

	// get stage instance from DB instance, and call callback function
	a_booksOld := backRepo.BackRepoA_books.Map_A_booksDBID_A_booksPtr[a_booksDB.ID]
	if a_booksOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_booksOld, a_booksNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_booksDB
	c.JSON(http.StatusOK, a_booksDB)
}

// DeleteA_books
//
// swagger:route DELETE /a_bookss/{ID} a_bookss deleteA_books
//
// # Delete a a_books
//
// default: genericError
//
//	200: a_booksDBResponse
func (controller *Controller) DeleteA_books(c *gin.Context) {

	mutexA_books.Lock()
	defer mutexA_books.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_books", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gongxsd/test/books/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoA_books.GetDB()

	// Get model if exist
	var a_booksDB orm.A_booksDB
	if _, err := db.First(&a_booksDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_booksDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_booksDeleted := new(models.A_books)
	a_booksDB.CopyBasicFieldsToA_books(a_booksDeleted)

	// get stage instance from DB instance, and call callback function
	a_booksStaged := backRepo.BackRepoA_books.Map_A_booksDBID_A_booksPtr[a_booksDB.ID]
	if a_booksStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_booksStaged, a_booksDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
