package shared

import (
	"time"
)

type FullTermWithUserPosEnum string

const (
	FullTermWithUserPosEnumNoun      FullTermWithUserPosEnum = "noun"
	FullTermWithUserPosEnumVerb      FullTermWithUserPosEnum = "verb"
	FullTermWithUserPosEnumAdverb    FullTermWithUserPosEnum = "adverb"
	FullTermWithUserPosEnumAdjective FullTermWithUserPosEnum = "adjective"
)

type FullTermWithUserTypeEnum string

const (
	FullTermWithUserTypeEnumApproved FullTermWithUserTypeEnum = "approved"
	FullTermWithUserTypeEnumBanned   FullTermWithUserTypeEnum = "banned"
	FullTermWithUserTypeEnumPending  FullTermWithUserTypeEnum = "pending"
)

type FullTermWithUser struct {
	ApprovedTermExtension *ApprovedTermExtension   `json:"approvedTermExtension,omitempty"`
	BacklinkedTerms       []FullLinkedTerm         `json:"backlinkedTerms,omitempty"`
	CaseSensitive         bool                     `json:"caseSensitive"`
	CreatedUser           TerminologyUser          `json:"createdUser"`
	CreationTime          time.Time                `json:"creationTime"`
	Description           *string                  `json:"description,omitempty"`
	Examples              []TermExample            `json:"examples,omitempty"`
	Highlight             bool                     `json:"highlight"`
	ID                    int64                    `json:"id"`
	LinkedTerms           []FullLinkedTerm         `json:"linkedTerms,omitempty"`
	Mistakes              []TermMistake            `json:"mistakes,omitempty"`
	ModificationTime      time.Time                `json:"modificationTime"`
	ModifiedUser          TerminologyUser          `json:"modifiedUser"`
	Pos                   *FullTermWithUserPosEnum `json:"pos,omitempty"`
	Tags                  []TermTagResponse        `json:"tags,omitempty"`
	Term                  string                   `json:"term"`
	TermBankID            int64                    `json:"termBankId"`
	Type                  FullTermWithUserTypeEnum `json:"type"`
}
