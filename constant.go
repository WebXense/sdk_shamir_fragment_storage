package sdk_shamir_fragment_storage

const (
	url_create_fragment = "/fragment"
	url_delete_fragment = "/fragment"
	url_get_fragment    = "/fragment"
	url_update_fragment = "/fragment"
)

const (
	ERR_FRAGMENT_ALREADY_EXISTS = "3ca7f58b-a9a6-45a3-9294-263fac310d67"
	ERR_FRAGMENT_CREATE_FAILED  = "f6253b06-7ef1-4f55-9372-9ce335f8ba71"
	ERR_FRAGMENT_NOT_FOUND      = "f6253b06-7ef1-4f55-9372-9ce335f8ba71"
	ERR_FRAGMENT_UPDATE_FAILED  = "277ceced-8ce8-47fc-8038-608765fd5397"
	ERR_FRAGMENT_DELETE_FAILED  = "480e7c26-188c-483b-8db0-e7283cd22080"
	ERR_SECRET_NOT_MATCH        = "9cd42ce3-86b3-45e0-a4af-6d3f360b4ae9"
)
