package resizeImage

import (
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

	ctx.Logger().Debugf("Setting: %s", s.ASetting)

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	ctx.Logger().Info("Input: %s", input.AnInput)

	output := &Output{AnOutput: input.AnInput}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	pic, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("Error Decoding file: %v", err)
	}

	maxdim := 256
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
	src = imaging.Resize(src, w, h, imaging.Lanczos)

	return src, nil
}
