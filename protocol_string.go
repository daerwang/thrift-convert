// generated by stringer -type Protocol; DO NOT EDIT

package main

import "fmt"

const _Protocol_name = "BinaryCompactJSON"

var _Protocol_index = [...]uint8{0, 6, 13, 17}

func (i Protocol) String() string {
	if i < 0 || i+1 >= Protocol(len(_Protocol_index)) {
		return fmt.Sprintf("Protocol(%d)", i)
	}
	return _Protocol_name[_Protocol_index[i]:_Protocol_index[i+1]]
}
