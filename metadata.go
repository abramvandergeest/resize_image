package resizeimage

import (
	"image"

	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	ResamplingFilter string `md:"resamplingFilter"`
}

type Input struct {
	File       []byte `md:"file,required"`
	MaxDimSize int         `md:"maxDimSize"`
	X int         `md:"x"`
	Y int         `md:"y"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	intVal, _ := coerce.ToInt(values["maxDimSize"])
	intValx, _ := coerce.ToInt(values["x"])
	intValy, _ := coerce.ToInt(values["y"])
	r.File = values["file"].([]byte)
	r.MaxDimSize = intVal
	r.X = intValx
	r.Y = intValy
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"file":       r.File,
		"maxDimSize": r.MaxDimSize,
		"x":r.X,
		"y":r.Y,
	}
}

type Output struct {
	ResizedImage image.Image `md:"resizedImage"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	o.ResizedImage = values["resizedImage"].(image.Image)
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"resizedImage": o.ResizedImage,
	}
}
