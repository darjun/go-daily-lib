package suite

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MyTestSuit struct {
	suite.Suite
	testCount uint32
}

func (s *MyTestSuit) SetupSuite() {
	fmt.Println("SetupSuite")
}

func (s *MyTestSuit) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

func (s *MyTestSuit) SetupTest() {
	fmt.Printf("SetupTest test count:%d\n", s.testCount)
}

func (s *MyTestSuit) TearDownTest() {
	s.testCount++
	fmt.Printf("TearDownTest test count:%d\n", s.testCount)
}

func (s *MyTestSuit) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest suite:%s test:%s\n", suiteName, testName)
}

func (s *MyTestSuit) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest suite:%s test:%s\n", suiteName, testName)
}

func (s *MyTestSuit) TestExample() {
	fmt.Println("TestExample")
}

func TestExample(t *testing.T) {
	suite.Run(t, new(MyTestSuit))
}
