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
var __BookDetailsGroup__dummysDeclaration__ models.BookDetailsGroup
var __BookDetailsGroup_time__dummyDeclaration time.Duration

var mutexBookDetailsGroup sync.Mutex

// An BookDetailsGroupID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBookDetailsGroup updateBookDetailsGroup deleteBookDetailsGroup
type BookDetailsGroupID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BookDetailsGroupInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBookDetailsGroup updateBookDetailsGroup
type BookDetailsGroupInput struct {
	// The BookDetailsGroup to submit or modify
	// in: body
	BookDetailsGroup *orm.BookDetailsGroupAPI
}

// GetBookDetailsGroups
//
// swagger:route GET /bookdetailsgroups bookdetailsgroups getBookDetailsGroups
//
// # Get all bookdetailsgroups
//
// Responses:
// default: genericError
//
//	200: bookdetailsgroupDBResponse
func (controller *Controller) GetBookDetailsGroups(c *gin.Context) {

	// source slice
	var bookdetailsgroupDBs []orm.BookDetailsGroupDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBookDetailsGroups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookDetailsGroup.GetDB()

	query := db.Find(&bookdetailsgroupDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bookdetailsgroupAPIs := make([]orm.BookDetailsGroupAPI, 0)

	// for each bookdetailsgroup, update fields from the database nullable fields
	for idx := range bookdetailsgroupDBs {
		bookdetailsgroupDB := &bookdetailsgroupDBs[idx]
		_ = bookdetailsgroupDB
		var bookdetailsgroupAPI orm.BookDetailsGroupAPI

		// insertion point for updating fields
		bookdetailsgroupAPI.ID = bookdetailsgroupDB.ID
		bookdetailsgroupDB.CopyBasicFieldsToBookDetailsGroup_WOP(&bookdetailsgroupAPI.BookDetailsGroup_WOP)
		bookdetailsgroupAPI.BookDetailsGroupPointersEncoding = bookdetailsgroupDB.BookDetailsGroupPointersEncoding
		bookdetailsgroupAPIs = append(bookdetailsgroupAPIs, bookdetailsgroupAPI)
	}

	c.JSON(http.StatusOK, bookdetailsgroupAPIs)
}

// PostBookDetailsGroup
//
// swagger:route POST /bookdetailsgroups bookdetailsgroups postBookDetailsGroup
//
// Creates a bookdetailsgroup
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBookDetailsGroup(c *gin.Context) {

	mutexBookDetailsGroup.Lock()
	defer mutexBookDetailsGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBookDetailsGroups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookDetailsGroup.GetDB()

	// Validate input
	var input orm.BookDetailsGroupAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bookdetailsgroup
	bookdetailsgroupDB := orm.BookDetailsGroupDB{}
	bookdetailsgroupDB.BookDetailsGroupPointersEncoding = input.BookDetailsGroupPointersEncoding
	bookdetailsgroupDB.CopyBasicFieldsFromBookDetailsGroup_WOP(&input.BookDetailsGroup_WOP)

	query := db.Create(&bookdetailsgroupDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBookDetailsGroup.CheckoutPhaseOneInstance(&bookdetailsgroupDB)
	bookdetailsgroup := backRepo.BackRepoBookDetailsGroup.Map_BookDetailsGroupDBID_BookDetailsGroupPtr[bookdetailsgroupDB.ID]

	if bookdetailsgroup != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), bookdetailsgroup)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bookdetailsgroupDB)
}

// GetBookDetailsGroup
//
// swagger:route GET /bookdetailsgroups/{ID} bookdetailsgroups getBookDetailsGroup
//
// Gets the details for a bookdetailsgroup.
//
// Responses:
// default: genericError
//
//	200: bookdetailsgroupDBResponse
func (controller *Controller) GetBookDetailsGroup(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBookDetailsGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookDetailsGroup.GetDB()

	// Get bookdetailsgroupDB in DB
	var bookdetailsgroupDB orm.BookDetailsGroupDB
	if err := db.First(&bookdetailsgroupDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bookdetailsgroupAPI orm.BookDetailsGroupAPI
	bookdetailsgroupAPI.ID = bookdetailsgroupDB.ID
	bookdetailsgroupAPI.BookDetailsGroupPointersEncoding = bookdetailsgroupDB.BookDetailsGroupPointersEncoding
	bookdetailsgroupDB.CopyBasicFieldsToBookDetailsGroup_WOP(&bookdetailsgroupAPI.BookDetailsGroup_WOP)

	c.JSON(http.StatusOK, bookdetailsgroupAPI)
}

// UpdateBookDetailsGroup
//
// swagger:route PATCH /bookdetailsgroups/{ID} bookdetailsgroups updateBookDetailsGroup
//
// # Update a bookdetailsgroup
//
// Responses:
// default: genericError
//
//	200: bookdetailsgroupDBResponse
func (controller *Controller) UpdateBookDetailsGroup(c *gin.Context) {

	mutexBookDetailsGroup.Lock()
	defer mutexBookDetailsGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBookDetailsGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookDetailsGroup.GetDB()

	// Validate input
	var input orm.BookDetailsGroupAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bookdetailsgroupDB orm.BookDetailsGroupDB

	// fetch the bookdetailsgroup
	query := db.First(&bookdetailsgroupDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bookdetailsgroupDB.CopyBasicFieldsFromBookDetailsGroup_WOP(&input.BookDetailsGroup_WOP)
	bookdetailsgroupDB.BookDetailsGroupPointersEncoding = input.BookDetailsGroupPointersEncoding

	query = db.Model(&bookdetailsgroupDB).Updates(bookdetailsgroupDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bookdetailsgroupNew := new(models.BookDetailsGroup)
	bookdetailsgroupDB.CopyBasicFieldsToBookDetailsGroup(bookdetailsgroupNew)

	// redeem pointers
	bookdetailsgroupDB.DecodePointers(backRepo, bookdetailsgroupNew)

	// get stage instance from DB instance, and call callback function
	bookdetailsgroupOld := backRepo.BackRepoBookDetailsGroup.Map_BookDetailsGroupDBID_BookDetailsGroupPtr[bookdetailsgroupDB.ID]
	if bookdetailsgroupOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bookdetailsgroupOld, bookdetailsgroupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bookdetailsgroupDB
	c.JSON(http.StatusOK, bookdetailsgroupDB)
}

// DeleteBookDetailsGroup
//
// swagger:route DELETE /bookdetailsgroups/{ID} bookdetailsgroups deleteBookDetailsGroup
//
// # Delete a bookdetailsgroup
//
// default: genericError
//
//	200: bookdetailsgroupDBResponse
func (controller *Controller) DeleteBookDetailsGroup(c *gin.Context) {

	mutexBookDetailsGroup.Lock()
	defer mutexBookDetailsGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBookDetailsGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookDetailsGroup.GetDB()

	// Get model if exist
	var bookdetailsgroupDB orm.BookDetailsGroupDB
	if err := db.First(&bookdetailsgroupDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&bookdetailsgroupDB)

	// get an instance (not staged) from DB instance, and call callback function
	bookdetailsgroupDeleted := new(models.BookDetailsGroup)
	bookdetailsgroupDB.CopyBasicFieldsToBookDetailsGroup(bookdetailsgroupDeleted)

	// get stage instance from DB instance, and call callback function
	bookdetailsgroupStaged := backRepo.BackRepoBookDetailsGroup.Map_BookDetailsGroupDBID_BookDetailsGroupPtr[bookdetailsgroupDB.ID]
	if bookdetailsgroupStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bookdetailsgroupStaged, bookdetailsgroupDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
