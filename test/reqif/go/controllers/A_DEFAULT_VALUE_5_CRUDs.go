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
var __A_DEFAULT_VALUE_5__dummysDeclaration__ models.A_DEFAULT_VALUE_5
var __A_DEFAULT_VALUE_5_time__dummyDeclaration time.Duration

var mutexA_DEFAULT_VALUE_5 sync.Mutex

// An A_DEFAULT_VALUE_5ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DEFAULT_VALUE_5 updateA_DEFAULT_VALUE_5 deleteA_DEFAULT_VALUE_5
type A_DEFAULT_VALUE_5ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DEFAULT_VALUE_5Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DEFAULT_VALUE_5 updateA_DEFAULT_VALUE_5
type A_DEFAULT_VALUE_5Input struct {
	// The A_DEFAULT_VALUE_5 to submit or modify
	// in: body
	A_DEFAULT_VALUE_5 *orm.A_DEFAULT_VALUE_5API
}

// GetA_DEFAULT_VALUE_5s
//
// swagger:route GET /a_default_value_5s a_default_value_5s getA_DEFAULT_VALUE_5s
//
// # Get all a_default_value_5s
//
// Responses:
// default: genericError
//
//	200: a_default_value_5DBResponse
func (controller *Controller) GetA_DEFAULT_VALUE_5s(c *gin.Context) {

	// source slice
	var a_default_value_5DBs []orm.A_DEFAULT_VALUE_5DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUE_5s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_5.GetDB()

	query := db.Find(&a_default_value_5DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_default_value_5APIs := make([]orm.A_DEFAULT_VALUE_5API, 0)

	// for each a_default_value_5, update fields from the database nullable fields
	for idx := range a_default_value_5DBs {
		a_default_value_5DB := &a_default_value_5DBs[idx]
		_ = a_default_value_5DB
		var a_default_value_5API orm.A_DEFAULT_VALUE_5API

		// insertion point for updating fields
		a_default_value_5API.ID = a_default_value_5DB.ID
		a_default_value_5DB.CopyBasicFieldsToA_DEFAULT_VALUE_5_WOP(&a_default_value_5API.A_DEFAULT_VALUE_5_WOP)
		a_default_value_5API.A_DEFAULT_VALUE_5PointersEncoding = a_default_value_5DB.A_DEFAULT_VALUE_5PointersEncoding
		a_default_value_5APIs = append(a_default_value_5APIs, a_default_value_5API)
	}

	c.JSON(http.StatusOK, a_default_value_5APIs)
}

// PostA_DEFAULT_VALUE_5
//
// swagger:route POST /a_default_value_5s a_default_value_5s postA_DEFAULT_VALUE_5
//
// Creates a a_default_value_5
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DEFAULT_VALUE_5(c *gin.Context) {

	mutexA_DEFAULT_VALUE_5.Lock()
	defer mutexA_DEFAULT_VALUE_5.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DEFAULT_VALUE_5s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_5.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUE_5API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_default_value_5
	a_default_value_5DB := orm.A_DEFAULT_VALUE_5DB{}
	a_default_value_5DB.A_DEFAULT_VALUE_5PointersEncoding = input.A_DEFAULT_VALUE_5PointersEncoding
	a_default_value_5DB.CopyBasicFieldsFromA_DEFAULT_VALUE_5_WOP(&input.A_DEFAULT_VALUE_5_WOP)

	query := db.Create(&a_default_value_5DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DEFAULT_VALUE_5.CheckoutPhaseOneInstance(&a_default_value_5DB)
	a_default_value_5 := backRepo.BackRepoA_DEFAULT_VALUE_5.Map_A_DEFAULT_VALUE_5DBID_A_DEFAULT_VALUE_5Ptr[a_default_value_5DB.ID]

	if a_default_value_5 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_default_value_5)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_default_value_5DB)
}

