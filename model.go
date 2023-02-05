package sdk_shamir_fragment_storage

type FragmentDTO struct {
	Key    string `json:"key"`
	Weight uint8  `json:"weight"`
	Value  string `json:"value"`
}

type CreateFragmentRequest struct {
	Key    string `json:"key" binding:"required"`
	Weight uint8  `json:"weight" binding:"required"`
	Value  string `json:"value" binding:"required"`
}

type DeleteFragmentRequest struct {
	Key string `json:"key" binding:"required"`
}

type GetFragmentRequest struct {
	Key string `form:"key" binding:"required"`
}

type UpdateFragmentRequest struct {
	Key    string `json:"key" binding:"required"`
	Weight uint8  `json:"weight" binding:"required"`
	Value  string `json:"value" binding:"required"`
}
