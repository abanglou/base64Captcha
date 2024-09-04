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
	"reflect"
	"testing"
)

func TestDriverDigitV2_DrawCaptcha(t *testing.T) {
	type fields struct {
		Height     int
		Width      int
		CaptchaLen int
		MaxSkew    float64
		BgColor    *color.RGBA
		DotCount   int
		Fonts      []string
	}
	type args struct {
		content string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantItem Item
		wantErr  bool
	}{
		{"DigitV2", fields{80, 240, 6, 0.6, nil, 8, []string{"3Dumb.ttf"}}, args{RandText(6, TxtNumbers)}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DriverDigitV2{
				Height:   tt.fields.Height,
				Width:    tt.fields.Width,
				Length:   tt.fields.CaptchaLen,
				MaxSkew:  tt.fields.MaxSkew,
				BgColor:  tt.fields.BgColor,
				DotCount: tt.fields.DotCount,
				Fonts:    tt.fields.Fonts,
			}
			gotItem, err := d.DrawCaptcha(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("DriverDigitV2.DrawCaptcha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			itemWriteFile(gotItem, "_builds", tt.args.content, "png")

		})
	}
}

func TestNewDriverV2Digit(t *testing.T) {
	type args struct {
		height   int
		width    int
		length   int
		maxSkew  float64
		dotCount int
	}
	tests := []struct {
		name string
		args args
		want *DriverDigitV2
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDriverDigit(tt.args.height, tt.args.width, tt.args.length, tt.args.maxSkew, tt.args.dotCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDriverDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriverDigitV2_GenerateIdQuestionAnswer(t *testing.T) {
	tests := []struct {
		name   string
		d      *DriverDigitV2
		wantId string
		wantQ  string
		wantA  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotQ, gotA := tt.d.GenerateIdQuestionAnswer()
			if gotId != tt.wantId {
				t.Errorf("DriverDigitV2.GenerateIdQuestionAnswer() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotQ != tt.wantQ {
				t.Errorf("DriverDigitV2.GenerateIdQuestionAnswer() gotQ = %v, want %v", gotQ, tt.wantQ)
			}
			if gotA != tt.wantA {
				t.Errorf("DriverDigitV2.GenerateIdQuestionAnswer() gotA = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}
