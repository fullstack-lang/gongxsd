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
var __A_DEFAULT_VALUE_3__dummysDeclaration__ models.A_DEFAULT_VALUE_3
var __A_DEFAULT_VALUE_3_time__dummyDeclaration time.Duration

var mutexA_DEFAULT_VALUE_3 sync.Mutex

// An A_DEFAULT_VALUE_3ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DEFAULT_VALUE_3 updateA_DEFAULT_VALUE_3 deleteA_DEFAULT_VALUE_3
type A_DEFAULT_VALUE_3ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DEFAULT_VALUE_3Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DEFAULT_VALUE_3 updateA_DEFAULT_VALUE_3
type A_DEFAULT_VALUE_3Input struct {
	// The A_DEFAULT_VALUE_3 to submit or modify
	// in: body
	A_DEFAULT_VALUE_3 *orm.A_DEFAULT_VALUE_3API
}

// GetA_DEFAULT_VALUE_3s
//
// swagger:route GET /a_default_value_3s a_default_value_3s getA_DEFAULT_VALUE_3s
//
// # Get all a_default_value_3s
//
// Responses:
// default: genericError
//
//	200: a_default_value_3DBResponse
func (controller *Controller) GetA_DEFAULT_VALUE_3s(c *gin.Context) {

	// source slice
	var a_default_value_3DBs []orm.A_DEFAULT_VALUE_3DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUE_3s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_3.GetDB()

	query := db.Find(&a_default_value_3DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_default_value_3APIs := make([]orm.A_DEFAULT_VALUE_3API, 0)

	// for each a_default_value_3, update fields from the database nullable fields
	for idx := range a_default_value_3DBs {
		a_default_value_3DB := &a_default_value_3DBs[idx]
		_ = a_default_value_3DB
		var a_default_value_3API orm.A_DEFAULT_VALUE_3API

		// insertion point for updating fields
		a_default_value_3API.ID = a_default_value_3DB.ID
		a_default_value_3DB.CopyBasicFieldsToA_DEFAULT_VALUE_3_WOP(&a_default_value_3API.A_DEFAULT_VALUE_3_WOP)
		a_default_value_3API.A_DEFAULT_VALUE_3PointersEncoding = a_default_value_3DB.A_DEFAULT_VALUE_3PointersEncoding
		a_default_value_3APIs = append(a_default_value_3APIs, a_default_value_3API)
	}

	c.JSON(http.StatusOK, a_default_value_3APIs)
}

// PostA_DEFAULT_VALUE_3
//
// swagger:route POST /a_default_value_3s a_default_value_3s postA_DEFAULT_VALUE_3
//
// Creates a a_default_value_3
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DEFAULT_VALUE_3(c *gin.Context) {

	mutexA_DEFAULT_VALUE_3.Lock()
	defer mutexA_DEFAULT_VALUE_3.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DEFAULT_VALUE_3s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_3.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUE_3API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_default_value_3
	a_default_value_3DB := orm.A_DEFAULT_VALUE_3DB{}
	a_default_value_3DB.A_DEFAULT_VALUE_3PointersEncoding = input.A_DEFAULT_VALUE_3PointersEncoding
	a_default_value_3DB.CopyBasicFieldsFromA_DEFAULT_VALUE_3_WOP(&input.A_DEFAULT_VALUE_3_WOP)

	query := db.Create(&a_default_value_3DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DEFAULT_VALUE_3.CheckoutPhaseOneInstance(&a_default_value_3DB)
	a_default_value_3 := backRepo.BackRepoA_DEFAULT_VALUE_3.Map_A_DEFAULT_VALUE_3DBID_A_DEFAULT_VALUE_3Ptr[a_default_value_3DB.ID]

	if a_default_value_3 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_default_value_3)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_default_value_3DB)
}

