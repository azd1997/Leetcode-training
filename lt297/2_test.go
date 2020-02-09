package lt297

import (
	"fmt"
	"testing"
)

func TestCodec(t *testing.T) {
	var input = "[1,2,3,null,null,4,5,null,null]"
	var codec Codec
	output := codec.serialize(codec.deserialize(input))
	//if input!=output {
	//	t.Errorf("input: %s, but get %s\n", input, output)
	//}
	fmt.Println(output)
}
