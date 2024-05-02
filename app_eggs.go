package crocgodyl

import (
	"encoding/json"
	"fmt"
	"time"
)

type Egg struct {
	Id           int               `json:"id"`
	Uuid         string            `json:"uuid"`
	Name         string            `json:"name"`
	Nest         int               `json:"nest"`
	Author       string            `json:"author"`
	Description  string            `json:"description"`
	DockerImage  string            `json:"docker_image"`
	DockerImages map[string]string `json:"docker_images"`
	Config       *struct {
		Files   map[string]interface{} `json:"files"`
		Startup struct {
			Done            string   `json:"done"`
			UserInteraction []string `json:"userInteraction"`
		} `json:"startup"`
		Stop         string        `json:"stop"`
		Logs         interface{}   `json:"logs"`
		FileDenylist []interface{} `json:"file_denylist"`
		Extends      interface{}   `json:"extends"`
	} `json:"config"`
	Startup string `json:"startup"`
	Script  struct {
		Privileged bool        `json:"privileged"`
		Install    string      `json:"install"`
		Entry      string      `json:"entry"`
		Container  string      `json:"container"`
		Extends    interface{} `json:"extends"`
	} `json:"script"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Relationships struct {
		//Nest struct {
		//	Object     string `json:"object"`
		//	Attributes struct {
		//		Id          int       `json:"id"`
		//		Uuid        string    `json:"uuid"`
		//		Author      string    `json:"author"`
		//		Name        string    `json:"name"`
		//		Description string    `json:"description"`
		//		CreatedAt   time.Time `json:"created_at"`
		//		UpdatedAt   time.Time `json:"updated_at"`
		//	} `json:"attributes"`
		//} `json:"nest"`
		//Servers struct {
		//	Object string        `json:"object"`
		//	Data   []interface{} `json:"data"`
		//} `json:"servers"`
		Variables struct {
			Object string `json:"object"`
			Data   []struct {
				Object     string `json:"object"`
				Attributes struct {
					Id           int       `json:"id"`
					EggId        int       `json:"egg_id"`
					Name         string    `json:"name"`
					Description  string    `json:"description"`
					EnvVariable  string    `json:"env_variable"`
					DefaultValue string    `json:"default_value"`
					UserViewable bool      `json:"user_viewable"`
					UserEditable bool      `json:"user_editable"`
					Rules        string    `json:"rules"`
					CreatedAt    time.Time `json:"created_at"`
					UpdatedAt    time.Time `json:"updated_at"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"variables"`
	} `json:"relationships"`
}

func (a *Application) GetEgg(nestId int, eggId int) (*Egg, error) {
	req := a.newRequest("GET", fmt.Sprintf("/nests/%d/eggs/%d?include=variables", nestId, eggId), nil)
	res, err := a.Http.Do(req)
	if err != nil {
		return nil, err
	}

	buf, err := validate(res)
	if err != nil {
		return nil, err
	}

	var model struct {
		Attributes Egg `json:"attributes"`
	}
	if err = json.Unmarshal(buf, &model); err != nil {
		return nil, err
	}

	return &model.Attributes, nil
}
