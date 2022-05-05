package pkg

import (
	"github.com/SourceFellows/go-fidl-dbus-generator/examples"
	"testing"
)

func TestParseFidl_Notofication(t *testing.T) {

	//given

	//when
	fidl, err := ParseFidl(examples.NotificationFidl)

	//then
	if err != nil {
		t.Errorf("could not parse fidl because of: %v", err)
		return
	}

	if fidl == nil {
		t.Error("fidl wasn't parsed")
		return
	}

	actionsParam := paramOfName(fidl, "actions")
	if !actionsParam.IsArray {
		t.Error("actions should be an array")
		return
	}

}

func TestParseFidl_SystemManager(t *testing.T) {

	//given

	//when
	fidl, err := ParseFidl(examples.SystemManagerFidl)

	//then
	if err != nil {
		t.Errorf("could not parse fidl because of: %v", err)
		return
	}

	if fidl == nil {
		t.Error("fidl wasn't parsed")
		return
	}

}

func paramOfName(fidl *Fidl, name string) *Param {

	for _, tr := range fidl.Entry.TypeRef {
		for _, p := range tr.Method.Params.InParams {
			if p.Name == name {
				return p
			}
		}
	}

	return nil

}
