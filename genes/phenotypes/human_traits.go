package phenotypes

import (
	genetic "github.com/sidewaiise/offspring/genetic/genome"
)

const (
	StrongMusclesTrait uint64 = iota
	WeakMusclesTrait
	FastTwitchMusclesTrait
	SlowTwitchMusclesTrait
	EnduranceTrait
	StrongMemoryTrait
	WeakMemoryTrait
	DistractedTrait
	FocusedTrait
	AdaptiveTrait
	UnadaptiveTrait
	IntrovertedTrait
	ExtrovertedTrait
	GoodHeartTrait
	BadHeartTrait
	CuriousTrait
	DestructiveTrait
	NatureLoverTrait
	CityLoverTrait
	KindnessTrait
	MeanTrait
	ScientificTrait
	NeuroticTrait
	PsychoticTrait
	OCDTrait
	AutismTrait
	AnxietyTrait
	CreativityTrait
	OrganisedTrait
	ConfusedGenderTrait
	HomosexualTrait
	LeftWingPoliticsTrait
	RightWingPoliticsTrait
	ExtremeLeftPoliticsTrait
	ExtremeRightPoliticsTrait
	ProblemSolverTrait
	CapitalistTrait
	SocialistTrait
	HumanitarianTrait
	ThrillseekerTrait
	ExplorerTrait
	HermitTrait
	MaleGenderTrait
	FemaleGenderTrait
	GoodLookingTrait
	AthleticTrait
	CuteTrait
	UglyTrait
	CriminalTendancyTrait
	RamblingSpeechTrait
	DryHumourTrait
	DarkHumourTrait
	ComedyTrait
)

