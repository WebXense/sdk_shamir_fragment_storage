package sdk_shamir_fragment_storage

type FragmentDTO struct {
	Key    string `json:"key"`
	Weight uint8  `json:"weight"`
	Value  string `json:"value"`
}

type createFragmentRequest struct {
	Key    string `json:"key" binding:"required"`
	Weight uint8  `json:"weight" binding:"required"`
	Value  string `json:"value" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}

type deleteFragmentRequest struct {
	Key    string `json:"key" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}

type updateFragmentRequest struct {
	Key    string `json:"key" binding:"required"`
	Weight uint8  `json:"weight" binding:"required"`
	Value  string `json:"value" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}
