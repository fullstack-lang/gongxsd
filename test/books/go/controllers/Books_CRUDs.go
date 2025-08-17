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
var __Books__dummysDeclaration__ models.Books
var __Books_time__dummyDeclaration time.Duration

var mutexBooks sync.Mutex

// An BooksID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBooks updateBooks deleteBooks
type BooksID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BooksInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBooks updateBooks
type BooksInput struct {
	// The Books to submit or modify
	// in: body
	Books *orm.BooksAPI
}

// GetBookss
//
// swagger:route GET /bookss bookss getBookss
//
// # Get all bookss
//
// Responses:
// default: genericError
//
//	200: booksDBResponse
func (controller *Controller) GetBookss(c *gin.Context) {

	// source slice
	var booksDBs []orm.BooksDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBookss", "Name", stackPath)
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
	db := backRepo.BackRepoBooks.GetDB()

	_, err := db.Find(&booksDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	booksAPIs := make([]orm.BooksAPI, 0)

	// for each books, update fields from the database nullable fields
	for idx := range booksDBs {
		booksDB := &booksDBs[idx]
		_ = booksDB
		var booksAPI orm.BooksAPI

		// insertion point for updating fields
		booksAPI.ID = booksDB.ID
		booksDB.CopyBasicFieldsToBooks_WOP(&booksAPI.Books_WOP)
		booksAPI.BooksPointersEncoding = booksDB.BooksPointersEncoding
		booksAPIs = append(booksAPIs, booksAPI)
	}

	c.JSON(http.StatusOK, booksAPIs)
}

// PostBooks
//
// swagger:route POST /bookss bookss postBooks
//
// Creates a books
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBooks(c *gin.Context) {

	mutexBooks.Lock()
	defer mutexBooks.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBookss", "Name", stackPath)
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
	db := backRepo.BackRepoBooks.GetDB()

	// Validate input
	var input orm.BooksAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create books
	booksDB := orm.BooksDB{}
	booksDB.BooksPointersEncoding = input.BooksPointersEncoding
	booksDB.CopyBasicFieldsFromBooks_WOP(&input.Books_WOP)

	_, err = db.Create(&booksDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBooks.CheckoutPhaseOneInstance(&booksDB)
	books := backRepo.BackRepoBooks.Map_BooksDBID_BooksPtr[booksDB.ID]

	if books != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), books)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, booksDB)
}

// GetBooks
//
// swagger:route GET /bookss/{ID} bookss getBooks
//
// Gets the details for a books.
//
// Responses:
// default: genericError
//
//	200: booksDBResponse
func (controller *Controller) GetBooks(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBooks", "Name", stackPath)
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
	db := backRepo.BackRepoBooks.GetDB()

	// Get booksDB in DB
	var booksDB orm.BooksDB
	if _, err := db.First(&booksDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var booksAPI orm.BooksAPI
	booksAPI.ID = booksDB.ID
	booksAPI.BooksPointersEncoding = booksDB.BooksPointersEncoding
	booksDB.CopyBasicFieldsToBooks_WOP(&booksAPI.Books_WOP)

	c.JSON(http.StatusOK, booksAPI)
}

// UpdateBooks
//
// swagger:route PATCH /bookss/{ID} bookss updateBooks
//
// # Update a books
//
// Responses:
// default: genericError
//
//	200: booksDBResponse
func (controller *Controller) UpdateBooks(c *gin.Context) {

	mutexBooks.Lock()
	defer mutexBooks.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBooks", "Name", stackPath)
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
	db := backRepo.BackRepoBooks.GetDB()

	// Validate input
	var input orm.BooksAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var booksDB orm.BooksDB

	// fetch the books
	_, err := db.First(&booksDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	booksDB.CopyBasicFieldsFromBooks_WOP(&input.Books_WOP)
	booksDB.BooksPointersEncoding = input.BooksPointersEncoding

	db, _ = db.Model(&booksDB)
	_, err = db.Updates(&booksDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	booksNew := new(models.Books)
	booksDB.CopyBasicFieldsToBooks(booksNew)

	// redeem pointers
	booksDB.DecodePointers(backRepo, booksNew)

	// get stage instance from DB instance, and call callback function
	booksOld := backRepo.BackRepoBooks.Map_BooksDBID_BooksPtr[booksDB.ID]
	if booksOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), booksOld, booksNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the booksDB
	c.JSON(http.StatusOK, booksDB)
}

// DeleteBooks
//
// swagger:route DELETE /bookss/{ID} bookss deleteBooks
//
// # Delete a books
//
// default: genericError
//
//	200: booksDBResponse
func (controller *Controller) DeleteBooks(c *gin.Context) {

	mutexBooks.Lock()
	defer mutexBooks.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBooks", "Name", stackPath)
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
	db := backRepo.BackRepoBooks.GetDB()

	// Get model if exist
	var booksDB orm.BooksDB
	if _, err := db.First(&booksDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&booksDB)

	// get an instance (not staged) from DB instance, and call callback function
	booksDeleted := new(models.Books)
	booksDB.CopyBasicFieldsToBooks(booksDeleted)

	// get stage instance from DB instance, and call callback function
	booksStaged := backRepo.BackRepoBooks.Map_BooksDBID_BooksPtr[booksDB.ID]
	if booksStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), booksStaged, booksDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
