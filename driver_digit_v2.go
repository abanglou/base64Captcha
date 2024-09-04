// Copyright 2017 Eric Zhou. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base64Captcha

import (
	"image/color"

	"github.com/golang/freetype/truetype"
)

// DriverDigit config for captcha-engine-digit.
type DriverDigitV2 struct {
	// Height png height in pixel.
	Height int
	// Width Captcha png width in pixel.
	Width int
	// DefaultLen Default number of digits in captcha solution.
	Length int
	// MaxSkew max absolute skew factor of a single digit.
	MaxSkew float64
	// DotCount Number of background circles.
	DotCount int
	//BgColor captcha image background color (optional)
	BgColor *color.RGBA
	//fontsStorage font storage (optional)
	fontsStorage FontsStorage

	//Fonts loads by name see fonts.go's comment
	Fonts      []string
	fontsArray []*truetype.Font
}

// NewDriverDigit creates a driver of digit
func NewDriverDigitV2(height int, width int, length int, maxSkew float64, dotCount int) *DriverDigit {
	return &DriverDigit{Height: height, Width: width, Length: length, MaxSkew: maxSkew, DotCount: dotCount}
}

// DefaultDriverDigit is a default driver of digit
var DefaultDriverDigitV2 = NewDriverDigit(80, 240, 5, 0.7, 80)

// ConvertFonts loads fonts from names
func (d *DriverDigitV2) ConvertFonts() *DriverDigitV2 {
	if d.fontsStorage == nil {
		d.fontsStorage = DefaultEmbeddedFonts
	}

	tfs := []*truetype.Font{}
	for _, fff := range d.Fonts {
		tf := d.fontsStorage.LoadFontByName("fonts/" + fff)
		tfs = append(tfs, tf)
	}
	if len(tfs) == 0 {
		tfs = fontsAll
	}
	d.fontsArray = tfs

	return d
}

// GenerateIdQuestionAnswer creates captcha content and answer
func (d *DriverDigitV2) GenerateIdQuestionAnswer() (id, q, a string) {
	id = RandomId()
	digits := randomDigits(d.Length)
	a = parseDigitsToString(digits)
	return id, a, a
}

// GenerateIdQuestionAnswer creates captcha content and answer
func (d *DriverDigitV2) GenerateSpecificIdQuestionAnswer(mId string) (id, q, a string) {
	id = mId
	digits := randomDigits(d.Length)
	a = parseDigitsToString(digits)
	return id, a, a
}

// DrawCaptcha creates digit captcha item
func (d *DriverDigitV2) DrawCaptcha(content string) (item Item, err error) {
	var bgc color.RGBA
	if d.BgColor != nil {
		bgc = *d.BgColor
	} else {
		bgc = RandLightColor()
	}
	// itemDigit := NewItemDigit(d.Width, d.Height, d.DotCount, d.MaxSkew)
	itemChar := NewItemChar(d.Width, d.Height, bgc)
	//draw question
	err = itemChar.drawTextV2(content, d.fontsArray)
	if err != nil {
		return
	}
	return itemChar, nil
}
