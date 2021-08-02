package notionapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type PropertyType string

type Property interface {
	GetType() PropertyType
}

type TitleProperty struct {
	ID    PropertyID   `json:"id,omitempty"`
	Type  PropertyType `json:"type,omitempty"`
	Title []RichText   `json:"title"`
}

func (p TitleProperty) GetType() PropertyType {
	return p.Type
}

type RichTextProperty struct {
	ID       PropertyID   `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	RichText []RichText   `json:"rich_text"`
}

func (p RichTextProperty) GetType() PropertyType {
	return p.Type
}

type NumberProperty struct {
	ID     PropertyID   `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Number int          `json:"number"`
}

func (p NumberProperty) GetType() PropertyType {
	return p.Type
}

type SelectProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Select Option       `json:"select"`
}

func (p SelectProperty) GetType() PropertyType {
	return p.Type
}

type MultiSelectProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type,omitempty"`
	MultiSelect []Option     `json:"multi_select"`
}

func (p MultiSelectProperty) GetType() PropertyType {
	return p.Type
}

type Option struct {
	ID    PropertyID `json:"id,omitempty"`
	Name  string     `json:"name"`
	Color Color      `json:"color,omitempty"`
}

type DateProperty struct {
	ID   ObjectID     `json:"id,omitempty"`
	Type PropertyType `json:"type,omitempty"`
	Date Date         `json:"date"`
}

type Date struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func (p DateProperty) GetType() PropertyType {
	return p.Type
}

type FormulaProperty struct {
	ID      ObjectID     `json:"id,omitempty"`
	Type    PropertyType `json:"type,omitempty"`
	Formula Formula      `json:"formula"`
}

type FormulaType string

type Formula struct {
	Type    FormulaType `json:"type,omitempty"`
	String  string      `json:"string,omitempty"`
	Number  int         `json:"number,omitempty"`
	Boolean bool        `json:"boolean,omitempty"`
	Date    *Date       `json:"date,omitempty"`
}

func (p FormulaProperty) GetType() PropertyType {
	return p.Type
}

type RelationProperty struct {
	ID       ObjectID     `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	Relation []Relation   `json:"relation"`
}

type Relation struct {
	ID PageID `json:"id"`
}

func (p RelationProperty) GetType() PropertyType {
	return p.Type
}

type RollupProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Rollup Rollup       `json:"rollup"`
}

type RollupType string

type Rollup struct {
	Type   RollupType `json:"type,omitempty"`
	Number int        `json:"number,omitempty"`
	Date   *Date      `json:"date,omitempty"`
	Array  []Property `json:"array,omitempty"` //todo: unmarshal
}

func (p RollupProperty) GetType() PropertyType {
	return p.Type
}

type PeopleProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	People []User       `json:"people"`
}

func (p PeopleProperty) GetType() PropertyType {
	return p.Type
}

type FilesProperty struct {
	ID    ObjectID     `json:"id,omitempty"`
	Type  PropertyType `json:"type,omitempty"`
	Files []File       `json:"files"`
}

func (p FilesProperty) GetType() PropertyType {
	return p.Type
}

type CheckboxProperty struct {
	ID       ObjectID     `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	Checkbox bool         `json:"checkbox"`
}

func (p CheckboxProperty) GetType() PropertyType {
	return p.Type
}

type URLProperty struct {
	ID   ObjectID     `json:"id,omitempty"`
	Type PropertyType `json:"type,omitempty"`
	URL  string       `json:"url"`
}

func (p URLProperty) GetType() PropertyType {
	return p.Type
}

type EmailProperty struct {
	ID    PropertyID   `json:"id,omitempty"`
	Type  PropertyType `json:"type,omitempty"`
	Email string       `json:"email"`
}

func (p EmailProperty) GetType() PropertyType {
	return p.Type
}

type PhoneNumberProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type,omitempty"`
	PhoneNumber string       `json:"phone_number"`
}

func (p PhoneNumberProperty) GetType() PropertyType {
	return p.Type
}

type CreatedTimeProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type,omitempty"`
	CreatedTime time.Time    `json:"created_time"`
}

func (p CreatedTimeProperty) GetType() PropertyType {
	return p.Type
}

type CreatedByProperty struct {
	ID        ObjectID     `json:"id,omitempty"`
	Type      PropertyType `json:"type,omitempty"`
	CreatedBy User         `json:"created_by"`
}

func (p CreatedByProperty) GetType() PropertyType {
	return p.Type
}

type LastEditedTimeProperty struct {
	ID             ObjectID     `json:"id,omitempty"`
	Type           PropertyType `json:"type,omitempty"`
	LastEditedTime time.Time    `json:"last_edited_time"`
}

func (p LastEditedTimeProperty) GetType() PropertyType {
	return p.Type
}

type LastEditedByProperty struct {
	ID           ObjectID     `json:"id,omitempty"`
	Type         PropertyType `json:"type,omitempty"`
	LastEditedBy User         `json:"last_edited_by"`
}

func (p LastEditedByProperty) GetType() PropertyType {
	return p.Type
}

type Properties map[string]Property

func (p *Properties) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	props, err := parsePageProperties(raw)
	if err != nil {
		return err
	}

	*p = props
	return nil
}

func parsePageProperties(raw map[string]interface{}) (map[string]Property, error) {
	result := make(map[string]Property)
	for k, v := range raw {
		var p Property
		switch rawProperty := v.(type) {
		case map[string]interface{}:
			switch PropertyType(rawProperty["type"].(string)) {
			case PropertyTypeTitle:
				p = &TitleProperty{}
			case PropertyTypeRichText:
				p = &RichTextProperty{}
			case PropertyTypeNumber:
				p = &NumberProperty{}
			case PropertyTypeSelect:
				p = &SelectProperty{}
			case PropertyTypeMultiSelect:
				p = &MultiSelectProperty{}
			case PropertyTypeDate:
				p = &DateProperty{}
			case PropertyTypeFormula:
				p = &FormulaProperty{}
			case PropertyTypeRelation:
				p = &RelationProperty{}
			case PropertyTypeRollup:
				p = &RollupProperty{}
			case PropertyTypePeople:
				p = &PeopleProperty{}
			case PropertyTypeFiles:
				p = &FilesProperty{}
			case PropertyTypeCheckbox:
				p = &CheckboxProperty{}
			case PropertyTypeURL:
				p = &URLProperty{}
			case PropertyTypeEmail:
				p = &EmailProperty{}
			case PropertyTypePhoneNumber:
				p = PhoneNumberProperty{}
			case PropertyTypeCreatedTime:
				p = &CreatedTimeProperty{}
			case PropertyTypeCreatedBy:
				p = &CreatedByProperty{}
			case PropertyTypeLastEditedTime:
				p = &LastEditedTimeProperty{}
			case PropertyTypeLastEditedBy:
				p = &LastEditedByProperty{}
			default:
				return nil, errors.New(fmt.Sprintf("unsupported property type: %s", rawProperty["type"].(string)))
			}
			b, err := json.Marshal(rawProperty)
			if err != nil {
				return nil, err
			}

			if err = json.Unmarshal(b, &p); err != nil {
				return nil, err
			}

			result[k] = p
		default:
			return nil, errors.New(fmt.Sprintf("unsupported property format %T", v))
		}
	}

	return result, nil
}
