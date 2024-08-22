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
var __A_DEFAULT_VALUE_1__dummysDeclaration__ models.A_DEFAULT_VALUE_1
var __A_DEFAULT_VALUE_1_time__dummyDeclaration time.Duration

var mutexA_DEFAULT_VALUE_1 sync.Mutex

// An A_DEFAULT_VALUE_1ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DEFAULT_VALUE_1 updateA_DEFAULT_VALUE_1 deleteA_DEFAULT_VALUE_1
type A_DEFAULT_VALUE_1ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DEFAULT_VALUE_1Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DEFAULT_VALUE_1 updateA_DEFAULT_VALUE_1
type A_DEFAULT_VALUE_1Input struct {
	// The A_DEFAULT_VALUE_1 to submit or modify
	// in: body
	A_DEFAULT_VALUE_1 *orm.A_DEFAULT_VALUE_1API
}

// GetA_DEFAULT_VALUE_1s
//
// swagger:route GET /a_default_value_1s a_default_value_1s getA_DEFAULT_VALUE_1s
//
// # Get all a_default_value_1s
//
// Responses:
// default: genericError
//
//	200: a_default_value_1DBResponse
func (controller *Controller) GetA_DEFAULT_VALUE_1s(c *gin.Context) {

	// source slice
	var a_default_value_1DBs []orm.A_DEFAULT_VALUE_1DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUE_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_1.GetDB()

	query := db.Find(&a_default_value_1DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_default_value_1APIs := make([]orm.A_DEFAULT_VALUE_1API, 0)

	// for each a_default_value_1, update fields from the database nullable fields
	for idx := range a_default_value_1DBs {
		a_default_value_1DB := &a_default_value_1DBs[idx]
		_ = a_default_value_1DB
		var a_default_value_1API orm.A_DEFAULT_VALUE_1API

		// insertion point for updating fields
		a_default_value_1API.ID = a_default_value_1DB.ID
		a_default_value_1DB.CopyBasicFieldsToA_DEFAULT_VALUE_1_WOP(&a_default_value_1API.A_DEFAULT_VALUE_1_WOP)
		a_default_value_1API.A_DEFAULT_VALUE_1PointersEncoding = a_default_value_1DB.A_DEFAULT_VALUE_1PointersEncoding
		a_default_value_1APIs = append(a_default_value_1APIs, a_default_value_1API)
	}

	c.JSON(http.StatusOK, a_default_value_1APIs)
}

// PostA_DEFAULT_VALUE_1
//
// swagger:route POST /a_default_value_1s a_default_value_1s postA_DEFAULT_VALUE_1
//
// Creates a a_default_value_1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DEFAULT_VALUE_1(c *gin.Context) {

	mutexA_DEFAULT_VALUE_1.Lock()
	defer mutexA_DEFAULT_VALUE_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DEFAULT_VALUE_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_1.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUE_1API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_default_value_1
	a_default_value_1DB := orm.A_DEFAULT_VALUE_1DB{}
	a_default_value_1DB.A_DEFAULT_VALUE_1PointersEncoding = input.A_DEFAULT_VALUE_1PointersEncoding
	a_default_value_1DB.CopyBasicFieldsFromA_DEFAULT_VALUE_1_WOP(&input.A_DEFAULT_VALUE_1_WOP)

	query := db.Create(&a_default_value_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DEFAULT_VALUE_1.CheckoutPhaseOneInstance(&a_default_value_1DB)
	a_default_value_1 := backRepo.BackRepoA_DEFAULT_VALUE_1.Map_A_DEFAULT_VALUE_1DBID_A_DEFAULT_VALUE_1Ptr[a_default_value_1DB.ID]

	if a_default_value_1 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_default_value_1)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_default_value_1DB)
}

