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
var __Attribute__dummysDeclaration__ models.Attribute
var __Attribute_time__dummyDeclaration time.Duration

var mutexAttribute sync.Mutex

// An AttributeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAttribute updateAttribute deleteAttribute
type AttributeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AttributeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAttribute updateAttribute
type AttributeInput struct {
	// The Attribute to submit or modify
	// in: body
	Attribute *orm.AttributeAPI
}

// GetAttributes
//
// swagger:route GET /attributes attributes getAttributes
//
// # Get all attributes
//
// Responses:
// default: genericError
//
//	200: attributeDBResponse
func (controller *Controller) GetAttributes(c *gin.Context) {

	// source slice
	var attributeDBs []orm.AttributeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAttributes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttribute.GetDB()

	_, err := db.Find(&attributeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributeAPIs := make([]orm.AttributeAPI, 0)

	// for each attribute, update fields from the database nullable fields
	for idx := range attributeDBs {
		attributeDB := &attributeDBs[idx]
		_ = attributeDB
		var attributeAPI orm.AttributeAPI

		// insertion point for updating fields
		attributeAPI.ID = attributeDB.ID
		attributeDB.CopyBasicFieldsToAttribute_WOP(&attributeAPI.Attribute_WOP)
		attributeAPI.AttributePointersEncoding = attributeDB.AttributePointersEncoding
		attributeAPIs = append(attributeAPIs, attributeAPI)
	}

	c.JSON(http.StatusOK, attributeAPIs)
}

// PostAttribute
//
// swagger:route POST /attributes attributes postAttribute
//
// Creates a attribute
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAttribute(c *gin.Context) {

	mutexAttribute.Lock()
	defer mutexAttribute.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAttributes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttribute.GetDB()

	// Validate input
	var input orm.AttributeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute
	attributeDB := orm.AttributeDB{}
	attributeDB.AttributePointersEncoding = input.AttributePointersEncoding
	attributeDB.CopyBasicFieldsFromAttribute_WOP(&input.Attribute_WOP)

	_, err = db.Create(&attributeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAttribute.CheckoutPhaseOneInstance(&attributeDB)
	attribute := backRepo.BackRepoAttribute.Map_AttributeDBID_AttributePtr[attributeDB.ID]

	if attribute != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributeDB)
}

// GetAttribute
//
// swagger:route GET /attributes/{ID} attributes getAttribute
//
// Gets the details for a attribute.
//
// Responses:
// default: genericError
//
//	200: attributeDBResponse
func (controller *Controller) GetAttribute(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAttribute", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttribute.GetDB()

	// Get attributeDB in DB
	var attributeDB orm.AttributeDB
	if _, err := db.First(&attributeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributeAPI orm.AttributeAPI
	attributeAPI.ID = attributeDB.ID
	attributeAPI.AttributePointersEncoding = attributeDB.AttributePointersEncoding
	attributeDB.CopyBasicFieldsToAttribute_WOP(&attributeAPI.Attribute_WOP)

	c.JSON(http.StatusOK, attributeAPI)
}

// UpdateAttribute
//
// swagger:route PATCH /attributes/{ID} attributes updateAttribute
//
// # Update a attribute
//
// Responses:
// default: genericError
//
//	200: attributeDBResponse
func (controller *Controller) UpdateAttribute(c *gin.Context) {

	mutexAttribute.Lock()
	defer mutexAttribute.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAttribute", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttribute.GetDB()

	// Validate input
	var input orm.AttributeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributeDB orm.AttributeDB

	// fetch the attribute
	_, err := db.First(&attributeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributeDB.CopyBasicFieldsFromAttribute_WOP(&input.Attribute_WOP)
	attributeDB.AttributePointersEncoding = input.AttributePointersEncoding

	db, _ = db.Model(&attributeDB)
	_, err = db.Updates(attributeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributeNew := new(models.Attribute)
	attributeDB.CopyBasicFieldsToAttribute(attributeNew)

	// redeem pointers
	attributeDB.DecodePointers(backRepo, attributeNew)

	// get stage instance from DB instance, and call callback function
	attributeOld := backRepo.BackRepoAttribute.Map_AttributeDBID_AttributePtr[attributeDB.ID]
	if attributeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributeOld, attributeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributeDB
	c.JSON(http.StatusOK, attributeDB)
}

// DeleteAttribute
//
// swagger:route DELETE /attributes/{ID} attributes deleteAttribute
//
// # Delete a attribute
//
// default: genericError
//
//	200: attributeDBResponse
func (controller *Controller) DeleteAttribute(c *gin.Context) {

	mutexAttribute.Lock()
	defer mutexAttribute.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAttribute", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttribute.GetDB()

	// Get model if exist
	var attributeDB orm.AttributeDB
	if _, err := db.First(&attributeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attributeDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributeDeleted := new(models.Attribute)
	attributeDB.CopyBasicFieldsToAttribute(attributeDeleted)

	// get stage instance from DB instance, and call callback function
	attributeStaged := backRepo.BackRepoAttribute.Map_AttributeDBID_AttributePtr[attributeDB.ID]
	if attributeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributeStaged, attributeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