// GetA_DEFAULT_VALUE_3
//
// swagger:route GET /a_default_value_3s/{ID} a_default_value_3s getA_DEFAULT_VALUE_3
//
// Gets the details for a a_default_value_3.
//
// Responses:
// default: genericError
//
//	200: a_default_value_3DBResponse
func (controller *Controller) GetA_DEFAULT_VALUE_3(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFAULT_VALUE_3", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_3.GetDB()

	// Get a_default_value_3DB in DB
	var a_default_value_3DB orm.A_DEFAULT_VALUE_3DB
	if err := db.First(&a_default_value_3DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_default_value_3API orm.A_DEFAULT_VALUE_3API
	a_default_value_3API.ID = a_default_value_3DB.ID
	a_default_value_3API.A_DEFAULT_VALUE_3PointersEncoding = a_default_value_3DB.A_DEFAULT_VALUE_3PointersEncoding
	a_default_value_3DB.CopyBasicFieldsToA_DEFAULT_VALUE_3_WOP(&a_default_value_3API.A_DEFAULT_VALUE_3_WOP)

	c.JSON(http.StatusOK, a_default_value_3API)
}

// UpdateA_DEFAULT_VALUE_3
//
// swagger:route PATCH /a_default_value_3s/{ID} a_default_value_3s updateA_DEFAULT_VALUE_3
//
// # Update a a_default_value_3
//
// Responses:
// default: genericError
//
//	200: a_default_value_3DBResponse
func (controller *Controller) UpdateA_DEFAULT_VALUE_3(c *gin.Context) {

	mutexA_DEFAULT_VALUE_3.Lock()
	defer mutexA_DEFAULT_VALUE_3.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DEFAULT_VALUE_3", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_3.GetDB()

	// Validate input
	var input orm.A_DEFAULT_VALUE_3API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_default_value_3DB orm.A_DEFAULT_VALUE_3DB

	// fetch the a_default_value_3
	query := db.First(&a_default_value_3DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_default_value_3DB.CopyBasicFieldsFromA_DEFAULT_VALUE_3_WOP(&input.A_DEFAULT_VALUE_3_WOP)
	a_default_value_3DB.A_DEFAULT_VALUE_3PointersEncoding = input.A_DEFAULT_VALUE_3PointersEncoding

	query = db.Model(&a_default_value_3DB).Updates(a_default_value_3DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_default_value_3New := new(models.A_DEFAULT_VALUE_3)
	a_default_value_3DB.CopyBasicFieldsToA_DEFAULT_VALUE_3(a_default_value_3New)

	// redeem pointers
	a_default_value_3DB.DecodePointers(backRepo, a_default_value_3New)

	// get stage instance from DB instance, and call callback function
	a_default_value_3Old := backRepo.BackRepoA_DEFAULT_VALUE_3.Map_A_DEFAULT_VALUE_3DBID_A_DEFAULT_VALUE_3Ptr[a_default_value_3DB.ID]
	if a_default_value_3Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_default_value_3Old, a_default_value_3New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_default_value_3DB
	c.JSON(http.StatusOK, a_default_value_3DB)
}

// DeleteA_DEFAULT_VALUE_3
//
// swagger:route DELETE /a_default_value_3s/{ID} a_default_value_3s deleteA_DEFAULT_VALUE_3
//
// # Delete a a_default_value_3
//
// default: genericError
//
//	200: a_default_value_3DBResponse
func (controller *Controller) DeleteA_DEFAULT_VALUE_3(c *gin.Context) {

	mutexA_DEFAULT_VALUE_3.Lock()
	defer mutexA_DEFAULT_VALUE_3.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DEFAULT_VALUE_3", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFAULT_VALUE_3.GetDB()

	// Get model if exist
	var a_default_value_3DB orm.A_DEFAULT_VALUE_3DB
	if err := db.First(&a_default_value_3DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_default_value_3DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_default_value_3Deleted := new(models.A_DEFAULT_VALUE_3)
	a_default_value_3DB.CopyBasicFieldsToA_DEFAULT_VALUE_3(a_default_value_3Deleted)

	// get stage instance from DB instance, and call callback function
	a_default_value_3Staged := backRepo.BackRepoA_DEFAULT_VALUE_3.Map_A_DEFAULT_VALUE_3DBID_A_DEFAULT_VALUE_3Ptr[a_default_value_3DB.ID]
	if a_default_value_3Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_default_value_3Staged, a_default_value_3Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
