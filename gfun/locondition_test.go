package gfun

import (
	"testing"
)

func TestTernary(t *testing.T) {
	result := Ternary(true, "a", "b")
	t.Logf("%v", result)
	// Output: a
}

func TestTernaryF(t *testing.T) {
	result := TernaryF(true, func() string { return "a" }, func() string { return "b" })

	t.Logf("%v", result)
	// Output: a
}

func TestIf(t *testing.T) {
	result1 := If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
	t.Logf("%v\n", result5)
	t.Logf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func TestIfF(t *testing.T) {
	result1 := If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
	t.Logf("%v\n", result5)
	t.Logf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func TestIfElseAndElseIf(t *testing.T) {
	result1 := If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
	t.Logf("%v\n", result5)
	t.Logf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func TestIfElseAndElseIfF(t *testing.T) {
	result1 := If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
	t.Logf("%v\n", result5)
	t.Logf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func TestIfElseAndElse(t *testing.T) {
	result1 := If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
	t.Logf("%v\n", result5)
	t.Logf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func TestIfElseAndElseF(t *testing.T) {
	result1 := If(true, 1).
		ElseIf(false, 2).
		Else(3)

	result2 := If(false, 1).
		ElseIf(true, 2).
		Else(3)

	result3 := If(false, 1).
		ElseIf(false, 2).
		Else(3)

	result4 := IfF(true, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result5 := IfF(false, func() int { return 1 }).
		ElseIfF(true, func() int { return 2 }).
		ElseF(func() int { return 3 })

	result6 := IfF(false, func() int { return 1 }).
		ElseIfF(false, func() int { return 2 }).
		ElseF(func() int { return 3 })

	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
	t.Logf("%v\n", result5)
	t.Logf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func TestSwitch(t *testing.T) {
	result1 := Switch[int, string](1).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result2 := Switch[int, string](2).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result3 := Switch[int, string](42).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result4 := Switch[int, string](1).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result5 := Switch[int, string](2).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result6 := Switch[int, string](42).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
	t.Logf("%v\n", result5)
	t.Logf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func TestSwitchCase(t *testing.T) {
	result1 := Switch[int, string](1).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result2 := Switch[int, string](2).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result3 := Switch[int, string](42).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result4 := Switch[int, string](1).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result5 := Switch[int, string](2).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result6 := Switch[int, string](42).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
	t.Logf("%v\n", result5)
	t.Logf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func TestSwitchCaseF(t *testing.T) {
	result1 := Switch[int, string](1).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result2 := Switch[int, string](2).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result3 := Switch[int, string](42).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result4 := Switch[int, string](1).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result5 := Switch[int, string](2).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result6 := Switch[int, string](42).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
	t.Logf("%v\n", result5)
	t.Logf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func TestSwitchCaseDefault(t *testing.T) {
	result1 := Switch[int, string](1).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result2 := Switch[int, string](2).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result3 := Switch[int, string](42).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result4 := Switch[int, string](1).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result5 := Switch[int, string](2).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result6 := Switch[int, string](42).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
	t.Logf("%v\n", result5)
	t.Logf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func TestSwitchCaseDefaultF(t *testing.T) {
	result1 := Switch[int, string](1).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result2 := Switch[int, string](2).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result3 := Switch[int, string](42).
		Case(1, "1").
		Case(2, "2").
		Default("3")

	result4 := Switch[int, string](1).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result5 := Switch[int, string](2).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	result6 := Switch[int, string](42).
		CaseF(1, func() string { return "1" }).
		CaseF(2, func() string { return "2" }).
		DefaultF(func() string { return "3" })

	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
	t.Logf("%v\n", result5)
	t.Logf("%v\n", result6)
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}
