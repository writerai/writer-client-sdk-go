package shared

type ContentIssueServiceEnum string

const (
	ContentIssueServiceEnumCommonMistakes     ContentIssueServiceEnum = "common-mistakes"
	ContentIssueServiceEnumBannedWords        ContentIssueServiceEnum = "banned-words"
	ContentIssueServiceEnumDictionary         ContentIssueServiceEnum = "dictionary"
	ContentIssueServiceEnumGec                ContentIssueServiceEnum = "gec"
	ContentIssueServiceEnumInfinitive         ContentIssueServiceEnum = "infinitive"
	ContentIssueServiceEnumSpelling           ContentIssueServiceEnum = "spelling"
	ContentIssueServiceEnumWritingStyle       ContentIssueServiceEnum = "writing-style"
	ContentIssueServiceEnumCustomRules        ContentIssueServiceEnum = "custom-rules"
	ContentIssueServiceEnumSentenceCase       ContentIssueServiceEnum = "sentence-case"
	ContentIssueServiceEnumAcronym            ContentIssueServiceEnum = "acronym"
	ContentIssueServiceEnumOxfordComma        ContentIssueServiceEnum = "oxford-comma"
	ContentIssueServiceEnumMlPunctuation      ContentIssueServiceEnum = "ml-punctuation"
	ContentIssueServiceEnumEmojis             ContentIssueServiceEnum = "emojis"
	ContentIssueServiceEnumGenderPronouns     ContentIssueServiceEnum = "gender-pronouns"
	ContentIssueServiceEnumSensitivity        ContentIssueServiceEnum = "sensitivity"
	ContentIssueServiceEnumPlagiarism         ContentIssueServiceEnum = "plagiarism"
	ContentIssueServiceEnumReadability        ContentIssueServiceEnum = "readability"
	ContentIssueServiceEnumSentenceComplexity ContentIssueServiceEnum = "sentence-complexity"
	ContentIssueServiceEnumVocabulary         ContentIssueServiceEnum = "vocabulary"
	ContentIssueServiceEnumParagraphLength    ContentIssueServiceEnum = "paragraph-length"
	ContentIssueServiceEnumPlainLanguage      ContentIssueServiceEnum = "plain-language"
	ContentIssueServiceEnumHealthyCommn       ContentIssueServiceEnum = "healthy-commn"
	ContentIssueServiceEnumConfidence         ContentIssueServiceEnum = "confidence"
	ContentIssueServiceEnumDataLossPrevention ContentIssueServiceEnum = "data-loss-prevention"
	ContentIssueServiceEnumHateSpeech         ContentIssueServiceEnum = "hate-speech"
	ContentIssueServiceEnumContentSafeguards  ContentIssueServiceEnum = "content-safeguards"
	ContentIssueServiceEnumFeedback           ContentIssueServiceEnum = "feedback"
	ContentIssueServiceEnumClaim              ContentIssueServiceEnum = "claim"
	ContentIssueServiceEnumQuote              ContentIssueServiceEnum = "quote"
	ContentIssueServiceEnumGenderNouns        ContentIssueServiceEnum = "gender-nouns"
	ContentIssueServiceEnumGenderTone         ContentIssueServiceEnum = "gender-tone"
	ContentIssueServiceEnumGrammar            ContentIssueServiceEnum = "grammar"
	ContentIssueServiceEnumPunctuationDark    ContentIssueServiceEnum = "punctuation-dark"
	ContentIssueServiceEnumFormatting         ContentIssueServiceEnum = "formatting"
	ContentIssueServiceEnumTwitter            ContentIssueServiceEnum = "twitter"
	ContentIssueServiceEnumGecDark            ContentIssueServiceEnum = "gec-dark"
	ContentIssueServiceEnumGecGpt3            ContentIssueServiceEnum = "gec-gpt3"
)

type ContentIssue struct {
	Description *string                 `json:"description,omitempty"`
	From        int64                   `json:"from"`
	Meta        interface{}             `json:"meta,omitempty"`
	Service     ContentIssueServiceEnum `json:"service"`
	Suggestions []string                `json:"suggestions,omitempty"`
	Until       int64                   `json:"until"`
}
