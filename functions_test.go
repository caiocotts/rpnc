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
		DescribeTable("Adding values together", func(a string, b string, expected string) {
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
		It("should return an error on an incorrect number of arguments", func() {
			err := Add(&stack)
			Expect(err).To(Not(BeNil()))
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
		It("should return an error on an incorrect number of arguments", func() {
			err := Subtract(&stack)
			Expect(err).To(Not(BeNil()))
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

		It("should return an error on an incorrect number of arguments", func() {
			err := Multiply(&stack)
			Expect(err).To(Not(BeNil()))
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

		It("should return an error on an incorrect number of arguments", func() {
			err := Divide(&stack)
			Expect(err).To(Not(BeNil()))
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

		It("should return an error on an incorrect number of arguments", func() {
			err := Drop(&stack)
			Expect(err).To(Not(BeNil()))
		})
	})

	Describe("Dup", func() {
		It("should duplicate the item at the first level of the stack once", func() {
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

		It("should return an error on an incorrect number of arguments", func() {
			err := Dup(&stack)
			Expect(err).To(Not(BeNil()))
		})
	})

	Describe("Clear", func() {
		It("should remove all values from the stack", func() {
			stack.Push("3")
			stack.Push("4")
			stack.Push("5")

			_ = Clear(&stack)

			numbers := stack.ToSlice()
			var expected []string

			Expect(numbers).To(Equal(expected))
		})
	})

	Describe("Swap", func() {
		It("should swap the values from level 1 and 2 of the stack", func() {
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

		It("should return an error on an incorrect number of arguments", func() {
			err := Divide(&stack)
			Expect(err).To(Not(BeNil()))
		})
	})
})
