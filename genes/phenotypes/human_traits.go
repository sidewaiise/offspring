package phenotypes

import (
	"github.com/sidewaiise/offspring/genetic/genome"
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

var HumanAnecdotalPhenome = genome.Genome{
	StrongMusclesTrait: genome.Gene{
		Desc:  "Strong muscles",
		Score: 1,
	},
	WeakMusclesTrait: genome.Gene{
		Desc:  "Weak muscles",
		Score: -1,
	},
	FastTwitchMusclesTrait: genome.Gene{
		Desc:  "Fast twitch muscles",
		Score: 1,
	},
	SlowTwitchMusclesTrait: genome.Gene{
		Desc:  "Slow twitch muscles",
		Score: 1,
	},
	EnduranceTrait: genome.Gene{
		Desc:  "Good endurance",
		Score: 1,
	},
	StrongMemoryTrait: genome.Gene{
		Desc:  "Good memory",
		Score: 1,
	},
	WeakMemoryTrait: genome.Gene{
		Desc:  "Trouble with gender identity",
		Score: 0,
	},
	DistractedTrait: genome.Gene{
		Desc:  "Easily distracted",
		Score: -1,
	},
	FocusedTrait: genome.Gene{
		Desc:  "Easily distracted",
		Score: 1,
	},
	AdaptiveTrait: genome.Gene{
		Desc:  "Adapts easily",
		Score: 1,
	},
	UnadaptiveTrait: genome.Gene{
		Desc:  "Does not adapt easily",
		Score: 1,
	},
	IntrovertedTrait: genome.Gene{
		Desc:  "Introverted, takes less social interaction for stimulation",
		Score: 1,
	},
	ExtrovertedTrait: genome.Gene{
		Desc:  "Extraverted, takes more social interaction for stimulation",
		Score: 1,
	},
	GoodHeartTrait: genome.Gene{
		Desc:  "Very healthy heart",
		Score: 1,
	},
	BadHeartTrait: genome.Gene{
		Desc:  "Heart problems",
		Score: -1,
	},
	CuriousTrait: genome.Gene{
		Desc:  "Curious mind",
		Score: 1,
	},
	DestructiveTrait: genome.Gene{
		Desc:  "Destructive mindset",
		Score: -1,
	},
	NatureLoverTrait: genome.Gene{
		Desc:  "Loves animals and nature",
		Score: 1,
	},
	CityLoverTrait: genome.Gene{
		Desc:  "Loves big cities",
		Score: 1,
	},
	KindnessTrait: genome.Gene{
		Desc:  "Kind to others",
		Score: 1,
	},
	MeanTrait: genome.Gene{
		Desc:  "Mean to others",
		Score: -1.5,
	},
	ScientificTrait: genome.Gene{
		Desc:  "Is scientific",
		Score: 1,
	},
	NeuroticTrait: genome.Gene{
		Desc:  "Neurotic",
		Score: -1,
	},
	PsychoticTrait: genome.Gene{
		Desc:  "Psychotic",
		Score: -1,
	},
	OCDTrait: genome.Gene{
		Desc:  "Obsessive compulsive",
		Score: -1,
	},
	AutismTrait: genome.Gene{
		Desc:  "Autistic",
		Score: -1,
	},
	AnxietyTrait: genome.Gene{
		Desc:  "Has anxiety",
		Score: -1,
	},
	CreativityTrait: genome.Gene{
		Desc:  "Is creative",
		Score: 1,
	},
	OrganisedTrait: genome.Gene{
		Desc:  "Is organised",
		Score: 1,
	},
	ConfusedGenderTrait: genome.Gene{
		Desc:  "Trouble with gender identity",
		Score: -1,
	},
	HomosexualTrait: genome.Gene{
		Desc:  "Non-hetero sexual preference",
		Score: 0,
	},
	LeftWingPoliticsTrait: genome.Gene{
		Desc:  "Left wing political views",
		Score: 0,
	},
	RightWingPoliticsTrait: genome.Gene{
		Desc:  "Right wing political views",
		Score: 0,
	},
	ExtremeLeftPoliticsTrait: genome.Gene{
		Desc:  "Extreme Left wing political views",
		Score: -1,
	},
	ExtremeRightPoliticsTrait: genome.Gene{
		Desc:  "Extreme Right wing political views",
		Score: -1,
	},
	ProblemSolverTrait: genome.Gene{
		Desc:  "Problem solver",
		Score: 1,
	},
	CapitalistTrait: genome.Gene{
		Desc:  "Capitalist",
		Score: 0,
	},
	SocialistTrait: genome.Gene{
		Desc:  "Socialist",
		Score: 0,
	},
	HumanitarianTrait: genome.Gene{
		Desc:  "Humanitarian",
		Score: 1,
	},
	ThrillseekerTrait: genome.Gene{
		Desc:  "Thrillseeker",
		Score: 1,
	},
	ExplorerTrait: genome.Gene{
		Desc:  "Explorer",
		Score: 1,
	},
	HermitTrait: genome.Gene{
		Desc:  "Is a Hermit",
		Score: 1,
	},
	MaleGenderTrait: genome.Gene{
		Desc:  "Has Male gender chomosomes XX",
		Score: 1,
	},
	FemaleGenderTrait: genome.Gene{
		Desc:  "Has Female gender chromosomes XY",
		Score: 1,
	},
	GoodLookingTrait: genome.Gene{
		Desc:  "Good Looking",
		Score: 1,
	},
	UglyTrait: genome.Gene{
		Desc:  "Not great looking",
		Score: -1,
	},
	CriminalTendancyTrait: genome.Gene{
		Desc:  "Criminal tendancies",
		Score: -1,
	},
	RamblingSpeechTrait: genome.Gene{
		Desc:  "Tends to ramble",
		Score: -1,
	},
	DryHumourTrait: genome.Gene{
		Desc:  "Tends to ramble",
		Score: -1,
	},
	AthleticTrait: genome.Gene{
		Desc:  "Athletic",
		Score: 1,
	},
	CuteTrait: genome.Gene{
		Desc:  "Cute",
		Score: 1,
	},
	DarkHumourTrait: genome.Gene{
		Desc:  "Has Dark humour",
		Score: 1,
	},
	ComedyTrait: genome.Gene{
		Desc:  "Likes all comedy",
		Score: 1,
	},
}
