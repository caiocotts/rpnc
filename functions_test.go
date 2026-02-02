package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Functions", func() {
	var stack Stack[string]

	BeforeEach(func() {
		stack = Stack[string]{}
	})

	Describe("Add", func() {
		DescribeTable("Adding values", func(a string, b string, expected string) {
			stack.Push(a)
			stack.Push(b)
			_ = Add(&stack)

			result, _ := stack.Pop()

			Expect(result).To(Equal(expected))
		},
			Entry("should be able to add whole numbers", "-4", "3", "-1"),
			Entry("should be able to add mixed float and whole numbers", "1", "1.5", "2.5"),
			Entry("should be able to add floating point numbers", "1.5", "1.5", "3"),
		)
		When("an incorrect number of arguments is passed", func() {
			It("should return an error", func() {
				err := Add(&stack)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Subtract", func() {
		DescribeTable("Subtracting values", func(a string, b string, expected string) {
			stack.Push(a)
			stack.Push(b)
			_ = Subtract(&stack)

			result, _ := stack.Pop()

			Expect(result).To(Equal(expected))
		},
			Entry("should be able to subtract whole numbers", "4", "3", "1"),
			Entry("should be able to subtract mixed float and whole numbers", "1", "1.5", "-0.5"),
			Entry("should be able to subtract two floating point numbers", "1.5", "1.5", "0"),
		)
		When("an incorrect number of arguments is passed", func() {
			It("should return an error", func() {
				err := Subtract(&stack)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Multiply", func() {
		DescribeTable("Multiplying values", func(a string, b string, expected string) {
			stack.Push(a)
			stack.Push(b)
			_ = Multiply(&stack)

			result, _ := stack.Pop()

			Expect(result).To(Equal(expected))
		},
			Entry("should be able to multiply whole numbers", "4", "3", "12"),
			Entry("should be able to multiply mixed float and whole numbers", "-1", "1.5", "-1.5"),
			Entry("should be able to multiply two floating point numbers", "1.5", "1.5", "2.25"),
		)
		When("an incorrect number of arguments is passed", func() {
			It("should return an error", func() {
				err := Multiply(&stack)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Divide", func() {
		DescribeTable("Dividing values", func(a string, b string, expected string) {
			stack.Push(a)
			stack.Push(b)
			_ = Divide(&stack)

			result, _ := stack.Pop()

			Expect(result).To(Equal(expected))
		},
			Entry("should be able to divide whole numbers", "4", "3", "1.3333333333333333"),
			Entry("should be able to divide mixed float and whole numbers", "1", "1.5", "0.6666666666666666"),
			Entry("should be able to divide two floating point numbers", "-1.5", "1.5", "-1"),
		)
		When("an incorrect number of arguments is passed", func() {
			It("should return an error", func() {
				err := Divide(&stack)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Drop", func() {
		It("should pop the stack once", func() {
			stack.Push("3")
			stack.Push("4")
			stack.Push("5")

			_ = Drop(&stack)

			numbers := stack.ToSlice()
			expected := []string{
				"3",
				"4",
			}

			Expect(numbers).To(Equal(expected))
		})
		When("an incorrect number of arguments is passed", func() {
			It("should return an error", func() {
				err := Drop(&stack)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Dup", func() {
		It("should duplicate the element at the first level of the stack once", func() {
			stack.Push("3")
			stack.Push("4")
			stack.Push("5")

			_ = Dup(&stack)

			numbers := stack.ToSlice()
			expected := []string{
				"3",
				"4",
				"5",
				"5",
			}

			Expect(numbers).To(Equal(expected))
		})
		When("an incorrect number of arguments is passed", func() {
			It("should return an error", func() {
				err := Dup(&stack)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Clear", func() {
		It("should remove all elements from the stack", func() {
			stack.Push("3")
			stack.Push("4")
			stack.Push("5")

			_ = Clear(&stack)

			numbers := stack.ToSlice()

			Expect(numbers).To(BeEmpty())
		})
	})

	Describe("Swap", func() {
		It("should swap the elements from level 1 and 2 of the stack", func() {
			stack.Push("3")
			stack.Push("4")
			stack.Push("5")

			_ = Swap(&stack)

			numbers := stack.ToSlice()
			expected := []string{
				"3",
				"5",
				"4",
			}

			Expect(numbers).To(Equal(expected))
		})
		When("an incorrect number of arguments is passed", func() {
			It("should return an error", func() {
				err := Divide(&stack)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Roll", func() {
		DescribeTable("when given an argument N and there are N number of elements on the stack", func(n string, expected []string) {
			stack.Push("3")
			stack.Push("4")
			stack.Push("32")
			stack.Push("7")
			stack.Push("23")

			stack.Push(n)

			_ = Roll(&stack)

			result := stack.ToSlice()

			Expect(result).To(Equal(expected))
		},
			Entry("should do nothing when N is a negative number", "-1",
				[]string{
					"3",
					"4",
					"32",
					"7",
					"23",
				}),
			Entry("should do nothing when N is 0", "0",
				[]string{
					"3",
					"4",
					"32",
					"7",
					"23",
				}),
			Entry("should do nothing when N is 1", "1",
				[]string{
					"3",
					"4",
					"32",
					"7",
					"23",
				}),
			Entry("should shift levels 1 to 2 up by one level and bring the 3rd element to level 1 when N is 3", "3",
				[]string{
					"3",
					"4",
					"7",
					"23",
					"32",
				}),
			Entry("should shift levels 1 to 4 up by one level and bring the 5th element to level 1 when N is 5", "5",
				[]string{
					"4",
					"32",
					"7",
					"23",
					"3",
				},
			),
		)
		When("an incorrect number of arguments is passed", func() {
			It("should return an error", func() {
				err := Roll(&stack)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Rot", func() {
		When("there are at least 3 elements on the stack", func() {
			It("should shift levels 1 to 2 up by one level and bring the 3rd element to level 1", func() {
				stack.Push("3")
				stack.Push("4")
				stack.Push("5")

				_ = Rot(&stack)

				elements := stack.ToSlice()
				expected := []string{
					"4",
					"5",
					"3",
				}

				Expect(elements).To(Equal(expected))
			})
		})
	})
	When("an incorrect number of arguments is passed", func() {
		It("should return an error", func() {
			err := Rot(&stack)
			Expect(err).ToNot(BeNil())
		})
	})
})
