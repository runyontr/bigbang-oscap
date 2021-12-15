package api

import (
	"time"

	// "github.com/gocomply/oscalkit/types/oscal/component_definition"
	"github.com/gocomply/oscalkit/types/oscal/component_definition"
	"github.com/google/uuid"
)

// type Component component_definition.Component

type Component struct {
	ComponentDefinition struct {
		UUID     string `yaml:"uuid"`
		Metadata struct {
			Title        string    `yaml:"title"`
			LastModified time.Time `yaml:"last-modified"`
			Version      int       `yaml:"version"`
			OscalVersion string    `yaml:"oscal-version"`
			Roles        []struct {
				ID    string `yaml:"id"`
				Title string `yaml:"title"`
			} `yaml:"roles"`
			Parties []struct {
				UUID  string `yaml:"uuid"`
				Type  string `yaml:"type"`
				Name  string `yaml:"name"`
				Links []struct {
					Href string `yaml:"href"`
					Rel  string `yaml:"rel"`
				} `yaml:"links"`
			} `yaml:"parties"`
		} `yaml:"metadata"`
		Components []struct {
			UUID             string `yaml:"uuid"`
			Type             string `yaml:"type"`
			Title            string `yaml:"title"`
			Description      string `yaml:"description"`
			Purpose          string `yaml:"purpose"`
			ResponsibleRoles []struct {
				RoleID    string `yaml:"role-id"`
				PartyUUID string `yaml:"party-uuid"`
			} `yaml:"responsible-roles"`
			ControlImplementations []struct {
				UUID                    string `yaml:"uuid"`
				Source                  string `yaml:"source"`
				Description             string `yaml:"description"`
				ImplementedRequirements []struct {
					UUID        interface{} `yaml:"uuid"`
					ControlID   string      `yaml:"control-id"`
					Description string      `yaml:"description"`
				} `yaml:"implemented-requirements"`
			} `yaml:"control-implementations"`
		} `yaml:"components"`
	} `yaml:"component-definition"`
}

func NewComponent(name string) Component {
	c := Component{
		Uuid: uuid.New().String(),
		Metadata: {
			Title:        name,
			Version:      20211019,
			OscalVersion: "1.0.0",
		},
	}
	c.ControlImplementations = make([]component_definition.ControlImplementation, 0)
	c.Links = make([]component_definition.Link, 0)

	return c
}

func (comp *Component) AddConntrol(name, id, description string) error {

	ci := component_definition.ControlImplementation{
		Uuid:        uuid.New().String(),
		Description: description,
	}

}
