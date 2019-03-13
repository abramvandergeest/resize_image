package resizeImage

import (
	"image"
	"mime/multipart"

	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	ResamplingFilter string `md:"resamplingFilter"`
}

type Input struct {
	File       multipart.File `md:"file,required"`
	MaxDimSize int            `md:"maxDimSize"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	intVal, _ := coerce.ToInt(values["maxDimSize"])
	r.File = values["file"].(multipart.File)
	r.MaxDimSize = intVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"file":       r.File,
		"maxDimSize": r.MaxDimSize,
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
