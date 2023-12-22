package averager

import "testing"

func Test_regexpAverager_interfaceTest(t *testing.T) {
	interfaceTest(t, &regexpAverager{})
}
