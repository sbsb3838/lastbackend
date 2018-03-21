//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package request

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/lastbackend/lastbackend/pkg/distribution/errors"
)

type VolumeRequest struct{}

func (VolumeRequest) CreateOptions() *VolumeCreateOptions {
	return new(VolumeCreateOptions)
}

func (v *VolumeCreateOptions) Validate() *errors.Err {
	return nil
}

func (v *VolumeCreateOptions) DecodeAndValidate(reader io.Reader) *errors.Err {

	if reader == nil {
		err := errors.New("data body can not be null")
		return errors.New("volume").IncorrectJSON(err)
	}

	if reader == nil {
		err := errors.New("data body can not be null")
		return errors.New("volume").IncorrectJSON(err)
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("volume").Unknown(err)
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return errors.New("volume").IncorrectJSON(err)
	}

	return v.Validate()
}

func (v *VolumeCreateOptions) ToJson() ([]byte, error) {
	return json.Marshal(v)
}

func (VolumeRequest) UpdateOptions() *VolumeUpdateOptions {
	return new(VolumeUpdateOptions)
}

func (v *VolumeUpdateOptions) Validate() *errors.Err {
	return nil
}

func (v *VolumeUpdateOptions) DecodeAndValidate(reader io.Reader) *errors.Err {

	if reader == nil {
		err := errors.New("data body can not be null")
		return errors.New("volume").IncorrectJSON(err)
	}

	if reader == nil {
		err := errors.New("data body can not be null")
		return errors.New("volume").IncorrectJSON(err)
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("volume").Unknown(err)
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return errors.New("volume").IncorrectJSON(err)
	}

	return v.Validate()
}

func (v *VolumeUpdateOptions) ToJson() ([]byte, error) {
	return json.Marshal(v)
}

func (VolumeRequest) RemoveOptions() *VolumeRemoveOptions {
	return new(VolumeRemoveOptions)
}

func (v *VolumeRemoveOptions) Validate() *errors.Err {
	return nil
}

func (v *VolumeRemoveOptions) DecodeAndValidate(reader io.Reader) *errors.Err {

	if reader == nil {
		return nil
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("volume").Unknown(err)
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return errors.New("volume").IncorrectJSON(err)
	}

	return v.Validate()
}