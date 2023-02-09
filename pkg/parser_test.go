package pkg

import (
	"bytes"
	"github.com/SourceFellows/go-fidl-dbus-generator/examples"
	"testing"
)

func TestParseFidl_Notification(t *testing.T) {

	//given
	parser := NewParser(bytes.NewReader(examples.NotificationFidl))

	//when
	fidl, err := parser.Parse()

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
	parser := NewParser(bytes.NewReader(examples.SystemManagerFidl))

	//when
	fidl, err := parser.Parse()

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

func TestParseFidl_FireAndForget(t *testing.T) {

	//given
	parser := NewParser(bytes.NewReader(examples.FireAndForgetsFidl))

	//when
	fidl, err := parser.Parse()

	//then
	if err != nil {
		t.Errorf("could not parse fidl because of: %v", err)
		return
	}

	if fidl == nil {
		t.Error("fidl wasn't parsed")
		return
	}

	if len(fidl.Methods) != 1 {
		t.Errorf("wrong number of methods. expected 1 but got %d", len(fidl.Methods))
		return
	}

	if !fidl.Methods[0].FireAndForget {
		t.Error("expected fireAndForget method")
		return
	}

}

func paramOfName(fidl *Fidl, name string) Param {

	for _, tr := range fidl.Methods {
		for _, p := range tr.In {
			if p.Name == name {
				return p
			}
		}
	}

	return Param{}

}