// GetA_DEFAULT_VALUE_5
//
// swagger:route GET /a_default_value_5s/{ID} a_default_value_5s getA_DEFAULT_VALUE_5
//
// Gets the details for a a_default_value_5.
//
// Responses:
// default: genericError
//
//	200: a_default_value_5DBResponse
func (controller *Controller) GetA_DEFAULT_VALUE_5(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUE_5", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_5.GetDB()

	// Get a_default_value_5DB in DB
	var a_default_value_5DB orm.A_DEFAULT_VALUE_5DB
	if err := db.First(&a_default_value_5DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_default_value_5API orm.A_DEFAULT_VALUE_5API
	a_default_value_5API.ID = a_default_value_5DB.ID
	a_default_value_5API.A_DEFAULT_VALUE_5PointersEncoding = a_default_value_5DB.A_DEFAULT_VALUE_5PointersEncoding
	a_default_value_5DB.CopyBasicFieldsToA_DEFAULT_VALUE_5_WOP(&a_default_value_5API.A_DEFAULT_VALUE_5_WOP)

	c.JSON(http.StatusOK, a_default_value_5API)
}

// UpdateA_DEFAULT_VALUE_5
//
// swagger:route PATCH /a_default_value_5s/{ID} a_default_value_5s updateA_DEFAULT_VALUE_5
//
// # Update a a_default_value_5
//
// Responses:
// default: genericError
//
//	200: a_default_value_5DBResponse
func (controller *Controller) UpdateA_DEFAULT_VALUE_5(c *gin.Context) {

	mutexA_DEFAULT_VALUE_5.Lock()
	defer mutexA_DEFAULT_VALUE_5.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DEFAULT_VALUE_5", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_5.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUE_5API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_default_value_5DB orm.A_DEFAULT_VALUE_5DB

	// fetch the a_default_value_5
	query := db.First(&a_default_value_5DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_default_value_5DB.CopyBasicFieldsFromA_DEFAULT_VALUE_5_WOP(&input.A_DEFAULT_VALUE_5_WOP)
	a_default_value_5DB.A_DEFAULT_VALUE_5PointersEncoding = input.A_DEFAULT_VALUE_5PointersEncoding

	query = db.Model(&a_default_value_5DB).Updates(a_default_value_5DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_default_value_5New := new(models.A_DEFAULT_VALUE_5)
	a_default_value_5DB.CopyBasicFieldsToA_DEFAULT_VALUE_5(a_default_value_5New)

	// redeem pointers
	a_default_value_5DB.DecodePointers(backRepo, a_default_value_5New)

	// get stage instance from DB instance, and call callback function
	a_default_value_5Old := backRepo.BackRepoA_DEFAULT_VALUE_5.Map_A_DEFAULT_VALUE_5DBID_A_DEFAULT_VALUE_5Ptr[a_default_value_5DB.ID]
	if a_default_value_5Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_default_value_5Old, a_default_value_5New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_default_value_5DB
	c.JSON(http.StatusOK, a_default_value_5DB)
}

// DeleteA_DEFAULT_VALUE_5
//
// swagger:route DELETE /a_default_value_5s/{ID} a_default_value_5s deleteA_DEFAULT_VALUE_5
//
// # Delete a a_default_value_5
//
// default: genericError
//
//	200: a_default_value_5DBResponse
func (controller *Controller) DeleteA_DEFAULT_VALUE_5(c *gin.Context) {

	mutexA_DEFAULT_VALUE_5.Lock()
	defer mutexA_DEFAULT_VALUE_5.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DEFAULT_VALUE_5", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_5.GetDB()

	// Get model if exist
	var a_default_value_5DB orm.A_DEFAULT_VALUE_5DB
	if err := db.First(&a_default_value_5DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_default_value_5DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_default_value_5Deleted := new(models.A_DEFAULT_VALUE_5)
	a_default_value_5DB.CopyBasicFieldsToA_DEFAULT_VALUE_5(a_default_value_5Deleted)

	// get stage instance from DB instance, and call callback function
	a_default_value_5Staged := backRepo.BackRepoA_DEFAULT_VALUE_5.Map_A_DEFAULT_VALUE_5DBID_A_DEFAULT_VALUE_5Ptr[a_default_value_5DB.ID]
	if a_default_value_5Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_default_value_5Staged, a_default_value_5Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
