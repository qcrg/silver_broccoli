package api

import "capnproto.org/go/capnp/v3"

func make_segment() (*capnp.Segment, error) {
	arena := capnp.SingleSegment(nil)
	_, seg, err := capnp.NewMessage(arena)
	if err != nil {
		arena.Release()
		return nil, err
	}
	return seg, nil
}

// func MakeGetBalanceInfo() (*GetBalanceInfo, error) {
// 	seg, err := make_segment()
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := NewGetBalanceInfo(seg)
// 	if err != nil {
// 		seg.Message().Release()
// 		return nil, err
// 	}
// 	return &res, nil
// }
