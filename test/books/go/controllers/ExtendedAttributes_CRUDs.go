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
var __ExtendedAttributes__dummysDeclaration__ models.ExtendedAttributes
var __ExtendedAttributes_time__dummyDeclaration time.Duration

var mutexExtendedAttributes sync.Mutex

// An ExtendedAttributesID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getExtendedAttributes updateExtendedAttributes deleteExtendedAttributes
type ExtendedAttributesID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ExtendedAttributesInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postExtendedAttributes updateExtendedAttributes
type ExtendedAttributesInput struct {
	// The ExtendedAttributes to submit or modify
	// in: body
	ExtendedAttributes *orm.ExtendedAttributesAPI
}

// GetExtendedAttributess
//
// swagger:route GET /extendedattributess extendedattributess getExtendedAttributess
//
// # Get all extendedattributess
//
// Responses:
// default: genericError
//
//	200: extendedattributesDBResponse
func (controller *Controller) GetExtendedAttributess(c *gin.Context) {

	// source slice
	var extendedattributesDBs []orm.ExtendedAttributesDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetExtendedAttributess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoExtendedAttributes.GetDB()

	query := db.Find(&extendedattributesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	extendedattributesAPIs := make([]orm.ExtendedAttributesAPI, 0)

	// for each extendedattributes, update fields from the database nullable fields
	for idx := range extendedattributesDBs {
		extendedattributesDB := &extendedattributesDBs[idx]
		_ = extendedattributesDB
		var extendedattributesAPI orm.ExtendedAttributesAPI

		// insertion point for updating fields
		extendedattributesAPI.ID = extendedattributesDB.ID
		extendedattributesDB.CopyBasicFieldsToExtendedAttributes_WOP(&extendedattributesAPI.ExtendedAttributes_WOP)
		extendedattributesAPI.ExtendedAttributesPointersEncoding = extendedattributesDB.ExtendedAttributesPointersEncoding
		extendedattributesAPIs = append(extendedattributesAPIs, extendedattributesAPI)
	}

	c.JSON(http.StatusOK, extendedattributesAPIs)
}

// PostExtendedAttributes
//
// swagger:route POST /extendedattributess extendedattributess postExtendedAttributes
//
// Creates a extendedattributes
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostExtendedAttributes(c *gin.Context) {

	mutexExtendedAttributes.Lock()
	defer mutexExtendedAttributes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostExtendedAttributess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoExtendedAttributes.GetDB()

	// Validate input
	var input orm.ExtendedAttributesAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create extendedattributes
	extendedattributesDB := orm.ExtendedAttributesDB{}
	extendedattributesDB.ExtendedAttributesPointersEncoding = input.ExtendedAttributesPointersEncoding
	extendedattributesDB.CopyBasicFieldsFromExtendedAttributes_WOP(&input.ExtendedAttributes_WOP)

	query := db.Create(&extendedattributesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoExtendedAttributes.CheckoutPhaseOneInstance(&extendedattributesDB)
	extendedattributes := backRepo.BackRepoExtendedAttributes.Map_ExtendedAttributesDBID_ExtendedAttributesPtr[extendedattributesDB.ID]

	if extendedattributes != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), extendedattributes)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, extendedattributesDB)
}

// GetExtendedAttributes
//
// swagger:route GET /extendedattributess/{ID} extendedattributess getExtendedAttributes
//
// Gets the details for a extendedattributes.
//
// Responses:
// default: genericError
//
//	200: extendedattributesDBResponse
func (controller *Controller) GetExtendedAttributes(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetExtendedAttributes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoExtendedAttributes.GetDB()

	// Get extendedattributesDB in DB
	var extendedattributesDB orm.ExtendedAttributesDB
	if err := db.First(&extendedattributesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var extendedattributesAPI orm.ExtendedAttributesAPI
	extendedattributesAPI.ID = extendedattributesDB.ID
	extendedattributesAPI.ExtendedAttributesPointersEncoding = extendedattributesDB.ExtendedAttributesPointersEncoding
	extendedattributesDB.CopyBasicFieldsToExtendedAttributes_WOP(&extendedattributesAPI.ExtendedAttributes_WOP)

	c.JSON(http.StatusOK, extendedattributesAPI)
}

// UpdateExtendedAttributes
//
// swagger:route PATCH /extendedattributess/{ID} extendedattributess updateExtendedAttributes
//
// # Update a extendedattributes
//
// Responses:
// default: genericError
//
//	200: extendedattributesDBResponse
func (controller *Controller) UpdateExtendedAttributes(c *gin.Context) {

	mutexExtendedAttributes.Lock()
	defer mutexExtendedAttributes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateExtendedAttributes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoExtendedAttributes.GetDB()

	// Validate input
	var input orm.ExtendedAttributesAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var extendedattributesDB orm.ExtendedAttributesDB

	// fetch the extendedattributes
	query := db.First(&extendedattributesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	extendedattributesDB.CopyBasicFieldsFromExtendedAttributes_WOP(&input.ExtendedAttributes_WOP)
	extendedattributesDB.ExtendedAttributesPointersEncoding = input.ExtendedAttributesPointersEncoding

	query = db.Model(&extendedattributesDB).Updates(extendedattributesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	extendedattributesNew := new(models.ExtendedAttributes)
	extendedattributesDB.CopyBasicFieldsToExtendedAttributes(extendedattributesNew)

	// redeem pointers
	extendedattributesDB.DecodePointers(backRepo, extendedattributesNew)

	// get stage instance from DB instance, and call callback function
	extendedattributesOld := backRepo.BackRepoExtendedAttributes.Map_ExtendedAttributesDBID_ExtendedAttributesPtr[extendedattributesDB.ID]
	if extendedattributesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), extendedattributesOld, extendedattributesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the extendedattributesDB
	c.JSON(http.StatusOK, extendedattributesDB)
}

// DeleteExtendedAttributes
//
// swagger:route DELETE /extendedattributess/{ID} extendedattributess deleteExtendedAttributes
//
// # Delete a extendedattributes
//
// default: genericError
//
//	200: extendedattributesDBResponse
func (controller *Controller) DeleteExtendedAttributes(c *gin.Context) {

	mutexExtendedAttributes.Lock()
	defer mutexExtendedAttributes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteExtendedAttributes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoExtendedAttributes.GetDB()

	// Get model if exist
	var extendedattributesDB orm.ExtendedAttributesDB
	if err := db.First(&extendedattributesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&extendedattributesDB)

	// get an instance (not staged) from DB instance, and call callback function
	extendedattributesDeleted := new(models.ExtendedAttributes)
	extendedattributesDB.CopyBasicFieldsToExtendedAttributes(extendedattributesDeleted)

	// get stage instance from DB instance, and call callback function
	extendedattributesStaged := backRepo.BackRepoExtendedAttributes.Map_ExtendedAttributesDBID_ExtendedAttributesPtr[extendedattributesDB.ID]
	if extendedattributesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), extendedattributesStaged, extendedattributesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
