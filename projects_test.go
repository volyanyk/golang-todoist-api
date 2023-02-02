package golang_todoist_api

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestGetProjects(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/projects", getProjects)
	expectedProjects := getTestProjects()

	once.Do(startServer)
	api := New("testing-token", OptionAPIURL("http://"+serverAddr+"/"))

	projects, err := api.GetProjects()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expectedProjects, *projects) {
		t.Fatal(ErrIncorrectResponse)
	}
}
func TestGetProjectById(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	var id = "1"
	http.HandleFunc("/projects/"+id, getProjectById(id))
	expectedProject := getTestProjectWithId(id)

	once.Do(startServer)
	api := New("testing-token", OptionAPIURL("http://"+serverAddr+"/"))

	project, err := api.GetProjectById(id)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expectedProject, *project) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func getTestProjects() []Project {
	return []Project{
		getTestProjectWithId("12345"),
		getTestProjectWithId("23456"),
	}
}

func getTestProjectWithId(id string) Project {
	return Project{
		ID:             id,
		ParentId:       "",
		Order:          0,
		Color:          "",
		Name:           "",
		CommentCount:   0,
		IsShared:       false,
		IsFavorite:     false,
		IsInboxProject: false,
		IsTeamInbox:    false,
		Url:            "",
		ViewStyle:      "",
	}
}

func getProjects(rw http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(
		getTestProjects(),
	)
	_, err := rw.Write(response)
	if err != nil {
		return
	}
}
func getProjectById(id string) func(rw http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(
		getTestProjectWithId(id),
	)

	return func(rw http.ResponseWriter, r *http.Request) {
		_, err := rw.Write(response)
		if err != nil {
			return
		}
	}
}
