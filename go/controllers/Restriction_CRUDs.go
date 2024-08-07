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
var __Restriction__dummysDeclaration__ models.Restriction
var __Restriction_time__dummyDeclaration time.Duration

var mutexRestriction sync.Mutex

// An RestrictionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRestriction updateRestriction deleteRestriction
type RestrictionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RestrictionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRestriction updateRestriction
type RestrictionInput struct {
	// The Restriction to submit or modify
	// in: body
	Restriction *orm.RestrictionAPI
}

// GetRestrictions
//
// swagger:route GET /restrictions restrictions getRestrictions
//
// # Get all restrictions
//
// Responses:
// default: genericError
//
//	200: restrictionDBResponse
func (controller *Controller) GetRestrictions(c *gin.Context) {

	// source slice
	var restrictionDBs []orm.RestrictionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRestrictions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRestriction.GetDB()

	query := db.Find(&restrictionDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	restrictionAPIs := make([]orm.RestrictionAPI, 0)

	// for each restriction, update fields from the database nullable fields
	for idx := range restrictionDBs {
		restrictionDB := &restrictionDBs[idx]
		_ = restrictionDB
		var restrictionAPI orm.RestrictionAPI

		// insertion point for updating fields
		restrictionAPI.ID = restrictionDB.ID
		restrictionDB.CopyBasicFieldsToRestriction_WOP(&restrictionAPI.Restriction_WOP)
		restrictionAPI.RestrictionPointersEncoding = restrictionDB.RestrictionPointersEncoding
		restrictionAPIs = append(restrictionAPIs, restrictionAPI)
	}

	c.JSON(http.StatusOK, restrictionAPIs)
}

// PostRestriction
//
// swagger:route POST /restrictions restrictions postRestriction
//
// Creates a restriction
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRestriction(c *gin.Context) {

	mutexRestriction.Lock()
	defer mutexRestriction.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRestrictions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRestriction.GetDB()

	// Validate input
	var input orm.RestrictionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create restriction
	restrictionDB := orm.RestrictionDB{}
	restrictionDB.RestrictionPointersEncoding = input.RestrictionPointersEncoding
	restrictionDB.CopyBasicFieldsFromRestriction_WOP(&input.Restriction_WOP)

	query := db.Create(&restrictionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRestriction.CheckoutPhaseOneInstance(&restrictionDB)
	restriction := backRepo.BackRepoRestriction.Map_RestrictionDBID_RestrictionPtr[restrictionDB.ID]

	if restriction != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), restriction)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, restrictionDB)
}

// GetRestriction
//
// swagger:route GET /restrictions/{ID} restrictions getRestriction
//
// Gets the details for a restriction.
//
// Responses:
// default: genericError
//
//	200: restrictionDBResponse
func (controller *Controller) GetRestriction(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRestriction", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRestriction.GetDB()

	// Get restrictionDB in DB
	var restrictionDB orm.RestrictionDB
	if err := db.First(&restrictionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var restrictionAPI orm.RestrictionAPI
	restrictionAPI.ID = restrictionDB.ID
	restrictionAPI.RestrictionPointersEncoding = restrictionDB.RestrictionPointersEncoding
	restrictionDB.CopyBasicFieldsToRestriction_WOP(&restrictionAPI.Restriction_WOP)

	c.JSON(http.StatusOK, restrictionAPI)
}

// UpdateRestriction
//
// swagger:route PATCH /restrictions/{ID} restrictions updateRestriction
//
// # Update a restriction
//
// Responses:
// default: genericError
//
//	200: restrictionDBResponse
func (controller *Controller) UpdateRestriction(c *gin.Context) {

	mutexRestriction.Lock()
	defer mutexRestriction.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRestriction", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRestriction.GetDB()

	// Validate input
	var input orm.RestrictionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var restrictionDB orm.RestrictionDB

	// fetch the restriction
	query := db.First(&restrictionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	restrictionDB.CopyBasicFieldsFromRestriction_WOP(&input.Restriction_WOP)
	restrictionDB.RestrictionPointersEncoding = input.RestrictionPointersEncoding

	query = db.Model(&restrictionDB).Updates(restrictionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	restrictionNew := new(models.Restriction)
	restrictionDB.CopyBasicFieldsToRestriction(restrictionNew)

	// redeem pointers
	restrictionDB.DecodePointers(backRepo, restrictionNew)

	// get stage instance from DB instance, and call callback function
	restrictionOld := backRepo.BackRepoRestriction.Map_RestrictionDBID_RestrictionPtr[restrictionDB.ID]
	if restrictionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), restrictionOld, restrictionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the restrictionDB
	c.JSON(http.StatusOK, restrictionDB)
}

// DeleteRestriction
//
// swagger:route DELETE /restrictions/{ID} restrictions deleteRestriction
//
// # Delete a restriction
//
// default: genericError
//
//	200: restrictionDBResponse
func (controller *Controller) DeleteRestriction(c *gin.Context) {

	mutexRestriction.Lock()
	defer mutexRestriction.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRestriction", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRestriction.GetDB()

	// Get model if exist
	var restrictionDB orm.RestrictionDB
	if err := db.First(&restrictionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&restrictionDB)

	// get an instance (not staged) from DB instance, and call callback function
	restrictionDeleted := new(models.Restriction)
	restrictionDB.CopyBasicFieldsToRestriction(restrictionDeleted)

	// get stage instance from DB instance, and call callback function
	restrictionStaged := backRepo.BackRepoRestriction.Map_RestrictionDBID_RestrictionPtr[restrictionDB.ID]
	if restrictionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), restrictionStaged, restrictionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
