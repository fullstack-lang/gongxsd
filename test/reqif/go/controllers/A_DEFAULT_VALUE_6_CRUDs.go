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
var __A_DEFAULT_VALUE_6__dummysDeclaration__ models.A_DEFAULT_VALUE_6
var __A_DEFAULT_VALUE_6_time__dummyDeclaration time.Duration

var mutexA_DEFAULT_VALUE_6 sync.Mutex

// An A_DEFAULT_VALUE_6ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DEFAULT_VALUE_6 updateA_DEFAULT_VALUE_6 deleteA_DEFAULT_VALUE_6
type A_DEFAULT_VALUE_6ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DEFAULT_VALUE_6Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DEFAULT_VALUE_6 updateA_DEFAULT_VALUE_6
type A_DEFAULT_VALUE_6Input struct {
	// The A_DEFAULT_VALUE_6 to submit or modify
	// in: body
	A_DEFAULT_VALUE_6 *orm.A_DEFAULT_VALUE_6API
}

// GetA_DEFAULT_VALUE_6s
//
// swagger:route GET /a_default_value_6s a_default_value_6s getA_DEFAULT_VALUE_6s
//
// # Get all a_default_value_6s
//
// Responses:
// default: genericError
//
//	200: a_default_value_6DBResponse
func (controller *Controller) GetA_DEFAULT_VALUE_6s(c *gin.Context) {

	// source slice
	var a_default_value_6DBs []orm.A_DEFAULT_VALUE_6DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUE_6s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_6.GetDB()

	query := db.Find(&a_default_value_6DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_default_value_6APIs := make([]orm.A_DEFAULT_VALUE_6API, 0)

	// for each a_default_value_6, update fields from the database nullable fields
	for idx := range a_default_value_6DBs {
		a_default_value_6DB := &a_default_value_6DBs[idx]
		_ = a_default_value_6DB
		var a_default_value_6API orm.A_DEFAULT_VALUE_6API

		// insertion point for updating fields
		a_default_value_6API.ID = a_default_value_6DB.ID
		a_default_value_6DB.CopyBasicFieldsToA_DEFAULT_VALUE_6_WOP(&a_default_value_6API.A_DEFAULT_VALUE_6_WOP)
		a_default_value_6API.A_DEFAULT_VALUE_6PointersEncoding = a_default_value_6DB.A_DEFAULT_VALUE_6PointersEncoding
		a_default_value_6APIs = append(a_default_value_6APIs, a_default_value_6API)
	}

	c.JSON(http.StatusOK, a_default_value_6APIs)
}

// PostA_DEFAULT_VALUE_6
//
// swagger:route POST /a_default_value_6s a_default_value_6s postA_DEFAULT_VALUE_6
//
// Creates a a_default_value_6
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DEFAULT_VALUE_6(c *gin.Context) {

	mutexA_DEFAULT_VALUE_6.Lock()
	defer mutexA_DEFAULT_VALUE_6.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DEFAULT_VALUE_6s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_6.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUE_6API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_default_value_6
	a_default_value_6DB := orm.A_DEFAULT_VALUE_6DB{}
	a_default_value_6DB.A_DEFAULT_VALUE_6PointersEncoding = input.A_DEFAULT_VALUE_6PointersEncoding
	a_default_value_6DB.CopyBasicFieldsFromA_DEFAULT_VALUE_6_WOP(&input.A_DEFAULT_VALUE_6_WOP)

	query := db.Create(&a_default_value_6DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DEFAULT_VALUE_6.CheckoutPhaseOneInstance(&a_default_value_6DB)
	a_default_value_6 := backRepo.BackRepoA_DEFAULT_VALUE_6.Map_A_DEFAULT_VALUE_6DBID_A_DEFAULT_VALUE_6Ptr[a_default_value_6DB.ID]

	if a_default_value_6 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_default_value_6)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_default_value_6DB)
}

