// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxsd/go/models"
	"github.com/fullstack-lang/gongxsd/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __All__dummysDeclaration__ models.All
var __All_time__dummyDeclaration time.Duration

var mutexAll sync.Mutex

// An AllID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAll updateAll deleteAll
type AllID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AllInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAll updateAll
type AllInput struct {
	// The All to submit or modify
	// in: body
	All *orm.AllAPI
}

// GetAlls
//
// swagger:route GET /alls alls getAlls
//
// # Get all alls
//
// Responses:
// default: genericError
//
//	200: allDBResponse
func (controller *Controller) GetAlls(c *gin.Context) {

	// source slice
	var allDBs []orm.AllDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAlls", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAll.GetDB()

	query := db.Find(&allDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	allAPIs := make([]orm.AllAPI, 0)

	// for each all, update fields from the database nullable fields
	for idx := range allDBs {
		allDB := &allDBs[idx]
		_ = allDB
		var allAPI orm.AllAPI

		// insertion point for updating fields
		allAPI.ID = allDB.ID
		allDB.CopyBasicFieldsToAll_WOP(&allAPI.All_WOP)
		allAPI.AllPointersEncoding = allDB.AllPointersEncoding
		allAPIs = append(allAPIs, allAPI)
	}

	c.JSON(http.StatusOK, allAPIs)
}

// PostAll
//
// swagger:route POST /alls alls postAll
//
// Creates a all
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAll(c *gin.Context) {

	mutexAll.Lock()
	defer mutexAll.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAlls", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAll.GetDB()

	// Validate input
	var input orm.AllAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create all
	allDB := orm.AllDB{}
	allDB.AllPointersEncoding = input.AllPointersEncoding
	allDB.CopyBasicFieldsFromAll_WOP(&input.All_WOP)

	query := db.Create(&allDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAll.CheckoutPhaseOneInstance(&allDB)
	all := backRepo.BackRepoAll.Map_AllDBID_AllPtr[allDB.ID]

	if all != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), all)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, allDB)
}

// GetAll
//
// swagger:route GET /alls/{ID} alls getAll
//
// Gets the details for a all.
//
// Responses:
// default: genericError
//
//	200: allDBResponse
func (controller *Controller) GetAll(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAll", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAll.GetDB()

	// Get allDB in DB
	var allDB orm.AllDB
	if err := db.First(&allDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var allAPI orm.AllAPI
	allAPI.ID = allDB.ID
	allAPI.AllPointersEncoding = allDB.AllPointersEncoding
	allDB.CopyBasicFieldsToAll_WOP(&allAPI.All_WOP)

	c.JSON(http.StatusOK, allAPI)
}

// UpdateAll
//
// swagger:route PATCH /alls/{ID} alls updateAll
//
// # Update a all
//
// Responses:
// default: genericError
//
//	200: allDBResponse
func (controller *Controller) UpdateAll(c *gin.Context) {

	mutexAll.Lock()
	defer mutexAll.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAll", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAll.GetDB()

	// Validate input
	var input orm.AllAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var allDB orm.AllDB

	// fetch the all
	query := db.First(&allDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	allDB.CopyBasicFieldsFromAll_WOP(&input.All_WOP)
	allDB.AllPointersEncoding = input.AllPointersEncoding

	query = db.Model(&allDB).Updates(allDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	allNew := new(models.All)
	allDB.CopyBasicFieldsToAll(allNew)

	// redeem pointers
	allDB.DecodePointers(backRepo, allNew)

	// get stage instance from DB instance, and call callback function
	allOld := backRepo.BackRepoAll.Map_AllDBID_AllPtr[allDB.ID]
	if allOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), allOld, allNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the allDB
	c.JSON(http.StatusOK, allDB)
}

// DeleteAll
//
// swagger:route DELETE /alls/{ID} alls deleteAll
//
// # Delete a all
//
// default: genericError
//
//	200: allDBResponse
func (controller *Controller) DeleteAll(c *gin.Context) {

	mutexAll.Lock()
	defer mutexAll.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAll", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAll.GetDB()

	// Get model if exist
	var allDB orm.AllDB
	if err := db.First(&allDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&allDB)

	// get an instance (not staged) from DB instance, and call callback function
	allDeleted := new(models.All)
	allDB.CopyBasicFieldsToAll(allDeleted)

	// get stage instance from DB instance, and call callback function
	allStaged := backRepo.BackRepoAll.Map_AllDBID_AllPtr[allDB.ID]
	if allStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), allStaged, allDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
