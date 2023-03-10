package shared

type ContentSettings struct {
	AgeAndFamilyStatus                  bool `json:"ageAndFamilyStatus"`
	Confidence                          bool `json:"confidence"`
	ContentSafeguards                   bool `json:"contentSafeguards"`
	Disability                          bool `json:"disability"`
	GenderIdentitySensitivity           bool `json:"genderIdentitySensitivity"`
	GenderInclusiveNouns                bool `json:"genderInclusiveNouns"`
	GenderInclusivePronouns             bool `json:"genderInclusivePronouns"`
	Grammar                             bool `json:"grammar"`
	HealthyCommunication                bool `json:"healthyCommunication"`
	PassiveVoice                        bool `json:"passiveVoice"`
	RaceEthnicityNationalitySensitivity bool `json:"raceEthnicityNationalitySensitivity"`
	SexualOrientationSensitivity        bool `json:"sexualOrientationSensitivity"`
	Spelling                            bool `json:"spelling"`
	SubstanceUseSensitivity             bool `json:"substanceUseSensitivity"`
	UnclearReference                    bool `json:"unclearReference"`
	Wordiness                           bool `json:"wordiness"`
}
