package resizeimage

import (
	"bytes"
	"fmt"
	"image"

	"github.com/disintegration/imaging"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: ResamplingFilter = %s", s.ResamplingFilter)

	act := &Activity{settings: s} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
	settings *Settings
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	fmt.Println("HERE")
	fmt.Println(a.settings)
	var rFilter imaging.ResampleFilter
	if a.settings.ResamplingFilter == "Lanczos" {
		rFilter = imaging.Lanczos
	} else if a.settings.ResamplingFilter == "NearestNeighbor" {
		rFilter = imaging.NearestNeighbor
	} else if a.settings.ResamplingFilter == "Linear" {
		rFilter = imaging.Linear
	} else if a.settings.ResamplingFilter == "CatmullRom" {
		rFilter = imaging.CatmullRom
	} else {
		rFilter = imaging.Lanczos
	}

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	pic, _, err := image.Decode(bytes.NewReader(input.File.([]byte)))
	if err != nil {
		return false, fmt.Errorf("Error Decoding file: %v", err)
	}

	maxdim := 256
	if input.MaxDimSize > 0 {
		maxdim = input.MaxDimSize
	}

	src := pic.(image.Image)
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	if w >= h {
		w = maxdim
		h = int(w * bounds.Max.Y / bounds.Max.X)

	} else {
		h = maxdim
		w = int(h * bounds.Max.X / bounds.Max.Y)
	}
	fmt.Println(w, h)
	src = imaging.Resize(src, w, h, rFilter)

	ctx.Logger().Infof("Input: maxDim is = %d", input.MaxDimSize)

	output := &Output{ResizedImage: src}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
