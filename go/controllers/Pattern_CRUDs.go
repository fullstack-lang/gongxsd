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
var __Pattern__dummysDeclaration__ models.Pattern
var __Pattern_time__dummyDeclaration time.Duration

var mutexPattern sync.Mutex

// An PatternID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPattern updatePattern deletePattern
type PatternID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PatternInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPattern updatePattern
type PatternInput struct {
	// The Pattern to submit or modify
	// in: body
	Pattern *orm.PatternAPI
}

// GetPatterns
//
// swagger:route GET /patterns patterns getPatterns
//
// # Get all patterns
//
// Responses:
// default: genericError
//
//	200: patternDBResponse
func (controller *Controller) GetPatterns(c *gin.Context) {

	// source slice
	var patternDBs []orm.PatternDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPatterns", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPattern.GetDB()

	query := db.Find(&patternDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	patternAPIs := make([]orm.PatternAPI, 0)

	// for each pattern, update fields from the database nullable fields
	for idx := range patternDBs {
		patternDB := &patternDBs[idx]
		_ = patternDB
		var patternAPI orm.PatternAPI

		// insertion point for updating fields
		patternAPI.ID = patternDB.ID
		patternDB.CopyBasicFieldsToPattern_WOP(&patternAPI.Pattern_WOP)
		patternAPI.PatternPointersEncoding = patternDB.PatternPointersEncoding
		patternAPIs = append(patternAPIs, patternAPI)
	}

	c.JSON(http.StatusOK, patternAPIs)
}

// PostPattern
//
// swagger:route POST /patterns patterns postPattern
//
// Creates a pattern
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPattern(c *gin.Context) {

	mutexPattern.Lock()
	defer mutexPattern.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPatterns", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPattern.GetDB()

	// Validate input
	var input orm.PatternAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create pattern
	patternDB := orm.PatternDB{}
	patternDB.PatternPointersEncoding = input.PatternPointersEncoding
	patternDB.CopyBasicFieldsFromPattern_WOP(&input.Pattern_WOP)

	query := db.Create(&patternDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPattern.CheckoutPhaseOneInstance(&patternDB)
	pattern := backRepo.BackRepoPattern.Map_PatternDBID_PatternPtr[patternDB.ID]

	if pattern != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), pattern)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, patternDB)
}

// GetPattern
//
// swagger:route GET /patterns/{ID} patterns getPattern
//
// Gets the details for a pattern.
//
// Responses:
// default: genericError
//
//	200: patternDBResponse
func (controller *Controller) GetPattern(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPattern", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPattern.GetDB()

	// Get patternDB in DB
	var patternDB orm.PatternDB
	if err := db.First(&patternDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var patternAPI orm.PatternAPI
	patternAPI.ID = patternDB.ID
	patternAPI.PatternPointersEncoding = patternDB.PatternPointersEncoding
	patternDB.CopyBasicFieldsToPattern_WOP(&patternAPI.Pattern_WOP)

	c.JSON(http.StatusOK, patternAPI)
}

// UpdatePattern
//
// swagger:route PATCH /patterns/{ID} patterns updatePattern
//
// # Update a pattern
//
// Responses:
// default: genericError
//
//	200: patternDBResponse
func (controller *Controller) UpdatePattern(c *gin.Context) {

	mutexPattern.Lock()
	defer mutexPattern.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePattern", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPattern.GetDB()

	// Validate input
	var input orm.PatternAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var patternDB orm.PatternDB

	// fetch the pattern
	query := db.First(&patternDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	patternDB.CopyBasicFieldsFromPattern_WOP(&input.Pattern_WOP)
	patternDB.PatternPointersEncoding = input.PatternPointersEncoding

	query = db.Model(&patternDB).Updates(patternDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	patternNew := new(models.Pattern)
	patternDB.CopyBasicFieldsToPattern(patternNew)

	// redeem pointers
	patternDB.DecodePointers(backRepo, patternNew)

	// get stage instance from DB instance, and call callback function
	patternOld := backRepo.BackRepoPattern.Map_PatternDBID_PatternPtr[patternDB.ID]
	if patternOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), patternOld, patternNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the patternDB
	c.JSON(http.StatusOK, patternDB)
}

// DeletePattern
//
// swagger:route DELETE /patterns/{ID} patterns deletePattern
//
// # Delete a pattern
//
// default: genericError
//
//	200: patternDBResponse
func (controller *Controller) DeletePattern(c *gin.Context) {

	mutexPattern.Lock()
	defer mutexPattern.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePattern", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPattern.GetDB()

	// Get model if exist
	var patternDB orm.PatternDB
	if err := db.First(&patternDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&patternDB)

	// get an instance (not staged) from DB instance, and call callback function
	patternDeleted := new(models.Pattern)
	patternDB.CopyBasicFieldsToPattern(patternDeleted)

	// get stage instance from DB instance, and call callback function
	patternStaged := backRepo.BackRepoPattern.Map_PatternDBID_PatternPtr[patternDB.ID]
	if patternStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), patternStaged, patternDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
