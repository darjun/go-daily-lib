package main

import (
  "errors"
  "fmt"
  "log"
  "strconv"
  "strings"

  "github.com/jessevdk/go-flags"
)

type MathCommand struct {
  Op string `long:"op" description:"operation to execute"`
  Args []string
  Result int64
}

func Calc(op string, nums []int64) int64 {
  if len(nums) == 0 {
    return 0
  }

  result := nums[0]
  for i := 1; i < len(nums); i++ {
	switch (op) {
	  case "+":
		result += nums[i]
	  case "-":
		result -= nums[i]
	  case "x":
		result *= nums[i]
	  case "÷":
		result /= nums[i]
	}
  }

  return result
}

func (this *MathCommand) Execute(args []string) error {
  if this.Op != "+" && this.Op != "-" && this.Op != "x" && this.Op != "÷" {
    return errors.New("invalid op")
  }

  nums := make([]int64, 0, len(args))
  for _, arg := range args {
    num, err := strconv.ParseInt(arg, 10, 64)
    if err != nil {
      return err
    }

    nums = append(nums, num)
  }

  this.Args = args
  this.Result = Calc(this.Op, nums)
  return nil
}

type Option struct {
  Math MathCommand `command:"math"`
}

func main() {
  var opt Option
  _, err := flags.Parse(&opt)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("The result of %s is %d", strings.Join(opt.Math.Args, opt.Math.Op), opt.Math.Result)
}