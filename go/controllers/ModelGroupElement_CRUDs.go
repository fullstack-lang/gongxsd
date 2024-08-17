// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxsd/go/models"
	"github.com/fullstack-lang/gongxsd/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __ModelGroupElement__dummysDeclaration__ models.ModelGroupElement
var __ModelGroupElement_time__dummyDeclaration time.Duration

var mutexModelGroupElement sync.Mutex

// An ModelGroupElementID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getModelGroupElement updateModelGroupElement deleteModelGroupElement
type ModelGroupElementID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ModelGroupElementInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postModelGroupElement updateModelGroupElement
type ModelGroupElementInput struct {
	// The ModelGroupElement to submit or modify
	// in: body
	ModelGroupElement *orm.ModelGroupElementAPI
}

// GetModelGroupElements
//
// swagger:route GET /modelgroupelements modelgroupelements getModelGroupElements
//
// # Get all modelgroupelements
//
// Responses:
// default: genericError
//
//	200: modelgroupelementDBResponse
func (controller *Controller) GetModelGroupElements(c *gin.Context) {

	// source slice
	var modelgroupelementDBs []orm.ModelGroupElementDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetModelGroupElements", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoModelGroupElement.GetDB()

	query := db.Find(&modelgroupelementDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	modelgroupelementAPIs := make([]orm.ModelGroupElementAPI, 0)

	// for each modelgroupelement, update fields from the database nullable fields
	for idx := range modelgroupelementDBs {
		modelgroupelementDB := &modelgroupelementDBs[idx]
		_ = modelgroupelementDB
		var modelgroupelementAPI orm.ModelGroupElementAPI

		// insertion point for updating fields
		modelgroupelementAPI.ID = modelgroupelementDB.ID
		modelgroupelementDB.CopyBasicFieldsToModelGroupElement_WOP(&modelgroupelementAPI.ModelGroupElement_WOP)
		modelgroupelementAPI.ModelGroupElementPointersEncoding = modelgroupelementDB.ModelGroupElementPointersEncoding
		modelgroupelementAPIs = append(modelgroupelementAPIs, modelgroupelementAPI)
	}

	c.JSON(http.StatusOK, modelgroupelementAPIs)
}

// PostModelGroupElement
//
// swagger:route POST /modelgroupelements modelgroupelements postModelGroupElement
//
// Creates a modelgroupelement
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostModelGroupElement(c *gin.Context) {

	mutexModelGroupElement.Lock()
	defer mutexModelGroupElement.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostModelGroupElements", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoModelGroupElement.GetDB()

	// Validate input
	var input orm.ModelGroupElementAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create modelgroupelement
	modelgroupelementDB := orm.ModelGroupElementDB{}
	modelgroupelementDB.ModelGroupElementPointersEncoding = input.ModelGroupElementPointersEncoding
	modelgroupelementDB.CopyBasicFieldsFromModelGroupElement_WOP(&input.ModelGroupElement_WOP)

	query := db.Create(&modelgroupelementDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoModelGroupElement.CheckoutPhaseOneInstance(&modelgroupelementDB)
	modelgroupelement := backRepo.BackRepoModelGroupElement.Map_ModelGroupElementDBID_ModelGroupElementPtr[modelgroupelementDB.ID]

	if modelgroupelement != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), modelgroupelement)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, modelgroupelementDB)
}

// GetModelGroupElement
//
// swagger:route GET /modelgroupelements/{ID} modelgroupelements getModelGroupElement
//
// Gets the details for a modelgroupelement.
//
// Responses:
// default: genericError
//
//	200: modelgroupelementDBResponse
func (controller *Controller) GetModelGroupElement(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetModelGroupElement", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoModelGroupElement.GetDB()

	// Get modelgroupelementDB in DB
	var modelgroupelementDB orm.ModelGroupElementDB
	if err := db.First(&modelgroupelementDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var modelgroupelementAPI orm.ModelGroupElementAPI
	modelgroupelementAPI.ID = modelgroupelementDB.ID
	modelgroupelementAPI.ModelGroupElementPointersEncoding = modelgroupelementDB.ModelGroupElementPointersEncoding
	modelgroupelementDB.CopyBasicFieldsToModelGroupElement_WOP(&modelgroupelementAPI.ModelGroupElement_WOP)

	c.JSON(http.StatusOK, modelgroupelementAPI)
}

// UpdateModelGroupElement
//
// swagger:route PATCH /modelgroupelements/{ID} modelgroupelements updateModelGroupElement
//
// # Update a modelgroupelement
//
// Responses:
// default: genericError
//
//	200: modelgroupelementDBResponse
func (controller *Controller) UpdateModelGroupElement(c *gin.Context) {

	mutexModelGroupElement.Lock()
	defer mutexModelGroupElement.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateModelGroupElement", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoModelGroupElement.GetDB()

	// Validate input
	var input orm.ModelGroupElementAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var modelgroupelementDB orm.ModelGroupElementDB

	// fetch the modelgroupelement
	query := db.First(&modelgroupelementDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	modelgroupelementDB.CopyBasicFieldsFromModelGroupElement_WOP(&input.ModelGroupElement_WOP)
	modelgroupelementDB.ModelGroupElementPointersEncoding = input.ModelGroupElementPointersEncoding

	query = db.Model(&modelgroupelementDB).Updates(modelgroupelementDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	modelgroupelementNew := new(models.ModelGroupElement)
	modelgroupelementDB.CopyBasicFieldsToModelGroupElement(modelgroupelementNew)

	// redeem pointers
	modelgroupelementDB.DecodePointers(backRepo, modelgroupelementNew)

	// get stage instance from DB instance, and call callback function
	modelgroupelementOld := backRepo.BackRepoModelGroupElement.Map_ModelGroupElementDBID_ModelGroupElementPtr[modelgroupelementDB.ID]
	if modelgroupelementOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), modelgroupelementOld, modelgroupelementNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the modelgroupelementDB
	c.JSON(http.StatusOK, modelgroupelementDB)
}

// DeleteModelGroupElement
//
// swagger:route DELETE /modelgroupelements/{ID} modelgroupelements deleteModelGroupElement
//
// # Delete a modelgroupelement
//
// default: genericError
//
//	200: modelgroupelementDBResponse
func (controller *Controller) DeleteModelGroupElement(c *gin.Context) {

	mutexModelGroupElement.Lock()
	defer mutexModelGroupElement.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteModelGroupElement", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoModelGroupElement.GetDB()

	// Get model if exist
	var modelgroupelementDB orm.ModelGroupElementDB
	if err := db.First(&modelgroupelementDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&modelgroupelementDB)

	// get an instance (not staged) from DB instance, and call callback function
	modelgroupelementDeleted := new(models.ModelGroupElement)
	modelgroupelementDB.CopyBasicFieldsToModelGroupElement(modelgroupelementDeleted)

	// get stage instance from DB instance, and call callback function
	modelgroupelementStaged := backRepo.BackRepoModelGroupElement.Map_ModelGroupElementDBID_ModelGroupElementPtr[modelgroupelementDB.ID]
	if modelgroupelementStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), modelgroupelementStaged, modelgroupelementDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
