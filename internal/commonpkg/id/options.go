package id

import "time"

type CodeOptions struct {
	chars []rune // 字符集（默认排除易混淆字符）
	n1    int    // 乘数（需与字符集长度互质）
	n2    int    // 位置偏移量（需与编码长度互质）
	l     int    // 编码长度（默认8）
	salt  uint64 // 盐值（默认123567369）
}

func WithCodeChars(arr []rune) func(*CodeOptions) {
	return func(options *CodeOptions) {
		if len(arr) > 0 {
			getCodeOptionsOrSetDefault(options).chars = arr
		}
	}
}

func WithCodeN1(n int) func(*CodeOptions) {
	return func(options *CodeOptions) {
		getCodeOptionsOrSetDefault(options).n1 = n
	}
}

func WithCodeN2(n int) func(*CodeOptions) {
	return func(options *CodeOptions) {
		getCodeOptionsOrSetDefault(options).n2 = n
	}
}

func WithCodeL(l int) func(*CodeOptions) {
	return func(options *CodeOptions) {
		if l > 0 {
			getCodeOptionsOrSetDefault(options).l = l
		}
	}
}

func WithCodeSalt(salt uint64) func(*CodeOptions) {
	return func(options *CodeOptions) {
		if salt > 0 {
			getCodeOptionsOrSetDefault(options).salt = salt
		}
	}
}

func getCodeOptionsOrSetDefault(options *CodeOptions) *CodeOptions {
	if options == nil {
		return &CodeOptions{
			// base string set, remove 0,1,I,O,U,Z
			chars: []rune{
				'2', '3', '4', '5', '6',
				'7', '8', '9', 'A', 'B',
				'C', 'D', 'E', 'F', 'G',
				'H', 'J', 'K', 'L', 'M',
				'N', 'P', 'Q', 'R', 'S',
				'T', 'V', 'W', 'X', 'Y',
			},
			// n1 / len(chars)=30 cop rime
			n1: 17,
			// n2 / l cop rime
			n2: 5,
			// code length
			l: 8,
			// random number
			salt: 123567369,
		}
	}
	return options
}

type SonyflakeOptions struct {
	machineId uint16
	startTime time.Time
}

// 机器ID (默认1)
func WithSonyflakeMachineId(id uint16) func(*SonyflakeOptions) {
	return func(options *SonyflakeOptions) {
		if id > 0 {
			getSonyflakeOptionsOrSetDefault(options).machineId = id
		}
	}
}

func WithSonyflakeStartTime(startTime time.Time) func(*SonyflakeOptions) {
	return func(options *SonyflakeOptions) {
		if !startTime.IsZero() {
			getSonyflakeOptionsOrSetDefault(options).startTime = startTime
		}
	}
}

// 时间起点 (默认2025-1-1)
func getSonyflakeOptionsOrSetDefault(options *SonyflakeOptions) *SonyflakeOptions {
	if options == nil {
		return &SonyflakeOptions{
			machineId: 1,
			startTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		}
	}
	return options
}