// GetA_DEFAULT_VALUE_1
//
// swagger:route GET /a_default_value_1s/{ID} a_default_value_1s getA_DEFAULT_VALUE_1
//
// Gets the details for a a_default_value_1.
//
// Responses:
// default: genericError
//
//	200: a_default_value_1DBResponse
func (controller *Controller) GetA_DEFAULT_VALUE_1(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUE_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_1.GetDB()

	// Get a_default_value_1DB in DB
	var a_default_value_1DB orm.A_DEFAULT_VALUE_1DB
	if err := db.First(&a_default_value_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_default_value_1API orm.A_DEFAULT_VALUE_1API
	a_default_value_1API.ID = a_default_value_1DB.ID
	a_default_value_1API.A_DEFAULT_VALUE_1PointersEncoding = a_default_value_1DB.A_DEFAULT_VALUE_1PointersEncoding
	a_default_value_1DB.CopyBasicFieldsToA_DEFAULT_VALUE_1_WOP(&a_default_value_1API.A_DEFAULT_VALUE_1_WOP)

	c.JSON(http.StatusOK, a_default_value_1API)
}

// UpdateA_DEFAULT_VALUE_1
//
// swagger:route PATCH /a_default_value_1s/{ID} a_default_value_1s updateA_DEFAULT_VALUE_1
//
// # Update a a_default_value_1
//
// Responses:
// default: genericError
//
//	200: a_default_value_1DBResponse
func (controller *Controller) UpdateA_DEFAULT_VALUE_1(c *gin.Context) {

	mutexA_DEFAULT_VALUE_1.Lock()
	defer mutexA_DEFAULT_VALUE_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DEFAULT_VALUE_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_1.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUE_1API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_default_value_1DB orm.A_DEFAULT_VALUE_1DB

	// fetch the a_default_value_1
	query := db.First(&a_default_value_1DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_default_value_1DB.CopyBasicFieldsFromA_DEFAULT_VALUE_1_WOP(&input.A_DEFAULT_VALUE_1_WOP)
	a_default_value_1DB.A_DEFAULT_VALUE_1PointersEncoding = input.A_DEFAULT_VALUE_1PointersEncoding

	query = db.Model(&a_default_value_1DB).Updates(a_default_value_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_default_value_1New := new(models.A_DEFAULT_VALUE_1)
	a_default_value_1DB.CopyBasicFieldsToA_DEFAULT_VALUE_1(a_default_value_1New)

	// redeem pointers
	a_default_value_1DB.DecodePointers(backRepo, a_default_value_1New)

	// get stage instance from DB instance, and call callback function
	a_default_value_1Old := backRepo.BackRepoA_DEFAULT_VALUE_1.Map_A_DEFAULT_VALUE_1DBID_A_DEFAULT_VALUE_1Ptr[a_default_value_1DB.ID]
	if a_default_value_1Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_default_value_1Old, a_default_value_1New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_default_value_1DB
	c.JSON(http.StatusOK, a_default_value_1DB)
}

// DeleteA_DEFAULT_VALUE_1
//
// swagger:route DELETE /a_default_value_1s/{ID} a_default_value_1s deleteA_DEFAULT_VALUE_1
//
// # Delete a a_default_value_1
//
// default: genericError
//
//	200: a_default_value_1DBResponse
func (controller *Controller) DeleteA_DEFAULT_VALUE_1(c *gin.Context) {

	mutexA_DEFAULT_VALUE_1.Lock()
	defer mutexA_DEFAULT_VALUE_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DEFAULT_VALUE_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_1.GetDB()

	// Get model if exist
	var a_default_value_1DB orm.A_DEFAULT_VALUE_1DB
	if err := db.First(&a_default_value_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_default_value_1DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_default_value_1Deleted := new(models.A_DEFAULT_VALUE_1)
	a_default_value_1DB.CopyBasicFieldsToA_DEFAULT_VALUE_1(a_default_value_1Deleted)

	// get stage instance from DB instance, and call callback function
	a_default_value_1Staged := backRepo.BackRepoA_DEFAULT_VALUE_1.Map_A_DEFAULT_VALUE_1DBID_A_DEFAULT_VALUE_1Ptr[a_default_value_1DB.ID]
	if a_default_value_1Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_default_value_1Staged, a_default_value_1Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
