package fetcher

import (
	"encoding/json"
	"fmt"
)

type SiteMap struct {
	Id string `json:"_id"`
	StartUrl []string `json:"startUrl"`
	Selectors []Selector `json:"selectors"`
}

func (s *SiteMap) UnmarshalJSON(data []byte) error {
	type Alias SiteMap
	alias := &struct {
		Selectors []json.RawMessage `json:"selectors"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	for _, rawSelector := range alias.Selectors {
		var base selectorBase
		if err := json.Unmarshal(rawSelector, &base); err != nil {
			return err
		}
		switch base.Type {
		case Element:
			var selector SelectorElement
			if err := json.Unmarshal(rawSelector, &selector); err != nil {
				return err
			}
			s.Selectors = append(s.Selectors, &selector)
		case ElementAttr:
			var selector SelectorElementAttr
			if err := json.Unmarshal(rawSelector, &selector); err != nil {
				return err
			}
			s.Selectors = append(s.Selectors, &selector)
		case Text:
			var selector SelectorText
			if err := json.Unmarshal(rawSelector, &selector); err != nil {
				return err
			}
			s.Selectors = append(s.Selectors, &selector)
		case Link:
			var selector SelectorLink
			if err := json.Unmarshal(rawSelector, &selector); err != nil {
				return err
			}
			s.Selectors = append(s.Selectors, &selector)
		case Image:
			var selector SelectorImage
			if err := json.Unmarshal(rawSelector, &selector); err != nil {
				return err
			}
			s.Selectors = append(s.Selectors, &selector)
		default:
			return fmt.Errorf("unknown type: %q", base.Type)
		}
	}
	return nil
}

type SelectorType string

const (
	Element = SelectorType("SelectorElement")
	ElementAttr = SelectorType("SelectorElementAttribute")
	Text = SelectorType("SelectorText")
	Link = SelectorType("SelectorLink")
	Image = SelectorType("SelectorImage")
)

type SelectorId string

type Selector interface {
	GetId() SelectorId
	GetType() SelectorType
	GetParentSelectors() []SelectorId
	GetSelector() string
	IsMultiple() bool
	GetDelay () int
}

type selectorBase struct {
	Id string `json:"id"`
	Type SelectorType `json:"type"`
	ParentSelectors []SelectorId `json:"parentSelectors"`
	Selector string `json:"selector"`
	Multiple bool `json:"multiple"`
	Delay int `json:"delay"`
}

func (s *selectorBase) GetId() SelectorId {
	return SelectorId(s.Id)
}

func (s *selectorBase) GetType() SelectorType {
	return s.Type
}

func (s *selectorBase) GetParentSelectors() []SelectorId {
	return s.ParentSelectors
}

func (s *selectorBase) GetSelector() string {
	return s.Selector
}

func (s *selectorBase) IsMultiple() bool {
	return s.Multiple
}

func (s *selectorBase) GetDelay() int {
	return s.Delay
}

type SelectorElement struct {
	selectorBase
}

type SelectorElementAttr struct {
	selectorBase
	ExtractAttribute string `json:"extractAttribute"`
}

type SelectorText struct {
	selectorBase
	Regex string `json:"regex"`
}

type SelectorLink struct {
	selectorBase
}

type SelectorImage struct {
	selectorBase
}
