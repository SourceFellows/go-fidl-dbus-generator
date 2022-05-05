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
