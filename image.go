package placemat

import (
	"context"
	"errors"
	"net/url"
)

// ImageSpec represents an Image specification in YAML.
type ImageSpec struct {
	Name              string `yaml:"name"`
	URL               string `yaml:"url,omitempty"`
	File              string `yaml:"file,omitempty"`
	CompressionMethod string `yaml:"compression,omitempty"`
}

// Image represents an image configuration
type Image struct {
	*ImageSpec
	u      *url.URL
	decomp Decompressor
	p      string
}

func NewImage(spec *ImageSpec) (*Image, error) {
	if len(spec.Name) == 0 {
		return nil, errors.New("invalid image spec: " + spec.Name)
	}

	if len(spec.URL) == 0 && len(spec.File) == 0 {
		return nil, errors.New("invalid image spec: " + spec.Name)
	}

	i := &Image{ImageSpec: spec}

	if len(spec.URL) > 0 {
		if len(spec.File) > 0 {
			return nil, errors.New("invalid image spec: " + spec.Name)
		}
		u, err := url.Parse(spec.URL)
		if err != nil {
			return nil, err
		}
		i.u = u
	}

	decomp, err := NewDecompressor(spec.CompressionMethod)
	if err != nil {
		return nil, err
	}
	i.decomp = decomp

	return i, nil
}

func (i *Image) Prepare(ctx context.Context, c *cache) error {
	err := downloadData(ctx, i.u, i.decomp, c)
	if err != nil {
		return err
	}

	i.p = c.Path(i.u.String())
	return nil
}

func (i *Image) Path() string {
	return i.p
}
