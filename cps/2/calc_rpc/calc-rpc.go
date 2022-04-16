package calc_rpc

import (
	"errors"
	"math"
	"strings"

	"github.com/CSProjectsAvatar/distri-systems/utils"
	"golang.org/x/exp/slices"
)

type RemoteCalc float64

type Args struct {
	A, B float64
}

func (t *RemoteCalc) Add(args *Args, reply *float64) error {
	*reply = args.A + args.B
	return nil
}

func (t *RemoteCalc) Sub(args *Args, reply *float64) error {
	*reply = args.A - args.B
	return nil
}

func (t *RemoteCalc) Mul(args *Args, reply *float64) error {
	*reply = args.A * args.B
	return nil
}

func (t *RemoteCalc) Div(args *Args, reply *float64) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	*reply = args.A / args.B
	return nil
}

func (t *RemoteCalc) Pow(args *Args, reply *float64) error {
	*reply = math.Pow(args.A, args.B)
	return nil
}

func (t *RemoteCalc) Iam(name *string, reply *string) error {
	students, err := utils.ReadLines("class.txt")
	utils.CheckErr(err)
	if !slices.Contains(students, strings.ToLower(*name)) {
		students = append(students, strings.ToLower(*name))
	}
	err = utils.WriteLines("class.txt", students)
	utils.CheckErr(err)

	*reply = "Welcome " + *name
	return nil
}

func (t *RemoteCalc) Doc(args *Args, reply *string) error {
	*reply = "Add(args *Args, reply *float64)\n" +
		"Sub(args *Args, reply *float64)\n" +
		"Mul(args *Args, reply *float64)\n" +
		"Div(args *Args, reply *float64)\n" +
		"Pow(args *Args, reply *float64)\n" +
		"Iam(name *string, reply *string)\n"
	return nil
}
