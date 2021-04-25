package parser

import (
	"errors"
	"io"
	"io/ioutil"

	"github.com/gihyeonsung/pott/internal/model"
)

func parse(r io.Reader) (*model.TextDocument, error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.New("could not read the text document: " + err.Error())
	}
	return model.NewTextDocument(string(content)), nil
}
