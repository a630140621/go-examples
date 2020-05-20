package param

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestParam(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Param Suite")
}

var _ = Describe("Param", func() {
	Describe("Unmarshal", func() {
		Context("和 json.Unmarshal 相似处", func() {
			It("可以处理 匿名 组合字段", func() {
				type Base struct {
					Field1 string
					Field2 int
				}
				type Req struct {
					Man bool
					Base
					Base2   Base
					Message string
					Code    int8
				}

				data := []byte(`{"Man": true, "Field1": "field1", "Field2": 1, "Message": "message", "Base2": {"Field1": "Base2 Field1", "Field2": 2}, "Code": 1}`)
				req := &Req{}
				Expect(Unmarshal(data, req)).Should(Succeed())
				Expect(req.Man).Should(BeTrue())
				Expect(req.Field1).Should(Equal("field1"))
				Expect(req.Field2).Should(Equal(1))
				Expect(req.Message).Should(Equal("message"))
				Expect(req.Code).Should(Equal(int8(1)))
				Expect(req.Base2.Field1).Should(Equal("Base2 Field1"))
				Expect(req.Base2.Field2).Should(Equal(2))
			})
		})
	})
})
