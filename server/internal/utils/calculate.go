package utils

import (
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func Calculate(userBalanceUnits, expenseAmountUnits int64, userBalanceNanos, expenseAmountNanos int32, calctype string) (int64, int64, string, error) {
	var resultBalanceUnits int64
	var resultBalanceNanos int64
	userBalUnits := strconv.Itoa(int(userBalanceUnits))
	userBalNanos := strconv.Itoa(int(userBalanceNanos))
	expenseAmtUnits := strconv.Itoa(int(expenseAmountUnits))
	expenseAmtNanos := strconv.Itoa(int(expenseAmountNanos))
	userBal := userBalUnits + "."
	for i := len(userBalNanos); i < 9; i++ {
		userBal += "0"
	}
	if strings.Contains(userBalNanos, "-") {
		userBal += strings.Replace(userBalNanos, "-", "", 1)
		if !strings.Contains(userBal, "-") {
			temp := "-" + userBal
			userBal = temp
		}
	} else {
		userBal += userBalNanos
	}
	expAmt := expenseAmtUnits + "."
	for i := len(expenseAmtNanos); i < 9; i++ {
		expAmt += "0"
	}
	if strings.Contains(expenseAmtNanos, "-") {
		expAmt += strings.Replace(expenseAmtNanos, "-", "", 1)
		if !strings.Contains(expAmt, "-") {
			temp := "-" + expAmt
			expAmt = temp
		}
	} else {
		expAmt += expenseAmtNanos
	}
	decimal.DivisionPrecision = 9
	userBalDec, err := decimal.NewFromString(userBal)
	if err != nil {
		return 0, 0, "", err
	}
	expAmtDec, err := decimal.NewFromString(expAmt)
	if err != nil {
		return 0, 0, "", err
	}
	var result string
	if calctype == "add" {
		result = userBalDec.Add(expAmtDec).String()

	} else if calctype == "sub" {
		result = userBalDec.Sub(expAmtDec).String()
	}
	resultArr := strings.Split(result, ".")
	// Set Nanos and Units
	resultUnit, _ := strconv.Atoi(string(resultArr[0]))
	resultBalanceUnits = int64(resultUnit)
	if len(resultArr) == 2 {
		currNanos := string(resultArr[1])
		currLength := len(currNanos)
		for currLength < 9 {
			currNanos += "0"
			currLength++
		}
		resultNanos, _ := strconv.Atoi(currNanos)
		resultBalanceNanos = int64(resultNanos)
		if strings.Contains(result, "-") {
			temp := -(resultBalanceNanos)
			resultBalanceNanos = temp
		}
	} else {
		resultBalanceNanos = 0
	}

	return resultBalanceUnits, resultBalanceNanos, result, nil
}
