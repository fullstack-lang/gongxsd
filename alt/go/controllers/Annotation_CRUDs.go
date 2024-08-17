// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxsd/alt/go/models"
	"github.com/fullstack-lang/gongxsd/alt/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Annotation__dummysDeclaration__ models.Annotation
var __Annotation_time__dummyDeclaration time.Duration

var mutexAnnotation sync.Mutex

// An AnnotationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAnnotation updateAnnotation deleteAnnotation
type AnnotationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AnnotationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAnnotation updateAnnotation
type AnnotationInput struct {
	// The Annotation to submit or modify
	// in: body
	Annotation *orm.AnnotationAPI
}

// GetAnnotations
//
// swagger:route GET /annotations annotations getAnnotations
//
// # Get all annotations
//
// Responses:
// default: genericError
//
//	200: annotationDBResponse
func (controller *Controller) GetAnnotations(c *gin.Context) {

	// source slice
	var annotationDBs []orm.AnnotationDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAnnotations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/alt/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnnotation.GetDB()

	query := db.Find(&annotationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	annotationAPIs := make([]orm.AnnotationAPI, 0)

	// for each annotation, update fields from the database nullable fields
	for idx := range annotationDBs {
		annotationDB := &annotationDBs[idx]
		_ = annotationDB
		var annotationAPI orm.AnnotationAPI

		// insertion point for updating fields
		annotationAPI.ID = annotationDB.ID
		annotationDB.CopyBasicFieldsToAnnotation_WOP(&annotationAPI.Annotation_WOP)
		annotationAPI.AnnotationPointersEncoding = annotationDB.AnnotationPointersEncoding
		annotationAPIs = append(annotationAPIs, annotationAPI)
	}

	c.JSON(http.StatusOK, annotationAPIs)
}

// PostAnnotation
//
// swagger:route POST /annotations annotations postAnnotation
//
// Creates a annotation
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAnnotation(c *gin.Context) {

	mutexAnnotation.Lock()
	defer mutexAnnotation.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAnnotations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/alt/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnnotation.GetDB()

	// Validate input
	var input orm.AnnotationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create annotation
	annotationDB := orm.AnnotationDB{}
	annotationDB.AnnotationPointersEncoding = input.AnnotationPointersEncoding
	annotationDB.CopyBasicFieldsFromAnnotation_WOP(&input.Annotation_WOP)

	query := db.Create(&annotationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAnnotation.CheckoutPhaseOneInstance(&annotationDB)
	annotation := backRepo.BackRepoAnnotation.Map_AnnotationDBID_AnnotationPtr[annotationDB.ID]

	if annotation != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), annotation)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, annotationDB)
}

// GetAnnotation
//
// swagger:route GET /annotations/{ID} annotations getAnnotation
//
// Gets the details for a annotation.
//
// Responses:
// default: genericError
//
//	200: annotationDBResponse
func (controller *Controller) GetAnnotation(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAnnotation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/alt/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnnotation.GetDB()

	// Get annotationDB in DB
	var annotationDB orm.AnnotationDB
	if err := db.First(&annotationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var annotationAPI orm.AnnotationAPI
	annotationAPI.ID = annotationDB.ID
	annotationAPI.AnnotationPointersEncoding = annotationDB.AnnotationPointersEncoding
	annotationDB.CopyBasicFieldsToAnnotation_WOP(&annotationAPI.Annotation_WOP)

	c.JSON(http.StatusOK, annotationAPI)
}

// UpdateAnnotation
//
// swagger:route PATCH /annotations/{ID} annotations updateAnnotation
//
// # Update a annotation
//
// Responses:
// default: genericError
//
//	200: annotationDBResponse
func (controller *Controller) UpdateAnnotation(c *gin.Context) {

	mutexAnnotation.Lock()
	defer mutexAnnotation.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAnnotation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/alt/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnnotation.GetDB()

	// Validate input
	var input orm.AnnotationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var annotationDB orm.AnnotationDB

	// fetch the annotation
	query := db.First(&annotationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	annotationDB.CopyBasicFieldsFromAnnotation_WOP(&input.Annotation_WOP)
	annotationDB.AnnotationPointersEncoding = input.AnnotationPointersEncoding

	query = db.Model(&annotationDB).Updates(annotationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	annotationNew := new(models.Annotation)
	annotationDB.CopyBasicFieldsToAnnotation(annotationNew)

	// redeem pointers
	annotationDB.DecodePointers(backRepo, annotationNew)

	// get stage instance from DB instance, and call callback function
	annotationOld := backRepo.BackRepoAnnotation.Map_AnnotationDBID_AnnotationPtr[annotationDB.ID]
	if annotationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), annotationOld, annotationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the annotationDB
	c.JSON(http.StatusOK, annotationDB)
}

// DeleteAnnotation
//
// swagger:route DELETE /annotations/{ID} annotations deleteAnnotation
//
// # Delete a annotation
//
// default: genericError
//
//	200: annotationDBResponse
func (controller *Controller) DeleteAnnotation(c *gin.Context) {

	mutexAnnotation.Lock()
	defer mutexAnnotation.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAnnotation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/alt/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnnotation.GetDB()

	// Get model if exist
	var annotationDB orm.AnnotationDB
	if err := db.First(&annotationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&annotationDB)

	// get an instance (not staged) from DB instance, and call callback function
	annotationDeleted := new(models.Annotation)
	annotationDB.CopyBasicFieldsToAnnotation(annotationDeleted)

	// get stage instance from DB instance, and call callback function
	annotationStaged := backRepo.BackRepoAnnotation.Map_AnnotationDBID_AnnotationPtr[annotationDB.ID]
	if annotationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), annotationStaged, annotationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
