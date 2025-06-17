package api

var next_request_id uint64 = 0

func GenRequestId() uint64 {
	res := next_request_id
	next_request_id++
	return res
}
