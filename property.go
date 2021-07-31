package notionapi

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type PropertyType string

type Property interface {
	GetType() PropertyType
}

type PropertyID string

func (pID PropertyID) String() string {
	return string(pID)
}

type TextProperty struct {
	ID    PropertyID   `json:"id,omitempty"`
	Type  PropertyType `json:"type"`
	Title []RichText   `json:"title"`
}

func (p TextProperty) GetType() PropertyType {
	return p.Type
}

type EmptyRichTextProperty struct {
	ID       PropertyID   `json:"id,omitempty"`
	Type     PropertyType `json:"type"`
	RichText struct{}     `json:"rich_text"`
}

func (p EmptyRichTextProperty) GetType() PropertyType {
	return p.Type
}

type RichTextProperty struct {
	ID       PropertyID   `json:"id,omitempty"`
	Type     PropertyType `json:"type"`
	RichText []RichText   `json:"rich_text"`
}

func (p RichTextProperty) GetType() PropertyType {
	return p.Type
}

type DatabaseTitleProperty struct {
	ID    PropertyID   `json:"id,omitempty"`
	Type  PropertyType `json:"type"`
	Title RichText     `json:"title"`
}

func (p DatabaseTitleProperty) GetType() PropertyType {
	return p.Type
}

type PageTitleProperty struct {
	ID    PropertyID   `json:"id,omitempty"`
	Type  PropertyType `json:"type,omitempty"`
	Title Paragraph    `json:"title"`
}

func (p PageTitleProperty) GetType() PropertyType {
	return p.Type
}

type FormatType string

func (ft FormatType) String() string {
	return string(ft)
}

type NumberProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type"`
	Format FormatType   `json:"format"`
}

func (p NumberProperty) GetType() PropertyType {
	return p.Type
}

type DatabaseSelectProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type"`
	Select Select       `json:"select"`
}

func (p DatabaseSelectProperty) GetType() PropertyType {
	return p.Type
}

type PageSelectProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type"`
	Select Option       `json:"select"`
}

func (p PageSelectProperty) GetType() PropertyType {
	return p.Type
}

type Select struct {
	Options []Option `json:"options"`
}

type DatabaseMultiSelectProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type"`
	MultiSelect Select       `json:"multi_select"`
}

func (p DatabaseMultiSelectProperty) GetType() PropertyType {
	return p.Type
}

type PageMultiSelectProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type"`
	MultiSelect []Option     `json:"multi_select"`
}

func (p PageMultiSelectProperty) GetType() PropertyType {
	return p.Type
}

type Option struct {
	ID    PropertyID `json:"id,omitempty"`
	Name  string     `json:"name"`
	Color Color      `json:"color,omitempty"`
}

type DatabaseDateProperty struct {
	ID   ObjectID     `json:"id,omitempty"`
	Type PropertyType `json:"type"`
	Date struct{}     `json:"date"`
}

func (p DatabaseDateProperty) GetType() PropertyType {
	return p.Type
}

type PageDateProperty struct {
	ID   ObjectID     `json:"id,omitempty"`
	Type PropertyType `json:"type"`
	Date Date         `json:"date"`
}

func (p PageDateProperty) GetType() PropertyType {
	return p.Type
}

type DatabasePeopleProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type"`
	People struct{}     `json:"people"`
}

func (p DatabasePeopleProperty) GetType() PropertyType {
	return p.Type
}

//TODO: is this a right property format?
type PagePeopleProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type"`
	People []User       `json:"people"`
}

func (p PagePeopleProperty) GetType() PropertyType {
	return p.Type
}

type DatabaseFilesProperty struct {
	ID    ObjectID     `json:"id,omitempty"`
	Type  PropertyType `json:"type"`
	Files struct{}     `json:"files"`
}

func (p DatabaseFilesProperty) GetType() PropertyType {
	return p.Type
}

type PageFilesProperty struct {
	ID    ObjectID     `json:"id,omitempty"`
	Type  PropertyType `json:"type"`
	Files []File       `json:"files"`
}

func (p PageFilesProperty) GetType() PropertyType {
	return p.Type
}

type DatabaseCheckboxProperty struct {
	ID       ObjectID     `json:"id,omitempty"`
	Type     PropertyType `json:"type"`
	Checkbox struct{}     `json:"checkbox"`
}

func (p DatabaseCheckboxProperty) GetType() PropertyType {
	return p.Type
}

type PageCheckboxProperty struct {
	ID       ObjectID     `json:"id,omitempty"`
	Type     PropertyType `json:"type"`
	Checkbox bool         `json:"checkbox"`
}

func (p PageCheckboxProperty) GetType() PropertyType {
	return p.Type
}

type DatabaseURLProperty struct {
	ID   ObjectID     `json:"id,omitempty"`
	Type PropertyType `json:"type"`
	URL  struct{}     `json:"url"`
}

func (p DatabaseURLProperty) GetType() PropertyType {
	return p.Type
}

type PageURLProperty struct {
	ID   ObjectID     `json:"id,omitempty"`
	Type PropertyType `json:"type"`
	URL  string       `json:"url"`
}

func (p PageURLProperty) GetType() PropertyType {
	return p.Type
}

type DatabaseEmailProperty struct {
	ID    PropertyID   `json:"id,omitempty"`
	Type  PropertyType `json:"type"`
	Email struct{}     `json:"email"`
}

func (p DatabaseEmailProperty) GetType() PropertyType {
	return p.Type
}

type PageEmailProperty struct {
	ID    PropertyID   `json:"id,omitempty"`
	Type  PropertyType `json:"type"`
	Email string       `json:"email"`
}

func (p PageEmailProperty) GetType() PropertyType {
	return p.Type
}

type DatabasePhoneNumberProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type"`
	PhoneNumber struct{}     `json:"phone_number"`
}

func (p DatabasePhoneNumberProperty) GetType() PropertyType {
	return p.Type
}

type PagePhoneNumberProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type"`
	PhoneNumber string       `json:"phone_number"`
}

func (p PagePhoneNumberProperty) GetType() PropertyType {
	return p.Type
}

type FormulaProperty struct {
	ID         ObjectID     `json:"id,omitempty"`
	Type       PropertyType `json:"type"`
	Expression string       `json:"expression"`
}

func (p FormulaProperty) GetType() PropertyType {
	return p.Type
}

type RelationProperty struct {
	Type     PropertyType `json:"type"`
	Relation Relation     `json:"relation"`
}

type Relation struct {
	DatabaseID         DatabaseID `json:"database_id"`
	SyncedPropertyID   PropertyID `json:"synced_property_id"`
	SyncedPropertyName string     `json:"synced_property_name"`
}

func (p RelationProperty) GetType() PropertyType {
	return p.Type
}

type RollupProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type"`
	Rollup Rollup       `json:"rollup"`
}

type Rollup struct {
	RelationPropertyName string       `json:"relation_property_name"`
	RelationPropertyID   PropertyID   `json:"relation_property_id"`
	RollupPropertyName   string       `json:"rollup_property_name"`
	RollupPropertyID     PropertyID   `json:"rollup_property_id"`
	Function             FunctionType `json:"function"`
}

func (p RollupProperty) GetType() PropertyType {
	return p.Type
}

type CreatedTimeProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type"`
	CreatedTime interface{}  `json:"created_time"`
}

func (p CreatedTimeProperty) GetType() PropertyType {
	return p.Type
}

type CreatedByProperty struct {
	ID        ObjectID     `json:"id"`
	Type      PropertyType `json:"type"`
	CreatedBy interface{}  `json:"created_by"`
}

func (p CreatedByProperty) GetType() PropertyType {
	return p.Type
}

type LastEditedTimeProperty struct {
	ID             ObjectID     `json:"id"`
	Type           PropertyType `json:"type"`
	LastEditedTime interface{}  `json:"last_edited_time"`
}

func (p LastEditedTimeProperty) GetType() PropertyType {
	return p.Type
}

type LastEditedByProperty struct {
	ID           ObjectID     `json:"id"`
	Type         PropertyType `json:"type"`
	LastEditedBy interface{}  `json:"last_edited_by"`
}

func (p LastEditedByProperty) GetType() PropertyType {
	return p.Type
}

type DatabaseProperties map[string]Property

func (p *DatabaseProperties) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	props, err := parseDatabaseProperties(raw)
	if err != nil {
		return err
	}

	*p = props
	return nil
}

func parseDatabaseProperties(raw map[string]interface{}) (map[string]Property, error) {
	result := make(map[string]Property)
	for k, v := range raw {
		var p Property
		switch rawProperty := v.(type) {
		case map[string]interface{}:
			switch PropertyType(rawProperty["type"].(string)) {
			case PropertyTypeTitle:
				p = &DatabaseTitleProperty{}
			case PropertyTypeRichText:
				switch v.(map[string]interface{})["rich_text"].(type) {
				case map[string]interface{}:
					p = &EmptyRichTextProperty{}
				default:
					p = &RichTextProperty{}
				}
			case PropertyTypeSelect:
				p = &DatabaseSelectProperty{}
			case PropertyTypeMultiSelect:
				p = &DatabaseMultiSelectProperty{}
			case PropertyTypeNumber:
				p = &NumberProperty{}
			case PropertyTypeCheckbox:
				p = &DatabaseCheckboxProperty{}
			case PropertyTypeEmail:
				p = &DatabaseEmailProperty{}
			case PropertyTypeURL:
				p = &DatabaseURLProperty{}
			case PropertyTypeFile:
				p = &DatabaseFilesProperty{}
			case PropertyTypePhoneNumber:
				p = DatabasePhoneNumberProperty{}
			case PropertyTypeFormula:
				p = &FormulaProperty{}
			case PropertyTypeDate:
				p = &DatabaseDateProperty{}
			case PropertyTypeRelation:
				p = &RelationProperty{}
			case PropertyTypeRollup:
				p = &RollupProperty{}
			case PropertyTypePeople:
				p = &DatabasePeopleProperty{}
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

type PageProperties map[string]Property

func (p *PageProperties) UnmarshalJSON(data []byte) error {
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
				p = &PageTitleProperty{}
			case PropertyTypeRichText:
				switch v.(map[string]interface{})["rich_text"].(type) {
				case map[string]interface{}:
					p = &EmptyRichTextProperty{}
				default:
					p = &RichTextProperty{}
				}
			case PropertyTypeSelect:
				p = &PageSelectProperty{}
			case PropertyTypeMultiSelect:
				p = &PageMultiSelectProperty{}
			case PropertyTypeNumber:
				p = &NumberProperty{}
			case PropertyTypeCheckbox:
				p = &PageCheckboxProperty{}
			case PropertyTypeEmail:
				p = &PageEmailProperty{}
			case PropertyTypeURL:
				p = &PageURLProperty{}
			case PropertyTypeFile:
				p = &PageFilesProperty{}
			case PropertyTypePhoneNumber:
				p = PagePhoneNumberProperty{}
			case PropertyTypeFormula:
				p = &FormulaProperty{}
			case PropertyTypeDate:
				p = &PageDateProperty{}
			case PropertyTypeRelation:
				p = &RelationProperty{}
			case PropertyTypeRollup:
				p = &RollupProperty{}
			case PropertyTypePeople:
				p = &PagePeopleProperty{}
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
