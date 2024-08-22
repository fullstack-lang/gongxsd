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
var __A_DEFINITION_3__dummysDeclaration__ models.A_DEFINITION_3
var __A_DEFINITION_3_time__dummyDeclaration time.Duration

var mutexA_DEFINITION_3 sync.Mutex

// An A_DEFINITION_3ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DEFINITION_3 updateA_DEFINITION_3 deleteA_DEFINITION_3
type A_DEFINITION_3ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DEFINITION_3Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DEFINITION_3 updateA_DEFINITION_3
type A_DEFINITION_3Input struct {
	// The A_DEFINITION_3 to submit or modify
	// in: body
	A_DEFINITION_3 *orm.A_DEFINITION_3API
}

// GetA_DEFINITION_3s
//
// swagger:route GET /a_definition_3s a_definition_3s getA_DEFINITION_3s
//
// # Get all a_definition_3s
//
// Responses:
// default: genericError
//
//	200: a_definition_3DBResponse
func (controller *Controller) GetA_DEFINITION_3s(c *gin.Context) {

	// source slice
	var a_definition_3DBs []orm.A_DEFINITION_3DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFINITION_3s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_3.GetDB()

	query := db.Find(&a_definition_3DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_definition_3APIs := make([]orm.A_DEFINITION_3API, 0)

	// for each a_definition_3, update fields from the database nullable fields
	for idx := range a_definition_3DBs {
		a_definition_3DB := &a_definition_3DBs[idx]
		_ = a_definition_3DB
		var a_definition_3API orm.A_DEFINITION_3API

		// insertion point for updating fields
		a_definition_3API.ID = a_definition_3DB.ID
		a_definition_3DB.CopyBasicFieldsToA_DEFINITION_3_WOP(&a_definition_3API.A_DEFINITION_3_WOP)
		a_definition_3API.A_DEFINITION_3PointersEncoding = a_definition_3DB.A_DEFINITION_3PointersEncoding
		a_definition_3APIs = append(a_definition_3APIs, a_definition_3API)
	}

	c.JSON(http.StatusOK, a_definition_3APIs)
}

// PostA_DEFINITION_3
//
// swagger:route POST /a_definition_3s a_definition_3s postA_DEFINITION_3
//
// Creates a a_definition_3
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DEFINITION_3(c *gin.Context) {

	mutexA_DEFINITION_3.Lock()
	defer mutexA_DEFINITION_3.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DEFINITION_3s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_3.GetDB()

	// Validate input
	var input orm.A_DEFINITION_3API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_definition_3
	a_definition_3DB := orm.A_DEFINITION_3DB{}
	a_definition_3DB.A_DEFINITION_3PointersEncoding = input.A_DEFINITION_3PointersEncoding
	a_definition_3DB.CopyBasicFieldsFromA_DEFINITION_3_WOP(&input.A_DEFINITION_3_WOP)

	query := db.Create(&a_definition_3DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DEFINITION_3.CheckoutPhaseOneInstance(&a_definition_3DB)
	a_definition_3 := backRepo.BackRepoA_DEFINITION_3.Map_A_DEFINITION_3DBID_A_DEFINITION_3Ptr[a_definition_3DB.ID]

	if a_definition_3 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_definition_3)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_definition_3DB)
}

// GetA_DEFINITION_3
//
// swagger:route GET /a_definition_3s/{ID} a_definition_3s getA_DEFINITION_3
//
// Gets the details for a a_definition_3.
//
// Responses:
// default: genericError
//
//	200: a_definition_3DBResponse
func (controller *Controller) GetA_DEFINITION_3(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DEFINITION_3", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_3.GetDB()

	// Get a_definition_3DB in DB
	var a_definition_3DB orm.A_DEFINITION_3DB
	if err := db.First(&a_definition_3DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_definition_3API orm.A_DEFINITION_3API
	a_definition_3API.ID = a_definition_3DB.ID
	a_definition_3API.A_DEFINITION_3PointersEncoding = a_definition_3DB.A_DEFINITION_3PointersEncoding
	a_definition_3DB.CopyBasicFieldsToA_DEFINITION_3_WOP(&a_definition_3API.A_DEFINITION_3_WOP)

	c.JSON(http.StatusOK, a_definition_3API)
}

// UpdateA_DEFINITION_3
//
// swagger:route PATCH /a_definition_3s/{ID} a_definition_3s updateA_DEFINITION_3
//
// # Update a a_definition_3
//
// Responses:
// default: genericError
//
//	200: a_definition_3DBResponse
func (controller *Controller) UpdateA_DEFINITION_3(c *gin.Context) {

	mutexA_DEFINITION_3.Lock()
	defer mutexA_DEFINITION_3.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DEFINITION_3", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_3.GetDB()

	// Validate input
	var input orm.A_DEFINITION_3API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_definition_3DB orm.A_DEFINITION_3DB

	// fetch the a_definition_3
	query := db.First(&a_definition_3DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_definition_3DB.CopyBasicFieldsFromA_DEFINITION_3_WOP(&input.A_DEFINITION_3_WOP)
	a_definition_3DB.A_DEFINITION_3PointersEncoding = input.A_DEFINITION_3PointersEncoding

	query = db.Model(&a_definition_3DB).Updates(a_definition_3DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_definition_3New := new(models.A_DEFINITION_3)
	a_definition_3DB.CopyBasicFieldsToA_DEFINITION_3(a_definition_3New)

	// redeem pointers
	a_definition_3DB.DecodePointers(backRepo, a_definition_3New)

	// get stage instance from DB instance, and call callback function
	a_definition_3Old := backRepo.BackRepoA_DEFINITION_3.Map_A_DEFINITION_3DBID_A_DEFINITION_3Ptr[a_definition_3DB.ID]
	if a_definition_3Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_definition_3Old, a_definition_3New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_definition_3DB
	c.JSON(http.StatusOK, a_definition_3DB)
}

// DeleteA_DEFINITION_3
//
// swagger:route DELETE /a_definition_3s/{ID} a_definition_3s deleteA_DEFINITION_3
//
// # Delete a a_definition_3
//
// default: genericError
//
//	200: a_definition_3DBResponse
func (controller *Controller) DeleteA_DEFINITION_3(c *gin.Context) {

	mutexA_DEFINITION_3.Lock()
	defer mutexA_DEFINITION_3.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DEFINITION_3", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_DEFINITION_3.GetDB()

	// Get model if exist
	var a_definition_3DB orm.A_DEFINITION_3DB
	if err := db.First(&a_definition_3DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_definition_3DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_definition_3Deleted := new(models.A_DEFINITION_3)
	a_definition_3DB.CopyBasicFieldsToA_DEFINITION_3(a_definition_3Deleted)

	// get stage instance from DB instance, and call callback function
	a_definition_3Staged := backRepo.BackRepoA_DEFINITION_3.Map_A_DEFINITION_3DBID_A_DEFINITION_3Ptr[a_definition_3DB.ID]
	if a_definition_3Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_definition_3Staged, a_definition_3Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
