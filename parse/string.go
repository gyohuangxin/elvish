// Code generated by "stringer -type=PrimaryType,RedirMode,ControlKind -output=string.go"; DO NOT EDIT

package parse

import "fmt"

const _PrimaryType_name = "BadPrimaryBarewordSingleQuotedDoubleQuotedVariableWildcardTildeErrorCaptureOutputCaptureListLambdaMapBraced"

var _PrimaryType_index = [...]uint8{0, 10, 18, 30, 42, 50, 58, 63, 75, 88, 92, 98, 101, 107}

func (i PrimaryType) String() string {
	if i < 0 || i >= PrimaryType(len(_PrimaryType_index)-1) {
		return fmt.Sprintf("PrimaryType(%d)", i)
	}
	return _PrimaryType_name[_PrimaryType_index[i]:_PrimaryType_index[i+1]]
}

const _RedirMode_name = "BadRedirModeReadWriteReadWriteAppend"

var _RedirMode_index = [...]uint8{0, 12, 16, 21, 30, 36}

func (i RedirMode) String() string {
	if i < 0 || i >= RedirMode(len(_RedirMode_index)-1) {
		return fmt.Sprintf("RedirMode(%d)", i)
	}
	return _RedirMode_name[_RedirMode_index[i]:_RedirMode_index[i+1]]
}

const _ControlKind_name = "BadControlIfControlWhileControlForControlBeginControl"

var _ControlKind_index = [...]uint8{0, 10, 19, 31, 41, 53}

func (i ControlKind) String() string {
	if i < 0 || i >= ControlKind(len(_ControlKind_index)-1) {
		return fmt.Sprintf("ControlKind(%d)", i)
	}
	return _ControlKind_name[_ControlKind_index[i]:_ControlKind_index[i+1]]
}
