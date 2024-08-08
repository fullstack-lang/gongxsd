package models

func PostProcessing(stage *StageStruct) {
	PostProcessingUpdateNames(stage)
	PostProcessingAnalyzeXSDStructure(stage)
}
