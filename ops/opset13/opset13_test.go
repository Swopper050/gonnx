package opset13

import (
	"testing"

	"github.com/advancedclimatesystems/gonnx/ops"
	"github.com/stretchr/testify/assert"
)

func TestGetOperator(t *testing.T) {
	tests := []struct {
		opType   string
		expected ops.Operator
		err      error
	}{
		{
			"Abs",
			newAbs(),
			nil,
		},
		{
			"Acos",
			newAcos(),
			nil,
		},
		{
			"Acosh",
			newAcosh(),
			nil,
		},
		{
			"Add",
			newAdd(),
			nil,
		},
		{
			"Atan",
			newAtan(),
			nil,
		},
		{
			"Atanh",
			newAtanh(),
			nil,
		},
		{
			"Asin",
			newAsin(),
			nil,
		},
		{
			"Asinh",
			newAsinh(),
			nil,
		},
		{
			"Cast",
			newCast(),
			nil,
		},
		{
			"Concat",
			newConcat(),
			nil,
		},
		{
			"Constant",
			newConstant(),
			nil,
		},
		{
			"ConstantOfShape",
			newConstantOfShape(),
			nil,
		},
		{
			"Conv",
			newConv(),
			nil,
		},
		{
			"Cos",
			newCos(),
			nil,
		},
		{
			"Cosh",
			newCosh(),
			nil,
		},
		{
			"Div",
			newDiv(),
			nil,
		},
		{
			"Gather",
			newGather(),
			nil,
		},
		{
			"Gemm",
			newGemm(),
			nil,
		},
		{
			"GRU",
			newGRU(),
			nil,
		},
		{
			"MatMul",
			newMatMul(),
			nil,
		},
		{
			"Mul",
			newMul(),
			nil,
		},
		{
			"Not",
			newNot(),
			nil,
		},
		{
			"Relu",
			newRelu(),
			nil,
		},
		{
			"Reshape",
			newReshape(),
			nil,
		},
		{
			"RNN",
			newRNN(),
			nil,
		},
		{
			"Scaler",
			newScaler(),
			nil,
		},
		{
			"Shape",
			newShape(),
			nil,
		},
		{
			"Sigmoid",
			newSigmoid(),
			nil,
		},
		{
			"Sin",
			newSin(),
			nil,
		},
		{
			"Sinh",
			newSinh(),
			nil,
		},
		{
			"Slice",
			newSlice(),
			nil,
		},
		{
			"Softmax",
			newSoftmax(),
			nil,
		},
		{
			"Squeeze",
			newSqueeze(),
			nil,
		},
		{
			"Sub",
			newSub(),
			nil,
		},
		{
			"Tan",
			newTan(),
			nil,
		},
		{
			"Tanh",
			newTanh(),
			nil,
		},
		{
			"Transpose",
			newTranspose(),
			nil,
		},
		{
			"Unsqueeze",
			newUnsqueeze(),
			nil,
		},
		{
			"NotYetImplemented",
			nil,
			ops.ErrUnknownOperatorType("NotYetImplemented"),
		},
	}

	for _, test := range tests {
		op, err := GetOperator(test.opType)

		assert.Equal(t, test.expected, op)
		assert.Equal(t, test.err, err)
	}
}