var HumanAnecdotalPhenome = genetic.Genome{
	StrongMusclesTrait: genetic.Gene{
		Desc:  "Strong muscles",
		Score: 1,
	},
	WeakMusclesTrait: genetic.Gene{
		Desc:  "Weak muscles",
		Score: -1,
	},
	FastTwitchMusclesTrait: genetic.Gene{
		Desc:  "Fast twitch muscles",
		Score: 1,
	},
	SlowTwitchMusclesTrait: genetic.Gene{
		Desc:  "Slow twitch muscles",
		Score: 1,
	},
	EnduranceTrait: genetic.Gene{
		Desc:  "Good endurance",
		Score: 1,
	},
	StrongMemoryTrait: genetic.Gene{
		Desc:  "Good memory",
		Score: 1,
	},
	WeakMemoryTrait: genetic.Gene{
		Desc:  "Trouble with gender identity",
		Score: 0,
	},
	DistractedTrait: genetic.Gene{
		Desc:  "Easily distracted",
		Score: -1,
	},
	FocusedTrait: genetic.Gene{
		Desc:  "Easily distracted",
		Score: 1,
	},
	AdaptiveTrait: genetic.Gene{
		Desc:  "Adapts easily",
		Score: 1,
	},
	UnadaptiveTrait: genetic.Gene{
		Desc:  "Does not adapt easily",
		Score: 1,
	},
	IntrovertedTrait: genetic.Gene{
		Desc:  "Introverted, takes less social interaction for stimulation",
		Score: 1,
	},
	ExtrovertedTrait: genetic.Gene{
		Desc:  "Extraverted, takes more social interaction for stimulation",
		Score: 1,
	},
	GoodHeartTrait: genetic.Gene{
		Desc:  "Very healthy heart",
		Score: 1,
	},
	BadHeartTrait: genetic.Gene{
		Desc:  "Heart problems",
		Score: -1,
	},
	CuriousTrait: genetic.Gene{
		Desc:  "Curious mind",
		Score: 1,
	},
	DestructiveTrait: genetic.Gene{
		Desc:  "Destructive mindset",
		Score: -1,
	},
	NatureLoverTrait: genetic.Gene{
		Desc:  "Loves animals and nature",
		Score: 1,
	},
	CityLoverTrait: genetic.Gene{
		Desc:  "Loves big cities",
		Score: 1,
	},
	KindnessTrait: genetic.Gene{
		Desc:  "Kind to others",
		Score: 1,
	},
	MeanTrait: genetic.Gene{
		Desc:  "Mean to others",
		Score: -1.5,
	},
	ScientificTrait: genetic.Gene{
		Desc:  "Is scientific",
		Score: 1,
	},
	NeuroticTrait: genetic.Gene{
		Desc:  "Neurotic",
		Score: -1,
	},
	PsychoticTrait: genetic.Gene{
		Desc:  "Psychotic",
		Score: -1,
	},
	OCDTrait: genetic.Gene{
		Desc:  "Obsessive compulsive",
		Score: -1,
	},
	AutismTrait: genetic.Gene{
		Desc:  "Autistic",
		Score: -1,
	},
	AnxietyTrait: genetic.Gene{
		Desc:  "Has anxiety",
		Score: -1,
	},
	CreativityTrait: genetic.Gene{
		Desc:  "Is creative",
		Score: 1,
	},
	OrganisedTrait: genetic.Gene{
		Desc:  "Is organised",
		Score: 1,
	},
	ConfusedGenderTrait: genetic.Gene{
		Desc:  "Trouble with gender identity",
		Score: -1,
	},
	HomosexualTrait: genetic.Gene{
		Desc:  "Non-hetero sexual preference",
		Score: 0,
	},
	LeftWingPoliticsTrait: genetic.Gene{
		Desc:  "Left wing political views",
		Score: 0,
	},
	RightWingPoliticsTrait: genetic.Gene{
		Desc:  "Right wing political views",
		Score: 0,
	},
	ExtremeLeftPoliticsTrait: genetic.Gene{
		Desc:  "Extreme Left wing political views",
		Score: -1,
	},
	ExtremeRightPoliticsTrait: genetic.Gene{
		Desc:  "Extreme Right wing political views",
		Score: -1,
	},
	ProblemSolverTrait: genetic.Gene{
		Desc:  "Problem solver",
		Score: 1,
	},
	CapitalistTrait: genetic.Gene{
		Desc:  "Capitalist",
		Score: 0,
	},
	SocialistTrait: genetic.Gene{
		Desc:  "Socialist",
		Score: 0,
	},
	HumanitarianTrait: genetic.Gene{
		Desc:  "Humanitarian",
		Score: 1,
	},
	ThrillseekerTrait: genetic.Gene{
		Desc:  "Thrillseeker",
		Score: 1,
	},
	ExplorerTrait: genetic.Gene{
		Desc:  "Explorer",
		Score: 1,
	},
	HermitTrait: genetic.Gene{
		Desc:  "Is a Hermit",
		Score: 1,
	},
	MaleGenderTrait: genetic.Gene{
		Desc:  "Has Male gender chomosomes XX",
		Score: 1,
	},
	FemaleGenderTrait: genetic.Gene{
		Desc:  "Has Female gender chromosomes XY",
		Score: 1,
	},
	GoodLookingTrait: genetic.Gene{
		Desc:  "Good Looking",
		Score: 1,
	},
	UglyTrait: genetic.Gene{
		Desc:  "Not great looking",
		Score: -1,
	},
	CriminalTendancyTrait: genetic.Gene{
		Desc:  "Criminal tendancies",
		Score: -1,
	},
	RamblingSpeechTrait: genetic.Gene{
		Desc:  "Tends to ramble",
		Score: -1,
	},
	DryHumourTrait: genetic.Gene{
		Desc:  "Tends to ramble",
		Score: -1,
	},
	AthleticTrait: genetic.Gene{
		Desc:  "Athletic",
		Score: 1,
	},
	CuteTrait: genetic.Gene{
		Desc:  "Cute",
		Score: 1,
	},
	DarkHumourTrait: genetic.Gene{
		Desc:  "Has Dark humour",
		Score: 1,
	},
	ComedyTrait: genetic.Gene{
		Desc:  "Likes all comedy",
		Score: 1,
	},
}

func NewHumanTraits(keys []uint64) genetic.Genes {
	genes := genetic.NewGenes(HumanAnecdotalPhenome)

	for _, gene := range keys {
		genes.Set(gene)
	}

	return genes
}
