package bizerrs

var (
	ErrSystem = New("-1", "系统异常")

	// 通用异常，3位code
	ErrIllegalArgument = New("001", "参数不正确")
)
