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
var __AttributeGroup__dummysDeclaration__ models.AttributeGroup
var __AttributeGroup_time__dummyDeclaration time.Duration

var mutexAttributeGroup sync.Mutex

// An AttributeGroupID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAttributeGroup updateAttributeGroup deleteAttributeGroup
type AttributeGroupID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AttributeGroupInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAttributeGroup updateAttributeGroup
type AttributeGroupInput struct {
	// The AttributeGroup to submit or modify
	// in: body
	AttributeGroup *orm.AttributeGroupAPI
}

// GetAttributeGroups
//
// swagger:route GET /attributegroups attributegroups getAttributeGroups
//
// # Get all attributegroups
//
// Responses:
// default: genericError
//
//	200: attributegroupDBResponse
func (controller *Controller) GetAttributeGroups(c *gin.Context) {

	// source slice
	var attributegroupDBs []orm.AttributeGroupDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAttributeGroups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttributeGroup.GetDB()

	_, err := db.Find(&attributegroupDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributegroupAPIs := make([]orm.AttributeGroupAPI, 0)

	// for each attributegroup, update fields from the database nullable fields
	for idx := range attributegroupDBs {
		attributegroupDB := &attributegroupDBs[idx]
		_ = attributegroupDB
		var attributegroupAPI orm.AttributeGroupAPI

		// insertion point for updating fields
		attributegroupAPI.ID = attributegroupDB.ID
		attributegroupDB.CopyBasicFieldsToAttributeGroup_WOP(&attributegroupAPI.AttributeGroup_WOP)
		attributegroupAPI.AttributeGroupPointersEncoding = attributegroupDB.AttributeGroupPointersEncoding
		attributegroupAPIs = append(attributegroupAPIs, attributegroupAPI)
	}

	c.JSON(http.StatusOK, attributegroupAPIs)
}

// PostAttributeGroup
//
// swagger:route POST /attributegroups attributegroups postAttributeGroup
//
// Creates a attributegroup
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAttributeGroup(c *gin.Context) {

	mutexAttributeGroup.Lock()
	defer mutexAttributeGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAttributeGroups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttributeGroup.GetDB()

	// Validate input
	var input orm.AttributeGroupAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributegroup
	attributegroupDB := orm.AttributeGroupDB{}
	attributegroupDB.AttributeGroupPointersEncoding = input.AttributeGroupPointersEncoding
	attributegroupDB.CopyBasicFieldsFromAttributeGroup_WOP(&input.AttributeGroup_WOP)

	_, err = db.Create(&attributegroupDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAttributeGroup.CheckoutPhaseOneInstance(&attributegroupDB)
	attributegroup := backRepo.BackRepoAttributeGroup.Map_AttributeGroupDBID_AttributeGroupPtr[attributegroupDB.ID]

	if attributegroup != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributegroup)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributegroupDB)
}

// GetAttributeGroup
//
// swagger:route GET /attributegroups/{ID} attributegroups getAttributeGroup
//
// Gets the details for a attributegroup.
//
// Responses:
// default: genericError
//
//	200: attributegroupDBResponse
func (controller *Controller) GetAttributeGroup(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAttributeGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttributeGroup.GetDB()

	// Get attributegroupDB in DB
	var attributegroupDB orm.AttributeGroupDB
	if _, err := db.First(&attributegroupDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributegroupAPI orm.AttributeGroupAPI
	attributegroupAPI.ID = attributegroupDB.ID
	attributegroupAPI.AttributeGroupPointersEncoding = attributegroupDB.AttributeGroupPointersEncoding
	attributegroupDB.CopyBasicFieldsToAttributeGroup_WOP(&attributegroupAPI.AttributeGroup_WOP)

	c.JSON(http.StatusOK, attributegroupAPI)
}

// UpdateAttributeGroup
//
// swagger:route PATCH /attributegroups/{ID} attributegroups updateAttributeGroup
//
// # Update a attributegroup
//
// Responses:
// default: genericError
//
//	200: attributegroupDBResponse
func (controller *Controller) UpdateAttributeGroup(c *gin.Context) {

	mutexAttributeGroup.Lock()
	defer mutexAttributeGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAttributeGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttributeGroup.GetDB()

	// Validate input
	var input orm.AttributeGroupAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributegroupDB orm.AttributeGroupDB

	// fetch the attributegroup
	_, err := db.First(&attributegroupDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributegroupDB.CopyBasicFieldsFromAttributeGroup_WOP(&input.AttributeGroup_WOP)
	attributegroupDB.AttributeGroupPointersEncoding = input.AttributeGroupPointersEncoding

	db, _ = db.Model(&attributegroupDB)
	_, err = db.Updates(&attributegroupDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributegroupNew := new(models.AttributeGroup)
	attributegroupDB.CopyBasicFieldsToAttributeGroup(attributegroupNew)

	// redeem pointers
	attributegroupDB.DecodePointers(backRepo, attributegroupNew)

	// get stage instance from DB instance, and call callback function
	attributegroupOld := backRepo.BackRepoAttributeGroup.Map_AttributeGroupDBID_AttributeGroupPtr[attributegroupDB.ID]
	if attributegroupOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributegroupOld, attributegroupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributegroupDB
	c.JSON(http.StatusOK, attributegroupDB)
}

// DeleteAttributeGroup
//
// swagger:route DELETE /attributegroups/{ID} attributegroups deleteAttributeGroup
//
// # Delete a attributegroup
//
// default: genericError
//
//	200: attributegroupDBResponse
func (controller *Controller) DeleteAttributeGroup(c *gin.Context) {

	mutexAttributeGroup.Lock()
	defer mutexAttributeGroup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAttributeGroup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttributeGroup.GetDB()

	// Get model if exist
	var attributegroupDB orm.AttributeGroupDB
	if _, err := db.First(&attributegroupDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attributegroupDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributegroupDeleted := new(models.AttributeGroup)
	attributegroupDB.CopyBasicFieldsToAttributeGroup(attributegroupDeleted)

	// get stage instance from DB instance, and call callback function
	attributegroupStaged := backRepo.BackRepoAttributeGroup.Map_AttributeGroupDBID_AttributeGroupPtr[attributegroupDB.ID]
	if attributegroupStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributegroupStaged, attributegroupDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
