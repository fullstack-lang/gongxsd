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
var __A_DEFAULT_VALUE__dummysDeclaration__ models.A_DEFAULT_VALUE
var __A_DEFAULT_VALUE_time__dummyDeclaration time.Duration

var mutexA_DEFAULT_VALUE sync.Mutex

// An A_DEFAULT_VALUEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DEFAULT_VALUE updateA_DEFAULT_VALUE deleteA_DEFAULT_VALUE
type A_DEFAULT_VALUEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DEFAULT_VALUEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DEFAULT_VALUE updateA_DEFAULT_VALUE
type A_DEFAULT_VALUEInput struct {
	// The A_DEFAULT_VALUE to submit or modify
	// in: body
	A_DEFAULT_VALUE *orm.A_DEFAULT_VALUEAPI
}

// GetA_DEFAULT_VALUEs
//
// swagger:route GET /a_default_values a_default_values getA_DEFAULT_VALUEs
//
// # Get all a_default_values
//
// Responses:
// default: genericError
//
//	200: a_default_valueDBResponse
func (controller *Controller) GetA_DEFAULT_VALUEs(c *gin.Context) {

	// source slice
	var a_default_valueDBs []orm.A_DEFAULT_VALUEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE.GetDB()

	query := db.Find(&a_default_valueDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_default_valueAPIs := make([]orm.A_DEFAULT_VALUEAPI, 0)

	// for each a_default_value, update fields from the database nullable fields
	for idx := range a_default_valueDBs {
		a_default_valueDB := &a_default_valueDBs[idx]
		_ = a_default_valueDB
		var a_default_valueAPI orm.A_DEFAULT_VALUEAPI

		// insertion point for updating fields
		a_default_valueAPI.ID = a_default_valueDB.ID
		a_default_valueDB.CopyBasicFieldsToA_DEFAULT_VALUE_WOP(&a_default_valueAPI.A_DEFAULT_VALUE_WOP)
		a_default_valueAPI.A_DEFAULT_VALUEPointersEncoding = a_default_valueDB.A_DEFAULT_VALUEPointersEncoding
		a_default_valueAPIs = append(a_default_valueAPIs, a_default_valueAPI)
	}

	c.JSON(http.StatusOK, a_default_valueAPIs)
}

// PostA_DEFAULT_VALUE
//
// swagger:route POST /a_default_values a_default_values postA_DEFAULT_VALUE
//
// Creates a a_default_value
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DEFAULT_VALUE(c *gin.Context) {

	mutexA_DEFAULT_VALUE.Lock()
	defer mutexA_DEFAULT_VALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DEFAULT_VALUEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_default_value
	a_default_valueDB := orm.A_DEFAULT_VALUEDB{}
	a_default_valueDB.A_DEFAULT_VALUEPointersEncoding = input.A_DEFAULT_VALUEPointersEncoding
	a_default_valueDB.CopyBasicFieldsFromA_DEFAULT_VALUE_WOP(&input.A_DEFAULT_VALUE_WOP)

	query := db.Create(&a_default_valueDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DEFAULT_VALUE.CheckoutPhaseOneInstance(&a_default_valueDB)
	a_default_value := backRepo.BackRepoA_DEFAULT_VALUE.Map_A_DEFAULT_VALUEDBID_A_DEFAULT_VALUEPtr[a_default_valueDB.ID]

	if a_default_value != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_default_value)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_default_valueDB)
}

// GetA_DEFAULT_VALUE
//
// swagger:route GET /a_default_values/{ID} a_default_values getA_DEFAULT_VALUE
//
// Gets the details for a a_default_value.
//
// Responses:
// default: genericError
//
//	200: a_default_valueDBResponse
func (controller *Controller) GetA_DEFAULT_VALUE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE.GetDB()

	// Get a_default_valueDB in DB
	var a_default_valueDB orm.A_DEFAULT_VALUEDB
	if err := db.First(&a_default_valueDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_default_valueAPI orm.A_DEFAULT_VALUEAPI
	a_default_valueAPI.ID = a_default_valueDB.ID
	a_default_valueAPI.A_DEFAULT_VALUEPointersEncoding = a_default_valueDB.A_DEFAULT_VALUEPointersEncoding
	a_default_valueDB.CopyBasicFieldsToA_DEFAULT_VALUE_WOP(&a_default_valueAPI.A_DEFAULT_VALUE_WOP)

	c.JSON(http.StatusOK, a_default_valueAPI)
}

// UpdateA_DEFAULT_VALUE
//
// swagger:route PATCH /a_default_values/{ID} a_default_values updateA_DEFAULT_VALUE
//
// # Update a a_default_value
//
// Responses:
// default: genericError
//
//	200: a_default_valueDBResponse
func (controller *Controller) UpdateA_DEFAULT_VALUE(c *gin.Context) {

	mutexA_DEFAULT_VALUE.Lock()
	defer mutexA_DEFAULT_VALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DEFAULT_VALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_default_valueDB orm.A_DEFAULT_VALUEDB

	// fetch the a_default_value
	query := db.First(&a_default_valueDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_default_valueDB.CopyBasicFieldsFromA_DEFAULT_VALUE_WOP(&input.A_DEFAULT_VALUE_WOP)
	a_default_valueDB.A_DEFAULT_VALUEPointersEncoding = input.A_DEFAULT_VALUEPointersEncoding

	query = db.Model(&a_default_valueDB).Updates(a_default_valueDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_default_valueNew := new(models.A_DEFAULT_VALUE)
	a_default_valueDB.CopyBasicFieldsToA_DEFAULT_VALUE(a_default_valueNew)

	// redeem pointers
	a_default_valueDB.DecodePointers(backRepo, a_default_valueNew)

	// get stage instance from DB instance, and call callback function
	a_default_valueOld := backRepo.BackRepoA_DEFAULT_VALUE.Map_A_DEFAULT_VALUEDBID_A_DEFAULT_VALUEPtr[a_default_valueDB.ID]
	if a_default_valueOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_default_valueOld, a_default_valueNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_default_valueDB
	c.JSON(http.StatusOK, a_default_valueDB)
}

// DeleteA_DEFAULT_VALUE
//
// swagger:route DELETE /a_default_values/{ID} a_default_values deleteA_DEFAULT_VALUE
//
// # Delete a a_default_value
//
// default: genericError
//
//	200: a_default_valueDBResponse
func (controller *Controller) DeleteA_DEFAULT_VALUE(c *gin.Context) {

	mutexA_DEFAULT_VALUE.Lock()
	defer mutexA_DEFAULT_VALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DEFAULT_VALUE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE.GetDB()

	// Get model if exist
	var a_default_valueDB orm.A_DEFAULT_VALUEDB
	if err := db.First(&a_default_valueDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_default_valueDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_default_valueDeleted := new(models.A_DEFAULT_VALUE)
	a_default_valueDB.CopyBasicFieldsToA_DEFAULT_VALUE(a_default_valueDeleted)

	// get stage instance from DB instance, and call callback function
	a_default_valueStaged := backRepo.BackRepoA_DEFAULT_VALUE.Map_A_DEFAULT_VALUEDBID_A_DEFAULT_VALUEPtr[a_default_valueDB.ID]
	if a_default_valueStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_default_valueStaged, a_default_valueDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
