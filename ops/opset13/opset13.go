package opset13

import (
	"github.com/advancedclimatesystems/gonnx/ops"
)

var operators13 = map[string]func() ops.Operator{
	"Abs":             newAbs,
	"Acos":            newAcos,
	"Acosh":           newAcosh,
	"Add":             newAdd,
	"Asin":            newAsin,
	"Asinh":           newAsinh,
	"Atan":            newAtan,
	"Atanh":           newAtanh,
	"Cast":            newCast,
	"Concat":          newConcat,
	"Constant":        newConstant,
	"ConstantOfShape": newConstantOfShape,
	"Conv":            newConv,
	"Cos":             newCos,
	"Cosh":            newCosh,
	"Div":             newDiv,
	"Gather":          newGather,
	"Gemm":            newGemm,
	"GRU":             newGRU,
	"MatMul":          newMatMul,
	"Mul":             newMul,
	"Not":             newNot,
	"PRelu":           newPRelu,
	"Relu":            newRelu,
	"Reshape":         newReshape,
	"Scaler":          newScaler,
	"Shape":           newShape,
	"Sigmoid":         newSigmoid,
	"Sin":             newSin,
	"Sinh":            newSinh,
	"Slice":           newSlice,
	"Softmax":         newSoftmax,
	"Squeeze":         newSqueeze,
	"Sub":             newSub,
	"Tan":             newTan,
	"Tanh":            newTanh,
	"Transpose":       newTranspose,
	"Unsqueeze":       newUnsqueeze,
}

// GetOperator maps strings as found in the ModelProto to Operators from opset 13.
func GetOperator(operatorType string) (ops.Operator, error) {
	if opInit, ok := operators13[operatorType]; ok {
		return opInit(), nil
	}

	return nil, ops.ErrUnknownOperatorType(operatorType)
}

// GetOpNames returns a list with operator names for opset 13.
func GetOpNames() []string {
	opList := make([]string, 0, len(operators13))

	for opName := range operators13 {
		opList = append(opList, opName)
	}

	return opList
}
