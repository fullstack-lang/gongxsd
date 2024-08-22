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
var __A_VALUES__dummysDeclaration__ models.A_VALUES
var __A_VALUES_time__dummyDeclaration time.Duration

var mutexA_VALUES sync.Mutex

// An A_VALUESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_VALUES updateA_VALUES deleteA_VALUES
type A_VALUESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_VALUESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_VALUES updateA_VALUES
type A_VALUESInput struct {
	// The A_VALUES to submit or modify
	// in: body
	A_VALUES *orm.A_VALUESAPI
}

// GetA_VALUESs
//
// swagger:route GET /a_valuess a_valuess getA_VALUESs
//
// # Get all a_valuess
//
// Responses:
// default: genericError
//
//	200: a_valuesDBResponse
func (controller *Controller) GetA_VALUESs(c *gin.Context) {

	// source slice
	var a_valuesDBs []orm.A_VALUESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_VALUESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_VALUES.GetDB()

	query := db.Find(&a_valuesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_valuesAPIs := make([]orm.A_VALUESAPI, 0)

	// for each a_values, update fields from the database nullable fields
	for idx := range a_valuesDBs {
		a_valuesDB := &a_valuesDBs[idx]
		_ = a_valuesDB
		var a_valuesAPI orm.A_VALUESAPI

		// insertion point for updating fields
		a_valuesAPI.ID = a_valuesDB.ID
		a_valuesDB.CopyBasicFieldsToA_VALUES_WOP(&a_valuesAPI.A_VALUES_WOP)
		a_valuesAPI.A_VALUESPointersEncoding = a_valuesDB.A_VALUESPointersEncoding
		a_valuesAPIs = append(a_valuesAPIs, a_valuesAPI)
	}

	c.JSON(http.StatusOK, a_valuesAPIs)
}

// PostA_VALUES
//
// swagger:route POST /a_valuess a_valuess postA_VALUES
//
// Creates a a_values
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_VALUES(c *gin.Context) {

	mutexA_VALUES.Lock()
	defer mutexA_VALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_VALUESs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_VALUES.GetDB()

	// Validate input
	var input orm.A_VALUESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_values
	a_valuesDB := orm.A_VALUESDB{}
	a_valuesDB.A_VALUESPointersEncoding = input.A_VALUESPointersEncoding
	a_valuesDB.CopyBasicFieldsFromA_VALUES_WOP(&input.A_VALUES_WOP)

	query := db.Create(&a_valuesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_VALUES.CheckoutPhaseOneInstance(&a_valuesDB)
	a_values := backRepo.BackRepoA_VALUES.Map_A_VALUESDBID_A_VALUESPtr[a_valuesDB.ID]

	if a_values != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_values)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_valuesDB)
}

// GetA_VALUES
//
// swagger:route GET /a_valuess/{ID} a_valuess getA_VALUES
//
// Gets the details for a a_values.
//
// Responses:
// default: genericError
//
//	200: a_valuesDBResponse
func (controller *Controller) GetA_VALUES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_VALUES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_VALUES.GetDB()

	// Get a_valuesDB in DB
	var a_valuesDB orm.A_VALUESDB
	if err := db.First(&a_valuesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_valuesAPI orm.A_VALUESAPI
	a_valuesAPI.ID = a_valuesDB.ID
	a_valuesAPI.A_VALUESPointersEncoding = a_valuesDB.A_VALUESPointersEncoding
	a_valuesDB.CopyBasicFieldsToA_VALUES_WOP(&a_valuesAPI.A_VALUES_WOP)

	c.JSON(http.StatusOK, a_valuesAPI)
}

// UpdateA_VALUES
//
// swagger:route PATCH /a_valuess/{ID} a_valuess updateA_VALUES
//
// # Update a a_values
//
// Responses:
// default: genericError
//
//	200: a_valuesDBResponse
func (controller *Controller) UpdateA_VALUES(c *gin.Context) {

	mutexA_VALUES.Lock()
	defer mutexA_VALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_VALUES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_VALUES.GetDB()

	// Validate input
	var input orm.A_VALUESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_valuesDB orm.A_VALUESDB

	// fetch the a_values
	query := db.First(&a_valuesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_valuesDB.CopyBasicFieldsFromA_VALUES_WOP(&input.A_VALUES_WOP)
	a_valuesDB.A_VALUESPointersEncoding = input.A_VALUESPointersEncoding

	query = db.Model(&a_valuesDB).Updates(a_valuesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_valuesNew := new(models.A_VALUES)
	a_valuesDB.CopyBasicFieldsToA_VALUES(a_valuesNew)

	// redeem pointers
	a_valuesDB.DecodePointers(backRepo, a_valuesNew)

	// get stage instance from DB instance, and call callback function
	a_valuesOld := backRepo.BackRepoA_VALUES.Map_A_VALUESDBID_A_VALUESPtr[a_valuesDB.ID]
	if a_valuesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_valuesOld, a_valuesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_valuesDB
	c.JSON(http.StatusOK, a_valuesDB)
}

// DeleteA_VALUES
//
// swagger:route DELETE /a_valuess/{ID} a_valuess deleteA_VALUES
//
// # Delete a a_values
//
// default: genericError
//
//	200: a_valuesDBResponse
func (controller *Controller) DeleteA_VALUES(c *gin.Context) {

	mutexA_VALUES.Lock()
	defer mutexA_VALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_VALUES", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_VALUES.GetDB()

	// Get model if exist
	var a_valuesDB orm.A_VALUESDB
	if err := db.First(&a_valuesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_valuesDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_valuesDeleted := new(models.A_VALUES)
	a_valuesDB.CopyBasicFieldsToA_VALUES(a_valuesDeleted)

	// get stage instance from DB instance, and call callback function
	a_valuesStaged := backRepo.BackRepoA_VALUES.Map_A_VALUESDBID_A_VALUESPtr[a_valuesDB.ID]
	if a_valuesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_valuesStaged, a_valuesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
