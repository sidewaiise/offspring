package genes

import (
	"github.com/sidewaiise/offspring/chromosomes"
)

var GenericHumanGenome = chromosomes.Genome{
	"strong_muscles": {
		Desc:  "Strong muscles",
		Score: 1,
	},
	"weak_muscles": {
		Desc:  "Weak muscles",
		Score: -1,
	},
	"fast_muscles": {
		Desc:  "Fast twitch muscles",
		Score: 1,
	},
	"slow_muscles": {
		Desc:  "Slow twitch muscles",
		Score: 1,
	},
	"endurance": {
		Desc:  "Good endurance",
		Score: 1,
	},
	"strong_memory": {
		Desc:  "Good memory",
		Score: 1,
	},
	"weak_memory": {
		Desc:  "Trouble with gender identity",
		Score: 0,
	},
	"distracted": {
		Desc:  "Easily distracted",
		Score: -1,
	},
	"focused": {
		Desc:  "Easily distracted",
		Score: 1,
	},
	"adaptive": {
		Desc:  "Adapts easily",
		Score: 1,
	},
	"rigid": {
		Desc:  "Does not adapt easily",
		Score: 1,
	},
	"introverted": {
		Desc:  "Introverted, takes less social interaction for stimulation",
		Score: 1,
	},
	"extroverted": {
		Desc:  "Extraverted, takes more social interaction for stimulation",
		Score: 1,
	},
	"healthy_heart": {
		Desc:  "Very healthy heart",
		Score: 1,
	},
	"unhealthy_heart": {
		Desc:  "Heart problems",
		Score: -1,
	},
	"curiosity": {
		Desc:  "Curious mind",
		Score: 1,
	},
	"destructive": {
		Desc:  "Destructive mindset",
		Score: -1,
	},
	"nature_lover": {
		Desc:  "Loves animals and nature",
		Score: 1,
	},
	"urban_lover": {
		Desc:  "Loves big cities",
		Score: 1,
	},
	"kindness": {
		Desc:  "Kind to others",
		Score: 1,
	},
	"mean": {
		Desc:  "Mean to others",
		Score: -1.5,
	},
	"scientific": {
		Desc:  "Is scientific",
		Score: 1,
	},
	"neurosis": {
		Desc:  "Neurotic",
		Score: -1,
	},
	"psychosis": {
		Desc:  "Psychotic",
		Score: -1,
	},
	"ocd": {
		Desc:  "Obsessive compulsive",
		Score: -1,
	},
	"autistic": {
		Desc:  "Autistic",
		Score: -1,
	},
	"anxious": {
		Desc:  "Has anxiety",
		Score: -1,
	},
	"creative": {
		Desc:  "Is creative",
		Score: 1,
	},
	"organised": {
		Desc:  "Is organised",
		Score: 1,
	},
	"gender_identity": {
		Desc:  "Trouble with gender identity",
		Score: -1,
	},
	"sexual_preference": {
		Desc:  "Non-hetero sexual preference",
		Score: 0,
	},
	"left_wing": {
		Desc:  "Left wing political views",
		Score: 0,
	},
	"right_wing": {
		Desc:  "Right wing political views",
		Score: 0,
	},
	"extreme_left": {
		Desc:  "Extreme Left wing political views",
		Score: -1,
	},
	"extreme_right": {
		Desc:  "Extreme Right wing political views",
		Score: -1,
	},
	"problem_solver": {
		Desc:  "Problem solver",
		Score: 1,
	},
	"capitalist": {
		Desc:  "Capitalist",
		Score: 0,
	},
	"humanitarian": {
		Desc:  "Humanitarian",
		Score: 1,
	},
	"thrillseeker": {
		Desc:  "Thrillseeker",
		Score: 1,
	},
	"explorer": {
		Desc:  "Explorer",
		Score: 1,
	},
	"hermit": {
		Desc:  "Is a Hermit",
		Score: 1,
	},
	"gender_chromosome_XX": {
		Desc:  "Has Male gender chomosomes XX",
		Score: 1,
	},
	"gender_chromosome_XY": {
		Desc:  "Has Female gender chromosomes XY",
		Score: 1,
	},
	"good_looking": {
		Desc:  "Good Looking",
		Score: 1,
	},
	"ugly": {
		Desc:  "Not great looking",
		Score: -1,
	},
	"criminal": {
		Desc:  "Criminal tendancies",
		Score: -1,
	},
	"rambles": {
		Desc:  "Tends to ramble",
		Score: -1,
	},
	"dry_humour": {
		Desc:  "Tends to ramble",
		Score: -1,
	},
}

func NewGenericHumanGenes(keys map[string]bool) chromosomes.Genes {
	genes := make(map[string]bool, 0)

	for name := range GenericHumanGenome {
		if !keys[name] {
			genes[name] = false
			continue
		}
		genes[name] = keys[name]
	}

	return genes
}
