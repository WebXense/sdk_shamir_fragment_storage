package sdk_shamir_fragment_storage

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/WebXense/ginger/ginger"
	"github.com/WebXense/http"
)

type shamirFragmentStorageSdk struct {
	host string
}

func NewShamirFragmentStorageSdk(host string) *shamirFragmentStorageSdk {
	return &shamirFragmentStorageSdk{
		host: host,
	}
}

func (s *shamirFragmentStorageSdk) CreateFragment(key string, weight uint8, value string, secret string) (*FragmentDTO, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](s.host+url_create_fragment, nil, nil, &createFragmentRequest{
		Key:    key,
		Weight: weight,
		Value:  value,
		Secret: secret,
	})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &FragmentDTO{}), resp, nil
}

func (s *shamirFragmentStorageSdk) DeleteFragment(key, secret string) (*ginger.Response, error) {
	status, resp, err := http.Delete[ginger.Response](s.host+url_delete_fragment, nil, nil, &deleteFragmentRequest{
		Key:    key,
		Secret: secret,
	})
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, errors.New("delete fragment failed with status: " + strconv.Itoa(status))
	}

	return resp, nil
}

func (s *shamirFragmentStorageSdk) GetFragment(key, secret string) (*FragmentDTO, *ginger.Response, error) {
	status, resp, err := http.Get[ginger.Response](s.host+url_get_fragment, nil, map[string]string{
		"key":    key,
		"secret": secret,
	}, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("get fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &FragmentDTO{}), resp, nil
}

func (s *shamirFragmentStorageSdk) UpdateFragment(key string, weight uint8, value string, secret string) (*FragmentDTO, *ginger.Response, error) {
	status, resp, err := http.Put[ginger.Response](s.host+url_update_fragment, nil, nil, &updateFragmentRequest{
		Key:    key,
		Weight: weight,
		Value:  value,
		Secret: secret,
	})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("update fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &FragmentDTO{}), resp, nil
}

func mapByJson[T any](from interface{}, to *T) *T {
	j, err := json.Marshal(from)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(j, to)
	if err != nil {
		return nil
	}
	return to
}
