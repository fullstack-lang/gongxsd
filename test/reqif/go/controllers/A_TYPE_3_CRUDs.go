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
var __A_TYPE_3__dummysDeclaration__ models.A_TYPE_3
var __A_TYPE_3_time__dummyDeclaration time.Duration

var mutexA_TYPE_3 sync.Mutex

// An A_TYPE_3ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TYPE_3 updateA_TYPE_3 deleteA_TYPE_3
type A_TYPE_3ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TYPE_3Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TYPE_3 updateA_TYPE_3
type A_TYPE_3Input struct {
	// The A_TYPE_3 to submit or modify
	// in: body
	A_TYPE_3 *orm.A_TYPE_3API
}

// GetA_TYPE_3s
//
// swagger:route GET /a_type_3s a_type_3s getA_TYPE_3s
//
// # Get all a_type_3s
//
// Responses:
// default: genericError
//
//	200: a_type_3DBResponse
func (controller *Controller) GetA_TYPE_3s(c *gin.Context) {

	// source slice
	var a_type_3DBs []orm.A_TYPE_3DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_3s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_3.GetDB()

	query := db.Find(&a_type_3DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_type_3APIs := make([]orm.A_TYPE_3API, 0)

	// for each a_type_3, update fields from the database nullable fields
	for idx := range a_type_3DBs {
		a_type_3DB := &a_type_3DBs[idx]
		_ = a_type_3DB
		var a_type_3API orm.A_TYPE_3API

		// insertion point for updating fields
		a_type_3API.ID = a_type_3DB.ID
		a_type_3DB.CopyBasicFieldsToA_TYPE_3_WOP(&a_type_3API.A_TYPE_3_WOP)
		a_type_3API.A_TYPE_3PointersEncoding = a_type_3DB.A_TYPE_3PointersEncoding
		a_type_3APIs = append(a_type_3APIs, a_type_3API)
	}

	c.JSON(http.StatusOK, a_type_3APIs)
}

// PostA_TYPE_3
//
// swagger:route POST /a_type_3s a_type_3s postA_TYPE_3
//
// Creates a a_type_3
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TYPE_3(c *gin.Context) {

	mutexA_TYPE_3.Lock()
	defer mutexA_TYPE_3.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TYPE_3s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_3.GetDB()

	// Validate input
	var input orm.A_TYPE_3API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_type_3
	a_type_3DB := orm.A_TYPE_3DB{}
	a_type_3DB.A_TYPE_3PointersEncoding = input.A_TYPE_3PointersEncoding
	a_type_3DB.CopyBasicFieldsFromA_TYPE_3_WOP(&input.A_TYPE_3_WOP)

	query := db.Create(&a_type_3DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TYPE_3.CheckoutPhaseOneInstance(&a_type_3DB)
	a_type_3 := backRepo.BackRepoA_TYPE_3.Map_A_TYPE_3DBID_A_TYPE_3Ptr[a_type_3DB.ID]

	if a_type_3 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_type_3)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_type_3DB)
}

// GetA_TYPE_3
//
// swagger:route GET /a_type_3s/{ID} a_type_3s getA_TYPE_3
//
// Gets the details for a a_type_3.
//
// Responses:
// default: genericError
//
//	200: a_type_3DBResponse
func (controller *Controller) GetA_TYPE_3(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_3", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_3.GetDB()

	// Get a_type_3DB in DB
	var a_type_3DB orm.A_TYPE_3DB
	if err := db.First(&a_type_3DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_type_3API orm.A_TYPE_3API
	a_type_3API.ID = a_type_3DB.ID
	a_type_3API.A_TYPE_3PointersEncoding = a_type_3DB.A_TYPE_3PointersEncoding
	a_type_3DB.CopyBasicFieldsToA_TYPE_3_WOP(&a_type_3API.A_TYPE_3_WOP)

	c.JSON(http.StatusOK, a_type_3API)
}

// UpdateA_TYPE_3
//
// swagger:route PATCH /a_type_3s/{ID} a_type_3s updateA_TYPE_3
//
// # Update a a_type_3
//
// Responses:
// default: genericError
//
//	200: a_type_3DBResponse
func (controller *Controller) UpdateA_TYPE_3(c *gin.Context) {

	mutexA_TYPE_3.Lock()
	defer mutexA_TYPE_3.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_TYPE_3", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_3.GetDB()

	// Validate input
	var input orm.A_TYPE_3API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_type_3DB orm.A_TYPE_3DB

	// fetch the a_type_3
	query := db.First(&a_type_3DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_type_3DB.CopyBasicFieldsFromA_TYPE_3_WOP(&input.A_TYPE_3_WOP)
	a_type_3DB.A_TYPE_3PointersEncoding = input.A_TYPE_3PointersEncoding

	query = db.Model(&a_type_3DB).Updates(a_type_3DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_type_3New := new(models.A_TYPE_3)
	a_type_3DB.CopyBasicFieldsToA_TYPE_3(a_type_3New)

	// redeem pointers
	a_type_3DB.DecodePointers(backRepo, a_type_3New)

	// get stage instance from DB instance, and call callback function
	a_type_3Old := backRepo.BackRepoA_TYPE_3.Map_A_TYPE_3DBID_A_TYPE_3Ptr[a_type_3DB.ID]
	if a_type_3Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_type_3Old, a_type_3New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_type_3DB
	c.JSON(http.StatusOK, a_type_3DB)
}

// DeleteA_TYPE_3
//
// swagger:route DELETE /a_type_3s/{ID} a_type_3s deleteA_TYPE_3
//
// # Delete a a_type_3
//
// default: genericError
//
//	200: a_type_3DBResponse
func (controller *Controller) DeleteA_TYPE_3(c *gin.Context) {

	mutexA_TYPE_3.Lock()
	defer mutexA_TYPE_3.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TYPE_3", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_3.GetDB()

	// Get model if exist
	var a_type_3DB orm.A_TYPE_3DB
	if err := db.First(&a_type_3DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_type_3DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_type_3Deleted := new(models.A_TYPE_3)
	a_type_3DB.CopyBasicFieldsToA_TYPE_3(a_type_3Deleted)

	// get stage instance from DB instance, and call callback function
	a_type_3Staged := backRepo.BackRepoA_TYPE_3.Map_A_TYPE_3DBID_A_TYPE_3Ptr[a_type_3DB.ID]
	if a_type_3Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_type_3Staged, a_type_3Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
