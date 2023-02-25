// Code generated by "stringer -type astNodeType"; DO NOT EDIT.

package jmespath

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ASTEmpty-0]
	_ = x[ASTArithmeticExpression-1]
	_ = x[ASTArithmeticUnaryExpression-2]
	_ = x[ASTComparator-3]
	_ = x[ASTCurrentNode-4]
	_ = x[ASTRootNode-5]
	_ = x[ASTExpRef-6]
	_ = x[ASTFunctionExpression-7]
	_ = x[ASTField-8]
	_ = x[ASTFilterProjection-9]
	_ = x[ASTFlatten-10]
	_ = x[ASTIdentity-11]
	_ = x[ASTIndex-12]
	_ = x[ASTIndexExpression-13]
	_ = x[ASTKeyValPair-14]
	_ = x[ASTLiteral-15]
	_ = x[ASTMultiSelectHash-16]
	_ = x[ASTMultiSelectList-17]
	_ = x[ASTOrExpression-18]
	_ = x[ASTAndExpression-19]
	_ = x[ASTNotExpression-20]
	_ = x[ASTPipe-21]
	_ = x[ASTProjection-22]
	_ = x[ASTSubexpression-23]
	_ = x[ASTSlice-24]
	_ = x[ASTValueProjection-25]
}

const _astNodeType_name = "ASTEmptyASTArithmeticExpressionASTArithmeticUnaryExpressionASTComparatorASTCurrentNodeASTRootNodeASTExpRefASTFunctionExpressionASTFieldASTFilterProjectionASTFlattenASTIdentityASTIndexASTIndexExpressionASTKeyValPairASTLiteralASTMultiSelectHashASTMultiSelectListASTOrExpressionASTAndExpressionASTNotExpressionASTPipeASTProjectionASTSubexpressionASTSliceASTValueProjection"

var _astNodeType_index = [...]uint16{0, 8, 31, 59, 72, 86, 97, 106, 127, 135, 154, 164, 175, 183, 201, 214, 224, 242, 260, 275, 291, 307, 314, 327, 343, 351, 369}

func (i astNodeType) String() string {
	if i < 0 || i >= astNodeType(len(_astNodeType_index)-1) {
		return "astNodeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _astNodeType_name[_astNodeType_index[i]:_astNodeType_index[i+1]]
}