// GetA_DEFAULT_VALUE_6
//
// swagger:route GET /a_default_value_6s/{ID} a_default_value_6s getA_DEFAULT_VALUE_6
//
// Gets the details for a a_default_value_6.
//
// Responses:
// default: genericError
//
//	200: a_default_value_6DBResponse
func (controller *Controller) GetA_DEFAULT_VALUE_6(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUE_6", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_6.GetDB()

	// Get a_default_value_6DB in DB
	var a_default_value_6DB orm.A_DEFAULT_VALUE_6DB
	if err := db.First(&a_default_value_6DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_default_value_6API orm.A_DEFAULT_VALUE_6API
	a_default_value_6API.ID = a_default_value_6DB.ID
	a_default_value_6API.A_DEFAULT_VALUE_6PointersEncoding = a_default_value_6DB.A_DEFAULT_VALUE_6PointersEncoding
	a_default_value_6DB.CopyBasicFieldsToA_DEFAULT_VALUE_6_WOP(&a_default_value_6API.A_DEFAULT_VALUE_6_WOP)

	c.JSON(http.StatusOK, a_default_value_6API)
}

// UpdateA_DEFAULT_VALUE_6
//
// swagger:route PATCH /a_default_value_6s/{ID} a_default_value_6s updateA_DEFAULT_VALUE_6
//
// # Update a a_default_value_6
//
// Responses:
// default: genericError
//
//	200: a_default_value_6DBResponse
func (controller *Controller) UpdateA_DEFAULT_VALUE_6(c *gin.Context) {

	mutexA_DEFAULT_VALUE_6.Lock()
	defer mutexA_DEFAULT_VALUE_6.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DEFAULT_VALUE_6", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_6.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUE_6API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_default_value_6DB orm.A_DEFAULT_VALUE_6DB

	// fetch the a_default_value_6
	query := db.First(&a_default_value_6DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_default_value_6DB.CopyBasicFieldsFromA_DEFAULT_VALUE_6_WOP(&input.A_DEFAULT_VALUE_6_WOP)
	a_default_value_6DB.A_DEFAULT_VALUE_6PointersEncoding = input.A_DEFAULT_VALUE_6PointersEncoding

	query = db.Model(&a_default_value_6DB).Updates(a_default_value_6DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_default_value_6New := new(models.A_DEFAULT_VALUE_6)
	a_default_value_6DB.CopyBasicFieldsToA_DEFAULT_VALUE_6(a_default_value_6New)

	// redeem pointers
	a_default_value_6DB.DecodePointers(backRepo, a_default_value_6New)

	// get stage instance from DB instance, and call callback function
	a_default_value_6Old := backRepo.BackRepoA_DEFAULT_VALUE_6.Map_A_DEFAULT_VALUE_6DBID_A_DEFAULT_VALUE_6Ptr[a_default_value_6DB.ID]
	if a_default_value_6Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_default_value_6Old, a_default_value_6New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_default_value_6DB
	c.JSON(http.StatusOK, a_default_value_6DB)
}

// DeleteA_DEFAULT_VALUE_6
//
// swagger:route DELETE /a_default_value_6s/{ID} a_default_value_6s deleteA_DEFAULT_VALUE_6
//
// # Delete a a_default_value_6
//
// default: genericError
//
//	200: a_default_value_6DBResponse
func (controller *Controller) DeleteA_DEFAULT_VALUE_6(c *gin.Context) {

	mutexA_DEFAULT_VALUE_6.Lock()
	defer mutexA_DEFAULT_VALUE_6.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DEFAULT_VALUE_6", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_6.GetDB()

	// Get model if exist
	var a_default_value_6DB orm.A_DEFAULT_VALUE_6DB
	if err := db.First(&a_default_value_6DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_default_value_6DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_default_value_6Deleted := new(models.A_DEFAULT_VALUE_6)
	a_default_value_6DB.CopyBasicFieldsToA_DEFAULT_VALUE_6(a_default_value_6Deleted)

	// get stage instance from DB instance, and call callback function
	a_default_value_6Staged := backRepo.BackRepoA_DEFAULT_VALUE_6.Map_A_DEFAULT_VALUE_6DBID_A_DEFAULT_VALUE_6Ptr[a_default_value_6DB.ID]
	if a_default_value_6Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_default_value_6Staged, a_default_value_6Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
