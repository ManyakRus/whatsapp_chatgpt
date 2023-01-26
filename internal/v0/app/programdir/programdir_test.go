package programdir

import "testing"

func TestCurrentFilename(t *testing.T) {
	s := CurrentFilename()
	if s == "" {
		t.Error("programdir_test.TestCurrentFilename() error")
	}

}

func TestProgramDir(t *testing.T) {
	s := ProgramDir()
	if s == "" {
		t.Error("programdir_test.TestCurrentFilename() error")
	}
}
