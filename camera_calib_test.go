package gogige

import (
	"path/filepath"
	"testing"
)

func TestWriteCalibFileRequiresConnection(t *testing.T) {
	path := filepath.Join(t.TempDir(), "calib.json")

	var nilCam *Camera
	if _, err := nilCam.WriteCalibFile(path); err == nil {
		t.Fatal("nil camera must report a connection error")
	}
	if _, err := (&Camera{}).WriteCalibFile(path); err == nil {
		t.Fatal("camera without GVCP must report a connection error")
	}

	closed := &device{}
	if _, err := closed.WriteCalibFile(path); err == nil {
		t.Fatal("closed device must report an error, not write a file")
	}
}
