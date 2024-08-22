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
var __A_TYPE_1__dummysDeclaration__ models.A_TYPE_1
var __A_TYPE_1_time__dummyDeclaration time.Duration

var mutexA_TYPE_1 sync.Mutex

// An A_TYPE_1ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TYPE_1 updateA_TYPE_1 deleteA_TYPE_1
type A_TYPE_1ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TYPE_1Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TYPE_1 updateA_TYPE_1
type A_TYPE_1Input struct {
	// The A_TYPE_1 to submit or modify
	// in: body
	A_TYPE_1 *orm.A_TYPE_1API
}

// GetA_TYPE_1s
//
// swagger:route GET /a_type_1s a_type_1s getA_TYPE_1s
//
// # Get all a_type_1s
//
// Responses:
// default: genericError
//
//	200: a_type_1DBResponse
func (controller *Controller) GetA_TYPE_1s(c *gin.Context) {

	// source slice
	var a_type_1DBs []orm.A_TYPE_1DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_1.GetDB()

	query := db.Find(&a_type_1DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_type_1APIs := make([]orm.A_TYPE_1API, 0)

	// for each a_type_1, update fields from the database nullable fields
	for idx := range a_type_1DBs {
		a_type_1DB := &a_type_1DBs[idx]
		_ = a_type_1DB
		var a_type_1API orm.A_TYPE_1API

		// insertion point for updating fields
		a_type_1API.ID = a_type_1DB.ID
		a_type_1DB.CopyBasicFieldsToA_TYPE_1_WOP(&a_type_1API.A_TYPE_1_WOP)
		a_type_1API.A_TYPE_1PointersEncoding = a_type_1DB.A_TYPE_1PointersEncoding
		a_type_1APIs = append(a_type_1APIs, a_type_1API)
	}

	c.JSON(http.StatusOK, a_type_1APIs)
}

// PostA_TYPE_1
//
// swagger:route POST /a_type_1s a_type_1s postA_TYPE_1
//
// Creates a a_type_1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TYPE_1(c *gin.Context) {

	mutexA_TYPE_1.Lock()
	defer mutexA_TYPE_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TYPE_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_1.GetDB()

	// Validate input
	var input orm.A_TYPE_1API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_type_1
	a_type_1DB := orm.A_TYPE_1DB{}
	a_type_1DB.A_TYPE_1PointersEncoding = input.A_TYPE_1PointersEncoding
	a_type_1DB.CopyBasicFieldsFromA_TYPE_1_WOP(&input.A_TYPE_1_WOP)

	query := db.Create(&a_type_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TYPE_1.CheckoutPhaseOneInstance(&a_type_1DB)
	a_type_1 := backRepo.BackRepoA_TYPE_1.Map_A_TYPE_1DBID_A_TYPE_1Ptr[a_type_1DB.ID]

	if a_type_1 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_type_1)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_type_1DB)
}

// GetA_TYPE_1
//
// swagger:route GET /a_type_1s/{ID} a_type_1s getA_TYPE_1
//
// Gets the details for a a_type_1.
//
// Responses:
// default: genericError
//
//	200: a_type_1DBResponse
func (controller *Controller) GetA_TYPE_1(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_1.GetDB()

	// Get a_type_1DB in DB
	var a_type_1DB orm.A_TYPE_1DB
	if err := db.First(&a_type_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_type_1API orm.A_TYPE_1API
	a_type_1API.ID = a_type_1DB.ID
	a_type_1API.A_TYPE_1PointersEncoding = a_type_1DB.A_TYPE_1PointersEncoding
	a_type_1DB.CopyBasicFieldsToA_TYPE_1_WOP(&a_type_1API.A_TYPE_1_WOP)

	c.JSON(http.StatusOK, a_type_1API)
}

// UpdateA_TYPE_1
//
// swagger:route PATCH /a_type_1s/{ID} a_type_1s updateA_TYPE_1
//
// # Update a a_type_1
//
// Responses:
// default: genericError
//
//	200: a_type_1DBResponse
func (controller *Controller) UpdateA_TYPE_1(c *gin.Context) {

	mutexA_TYPE_1.Lock()
	defer mutexA_TYPE_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_TYPE_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_1.GetDB()

	// Validate input
	var input orm.A_TYPE_1API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_type_1DB orm.A_TYPE_1DB

	// fetch the a_type_1
	query := db.First(&a_type_1DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_type_1DB.CopyBasicFieldsFromA_TYPE_1_WOP(&input.A_TYPE_1_WOP)
	a_type_1DB.A_TYPE_1PointersEncoding = input.A_TYPE_1PointersEncoding

	query = db.Model(&a_type_1DB).Updates(a_type_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_type_1New := new(models.A_TYPE_1)
	a_type_1DB.CopyBasicFieldsToA_TYPE_1(a_type_1New)

	// redeem pointers
	a_type_1DB.DecodePointers(backRepo, a_type_1New)

	// get stage instance from DB instance, and call callback function
	a_type_1Old := backRepo.BackRepoA_TYPE_1.Map_A_TYPE_1DBID_A_TYPE_1Ptr[a_type_1DB.ID]
	if a_type_1Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_type_1Old, a_type_1New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_type_1DB
	c.JSON(http.StatusOK, a_type_1DB)
}

// DeleteA_TYPE_1
//
// swagger:route DELETE /a_type_1s/{ID} a_type_1s deleteA_TYPE_1
//
// # Delete a a_type_1
//
// default: genericError
//
//	200: a_type_1DBResponse
func (controller *Controller) DeleteA_TYPE_1(c *gin.Context) {

	mutexA_TYPE_1.Lock()
	defer mutexA_TYPE_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TYPE_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_1.GetDB()

	// Get model if exist
	var a_type_1DB orm.A_TYPE_1DB
	if err := db.First(&a_type_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_type_1DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_type_1Deleted := new(models.A_TYPE_1)
	a_type_1DB.CopyBasicFieldsToA_TYPE_1(a_type_1Deleted)

	// get stage instance from DB instance, and call callback function
	a_type_1Staged := backRepo.BackRepoA_TYPE_1.Map_A_TYPE_1DBID_A_TYPE_1Ptr[a_type_1DB.ID]
	if a_type_1Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_type_1Staged, a_type_1Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
