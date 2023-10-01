package noteshandler

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetNoteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}
