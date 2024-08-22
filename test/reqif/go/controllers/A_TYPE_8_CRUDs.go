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
var __A_TYPE_8__dummysDeclaration__ models.A_TYPE_8
var __A_TYPE_8_time__dummyDeclaration time.Duration

var mutexA_TYPE_8 sync.Mutex

// An A_TYPE_8ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TYPE_8 updateA_TYPE_8 deleteA_TYPE_8
type A_TYPE_8ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TYPE_8Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TYPE_8 updateA_TYPE_8
type A_TYPE_8Input struct {
	// The A_TYPE_8 to submit or modify
	// in: body
	A_TYPE_8 *orm.A_TYPE_8API
}

// GetA_TYPE_8s
//
// swagger:route GET /a_type_8s a_type_8s getA_TYPE_8s
//
// # Get all a_type_8s
//
// Responses:
// default: genericError
//
//	200: a_type_8DBResponse
func (controller *Controller) GetA_TYPE_8s(c *gin.Context) {

	// source slice
	var a_type_8DBs []orm.A_TYPE_8DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_8s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_8.GetDB()

	query := db.Find(&a_type_8DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_type_8APIs := make([]orm.A_TYPE_8API, 0)

	// for each a_type_8, update fields from the database nullable fields
	for idx := range a_type_8DBs {
		a_type_8DB := &a_type_8DBs[idx]
		_ = a_type_8DB
		var a_type_8API orm.A_TYPE_8API

		// insertion point for updating fields
		a_type_8API.ID = a_type_8DB.ID
		a_type_8DB.CopyBasicFieldsToA_TYPE_8_WOP(&a_type_8API.A_TYPE_8_WOP)
		a_type_8API.A_TYPE_8PointersEncoding = a_type_8DB.A_TYPE_8PointersEncoding
		a_type_8APIs = append(a_type_8APIs, a_type_8API)
	}

	c.JSON(http.StatusOK, a_type_8APIs)
}

// PostA_TYPE_8
//
// swagger:route POST /a_type_8s a_type_8s postA_TYPE_8
//
// Creates a a_type_8
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TYPE_8(c *gin.Context) {

	mutexA_TYPE_8.Lock()
	defer mutexA_TYPE_8.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TYPE_8s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_8.GetDB()

	// Validate input
	var input orm.A_TYPE_8API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_type_8
	a_type_8DB := orm.A_TYPE_8DB{}
	a_type_8DB.A_TYPE_8PointersEncoding = input.A_TYPE_8PointersEncoding
	a_type_8DB.CopyBasicFieldsFromA_TYPE_8_WOP(&input.A_TYPE_8_WOP)

	query := db.Create(&a_type_8DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TYPE_8.CheckoutPhaseOneInstance(&a_type_8DB)
	a_type_8 := backRepo.BackRepoA_TYPE_8.Map_A_TYPE_8DBID_A_TYPE_8Ptr[a_type_8DB.ID]

	if a_type_8 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_type_8)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_type_8DB)
}

// GetA_TYPE_8
//
// swagger:route GET /a_type_8s/{ID} a_type_8s getA_TYPE_8
//
// Gets the details for a a_type_8.
//
// Responses:
// default: genericError
//
//	200: a_type_8DBResponse
func (controller *Controller) GetA_TYPE_8(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_8", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_8.GetDB()

	// Get a_type_8DB in DB
	var a_type_8DB orm.A_TYPE_8DB
	if err := db.First(&a_type_8DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_type_8API orm.A_TYPE_8API
	a_type_8API.ID = a_type_8DB.ID
	a_type_8API.A_TYPE_8PointersEncoding = a_type_8DB.A_TYPE_8PointersEncoding
	a_type_8DB.CopyBasicFieldsToA_TYPE_8_WOP(&a_type_8API.A_TYPE_8_WOP)

	c.JSON(http.StatusOK, a_type_8API)
}

// UpdateA_TYPE_8
//
// swagger:route PATCH /a_type_8s/{ID} a_type_8s updateA_TYPE_8
//
// # Update a a_type_8
//
// Responses:
// default: genericError
//
//	200: a_type_8DBResponse
func (controller *Controller) UpdateA_TYPE_8(c *gin.Context) {

	mutexA_TYPE_8.Lock()
	defer mutexA_TYPE_8.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_TYPE_8", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_8.GetDB()

	// Validate input
	var input orm.A_TYPE_8API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_type_8DB orm.A_TYPE_8DB

	// fetch the a_type_8
	query := db.First(&a_type_8DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_type_8DB.CopyBasicFieldsFromA_TYPE_8_WOP(&input.A_TYPE_8_WOP)
	a_type_8DB.A_TYPE_8PointersEncoding = input.A_TYPE_8PointersEncoding

	query = db.Model(&a_type_8DB).Updates(a_type_8DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_type_8New := new(models.A_TYPE_8)
	a_type_8DB.CopyBasicFieldsToA_TYPE_8(a_type_8New)

	// redeem pointers
	a_type_8DB.DecodePointers(backRepo, a_type_8New)

	// get stage instance from DB instance, and call callback function
	a_type_8Old := backRepo.BackRepoA_TYPE_8.Map_A_TYPE_8DBID_A_TYPE_8Ptr[a_type_8DB.ID]
	if a_type_8Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_type_8Old, a_type_8New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_type_8DB
	c.JSON(http.StatusOK, a_type_8DB)
}

// DeleteA_TYPE_8
//
// swagger:route DELETE /a_type_8s/{ID} a_type_8s deleteA_TYPE_8
//
// # Delete a a_type_8
//
// default: genericError
//
//	200: a_type_8DBResponse
func (controller *Controller) DeleteA_TYPE_8(c *gin.Context) {

	mutexA_TYPE_8.Lock()
	defer mutexA_TYPE_8.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TYPE_8", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_8.GetDB()

	// Get model if exist
	var a_type_8DB orm.A_TYPE_8DB
	if err := db.First(&a_type_8DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_type_8DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_type_8Deleted := new(models.A_TYPE_8)
	a_type_8DB.CopyBasicFieldsToA_TYPE_8(a_type_8Deleted)

	// get stage instance from DB instance, and call callback function
	a_type_8Staged := backRepo.BackRepoA_TYPE_8.Map_A_TYPE_8DBID_A_TYPE_8Ptr[a_type_8DB.ID]
	if a_type_8Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_type_8Staged, a_type_8Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
