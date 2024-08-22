// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
	"github.com/fullstack-lang/gongxsd/test/reqif/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __A_PROPERTIES__dummysDeclaration__ models.A_PROPERTIES
var __A_PROPERTIES_time__dummyDeclaration time.Duration

var mutexA_PROPERTIES sync.Mutex

// An A_PROPERTIESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_PROPERTIES updateA_PROPERTIES deleteA_PROPERTIES
type A_PROPERTIESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_PROPERTIESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_PROPERTIES updateA_PROPERTIES
type A_PROPERTIESInput struct {
	// The A_PROPERTIES to submit or modify
	// in: body
	A_PROPERTIES *orm.A_PROPERTIESAPI
}

// GetA_PROPERTIESs
//
// swagger:route GET /a_propertiess a_propertiess getA_PROPERTIESs
//
// # Get all a_propertiess
//
// Responses:
// default: genericError
//
//	200: a_propertiesDBResponse
func (controller *Controller) GetA_PROPERTIESs(c *gin.Context) {

	// source slice
	var a_propertiesDBs []orm.A_PROPERTIESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_PROPERTIESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_PROPERTIES.GetDB()

	query := db.Find(&a_propertiesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_propertiesAPIs := make([]orm.A_PROPERTIESAPI, 0)

	// for each a_properties, update fields from the database nullable fields
	for idx := range a_propertiesDBs {
		a_propertiesDB := &a_propertiesDBs[idx]
		_ = a_propertiesDB
		var a_propertiesAPI orm.A_PROPERTIESAPI

		// insertion point for updating fields
		a_propertiesAPI.ID = a_propertiesDB.ID
		a_propertiesDB.CopyBasicFieldsToA_PROPERTIES_WOP(&a_propertiesAPI.A_PROPERTIES_WOP)
		a_propertiesAPI.A_PROPERTIESPointersEncoding = a_propertiesDB.A_PROPERTIESPointersEncoding
		a_propertiesAPIs = append(a_propertiesAPIs, a_propertiesAPI)
	}

	c.JSON(http.StatusOK, a_propertiesAPIs)
}

// PostA_PROPERTIES
//
// swagger:route POST /a_propertiess a_propertiess postA_PROPERTIES
//
// Creates a a_properties
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_PROPERTIES(c *gin.Context) {

	mutexA_PROPERTIES.Lock()
	defer mutexA_PROPERTIES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_PROPERTIESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_PROPERTIES.GetDB()

	// Validate input
	var input orm.A_PROPERTIESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_properties
	a_propertiesDB := orm.A_PROPERTIESDB{}
	a_propertiesDB.A_PROPERTIESPointersEncoding = input.A_PROPERTIESPointersEncoding
	a_propertiesDB.CopyBasicFieldsFromA_PROPERTIES_WOP(&input.A_PROPERTIES_WOP)

	query := db.Create(&a_propertiesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_PROPERTIES.CheckoutPhaseOneInstance(&a_propertiesDB)
	a_properties := backRepo.BackRepoA_PROPERTIES.Map_A_PROPERTIESDBID_A_PROPERTIESPtr[a_propertiesDB.ID]

	if a_properties != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_properties)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_propertiesDB)
}

// GetA_PROPERTIES
//
// swagger:route GET /a_propertiess/{ID} a_propertiess getA_PROPERTIES
//
// Gets the details for a a_properties.
//
// Responses:
// default: genericError
//
//	200: a_propertiesDBResponse
func (controller *Controller) GetA_PROPERTIES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_PROPERTIES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_PROPERTIES.GetDB()

	// Get a_propertiesDB in DB
	var a_propertiesDB orm.A_PROPERTIESDB
	if err := db.First(&a_propertiesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_propertiesAPI orm.A_PROPERTIESAPI
	a_propertiesAPI.ID = a_propertiesDB.ID
	a_propertiesAPI.A_PROPERTIESPointersEncoding = a_propertiesDB.A_PROPERTIESPointersEncoding
	a_propertiesDB.CopyBasicFieldsToA_PROPERTIES_WOP(&a_propertiesAPI.A_PROPERTIES_WOP)

	c.JSON(http.StatusOK, a_propertiesAPI)
}

// UpdateA_PROPERTIES
//
// swagger:route PATCH /a_propertiess/{ID} a_propertiess updateA_PROPERTIES
//
// # Update a a_properties
//
// Responses:
// default: genericError
//
//	200: a_propertiesDBResponse
func (controller *Controller) UpdateA_PROPERTIES(c *gin.Context) {

	mutexA_PROPERTIES.Lock()
	defer mutexA_PROPERTIES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_PROPERTIES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_PROPERTIES.GetDB()

	// Validate input
	var input orm.A_PROPERTIESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_propertiesDB orm.A_PROPERTIESDB

	// fetch the a_properties
	query := db.First(&a_propertiesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_propertiesDB.CopyBasicFieldsFromA_PROPERTIES_WOP(&input.A_PROPERTIES_WOP)
	a_propertiesDB.A_PROPERTIESPointersEncoding = input.A_PROPERTIESPointersEncoding

	query = db.Model(&a_propertiesDB).Updates(a_propertiesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_propertiesNew := new(models.A_PROPERTIES)
	a_propertiesDB.CopyBasicFieldsToA_PROPERTIES(a_propertiesNew)

	// redeem pointers
	a_propertiesDB.DecodePointers(backRepo, a_propertiesNew)

	// get stage instance from DB instance, and call callback function
	a_propertiesOld := backRepo.BackRepoA_PROPERTIES.Map_A_PROPERTIESDBID_A_PROPERTIESPtr[a_propertiesDB.ID]
	if a_propertiesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_propertiesOld, a_propertiesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_propertiesDB
	c.JSON(http.StatusOK, a_propertiesDB)
}

// DeleteA_PROPERTIES
//
// swagger:route DELETE /a_propertiess/{ID} a_propertiess deleteA_PROPERTIES
//
// # Delete a a_properties
//
// default: genericError
//
//	200: a_propertiesDBResponse
func (controller *Controller) DeleteA_PROPERTIES(c *gin.Context) {

	mutexA_PROPERTIES.Lock()
	defer mutexA_PROPERTIES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_PROPERTIES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_PROPERTIES.GetDB()

	// Get model if exist
	var a_propertiesDB orm.A_PROPERTIESDB
	if err := db.First(&a_propertiesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_propertiesDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_propertiesDeleted := new(models.A_PROPERTIES)
	a_propertiesDB.CopyBasicFieldsToA_PROPERTIES(a_propertiesDeleted)

	// get stage instance from DB instance, and call callback function
	a_propertiesStaged := backRepo.BackRepoA_PROPERTIES.Map_A_PROPERTIESDBID_A_PROPERTIESPtr[a_propertiesDB.ID]
	if a_propertiesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_propertiesStaged, a_propertiesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
