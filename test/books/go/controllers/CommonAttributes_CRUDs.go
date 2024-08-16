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
var __CommonAttributes__dummysDeclaration__ models.CommonAttributes
var __CommonAttributes_time__dummyDeclaration time.Duration

var mutexCommonAttributes sync.Mutex

// An CommonAttributesID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCommonAttributes updateCommonAttributes deleteCommonAttributes
type CommonAttributesID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CommonAttributesInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCommonAttributes updateCommonAttributes
type CommonAttributesInput struct {
	// The CommonAttributes to submit or modify
	// in: body
	CommonAttributes *orm.CommonAttributesAPI
}

// GetCommonAttributess
//
// swagger:route GET /commonattributess commonattributess getCommonAttributess
//
// # Get all commonattributess
//
// Responses:
// default: genericError
//
//	200: commonattributesDBResponse
func (controller *Controller) GetCommonAttributess(c *gin.Context) {

	// source slice
	var commonattributesDBs []orm.CommonAttributesDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCommonAttributess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCommonAttributes.GetDB()

	query := db.Find(&commonattributesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	commonattributesAPIs := make([]orm.CommonAttributesAPI, 0)

	// for each commonattributes, update fields from the database nullable fields
	for idx := range commonattributesDBs {
		commonattributesDB := &commonattributesDBs[idx]
		_ = commonattributesDB
		var commonattributesAPI orm.CommonAttributesAPI

		// insertion point for updating fields
		commonattributesAPI.ID = commonattributesDB.ID
		commonattributesDB.CopyBasicFieldsToCommonAttributes_WOP(&commonattributesAPI.CommonAttributes_WOP)
		commonattributesAPI.CommonAttributesPointersEncoding = commonattributesDB.CommonAttributesPointersEncoding
		commonattributesAPIs = append(commonattributesAPIs, commonattributesAPI)
	}

	c.JSON(http.StatusOK, commonattributesAPIs)
}

// PostCommonAttributes
//
// swagger:route POST /commonattributess commonattributess postCommonAttributes
//
// Creates a commonattributes
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCommonAttributes(c *gin.Context) {

	mutexCommonAttributes.Lock()
	defer mutexCommonAttributes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCommonAttributess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCommonAttributes.GetDB()

	// Validate input
	var input orm.CommonAttributesAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create commonattributes
	commonattributesDB := orm.CommonAttributesDB{}
	commonattributesDB.CommonAttributesPointersEncoding = input.CommonAttributesPointersEncoding
	commonattributesDB.CopyBasicFieldsFromCommonAttributes_WOP(&input.CommonAttributes_WOP)

	query := db.Create(&commonattributesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCommonAttributes.CheckoutPhaseOneInstance(&commonattributesDB)
	commonattributes := backRepo.BackRepoCommonAttributes.Map_CommonAttributesDBID_CommonAttributesPtr[commonattributesDB.ID]

	if commonattributes != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), commonattributes)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, commonattributesDB)
}

// GetCommonAttributes
//
// swagger:route GET /commonattributess/{ID} commonattributess getCommonAttributes
//
// Gets the details for a commonattributes.
//
// Responses:
// default: genericError
//
//	200: commonattributesDBResponse
func (controller *Controller) GetCommonAttributes(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCommonAttributes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCommonAttributes.GetDB()

	// Get commonattributesDB in DB
	var commonattributesDB orm.CommonAttributesDB
	if err := db.First(&commonattributesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var commonattributesAPI orm.CommonAttributesAPI
	commonattributesAPI.ID = commonattributesDB.ID
	commonattributesAPI.CommonAttributesPointersEncoding = commonattributesDB.CommonAttributesPointersEncoding
	commonattributesDB.CopyBasicFieldsToCommonAttributes_WOP(&commonattributesAPI.CommonAttributes_WOP)

	c.JSON(http.StatusOK, commonattributesAPI)
}

// UpdateCommonAttributes
//
// swagger:route PATCH /commonattributess/{ID} commonattributess updateCommonAttributes
//
// # Update a commonattributes
//
// Responses:
// default: genericError
//
//	200: commonattributesDBResponse
func (controller *Controller) UpdateCommonAttributes(c *gin.Context) {

	mutexCommonAttributes.Lock()
	defer mutexCommonAttributes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCommonAttributes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCommonAttributes.GetDB()

	// Validate input
	var input orm.CommonAttributesAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var commonattributesDB orm.CommonAttributesDB

	// fetch the commonattributes
	query := db.First(&commonattributesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	commonattributesDB.CopyBasicFieldsFromCommonAttributes_WOP(&input.CommonAttributes_WOP)
	commonattributesDB.CommonAttributesPointersEncoding = input.CommonAttributesPointersEncoding

	query = db.Model(&commonattributesDB).Updates(commonattributesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	commonattributesNew := new(models.CommonAttributes)
	commonattributesDB.CopyBasicFieldsToCommonAttributes(commonattributesNew)

	// redeem pointers
	commonattributesDB.DecodePointers(backRepo, commonattributesNew)

	// get stage instance from DB instance, and call callback function
	commonattributesOld := backRepo.BackRepoCommonAttributes.Map_CommonAttributesDBID_CommonAttributesPtr[commonattributesDB.ID]
	if commonattributesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), commonattributesOld, commonattributesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the commonattributesDB
	c.JSON(http.StatusOK, commonattributesDB)
}

// DeleteCommonAttributes
//
// swagger:route DELETE /commonattributess/{ID} commonattributess deleteCommonAttributes
//
// # Delete a commonattributes
//
// default: genericError
//
//	200: commonattributesDBResponse
func (controller *Controller) DeleteCommonAttributes(c *gin.Context) {

	mutexCommonAttributes.Lock()
	defer mutexCommonAttributes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCommonAttributes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/books/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCommonAttributes.GetDB()

	// Get model if exist
	var commonattributesDB orm.CommonAttributesDB
	if err := db.First(&commonattributesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&commonattributesDB)

	// get an instance (not staged) from DB instance, and call callback function
	commonattributesDeleted := new(models.CommonAttributes)
	commonattributesDB.CopyBasicFieldsToCommonAttributes(commonattributesDeleted)

	// get stage instance from DB instance, and call callback function
	commonattributesStaged := backRepo.BackRepoCommonAttributes.Map_CommonAttributesDBID_CommonAttributesPtr[commonattributesDB.ID]
	if commonattributesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), commonattributesStaged, commonattributesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
