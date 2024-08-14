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
var __BookType__dummysDeclaration__ models.BookType
var __BookType_time__dummyDeclaration time.Duration

var mutexBookType sync.Mutex

// An BookTypeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBookType updateBookType deleteBookType
type BookTypeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BookTypeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBookType updateBookType
type BookTypeInput struct {
	// The BookType to submit or modify
	// in: body
	BookType *orm.BookTypeAPI
}

// GetBookTypes
//
// swagger:route GET /booktypes booktypes getBookTypes
//
// # Get all booktypes
//
// Responses:
// default: genericError
//
//	200: booktypeDBResponse
func (controller *Controller) GetBookTypes(c *gin.Context) {

	// source slice
	var booktypeDBs []orm.BookTypeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBookTypes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookType.GetDB()

	query := db.Find(&booktypeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	booktypeAPIs := make([]orm.BookTypeAPI, 0)

	// for each booktype, update fields from the database nullable fields
	for idx := range booktypeDBs {
		booktypeDB := &booktypeDBs[idx]
		_ = booktypeDB
		var booktypeAPI orm.BookTypeAPI

		// insertion point for updating fields
		booktypeAPI.ID = booktypeDB.ID
		booktypeDB.CopyBasicFieldsToBookType_WOP(&booktypeAPI.BookType_WOP)
		booktypeAPI.BookTypePointersEncoding = booktypeDB.BookTypePointersEncoding
		booktypeAPIs = append(booktypeAPIs, booktypeAPI)
	}

	c.JSON(http.StatusOK, booktypeAPIs)
}

// PostBookType
//
// swagger:route POST /booktypes booktypes postBookType
//
// Creates a booktype
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBookType(c *gin.Context) {

	mutexBookType.Lock()
	defer mutexBookType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBookTypes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookType.GetDB()

	// Validate input
	var input orm.BookTypeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create booktype
	booktypeDB := orm.BookTypeDB{}
	booktypeDB.BookTypePointersEncoding = input.BookTypePointersEncoding
	booktypeDB.CopyBasicFieldsFromBookType_WOP(&input.BookType_WOP)

	query := db.Create(&booktypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBookType.CheckoutPhaseOneInstance(&booktypeDB)
	booktype := backRepo.BackRepoBookType.Map_BookTypeDBID_BookTypePtr[booktypeDB.ID]

	if booktype != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), booktype)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, booktypeDB)
}

// GetBookType
//
// swagger:route GET /booktypes/{ID} booktypes getBookType
//
// Gets the details for a booktype.
//
// Responses:
// default: genericError
//
//	200: booktypeDBResponse
func (controller *Controller) GetBookType(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBookType", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookType.GetDB()

	// Get booktypeDB in DB
	var booktypeDB orm.BookTypeDB
	if err := db.First(&booktypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var booktypeAPI orm.BookTypeAPI
	booktypeAPI.ID = booktypeDB.ID
	booktypeAPI.BookTypePointersEncoding = booktypeDB.BookTypePointersEncoding
	booktypeDB.CopyBasicFieldsToBookType_WOP(&booktypeAPI.BookType_WOP)

	c.JSON(http.StatusOK, booktypeAPI)
}

// UpdateBookType
//
// swagger:route PATCH /booktypes/{ID} booktypes updateBookType
//
// # Update a booktype
//
// Responses:
// default: genericError
//
//	200: booktypeDBResponse
func (controller *Controller) UpdateBookType(c *gin.Context) {

	mutexBookType.Lock()
	defer mutexBookType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBookType", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookType.GetDB()

	// Validate input
	var input orm.BookTypeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var booktypeDB orm.BookTypeDB

	// fetch the booktype
	query := db.First(&booktypeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	booktypeDB.CopyBasicFieldsFromBookType_WOP(&input.BookType_WOP)
	booktypeDB.BookTypePointersEncoding = input.BookTypePointersEncoding

	query = db.Model(&booktypeDB).Updates(booktypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	booktypeNew := new(models.BookType)
	booktypeDB.CopyBasicFieldsToBookType(booktypeNew)

	// redeem pointers
	booktypeDB.DecodePointers(backRepo, booktypeNew)

	// get stage instance from DB instance, and call callback function
	booktypeOld := backRepo.BackRepoBookType.Map_BookTypeDBID_BookTypePtr[booktypeDB.ID]
	if booktypeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), booktypeOld, booktypeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the booktypeDB
	c.JSON(http.StatusOK, booktypeDB)
}

// DeleteBookType
//
// swagger:route DELETE /booktypes/{ID} booktypes deleteBookType
//
// # Delete a booktype
//
// default: genericError
//
//	200: booktypeDBResponse
func (controller *Controller) DeleteBookType(c *gin.Context) {

	mutexBookType.Lock()
	defer mutexBookType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBookType", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookType.GetDB()

	// Get model if exist
	var booktypeDB orm.BookTypeDB
	if err := db.First(&booktypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&booktypeDB)

	// get an instance (not staged) from DB instance, and call callback function
	booktypeDeleted := new(models.BookType)
	booktypeDB.CopyBasicFieldsToBookType(booktypeDeleted)

	// get stage instance from DB instance, and call callback function
	booktypeStaged := backRepo.BackRepoBookType.Map_BookTypeDBID_BookTypePtr[booktypeDB.ID]
	if booktypeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), booktypeStaged, booktypeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
