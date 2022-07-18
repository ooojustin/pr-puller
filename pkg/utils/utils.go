package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func GetResponseBody(r *http.Response) (string, bool) {
	defer r.Body.Close()

	rbytes, err := io.ReadAll(r.Body)
	if err != nil {
		return "", false
	}

	return string(rbytes), true
}

func FindHiddenValue(key string, pageSource string) (string, bool) {
	pattern := fmt.Sprintf("<input type=\"hidden\" name=\"%s\" value=\"(.+?)\"", key)
	exp := regexp.MustCompile(pattern)

	fss := exp.FindStringSubmatch(pageSource)
	if len(fss) != 2 {
		// Expect the whole match (1) and the grouped value (2)
		return "", false
	}

	return fss[1], true
}

func GetAttribute(node *html.Node, key string) (string, bool) {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func AddFormField(w *multipart.Writer, key string, value string) error {
	writer, err := w.CreateFormField(key)
	if err != nil {
		return err
	}
	reader := strings.NewReader(value)
	_, err = io.Copy(writer, reader)
	return err
}
