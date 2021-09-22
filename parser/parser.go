package parser

import (
	"bufio"
	"errors"
	"io"
	"net/http"

	"github.com/wahyuhadi/gorace/models"
)

func ReadHTTPFromFile(r io.Reader) (*models.Connection, error) {
	buf := bufio.NewReader(r)
	req, err := http.ReadRequest(buf)
	if err == io.EOF {
		return nil, errors.New("error read in file request EOF")
	}

	if err != nil {
		return nil, errors.New("error read in file request")
	}

	return &models.Connection{Request: req, Response: nil}, nil

}
