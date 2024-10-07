package errors

import "sync"

type ErrorCode int

const (
	AuthNeedLogin             = ErrorCode(401001)
	AuthInvalidToken          = ErrorCode(401002)
	NoPermission              = ErrorCode(401003)
	UserNamePasswordInCorrect = ErrorCode(401004)
	NoRefreshCodeFound        = ErrorCode(401005)
	InvalidRefreshCode        = ErrorCode(401006)
	PortfolioIsLocked         = ErrorCode(403007)
	ParameterInvalid          = ErrorCode(400001)
	ValidationErrors          = ErrorCode(400002)
	NoExchangeSelect          = ErrorCode(400003)
	DuplicateResource         = ErrorCode(400004)
	ArticleCanNotComment      = ErrorCode(400005)
	PasswordNotValid          = ErrorCode(400006)
	CreditNotEnough           = ErrorCode(400007)

	OrderTypeInvalid       = ErrorCode(400708)
	OrderAlreadyConfirmed  = ErrorCode(400709)
	OrderAlreadyPaid       = ErrorCode(400710)
	OrderNotFound          = ErrorCode(400711)
	OrderNotConfirmed      = ErrorCode(400712)
	OrderCustomerNotSame   = ErrorCode(400713)
	OrderAlreadyHasPayment = ErrorCode(400714)
	OrderAlreadyReturned   = ErrorCode(400717)
	OrderNotPaid           = ErrorCode(400718)
	OrderDetailNotFound    = ErrorCode(400719)

	InventoryNotUniqueCode = ErrorCode(400715)
	InventoryIsUniqueCode  = ErrorCode(400716)

	TooManyRequest          = ErrorCode(400200)
	LoginSessionExpired     = ErrorCode(400100)
	Login3rdSourceNotFound  = ErrorCode(400101)
	Login3rdBindExist       = ErrorCode(400102)
	NotFound                = ErrorCode(404001)
	InviteCodeAlreadySetted = ErrorCode(400302)
	InviteCodeNotValid      = ErrorCode(400303)
	PromotionNotValid       = ErrorCode(400301)
	CreditActionNotValid    = ErrorCode(400401)

	RequestReachLimit = ErrorCode(429001)

	InternalErrorHappened       = ErrorCode(500001)
	TransactionAlreadyProcessed = ErrorCode(600001)

	NotValidFile = ErrorCode(400405)
)

var mapping *sync.Map
var codeMessages = map[ErrorCode]string{
	AuthNeedLogin:             "401001: Please login",
	AuthInvalidToken:          "401002: Invalid token",
	NoPermission:              "401003: No permission to check resource",
	UserNamePasswordInCorrect: "401004: Username or password incorrect",
	NoRefreshCodeFound:        "401005: Refresh token is not found or invalid",
	InvalidRefreshCode:        "401006: Refresh token is invalid",
	PortfolioIsLocked:         "403007: Portfolio is locked",

	ParameterInvalid:     "400001: Invalid parameters",
	ValidationErrors:     "400002: Not validate",
	NoExchangeSelect:     "400003: Please select at least one exchange",
	DuplicateResource:    "400004: Already exists",
	ArticleCanNotComment: "400005: Article can not be commented",
	PasswordNotValid:     "400006: Password is not valid",

	OrderTypeInvalid:       "400708: Order type is not valid",
	OrderAlreadyConfirmed:  "400709: Order already confirmed",
	OrderAlreadyPaid:       "400710: Order already paid",
	OrderNotFound:          "400711: Order not found",
	OrderNotConfirmed:      "400712: Order not confirmed",
	OrderCustomerNotSame:   "400713: Order customer not same",
	OrderAlreadyHasPayment: "400714: Order already has payment",
	OrderAlreadyReturned:   "400717: Order already returned",
	OrderNotPaid:           "400718: Order not paid",
	OrderDetailNotFound:    "400719: Order detail not found",

	InventoryNotUniqueCode: "400715: Product is not unique code",
	InventoryIsUniqueCode:  "400716: Product is unique code",

	LoginSessionExpired:         "400100: Login Session Expired",
	Login3rdSourceNotFound:      "400101: Third party source not found",
	Login3rdBindExist:           "400102: Third party account already be taken or user already bind an account",
	NotFound:                    "404001: Resource not found",
	InternalErrorHappened:       "500001: There are internal error",
	TransactionAlreadyProcessed: "600001: Transaction already processed",
	TooManyRequest:              "400200: Too many request sent",
	CreditNotEnough:             "400007: credit not enough",
	PromotionNotValid:           "400301: Promotion code not valid",
	CreditActionNotValid:        "400401: Credit action not valid",
	InviteCodeAlreadySetted:     "400302: Already setted invite code",
	NotValidFile:                "400405: Not a valid file",
	RequestReachLimit:           "429001: Reached request limit",
}

func init() {
	mapping = &sync.Map{}
	for k, v := range codeMessages {
		Register(k, v)
	}
}

func Register(errCode ErrorCode, template string) {
	mapping.Store(errCode, template)
}

func (errCode ErrorCode) Error(err ...error) *HTTPError {
	status := int(errCode) / 1000
	e := NewHTTPError(errCode, status, err...)

	if v, ok := mapping.Load(errCode); ok {
		e.Message = v.(string)
	} else if e.error != nil {
		e.Message = e.error.Error()
	} else {
		e.Message = "unknown error"
	}
	if status == 500 {
		e.Log = true
	}
	return e
}
