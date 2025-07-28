package mathutils_test

import (
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestMathutils(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Mathutils Suite")
}
