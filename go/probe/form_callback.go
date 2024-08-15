// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/go/models"
	"github.com/fullstack-lang/gongxsd/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__AllFormCallback(
	all *models.All,
	probe *Probe,
	formGroup *table.FormGroup,
) (allFormCallback *AllFormCallback) {
	allFormCallback = new(AllFormCallback)
	allFormCallback.probe = probe
	allFormCallback.all = all
	allFormCallback.formGroup = formGroup

	allFormCallback.CreationMode = (all == nil)

	return
}

type AllFormCallback struct {
	all *models.All

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (allFormCallback *AllFormCallback) OnSave() {

	log.Println("AllFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	allFormCallback.probe.formStage.Checkout()

	if allFormCallback.all == nil {
		allFormCallback.all = new(models.All).Stage(allFormCallback.probe.stageOfInterest)
	}
	all_ := allFormCallback.all
	_ = all_

	for _, formDiv := range allFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(all_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(all_.Annotation), allFormCallback.probe.stageOfInterest, formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(all_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(all_.MaxOccurs), formDiv)
		case "All:Alls":
			// we need to retrieve the field owner before the change
			var pastAllOwner *models.All
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "All"
			rf.Fieldname = "Alls"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				allFormCallback.probe.stageOfInterest,
				allFormCallback.probe.backRepoOfInterest,
				all_,
				&rf)

			if reverseFieldOwner != nil {
				pastAllOwner = reverseFieldOwner.(*models.All)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAllOwner != nil {
					idx := slices.Index(pastAllOwner.Alls, all_)
					pastAllOwner.Alls = slices.Delete(pastAllOwner.Alls, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _all := range *models.GetGongstructInstancesSet[models.All](allFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _all.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAllOwner := _all // we have a match
						if pastAllOwner != nil {
							if newAllOwner != pastAllOwner {
								idx := slices.Index(pastAllOwner.Alls, all_)
								pastAllOwner.Alls = slices.Delete(pastAllOwner.Alls, idx, idx+1)
								newAllOwner.Alls = append(newAllOwner.Alls, all_)
							}
						} else {
							newAllOwner.Alls = append(newAllOwner.Alls, all_)
						}
					}
				}
			}
		case "Choice:Alls":
			// we need to retrieve the field owner before the change
			var pastChoiceOwner *models.Choice
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Choice"
			rf.Fieldname = "Alls"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				allFormCallback.probe.stageOfInterest,
				allFormCallback.probe.backRepoOfInterest,
				all_,
				&rf)

			if reverseFieldOwner != nil {
				pastChoiceOwner = reverseFieldOwner.(*models.Choice)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastChoiceOwner != nil {
					idx := slices.Index(pastChoiceOwner.Alls, all_)
					pastChoiceOwner.Alls = slices.Delete(pastChoiceOwner.Alls, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _choice := range *models.GetGongstructInstancesSet[models.Choice](allFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _choice.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newChoiceOwner := _choice // we have a match
						if pastChoiceOwner != nil {
							if newChoiceOwner != pastChoiceOwner {
								idx := slices.Index(pastChoiceOwner.Alls, all_)
								pastChoiceOwner.Alls = slices.Delete(pastChoiceOwner.Alls, idx, idx+1)
								newChoiceOwner.Alls = append(newChoiceOwner.Alls, all_)
							}
						} else {
							newChoiceOwner.Alls = append(newChoiceOwner.Alls, all_)
						}
					}
				}
			}
		case "ComplexType:Alls":
			// we need to retrieve the field owner before the change
			var pastComplexTypeOwner *models.ComplexType
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ComplexType"
			rf.Fieldname = "Alls"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				allFormCallback.probe.stageOfInterest,
				allFormCallback.probe.backRepoOfInterest,
				all_,
				&rf)

			if reverseFieldOwner != nil {
				pastComplexTypeOwner = reverseFieldOwner.(*models.ComplexType)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastComplexTypeOwner != nil {
					idx := slices.Index(pastComplexTypeOwner.Alls, all_)
					pastComplexTypeOwner.Alls = slices.Delete(pastComplexTypeOwner.Alls, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](allFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _complextype.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newComplexTypeOwner := _complextype // we have a match
						if pastComplexTypeOwner != nil {
							if newComplexTypeOwner != pastComplexTypeOwner {
								idx := slices.Index(pastComplexTypeOwner.Alls, all_)
								pastComplexTypeOwner.Alls = slices.Delete(pastComplexTypeOwner.Alls, idx, idx+1)
								newComplexTypeOwner.Alls = append(newComplexTypeOwner.Alls, all_)
							}
						} else {
							newComplexTypeOwner.Alls = append(newComplexTypeOwner.Alls, all_)
						}
					}
				}
			}
		case "Extension:Alls":
			// we need to retrieve the field owner before the change
			var pastExtensionOwner *models.Extension
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Extension"
			rf.Fieldname = "Alls"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				allFormCallback.probe.stageOfInterest,
				allFormCallback.probe.backRepoOfInterest,
				all_,
				&rf)

			if reverseFieldOwner != nil {
				pastExtensionOwner = reverseFieldOwner.(*models.Extension)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastExtensionOwner != nil {
					idx := slices.Index(pastExtensionOwner.Alls, all_)
					pastExtensionOwner.Alls = slices.Delete(pastExtensionOwner.Alls, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _extension := range *models.GetGongstructInstancesSet[models.Extension](allFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _extension.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newExtensionOwner := _extension // we have a match
						if pastExtensionOwner != nil {
							if newExtensionOwner != pastExtensionOwner {
								idx := slices.Index(pastExtensionOwner.Alls, all_)
								pastExtensionOwner.Alls = slices.Delete(pastExtensionOwner.Alls, idx, idx+1)
								newExtensionOwner.Alls = append(newExtensionOwner.Alls, all_)
							}
						} else {
							newExtensionOwner.Alls = append(newExtensionOwner.Alls, all_)
						}
					}
				}
			}
		case "Group:Alls":
			// we need to retrieve the field owner before the change
			var pastGroupOwner *models.Group
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Alls"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				allFormCallback.probe.stageOfInterest,
				allFormCallback.probe.backRepoOfInterest,
				all_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupOwner = reverseFieldOwner.(*models.Group)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupOwner != nil {
					idx := slices.Index(pastGroupOwner.Alls, all_)
					pastGroupOwner.Alls = slices.Delete(pastGroupOwner.Alls, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _group := range *models.GetGongstructInstancesSet[models.Group](allFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupOwner := _group // we have a match
						if pastGroupOwner != nil {
							if newGroupOwner != pastGroupOwner {
								idx := slices.Index(pastGroupOwner.Alls, all_)
								pastGroupOwner.Alls = slices.Delete(pastGroupOwner.Alls, idx, idx+1)
								newGroupOwner.Alls = append(newGroupOwner.Alls, all_)
							}
						} else {
							newGroupOwner.Alls = append(newGroupOwner.Alls, all_)
						}
					}
				}
			}
		case "Sequence:Alls":
			// we need to retrieve the field owner before the change
			var pastSequenceOwner *models.Sequence
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sequence"
			rf.Fieldname = "Alls"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				allFormCallback.probe.stageOfInterest,
				allFormCallback.probe.backRepoOfInterest,
				all_,
				&rf)

			if reverseFieldOwner != nil {
				pastSequenceOwner = reverseFieldOwner.(*models.Sequence)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSequenceOwner != nil {
					idx := slices.Index(pastSequenceOwner.Alls, all_)
					pastSequenceOwner.Alls = slices.Delete(pastSequenceOwner.Alls, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _sequence := range *models.GetGongstructInstancesSet[models.Sequence](allFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _sequence.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSequenceOwner := _sequence // we have a match
						if pastSequenceOwner != nil {
							if newSequenceOwner != pastSequenceOwner {
								idx := slices.Index(pastSequenceOwner.Alls, all_)
								pastSequenceOwner.Alls = slices.Delete(pastSequenceOwner.Alls, idx, idx+1)
								newSequenceOwner.Alls = append(newSequenceOwner.Alls, all_)
							}
						} else {
							newSequenceOwner.Alls = append(newSequenceOwner.Alls, all_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if allFormCallback.formGroup.HasSuppressButtonBeenPressed {
		all_.Unstage(allFormCallback.probe.stageOfInterest)
	}

	allFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.All](
		allFormCallback.probe,
	)
	allFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if allFormCallback.CreationMode || allFormCallback.formGroup.HasSuppressButtonBeenPressed {
		allFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(allFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AllFormCallback(
			nil,
			allFormCallback.probe,
			newFormGroup,
		)
		all := new(models.All)
		FillUpForm(all, newFormGroup, allFormCallback.probe)
		allFormCallback.probe.formStage.Commit()
	}

	fillUpTree(allFormCallback.probe)
}
func __gong__New__AnnotationFormCallback(
	annotation *models.Annotation,
	probe *Probe,
	formGroup *table.FormGroup,
) (annotationFormCallback *AnnotationFormCallback) {
	annotationFormCallback = new(AnnotationFormCallback)
	annotationFormCallback.probe = probe
	annotationFormCallback.annotation = annotation
	annotationFormCallback.formGroup = formGroup

	annotationFormCallback.CreationMode = (annotation == nil)

	return
}

type AnnotationFormCallback struct {
	annotation *models.Annotation

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (annotationFormCallback *AnnotationFormCallback) OnSave() {

	log.Println("AnnotationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	annotationFormCallback.probe.formStage.Checkout()

	if annotationFormCallback.annotation == nil {
		annotationFormCallback.annotation = new(models.Annotation).Stage(annotationFormCallback.probe.stageOfInterest)
	}
	annotation_ := annotationFormCallback.annotation
	_ = annotation_

	for _, formDiv := range annotationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(annotation_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if annotationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		annotation_.Unstage(annotationFormCallback.probe.stageOfInterest)
	}

	annotationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Annotation](
		annotationFormCallback.probe,
	)
	annotationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if annotationFormCallback.CreationMode || annotationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		annotationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(annotationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AnnotationFormCallback(
			nil,
			annotationFormCallback.probe,
			newFormGroup,
		)
		annotation := new(models.Annotation)
		FillUpForm(annotation, newFormGroup, annotationFormCallback.probe)
		annotationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(annotationFormCallback.probe)
}
func __gong__New__AttributeFormCallback(
	attribute *models.Attribute,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributeFormCallback *AttributeFormCallback) {
	attributeFormCallback = new(AttributeFormCallback)
	attributeFormCallback.probe = probe
	attributeFormCallback.attribute = attribute
	attributeFormCallback.formGroup = formGroup

	attributeFormCallback.CreationMode = (attribute == nil)

	return
}

type AttributeFormCallback struct {
	attribute *models.Attribute

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributeFormCallback *AttributeFormCallback) OnSave() {

	log.Println("AttributeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributeFormCallback.probe.formStage.Checkout()

	if attributeFormCallback.attribute == nil {
		attributeFormCallback.attribute = new(models.Attribute).Stage(attributeFormCallback.probe.stageOfInterest)
	}
	attribute_ := attributeFormCallback.attribute
	_ = attribute_

	for _, formDiv := range attributeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attribute_.Name), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(attribute_.NameXSD), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(attribute_.Type), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(attribute_.Annotation), attributeFormCallback.probe.stageOfInterest, formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(attribute_.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(attribute_.GoIdentifier), formDiv)
		case "Default":
			FormDivBasicFieldToField(&(attribute_.Default), formDiv)
		case "Use":
			FormDivBasicFieldToField(&(attribute_.Use), formDiv)
		case "Form":
			FormDivBasicFieldToField(&(attribute_.Form), formDiv)
		case "Fixed":
			FormDivBasicFieldToField(&(attribute_.Fixed), formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(attribute_.Ref), formDiv)
		case "TargetNamespace":
			FormDivBasicFieldToField(&(attribute_.TargetNamespace), formDiv)
		case "SimpleType":
			FormDivBasicFieldToField(&(attribute_.SimpleType), formDiv)
		case "IDXSD":
			FormDivBasicFieldToField(&(attribute_.IDXSD), formDiv)
		case "AttributeGroup:Attributes":
			// we need to retrieve the field owner before the change
			var pastAttributeGroupOwner *models.AttributeGroup
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "AttributeGroup"
			rf.Fieldname = "Attributes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributeFormCallback.probe.stageOfInterest,
				attributeFormCallback.probe.backRepoOfInterest,
				attribute_,
				&rf)

			if reverseFieldOwner != nil {
				pastAttributeGroupOwner = reverseFieldOwner.(*models.AttributeGroup)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAttributeGroupOwner != nil {
					idx := slices.Index(pastAttributeGroupOwner.Attributes, attribute_)
					pastAttributeGroupOwner.Attributes = slices.Delete(pastAttributeGroupOwner.Attributes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attributegroup := range *models.GetGongstructInstancesSet[models.AttributeGroup](attributeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attributegroup.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAttributeGroupOwner := _attributegroup // we have a match
						if pastAttributeGroupOwner != nil {
							if newAttributeGroupOwner != pastAttributeGroupOwner {
								idx := slices.Index(pastAttributeGroupOwner.Attributes, attribute_)
								pastAttributeGroupOwner.Attributes = slices.Delete(pastAttributeGroupOwner.Attributes, idx, idx+1)
								newAttributeGroupOwner.Attributes = append(newAttributeGroupOwner.Attributes, attribute_)
							}
						} else {
							newAttributeGroupOwner.Attributes = append(newAttributeGroupOwner.Attributes, attribute_)
						}
					}
				}
			}
		case "ComplexType:Attributes":
			// we need to retrieve the field owner before the change
			var pastComplexTypeOwner *models.ComplexType
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ComplexType"
			rf.Fieldname = "Attributes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributeFormCallback.probe.stageOfInterest,
				attributeFormCallback.probe.backRepoOfInterest,
				attribute_,
				&rf)

			if reverseFieldOwner != nil {
				pastComplexTypeOwner = reverseFieldOwner.(*models.ComplexType)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastComplexTypeOwner != nil {
					idx := slices.Index(pastComplexTypeOwner.Attributes, attribute_)
					pastComplexTypeOwner.Attributes = slices.Delete(pastComplexTypeOwner.Attributes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](attributeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _complextype.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newComplexTypeOwner := _complextype // we have a match
						if pastComplexTypeOwner != nil {
							if newComplexTypeOwner != pastComplexTypeOwner {
								idx := slices.Index(pastComplexTypeOwner.Attributes, attribute_)
								pastComplexTypeOwner.Attributes = slices.Delete(pastComplexTypeOwner.Attributes, idx, idx+1)
								newComplexTypeOwner.Attributes = append(newComplexTypeOwner.Attributes, attribute_)
							}
						} else {
							newComplexTypeOwner.Attributes = append(newComplexTypeOwner.Attributes, attribute_)
						}
					}
				}
			}
		case "Extension:Attributes":
			// we need to retrieve the field owner before the change
			var pastExtensionOwner *models.Extension
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Extension"
			rf.Fieldname = "Attributes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributeFormCallback.probe.stageOfInterest,
				attributeFormCallback.probe.backRepoOfInterest,
				attribute_,
				&rf)

			if reverseFieldOwner != nil {
				pastExtensionOwner = reverseFieldOwner.(*models.Extension)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastExtensionOwner != nil {
					idx := slices.Index(pastExtensionOwner.Attributes, attribute_)
					pastExtensionOwner.Attributes = slices.Delete(pastExtensionOwner.Attributes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _extension := range *models.GetGongstructInstancesSet[models.Extension](attributeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _extension.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newExtensionOwner := _extension // we have a match
						if pastExtensionOwner != nil {
							if newExtensionOwner != pastExtensionOwner {
								idx := slices.Index(pastExtensionOwner.Attributes, attribute_)
								pastExtensionOwner.Attributes = slices.Delete(pastExtensionOwner.Attributes, idx, idx+1)
								newExtensionOwner.Attributes = append(newExtensionOwner.Attributes, attribute_)
							}
						} else {
							newExtensionOwner.Attributes = append(newExtensionOwner.Attributes, attribute_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attributeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attribute_.Unstage(attributeFormCallback.probe.stageOfInterest)
	}

	attributeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Attribute](
		attributeFormCallback.probe,
	)
	attributeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributeFormCallback.CreationMode || attributeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AttributeFormCallback(
			nil,
			attributeFormCallback.probe,
			newFormGroup,
		)
		attribute := new(models.Attribute)
		FillUpForm(attribute, newFormGroup, attributeFormCallback.probe)
		attributeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributeFormCallback.probe)
}
func __gong__New__AttributeGroupFormCallback(
	attributegroup *models.AttributeGroup,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributegroupFormCallback *AttributeGroupFormCallback) {
	attributegroupFormCallback = new(AttributeGroupFormCallback)
	attributegroupFormCallback.probe = probe
	attributegroupFormCallback.attributegroup = attributegroup
	attributegroupFormCallback.formGroup = formGroup

	attributegroupFormCallback.CreationMode = (attributegroup == nil)

	return
}

type AttributeGroupFormCallback struct {
	attributegroup *models.AttributeGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributegroupFormCallback *AttributeGroupFormCallback) OnSave() {

	log.Println("AttributeGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributegroupFormCallback.probe.formStage.Checkout()

	if attributegroupFormCallback.attributegroup == nil {
		attributegroupFormCallback.attributegroup = new(models.AttributeGroup).Stage(attributegroupFormCallback.probe.stageOfInterest)
	}
	attributegroup_ := attributegroupFormCallback.attributegroup
	_ = attributegroup_

	for _, formDiv := range attributegroupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributegroup_.Name), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(attributegroup_.NameXSD), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(attributegroup_.Annotation), attributegroupFormCallback.probe.stageOfInterest, formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(attributegroup_.Ref), formDiv)
		case "AttributeGroup:AttributeGroups":
			// we need to retrieve the field owner before the change
			var pastAttributeGroupOwner *models.AttributeGroup
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "AttributeGroup"
			rf.Fieldname = "AttributeGroups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributegroupFormCallback.probe.stageOfInterest,
				attributegroupFormCallback.probe.backRepoOfInterest,
				attributegroup_,
				&rf)

			if reverseFieldOwner != nil {
				pastAttributeGroupOwner = reverseFieldOwner.(*models.AttributeGroup)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAttributeGroupOwner != nil {
					idx := slices.Index(pastAttributeGroupOwner.AttributeGroups, attributegroup_)
					pastAttributeGroupOwner.AttributeGroups = slices.Delete(pastAttributeGroupOwner.AttributeGroups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _attributegroup := range *models.GetGongstructInstancesSet[models.AttributeGroup](attributegroupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _attributegroup.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAttributeGroupOwner := _attributegroup // we have a match
						if pastAttributeGroupOwner != nil {
							if newAttributeGroupOwner != pastAttributeGroupOwner {
								idx := slices.Index(pastAttributeGroupOwner.AttributeGroups, attributegroup_)
								pastAttributeGroupOwner.AttributeGroups = slices.Delete(pastAttributeGroupOwner.AttributeGroups, idx, idx+1)
								newAttributeGroupOwner.AttributeGroups = append(newAttributeGroupOwner.AttributeGroups, attributegroup_)
							}
						} else {
							newAttributeGroupOwner.AttributeGroups = append(newAttributeGroupOwner.AttributeGroups, attributegroup_)
						}
					}
				}
			}
		case "ComplexType:AttributeGroups":
			// we need to retrieve the field owner before the change
			var pastComplexTypeOwner *models.ComplexType
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ComplexType"
			rf.Fieldname = "AttributeGroups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributegroupFormCallback.probe.stageOfInterest,
				attributegroupFormCallback.probe.backRepoOfInterest,
				attributegroup_,
				&rf)

			if reverseFieldOwner != nil {
				pastComplexTypeOwner = reverseFieldOwner.(*models.ComplexType)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastComplexTypeOwner != nil {
					idx := slices.Index(pastComplexTypeOwner.AttributeGroups, attributegroup_)
					pastComplexTypeOwner.AttributeGroups = slices.Delete(pastComplexTypeOwner.AttributeGroups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](attributegroupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _complextype.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newComplexTypeOwner := _complextype // we have a match
						if pastComplexTypeOwner != nil {
							if newComplexTypeOwner != pastComplexTypeOwner {
								idx := slices.Index(pastComplexTypeOwner.AttributeGroups, attributegroup_)
								pastComplexTypeOwner.AttributeGroups = slices.Delete(pastComplexTypeOwner.AttributeGroups, idx, idx+1)
								newComplexTypeOwner.AttributeGroups = append(newComplexTypeOwner.AttributeGroups, attributegroup_)
							}
						} else {
							newComplexTypeOwner.AttributeGroups = append(newComplexTypeOwner.AttributeGroups, attributegroup_)
						}
					}
				}
			}
		case "Schema:AttributeGroups":
			// we need to retrieve the field owner before the change
			var pastSchemaOwner *models.Schema
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "AttributeGroups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				attributegroupFormCallback.probe.stageOfInterest,
				attributegroupFormCallback.probe.backRepoOfInterest,
				attributegroup_,
				&rf)

			if reverseFieldOwner != nil {
				pastSchemaOwner = reverseFieldOwner.(*models.Schema)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSchemaOwner != nil {
					idx := slices.Index(pastSchemaOwner.AttributeGroups, attributegroup_)
					pastSchemaOwner.AttributeGroups = slices.Delete(pastSchemaOwner.AttributeGroups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _schema := range *models.GetGongstructInstancesSet[models.Schema](attributegroupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _schema.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSchemaOwner := _schema // we have a match
						if pastSchemaOwner != nil {
							if newSchemaOwner != pastSchemaOwner {
								idx := slices.Index(pastSchemaOwner.AttributeGroups, attributegroup_)
								pastSchemaOwner.AttributeGroups = slices.Delete(pastSchemaOwner.AttributeGroups, idx, idx+1)
								newSchemaOwner.AttributeGroups = append(newSchemaOwner.AttributeGroups, attributegroup_)
							}
						} else {
							newSchemaOwner.AttributeGroups = append(newSchemaOwner.AttributeGroups, attributegroup_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attributegroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributegroup_.Unstage(attributegroupFormCallback.probe.stageOfInterest)
	}

	attributegroupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.AttributeGroup](
		attributegroupFormCallback.probe,
	)
	attributegroupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributegroupFormCallback.CreationMode || attributegroupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributegroupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(attributegroupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AttributeGroupFormCallback(
			nil,
			attributegroupFormCallback.probe,
			newFormGroup,
		)
		attributegroup := new(models.AttributeGroup)
		FillUpForm(attributegroup, newFormGroup, attributegroupFormCallback.probe)
		attributegroupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(attributegroupFormCallback.probe)
}
func __gong__New__ChoiceFormCallback(
	choice *models.Choice,
	probe *Probe,
	formGroup *table.FormGroup,
) (choiceFormCallback *ChoiceFormCallback) {
	choiceFormCallback = new(ChoiceFormCallback)
	choiceFormCallback.probe = probe
	choiceFormCallback.choice = choice
	choiceFormCallback.formGroup = formGroup

	choiceFormCallback.CreationMode = (choice == nil)

	return
}

type ChoiceFormCallback struct {
	choice *models.Choice

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (choiceFormCallback *ChoiceFormCallback) OnSave() {

	log.Println("ChoiceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	choiceFormCallback.probe.formStage.Checkout()

	if choiceFormCallback.choice == nil {
		choiceFormCallback.choice = new(models.Choice).Stage(choiceFormCallback.probe.stageOfInterest)
	}
	choice_ := choiceFormCallback.choice
	_ = choice_

	for _, formDiv := range choiceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(choice_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(choice_.Annotation), choiceFormCallback.probe.stageOfInterest, formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(choice_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(choice_.MaxOccurs), formDiv)
		case "All:Choices":
			// we need to retrieve the field owner before the change
			var pastAllOwner *models.All
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "All"
			rf.Fieldname = "Choices"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				choiceFormCallback.probe.stageOfInterest,
				choiceFormCallback.probe.backRepoOfInterest,
				choice_,
				&rf)

			if reverseFieldOwner != nil {
				pastAllOwner = reverseFieldOwner.(*models.All)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAllOwner != nil {
					idx := slices.Index(pastAllOwner.Choices, choice_)
					pastAllOwner.Choices = slices.Delete(pastAllOwner.Choices, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _all := range *models.GetGongstructInstancesSet[models.All](choiceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _all.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAllOwner := _all // we have a match
						if pastAllOwner != nil {
							if newAllOwner != pastAllOwner {
								idx := slices.Index(pastAllOwner.Choices, choice_)
								pastAllOwner.Choices = slices.Delete(pastAllOwner.Choices, idx, idx+1)
								newAllOwner.Choices = append(newAllOwner.Choices, choice_)
							}
						} else {
							newAllOwner.Choices = append(newAllOwner.Choices, choice_)
						}
					}
				}
			}
		case "Choice:Choices":
			// we need to retrieve the field owner before the change
			var pastChoiceOwner *models.Choice
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Choice"
			rf.Fieldname = "Choices"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				choiceFormCallback.probe.stageOfInterest,
				choiceFormCallback.probe.backRepoOfInterest,
				choice_,
				&rf)

			if reverseFieldOwner != nil {
				pastChoiceOwner = reverseFieldOwner.(*models.Choice)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastChoiceOwner != nil {
					idx := slices.Index(pastChoiceOwner.Choices, choice_)
					pastChoiceOwner.Choices = slices.Delete(pastChoiceOwner.Choices, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _choice := range *models.GetGongstructInstancesSet[models.Choice](choiceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _choice.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newChoiceOwner := _choice // we have a match
						if pastChoiceOwner != nil {
							if newChoiceOwner != pastChoiceOwner {
								idx := slices.Index(pastChoiceOwner.Choices, choice_)
								pastChoiceOwner.Choices = slices.Delete(pastChoiceOwner.Choices, idx, idx+1)
								newChoiceOwner.Choices = append(newChoiceOwner.Choices, choice_)
							}
						} else {
							newChoiceOwner.Choices = append(newChoiceOwner.Choices, choice_)
						}
					}
				}
			}
		case "ComplexType:Choices":
			// we need to retrieve the field owner before the change
			var pastComplexTypeOwner *models.ComplexType
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ComplexType"
			rf.Fieldname = "Choices"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				choiceFormCallback.probe.stageOfInterest,
				choiceFormCallback.probe.backRepoOfInterest,
				choice_,
				&rf)

			if reverseFieldOwner != nil {
				pastComplexTypeOwner = reverseFieldOwner.(*models.ComplexType)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastComplexTypeOwner != nil {
					idx := slices.Index(pastComplexTypeOwner.Choices, choice_)
					pastComplexTypeOwner.Choices = slices.Delete(pastComplexTypeOwner.Choices, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](choiceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _complextype.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newComplexTypeOwner := _complextype // we have a match
						if pastComplexTypeOwner != nil {
							if newComplexTypeOwner != pastComplexTypeOwner {
								idx := slices.Index(pastComplexTypeOwner.Choices, choice_)
								pastComplexTypeOwner.Choices = slices.Delete(pastComplexTypeOwner.Choices, idx, idx+1)
								newComplexTypeOwner.Choices = append(newComplexTypeOwner.Choices, choice_)
							}
						} else {
							newComplexTypeOwner.Choices = append(newComplexTypeOwner.Choices, choice_)
						}
					}
				}
			}
		case "Extension:Choices":
			// we need to retrieve the field owner before the change
			var pastExtensionOwner *models.Extension
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Extension"
			rf.Fieldname = "Choices"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				choiceFormCallback.probe.stageOfInterest,
				choiceFormCallback.probe.backRepoOfInterest,
				choice_,
				&rf)

			if reverseFieldOwner != nil {
				pastExtensionOwner = reverseFieldOwner.(*models.Extension)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastExtensionOwner != nil {
					idx := slices.Index(pastExtensionOwner.Choices, choice_)
					pastExtensionOwner.Choices = slices.Delete(pastExtensionOwner.Choices, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _extension := range *models.GetGongstructInstancesSet[models.Extension](choiceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _extension.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newExtensionOwner := _extension // we have a match
						if pastExtensionOwner != nil {
							if newExtensionOwner != pastExtensionOwner {
								idx := slices.Index(pastExtensionOwner.Choices, choice_)
								pastExtensionOwner.Choices = slices.Delete(pastExtensionOwner.Choices, idx, idx+1)
								newExtensionOwner.Choices = append(newExtensionOwner.Choices, choice_)
							}
						} else {
							newExtensionOwner.Choices = append(newExtensionOwner.Choices, choice_)
						}
					}
				}
			}
		case "Group:Choices":
			// we need to retrieve the field owner before the change
			var pastGroupOwner *models.Group
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Choices"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				choiceFormCallback.probe.stageOfInterest,
				choiceFormCallback.probe.backRepoOfInterest,
				choice_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupOwner = reverseFieldOwner.(*models.Group)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupOwner != nil {
					idx := slices.Index(pastGroupOwner.Choices, choice_)
					pastGroupOwner.Choices = slices.Delete(pastGroupOwner.Choices, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _group := range *models.GetGongstructInstancesSet[models.Group](choiceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupOwner := _group // we have a match
						if pastGroupOwner != nil {
							if newGroupOwner != pastGroupOwner {
								idx := slices.Index(pastGroupOwner.Choices, choice_)
								pastGroupOwner.Choices = slices.Delete(pastGroupOwner.Choices, idx, idx+1)
								newGroupOwner.Choices = append(newGroupOwner.Choices, choice_)
							}
						} else {
							newGroupOwner.Choices = append(newGroupOwner.Choices, choice_)
						}
					}
				}
			}
		case "Sequence:Choices":
			// we need to retrieve the field owner before the change
			var pastSequenceOwner *models.Sequence
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sequence"
			rf.Fieldname = "Choices"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				choiceFormCallback.probe.stageOfInterest,
				choiceFormCallback.probe.backRepoOfInterest,
				choice_,
				&rf)

			if reverseFieldOwner != nil {
				pastSequenceOwner = reverseFieldOwner.(*models.Sequence)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSequenceOwner != nil {
					idx := slices.Index(pastSequenceOwner.Choices, choice_)
					pastSequenceOwner.Choices = slices.Delete(pastSequenceOwner.Choices, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _sequence := range *models.GetGongstructInstancesSet[models.Sequence](choiceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _sequence.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSequenceOwner := _sequence // we have a match
						if pastSequenceOwner != nil {
							if newSequenceOwner != pastSequenceOwner {
								idx := slices.Index(pastSequenceOwner.Choices, choice_)
								pastSequenceOwner.Choices = slices.Delete(pastSequenceOwner.Choices, idx, idx+1)
								newSequenceOwner.Choices = append(newSequenceOwner.Choices, choice_)
							}
						} else {
							newSequenceOwner.Choices = append(newSequenceOwner.Choices, choice_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if choiceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		choice_.Unstage(choiceFormCallback.probe.stageOfInterest)
	}

	choiceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Choice](
		choiceFormCallback.probe,
	)
	choiceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if choiceFormCallback.CreationMode || choiceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		choiceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(choiceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ChoiceFormCallback(
			nil,
			choiceFormCallback.probe,
			newFormGroup,
		)
		choice := new(models.Choice)
		FillUpForm(choice, newFormGroup, choiceFormCallback.probe)
		choiceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(choiceFormCallback.probe)
}
func __gong__New__ComplexContentFormCallback(
	complexcontent *models.ComplexContent,
	probe *Probe,
	formGroup *table.FormGroup,
) (complexcontentFormCallback *ComplexContentFormCallback) {
	complexcontentFormCallback = new(ComplexContentFormCallback)
	complexcontentFormCallback.probe = probe
	complexcontentFormCallback.complexcontent = complexcontent
	complexcontentFormCallback.formGroup = formGroup

	complexcontentFormCallback.CreationMode = (complexcontent == nil)

	return
}

type ComplexContentFormCallback struct {
	complexcontent *models.ComplexContent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (complexcontentFormCallback *ComplexContentFormCallback) OnSave() {

	log.Println("ComplexContentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	complexcontentFormCallback.probe.formStage.Checkout()

	if complexcontentFormCallback.complexcontent == nil {
		complexcontentFormCallback.complexcontent = new(models.ComplexContent).Stage(complexcontentFormCallback.probe.stageOfInterest)
	}
	complexcontent_ := complexcontentFormCallback.complexcontent
	_ = complexcontent_

	for _, formDiv := range complexcontentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(complexcontent_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if complexcontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complexcontent_.Unstage(complexcontentFormCallback.probe.stageOfInterest)
	}

	complexcontentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ComplexContent](
		complexcontentFormCallback.probe,
	)
	complexcontentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if complexcontentFormCallback.CreationMode || complexcontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complexcontentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(complexcontentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ComplexContentFormCallback(
			nil,
			complexcontentFormCallback.probe,
			newFormGroup,
		)
		complexcontent := new(models.ComplexContent)
		FillUpForm(complexcontent, newFormGroup, complexcontentFormCallback.probe)
		complexcontentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(complexcontentFormCallback.probe)
}
func __gong__New__ComplexTypeFormCallback(
	complextype *models.ComplexType,
	probe *Probe,
	formGroup *table.FormGroup,
) (complextypeFormCallback *ComplexTypeFormCallback) {
	complextypeFormCallback = new(ComplexTypeFormCallback)
	complextypeFormCallback.probe = probe
	complextypeFormCallback.complextype = complextype
	complextypeFormCallback.formGroup = formGroup

	complextypeFormCallback.CreationMode = (complextype == nil)

	return
}

type ComplexTypeFormCallback struct {
	complextype *models.ComplexType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (complextypeFormCallback *ComplexTypeFormCallback) OnSave() {

	log.Println("ComplexTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	complextypeFormCallback.probe.formStage.Checkout()

	if complextypeFormCallback.complextype == nil {
		complextypeFormCallback.complextype = new(models.ComplexType).Stage(complextypeFormCallback.probe.stageOfInterest)
	}
	complextype_ := complextypeFormCallback.complextype
	_ = complextype_

	for _, formDiv := range complextypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(complextype_.Name), formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(complextype_.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(complextype_.GoIdentifier), formDiv)
		case "IsAnonymous":
			FormDivBasicFieldToField(&(complextype_.IsAnonymous), formDiv)
		case "OuterElement":
			FormDivSelectFieldToField(&(complextype_.OuterElement), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(complextype_.Annotation), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(complextype_.NameXSD), formDiv)
		case "Extension":
			FormDivSelectFieldToField(&(complextype_.Extension), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "SimpleContent":
			FormDivSelectFieldToField(&(complextype_.SimpleContent), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "ComplexContent":
			FormDivSelectFieldToField(&(complextype_.ComplexContent), complextypeFormCallback.probe.stageOfInterest, formDiv)
		case "Schema:ComplexTypes":
			// we need to retrieve the field owner before the change
			var pastSchemaOwner *models.Schema
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "ComplexTypes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				complextypeFormCallback.probe.stageOfInterest,
				complextypeFormCallback.probe.backRepoOfInterest,
				complextype_,
				&rf)

			if reverseFieldOwner != nil {
				pastSchemaOwner = reverseFieldOwner.(*models.Schema)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSchemaOwner != nil {
					idx := slices.Index(pastSchemaOwner.ComplexTypes, complextype_)
					pastSchemaOwner.ComplexTypes = slices.Delete(pastSchemaOwner.ComplexTypes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _schema := range *models.GetGongstructInstancesSet[models.Schema](complextypeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _schema.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSchemaOwner := _schema // we have a match
						if pastSchemaOwner != nil {
							if newSchemaOwner != pastSchemaOwner {
								idx := slices.Index(pastSchemaOwner.ComplexTypes, complextype_)
								pastSchemaOwner.ComplexTypes = slices.Delete(pastSchemaOwner.ComplexTypes, idx, idx+1)
								newSchemaOwner.ComplexTypes = append(newSchemaOwner.ComplexTypes, complextype_)
							}
						} else {
							newSchemaOwner.ComplexTypes = append(newSchemaOwner.ComplexTypes, complextype_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if complextypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complextype_.Unstage(complextypeFormCallback.probe.stageOfInterest)
	}

	complextypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ComplexType](
		complextypeFormCallback.probe,
	)
	complextypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if complextypeFormCallback.CreationMode || complextypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complextypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(complextypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ComplexTypeFormCallback(
			nil,
			complextypeFormCallback.probe,
			newFormGroup,
		)
		complextype := new(models.ComplexType)
		FillUpForm(complextype, newFormGroup, complextypeFormCallback.probe)
		complextypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(complextypeFormCallback.probe)
}
func __gong__New__DocumentationFormCallback(
	documentation *models.Documentation,
	probe *Probe,
	formGroup *table.FormGroup,
) (documentationFormCallback *DocumentationFormCallback) {
	documentationFormCallback = new(DocumentationFormCallback)
	documentationFormCallback.probe = probe
	documentationFormCallback.documentation = documentation
	documentationFormCallback.formGroup = formGroup

	documentationFormCallback.CreationMode = (documentation == nil)

	return
}

type DocumentationFormCallback struct {
	documentation *models.Documentation

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (documentationFormCallback *DocumentationFormCallback) OnSave() {

	log.Println("DocumentationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	documentationFormCallback.probe.formStage.Checkout()

	if documentationFormCallback.documentation == nil {
		documentationFormCallback.documentation = new(models.Documentation).Stage(documentationFormCallback.probe.stageOfInterest)
	}
	documentation_ := documentationFormCallback.documentation
	_ = documentation_

	for _, formDiv := range documentationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(documentation_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(documentation_.Text), formDiv)
		case "Source":
			FormDivBasicFieldToField(&(documentation_.Source), formDiv)
		case "Lang":
			FormDivBasicFieldToField(&(documentation_.Lang), formDiv)
		case "Annotation:Documentations":
			// we need to retrieve the field owner before the change
			var pastAnnotationOwner *models.Annotation
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Annotation"
			rf.Fieldname = "Documentations"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				documentationFormCallback.probe.stageOfInterest,
				documentationFormCallback.probe.backRepoOfInterest,
				documentation_,
				&rf)

			if reverseFieldOwner != nil {
				pastAnnotationOwner = reverseFieldOwner.(*models.Annotation)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAnnotationOwner != nil {
					idx := slices.Index(pastAnnotationOwner.Documentations, documentation_)
					pastAnnotationOwner.Documentations = slices.Delete(pastAnnotationOwner.Documentations, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _annotation := range *models.GetGongstructInstancesSet[models.Annotation](documentationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _annotation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAnnotationOwner := _annotation // we have a match
						if pastAnnotationOwner != nil {
							if newAnnotationOwner != pastAnnotationOwner {
								idx := slices.Index(pastAnnotationOwner.Documentations, documentation_)
								pastAnnotationOwner.Documentations = slices.Delete(pastAnnotationOwner.Documentations, idx, idx+1)
								newAnnotationOwner.Documentations = append(newAnnotationOwner.Documentations, documentation_)
							}
						} else {
							newAnnotationOwner.Documentations = append(newAnnotationOwner.Documentations, documentation_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if documentationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentation_.Unstage(documentationFormCallback.probe.stageOfInterest)
	}

	documentationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Documentation](
		documentationFormCallback.probe,
	)
	documentationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if documentationFormCallback.CreationMode || documentationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		documentationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(documentationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DocumentationFormCallback(
			nil,
			documentationFormCallback.probe,
			newFormGroup,
		)
		documentation := new(models.Documentation)
		FillUpForm(documentation, newFormGroup, documentationFormCallback.probe)
		documentationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(documentationFormCallback.probe)
}
func __gong__New__ElementFormCallback(
	element *models.Element,
	probe *Probe,
	formGroup *table.FormGroup,
) (elementFormCallback *ElementFormCallback) {
	elementFormCallback = new(ElementFormCallback)
	elementFormCallback.probe = probe
	elementFormCallback.element = element
	elementFormCallback.formGroup = formGroup

	elementFormCallback.CreationMode = (element == nil)

	return
}

type ElementFormCallback struct {
	element *models.Element

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (elementFormCallback *ElementFormCallback) OnSave() {

	log.Println("ElementFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	elementFormCallback.probe.formStage.Checkout()

	if elementFormCallback.element == nil {
		elementFormCallback.element = new(models.Element).Stage(elementFormCallback.probe.stageOfInterest)
	}
	element_ := elementFormCallback.element
	_ = element_

	for _, formDiv := range elementFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(element_.Name), formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(element_.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(element_.GoIdentifier), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(element_.Annotation), elementFormCallback.probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(element_.NameXSD), formDiv)
		case "Type":
			FormDivBasicFieldToField(&(element_.Type), formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(element_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(element_.MaxOccurs), formDiv)
		case "Default":
			FormDivBasicFieldToField(&(element_.Default), formDiv)
		case "Fixed":
			FormDivBasicFieldToField(&(element_.Fixed), formDiv)
		case "Nillable":
			FormDivBasicFieldToField(&(element_.Nillable), formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(element_.Ref), formDiv)
		case "Abstract":
			FormDivBasicFieldToField(&(element_.Abstract), formDiv)
		case "Form":
			FormDivBasicFieldToField(&(element_.Form), formDiv)
		case "Block":
			FormDivBasicFieldToField(&(element_.Block), formDiv)
		case "Final":
			FormDivBasicFieldToField(&(element_.Final), formDiv)
		case "SimpleType":
			FormDivSelectFieldToField(&(element_.SimpleType), elementFormCallback.probe.stageOfInterest, formDiv)
		case "ComplexType":
			FormDivSelectFieldToField(&(element_.ComplexType), elementFormCallback.probe.stageOfInterest, formDiv)
		case "All:Elements":
			// we need to retrieve the field owner before the change
			var pastAllOwner *models.All
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "All"
			rf.Fieldname = "Elements"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				elementFormCallback.probe.stageOfInterest,
				elementFormCallback.probe.backRepoOfInterest,
				element_,
				&rf)

			if reverseFieldOwner != nil {
				pastAllOwner = reverseFieldOwner.(*models.All)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAllOwner != nil {
					idx := slices.Index(pastAllOwner.Elements, element_)
					pastAllOwner.Elements = slices.Delete(pastAllOwner.Elements, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _all := range *models.GetGongstructInstancesSet[models.All](elementFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _all.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAllOwner := _all // we have a match
						if pastAllOwner != nil {
							if newAllOwner != pastAllOwner {
								idx := slices.Index(pastAllOwner.Elements, element_)
								pastAllOwner.Elements = slices.Delete(pastAllOwner.Elements, idx, idx+1)
								newAllOwner.Elements = append(newAllOwner.Elements, element_)
							}
						} else {
							newAllOwner.Elements = append(newAllOwner.Elements, element_)
						}
					}
				}
			}
		case "Choice:Elements":
			// we need to retrieve the field owner before the change
			var pastChoiceOwner *models.Choice
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Choice"
			rf.Fieldname = "Elements"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				elementFormCallback.probe.stageOfInterest,
				elementFormCallback.probe.backRepoOfInterest,
				element_,
				&rf)

			if reverseFieldOwner != nil {
				pastChoiceOwner = reverseFieldOwner.(*models.Choice)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastChoiceOwner != nil {
					idx := slices.Index(pastChoiceOwner.Elements, element_)
					pastChoiceOwner.Elements = slices.Delete(pastChoiceOwner.Elements, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _choice := range *models.GetGongstructInstancesSet[models.Choice](elementFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _choice.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newChoiceOwner := _choice // we have a match
						if pastChoiceOwner != nil {
							if newChoiceOwner != pastChoiceOwner {
								idx := slices.Index(pastChoiceOwner.Elements, element_)
								pastChoiceOwner.Elements = slices.Delete(pastChoiceOwner.Elements, idx, idx+1)
								newChoiceOwner.Elements = append(newChoiceOwner.Elements, element_)
							}
						} else {
							newChoiceOwner.Elements = append(newChoiceOwner.Elements, element_)
						}
					}
				}
			}
		case "ComplexType:Elements":
			// we need to retrieve the field owner before the change
			var pastComplexTypeOwner *models.ComplexType
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ComplexType"
			rf.Fieldname = "Elements"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				elementFormCallback.probe.stageOfInterest,
				elementFormCallback.probe.backRepoOfInterest,
				element_,
				&rf)

			if reverseFieldOwner != nil {
				pastComplexTypeOwner = reverseFieldOwner.(*models.ComplexType)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastComplexTypeOwner != nil {
					idx := slices.Index(pastComplexTypeOwner.Elements, element_)
					pastComplexTypeOwner.Elements = slices.Delete(pastComplexTypeOwner.Elements, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](elementFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _complextype.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newComplexTypeOwner := _complextype // we have a match
						if pastComplexTypeOwner != nil {
							if newComplexTypeOwner != pastComplexTypeOwner {
								idx := slices.Index(pastComplexTypeOwner.Elements, element_)
								pastComplexTypeOwner.Elements = slices.Delete(pastComplexTypeOwner.Elements, idx, idx+1)
								newComplexTypeOwner.Elements = append(newComplexTypeOwner.Elements, element_)
							}
						} else {
							newComplexTypeOwner.Elements = append(newComplexTypeOwner.Elements, element_)
						}
					}
				}
			}
		case "Extension:Elements":
			// we need to retrieve the field owner before the change
			var pastExtensionOwner *models.Extension
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Extension"
			rf.Fieldname = "Elements"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				elementFormCallback.probe.stageOfInterest,
				elementFormCallback.probe.backRepoOfInterest,
				element_,
				&rf)

			if reverseFieldOwner != nil {
				pastExtensionOwner = reverseFieldOwner.(*models.Extension)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastExtensionOwner != nil {
					idx := slices.Index(pastExtensionOwner.Elements, element_)
					pastExtensionOwner.Elements = slices.Delete(pastExtensionOwner.Elements, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _extension := range *models.GetGongstructInstancesSet[models.Extension](elementFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _extension.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newExtensionOwner := _extension // we have a match
						if pastExtensionOwner != nil {
							if newExtensionOwner != pastExtensionOwner {
								idx := slices.Index(pastExtensionOwner.Elements, element_)
								pastExtensionOwner.Elements = slices.Delete(pastExtensionOwner.Elements, idx, idx+1)
								newExtensionOwner.Elements = append(newExtensionOwner.Elements, element_)
							}
						} else {
							newExtensionOwner.Elements = append(newExtensionOwner.Elements, element_)
						}
					}
				}
			}
		case "Group:Elements":
			// we need to retrieve the field owner before the change
			var pastGroupOwner *models.Group
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Elements"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				elementFormCallback.probe.stageOfInterest,
				elementFormCallback.probe.backRepoOfInterest,
				element_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupOwner = reverseFieldOwner.(*models.Group)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupOwner != nil {
					idx := slices.Index(pastGroupOwner.Elements, element_)
					pastGroupOwner.Elements = slices.Delete(pastGroupOwner.Elements, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _group := range *models.GetGongstructInstancesSet[models.Group](elementFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupOwner := _group // we have a match
						if pastGroupOwner != nil {
							if newGroupOwner != pastGroupOwner {
								idx := slices.Index(pastGroupOwner.Elements, element_)
								pastGroupOwner.Elements = slices.Delete(pastGroupOwner.Elements, idx, idx+1)
								newGroupOwner.Elements = append(newGroupOwner.Elements, element_)
							}
						} else {
							newGroupOwner.Elements = append(newGroupOwner.Elements, element_)
						}
					}
				}
			}
		case "Schema:Elements":
			// we need to retrieve the field owner before the change
			var pastSchemaOwner *models.Schema
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "Elements"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				elementFormCallback.probe.stageOfInterest,
				elementFormCallback.probe.backRepoOfInterest,
				element_,
				&rf)

			if reverseFieldOwner != nil {
				pastSchemaOwner = reverseFieldOwner.(*models.Schema)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSchemaOwner != nil {
					idx := slices.Index(pastSchemaOwner.Elements, element_)
					pastSchemaOwner.Elements = slices.Delete(pastSchemaOwner.Elements, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _schema := range *models.GetGongstructInstancesSet[models.Schema](elementFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _schema.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSchemaOwner := _schema // we have a match
						if pastSchemaOwner != nil {
							if newSchemaOwner != pastSchemaOwner {
								idx := slices.Index(pastSchemaOwner.Elements, element_)
								pastSchemaOwner.Elements = slices.Delete(pastSchemaOwner.Elements, idx, idx+1)
								newSchemaOwner.Elements = append(newSchemaOwner.Elements, element_)
							}
						} else {
							newSchemaOwner.Elements = append(newSchemaOwner.Elements, element_)
						}
					}
				}
			}
		case "Sequence:Elements":
			// we need to retrieve the field owner before the change
			var pastSequenceOwner *models.Sequence
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sequence"
			rf.Fieldname = "Elements"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				elementFormCallback.probe.stageOfInterest,
				elementFormCallback.probe.backRepoOfInterest,
				element_,
				&rf)

			if reverseFieldOwner != nil {
				pastSequenceOwner = reverseFieldOwner.(*models.Sequence)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSequenceOwner != nil {
					idx := slices.Index(pastSequenceOwner.Elements, element_)
					pastSequenceOwner.Elements = slices.Delete(pastSequenceOwner.Elements, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _sequence := range *models.GetGongstructInstancesSet[models.Sequence](elementFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _sequence.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSequenceOwner := _sequence // we have a match
						if pastSequenceOwner != nil {
							if newSequenceOwner != pastSequenceOwner {
								idx := slices.Index(pastSequenceOwner.Elements, element_)
								pastSequenceOwner.Elements = slices.Delete(pastSequenceOwner.Elements, idx, idx+1)
								newSequenceOwner.Elements = append(newSequenceOwner.Elements, element_)
							}
						} else {
							newSequenceOwner.Elements = append(newSequenceOwner.Elements, element_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if elementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		element_.Unstage(elementFormCallback.probe.stageOfInterest)
	}

	elementFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Element](
		elementFormCallback.probe,
	)
	elementFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if elementFormCallback.CreationMode || elementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		elementFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(elementFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ElementFormCallback(
			nil,
			elementFormCallback.probe,
			newFormGroup,
		)
		element := new(models.Element)
		FillUpForm(element, newFormGroup, elementFormCallback.probe)
		elementFormCallback.probe.formStage.Commit()
	}

	fillUpTree(elementFormCallback.probe)
}
func __gong__New__EnumerationFormCallback(
	enumeration *models.Enumeration,
	probe *Probe,
	formGroup *table.FormGroup,
) (enumerationFormCallback *EnumerationFormCallback) {
	enumerationFormCallback = new(EnumerationFormCallback)
	enumerationFormCallback.probe = probe
	enumerationFormCallback.enumeration = enumeration
	enumerationFormCallback.formGroup = formGroup

	enumerationFormCallback.CreationMode = (enumeration == nil)

	return
}

type EnumerationFormCallback struct {
	enumeration *models.Enumeration

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (enumerationFormCallback *EnumerationFormCallback) OnSave() {

	log.Println("EnumerationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	enumerationFormCallback.probe.formStage.Checkout()

	if enumerationFormCallback.enumeration == nil {
		enumerationFormCallback.enumeration = new(models.Enumeration).Stage(enumerationFormCallback.probe.stageOfInterest)
	}
	enumeration_ := enumerationFormCallback.enumeration
	_ = enumeration_

	for _, formDiv := range enumerationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(enumeration_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(enumeration_.Annotation), enumerationFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(enumeration_.Value), formDiv)
		case "Restriction:Enumerations":
			// we need to retrieve the field owner before the change
			var pastRestrictionOwner *models.Restriction
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Restriction"
			rf.Fieldname = "Enumerations"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				enumerationFormCallback.probe.stageOfInterest,
				enumerationFormCallback.probe.backRepoOfInterest,
				enumeration_,
				&rf)

			if reverseFieldOwner != nil {
				pastRestrictionOwner = reverseFieldOwner.(*models.Restriction)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRestrictionOwner != nil {
					idx := slices.Index(pastRestrictionOwner.Enumerations, enumeration_)
					pastRestrictionOwner.Enumerations = slices.Delete(pastRestrictionOwner.Enumerations, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _restriction := range *models.GetGongstructInstancesSet[models.Restriction](enumerationFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _restriction.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRestrictionOwner := _restriction // we have a match
						if pastRestrictionOwner != nil {
							if newRestrictionOwner != pastRestrictionOwner {
								idx := slices.Index(pastRestrictionOwner.Enumerations, enumeration_)
								pastRestrictionOwner.Enumerations = slices.Delete(pastRestrictionOwner.Enumerations, idx, idx+1)
								newRestrictionOwner.Enumerations = append(newRestrictionOwner.Enumerations, enumeration_)
							}
						} else {
							newRestrictionOwner.Enumerations = append(newRestrictionOwner.Enumerations, enumeration_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enumeration_.Unstage(enumerationFormCallback.probe.stageOfInterest)
	}

	enumerationFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Enumeration](
		enumerationFormCallback.probe,
	)
	enumerationFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if enumerationFormCallback.CreationMode || enumerationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		enumerationFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(enumerationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EnumerationFormCallback(
			nil,
			enumerationFormCallback.probe,
			newFormGroup,
		)
		enumeration := new(models.Enumeration)
		FillUpForm(enumeration, newFormGroup, enumerationFormCallback.probe)
		enumerationFormCallback.probe.formStage.Commit()
	}

	fillUpTree(enumerationFormCallback.probe)
}
func __gong__New__ExtensionFormCallback(
	extension *models.Extension,
	probe *Probe,
	formGroup *table.FormGroup,
) (extensionFormCallback *ExtensionFormCallback) {
	extensionFormCallback = new(ExtensionFormCallback)
	extensionFormCallback.probe = probe
	extensionFormCallback.extension = extension
	extensionFormCallback.formGroup = formGroup

	extensionFormCallback.CreationMode = (extension == nil)

	return
}

type ExtensionFormCallback struct {
	extension *models.Extension

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (extensionFormCallback *ExtensionFormCallback) OnSave() {

	log.Println("ExtensionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	extensionFormCallback.probe.formStage.Checkout()

	if extensionFormCallback.extension == nil {
		extensionFormCallback.extension = new(models.Extension).Stage(extensionFormCallback.probe.stageOfInterest)
	}
	extension_ := extensionFormCallback.extension
	_ = extension_

	for _, formDiv := range extensionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(extension_.Name), formDiv)
		case "Base":
			FormDivBasicFieldToField(&(extension_.Base), formDiv)
		}
	}

	// manage the suppress operation
	if extensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		extension_.Unstage(extensionFormCallback.probe.stageOfInterest)
	}

	extensionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Extension](
		extensionFormCallback.probe,
	)
	extensionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if extensionFormCallback.CreationMode || extensionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		extensionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(extensionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ExtensionFormCallback(
			nil,
			extensionFormCallback.probe,
			newFormGroup,
		)
		extension := new(models.Extension)
		FillUpForm(extension, newFormGroup, extensionFormCallback.probe)
		extensionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(extensionFormCallback.probe)
}
func __gong__New__GroupFormCallback(
	group *models.Group,
	probe *Probe,
	formGroup *table.FormGroup,
) (groupFormCallback *GroupFormCallback) {
	groupFormCallback = new(GroupFormCallback)
	groupFormCallback.probe = probe
	groupFormCallback.group = group
	groupFormCallback.formGroup = formGroup

	groupFormCallback.CreationMode = (group == nil)

	return
}

type GroupFormCallback struct {
	group *models.Group

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (groupFormCallback *GroupFormCallback) OnSave() {

	log.Println("GroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	groupFormCallback.probe.formStage.Checkout()

	if groupFormCallback.group == nil {
		groupFormCallback.group = new(models.Group).Stage(groupFormCallback.probe.stageOfInterest)
	}
	group_ := groupFormCallback.group
	_ = group_

	for _, formDiv := range groupFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(group_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(group_.Annotation), groupFormCallback.probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(group_.NameXSD), formDiv)
		case "Ref":
			FormDivBasicFieldToField(&(group_.Ref), formDiv)
		case "IsAnonymous":
			FormDivBasicFieldToField(&(group_.IsAnonymous), formDiv)
		case "OuterElement":
			FormDivSelectFieldToField(&(group_.OuterElement), groupFormCallback.probe.stageOfInterest, formDiv)
		case "HasNameConflict":
			FormDivBasicFieldToField(&(group_.HasNameConflict), formDiv)
		case "GoIdentifier":
			FormDivBasicFieldToField(&(group_.GoIdentifier), formDiv)
		case "All:Groups":
			// we need to retrieve the field owner before the change
			var pastAllOwner *models.All
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "All"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupFormCallback.probe.stageOfInterest,
				groupFormCallback.probe.backRepoOfInterest,
				group_,
				&rf)

			if reverseFieldOwner != nil {
				pastAllOwner = reverseFieldOwner.(*models.All)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAllOwner != nil {
					idx := slices.Index(pastAllOwner.Groups, group_)
					pastAllOwner.Groups = slices.Delete(pastAllOwner.Groups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _all := range *models.GetGongstructInstancesSet[models.All](groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _all.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAllOwner := _all // we have a match
						if pastAllOwner != nil {
							if newAllOwner != pastAllOwner {
								idx := slices.Index(pastAllOwner.Groups, group_)
								pastAllOwner.Groups = slices.Delete(pastAllOwner.Groups, idx, idx+1)
								newAllOwner.Groups = append(newAllOwner.Groups, group_)
							}
						} else {
							newAllOwner.Groups = append(newAllOwner.Groups, group_)
						}
					}
				}
			}
		case "Choice:Groups":
			// we need to retrieve the field owner before the change
			var pastChoiceOwner *models.Choice
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Choice"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupFormCallback.probe.stageOfInterest,
				groupFormCallback.probe.backRepoOfInterest,
				group_,
				&rf)

			if reverseFieldOwner != nil {
				pastChoiceOwner = reverseFieldOwner.(*models.Choice)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastChoiceOwner != nil {
					idx := slices.Index(pastChoiceOwner.Groups, group_)
					pastChoiceOwner.Groups = slices.Delete(pastChoiceOwner.Groups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _choice := range *models.GetGongstructInstancesSet[models.Choice](groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _choice.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newChoiceOwner := _choice // we have a match
						if pastChoiceOwner != nil {
							if newChoiceOwner != pastChoiceOwner {
								idx := slices.Index(pastChoiceOwner.Groups, group_)
								pastChoiceOwner.Groups = slices.Delete(pastChoiceOwner.Groups, idx, idx+1)
								newChoiceOwner.Groups = append(newChoiceOwner.Groups, group_)
							}
						} else {
							newChoiceOwner.Groups = append(newChoiceOwner.Groups, group_)
						}
					}
				}
			}
		case "ComplexType:Groups":
			// we need to retrieve the field owner before the change
			var pastComplexTypeOwner *models.ComplexType
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ComplexType"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupFormCallback.probe.stageOfInterest,
				groupFormCallback.probe.backRepoOfInterest,
				group_,
				&rf)

			if reverseFieldOwner != nil {
				pastComplexTypeOwner = reverseFieldOwner.(*models.ComplexType)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastComplexTypeOwner != nil {
					idx := slices.Index(pastComplexTypeOwner.Groups, group_)
					pastComplexTypeOwner.Groups = slices.Delete(pastComplexTypeOwner.Groups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _complextype.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newComplexTypeOwner := _complextype // we have a match
						if pastComplexTypeOwner != nil {
							if newComplexTypeOwner != pastComplexTypeOwner {
								idx := slices.Index(pastComplexTypeOwner.Groups, group_)
								pastComplexTypeOwner.Groups = slices.Delete(pastComplexTypeOwner.Groups, idx, idx+1)
								newComplexTypeOwner.Groups = append(newComplexTypeOwner.Groups, group_)
							}
						} else {
							newComplexTypeOwner.Groups = append(newComplexTypeOwner.Groups, group_)
						}
					}
				}
			}
		case "Element:Groups":
			// we need to retrieve the field owner before the change
			var pastElementOwner *models.Element
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Element"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupFormCallback.probe.stageOfInterest,
				groupFormCallback.probe.backRepoOfInterest,
				group_,
				&rf)

			if reverseFieldOwner != nil {
				pastElementOwner = reverseFieldOwner.(*models.Element)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastElementOwner != nil {
					idx := slices.Index(pastElementOwner.Groups, group_)
					pastElementOwner.Groups = slices.Delete(pastElementOwner.Groups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _element := range *models.GetGongstructInstancesSet[models.Element](groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _element.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newElementOwner := _element // we have a match
						if pastElementOwner != nil {
							if newElementOwner != pastElementOwner {
								idx := slices.Index(pastElementOwner.Groups, group_)
								pastElementOwner.Groups = slices.Delete(pastElementOwner.Groups, idx, idx+1)
								newElementOwner.Groups = append(newElementOwner.Groups, group_)
							}
						} else {
							newElementOwner.Groups = append(newElementOwner.Groups, group_)
						}
					}
				}
			}
		case "Extension:Groups":
			// we need to retrieve the field owner before the change
			var pastExtensionOwner *models.Extension
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Extension"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupFormCallback.probe.stageOfInterest,
				groupFormCallback.probe.backRepoOfInterest,
				group_,
				&rf)

			if reverseFieldOwner != nil {
				pastExtensionOwner = reverseFieldOwner.(*models.Extension)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastExtensionOwner != nil {
					idx := slices.Index(pastExtensionOwner.Groups, group_)
					pastExtensionOwner.Groups = slices.Delete(pastExtensionOwner.Groups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _extension := range *models.GetGongstructInstancesSet[models.Extension](groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _extension.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newExtensionOwner := _extension // we have a match
						if pastExtensionOwner != nil {
							if newExtensionOwner != pastExtensionOwner {
								idx := slices.Index(pastExtensionOwner.Groups, group_)
								pastExtensionOwner.Groups = slices.Delete(pastExtensionOwner.Groups, idx, idx+1)
								newExtensionOwner.Groups = append(newExtensionOwner.Groups, group_)
							}
						} else {
							newExtensionOwner.Groups = append(newExtensionOwner.Groups, group_)
						}
					}
				}
			}
		case "Group:Groups":
			// we need to retrieve the field owner before the change
			var pastGroupOwner *models.Group
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupFormCallback.probe.stageOfInterest,
				groupFormCallback.probe.backRepoOfInterest,
				group_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupOwner = reverseFieldOwner.(*models.Group)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupOwner != nil {
					idx := slices.Index(pastGroupOwner.Groups, group_)
					pastGroupOwner.Groups = slices.Delete(pastGroupOwner.Groups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _group := range *models.GetGongstructInstancesSet[models.Group](groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupOwner := _group // we have a match
						if pastGroupOwner != nil {
							if newGroupOwner != pastGroupOwner {
								idx := slices.Index(pastGroupOwner.Groups, group_)
								pastGroupOwner.Groups = slices.Delete(pastGroupOwner.Groups, idx, idx+1)
								newGroupOwner.Groups = append(newGroupOwner.Groups, group_)
							}
						} else {
							newGroupOwner.Groups = append(newGroupOwner.Groups, group_)
						}
					}
				}
			}
		case "Schema:Groups":
			// we need to retrieve the field owner before the change
			var pastSchemaOwner *models.Schema
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupFormCallback.probe.stageOfInterest,
				groupFormCallback.probe.backRepoOfInterest,
				group_,
				&rf)

			if reverseFieldOwner != nil {
				pastSchemaOwner = reverseFieldOwner.(*models.Schema)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSchemaOwner != nil {
					idx := slices.Index(pastSchemaOwner.Groups, group_)
					pastSchemaOwner.Groups = slices.Delete(pastSchemaOwner.Groups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _schema := range *models.GetGongstructInstancesSet[models.Schema](groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _schema.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSchemaOwner := _schema // we have a match
						if pastSchemaOwner != nil {
							if newSchemaOwner != pastSchemaOwner {
								idx := slices.Index(pastSchemaOwner.Groups, group_)
								pastSchemaOwner.Groups = slices.Delete(pastSchemaOwner.Groups, idx, idx+1)
								newSchemaOwner.Groups = append(newSchemaOwner.Groups, group_)
							}
						} else {
							newSchemaOwner.Groups = append(newSchemaOwner.Groups, group_)
						}
					}
				}
			}
		case "Sequence:Groups":
			// we need to retrieve the field owner before the change
			var pastSequenceOwner *models.Sequence
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sequence"
			rf.Fieldname = "Groups"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				groupFormCallback.probe.stageOfInterest,
				groupFormCallback.probe.backRepoOfInterest,
				group_,
				&rf)

			if reverseFieldOwner != nil {
				pastSequenceOwner = reverseFieldOwner.(*models.Sequence)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSequenceOwner != nil {
					idx := slices.Index(pastSequenceOwner.Groups, group_)
					pastSequenceOwner.Groups = slices.Delete(pastSequenceOwner.Groups, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _sequence := range *models.GetGongstructInstancesSet[models.Sequence](groupFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _sequence.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSequenceOwner := _sequence // we have a match
						if pastSequenceOwner != nil {
							if newSequenceOwner != pastSequenceOwner {
								idx := slices.Index(pastSequenceOwner.Groups, group_)
								pastSequenceOwner.Groups = slices.Delete(pastSequenceOwner.Groups, idx, idx+1)
								newSequenceOwner.Groups = append(newSequenceOwner.Groups, group_)
							}
						} else {
							newSequenceOwner.Groups = append(newSequenceOwner.Groups, group_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		group_.Unstage(groupFormCallback.probe.stageOfInterest)
	}

	groupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Group](
		groupFormCallback.probe,
	)
	groupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if groupFormCallback.CreationMode || groupFormCallback.formGroup.HasSuppressButtonBeenPressed {
		groupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(groupFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GroupFormCallback(
			nil,
			groupFormCallback.probe,
			newFormGroup,
		)
		group := new(models.Group)
		FillUpForm(group, newFormGroup, groupFormCallback.probe)
		groupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(groupFormCallback.probe)
}
func __gong__New__LengthFormCallback(
	length *models.Length,
	probe *Probe,
	formGroup *table.FormGroup,
) (lengthFormCallback *LengthFormCallback) {
	lengthFormCallback = new(LengthFormCallback)
	lengthFormCallback.probe = probe
	lengthFormCallback.length = length
	lengthFormCallback.formGroup = formGroup

	lengthFormCallback.CreationMode = (length == nil)

	return
}

type LengthFormCallback struct {
	length *models.Length

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (lengthFormCallback *LengthFormCallback) OnSave() {

	log.Println("LengthFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lengthFormCallback.probe.formStage.Checkout()

	if lengthFormCallback.length == nil {
		lengthFormCallback.length = new(models.Length).Stage(lengthFormCallback.probe.stageOfInterest)
	}
	length_ := lengthFormCallback.length
	_ = length_

	for _, formDiv := range lengthFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(length_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(length_.Annotation), lengthFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(length_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if lengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		length_.Unstage(lengthFormCallback.probe.stageOfInterest)
	}

	lengthFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Length](
		lengthFormCallback.probe,
	)
	lengthFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if lengthFormCallback.CreationMode || lengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		lengthFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(lengthFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LengthFormCallback(
			nil,
			lengthFormCallback.probe,
			newFormGroup,
		)
		length := new(models.Length)
		FillUpForm(length, newFormGroup, lengthFormCallback.probe)
		lengthFormCallback.probe.formStage.Commit()
	}

	fillUpTree(lengthFormCallback.probe)
}
func __gong__New__MaxInclusiveFormCallback(
	maxinclusive *models.MaxInclusive,
	probe *Probe,
	formGroup *table.FormGroup,
) (maxinclusiveFormCallback *MaxInclusiveFormCallback) {
	maxinclusiveFormCallback = new(MaxInclusiveFormCallback)
	maxinclusiveFormCallback.probe = probe
	maxinclusiveFormCallback.maxinclusive = maxinclusive
	maxinclusiveFormCallback.formGroup = formGroup

	maxinclusiveFormCallback.CreationMode = (maxinclusive == nil)

	return
}

type MaxInclusiveFormCallback struct {
	maxinclusive *models.MaxInclusive

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (maxinclusiveFormCallback *MaxInclusiveFormCallback) OnSave() {

	log.Println("MaxInclusiveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	maxinclusiveFormCallback.probe.formStage.Checkout()

	if maxinclusiveFormCallback.maxinclusive == nil {
		maxinclusiveFormCallback.maxinclusive = new(models.MaxInclusive).Stage(maxinclusiveFormCallback.probe.stageOfInterest)
	}
	maxinclusive_ := maxinclusiveFormCallback.maxinclusive
	_ = maxinclusive_

	for _, formDiv := range maxinclusiveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(maxinclusive_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(maxinclusive_.Annotation), maxinclusiveFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(maxinclusive_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if maxinclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		maxinclusive_.Unstage(maxinclusiveFormCallback.probe.stageOfInterest)
	}

	maxinclusiveFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MaxInclusive](
		maxinclusiveFormCallback.probe,
	)
	maxinclusiveFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if maxinclusiveFormCallback.CreationMode || maxinclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		maxinclusiveFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(maxinclusiveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MaxInclusiveFormCallback(
			nil,
			maxinclusiveFormCallback.probe,
			newFormGroup,
		)
		maxinclusive := new(models.MaxInclusive)
		FillUpForm(maxinclusive, newFormGroup, maxinclusiveFormCallback.probe)
		maxinclusiveFormCallback.probe.formStage.Commit()
	}

	fillUpTree(maxinclusiveFormCallback.probe)
}
func __gong__New__MaxLengthFormCallback(
	maxlength *models.MaxLength,
	probe *Probe,
	formGroup *table.FormGroup,
) (maxlengthFormCallback *MaxLengthFormCallback) {
	maxlengthFormCallback = new(MaxLengthFormCallback)
	maxlengthFormCallback.probe = probe
	maxlengthFormCallback.maxlength = maxlength
	maxlengthFormCallback.formGroup = formGroup

	maxlengthFormCallback.CreationMode = (maxlength == nil)

	return
}

type MaxLengthFormCallback struct {
	maxlength *models.MaxLength

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (maxlengthFormCallback *MaxLengthFormCallback) OnSave() {

	log.Println("MaxLengthFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	maxlengthFormCallback.probe.formStage.Checkout()

	if maxlengthFormCallback.maxlength == nil {
		maxlengthFormCallback.maxlength = new(models.MaxLength).Stage(maxlengthFormCallback.probe.stageOfInterest)
	}
	maxlength_ := maxlengthFormCallback.maxlength
	_ = maxlength_

	for _, formDiv := range maxlengthFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(maxlength_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(maxlength_.Annotation), maxlengthFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(maxlength_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if maxlengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		maxlength_.Unstage(maxlengthFormCallback.probe.stageOfInterest)
	}

	maxlengthFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MaxLength](
		maxlengthFormCallback.probe,
	)
	maxlengthFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if maxlengthFormCallback.CreationMode || maxlengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		maxlengthFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(maxlengthFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MaxLengthFormCallback(
			nil,
			maxlengthFormCallback.probe,
			newFormGroup,
		)
		maxlength := new(models.MaxLength)
		FillUpForm(maxlength, newFormGroup, maxlengthFormCallback.probe)
		maxlengthFormCallback.probe.formStage.Commit()
	}

	fillUpTree(maxlengthFormCallback.probe)
}
func __gong__New__MinInclusiveFormCallback(
	mininclusive *models.MinInclusive,
	probe *Probe,
	formGroup *table.FormGroup,
) (mininclusiveFormCallback *MinInclusiveFormCallback) {
	mininclusiveFormCallback = new(MinInclusiveFormCallback)
	mininclusiveFormCallback.probe = probe
	mininclusiveFormCallback.mininclusive = mininclusive
	mininclusiveFormCallback.formGroup = formGroup

	mininclusiveFormCallback.CreationMode = (mininclusive == nil)

	return
}

type MinInclusiveFormCallback struct {
	mininclusive *models.MinInclusive

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (mininclusiveFormCallback *MinInclusiveFormCallback) OnSave() {

	log.Println("MinInclusiveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	mininclusiveFormCallback.probe.formStage.Checkout()

	if mininclusiveFormCallback.mininclusive == nil {
		mininclusiveFormCallback.mininclusive = new(models.MinInclusive).Stage(mininclusiveFormCallback.probe.stageOfInterest)
	}
	mininclusive_ := mininclusiveFormCallback.mininclusive
	_ = mininclusive_

	for _, formDiv := range mininclusiveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(mininclusive_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(mininclusive_.Annotation), mininclusiveFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(mininclusive_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if mininclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mininclusive_.Unstage(mininclusiveFormCallback.probe.stageOfInterest)
	}

	mininclusiveFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MinInclusive](
		mininclusiveFormCallback.probe,
	)
	mininclusiveFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if mininclusiveFormCallback.CreationMode || mininclusiveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		mininclusiveFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(mininclusiveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MinInclusiveFormCallback(
			nil,
			mininclusiveFormCallback.probe,
			newFormGroup,
		)
		mininclusive := new(models.MinInclusive)
		FillUpForm(mininclusive, newFormGroup, mininclusiveFormCallback.probe)
		mininclusiveFormCallback.probe.formStage.Commit()
	}

	fillUpTree(mininclusiveFormCallback.probe)
}
func __gong__New__MinLengthFormCallback(
	minlength *models.MinLength,
	probe *Probe,
	formGroup *table.FormGroup,
) (minlengthFormCallback *MinLengthFormCallback) {
	minlengthFormCallback = new(MinLengthFormCallback)
	minlengthFormCallback.probe = probe
	minlengthFormCallback.minlength = minlength
	minlengthFormCallback.formGroup = formGroup

	minlengthFormCallback.CreationMode = (minlength == nil)

	return
}

type MinLengthFormCallback struct {
	minlength *models.MinLength

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (minlengthFormCallback *MinLengthFormCallback) OnSave() {

	log.Println("MinLengthFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	minlengthFormCallback.probe.formStage.Checkout()

	if minlengthFormCallback.minlength == nil {
		minlengthFormCallback.minlength = new(models.MinLength).Stage(minlengthFormCallback.probe.stageOfInterest)
	}
	minlength_ := minlengthFormCallback.minlength
	_ = minlength_

	for _, formDiv := range minlengthFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(minlength_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(minlength_.Annotation), minlengthFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(minlength_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if minlengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		minlength_.Unstage(minlengthFormCallback.probe.stageOfInterest)
	}

	minlengthFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.MinLength](
		minlengthFormCallback.probe,
	)
	minlengthFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if minlengthFormCallback.CreationMode || minlengthFormCallback.formGroup.HasSuppressButtonBeenPressed {
		minlengthFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(minlengthFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MinLengthFormCallback(
			nil,
			minlengthFormCallback.probe,
			newFormGroup,
		)
		minlength := new(models.MinLength)
		FillUpForm(minlength, newFormGroup, minlengthFormCallback.probe)
		minlengthFormCallback.probe.formStage.Commit()
	}

	fillUpTree(minlengthFormCallback.probe)
}
func __gong__New__PatternFormCallback(
	pattern *models.Pattern,
	probe *Probe,
	formGroup *table.FormGroup,
) (patternFormCallback *PatternFormCallback) {
	patternFormCallback = new(PatternFormCallback)
	patternFormCallback.probe = probe
	patternFormCallback.pattern = pattern
	patternFormCallback.formGroup = formGroup

	patternFormCallback.CreationMode = (pattern == nil)

	return
}

type PatternFormCallback struct {
	pattern *models.Pattern

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (patternFormCallback *PatternFormCallback) OnSave() {

	log.Println("PatternFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	patternFormCallback.probe.formStage.Checkout()

	if patternFormCallback.pattern == nil {
		patternFormCallback.pattern = new(models.Pattern).Stage(patternFormCallback.probe.stageOfInterest)
	}
	pattern_ := patternFormCallback.pattern
	_ = pattern_

	for _, formDiv := range patternFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pattern_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(pattern_.Annotation), patternFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(pattern_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if patternFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pattern_.Unstage(patternFormCallback.probe.stageOfInterest)
	}

	patternFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Pattern](
		patternFormCallback.probe,
	)
	patternFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if patternFormCallback.CreationMode || patternFormCallback.formGroup.HasSuppressButtonBeenPressed {
		patternFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(patternFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PatternFormCallback(
			nil,
			patternFormCallback.probe,
			newFormGroup,
		)
		pattern := new(models.Pattern)
		FillUpForm(pattern, newFormGroup, patternFormCallback.probe)
		patternFormCallback.probe.formStage.Commit()
	}

	fillUpTree(patternFormCallback.probe)
}
func __gong__New__RestrictionFormCallback(
	restriction *models.Restriction,
	probe *Probe,
	formGroup *table.FormGroup,
) (restrictionFormCallback *RestrictionFormCallback) {
	restrictionFormCallback = new(RestrictionFormCallback)
	restrictionFormCallback.probe = probe
	restrictionFormCallback.restriction = restriction
	restrictionFormCallback.formGroup = formGroup

	restrictionFormCallback.CreationMode = (restriction == nil)

	return
}

type RestrictionFormCallback struct {
	restriction *models.Restriction

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (restrictionFormCallback *RestrictionFormCallback) OnSave() {

	log.Println("RestrictionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	restrictionFormCallback.probe.formStage.Checkout()

	if restrictionFormCallback.restriction == nil {
		restrictionFormCallback.restriction = new(models.Restriction).Stage(restrictionFormCallback.probe.stageOfInterest)
	}
	restriction_ := restrictionFormCallback.restriction
	_ = restriction_

	for _, formDiv := range restrictionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(restriction_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(restriction_.Annotation), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "Base":
			FormDivBasicFieldToField(&(restriction_.Base), formDiv)
		case "MinInclusive":
			FormDivSelectFieldToField(&(restriction_.MinInclusive), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "MaxInclusive":
			FormDivSelectFieldToField(&(restriction_.MaxInclusive), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "Pattern":
			FormDivSelectFieldToField(&(restriction_.Pattern), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "WhiteSpace":
			FormDivSelectFieldToField(&(restriction_.WhiteSpace), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "MinLength":
			FormDivSelectFieldToField(&(restriction_.MinLength), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "MaxLength":
			FormDivSelectFieldToField(&(restriction_.MaxLength), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "Length":
			FormDivSelectFieldToField(&(restriction_.Length), restrictionFormCallback.probe.stageOfInterest, formDiv)
		case "TotalDigit":
			FormDivSelectFieldToField(&(restriction_.TotalDigit), restrictionFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if restrictionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		restriction_.Unstage(restrictionFormCallback.probe.stageOfInterest)
	}

	restrictionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Restriction](
		restrictionFormCallback.probe,
	)
	restrictionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if restrictionFormCallback.CreationMode || restrictionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		restrictionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(restrictionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RestrictionFormCallback(
			nil,
			restrictionFormCallback.probe,
			newFormGroup,
		)
		restriction := new(models.Restriction)
		FillUpForm(restriction, newFormGroup, restrictionFormCallback.probe)
		restrictionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(restrictionFormCallback.probe)
}
func __gong__New__SchemaFormCallback(
	schema *models.Schema,
	probe *Probe,
	formGroup *table.FormGroup,
) (schemaFormCallback *SchemaFormCallback) {
	schemaFormCallback = new(SchemaFormCallback)
	schemaFormCallback.probe = probe
	schemaFormCallback.schema = schema
	schemaFormCallback.formGroup = formGroup

	schemaFormCallback.CreationMode = (schema == nil)

	return
}

type SchemaFormCallback struct {
	schema *models.Schema

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (schemaFormCallback *SchemaFormCallback) OnSave() {

	log.Println("SchemaFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	schemaFormCallback.probe.formStage.Checkout()

	if schemaFormCallback.schema == nil {
		schemaFormCallback.schema = new(models.Schema).Stage(schemaFormCallback.probe.stageOfInterest)
	}
	schema_ := schemaFormCallback.schema
	_ = schema_

	for _, formDiv := range schemaFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(schema_.Name), formDiv)
		case "Xs":
			FormDivBasicFieldToField(&(schema_.Xs), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(schema_.Annotation), schemaFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if schemaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		schema_.Unstage(schemaFormCallback.probe.stageOfInterest)
	}

	schemaFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Schema](
		schemaFormCallback.probe,
	)
	schemaFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if schemaFormCallback.CreationMode || schemaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		schemaFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(schemaFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SchemaFormCallback(
			nil,
			schemaFormCallback.probe,
			newFormGroup,
		)
		schema := new(models.Schema)
		FillUpForm(schema, newFormGroup, schemaFormCallback.probe)
		schemaFormCallback.probe.formStage.Commit()
	}

	fillUpTree(schemaFormCallback.probe)
}
func __gong__New__SequenceFormCallback(
	sequence *models.Sequence,
	probe *Probe,
	formGroup *table.FormGroup,
) (sequenceFormCallback *SequenceFormCallback) {
	sequenceFormCallback = new(SequenceFormCallback)
	sequenceFormCallback.probe = probe
	sequenceFormCallback.sequence = sequence
	sequenceFormCallback.formGroup = formGroup

	sequenceFormCallback.CreationMode = (sequence == nil)

	return
}

type SequenceFormCallback struct {
	sequence *models.Sequence

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (sequenceFormCallback *SequenceFormCallback) OnSave() {

	log.Println("SequenceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sequenceFormCallback.probe.formStage.Checkout()

	if sequenceFormCallback.sequence == nil {
		sequenceFormCallback.sequence = new(models.Sequence).Stage(sequenceFormCallback.probe.stageOfInterest)
	}
	sequence_ := sequenceFormCallback.sequence
	_ = sequence_

	for _, formDiv := range sequenceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(sequence_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(sequence_.Annotation), sequenceFormCallback.probe.stageOfInterest, formDiv)
		case "MinOccurs":
			FormDivBasicFieldToField(&(sequence_.MinOccurs), formDiv)
		case "MaxOccurs":
			FormDivBasicFieldToField(&(sequence_.MaxOccurs), formDiv)
		case "All:Sequences":
			// we need to retrieve the field owner before the change
			var pastAllOwner *models.All
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "All"
			rf.Fieldname = "Sequences"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				sequenceFormCallback.probe.stageOfInterest,
				sequenceFormCallback.probe.backRepoOfInterest,
				sequence_,
				&rf)

			if reverseFieldOwner != nil {
				pastAllOwner = reverseFieldOwner.(*models.All)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAllOwner != nil {
					idx := slices.Index(pastAllOwner.Sequences, sequence_)
					pastAllOwner.Sequences = slices.Delete(pastAllOwner.Sequences, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _all := range *models.GetGongstructInstancesSet[models.All](sequenceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _all.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAllOwner := _all // we have a match
						if pastAllOwner != nil {
							if newAllOwner != pastAllOwner {
								idx := slices.Index(pastAllOwner.Sequences, sequence_)
								pastAllOwner.Sequences = slices.Delete(pastAllOwner.Sequences, idx, idx+1)
								newAllOwner.Sequences = append(newAllOwner.Sequences, sequence_)
							}
						} else {
							newAllOwner.Sequences = append(newAllOwner.Sequences, sequence_)
						}
					}
				}
			}
		case "Choice:Sequences":
			// we need to retrieve the field owner before the change
			var pastChoiceOwner *models.Choice
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Choice"
			rf.Fieldname = "Sequences"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				sequenceFormCallback.probe.stageOfInterest,
				sequenceFormCallback.probe.backRepoOfInterest,
				sequence_,
				&rf)

			if reverseFieldOwner != nil {
				pastChoiceOwner = reverseFieldOwner.(*models.Choice)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastChoiceOwner != nil {
					idx := slices.Index(pastChoiceOwner.Sequences, sequence_)
					pastChoiceOwner.Sequences = slices.Delete(pastChoiceOwner.Sequences, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _choice := range *models.GetGongstructInstancesSet[models.Choice](sequenceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _choice.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newChoiceOwner := _choice // we have a match
						if pastChoiceOwner != nil {
							if newChoiceOwner != pastChoiceOwner {
								idx := slices.Index(pastChoiceOwner.Sequences, sequence_)
								pastChoiceOwner.Sequences = slices.Delete(pastChoiceOwner.Sequences, idx, idx+1)
								newChoiceOwner.Sequences = append(newChoiceOwner.Sequences, sequence_)
							}
						} else {
							newChoiceOwner.Sequences = append(newChoiceOwner.Sequences, sequence_)
						}
					}
				}
			}
		case "ComplexType:Sequences":
			// we need to retrieve the field owner before the change
			var pastComplexTypeOwner *models.ComplexType
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "ComplexType"
			rf.Fieldname = "Sequences"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				sequenceFormCallback.probe.stageOfInterest,
				sequenceFormCallback.probe.backRepoOfInterest,
				sequence_,
				&rf)

			if reverseFieldOwner != nil {
				pastComplexTypeOwner = reverseFieldOwner.(*models.ComplexType)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastComplexTypeOwner != nil {
					idx := slices.Index(pastComplexTypeOwner.Sequences, sequence_)
					pastComplexTypeOwner.Sequences = slices.Delete(pastComplexTypeOwner.Sequences, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _complextype := range *models.GetGongstructInstancesSet[models.ComplexType](sequenceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _complextype.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newComplexTypeOwner := _complextype // we have a match
						if pastComplexTypeOwner != nil {
							if newComplexTypeOwner != pastComplexTypeOwner {
								idx := slices.Index(pastComplexTypeOwner.Sequences, sequence_)
								pastComplexTypeOwner.Sequences = slices.Delete(pastComplexTypeOwner.Sequences, idx, idx+1)
								newComplexTypeOwner.Sequences = append(newComplexTypeOwner.Sequences, sequence_)
							}
						} else {
							newComplexTypeOwner.Sequences = append(newComplexTypeOwner.Sequences, sequence_)
						}
					}
				}
			}
		case "Extension:Sequences":
			// we need to retrieve the field owner before the change
			var pastExtensionOwner *models.Extension
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Extension"
			rf.Fieldname = "Sequences"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				sequenceFormCallback.probe.stageOfInterest,
				sequenceFormCallback.probe.backRepoOfInterest,
				sequence_,
				&rf)

			if reverseFieldOwner != nil {
				pastExtensionOwner = reverseFieldOwner.(*models.Extension)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastExtensionOwner != nil {
					idx := slices.Index(pastExtensionOwner.Sequences, sequence_)
					pastExtensionOwner.Sequences = slices.Delete(pastExtensionOwner.Sequences, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _extension := range *models.GetGongstructInstancesSet[models.Extension](sequenceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _extension.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newExtensionOwner := _extension // we have a match
						if pastExtensionOwner != nil {
							if newExtensionOwner != pastExtensionOwner {
								idx := slices.Index(pastExtensionOwner.Sequences, sequence_)
								pastExtensionOwner.Sequences = slices.Delete(pastExtensionOwner.Sequences, idx, idx+1)
								newExtensionOwner.Sequences = append(newExtensionOwner.Sequences, sequence_)
							}
						} else {
							newExtensionOwner.Sequences = append(newExtensionOwner.Sequences, sequence_)
						}
					}
				}
			}
		case "Group:Sequences":
			// we need to retrieve the field owner before the change
			var pastGroupOwner *models.Group
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Group"
			rf.Fieldname = "Sequences"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				sequenceFormCallback.probe.stageOfInterest,
				sequenceFormCallback.probe.backRepoOfInterest,
				sequence_,
				&rf)

			if reverseFieldOwner != nil {
				pastGroupOwner = reverseFieldOwner.(*models.Group)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGroupOwner != nil {
					idx := slices.Index(pastGroupOwner.Sequences, sequence_)
					pastGroupOwner.Sequences = slices.Delete(pastGroupOwner.Sequences, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _group := range *models.GetGongstructInstancesSet[models.Group](sequenceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _group.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGroupOwner := _group // we have a match
						if pastGroupOwner != nil {
							if newGroupOwner != pastGroupOwner {
								idx := slices.Index(pastGroupOwner.Sequences, sequence_)
								pastGroupOwner.Sequences = slices.Delete(pastGroupOwner.Sequences, idx, idx+1)
								newGroupOwner.Sequences = append(newGroupOwner.Sequences, sequence_)
							}
						} else {
							newGroupOwner.Sequences = append(newGroupOwner.Sequences, sequence_)
						}
					}
				}
			}
		case "Sequence:Sequences":
			// we need to retrieve the field owner before the change
			var pastSequenceOwner *models.Sequence
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Sequence"
			rf.Fieldname = "Sequences"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				sequenceFormCallback.probe.stageOfInterest,
				sequenceFormCallback.probe.backRepoOfInterest,
				sequence_,
				&rf)

			if reverseFieldOwner != nil {
				pastSequenceOwner = reverseFieldOwner.(*models.Sequence)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSequenceOwner != nil {
					idx := slices.Index(pastSequenceOwner.Sequences, sequence_)
					pastSequenceOwner.Sequences = slices.Delete(pastSequenceOwner.Sequences, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _sequence := range *models.GetGongstructInstancesSet[models.Sequence](sequenceFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _sequence.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSequenceOwner := _sequence // we have a match
						if pastSequenceOwner != nil {
							if newSequenceOwner != pastSequenceOwner {
								idx := slices.Index(pastSequenceOwner.Sequences, sequence_)
								pastSequenceOwner.Sequences = slices.Delete(pastSequenceOwner.Sequences, idx, idx+1)
								newSequenceOwner.Sequences = append(newSequenceOwner.Sequences, sequence_)
							}
						} else {
							newSequenceOwner.Sequences = append(newSequenceOwner.Sequences, sequence_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if sequenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sequence_.Unstage(sequenceFormCallback.probe.stageOfInterest)
	}

	sequenceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Sequence](
		sequenceFormCallback.probe,
	)
	sequenceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if sequenceFormCallback.CreationMode || sequenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sequenceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(sequenceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SequenceFormCallback(
			nil,
			sequenceFormCallback.probe,
			newFormGroup,
		)
		sequence := new(models.Sequence)
		FillUpForm(sequence, newFormGroup, sequenceFormCallback.probe)
		sequenceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(sequenceFormCallback.probe)
}
func __gong__New__SimpleContentFormCallback(
	simplecontent *models.SimpleContent,
	probe *Probe,
	formGroup *table.FormGroup,
) (simplecontentFormCallback *SimpleContentFormCallback) {
	simplecontentFormCallback = new(SimpleContentFormCallback)
	simplecontentFormCallback.probe = probe
	simplecontentFormCallback.simplecontent = simplecontent
	simplecontentFormCallback.formGroup = formGroup

	simplecontentFormCallback.CreationMode = (simplecontent == nil)

	return
}

type SimpleContentFormCallback struct {
	simplecontent *models.SimpleContent

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (simplecontentFormCallback *SimpleContentFormCallback) OnSave() {

	log.Println("SimpleContentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	simplecontentFormCallback.probe.formStage.Checkout()

	if simplecontentFormCallback.simplecontent == nil {
		simplecontentFormCallback.simplecontent = new(models.SimpleContent).Stage(simplecontentFormCallback.probe.stageOfInterest)
	}
	simplecontent_ := simplecontentFormCallback.simplecontent
	_ = simplecontent_

	for _, formDiv := range simplecontentFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(simplecontent_.Name), formDiv)
		case "Extension":
			FormDivSelectFieldToField(&(simplecontent_.Extension), simplecontentFormCallback.probe.stageOfInterest, formDiv)
		case "Restriction":
			FormDivSelectFieldToField(&(simplecontent_.Restriction), simplecontentFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if simplecontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		simplecontent_.Unstage(simplecontentFormCallback.probe.stageOfInterest)
	}

	simplecontentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SimpleContent](
		simplecontentFormCallback.probe,
	)
	simplecontentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if simplecontentFormCallback.CreationMode || simplecontentFormCallback.formGroup.HasSuppressButtonBeenPressed {
		simplecontentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(simplecontentFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SimpleContentFormCallback(
			nil,
			simplecontentFormCallback.probe,
			newFormGroup,
		)
		simplecontent := new(models.SimpleContent)
		FillUpForm(simplecontent, newFormGroup, simplecontentFormCallback.probe)
		simplecontentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(simplecontentFormCallback.probe)
}
func __gong__New__SimpleTypeFormCallback(
	simpletype *models.SimpleType,
	probe *Probe,
	formGroup *table.FormGroup,
) (simpletypeFormCallback *SimpleTypeFormCallback) {
	simpletypeFormCallback = new(SimpleTypeFormCallback)
	simpletypeFormCallback.probe = probe
	simpletypeFormCallback.simpletype = simpletype
	simpletypeFormCallback.formGroup = formGroup

	simpletypeFormCallback.CreationMode = (simpletype == nil)

	return
}

type SimpleTypeFormCallback struct {
	simpletype *models.SimpleType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (simpletypeFormCallback *SimpleTypeFormCallback) OnSave() {

	log.Println("SimpleTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	simpletypeFormCallback.probe.formStage.Checkout()

	if simpletypeFormCallback.simpletype == nil {
		simpletypeFormCallback.simpletype = new(models.SimpleType).Stage(simpletypeFormCallback.probe.stageOfInterest)
	}
	simpletype_ := simpletypeFormCallback.simpletype
	_ = simpletype_

	for _, formDiv := range simpletypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(simpletype_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(simpletype_.Annotation), simpletypeFormCallback.probe.stageOfInterest, formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(simpletype_.NameXSD), formDiv)
		case "Restriction":
			FormDivSelectFieldToField(&(simpletype_.Restriction), simpletypeFormCallback.probe.stageOfInterest, formDiv)
		case "Union":
			FormDivSelectFieldToField(&(simpletype_.Union), simpletypeFormCallback.probe.stageOfInterest, formDiv)
		case "Schema:SimpleTypes":
			// we need to retrieve the field owner before the change
			var pastSchemaOwner *models.Schema
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Schema"
			rf.Fieldname = "SimpleTypes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				simpletypeFormCallback.probe.stageOfInterest,
				simpletypeFormCallback.probe.backRepoOfInterest,
				simpletype_,
				&rf)

			if reverseFieldOwner != nil {
				pastSchemaOwner = reverseFieldOwner.(*models.Schema)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSchemaOwner != nil {
					idx := slices.Index(pastSchemaOwner.SimpleTypes, simpletype_)
					pastSchemaOwner.SimpleTypes = slices.Delete(pastSchemaOwner.SimpleTypes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _schema := range *models.GetGongstructInstancesSet[models.Schema](simpletypeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _schema.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSchemaOwner := _schema // we have a match
						if pastSchemaOwner != nil {
							if newSchemaOwner != pastSchemaOwner {
								idx := slices.Index(pastSchemaOwner.SimpleTypes, simpletype_)
								pastSchemaOwner.SimpleTypes = slices.Delete(pastSchemaOwner.SimpleTypes, idx, idx+1)
								newSchemaOwner.SimpleTypes = append(newSchemaOwner.SimpleTypes, simpletype_)
							}
						} else {
							newSchemaOwner.SimpleTypes = append(newSchemaOwner.SimpleTypes, simpletype_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if simpletypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		simpletype_.Unstage(simpletypeFormCallback.probe.stageOfInterest)
	}

	simpletypeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SimpleType](
		simpletypeFormCallback.probe,
	)
	simpletypeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if simpletypeFormCallback.CreationMode || simpletypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		simpletypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(simpletypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SimpleTypeFormCallback(
			nil,
			simpletypeFormCallback.probe,
			newFormGroup,
		)
		simpletype := new(models.SimpleType)
		FillUpForm(simpletype, newFormGroup, simpletypeFormCallback.probe)
		simpletypeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(simpletypeFormCallback.probe)
}
func __gong__New__TotalDigitFormCallback(
	totaldigit *models.TotalDigit,
	probe *Probe,
	formGroup *table.FormGroup,
) (totaldigitFormCallback *TotalDigitFormCallback) {
	totaldigitFormCallback = new(TotalDigitFormCallback)
	totaldigitFormCallback.probe = probe
	totaldigitFormCallback.totaldigit = totaldigit
	totaldigitFormCallback.formGroup = formGroup

	totaldigitFormCallback.CreationMode = (totaldigit == nil)

	return
}

type TotalDigitFormCallback struct {
	totaldigit *models.TotalDigit

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (totaldigitFormCallback *TotalDigitFormCallback) OnSave() {

	log.Println("TotalDigitFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	totaldigitFormCallback.probe.formStage.Checkout()

	if totaldigitFormCallback.totaldigit == nil {
		totaldigitFormCallback.totaldigit = new(models.TotalDigit).Stage(totaldigitFormCallback.probe.stageOfInterest)
	}
	totaldigit_ := totaldigitFormCallback.totaldigit
	_ = totaldigit_

	for _, formDiv := range totaldigitFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(totaldigit_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(totaldigit_.Annotation), totaldigitFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(totaldigit_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if totaldigitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		totaldigit_.Unstage(totaldigitFormCallback.probe.stageOfInterest)
	}

	totaldigitFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.TotalDigit](
		totaldigitFormCallback.probe,
	)
	totaldigitFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if totaldigitFormCallback.CreationMode || totaldigitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		totaldigitFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(totaldigitFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TotalDigitFormCallback(
			nil,
			totaldigitFormCallback.probe,
			newFormGroup,
		)
		totaldigit := new(models.TotalDigit)
		FillUpForm(totaldigit, newFormGroup, totaldigitFormCallback.probe)
		totaldigitFormCallback.probe.formStage.Commit()
	}

	fillUpTree(totaldigitFormCallback.probe)
}
func __gong__New__UnionFormCallback(
	union *models.Union,
	probe *Probe,
	formGroup *table.FormGroup,
) (unionFormCallback *UnionFormCallback) {
	unionFormCallback = new(UnionFormCallback)
	unionFormCallback.probe = probe
	unionFormCallback.union = union
	unionFormCallback.formGroup = formGroup

	unionFormCallback.CreationMode = (union == nil)

	return
}

type UnionFormCallback struct {
	union *models.Union

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (unionFormCallback *UnionFormCallback) OnSave() {

	log.Println("UnionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	unionFormCallback.probe.formStage.Checkout()

	if unionFormCallback.union == nil {
		unionFormCallback.union = new(models.Union).Stage(unionFormCallback.probe.stageOfInterest)
	}
	union_ := unionFormCallback.union
	_ = union_

	for _, formDiv := range unionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(union_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(union_.Annotation), unionFormCallback.probe.stageOfInterest, formDiv)
		case "MemberTypes":
			FormDivBasicFieldToField(&(union_.MemberTypes), formDiv)
		}
	}

	// manage the suppress operation
	if unionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		union_.Unstage(unionFormCallback.probe.stageOfInterest)
	}

	unionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Union](
		unionFormCallback.probe,
	)
	unionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if unionFormCallback.CreationMode || unionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		unionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(unionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UnionFormCallback(
			nil,
			unionFormCallback.probe,
			newFormGroup,
		)
		union := new(models.Union)
		FillUpForm(union, newFormGroup, unionFormCallback.probe)
		unionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(unionFormCallback.probe)
}
func __gong__New__WhiteSpaceFormCallback(
	whitespace *models.WhiteSpace,
	probe *Probe,
	formGroup *table.FormGroup,
) (whitespaceFormCallback *WhiteSpaceFormCallback) {
	whitespaceFormCallback = new(WhiteSpaceFormCallback)
	whitespaceFormCallback.probe = probe
	whitespaceFormCallback.whitespace = whitespace
	whitespaceFormCallback.formGroup = formGroup

	whitespaceFormCallback.CreationMode = (whitespace == nil)

	return
}

type WhiteSpaceFormCallback struct {
	whitespace *models.WhiteSpace

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (whitespaceFormCallback *WhiteSpaceFormCallback) OnSave() {

	log.Println("WhiteSpaceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	whitespaceFormCallback.probe.formStage.Checkout()

	if whitespaceFormCallback.whitespace == nil {
		whitespaceFormCallback.whitespace = new(models.WhiteSpace).Stage(whitespaceFormCallback.probe.stageOfInterest)
	}
	whitespace_ := whitespaceFormCallback.whitespace
	_ = whitespace_

	for _, formDiv := range whitespaceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(whitespace_.Name), formDiv)
		case "Annotation":
			FormDivSelectFieldToField(&(whitespace_.Annotation), whitespaceFormCallback.probe.stageOfInterest, formDiv)
		case "Value":
			FormDivBasicFieldToField(&(whitespace_.Value), formDiv)
		}
	}

	// manage the suppress operation
	if whitespaceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		whitespace_.Unstage(whitespaceFormCallback.probe.stageOfInterest)
	}

	whitespaceFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.WhiteSpace](
		whitespaceFormCallback.probe,
	)
	whitespaceFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if whitespaceFormCallback.CreationMode || whitespaceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		whitespaceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(whitespaceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__WhiteSpaceFormCallback(
			nil,
			whitespaceFormCallback.probe,
			newFormGroup,
		)
		whitespace := new(models.WhiteSpace)
		FillUpForm(whitespace, newFormGroup, whitespaceFormCallback.probe)
		whitespaceFormCallback.probe.formStage.Commit()
	}

	fillUpTree(whitespaceFormCallback.probe)
}
