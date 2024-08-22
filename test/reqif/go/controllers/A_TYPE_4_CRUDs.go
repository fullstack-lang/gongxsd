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
var __A_TYPE_4__dummysDeclaration__ models.A_TYPE_4
var __A_TYPE_4_time__dummyDeclaration time.Duration

var mutexA_TYPE_4 sync.Mutex

// An A_TYPE_4ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TYPE_4 updateA_TYPE_4 deleteA_TYPE_4
type A_TYPE_4ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TYPE_4Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TYPE_4 updateA_TYPE_4
type A_TYPE_4Input struct {
	// The A_TYPE_4 to submit or modify
	// in: body
	A_TYPE_4 *orm.A_TYPE_4API
}

// GetA_TYPE_4s
//
// swagger:route GET /a_type_4s a_type_4s getA_TYPE_4s
//
// # Get all a_type_4s
//
// Responses:
// default: genericError
//
//	200: a_type_4DBResponse
func (controller *Controller) GetA_TYPE_4s(c *gin.Context) {

	// source slice
	var a_type_4DBs []orm.A_TYPE_4DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_4s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_4.GetDB()

	query := db.Find(&a_type_4DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_type_4APIs := make([]orm.A_TYPE_4API, 0)

	// for each a_type_4, update fields from the database nullable fields
	for idx := range a_type_4DBs {
		a_type_4DB := &a_type_4DBs[idx]
		_ = a_type_4DB
		var a_type_4API orm.A_TYPE_4API

		// insertion point for updating fields
		a_type_4API.ID = a_type_4DB.ID
		a_type_4DB.CopyBasicFieldsToA_TYPE_4_WOP(&a_type_4API.A_TYPE_4_WOP)
		a_type_4API.A_TYPE_4PointersEncoding = a_type_4DB.A_TYPE_4PointersEncoding
		a_type_4APIs = append(a_type_4APIs, a_type_4API)
	}

	c.JSON(http.StatusOK, a_type_4APIs)
}

// PostA_TYPE_4
//
// swagger:route POST /a_type_4s a_type_4s postA_TYPE_4
//
// Creates a a_type_4
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TYPE_4(c *gin.Context) {

	mutexA_TYPE_4.Lock()
	defer mutexA_TYPE_4.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TYPE_4s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_4.GetDB()

	// Validate input
	var input orm.A_TYPE_4API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_type_4
	a_type_4DB := orm.A_TYPE_4DB{}
	a_type_4DB.A_TYPE_4PointersEncoding = input.A_TYPE_4PointersEncoding
	a_type_4DB.CopyBasicFieldsFromA_TYPE_4_WOP(&input.A_TYPE_4_WOP)

	query := db.Create(&a_type_4DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TYPE_4.CheckoutPhaseOneInstance(&a_type_4DB)
	a_type_4 := backRepo.BackRepoA_TYPE_4.Map_A_TYPE_4DBID_A_TYPE_4Ptr[a_type_4DB.ID]

	if a_type_4 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_type_4)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_type_4DB)
}

// GetA_TYPE_4
//
// swagger:route GET /a_type_4s/{ID} a_type_4s getA_TYPE_4
//
// Gets the details for a a_type_4.
//
// Responses:
// default: genericError
//
//	200: a_type_4DBResponse
func (controller *Controller) GetA_TYPE_4(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_4", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_4.GetDB()

	// Get a_type_4DB in DB
	var a_type_4DB orm.A_TYPE_4DB
	if err := db.First(&a_type_4DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_type_4API orm.A_TYPE_4API
	a_type_4API.ID = a_type_4DB.ID
	a_type_4API.A_TYPE_4PointersEncoding = a_type_4DB.A_TYPE_4PointersEncoding
	a_type_4DB.CopyBasicFieldsToA_TYPE_4_WOP(&a_type_4API.A_TYPE_4_WOP)

	c.JSON(http.StatusOK, a_type_4API)
}

// UpdateA_TYPE_4
//
// swagger:route PATCH /a_type_4s/{ID} a_type_4s updateA_TYPE_4
//
// # Update a a_type_4
//
// Responses:
// default: genericError
//
//	200: a_type_4DBResponse
func (controller *Controller) UpdateA_TYPE_4(c *gin.Context) {

	mutexA_TYPE_4.Lock()
	defer mutexA_TYPE_4.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_TYPE_4", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_4.GetDB()

	// Validate input
	var input orm.A_TYPE_4API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_type_4DB orm.A_TYPE_4DB

	// fetch the a_type_4
	query := db.First(&a_type_4DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_type_4DB.CopyBasicFieldsFromA_TYPE_4_WOP(&input.A_TYPE_4_WOP)
	a_type_4DB.A_TYPE_4PointersEncoding = input.A_TYPE_4PointersEncoding

	query = db.Model(&a_type_4DB).Updates(a_type_4DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_type_4New := new(models.A_TYPE_4)
	a_type_4DB.CopyBasicFieldsToA_TYPE_4(a_type_4New)

	// redeem pointers
	a_type_4DB.DecodePointers(backRepo, a_type_4New)

	// get stage instance from DB instance, and call callback function
	a_type_4Old := backRepo.BackRepoA_TYPE_4.Map_A_TYPE_4DBID_A_TYPE_4Ptr[a_type_4DB.ID]
	if a_type_4Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_type_4Old, a_type_4New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_type_4DB
	c.JSON(http.StatusOK, a_type_4DB)
}

// DeleteA_TYPE_4
//
// swagger:route DELETE /a_type_4s/{ID} a_type_4s deleteA_TYPE_4
//
// # Delete a a_type_4
//
// default: genericError
//
//	200: a_type_4DBResponse
func (controller *Controller) DeleteA_TYPE_4(c *gin.Context) {

	mutexA_TYPE_4.Lock()
	defer mutexA_TYPE_4.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TYPE_4", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_4.GetDB()

	// Get model if exist
	var a_type_4DB orm.A_TYPE_4DB
	if err := db.First(&a_type_4DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_type_4DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_type_4Deleted := new(models.A_TYPE_4)
	a_type_4DB.CopyBasicFieldsToA_TYPE_4(a_type_4Deleted)

	// get stage instance from DB instance, and call callback function
	a_type_4Staged := backRepo.BackRepoA_TYPE_4.Map_A_TYPE_4DBID_A_TYPE_4Ptr[a_type_4DB.ID]
	if a_type_4Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_type_4Staged, a_type_4Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
