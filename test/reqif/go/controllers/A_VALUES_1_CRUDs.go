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
var __A_VALUES_1__dummysDeclaration__ models.A_VALUES_1
var __A_VALUES_1_time__dummyDeclaration time.Duration

var mutexA_VALUES_1 sync.Mutex

// An A_VALUES_1ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_VALUES_1 updateA_VALUES_1 deleteA_VALUES_1
type A_VALUES_1ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_VALUES_1Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_VALUES_1 updateA_VALUES_1
type A_VALUES_1Input struct {
	// The A_VALUES_1 to submit or modify
	// in: body
	A_VALUES_1 *orm.A_VALUES_1API
}

// GetA_VALUES_1s
//
// swagger:route GET /a_values_1s a_values_1s getA_VALUES_1s
//
// # Get all a_values_1s
//
// Responses:
// default: genericError
//
//	200: a_values_1DBResponse
func (controller *Controller) GetA_VALUES_1s(c *gin.Context) {

	// source slice
	var a_values_1DBs []orm.A_VALUES_1DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_VALUES_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_VALUES_1.GetDB()

	query := db.Find(&a_values_1DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_values_1APIs := make([]orm.A_VALUES_1API, 0)

	// for each a_values_1, update fields from the database nullable fields
	for idx := range a_values_1DBs {
		a_values_1DB := &a_values_1DBs[idx]
		_ = a_values_1DB
		var a_values_1API orm.A_VALUES_1API

		// insertion point for updating fields
		a_values_1API.ID = a_values_1DB.ID
		a_values_1DB.CopyBasicFieldsToA_VALUES_1_WOP(&a_values_1API.A_VALUES_1_WOP)
		a_values_1API.A_VALUES_1PointersEncoding = a_values_1DB.A_VALUES_1PointersEncoding
		a_values_1APIs = append(a_values_1APIs, a_values_1API)
	}

	c.JSON(http.StatusOK, a_values_1APIs)
}

// PostA_VALUES_1
//
// swagger:route POST /a_values_1s a_values_1s postA_VALUES_1
//
// Creates a a_values_1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_VALUES_1(c *gin.Context) {

	mutexA_VALUES_1.Lock()
	defer mutexA_VALUES_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_VALUES_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_VALUES_1.GetDB()

	// Validate input
	var input orm.A_VALUES_1API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_values_1
	a_values_1DB := orm.A_VALUES_1DB{}
	a_values_1DB.A_VALUES_1PointersEncoding = input.A_VALUES_1PointersEncoding
	a_values_1DB.CopyBasicFieldsFromA_VALUES_1_WOP(&input.A_VALUES_1_WOP)

	query := db.Create(&a_values_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_VALUES_1.CheckoutPhaseOneInstance(&a_values_1DB)
	a_values_1 := backRepo.BackRepoA_VALUES_1.Map_A_VALUES_1DBID_A_VALUES_1Ptr[a_values_1DB.ID]

	if a_values_1 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_values_1)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_values_1DB)
}

// GetA_VALUES_1
//
// swagger:route GET /a_values_1s/{ID} a_values_1s getA_VALUES_1
//
// Gets the details for a a_values_1.
//
// Responses:
// default: genericError
//
//	200: a_values_1DBResponse
func (controller *Controller) GetA_VALUES_1(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_VALUES_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_VALUES_1.GetDB()

	// Get a_values_1DB in DB
	var a_values_1DB orm.A_VALUES_1DB
	if err := db.First(&a_values_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_values_1API orm.A_VALUES_1API
	a_values_1API.ID = a_values_1DB.ID
	a_values_1API.A_VALUES_1PointersEncoding = a_values_1DB.A_VALUES_1PointersEncoding
	a_values_1DB.CopyBasicFieldsToA_VALUES_1_WOP(&a_values_1API.A_VALUES_1_WOP)

	c.JSON(http.StatusOK, a_values_1API)
}

// UpdateA_VALUES_1
//
// swagger:route PATCH /a_values_1s/{ID} a_values_1s updateA_VALUES_1
//
// # Update a a_values_1
//
// Responses:
// default: genericError
//
//	200: a_values_1DBResponse
func (controller *Controller) UpdateA_VALUES_1(c *gin.Context) {

	mutexA_VALUES_1.Lock()
	defer mutexA_VALUES_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_VALUES_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_VALUES_1.GetDB()

	// Validate input
	var input orm.A_VALUES_1API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_values_1DB orm.A_VALUES_1DB

	// fetch the a_values_1
	query := db.First(&a_values_1DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_values_1DB.CopyBasicFieldsFromA_VALUES_1_WOP(&input.A_VALUES_1_WOP)
	a_values_1DB.A_VALUES_1PointersEncoding = input.A_VALUES_1PointersEncoding

	query = db.Model(&a_values_1DB).Updates(a_values_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_values_1New := new(models.A_VALUES_1)
	a_values_1DB.CopyBasicFieldsToA_VALUES_1(a_values_1New)

	// redeem pointers
	a_values_1DB.DecodePointers(backRepo, a_values_1New)

	// get stage instance from DB instance, and call callback function
	a_values_1Old := backRepo.BackRepoA_VALUES_1.Map_A_VALUES_1DBID_A_VALUES_1Ptr[a_values_1DB.ID]
	if a_values_1Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_values_1Old, a_values_1New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_values_1DB
	c.JSON(http.StatusOK, a_values_1DB)
}

// DeleteA_VALUES_1
//
// swagger:route DELETE /a_values_1s/{ID} a_values_1s deleteA_VALUES_1
//
// # Delete a a_values_1
//
// default: genericError
//
//	200: a_values_1DBResponse
func (controller *Controller) DeleteA_VALUES_1(c *gin.Context) {

	mutexA_VALUES_1.Lock()
	defer mutexA_VALUES_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_VALUES_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_VALUES_1.GetDB()

	// Get model if exist
	var a_values_1DB orm.A_VALUES_1DB
	if err := db.First(&a_values_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_values_1DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_values_1Deleted := new(models.A_VALUES_1)
	a_values_1DB.CopyBasicFieldsToA_VALUES_1(a_values_1Deleted)

	// get stage instance from DB instance, and call callback function
	a_values_1Staged := backRepo.BackRepoA_VALUES_1.Map_A_VALUES_1DBID_A_VALUES_1Ptr[a_values_1DB.ID]
	if a_values_1Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_values_1Staged, a_values_1Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
