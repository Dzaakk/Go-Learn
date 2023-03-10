package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})
	t.Run("Joko", func(t *testing.T) {
		result := HelloWorld("Joko")
		require.Equal(t, "Hello Joko", result, "Result must be 'Hello Joko'")

	})
}

//test seluruh func testing yang ada
func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

//require memanggil failnow()
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Require Done")
}

//assert memanggil fail()
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}
func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		//error
		t.Error("Result must be 'Hello Eko'")
	}
}
func TestHelloWorldJoko(t *testing.T) {
	result := HelloWorld("Joko")

	if result != "Hello Joko" {
		//error
		t.Fatal("Result must be 'hello Joko'")
	}
}
