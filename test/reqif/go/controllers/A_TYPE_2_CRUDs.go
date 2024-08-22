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
var __A_TYPE_2__dummysDeclaration__ models.A_TYPE_2
var __A_TYPE_2_time__dummyDeclaration time.Duration

var mutexA_TYPE_2 sync.Mutex

// An A_TYPE_2ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TYPE_2 updateA_TYPE_2 deleteA_TYPE_2
type A_TYPE_2ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TYPE_2Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TYPE_2 updateA_TYPE_2
type A_TYPE_2Input struct {
	// The A_TYPE_2 to submit or modify
	// in: body
	A_TYPE_2 *orm.A_TYPE_2API
}

// GetA_TYPE_2s
//
// swagger:route GET /a_type_2s a_type_2s getA_TYPE_2s
//
// # Get all a_type_2s
//
// Responses:
// default: genericError
//
//	200: a_type_2DBResponse
func (controller *Controller) GetA_TYPE_2s(c *gin.Context) {

	// source slice
	var a_type_2DBs []orm.A_TYPE_2DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_2s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_2.GetDB()

	query := db.Find(&a_type_2DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_type_2APIs := make([]orm.A_TYPE_2API, 0)

	// for each a_type_2, update fields from the database nullable fields
	for idx := range a_type_2DBs {
		a_type_2DB := &a_type_2DBs[idx]
		_ = a_type_2DB
		var a_type_2API orm.A_TYPE_2API

		// insertion point for updating fields
		a_type_2API.ID = a_type_2DB.ID
		a_type_2DB.CopyBasicFieldsToA_TYPE_2_WOP(&a_type_2API.A_TYPE_2_WOP)
		a_type_2API.A_TYPE_2PointersEncoding = a_type_2DB.A_TYPE_2PointersEncoding
		a_type_2APIs = append(a_type_2APIs, a_type_2API)
	}

	c.JSON(http.StatusOK, a_type_2APIs)
}

// PostA_TYPE_2
//
// swagger:route POST /a_type_2s a_type_2s postA_TYPE_2
//
// Creates a a_type_2
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TYPE_2(c *gin.Context) {

	mutexA_TYPE_2.Lock()
	defer mutexA_TYPE_2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TYPE_2s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_2.GetDB()

	// Validate input
	var input orm.A_TYPE_2API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_type_2
	a_type_2DB := orm.A_TYPE_2DB{}
	a_type_2DB.A_TYPE_2PointersEncoding = input.A_TYPE_2PointersEncoding
	a_type_2DB.CopyBasicFieldsFromA_TYPE_2_WOP(&input.A_TYPE_2_WOP)

	query := db.Create(&a_type_2DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TYPE_2.CheckoutPhaseOneInstance(&a_type_2DB)
	a_type_2 := backRepo.BackRepoA_TYPE_2.Map_A_TYPE_2DBID_A_TYPE_2Ptr[a_type_2DB.ID]

	if a_type_2 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_type_2)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_type_2DB)
}

// GetA_TYPE_2
//
// swagger:route GET /a_type_2s/{ID} a_type_2s getA_TYPE_2
//
// Gets the details for a a_type_2.
//
// Responses:
// default: genericError
//
//	200: a_type_2DBResponse
func (controller *Controller) GetA_TYPE_2(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TYPE_2", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_2.GetDB()

	// Get a_type_2DB in DB
	var a_type_2DB orm.A_TYPE_2DB
	if err := db.First(&a_type_2DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_type_2API orm.A_TYPE_2API
	a_type_2API.ID = a_type_2DB.ID
	a_type_2API.A_TYPE_2PointersEncoding = a_type_2DB.A_TYPE_2PointersEncoding
	a_type_2DB.CopyBasicFieldsToA_TYPE_2_WOP(&a_type_2API.A_TYPE_2_WOP)

	c.JSON(http.StatusOK, a_type_2API)
}

// UpdateA_TYPE_2
//
// swagger:route PATCH /a_type_2s/{ID} a_type_2s updateA_TYPE_2
//
// # Update a a_type_2
//
// Responses:
// default: genericError
//
//	200: a_type_2DBResponse
func (controller *Controller) UpdateA_TYPE_2(c *gin.Context) {

	mutexA_TYPE_2.Lock()
	defer mutexA_TYPE_2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_TYPE_2", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_2.GetDB()

	// Validate input
	var input orm.A_TYPE_2API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_type_2DB orm.A_TYPE_2DB

	// fetch the a_type_2
	query := db.First(&a_type_2DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_type_2DB.CopyBasicFieldsFromA_TYPE_2_WOP(&input.A_TYPE_2_WOP)
	a_type_2DB.A_TYPE_2PointersEncoding = input.A_TYPE_2PointersEncoding

	query = db.Model(&a_type_2DB).Updates(a_type_2DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_type_2New := new(models.A_TYPE_2)
	a_type_2DB.CopyBasicFieldsToA_TYPE_2(a_type_2New)

	// redeem pointers
	a_type_2DB.DecodePointers(backRepo, a_type_2New)

	// get stage instance from DB instance, and call callback function
	a_type_2Old := backRepo.BackRepoA_TYPE_2.Map_A_TYPE_2DBID_A_TYPE_2Ptr[a_type_2DB.ID]
	if a_type_2Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_type_2Old, a_type_2New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_type_2DB
	c.JSON(http.StatusOK, a_type_2DB)
}

// DeleteA_TYPE_2
//
// swagger:route DELETE /a_type_2s/{ID} a_type_2s deleteA_TYPE_2
//
// # Delete a a_type_2
//
// default: genericError
//
//	200: a_type_2DBResponse
func (controller *Controller) DeleteA_TYPE_2(c *gin.Context) {

	mutexA_TYPE_2.Lock()
	defer mutexA_TYPE_2.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TYPE_2", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoA_TYPE_2.GetDB()

	// Get model if exist
	var a_type_2DB orm.A_TYPE_2DB
	if err := db.First(&a_type_2DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&a_type_2DB)

	// get an instance (not staged) from DB instance, and call callback function
	a_type_2Deleted := new(models.A_TYPE_2)
	a_type_2DB.CopyBasicFieldsToA_TYPE_2(a_type_2Deleted)

	// get stage instance from DB instance, and call callback function
	a_type_2Staged := backRepo.BackRepoA_TYPE_2.Map_A_TYPE_2DBID_A_TYPE_2Ptr[a_type_2DB.ID]
	if a_type_2Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_type_2Staged, a_type_2Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
