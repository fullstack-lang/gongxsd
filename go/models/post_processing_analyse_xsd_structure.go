package models

func PostProcessingAnalyzeXSDStructure(stage *StageStruct) {

	// characterize complex types that are inlined
	for element := range *GetGongstructInstancesSet[Element](stage) {

		if element.ComplexType != nil {
			element.ComplexType.IsInlined = true
			element.ComplexType.EnclosingElement = element
		}
	}


}
