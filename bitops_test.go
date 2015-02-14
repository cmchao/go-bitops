package bitops

import "testing"

func TestExtract32(t *testing.T) {
    var value uint32 = 0xF0F0F0F0
    var length uint
    var start uint
    var field uint32

    //check error
    start = 32
    length = 0
    _, err := Extract32(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid start")
    }

    start = 0
    length = 33
    _, err = Extract32(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid length")
    }

    start = 31
    length = 2
    _, err = Extract32(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid length from valid start")
    }

    //check pass case
    start = 0
    length = 1
    field, err = Extract32(value, start, length)
    if err != nil || field != 0x0 {
        t.Fail()
        t.Log("LSB")
    }

    start = 31
    length = 1
    field, err = Extract32(value, start, length)
    if err != nil || field != 0x1 {
        t.Fail()
        t.Log("MSB")
    }

    start = 4
    length = 4
    field, err = Extract32(value, start, length)
    if err != nil || field != 0xF {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestExtract64(t *testing.T) {
    var value uint64 = 0xF0F0F0F0F0F0F0F0
    var length uint
    var start uint
    var field uint64

    //check error
    start = 64
    length = 0
    _, err := Extract64(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid start")
    }

    start = 0
    length = 65
    _, err = Extract64(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid length")
    }

    start = 63
    length = 2
    _, err = Extract64(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid length from valid start")
    }

    //check pass case
    start = 0
    length = 1
    field, err = Extract64(value, start, length)
    if err != nil || field != 0x0 {
        t.Fail()
        t.Log("LSB")
    }

    start = 63
    length = 1
    field, err = Extract64(value, start, length)
    if err != nil || field != 0x1 {
        t.Fail()
        t.Log("MSB")
    }

    start = 12 
    length = 4
    field, err = Extract64(value, start, length)
    if err != nil || field != 0xF {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestGetField32(t *testing.T) {
    var value uint32 = 0xF0F0F0F0
    var high uint
    var low uint
    var field uint32

    //check error
    high = 32
    low = 0 
    _, err := GetField32(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("invalid high")
    }

    high = 0
    low = 32
    _, err = GetField32(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("invalid low")
    }

    high = 10
    low = 20
    _, err = GetField32(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("high < low")
    }

    //check pass case
    high = 0
    low = 0
    field, err = GetField32(value, high, low)
    if err != nil || field != 0x0 {
        t.Fail()
        t.Log("LSB")
    }

    high = 31
    low = 31
    field, err = GetField32(value, high, low)
    if err != nil || field != 0x1 {
        t.Fail()
        t.Log("MSB")
    }

    high = 7
    low = 4
    field, err = GetField32(value, high, low)
    if err != nil || field != 0xF {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestGetField64(t *testing.T) {
    var value uint64 = 0xF0F0F0F0F0F0F0F0
    var high uint
    var low uint
    var field uint64

    //check error
    high = 64
    low = 0 
    _, err := GetField64(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("invalid high")
    }

    high = 0
    low = 64
    _, err = GetField64(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("invalid low")
    }

    high = 10
    low = 20
    _, err = GetField64(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("high < low")
    }

    //check pass case
    high = 0
    low = 0
    field, err = GetField64(value, high, low)
    if err != nil || field != 0x0 {
        t.Fail()
        t.Log("LSB")
    }

    high = 63
    low = 63
    field, err = GetField64(value, high, low)
    if err != nil || field != 0x1 {
        t.Fail()
        t.Log("MSB")
    }

    high = 7
    low = 4
    field, err = GetField64(value, high, low)
    if err != nil || field != 0xF {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestDeposit32(t *testing.T) {
    var value uint32 = 0xF0F0F0F0
    var length uint
    var start uint
    var field uint32

    //check error
    start = 32
    length = 0
    field, err := Deposit32(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid start")
    }

    if value != field {
        t.Fail()
        t.Logf("invalid return %x", field)
    }

    start = 0
    length = 33
    field, err = Deposit32(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid length")
    }

    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    start = 31
    length = 2
    field, err = Deposit32(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid length from valid start")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    //check pass case
    start = 0
    length = 1
    field, err = Deposit32(value, start, length, 0x7)
    if err != nil || field != 0xF0F0F0F1 {
        t.Fail()
        t.Log("LSB")
    }

    start = 31
    length = 1
    field, err = Deposit32(value, start, length, 0x0)
    if err != nil || field != 0x70F0F0F0 {
        t.Fail()
        t.Log("MSB")
    }

    start = 4
    length = 16
    field, err = Deposit32(value, start, length, 0x1234)
    if err != nil || field != 0xF0F12340 {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestDeposit64(t *testing.T) {
    var value uint64 = 0xF0F0F0F0F0F0F0F0
    var length uint
    var start uint
    var field uint64

    //check error
    start = 64
    length = 0
    field, err := Deposit64(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid start")
    }

    if value != field {
        t.Fail()
        t.Logf("invalid return %x", field)
    }

    start = 0
    length = 65
    field, err = Deposit64(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid length")
    }

    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    start = 63
    length = 2
    field, err = Deposit64(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid length from valid start")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    //check pass case
    start = 0
    length = 1
    field, err = Deposit64(value, start, length, 0x7)
    if err != nil || field != 0xF0F0F0F0F0F0F0F1 {
        t.Fail()
        t.Log("LSB")
    }

    start = 63
    length = 1
    field, err = Deposit64(value, start, length, 0x0)
    if err != nil || field != 0x70F0F0F0F0F0F0F0 {
        t.Fail()
        t.Log("MSB")
    }

    start = 4
    length = 16
    field, err = Deposit64(value, start, length, 0x1234)
    if err != nil || field != 0xF0F0F0F0F0F12340 {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestSetField32(t *testing.T) {
    var value uint32 = 0xF0F0F0F0
    var high uint
    var low uint
    var field uint32

    //check error
    high = 32
    low = 0
    field, err := SetField32(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("invalid high")
    }
    if value != field {
        t.Fail()
        t.Logf("invalid return %x", field)
    }

    high = 0
    low = 32
    field, err = SetField32(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("invalid low")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    high = 10
    low = 20
    field, err = SetField32(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("high < low")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    //check pass case
    high = 0
    low = 0
    field, err = SetField32(value, high, low, 0x7)
    if err != nil || field != 0xF0F0F0F1 {
        t.Fail()
        t.Log("LSB")
    }

    high = 31
    low = 31
    field, err = SetField32(value, high, low, 0x0)
    if err != nil || field != 0x70F0F0F0 {
        t.Fail()
        t.Log("MSB")
    }

    high = 19
    low = 4
    field, err = SetField32(value, high, low, 0x1234)
    if err != nil || field != 0xF0F12340 {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestSetField64(t *testing.T) {
    var value uint64 = 0xF0F0F0F0F0F0F0F0
    var high uint
    var low uint
    var field uint64

    //check error
    high = 64
    low = 0
    field, err := SetField64(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("invalid high")
    }
    if value != field {
        t.Fail()
        t.Logf("invalid return %x", field)
    }

    high = 0
    low = 64
    field, err = SetField64(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("invalid low")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    high = 10
    low = 20
    field, err = SetField64(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("high < low")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    //check pass case
    high = 0
    low = 0
    field, err = SetField64(value, high, low, 0x7)
    if err != nil || field != 0xF0F0F0F0F0F0F0F1 {
        t.Fail()
        t.Log("LSB")
    }

    high = 63
    low = 63
    field, err = SetField64(value, high, low, 0x0)
    if err != nil || field != 0x70F0F0F0F0F0F0F0 {
        t.Fail()
        t.Log("MSB")
    }

    high = 19
    low = 4
    field, err = SetField64(value, high, low, 0x1234)
    if err != nil || field != 0xF0F0F0F0F0F12340 {
        t.Fail()
        t.Log("get valid field")
    }
}
