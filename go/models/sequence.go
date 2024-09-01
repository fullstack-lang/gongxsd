package models

type Sequence struct {
	Name string
	Annotated

	OccurrenceDefinitionAbstract

	ModelGroup

	ParticleAbstract

	OuterParticleOwnerAbstract
}
