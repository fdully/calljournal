package templates

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/fdully/calljournal/bindata/tmpl"
	"github.com/fdully/calljournal/internal/logging"
	"github.com/fdully/calljournal/internal/util"
	"github.com/oxtoacart/bpool"
)

var (
	bufPool *bpool.BufferPool
	store   map[string]*template.Template
)

const (
	bufNumber = 48
)

func Init(ctx context.Context) {
	logger := logging.FromContext(ctx)

	bufPool = bpool.NewBufferPool(bufNumber)
	store = make(map[string]*template.Template)

	if err := parseAndStore(); err != nil {
		logger.Fatalf("failed to parse templates: %w", err)
	}
}

func Check(name string, data interface{}) error {
	if bufPool == nil {
		return fmt.Errorf("failed bufpool: %w", util.ErrIsNil)
	}

	t, ok := store[name]
	if !ok {
		return fmt.Errorf("failed template %s: %w", name, util.ErrNotExist)
	}

	buf := bufPool.Get()
	defer bufPool.Put(buf)

	err := t.Execute(buf, data)
	if err != nil {
		return fmt.Errorf("failed to check execute template %s: %w", name, err)
	}

	return nil
}

func Execute(w io.Writer, name string, data interface{}) error {
	if store == nil {
		return fmt.Errorf("failed store: %w", util.ErrIsNil)
	}

	t, ok := store[name]
	if !ok {
		return fmt.Errorf("failed template %s: %w", name, util.ErrNotExist)
	}

	return t.Execute(w, data)
}

func parseAndStore() error {
	if store == nil {
		return fmt.Errorf("failed store: %w", util.ErrIsNil)
	}

	// loop through templates pages
	for _, name := range tmpl.AssetNames() {
		if !strings.Contains(name, "page") {
			continue
		}

		b, err := tmpl.Asset(name)
		if err != nil {
			return fmt.Errorf("failed to asset %s: %w", name, err)
		}

		t := template.New(name)

		t, err = t.Parse(string(b))
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", name, err)
		}

		// loop through templates layout and includes, not pages
		for _, other := range tmpl.AssetNames() {
			if strings.Contains(other, "page") {
				continue
			}

			b, err := tmpl.Asset(other)
			if err != nil {
				return fmt.Errorf("failed to asset %s: %w", other, err)
			}

			t, err = t.Parse(string(b))
			if err != nil {
				return fmt.Errorf("failed to parse template %s: %w", other, err)
			}
		}

		// store each template page with name, for instance "debug.page.gohtml"
		store[name] = t
	}

	return nil
}
