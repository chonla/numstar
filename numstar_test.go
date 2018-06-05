package numstar_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chonla/numstar"
)

func TestReturnEmptyIfNotASingleNumber(t *testing.T) {
	ns := numstar.New()
	data := []string{"12", "a", "_", ""}
	for _, v := range data {
		result := ns.BigNum(v)
		assert.Empty(t, result, "must be empty")
	}
}

func TestReturnNonEmptyIfASingleNumber(t *testing.T) {
	ns := numstar.New()
	result := ns.BigNum("1")
	assert.NotEmpty(t, result, "must be not empty")
}

func TestConcatBigNumber(t *testing.T) {
	m := ` ** 
 ** 
****`
	n := ` A* 
 *X 
****`
	expected := ` **     A* 
 **     *X 
****   ****`
	ns := numstar.New()
	result := ns.Concat(m, n, 3)
	assert.Equal(t, expected, result, "must be equal")
}

func TestConcatTwoEmptyStringShouldReturnEmptyString(t *testing.T) {
	m := ``
	n := ``
	expected := ``
	ns := numstar.New()
	result := ns.Concat(m, n, 3)
	assert.Equal(t, expected, result, "must be equal")
}

func TestConcatBigNumberWithEmptyStringShouldReturnOneNotEmpty(t *testing.T) {
	m := ``
	n := ` A* 
 *X 
****`
	expected := n
	ns := numstar.New()
	result := ns.Concat(m, n, 3)
	assert.Equal(t, expected, result, "must be equal")
}

func TestConcatBigNumberWithEmptyStringShouldReturnOneNotEmpty2(t *testing.T) {
	m := ` A* 
 *X 
****`
	n := ``
	expected := m
	ns := numstar.New()
	result := ns.Concat(m, n, 3)
	assert.Equal(t, expected, result, "must be equal")
}

func TestResultShouldContainOnlySpaceStarAndNewLines(t *testing.T) {
	m := `1234567890`
	ns := numstar.New()
	result := ns.Big(m)
	result = strings.Replace(result, "\n", "", -1)
	result = strings.Replace(result, " ", "", -1)
	result = strings.Replace(result, "*", "", -1)
	assert.Empty(t, result, "must contain only spaces, stars and new lines")
}

func TestResultShouldContainOnlyGivenSpaceStarAndNewLines(t *testing.T) {
	m := `1234567890`
	ns := numstar.New()
	ns.Space("_")
	ns.Star("^")
	ns.Gap(10)
	result := ns.Big(m)
	result = strings.Replace(result, "\n", "", -1)
	result = strings.Replace(result, "_", "", -1)
	result = strings.Replace(result, "^", "", -1)
	assert.Empty(t, result, "must contain only spaces, stars and new lines")
}
