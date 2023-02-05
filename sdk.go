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

func (s *shamirFragmentStorageSdk) CreateFragment(request *CreateFragmentRequest) (*ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](s.host+url_create_fragment, nil, nil, request)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}

	resp.Data = mapByJson(resp.Data, &FragmentDTO{})
	return resp, nil
}

func (s *shamirFragmentStorageSdk) DeleteFragment(request *DeleteFragmentRequest) (*ginger.Response, error) {
	status, resp, err := http.Delete[ginger.Response](s.host+url_delete_fragment, nil, nil, request)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, errors.New("delete fragment failed with status: " + strconv.Itoa(status))
	}

	return resp, nil
}

func (s *shamirFragmentStorageSdk) GetFragment(request *GetFragmentRequest) (*ginger.Response, error) {
	status, resp, err := http.Get[ginger.Response](s.host+url_get_fragment, nil, map[string]string{
		"key": request.Key,
	}, nil)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, errors.New("get fragment failed with status: " + strconv.Itoa(status))
	}

	resp.Data = mapByJson(resp.Data, &FragmentDTO{})
	return resp, nil
}

func (s *shamirFragmentStorageSdk) UpdateFragment(request *UpdateFragmentRequest) (*ginger.Response, error) {
	status, resp, err := http.Put[ginger.Response](s.host+url_update_fragment, nil, nil, request)
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, errors.New("update fragment failed with status: " + strconv.Itoa(status))
	}

	resp.Data = mapByJson(resp.Data, &FragmentDTO{})
	return resp, nil
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
