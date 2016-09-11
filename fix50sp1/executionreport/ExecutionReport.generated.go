package executionreport

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//ExecutionReport is the fix50sp1 ExecutionReport type, MsgType = 8
type ExecutionReport struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ExecutionReport from a quickfix.Message instance
func FromMessage(m quickfix.Message) ExecutionReport {
	return ExecutionReport{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ExecutionReport) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ExecutionReport initialized with the required fields for ExecutionReport
func New(orderid field.OrderIDField, execid field.ExecIDField, exectype field.ExecTypeField, ordstatus field.OrdStatusField, side field.SideField, leavesqty field.LeavesQtyField, cumqty field.CumQtyField) (m ExecutionReport) {
	m.Header = fixt11.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("8"))
	m.Set(orderid)
	m.Set(execid)
	m.Set(exectype)
	m.Set(ordstatus)
	m.Set(side)
	m.Set(leavesqty)
	m.Set(cumqty)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ExecutionReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "8", r
}

//SetAccount sets Account, Tag 1
func (m ExecutionReport) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetAvgPx sets AvgPx, Tag 6
func (m ExecutionReport) SetAvgPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewAvgPx(value, scale))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m ExecutionReport) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetCommission sets Commission, Tag 12
func (m ExecutionReport) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m ExecutionReport) SetCommType(v enum.CommType) {
	m.Set(field.NewCommType(v))
}

//SetCumQty sets CumQty, Tag 14
func (m ExecutionReport) SetCumQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCumQty(value, scale))
}

//SetCurrency sets Currency, Tag 15
func (m ExecutionReport) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecID sets ExecID, Tag 17
func (m ExecutionReport) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m ExecutionReport) SetExecInst(v enum.ExecInst) {
	m.Set(field.NewExecInst(v))
}

//SetExecRefID sets ExecRefID, Tag 19
func (m ExecutionReport) SetExecRefID(v string) {
	m.Set(field.NewExecRefID(v))
}

//SetHandlInst sets HandlInst, Tag 21
func (m ExecutionReport) SetHandlInst(v enum.HandlInst) {
	m.Set(field.NewHandlInst(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m ExecutionReport) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetLastCapacity sets LastCapacity, Tag 29
func (m ExecutionReport) SetLastCapacity(v enum.LastCapacity) {
	m.Set(field.NewLastCapacity(v))
}

//SetLastMkt sets LastMkt, Tag 30
func (m ExecutionReport) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//SetLastPx sets LastPx, Tag 31
func (m ExecutionReport) SetLastPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastPx(value, scale))
}

//SetLastQty sets LastQty, Tag 32
func (m ExecutionReport) SetLastQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastQty(value, scale))
}

//SetOrderID sets OrderID, Tag 37
func (m ExecutionReport) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m ExecutionReport) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

//SetOrdStatus sets OrdStatus, Tag 39
func (m ExecutionReport) SetOrdStatus(v enum.OrdStatus) {
	m.Set(field.NewOrdStatus(v))
}

//SetOrdType sets OrdType, Tag 40
func (m ExecutionReport) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

//SetOrigClOrdID sets OrigClOrdID, Tag 41
func (m ExecutionReport) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

//SetPrice sets Price, Tag 44
func (m ExecutionReport) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetSecurityID sets SecurityID, Tag 48
func (m ExecutionReport) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m ExecutionReport) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m ExecutionReport) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m ExecutionReport) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m ExecutionReport) SetTimeInForce(v enum.TimeInForce) {
	m.Set(field.NewTimeInForce(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m ExecutionReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlType sets SettlType, Tag 63
func (m ExecutionReport) SetSettlType(v enum.SettlType) {
	m.Set(field.NewSettlType(v))
}

//SetSettlDate sets SettlDate, Tag 64
func (m ExecutionReport) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m ExecutionReport) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetListID sets ListID, Tag 66
func (m ExecutionReport) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetAllocID sets AllocID, Tag 70
func (m ExecutionReport) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m ExecutionReport) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetPositionEffect sets PositionEffect, Tag 77
func (m ExecutionReport) SetPositionEffect(v enum.PositionEffect) {
	m.Set(field.NewPositionEffect(v))
}

//SetNoAllocs sets NoAllocs, Tag 78
func (m ExecutionReport) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetStopPx sets StopPx, Tag 99
func (m ExecutionReport) SetStopPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewStopPx(value, scale))
}

//SetOrdRejReason sets OrdRejReason, Tag 103
func (m ExecutionReport) SetOrdRejReason(v enum.OrdRejReason) {
	m.Set(field.NewOrdRejReason(v))
}

//SetIssuer sets Issuer, Tag 106
func (m ExecutionReport) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m ExecutionReport) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetMinQty sets MinQty, Tag 110
func (m ExecutionReport) SetMinQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinQty(value, scale))
}

//SetMaxFloor sets MaxFloor, Tag 111
func (m ExecutionReport) SetMaxFloor(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxFloor(value, scale))
}

//SetReportToExch sets ReportToExch, Tag 113
func (m ExecutionReport) SetReportToExch(v bool) {
	m.Set(field.NewReportToExch(v))
}

//SetNetMoney sets NetMoney, Tag 118
func (m ExecutionReport) SetNetMoney(value decimal.Decimal, scale int32) {
	m.Set(field.NewNetMoney(value, scale))
}

//SetSettlCurrAmt sets SettlCurrAmt, Tag 119
func (m ExecutionReport) SetSettlCurrAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlCurrAmt(value, scale))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m ExecutionReport) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m ExecutionReport) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetNoMiscFees sets NoMiscFees, Tag 136
func (m ExecutionReport) SetNoMiscFees(f NoMiscFeesRepeatingGroup) {
	m.SetGroup(f)
}

//SetExecType sets ExecType, Tag 150
func (m ExecutionReport) SetExecType(v enum.ExecType) {
	m.Set(field.NewExecType(v))
}

//SetLeavesQty sets LeavesQty, Tag 151
func (m ExecutionReport) SetLeavesQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLeavesQty(value, scale))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m ExecutionReport) SetCashOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOrderQty(value, scale))
}

//SetSettlCurrFxRate sets SettlCurrFxRate, Tag 155
func (m ExecutionReport) SetSettlCurrFxRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlCurrFxRate(value, scale))
}

//SetSettlCurrFxRateCalc sets SettlCurrFxRateCalc, Tag 156
func (m ExecutionReport) SetSettlCurrFxRateCalc(v enum.SettlCurrFxRateCalc) {
	m.Set(field.NewSettlCurrFxRateCalc(v))
}

//SetNumDaysInterest sets NumDaysInterest, Tag 157
func (m ExecutionReport) SetNumDaysInterest(v int) {
	m.Set(field.NewNumDaysInterest(v))
}

//SetAccruedInterestRate sets AccruedInterestRate, Tag 158
func (m ExecutionReport) SetAccruedInterestRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewAccruedInterestRate(value, scale))
}

//SetAccruedInterestAmt sets AccruedInterestAmt, Tag 159
func (m ExecutionReport) SetAccruedInterestAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewAccruedInterestAmt(value, scale))
}

//SetSecurityType sets SecurityType, Tag 167
func (m ExecutionReport) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m ExecutionReport) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m ExecutionReport) SetOrderQty2(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty2(value, scale))
}

//SetSettlDate2 sets SettlDate2, Tag 193
func (m ExecutionReport) SetSettlDate2(v string) {
	m.Set(field.NewSettlDate2(v))
}

//SetLastSpotRate sets LastSpotRate, Tag 194
func (m ExecutionReport) SetLastSpotRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastSpotRate(value, scale))
}

//SetLastForwardPoints sets LastForwardPoints, Tag 195
func (m ExecutionReport) SetLastForwardPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastForwardPoints(value, scale))
}

//SetSecondaryOrderID sets SecondaryOrderID, Tag 198
func (m ExecutionReport) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m ExecutionReport) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m ExecutionReport) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m ExecutionReport) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m ExecutionReport) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m ExecutionReport) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetMaxShow sets MaxShow, Tag 210
func (m ExecutionReport) SetMaxShow(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxShow(value, scale))
}

//SetPegOffsetValue sets PegOffsetValue, Tag 211
func (m ExecutionReport) SetPegOffsetValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewPegOffsetValue(value, scale))
}

//SetSpread sets Spread, Tag 218
func (m ExecutionReport) SetSpread(value decimal.Decimal, scale int32) {
	m.Set(field.NewSpread(value, scale))
}

//SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220
func (m ExecutionReport) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

//SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221
func (m ExecutionReport) SetBenchmarkCurveName(v enum.BenchmarkCurveName) {
	m.Set(field.NewBenchmarkCurveName(v))
}

//SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222
func (m ExecutionReport) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m ExecutionReport) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m ExecutionReport) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m ExecutionReport) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m ExecutionReport) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m ExecutionReport) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m ExecutionReport) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetTradeOriginationDate sets TradeOriginationDate, Tag 229
func (m ExecutionReport) SetTradeOriginationDate(v string) {
	m.Set(field.NewTradeOriginationDate(v))
}

//SetExDate sets ExDate, Tag 230
func (m ExecutionReport) SetExDate(v string) {
	m.Set(field.NewExDate(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m ExecutionReport) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetNoStipulations sets NoStipulations, Tag 232
func (m ExecutionReport) SetNoStipulations(f NoStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

//SetYieldType sets YieldType, Tag 235
func (m ExecutionReport) SetYieldType(v enum.YieldType) {
	m.Set(field.NewYieldType(v))
}

//SetYield sets Yield, Tag 236
func (m ExecutionReport) SetYield(value decimal.Decimal, scale int32) {
	m.Set(field.NewYield(value, scale))
}

//SetTotalTakedown sets TotalTakedown, Tag 237
func (m ExecutionReport) SetTotalTakedown(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalTakedown(value, scale))
}

//SetConcession sets Concession, Tag 238
func (m ExecutionReport) SetConcession(value decimal.Decimal, scale int32) {
	m.Set(field.NewConcession(value, scale))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m ExecutionReport) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m ExecutionReport) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m ExecutionReport) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetTradedFlatSwitch sets TradedFlatSwitch, Tag 258
func (m ExecutionReport) SetTradedFlatSwitch(v bool) {
	m.Set(field.NewTradedFlatSwitch(v))
}

//SetBasisFeatureDate sets BasisFeatureDate, Tag 259
func (m ExecutionReport) SetBasisFeatureDate(v string) {
	m.Set(field.NewBasisFeatureDate(v))
}

//SetBasisFeaturePrice sets BasisFeaturePrice, Tag 260
func (m ExecutionReport) SetBasisFeaturePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewBasisFeaturePrice(value, scale))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m ExecutionReport) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m ExecutionReport) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m ExecutionReport) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m ExecutionReport) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m ExecutionReport) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m ExecutionReport) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m ExecutionReport) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetComplianceID sets ComplianceID, Tag 376
func (m ExecutionReport) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

//SetSolicitedFlag sets SolicitedFlag, Tag 377
func (m ExecutionReport) SetSolicitedFlag(v bool) {
	m.Set(field.NewSolicitedFlag(v))
}

//SetExecRestatementReason sets ExecRestatementReason, Tag 378
func (m ExecutionReport) SetExecRestatementReason(v enum.ExecRestatementReason) {
	m.Set(field.NewExecRestatementReason(v))
}

//SetGrossTradeAmt sets GrossTradeAmt, Tag 381
func (m ExecutionReport) SetGrossTradeAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewGrossTradeAmt(value, scale))
}

//SetNoContraBrokers sets NoContraBrokers, Tag 382
func (m ExecutionReport) SetNoContraBrokers(f NoContraBrokersRepeatingGroup) {
	m.SetGroup(f)
}

//SetDiscretionInst sets DiscretionInst, Tag 388
func (m ExecutionReport) SetDiscretionInst(v enum.DiscretionInst) {
	m.Set(field.NewDiscretionInst(v))
}

//SetDiscretionOffsetValue sets DiscretionOffsetValue, Tag 389
func (m ExecutionReport) SetDiscretionOffsetValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewDiscretionOffsetValue(value, scale))
}

//SetPriceType sets PriceType, Tag 423
func (m ExecutionReport) SetPriceType(v enum.PriceType) {
	m.Set(field.NewPriceType(v))
}

//SetDayOrderQty sets DayOrderQty, Tag 424
func (m ExecutionReport) SetDayOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDayOrderQty(value, scale))
}

//SetDayCumQty sets DayCumQty, Tag 425
func (m ExecutionReport) SetDayCumQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDayCumQty(value, scale))
}

//SetDayAvgPx sets DayAvgPx, Tag 426
func (m ExecutionReport) SetDayAvgPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewDayAvgPx(value, scale))
}

//SetGTBookingInst sets GTBookingInst, Tag 427
func (m ExecutionReport) SetGTBookingInst(v enum.GTBookingInst) {
	m.Set(field.NewGTBookingInst(v))
}

//SetExpireDate sets ExpireDate, Tag 432
func (m ExecutionReport) SetExpireDate(v string) {
	m.Set(field.NewExpireDate(v))
}

//SetMultiLegReportingType sets MultiLegReportingType, Tag 442
func (m ExecutionReport) SetMultiLegReportingType(v enum.MultiLegReportingType) {
	m.Set(field.NewMultiLegReportingType(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m ExecutionReport) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m ExecutionReport) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m ExecutionReport) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m ExecutionReport) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m ExecutionReport) SetRoundingDirection(v enum.RoundingDirection) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m ExecutionReport) SetRoundingModulus(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundingModulus(value, scale))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m ExecutionReport) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m ExecutionReport) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m ExecutionReport) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetCommCurrency sets CommCurrency, Tag 479
func (m ExecutionReport) SetCommCurrency(v string) {
	m.Set(field.NewCommCurrency(v))
}

//SetCancellationRights sets CancellationRights, Tag 480
func (m ExecutionReport) SetCancellationRights(v enum.CancellationRights) {
	m.Set(field.NewCancellationRights(v))
}

//SetMoneyLaunderingStatus sets MoneyLaunderingStatus, Tag 481
func (m ExecutionReport) SetMoneyLaunderingStatus(v enum.MoneyLaunderingStatus) {
	m.Set(field.NewMoneyLaunderingStatus(v))
}

//SetTransBkdTime sets TransBkdTime, Tag 483
func (m ExecutionReport) SetTransBkdTime(v time.Time) {
	m.Set(field.NewTransBkdTime(v))
}

//SetExecPriceType sets ExecPriceType, Tag 484
func (m ExecutionReport) SetExecPriceType(v enum.ExecPriceType) {
	m.Set(field.NewExecPriceType(v))
}

//SetExecPriceAdjustment sets ExecPriceAdjustment, Tag 485
func (m ExecutionReport) SetExecPriceAdjustment(value decimal.Decimal, scale int32) {
	m.Set(field.NewExecPriceAdjustment(value, scale))
}

//SetDesignation sets Designation, Tag 494
func (m ExecutionReport) SetDesignation(v string) {
	m.Set(field.NewDesignation(v))
}

//SetFundRenewWaiv sets FundRenewWaiv, Tag 497
func (m ExecutionReport) SetFundRenewWaiv(v enum.FundRenewWaiv) {
	m.Set(field.NewFundRenewWaiv(v))
}

//SetRegistID sets RegistID, Tag 513
func (m ExecutionReport) SetRegistID(v string) {
	m.Set(field.NewRegistID(v))
}

//SetExecValuationPoint sets ExecValuationPoint, Tag 515
func (m ExecutionReport) SetExecValuationPoint(v time.Time) {
	m.Set(field.NewExecValuationPoint(v))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m ExecutionReport) SetOrderPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderPercent(value, scale))
}

//SetNoContAmts sets NoContAmts, Tag 518
func (m ExecutionReport) SetNoContAmts(f NoContAmtsRepeatingGroup) {
	m.SetGroup(f)
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m ExecutionReport) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetSecondaryExecID sets SecondaryExecID, Tag 527
func (m ExecutionReport) SetSecondaryExecID(v string) {
	m.Set(field.NewSecondaryExecID(v))
}

//SetOrderCapacity sets OrderCapacity, Tag 528
func (m ExecutionReport) SetOrderCapacity(v enum.OrderCapacity) {
	m.Set(field.NewOrderCapacity(v))
}

//SetOrderRestrictions sets OrderRestrictions, Tag 529
func (m ExecutionReport) SetOrderRestrictions(v enum.OrderRestrictions) {
	m.Set(field.NewOrderRestrictions(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m ExecutionReport) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m ExecutionReport) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

//SetCashMargin sets CashMargin, Tag 544
func (m ExecutionReport) SetCashMargin(v enum.CashMargin) {
	m.Set(field.NewCashMargin(v))
}

//SetCrossID sets CrossID, Tag 548
func (m ExecutionReport) SetCrossID(v string) {
	m.Set(field.NewCrossID(v))
}

//SetCrossType sets CrossType, Tag 549
func (m ExecutionReport) SetCrossType(v enum.CrossType) {
	m.Set(field.NewCrossType(v))
}

//SetOrigCrossID sets OrigCrossID, Tag 551
func (m ExecutionReport) SetOrigCrossID(v string) {
	m.Set(field.NewOrigCrossID(v))
}

//SetNoLegs sets NoLegs, Tag 555
func (m ExecutionReport) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

//SetMatchType sets MatchType, Tag 574
func (m ExecutionReport) SetMatchType(v enum.MatchType) {
	m.Set(field.NewMatchType(v))
}

//SetAccountType sets AccountType, Tag 581
func (m ExecutionReport) SetAccountType(v enum.AccountType) {
	m.Set(field.NewAccountType(v))
}

//SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m ExecutionReport) SetCustOrderCapacity(v enum.CustOrderCapacity) {
	m.Set(field.NewCustOrderCapacity(v))
}

//SetClOrdLinkID sets ClOrdLinkID, Tag 583
func (m ExecutionReport) SetClOrdLinkID(v string) {
	m.Set(field.NewClOrdLinkID(v))
}

//SetMassStatusReqID sets MassStatusReqID, Tag 584
func (m ExecutionReport) SetMassStatusReqID(v string) {
	m.Set(field.NewMassStatusReqID(v))
}

//SetDayBookingInst sets DayBookingInst, Tag 589
func (m ExecutionReport) SetDayBookingInst(v enum.DayBookingInst) {
	m.Set(field.NewDayBookingInst(v))
}

//SetBookingUnit sets BookingUnit, Tag 590
func (m ExecutionReport) SetBookingUnit(v enum.BookingUnit) {
	m.Set(field.NewBookingUnit(v))
}

//SetPreallocMethod sets PreallocMethod, Tag 591
func (m ExecutionReport) SetPreallocMethod(v enum.PreallocMethod) {
	m.Set(field.NewPreallocMethod(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m ExecutionReport) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetClearingFeeIndicator sets ClearingFeeIndicator, Tag 635
func (m ExecutionReport) SetClearingFeeIndicator(v enum.ClearingFeeIndicator) {
	m.Set(field.NewClearingFeeIndicator(v))
}

//SetWorkingIndicator sets WorkingIndicator, Tag 636
func (m ExecutionReport) SetWorkingIndicator(v bool) {
	m.Set(field.NewWorkingIndicator(v))
}

//SetPriorityIndicator sets PriorityIndicator, Tag 638
func (m ExecutionReport) SetPriorityIndicator(v enum.PriorityIndicator) {
	m.Set(field.NewPriorityIndicator(v))
}

//SetPriceImprovement sets PriceImprovement, Tag 639
func (m ExecutionReport) SetPriceImprovement(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceImprovement(value, scale))
}

//SetLastForwardPoints2 sets LastForwardPoints2, Tag 641
func (m ExecutionReport) SetLastForwardPoints2(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastForwardPoints2(value, scale))
}

//SetUnderlyingLastPx sets UnderlyingLastPx, Tag 651
func (m ExecutionReport) SetUnderlyingLastPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingLastPx(value, scale))
}

//SetUnderlyingLastQty sets UnderlyingLastQty, Tag 652
func (m ExecutionReport) SetUnderlyingLastQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingLastQty(value, scale))
}

//SetAcctIDSource sets AcctIDSource, Tag 660
func (m ExecutionReport) SetAcctIDSource(v enum.AcctIDSource) {
	m.Set(field.NewAcctIDSource(v))
}

//SetBenchmarkPrice sets BenchmarkPrice, Tag 662
func (m ExecutionReport) SetBenchmarkPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewBenchmarkPrice(value, scale))
}

//SetBenchmarkPriceType sets BenchmarkPriceType, Tag 663
func (m ExecutionReport) SetBenchmarkPriceType(v int) {
	m.Set(field.NewBenchmarkPriceType(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m ExecutionReport) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetLastParPx sets LastParPx, Tag 669
func (m ExecutionReport) SetLastParPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastParPx(value, scale))
}

//SetPool sets Pool, Tag 691
func (m ExecutionReport) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetQuoteRespID sets QuoteRespID, Tag 693
func (m ExecutionReport) SetQuoteRespID(v string) {
	m.Set(field.NewQuoteRespID(v))
}

//SetYieldRedemptionDate sets YieldRedemptionDate, Tag 696
func (m ExecutionReport) SetYieldRedemptionDate(v string) {
	m.Set(field.NewYieldRedemptionDate(v))
}

//SetYieldRedemptionPrice sets YieldRedemptionPrice, Tag 697
func (m ExecutionReport) SetYieldRedemptionPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewYieldRedemptionPrice(value, scale))
}

//SetYieldRedemptionPriceType sets YieldRedemptionPriceType, Tag 698
func (m ExecutionReport) SetYieldRedemptionPriceType(v int) {
	m.Set(field.NewYieldRedemptionPriceType(v))
}

//SetBenchmarkSecurityID sets BenchmarkSecurityID, Tag 699
func (m ExecutionReport) SetBenchmarkSecurityID(v string) {
	m.Set(field.NewBenchmarkSecurityID(v))
}

//SetYieldCalcDate sets YieldCalcDate, Tag 701
func (m ExecutionReport) SetYieldCalcDate(v string) {
	m.Set(field.NewYieldCalcDate(v))
}

//SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m ExecutionReport) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetInterestAtMaturity sets InterestAtMaturity, Tag 738
func (m ExecutionReport) SetInterestAtMaturity(value decimal.Decimal, scale int32) {
	m.Set(field.NewInterestAtMaturity(value, scale))
}

//SetBenchmarkSecurityIDSource sets BenchmarkSecurityIDSource, Tag 761
func (m ExecutionReport) SetBenchmarkSecurityIDSource(v string) {
	m.Set(field.NewBenchmarkSecurityIDSource(v))
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m ExecutionReport) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetNoTrdRegTimestamps sets NoTrdRegTimestamps, Tag 768
func (m ExecutionReport) SetNoTrdRegTimestamps(f NoTrdRegTimestampsRepeatingGroup) {
	m.SetGroup(f)
}

//SetBookingType sets BookingType, Tag 775
func (m ExecutionReport) SetBookingType(v enum.BookingType) {
	m.Set(field.NewBookingType(v))
}

//SetTerminationType sets TerminationType, Tag 788
func (m ExecutionReport) SetTerminationType(v enum.TerminationType) {
	m.Set(field.NewTerminationType(v))
}

//SetOrdStatusReqID sets OrdStatusReqID, Tag 790
func (m ExecutionReport) SetOrdStatusReqID(v string) {
	m.Set(field.NewOrdStatusReqID(v))
}

//SetCopyMsgIndicator sets CopyMsgIndicator, Tag 797
func (m ExecutionReport) SetCopyMsgIndicator(v bool) {
	m.Set(field.NewCopyMsgIndicator(v))
}

//SetPriceDelta sets PriceDelta, Tag 811
func (m ExecutionReport) SetPriceDelta(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceDelta(value, scale))
}

//SetPegMoveType sets PegMoveType, Tag 835
func (m ExecutionReport) SetPegMoveType(v enum.PegMoveType) {
	m.Set(field.NewPegMoveType(v))
}

//SetPegOffsetType sets PegOffsetType, Tag 836
func (m ExecutionReport) SetPegOffsetType(v enum.PegOffsetType) {
	m.Set(field.NewPegOffsetType(v))
}

//SetPegLimitType sets PegLimitType, Tag 837
func (m ExecutionReport) SetPegLimitType(v enum.PegLimitType) {
	m.Set(field.NewPegLimitType(v))
}

//SetPegRoundDirection sets PegRoundDirection, Tag 838
func (m ExecutionReport) SetPegRoundDirection(v enum.PegRoundDirection) {
	m.Set(field.NewPegRoundDirection(v))
}

//SetPeggedPrice sets PeggedPrice, Tag 839
func (m ExecutionReport) SetPeggedPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPeggedPrice(value, scale))
}

//SetPegScope sets PegScope, Tag 840
func (m ExecutionReport) SetPegScope(v enum.PegScope) {
	m.Set(field.NewPegScope(v))
}

//SetDiscretionMoveType sets DiscretionMoveType, Tag 841
func (m ExecutionReport) SetDiscretionMoveType(v enum.DiscretionMoveType) {
	m.Set(field.NewDiscretionMoveType(v))
}

//SetDiscretionOffsetType sets DiscretionOffsetType, Tag 842
func (m ExecutionReport) SetDiscretionOffsetType(v enum.DiscretionOffsetType) {
	m.Set(field.NewDiscretionOffsetType(v))
}

//SetDiscretionLimitType sets DiscretionLimitType, Tag 843
func (m ExecutionReport) SetDiscretionLimitType(v enum.DiscretionLimitType) {
	m.Set(field.NewDiscretionLimitType(v))
}

//SetDiscretionRoundDirection sets DiscretionRoundDirection, Tag 844
func (m ExecutionReport) SetDiscretionRoundDirection(v enum.DiscretionRoundDirection) {
	m.Set(field.NewDiscretionRoundDirection(v))
}

//SetDiscretionPrice sets DiscretionPrice, Tag 845
func (m ExecutionReport) SetDiscretionPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewDiscretionPrice(value, scale))
}

//SetDiscretionScope sets DiscretionScope, Tag 846
func (m ExecutionReport) SetDiscretionScope(v enum.DiscretionScope) {
	m.Set(field.NewDiscretionScope(v))
}

//SetTargetStrategy sets TargetStrategy, Tag 847
func (m ExecutionReport) SetTargetStrategy(v enum.TargetStrategy) {
	m.Set(field.NewTargetStrategy(v))
}

//SetTargetStrategyParameters sets TargetStrategyParameters, Tag 848
func (m ExecutionReport) SetTargetStrategyParameters(v string) {
	m.Set(field.NewTargetStrategyParameters(v))
}

//SetParticipationRate sets ParticipationRate, Tag 849
func (m ExecutionReport) SetParticipationRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewParticipationRate(value, scale))
}

//SetTargetStrategyPerformance sets TargetStrategyPerformance, Tag 850
func (m ExecutionReport) SetTargetStrategyPerformance(value decimal.Decimal, scale int32) {
	m.Set(field.NewTargetStrategyPerformance(value, scale))
}

//SetLastLiquidityInd sets LastLiquidityInd, Tag 851
func (m ExecutionReport) SetLastLiquidityInd(v enum.LastLiquidityInd) {
	m.Set(field.NewLastLiquidityInd(v))
}

//SetQtyType sets QtyType, Tag 854
func (m ExecutionReport) SetQtyType(v enum.QtyType) {
	m.Set(field.NewQtyType(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m ExecutionReport) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m ExecutionReport) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m ExecutionReport) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m ExecutionReport) SetCPProgram(v enum.CPProgram) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m ExecutionReport) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetTrdMatchID sets TrdMatchID, Tag 880
func (m ExecutionReport) SetTrdMatchID(v string) {
	m.Set(field.NewTrdMatchID(v))
}

//SetLastFragment sets LastFragment, Tag 893
func (m ExecutionReport) SetLastFragment(v bool) {
	m.Set(field.NewLastFragment(v))
}

//SetMarginRatio sets MarginRatio, Tag 898
func (m ExecutionReport) SetMarginRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginRatio(value, scale))
}

//SetTotNumReports sets TotNumReports, Tag 911
func (m ExecutionReport) SetTotNumReports(v int) {
	m.Set(field.NewTotNumReports(v))
}

//SetLastRptRequested sets LastRptRequested, Tag 912
func (m ExecutionReport) SetLastRptRequested(v bool) {
	m.Set(field.NewLastRptRequested(v))
}

//SetAgreementDesc sets AgreementDesc, Tag 913
func (m ExecutionReport) SetAgreementDesc(v string) {
	m.Set(field.NewAgreementDesc(v))
}

//SetAgreementID sets AgreementID, Tag 914
func (m ExecutionReport) SetAgreementID(v string) {
	m.Set(field.NewAgreementID(v))
}

//SetAgreementDate sets AgreementDate, Tag 915
func (m ExecutionReport) SetAgreementDate(v string) {
	m.Set(field.NewAgreementDate(v))
}

//SetStartDate sets StartDate, Tag 916
func (m ExecutionReport) SetStartDate(v string) {
	m.Set(field.NewStartDate(v))
}

//SetEndDate sets EndDate, Tag 917
func (m ExecutionReport) SetEndDate(v string) {
	m.Set(field.NewEndDate(v))
}

//SetAgreementCurrency sets AgreementCurrency, Tag 918
func (m ExecutionReport) SetAgreementCurrency(v string) {
	m.Set(field.NewAgreementCurrency(v))
}

//SetDeliveryType sets DeliveryType, Tag 919
func (m ExecutionReport) SetDeliveryType(v enum.DeliveryType) {
	m.Set(field.NewDeliveryType(v))
}

//SetEndAccruedInterestAmt sets EndAccruedInterestAmt, Tag 920
func (m ExecutionReport) SetEndAccruedInterestAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewEndAccruedInterestAmt(value, scale))
}

//SetStartCash sets StartCash, Tag 921
func (m ExecutionReport) SetStartCash(value decimal.Decimal, scale int32) {
	m.Set(field.NewStartCash(value, scale))
}

//SetEndCash sets EndCash, Tag 922
func (m ExecutionReport) SetEndCash(value decimal.Decimal, scale int32) {
	m.Set(field.NewEndCash(value, scale))
}

//SetTimeBracket sets TimeBracket, Tag 943
func (m ExecutionReport) SetTimeBracket(v string) {
	m.Set(field.NewTimeBracket(v))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m ExecutionReport) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//SetNoStrategyParameters sets NoStrategyParameters, Tag 957
func (m ExecutionReport) SetNoStrategyParameters(f NoStrategyParametersRepeatingGroup) {
	m.SetGroup(f)
}

//SetHostCrossID sets HostCrossID, Tag 961
func (m ExecutionReport) SetHostCrossID(v string) {
	m.Set(field.NewHostCrossID(v))
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m ExecutionReport) SetSecurityStatus(v enum.SecurityStatus) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m ExecutionReport) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m ExecutionReport) SetStrikeMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeMultiplier(value, scale))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m ExecutionReport) SetStrikeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeValue(value, scale))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m ExecutionReport) SetMinPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrement(value, scale))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m ExecutionReport) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m ExecutionReport) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m ExecutionReport) SetUnitOfMeasure(v enum.UnitOfMeasure) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m ExecutionReport) SetTimeUnit(v enum.TimeUnit) {
	m.Set(field.NewTimeUnit(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m ExecutionReport) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetManualOrderIndicator sets ManualOrderIndicator, Tag 1028
func (m ExecutionReport) SetManualOrderIndicator(v bool) {
	m.Set(field.NewManualOrderIndicator(v))
}

//SetCustDirectedOrder sets CustDirectedOrder, Tag 1029
func (m ExecutionReport) SetCustDirectedOrder(v bool) {
	m.Set(field.NewCustDirectedOrder(v))
}

//SetReceivedDeptID sets ReceivedDeptID, Tag 1030
func (m ExecutionReport) SetReceivedDeptID(v string) {
	m.Set(field.NewReceivedDeptID(v))
}

//SetCustOrderHandlingInst sets CustOrderHandlingInst, Tag 1031
func (m ExecutionReport) SetCustOrderHandlingInst(v enum.CustOrderHandlingInst) {
	m.Set(field.NewCustOrderHandlingInst(v))
}

//SetOrderHandlingInstSource sets OrderHandlingInstSource, Tag 1032
func (m ExecutionReport) SetOrderHandlingInstSource(v enum.OrderHandlingInstSource) {
	m.Set(field.NewOrderHandlingInstSource(v))
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m ExecutionReport) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetCalculatedCcyLastQty sets CalculatedCcyLastQty, Tag 1056
func (m ExecutionReport) SetCalculatedCcyLastQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCalculatedCcyLastQty(value, scale))
}

//SetAggressorIndicator sets AggressorIndicator, Tag 1057
func (m ExecutionReport) SetAggressorIndicator(v bool) {
	m.Set(field.NewAggressorIndicator(v))
}

//SetLastSwapPoints sets LastSwapPoints, Tag 1071
func (m ExecutionReport) SetLastSwapPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastSwapPoints(value, scale))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m ExecutionReport) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//SetSecondaryDisplayQty sets SecondaryDisplayQty, Tag 1082
func (m ExecutionReport) SetSecondaryDisplayQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewSecondaryDisplayQty(value, scale))
}

//SetDisplayWhen sets DisplayWhen, Tag 1083
func (m ExecutionReport) SetDisplayWhen(v enum.DisplayWhen) {
	m.Set(field.NewDisplayWhen(v))
}

//SetDisplayMethod sets DisplayMethod, Tag 1084
func (m ExecutionReport) SetDisplayMethod(v enum.DisplayMethod) {
	m.Set(field.NewDisplayMethod(v))
}

//SetDisplayLowQty sets DisplayLowQty, Tag 1085
func (m ExecutionReport) SetDisplayLowQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDisplayLowQty(value, scale))
}

//SetDisplayHighQty sets DisplayHighQty, Tag 1086
func (m ExecutionReport) SetDisplayHighQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDisplayHighQty(value, scale))
}

//SetDisplayMinIncr sets DisplayMinIncr, Tag 1087
func (m ExecutionReport) SetDisplayMinIncr(value decimal.Decimal, scale int32) {
	m.Set(field.NewDisplayMinIncr(value, scale))
}

//SetRefreshQty sets RefreshQty, Tag 1088
func (m ExecutionReport) SetRefreshQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewRefreshQty(value, scale))
}

//SetMatchIncrement sets MatchIncrement, Tag 1089
func (m ExecutionReport) SetMatchIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewMatchIncrement(value, scale))
}

//SetMaxPriceLevels sets MaxPriceLevels, Tag 1090
func (m ExecutionReport) SetMaxPriceLevels(v int) {
	m.Set(field.NewMaxPriceLevels(v))
}

//SetPreTradeAnonymity sets PreTradeAnonymity, Tag 1091
func (m ExecutionReport) SetPreTradeAnonymity(v bool) {
	m.Set(field.NewPreTradeAnonymity(v))
}

//SetPriceProtectionScope sets PriceProtectionScope, Tag 1092
func (m ExecutionReport) SetPriceProtectionScope(v enum.PriceProtectionScope) {
	m.Set(field.NewPriceProtectionScope(v))
}

//SetLotType sets LotType, Tag 1093
func (m ExecutionReport) SetLotType(v enum.LotType) {
	m.Set(field.NewLotType(v))
}

//SetPegPriceType sets PegPriceType, Tag 1094
func (m ExecutionReport) SetPegPriceType(v enum.PegPriceType) {
	m.Set(field.NewPegPriceType(v))
}

//SetPeggedRefPrice sets PeggedRefPrice, Tag 1095
func (m ExecutionReport) SetPeggedRefPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPeggedRefPrice(value, scale))
}

//SetPegSecurityIDSource sets PegSecurityIDSource, Tag 1096
func (m ExecutionReport) SetPegSecurityIDSource(v string) {
	m.Set(field.NewPegSecurityIDSource(v))
}

//SetPegSecurityID sets PegSecurityID, Tag 1097
func (m ExecutionReport) SetPegSecurityID(v string) {
	m.Set(field.NewPegSecurityID(v))
}

//SetPegSymbol sets PegSymbol, Tag 1098
func (m ExecutionReport) SetPegSymbol(v string) {
	m.Set(field.NewPegSymbol(v))
}

//SetPegSecurityDesc sets PegSecurityDesc, Tag 1099
func (m ExecutionReport) SetPegSecurityDesc(v string) {
	m.Set(field.NewPegSecurityDesc(v))
}

//SetTriggerType sets TriggerType, Tag 1100
func (m ExecutionReport) SetTriggerType(v enum.TriggerType) {
	m.Set(field.NewTriggerType(v))
}

//SetTriggerAction sets TriggerAction, Tag 1101
func (m ExecutionReport) SetTriggerAction(v enum.TriggerAction) {
	m.Set(field.NewTriggerAction(v))
}

//SetTriggerPrice sets TriggerPrice, Tag 1102
func (m ExecutionReport) SetTriggerPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewTriggerPrice(value, scale))
}

//SetTriggerSymbol sets TriggerSymbol, Tag 1103
func (m ExecutionReport) SetTriggerSymbol(v string) {
	m.Set(field.NewTriggerSymbol(v))
}

//SetTriggerSecurityID sets TriggerSecurityID, Tag 1104
func (m ExecutionReport) SetTriggerSecurityID(v string) {
	m.Set(field.NewTriggerSecurityID(v))
}

//SetTriggerSecurityIDSource sets TriggerSecurityIDSource, Tag 1105
func (m ExecutionReport) SetTriggerSecurityIDSource(v string) {
	m.Set(field.NewTriggerSecurityIDSource(v))
}

//SetTriggerSecurityDesc sets TriggerSecurityDesc, Tag 1106
func (m ExecutionReport) SetTriggerSecurityDesc(v string) {
	m.Set(field.NewTriggerSecurityDesc(v))
}

//SetTriggerPriceType sets TriggerPriceType, Tag 1107
func (m ExecutionReport) SetTriggerPriceType(v enum.TriggerPriceType) {
	m.Set(field.NewTriggerPriceType(v))
}

//SetTriggerPriceTypeScope sets TriggerPriceTypeScope, Tag 1108
func (m ExecutionReport) SetTriggerPriceTypeScope(v enum.TriggerPriceTypeScope) {
	m.Set(field.NewTriggerPriceTypeScope(v))
}

//SetTriggerPriceDirection sets TriggerPriceDirection, Tag 1109
func (m ExecutionReport) SetTriggerPriceDirection(v enum.TriggerPriceDirection) {
	m.Set(field.NewTriggerPriceDirection(v))
}

//SetTriggerNewPrice sets TriggerNewPrice, Tag 1110
func (m ExecutionReport) SetTriggerNewPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewTriggerNewPrice(value, scale))
}

//SetTriggerOrderType sets TriggerOrderType, Tag 1111
func (m ExecutionReport) SetTriggerOrderType(v enum.TriggerOrderType) {
	m.Set(field.NewTriggerOrderType(v))
}

//SetTriggerNewQty sets TriggerNewQty, Tag 1112
func (m ExecutionReport) SetTriggerNewQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewTriggerNewQty(value, scale))
}

//SetTriggerTradingSessionID sets TriggerTradingSessionID, Tag 1113
func (m ExecutionReport) SetTriggerTradingSessionID(v string) {
	m.Set(field.NewTriggerTradingSessionID(v))
}

//SetTriggerTradingSessionSubID sets TriggerTradingSessionSubID, Tag 1114
func (m ExecutionReport) SetTriggerTradingSessionSubID(v string) {
	m.Set(field.NewTriggerTradingSessionSubID(v))
}

//SetOrderCategory sets OrderCategory, Tag 1115
func (m ExecutionReport) SetOrderCategory(v enum.OrderCategory) {
	m.Set(field.NewOrderCategory(v))
}

//SetDisplayQty sets DisplayQty, Tag 1138
func (m ExecutionReport) SetDisplayQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDisplayQty(value, scale))
}

//SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146
func (m ExecutionReport) SetMinPriceIncrementAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrementAmount(value, scale))
}

//SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147
func (m ExecutionReport) SetUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnitOfMeasureQty(value, scale))
}

//SetSecurityGroup sets SecurityGroup, Tag 1151
func (m ExecutionReport) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

//SetApplID sets ApplID, Tag 1180
func (m ExecutionReport) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

//SetApplSeqNum sets ApplSeqNum, Tag 1181
func (m ExecutionReport) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

//SetSecurityXMLLen sets SecurityXMLLen, Tag 1184
func (m ExecutionReport) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

//SetSecurityXML sets SecurityXML, Tag 1185
func (m ExecutionReport) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

//SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186
func (m ExecutionReport) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

//SetVolatility sets Volatility, Tag 1188
func (m ExecutionReport) SetVolatility(value decimal.Decimal, scale int32) {
	m.Set(field.NewVolatility(value, scale))
}

//SetTimeToExpiration sets TimeToExpiration, Tag 1189
func (m ExecutionReport) SetTimeToExpiration(value decimal.Decimal, scale int32) {
	m.Set(field.NewTimeToExpiration(value, scale))
}

//SetRiskFreeRate sets RiskFreeRate, Tag 1190
func (m ExecutionReport) SetRiskFreeRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRiskFreeRate(value, scale))
}

//SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191
func (m ExecutionReport) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

//SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192
func (m ExecutionReport) SetPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceUnitOfMeasureQty(value, scale))
}

//SetSettlMethod sets SettlMethod, Tag 1193
func (m ExecutionReport) SetSettlMethod(v enum.SettlMethod) {
	m.Set(field.NewSettlMethod(v))
}

//SetExerciseStyle sets ExerciseStyle, Tag 1194
func (m ExecutionReport) SetExerciseStyle(v enum.ExerciseStyle) {
	m.Set(field.NewExerciseStyle(v))
}

//SetOptPayAmount sets OptPayAmount, Tag 1195
func (m ExecutionReport) SetOptPayAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewOptPayAmount(value, scale))
}

//SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196
func (m ExecutionReport) SetPriceQuoteMethod(v enum.PriceQuoteMethod) {
	m.Set(field.NewPriceQuoteMethod(v))
}

//SetFuturesValuationMethod sets FuturesValuationMethod, Tag 1197
func (m ExecutionReport) SetFuturesValuationMethod(v enum.FuturesValuationMethod) {
	m.Set(field.NewFuturesValuationMethod(v))
}

//SetListMethod sets ListMethod, Tag 1198
func (m ExecutionReport) SetListMethod(v enum.ListMethod) {
	m.Set(field.NewListMethod(v))
}

//SetCapPrice sets CapPrice, Tag 1199
func (m ExecutionReport) SetCapPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewCapPrice(value, scale))
}

//SetFloorPrice sets FloorPrice, Tag 1200
func (m ExecutionReport) SetFloorPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewFloorPrice(value, scale))
}

//SetProductComplex sets ProductComplex, Tag 1227
func (m ExecutionReport) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

//SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242
func (m ExecutionReport) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

//SetFlexibleIndicator sets FlexibleIndicator, Tag 1244
func (m ExecutionReport) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

//SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350
func (m ExecutionReport) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

//SetApplResendFlag sets ApplResendFlag, Tag 1352
func (m ExecutionReport) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

//SetTotNoFills sets TotNoFills, Tag 1361
func (m ExecutionReport) SetTotNoFills(v int) {
	m.Set(field.NewTotNoFills(v))
}

//SetNoFills sets NoFills, Tag 1362
func (m ExecutionReport) SetNoFills(f NoFillsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDividendYield sets DividendYield, Tag 1380
func (m ExecutionReport) SetDividendYield(value decimal.Decimal, scale int32) {
	m.Set(field.NewDividendYield(value, scale))
}

//GetAccount gets Account, Tag 1
func (m ExecutionReport) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAvgPx gets AvgPx, Tag 6
func (m ExecutionReport) GetAvgPx() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.AvgPxField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m ExecutionReport) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommission gets Commission, Tag 12
func (m ExecutionReport) GetCommission() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.CommissionField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetCommType gets CommType, Tag 13
func (m ExecutionReport) GetCommType() (v enum.CommType, err quickfix.MessageRejectError) {
	var f field.CommTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCumQty gets CumQty, Tag 14
func (m ExecutionReport) GetCumQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.CumQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetCurrency gets Currency, Tag 15
func (m ExecutionReport) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecID gets ExecID, Tag 17
func (m ExecutionReport) GetExecID() (v string, err quickfix.MessageRejectError) {
	var f field.ExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m ExecutionReport) GetExecInst() (v enum.ExecInst, err quickfix.MessageRejectError) {
	var f field.ExecInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecRefID gets ExecRefID, Tag 19
func (m ExecutionReport) GetExecRefID() (v string, err quickfix.MessageRejectError) {
	var f field.ExecRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetHandlInst gets HandlInst, Tag 21
func (m ExecutionReport) GetHandlInst() (v enum.HandlInst, err quickfix.MessageRejectError) {
	var f field.HandlInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m ExecutionReport) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastCapacity gets LastCapacity, Tag 29
func (m ExecutionReport) GetLastCapacity() (v enum.LastCapacity, err quickfix.MessageRejectError) {
	var f field.LastCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastMkt gets LastMkt, Tag 30
func (m ExecutionReport) GetLastMkt() (v string, err quickfix.MessageRejectError) {
	var f field.LastMktField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastPx gets LastPx, Tag 31
func (m ExecutionReport) GetLastPx() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LastPxField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLastQty gets LastQty, Tag 32
func (m ExecutionReport) GetLastQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LastQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetOrderID gets OrderID, Tag 37
func (m ExecutionReport) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m ExecutionReport) GetOrderQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.OrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetOrdStatus gets OrdStatus, Tag 39
func (m ExecutionReport) GetOrdStatus() (v enum.OrdStatus, err quickfix.MessageRejectError) {
	var f field.OrdStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrdType gets OrdType, Tag 40
func (m ExecutionReport) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrigClOrdID gets OrigClOrdID, Tag 41
func (m ExecutionReport) GetOrigClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.OrigClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPrice gets Price, Tag 44
func (m ExecutionReport) GetPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m ExecutionReport) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSide gets Side, Tag 54
func (m ExecutionReport) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m ExecutionReport) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m ExecutionReport) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m ExecutionReport) GetTimeInForce() (v enum.TimeInForce, err quickfix.MessageRejectError) {
	var f field.TimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m ExecutionReport) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlType gets SettlType, Tag 63
func (m ExecutionReport) GetSettlType() (v enum.SettlType, err quickfix.MessageRejectError) {
	var f field.SettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlDate gets SettlDate, Tag 64
func (m ExecutionReport) GetSettlDate() (v string, err quickfix.MessageRejectError) {
	var f field.SettlDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m ExecutionReport) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetListID gets ListID, Tag 66
func (m ExecutionReport) GetListID() (v string, err quickfix.MessageRejectError) {
	var f field.ListIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocID gets AllocID, Tag 70
func (m ExecutionReport) GetAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m ExecutionReport) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPositionEffect gets PositionEffect, Tag 77
func (m ExecutionReport) GetPositionEffect() (v enum.PositionEffect, err quickfix.MessageRejectError) {
	var f field.PositionEffectField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m ExecutionReport) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetStopPx gets StopPx, Tag 99
func (m ExecutionReport) GetStopPx() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.StopPxField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetOrdRejReason gets OrdRejReason, Tag 103
func (m ExecutionReport) GetOrdRejReason() (v enum.OrdRejReason, err quickfix.MessageRejectError) {
	var f field.OrdRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m ExecutionReport) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m ExecutionReport) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinQty gets MinQty, Tag 110
func (m ExecutionReport) GetMinQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.MinQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetMaxFloor gets MaxFloor, Tag 111
func (m ExecutionReport) GetMaxFloor() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.MaxFloorField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetReportToExch gets ReportToExch, Tag 113
func (m ExecutionReport) GetReportToExch() (v bool, err quickfix.MessageRejectError) {
	var f field.ReportToExchField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNetMoney gets NetMoney, Tag 118
func (m ExecutionReport) GetNetMoney() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.NetMoneyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetSettlCurrAmt gets SettlCurrAmt, Tag 119
func (m ExecutionReport) GetSettlCurrAmt() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.SettlCurrAmtField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m ExecutionReport) GetSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.SettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m ExecutionReport) GetExpireTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ExpireTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoMiscFees gets NoMiscFees, Tag 136
func (m ExecutionReport) GetNoMiscFees() (NoMiscFeesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMiscFeesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetExecType gets ExecType, Tag 150
func (m ExecutionReport) GetExecType() (v enum.ExecType, err quickfix.MessageRejectError) {
	var f field.ExecTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLeavesQty gets LeavesQty, Tag 151
func (m ExecutionReport) GetLeavesQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LeavesQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m ExecutionReport) GetCashOrderQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.CashOrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetSettlCurrFxRate gets SettlCurrFxRate, Tag 155
func (m ExecutionReport) GetSettlCurrFxRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.SettlCurrFxRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetSettlCurrFxRateCalc gets SettlCurrFxRateCalc, Tag 156
func (m ExecutionReport) GetSettlCurrFxRateCalc() (v enum.SettlCurrFxRateCalc, err quickfix.MessageRejectError) {
	var f field.SettlCurrFxRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNumDaysInterest gets NumDaysInterest, Tag 157
func (m ExecutionReport) GetNumDaysInterest() (v int, err quickfix.MessageRejectError) {
	var f field.NumDaysInterestField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccruedInterestRate gets AccruedInterestRate, Tag 158
func (m ExecutionReport) GetAccruedInterestRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.AccruedInterestRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetAccruedInterestAmt gets AccruedInterestAmt, Tag 159
func (m ExecutionReport) GetAccruedInterestAmt() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.AccruedInterestAmtField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m ExecutionReport) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m ExecutionReport) GetEffectiveTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.EffectiveTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m ExecutionReport) GetOrderQty2() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.OrderQty2Field
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetSettlDate2 gets SettlDate2, Tag 193
func (m ExecutionReport) GetSettlDate2() (v string, err quickfix.MessageRejectError) {
	var f field.SettlDate2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastSpotRate gets LastSpotRate, Tag 194
func (m ExecutionReport) GetLastSpotRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LastSpotRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLastForwardPoints gets LastForwardPoints, Tag 195
func (m ExecutionReport) GetLastForwardPoints() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LastForwardPointsField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetSecondaryOrderID gets SecondaryOrderID, Tag 198
func (m ExecutionReport) GetSecondaryOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryOrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m ExecutionReport) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m ExecutionReport) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m ExecutionReport) GetStrikePrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m ExecutionReport) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m ExecutionReport) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaxShow gets MaxShow, Tag 210
func (m ExecutionReport) GetMaxShow() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.MaxShowField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetPegOffsetValue gets PegOffsetValue, Tag 211
func (m ExecutionReport) GetPegOffsetValue() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.PegOffsetValueField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetSpread gets Spread, Tag 218
func (m ExecutionReport) GetSpread() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.SpreadField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220
func (m ExecutionReport) GetBenchmarkCurveCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurveCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221
func (m ExecutionReport) GetBenchmarkCurveName() (v enum.BenchmarkCurveName, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurveNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222
func (m ExecutionReport) GetBenchmarkCurvePoint() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurvePointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m ExecutionReport) GetCouponRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m ExecutionReport) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m ExecutionReport) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m ExecutionReport) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m ExecutionReport) GetRepurchaseRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetFactor gets Factor, Tag 228
func (m ExecutionReport) GetFactor() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetTradeOriginationDate gets TradeOriginationDate, Tag 229
func (m ExecutionReport) GetTradeOriginationDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeOriginationDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExDate gets ExDate, Tag 230
func (m ExecutionReport) GetExDate() (v string, err quickfix.MessageRejectError) {
	var f field.ExDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m ExecutionReport) GetContractMultiplier() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetNoStipulations gets NoStipulations, Tag 232
func (m ExecutionReport) GetNoStipulations() (NoStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetYieldType gets YieldType, Tag 235
func (m ExecutionReport) GetYieldType() (v enum.YieldType, err quickfix.MessageRejectError) {
	var f field.YieldTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetYield gets Yield, Tag 236
func (m ExecutionReport) GetYield() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.YieldField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetTotalTakedown gets TotalTakedown, Tag 237
func (m ExecutionReport) GetTotalTakedown() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.TotalTakedownField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetConcession gets Concession, Tag 238
func (m ExecutionReport) GetConcession() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.ConcessionField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m ExecutionReport) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m ExecutionReport) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m ExecutionReport) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradedFlatSwitch gets TradedFlatSwitch, Tag 258
func (m ExecutionReport) GetTradedFlatSwitch() (v bool, err quickfix.MessageRejectError) {
	var f field.TradedFlatSwitchField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBasisFeatureDate gets BasisFeatureDate, Tag 259
func (m ExecutionReport) GetBasisFeatureDate() (v string, err quickfix.MessageRejectError) {
	var f field.BasisFeatureDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBasisFeaturePrice gets BasisFeaturePrice, Tag 260
func (m ExecutionReport) GetBasisFeaturePrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.BasisFeaturePriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m ExecutionReport) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m ExecutionReport) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m ExecutionReport) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m ExecutionReport) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m ExecutionReport) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m ExecutionReport) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m ExecutionReport) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m ExecutionReport) GetComplianceID() (v string, err quickfix.MessageRejectError) {
	var f field.ComplianceIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSolicitedFlag gets SolicitedFlag, Tag 377
func (m ExecutionReport) GetSolicitedFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.SolicitedFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecRestatementReason gets ExecRestatementReason, Tag 378
func (m ExecutionReport) GetExecRestatementReason() (v enum.ExecRestatementReason, err quickfix.MessageRejectError) {
	var f field.ExecRestatementReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetGrossTradeAmt gets GrossTradeAmt, Tag 381
func (m ExecutionReport) GetGrossTradeAmt() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.GrossTradeAmtField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetNoContraBrokers gets NoContraBrokers, Tag 382
func (m ExecutionReport) GetNoContraBrokers() (NoContraBrokersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoContraBrokersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDiscretionInst gets DiscretionInst, Tag 388
func (m ExecutionReport) GetDiscretionInst() (v enum.DiscretionInst, err quickfix.MessageRejectError) {
	var f field.DiscretionInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionOffsetValue gets DiscretionOffsetValue, Tag 389
func (m ExecutionReport) GetDiscretionOffsetValue() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.DiscretionOffsetValueField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetPriceType gets PriceType, Tag 423
func (m ExecutionReport) GetPriceType() (v enum.PriceType, err quickfix.MessageRejectError) {
	var f field.PriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDayOrderQty gets DayOrderQty, Tag 424
func (m ExecutionReport) GetDayOrderQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.DayOrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetDayCumQty gets DayCumQty, Tag 425
func (m ExecutionReport) GetDayCumQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.DayCumQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetDayAvgPx gets DayAvgPx, Tag 426
func (m ExecutionReport) GetDayAvgPx() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.DayAvgPxField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetGTBookingInst gets GTBookingInst, Tag 427
func (m ExecutionReport) GetGTBookingInst() (v enum.GTBookingInst, err quickfix.MessageRejectError) {
	var f field.GTBookingInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExpireDate gets ExpireDate, Tag 432
func (m ExecutionReport) GetExpireDate() (v string, err quickfix.MessageRejectError) {
	var f field.ExpireDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMultiLegReportingType gets MultiLegReportingType, Tag 442
func (m ExecutionReport) GetMultiLegReportingType() (v enum.MultiLegReportingType, err quickfix.MessageRejectError) {
	var f field.MultiLegReportingTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m ExecutionReport) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m ExecutionReport) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m ExecutionReport) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCFICode gets CFICode, Tag 461
func (m ExecutionReport) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m ExecutionReport) GetRoundingDirection() (v enum.RoundingDirection, err quickfix.MessageRejectError) {
	var f field.RoundingDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m ExecutionReport) GetRoundingModulus() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.RoundingModulusField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m ExecutionReport) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m ExecutionReport) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m ExecutionReport) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommCurrency gets CommCurrency, Tag 479
func (m ExecutionReport) GetCommCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CommCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCancellationRights gets CancellationRights, Tag 480
func (m ExecutionReport) GetCancellationRights() (v enum.CancellationRights, err quickfix.MessageRejectError) {
	var f field.CancellationRightsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMoneyLaunderingStatus gets MoneyLaunderingStatus, Tag 481
func (m ExecutionReport) GetMoneyLaunderingStatus() (v enum.MoneyLaunderingStatus, err quickfix.MessageRejectError) {
	var f field.MoneyLaunderingStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransBkdTime gets TransBkdTime, Tag 483
func (m ExecutionReport) GetTransBkdTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransBkdTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecPriceType gets ExecPriceType, Tag 484
func (m ExecutionReport) GetExecPriceType() (v enum.ExecPriceType, err quickfix.MessageRejectError) {
	var f field.ExecPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecPriceAdjustment gets ExecPriceAdjustment, Tag 485
func (m ExecutionReport) GetExecPriceAdjustment() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.ExecPriceAdjustmentField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetDesignation gets Designation, Tag 494
func (m ExecutionReport) GetDesignation() (v string, err quickfix.MessageRejectError) {
	var f field.DesignationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFundRenewWaiv gets FundRenewWaiv, Tag 497
func (m ExecutionReport) GetFundRenewWaiv() (v enum.FundRenewWaiv, err quickfix.MessageRejectError) {
	var f field.FundRenewWaivField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRegistID gets RegistID, Tag 513
func (m ExecutionReport) GetRegistID() (v string, err quickfix.MessageRejectError) {
	var f field.RegistIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecValuationPoint gets ExecValuationPoint, Tag 515
func (m ExecutionReport) GetExecValuationPoint() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ExecValuationPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m ExecutionReport) GetOrderPercent() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.OrderPercentField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetNoContAmts gets NoContAmts, Tag 518
func (m ExecutionReport) GetNoContAmts() (NoContAmtsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoContAmtsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m ExecutionReport) GetSecondaryClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryExecID gets SecondaryExecID, Tag 527
func (m ExecutionReport) GetSecondaryExecID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderCapacity gets OrderCapacity, Tag 528
func (m ExecutionReport) GetOrderCapacity() (v enum.OrderCapacity, err quickfix.MessageRejectError) {
	var f field.OrderCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderRestrictions gets OrderRestrictions, Tag 529
func (m ExecutionReport) GetOrderRestrictions() (v enum.OrderRestrictions, err quickfix.MessageRejectError) {
	var f field.OrderRestrictionsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m ExecutionReport) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m ExecutionReport) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashMargin gets CashMargin, Tag 544
func (m ExecutionReport) GetCashMargin() (v enum.CashMargin, err quickfix.MessageRejectError) {
	var f field.CashMarginField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCrossID gets CrossID, Tag 548
func (m ExecutionReport) GetCrossID() (v string, err quickfix.MessageRejectError) {
	var f field.CrossIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCrossType gets CrossType, Tag 549
func (m ExecutionReport) GetCrossType() (v enum.CrossType, err quickfix.MessageRejectError) {
	var f field.CrossTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrigCrossID gets OrigCrossID, Tag 551
func (m ExecutionReport) GetOrigCrossID() (v string, err quickfix.MessageRejectError) {
	var f field.OrigCrossIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoLegs gets NoLegs, Tag 555
func (m ExecutionReport) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetMatchType gets MatchType, Tag 574
func (m ExecutionReport) GetMatchType() (v enum.MatchType, err quickfix.MessageRejectError) {
	var f field.MatchTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccountType gets AccountType, Tag 581
func (m ExecutionReport) GetAccountType() (v enum.AccountType, err quickfix.MessageRejectError) {
	var f field.AccountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m ExecutionReport) GetCustOrderCapacity() (v enum.CustOrderCapacity, err quickfix.MessageRejectError) {
	var f field.CustOrderCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClOrdLinkID gets ClOrdLinkID, Tag 583
func (m ExecutionReport) GetClOrdLinkID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdLinkIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMassStatusReqID gets MassStatusReqID, Tag 584
func (m ExecutionReport) GetMassStatusReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MassStatusReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDayBookingInst gets DayBookingInst, Tag 589
func (m ExecutionReport) GetDayBookingInst() (v enum.DayBookingInst, err quickfix.MessageRejectError) {
	var f field.DayBookingInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBookingUnit gets BookingUnit, Tag 590
func (m ExecutionReport) GetBookingUnit() (v enum.BookingUnit, err quickfix.MessageRejectError) {
	var f field.BookingUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPreallocMethod gets PreallocMethod, Tag 591
func (m ExecutionReport) GetPreallocMethod() (v enum.PreallocMethod, err quickfix.MessageRejectError) {
	var f field.PreallocMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m ExecutionReport) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClearingFeeIndicator gets ClearingFeeIndicator, Tag 635
func (m ExecutionReport) GetClearingFeeIndicator() (v enum.ClearingFeeIndicator, err quickfix.MessageRejectError) {
	var f field.ClearingFeeIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetWorkingIndicator gets WorkingIndicator, Tag 636
func (m ExecutionReport) GetWorkingIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.WorkingIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriorityIndicator gets PriorityIndicator, Tag 638
func (m ExecutionReport) GetPriorityIndicator() (v enum.PriorityIndicator, err quickfix.MessageRejectError) {
	var f field.PriorityIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceImprovement gets PriceImprovement, Tag 639
func (m ExecutionReport) GetPriceImprovement() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.PriceImprovementField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLastForwardPoints2 gets LastForwardPoints2, Tag 641
func (m ExecutionReport) GetLastForwardPoints2() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LastForwardPoints2Field
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingLastPx gets UnderlyingLastPx, Tag 651
func (m ExecutionReport) GetUnderlyingLastPx() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingLastPxField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingLastQty gets UnderlyingLastQty, Tag 652
func (m ExecutionReport) GetUnderlyingLastQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingLastQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetAcctIDSource gets AcctIDSource, Tag 660
func (m ExecutionReport) GetAcctIDSource() (v enum.AcctIDSource, err quickfix.MessageRejectError) {
	var f field.AcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBenchmarkPrice gets BenchmarkPrice, Tag 662
func (m ExecutionReport) GetBenchmarkPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.BenchmarkPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetBenchmarkPriceType gets BenchmarkPriceType, Tag 663
func (m ExecutionReport) GetBenchmarkPriceType() (v int, err quickfix.MessageRejectError) {
	var f field.BenchmarkPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m ExecutionReport) GetContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.ContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastParPx gets LastParPx, Tag 669
func (m ExecutionReport) GetLastParPx() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LastParPxField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetPool gets Pool, Tag 691
func (m ExecutionReport) GetPool() (v string, err quickfix.MessageRejectError) {
	var f field.PoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteRespID gets QuoteRespID, Tag 693
func (m ExecutionReport) GetQuoteRespID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteRespIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetYieldRedemptionDate gets YieldRedemptionDate, Tag 696
func (m ExecutionReport) GetYieldRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetYieldRedemptionPrice gets YieldRedemptionPrice, Tag 697
func (m ExecutionReport) GetYieldRedemptionPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetYieldRedemptionPriceType gets YieldRedemptionPriceType, Tag 698
func (m ExecutionReport) GetYieldRedemptionPriceType() (v int, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBenchmarkSecurityID gets BenchmarkSecurityID, Tag 699
func (m ExecutionReport) GetBenchmarkSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetYieldCalcDate gets YieldCalcDate, Tag 701
func (m ExecutionReport) GetYieldCalcDate() (v string, err quickfix.MessageRejectError) {
	var f field.YieldCalcDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m ExecutionReport) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetInterestAtMaturity gets InterestAtMaturity, Tag 738
func (m ExecutionReport) GetInterestAtMaturity() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.InterestAtMaturityField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetBenchmarkSecurityIDSource gets BenchmarkSecurityIDSource, Tag 761
func (m ExecutionReport) GetBenchmarkSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m ExecutionReport) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoTrdRegTimestamps gets NoTrdRegTimestamps, Tag 768
func (m ExecutionReport) GetNoTrdRegTimestamps() (NoTrdRegTimestampsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTrdRegTimestampsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetBookingType gets BookingType, Tag 775
func (m ExecutionReport) GetBookingType() (v enum.BookingType, err quickfix.MessageRejectError) {
	var f field.BookingTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTerminationType gets TerminationType, Tag 788
func (m ExecutionReport) GetTerminationType() (v enum.TerminationType, err quickfix.MessageRejectError) {
	var f field.TerminationTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrdStatusReqID gets OrdStatusReqID, Tag 790
func (m ExecutionReport) GetOrdStatusReqID() (v string, err quickfix.MessageRejectError) {
	var f field.OrdStatusReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCopyMsgIndicator gets CopyMsgIndicator, Tag 797
func (m ExecutionReport) GetCopyMsgIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.CopyMsgIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceDelta gets PriceDelta, Tag 811
func (m ExecutionReport) GetPriceDelta() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.PriceDeltaField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetPegMoveType gets PegMoveType, Tag 835
func (m ExecutionReport) GetPegMoveType() (v enum.PegMoveType, err quickfix.MessageRejectError) {
	var f field.PegMoveTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegOffsetType gets PegOffsetType, Tag 836
func (m ExecutionReport) GetPegOffsetType() (v enum.PegOffsetType, err quickfix.MessageRejectError) {
	var f field.PegOffsetTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegLimitType gets PegLimitType, Tag 837
func (m ExecutionReport) GetPegLimitType() (v enum.PegLimitType, err quickfix.MessageRejectError) {
	var f field.PegLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegRoundDirection gets PegRoundDirection, Tag 838
func (m ExecutionReport) GetPegRoundDirection() (v enum.PegRoundDirection, err quickfix.MessageRejectError) {
	var f field.PegRoundDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPeggedPrice gets PeggedPrice, Tag 839
func (m ExecutionReport) GetPeggedPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.PeggedPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetPegScope gets PegScope, Tag 840
func (m ExecutionReport) GetPegScope() (v enum.PegScope, err quickfix.MessageRejectError) {
	var f field.PegScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionMoveType gets DiscretionMoveType, Tag 841
func (m ExecutionReport) GetDiscretionMoveType() (v enum.DiscretionMoveType, err quickfix.MessageRejectError) {
	var f field.DiscretionMoveTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionOffsetType gets DiscretionOffsetType, Tag 842
func (m ExecutionReport) GetDiscretionOffsetType() (v enum.DiscretionOffsetType, err quickfix.MessageRejectError) {
	var f field.DiscretionOffsetTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionLimitType gets DiscretionLimitType, Tag 843
func (m ExecutionReport) GetDiscretionLimitType() (v enum.DiscretionLimitType, err quickfix.MessageRejectError) {
	var f field.DiscretionLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionRoundDirection gets DiscretionRoundDirection, Tag 844
func (m ExecutionReport) GetDiscretionRoundDirection() (v enum.DiscretionRoundDirection, err quickfix.MessageRejectError) {
	var f field.DiscretionRoundDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDiscretionPrice gets DiscretionPrice, Tag 845
func (m ExecutionReport) GetDiscretionPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.DiscretionPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetDiscretionScope gets DiscretionScope, Tag 846
func (m ExecutionReport) GetDiscretionScope() (v enum.DiscretionScope, err quickfix.MessageRejectError) {
	var f field.DiscretionScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTargetStrategy gets TargetStrategy, Tag 847
func (m ExecutionReport) GetTargetStrategy() (v enum.TargetStrategy, err quickfix.MessageRejectError) {
	var f field.TargetStrategyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTargetStrategyParameters gets TargetStrategyParameters, Tag 848
func (m ExecutionReport) GetTargetStrategyParameters() (v string, err quickfix.MessageRejectError) {
	var f field.TargetStrategyParametersField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetParticipationRate gets ParticipationRate, Tag 849
func (m ExecutionReport) GetParticipationRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.ParticipationRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetTargetStrategyPerformance gets TargetStrategyPerformance, Tag 850
func (m ExecutionReport) GetTargetStrategyPerformance() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.TargetStrategyPerformanceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLastLiquidityInd gets LastLiquidityInd, Tag 851
func (m ExecutionReport) GetLastLiquidityInd() (v enum.LastLiquidityInd, err quickfix.MessageRejectError) {
	var f field.LastLiquidityIndField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQtyType gets QtyType, Tag 854
func (m ExecutionReport) GetQtyType() (v enum.QtyType, err quickfix.MessageRejectError) {
	var f field.QtyTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m ExecutionReport) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m ExecutionReport) GetDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.DatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m ExecutionReport) GetInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.InterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m ExecutionReport) GetCPProgram() (v enum.CPProgram, err quickfix.MessageRejectError) {
	var f field.CPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m ExecutionReport) GetCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.CPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTrdMatchID gets TrdMatchID, Tag 880
func (m ExecutionReport) GetTrdMatchID() (v string, err quickfix.MessageRejectError) {
	var f field.TrdMatchIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastFragment gets LastFragment, Tag 893
func (m ExecutionReport) GetLastFragment() (v bool, err quickfix.MessageRejectError) {
	var f field.LastFragmentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarginRatio gets MarginRatio, Tag 898
func (m ExecutionReport) GetMarginRatio() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.MarginRatioField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetTotNumReports gets TotNumReports, Tag 911
func (m ExecutionReport) GetTotNumReports() (v int, err quickfix.MessageRejectError) {
	var f field.TotNumReportsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastRptRequested gets LastRptRequested, Tag 912
func (m ExecutionReport) GetLastRptRequested() (v bool, err quickfix.MessageRejectError) {
	var f field.LastRptRequestedField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAgreementDesc gets AgreementDesc, Tag 913
func (m ExecutionReport) GetAgreementDesc() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAgreementID gets AgreementID, Tag 914
func (m ExecutionReport) GetAgreementID() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAgreementDate gets AgreementDate, Tag 915
func (m ExecutionReport) GetAgreementDate() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStartDate gets StartDate, Tag 916
func (m ExecutionReport) GetStartDate() (v string, err quickfix.MessageRejectError) {
	var f field.StartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEndDate gets EndDate, Tag 917
func (m ExecutionReport) GetEndDate() (v string, err quickfix.MessageRejectError) {
	var f field.EndDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAgreementCurrency gets AgreementCurrency, Tag 918
func (m ExecutionReport) GetAgreementCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDeliveryType gets DeliveryType, Tag 919
func (m ExecutionReport) GetDeliveryType() (v enum.DeliveryType, err quickfix.MessageRejectError) {
	var f field.DeliveryTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEndAccruedInterestAmt gets EndAccruedInterestAmt, Tag 920
func (m ExecutionReport) GetEndAccruedInterestAmt() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.EndAccruedInterestAmtField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetStartCash gets StartCash, Tag 921
func (m ExecutionReport) GetStartCash() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.StartCashField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetEndCash gets EndCash, Tag 922
func (m ExecutionReport) GetEndCash() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.EndCashField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetTimeBracket gets TimeBracket, Tag 943
func (m ExecutionReport) GetTimeBracket() (v string, err quickfix.MessageRejectError) {
	var f field.TimeBracketField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m ExecutionReport) GetStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoStrategyParameters gets NoStrategyParameters, Tag 957
func (m ExecutionReport) GetNoStrategyParameters() (NoStrategyParametersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStrategyParametersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetHostCrossID gets HostCrossID, Tag 961
func (m ExecutionReport) GetHostCrossID() (v string, err quickfix.MessageRejectError) {
	var f field.HostCrossIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m ExecutionReport) GetSecurityStatus() (v enum.SecurityStatus, err quickfix.MessageRejectError) {
	var f field.SecurityStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m ExecutionReport) GetSettleOnOpenFlag() (v string, err quickfix.MessageRejectError) {
	var f field.SettleOnOpenFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m ExecutionReport) GetStrikeMultiplier() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.StrikeMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m ExecutionReport) GetStrikeValue() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.StrikeValueField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m ExecutionReport) GetMinPriceIncrement() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m ExecutionReport) GetPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.PositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m ExecutionReport) GetNTPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.NTPositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m ExecutionReport) GetUnitOfMeasure() (v enum.UnitOfMeasure, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m ExecutionReport) GetTimeUnit() (v enum.TimeUnit, err quickfix.MessageRejectError) {
	var f field.TimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m ExecutionReport) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetManualOrderIndicator gets ManualOrderIndicator, Tag 1028
func (m ExecutionReport) GetManualOrderIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.ManualOrderIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCustDirectedOrder gets CustDirectedOrder, Tag 1029
func (m ExecutionReport) GetCustDirectedOrder() (v bool, err quickfix.MessageRejectError) {
	var f field.CustDirectedOrderField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetReceivedDeptID gets ReceivedDeptID, Tag 1030
func (m ExecutionReport) GetReceivedDeptID() (v string, err quickfix.MessageRejectError) {
	var f field.ReceivedDeptIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCustOrderHandlingInst gets CustOrderHandlingInst, Tag 1031
func (m ExecutionReport) GetCustOrderHandlingInst() (v enum.CustOrderHandlingInst, err quickfix.MessageRejectError) {
	var f field.CustOrderHandlingInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderHandlingInstSource gets OrderHandlingInstSource, Tag 1032
func (m ExecutionReport) GetOrderHandlingInstSource() (v enum.OrderHandlingInstSource, err quickfix.MessageRejectError) {
	var f field.OrderHandlingInstSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m ExecutionReport) GetInstrmtAssignmentMethod() (v string, err quickfix.MessageRejectError) {
	var f field.InstrmtAssignmentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCalculatedCcyLastQty gets CalculatedCcyLastQty, Tag 1056
func (m ExecutionReport) GetCalculatedCcyLastQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.CalculatedCcyLastQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetAggressorIndicator gets AggressorIndicator, Tag 1057
func (m ExecutionReport) GetAggressorIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.AggressorIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastSwapPoints gets LastSwapPoints, Tag 1071
func (m ExecutionReport) GetLastSwapPoints() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LastSwapPointsField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m ExecutionReport) GetMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryDisplayQty gets SecondaryDisplayQty, Tag 1082
func (m ExecutionReport) GetSecondaryDisplayQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.SecondaryDisplayQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetDisplayWhen gets DisplayWhen, Tag 1083
func (m ExecutionReport) GetDisplayWhen() (v enum.DisplayWhen, err quickfix.MessageRejectError) {
	var f field.DisplayWhenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDisplayMethod gets DisplayMethod, Tag 1084
func (m ExecutionReport) GetDisplayMethod() (v enum.DisplayMethod, err quickfix.MessageRejectError) {
	var f field.DisplayMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDisplayLowQty gets DisplayLowQty, Tag 1085
func (m ExecutionReport) GetDisplayLowQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.DisplayLowQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetDisplayHighQty gets DisplayHighQty, Tag 1086
func (m ExecutionReport) GetDisplayHighQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.DisplayHighQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetDisplayMinIncr gets DisplayMinIncr, Tag 1087
func (m ExecutionReport) GetDisplayMinIncr() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.DisplayMinIncrField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetRefreshQty gets RefreshQty, Tag 1088
func (m ExecutionReport) GetRefreshQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.RefreshQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetMatchIncrement gets MatchIncrement, Tag 1089
func (m ExecutionReport) GetMatchIncrement() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.MatchIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetMaxPriceLevels gets MaxPriceLevels, Tag 1090
func (m ExecutionReport) GetMaxPriceLevels() (v int, err quickfix.MessageRejectError) {
	var f field.MaxPriceLevelsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPreTradeAnonymity gets PreTradeAnonymity, Tag 1091
func (m ExecutionReport) GetPreTradeAnonymity() (v bool, err quickfix.MessageRejectError) {
	var f field.PreTradeAnonymityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceProtectionScope gets PriceProtectionScope, Tag 1092
func (m ExecutionReport) GetPriceProtectionScope() (v enum.PriceProtectionScope, err quickfix.MessageRejectError) {
	var f field.PriceProtectionScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLotType gets LotType, Tag 1093
func (m ExecutionReport) GetLotType() (v enum.LotType, err quickfix.MessageRejectError) {
	var f field.LotTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegPriceType gets PegPriceType, Tag 1094
func (m ExecutionReport) GetPegPriceType() (v enum.PegPriceType, err quickfix.MessageRejectError) {
	var f field.PegPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPeggedRefPrice gets PeggedRefPrice, Tag 1095
func (m ExecutionReport) GetPeggedRefPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.PeggedRefPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetPegSecurityIDSource gets PegSecurityIDSource, Tag 1096
func (m ExecutionReport) GetPegSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.PegSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegSecurityID gets PegSecurityID, Tag 1097
func (m ExecutionReport) GetPegSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.PegSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegSymbol gets PegSymbol, Tag 1098
func (m ExecutionReport) GetPegSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.PegSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPegSecurityDesc gets PegSecurityDesc, Tag 1099
func (m ExecutionReport) GetPegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.PegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerType gets TriggerType, Tag 1100
func (m ExecutionReport) GetTriggerType() (v enum.TriggerType, err quickfix.MessageRejectError) {
	var f field.TriggerTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerAction gets TriggerAction, Tag 1101
func (m ExecutionReport) GetTriggerAction() (v enum.TriggerAction, err quickfix.MessageRejectError) {
	var f field.TriggerActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerPrice gets TriggerPrice, Tag 1102
func (m ExecutionReport) GetTriggerPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.TriggerPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetTriggerSymbol gets TriggerSymbol, Tag 1103
func (m ExecutionReport) GetTriggerSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerSecurityID gets TriggerSecurityID, Tag 1104
func (m ExecutionReport) GetTriggerSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerSecurityIDSource gets TriggerSecurityIDSource, Tag 1105
func (m ExecutionReport) GetTriggerSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerSecurityDesc gets TriggerSecurityDesc, Tag 1106
func (m ExecutionReport) GetTriggerSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerPriceType gets TriggerPriceType, Tag 1107
func (m ExecutionReport) GetTriggerPriceType() (v enum.TriggerPriceType, err quickfix.MessageRejectError) {
	var f field.TriggerPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerPriceTypeScope gets TriggerPriceTypeScope, Tag 1108
func (m ExecutionReport) GetTriggerPriceTypeScope() (v enum.TriggerPriceTypeScope, err quickfix.MessageRejectError) {
	var f field.TriggerPriceTypeScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerPriceDirection gets TriggerPriceDirection, Tag 1109
func (m ExecutionReport) GetTriggerPriceDirection() (v enum.TriggerPriceDirection, err quickfix.MessageRejectError) {
	var f field.TriggerPriceDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerNewPrice gets TriggerNewPrice, Tag 1110
func (m ExecutionReport) GetTriggerNewPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.TriggerNewPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetTriggerOrderType gets TriggerOrderType, Tag 1111
func (m ExecutionReport) GetTriggerOrderType() (v enum.TriggerOrderType, err quickfix.MessageRejectError) {
	var f field.TriggerOrderTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerNewQty gets TriggerNewQty, Tag 1112
func (m ExecutionReport) GetTriggerNewQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.TriggerNewQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetTriggerTradingSessionID gets TriggerTradingSessionID, Tag 1113
func (m ExecutionReport) GetTriggerTradingSessionID() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerTradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTriggerTradingSessionSubID gets TriggerTradingSessionSubID, Tag 1114
func (m ExecutionReport) GetTriggerTradingSessionSubID() (v string, err quickfix.MessageRejectError) {
	var f field.TriggerTradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderCategory gets OrderCategory, Tag 1115
func (m ExecutionReport) GetOrderCategory() (v enum.OrderCategory, err quickfix.MessageRejectError) {
	var f field.OrderCategoryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDisplayQty gets DisplayQty, Tag 1138
func (m ExecutionReport) GetDisplayQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.DisplayQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146
func (m ExecutionReport) GetMinPriceIncrementAmount() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementAmountField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147
func (m ExecutionReport) GetUnitOfMeasureQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetSecurityGroup gets SecurityGroup, Tag 1151
func (m ExecutionReport) GetSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplID gets ApplID, Tag 1180
func (m ExecutionReport) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplSeqNum gets ApplSeqNum, Tag 1181
func (m ExecutionReport) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityXMLLen gets SecurityXMLLen, Tag 1184
func (m ExecutionReport) GetSecurityXMLLen() (v int, err quickfix.MessageRejectError) {
	var f field.SecurityXMLLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityXML gets SecurityXML, Tag 1185
func (m ExecutionReport) GetSecurityXML() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186
func (m ExecutionReport) GetSecurityXMLSchema() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLSchemaField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetVolatility gets Volatility, Tag 1188
func (m ExecutionReport) GetVolatility() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.VolatilityField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetTimeToExpiration gets TimeToExpiration, Tag 1189
func (m ExecutionReport) GetTimeToExpiration() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.TimeToExpirationField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetRiskFreeRate gets RiskFreeRate, Tag 1190
func (m ExecutionReport) GetRiskFreeRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.RiskFreeRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191
func (m ExecutionReport) GetPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192
func (m ExecutionReport) GetPriceUnitOfMeasureQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetSettlMethod gets SettlMethod, Tag 1193
func (m ExecutionReport) GetSettlMethod() (v enum.SettlMethod, err quickfix.MessageRejectError) {
	var f field.SettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExerciseStyle gets ExerciseStyle, Tag 1194
func (m ExecutionReport) GetExerciseStyle() (v enum.ExerciseStyle, err quickfix.MessageRejectError) {
	var f field.ExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptPayAmount gets OptPayAmount, Tag 1195
func (m ExecutionReport) GetOptPayAmount() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.OptPayAmountField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196
func (m ExecutionReport) GetPriceQuoteMethod() (v enum.PriceQuoteMethod, err quickfix.MessageRejectError) {
	var f field.PriceQuoteMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFuturesValuationMethod gets FuturesValuationMethod, Tag 1197
func (m ExecutionReport) GetFuturesValuationMethod() (v enum.FuturesValuationMethod, err quickfix.MessageRejectError) {
	var f field.FuturesValuationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetListMethod gets ListMethod, Tag 1198
func (m ExecutionReport) GetListMethod() (v enum.ListMethod, err quickfix.MessageRejectError) {
	var f field.ListMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCapPrice gets CapPrice, Tag 1199
func (m ExecutionReport) GetCapPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.CapPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetFloorPrice gets FloorPrice, Tag 1200
func (m ExecutionReport) GetFloorPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.FloorPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetProductComplex gets ProductComplex, Tag 1227
func (m ExecutionReport) GetProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.ProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242
func (m ExecutionReport) GetFlexProductEligibilityIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexProductEligibilityIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFlexibleIndicator gets FlexibleIndicator, Tag 1244
func (m ExecutionReport) GetFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350
func (m ExecutionReport) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplResendFlag gets ApplResendFlag, Tag 1352
func (m ExecutionReport) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotNoFills gets TotNoFills, Tag 1361
func (m ExecutionReport) GetTotNoFills() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoFillsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoFills gets NoFills, Tag 1362
func (m ExecutionReport) GetNoFills() (NoFillsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoFillsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDividendYield gets DividendYield, Tag 1380
func (m ExecutionReport) GetDividendYield() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.DividendYieldField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m ExecutionReport) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasAvgPx returns true if AvgPx is present, Tag 6
func (m ExecutionReport) HasAvgPx() bool {
	return m.Has(tag.AvgPx)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m ExecutionReport) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasCommission returns true if Commission is present, Tag 12
func (m ExecutionReport) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m ExecutionReport) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCumQty returns true if CumQty is present, Tag 14
func (m ExecutionReport) HasCumQty() bool {
	return m.Has(tag.CumQty)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m ExecutionReport) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasExecID returns true if ExecID is present, Tag 17
func (m ExecutionReport) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m ExecutionReport) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasExecRefID returns true if ExecRefID is present, Tag 19
func (m ExecutionReport) HasExecRefID() bool {
	return m.Has(tag.ExecRefID)
}

//HasHandlInst returns true if HandlInst is present, Tag 21
func (m ExecutionReport) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m ExecutionReport) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasLastCapacity returns true if LastCapacity is present, Tag 29
func (m ExecutionReport) HasLastCapacity() bool {
	return m.Has(tag.LastCapacity)
}

//HasLastMkt returns true if LastMkt is present, Tag 30
func (m ExecutionReport) HasLastMkt() bool {
	return m.Has(tag.LastMkt)
}

//HasLastPx returns true if LastPx is present, Tag 31
func (m ExecutionReport) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

//HasLastQty returns true if LastQty is present, Tag 32
func (m ExecutionReport) HasLastQty() bool {
	return m.Has(tag.LastQty)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m ExecutionReport) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m ExecutionReport) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrdStatus returns true if OrdStatus is present, Tag 39
func (m ExecutionReport) HasOrdStatus() bool {
	return m.Has(tag.OrdStatus)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m ExecutionReport) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41
func (m ExecutionReport) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

//HasPrice returns true if Price is present, Tag 44
func (m ExecutionReport) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m ExecutionReport) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m ExecutionReport) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m ExecutionReport) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m ExecutionReport) HasText() bool {
	return m.Has(tag.Text)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m ExecutionReport) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m ExecutionReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlType returns true if SettlType is present, Tag 63
func (m ExecutionReport) HasSettlType() bool {
	return m.Has(tag.SettlType)
}

//HasSettlDate returns true if SettlDate is present, Tag 64
func (m ExecutionReport) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m ExecutionReport) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasListID returns true if ListID is present, Tag 66
func (m ExecutionReport) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m ExecutionReport) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m ExecutionReport) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasPositionEffect returns true if PositionEffect is present, Tag 77
func (m ExecutionReport) HasPositionEffect() bool {
	return m.Has(tag.PositionEffect)
}

//HasNoAllocs returns true if NoAllocs is present, Tag 78
func (m ExecutionReport) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

//HasStopPx returns true if StopPx is present, Tag 99
func (m ExecutionReport) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

//HasOrdRejReason returns true if OrdRejReason is present, Tag 103
func (m ExecutionReport) HasOrdRejReason() bool {
	return m.Has(tag.OrdRejReason)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m ExecutionReport) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m ExecutionReport) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasMinQty returns true if MinQty is present, Tag 110
func (m ExecutionReport) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

//HasMaxFloor returns true if MaxFloor is present, Tag 111
func (m ExecutionReport) HasMaxFloor() bool {
	return m.Has(tag.MaxFloor)
}

//HasReportToExch returns true if ReportToExch is present, Tag 113
func (m ExecutionReport) HasReportToExch() bool {
	return m.Has(tag.ReportToExch)
}

//HasNetMoney returns true if NetMoney is present, Tag 118
func (m ExecutionReport) HasNetMoney() bool {
	return m.Has(tag.NetMoney)
}

//HasSettlCurrAmt returns true if SettlCurrAmt is present, Tag 119
func (m ExecutionReport) HasSettlCurrAmt() bool {
	return m.Has(tag.SettlCurrAmt)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m ExecutionReport) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m ExecutionReport) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasNoMiscFees returns true if NoMiscFees is present, Tag 136
func (m ExecutionReport) HasNoMiscFees() bool {
	return m.Has(tag.NoMiscFees)
}

//HasExecType returns true if ExecType is present, Tag 150
func (m ExecutionReport) HasExecType() bool {
	return m.Has(tag.ExecType)
}

//HasLeavesQty returns true if LeavesQty is present, Tag 151
func (m ExecutionReport) HasLeavesQty() bool {
	return m.Has(tag.LeavesQty)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m ExecutionReport) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSettlCurrFxRate returns true if SettlCurrFxRate is present, Tag 155
func (m ExecutionReport) HasSettlCurrFxRate() bool {
	return m.Has(tag.SettlCurrFxRate)
}

//HasSettlCurrFxRateCalc returns true if SettlCurrFxRateCalc is present, Tag 156
func (m ExecutionReport) HasSettlCurrFxRateCalc() bool {
	return m.Has(tag.SettlCurrFxRateCalc)
}

//HasNumDaysInterest returns true if NumDaysInterest is present, Tag 157
func (m ExecutionReport) HasNumDaysInterest() bool {
	return m.Has(tag.NumDaysInterest)
}

//HasAccruedInterestRate returns true if AccruedInterestRate is present, Tag 158
func (m ExecutionReport) HasAccruedInterestRate() bool {
	return m.Has(tag.AccruedInterestRate)
}

//HasAccruedInterestAmt returns true if AccruedInterestAmt is present, Tag 159
func (m ExecutionReport) HasAccruedInterestAmt() bool {
	return m.Has(tag.AccruedInterestAmt)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m ExecutionReport) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m ExecutionReport) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m ExecutionReport) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasSettlDate2 returns true if SettlDate2 is present, Tag 193
func (m ExecutionReport) HasSettlDate2() bool {
	return m.Has(tag.SettlDate2)
}

//HasLastSpotRate returns true if LastSpotRate is present, Tag 194
func (m ExecutionReport) HasLastSpotRate() bool {
	return m.Has(tag.LastSpotRate)
}

//HasLastForwardPoints returns true if LastForwardPoints is present, Tag 195
func (m ExecutionReport) HasLastForwardPoints() bool {
	return m.Has(tag.LastForwardPoints)
}

//HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198
func (m ExecutionReport) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m ExecutionReport) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m ExecutionReport) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m ExecutionReport) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m ExecutionReport) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m ExecutionReport) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasMaxShow returns true if MaxShow is present, Tag 210
func (m ExecutionReport) HasMaxShow() bool {
	return m.Has(tag.MaxShow)
}

//HasPegOffsetValue returns true if PegOffsetValue is present, Tag 211
func (m ExecutionReport) HasPegOffsetValue() bool {
	return m.Has(tag.PegOffsetValue)
}

//HasSpread returns true if Spread is present, Tag 218
func (m ExecutionReport) HasSpread() bool {
	return m.Has(tag.Spread)
}

//HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220
func (m ExecutionReport) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

//HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221
func (m ExecutionReport) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

//HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222
func (m ExecutionReport) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m ExecutionReport) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m ExecutionReport) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m ExecutionReport) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m ExecutionReport) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m ExecutionReport) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m ExecutionReport) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasTradeOriginationDate returns true if TradeOriginationDate is present, Tag 229
func (m ExecutionReport) HasTradeOriginationDate() bool {
	return m.Has(tag.TradeOriginationDate)
}

//HasExDate returns true if ExDate is present, Tag 230
func (m ExecutionReport) HasExDate() bool {
	return m.Has(tag.ExDate)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m ExecutionReport) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasNoStipulations returns true if NoStipulations is present, Tag 232
func (m ExecutionReport) HasNoStipulations() bool {
	return m.Has(tag.NoStipulations)
}

//HasYieldType returns true if YieldType is present, Tag 235
func (m ExecutionReport) HasYieldType() bool {
	return m.Has(tag.YieldType)
}

//HasYield returns true if Yield is present, Tag 236
func (m ExecutionReport) HasYield() bool {
	return m.Has(tag.Yield)
}

//HasTotalTakedown returns true if TotalTakedown is present, Tag 237
func (m ExecutionReport) HasTotalTakedown() bool {
	return m.Has(tag.TotalTakedown)
}

//HasConcession returns true if Concession is present, Tag 238
func (m ExecutionReport) HasConcession() bool {
	return m.Has(tag.Concession)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m ExecutionReport) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m ExecutionReport) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m ExecutionReport) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasTradedFlatSwitch returns true if TradedFlatSwitch is present, Tag 258
func (m ExecutionReport) HasTradedFlatSwitch() bool {
	return m.Has(tag.TradedFlatSwitch)
}

//HasBasisFeatureDate returns true if BasisFeatureDate is present, Tag 259
func (m ExecutionReport) HasBasisFeatureDate() bool {
	return m.Has(tag.BasisFeatureDate)
}

//HasBasisFeaturePrice returns true if BasisFeaturePrice is present, Tag 260
func (m ExecutionReport) HasBasisFeaturePrice() bool {
	return m.Has(tag.BasisFeaturePrice)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m ExecutionReport) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m ExecutionReport) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m ExecutionReport) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m ExecutionReport) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m ExecutionReport) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m ExecutionReport) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m ExecutionReport) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasComplianceID returns true if ComplianceID is present, Tag 376
func (m ExecutionReport) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

//HasSolicitedFlag returns true if SolicitedFlag is present, Tag 377
func (m ExecutionReport) HasSolicitedFlag() bool {
	return m.Has(tag.SolicitedFlag)
}

//HasExecRestatementReason returns true if ExecRestatementReason is present, Tag 378
func (m ExecutionReport) HasExecRestatementReason() bool {
	return m.Has(tag.ExecRestatementReason)
}

//HasGrossTradeAmt returns true if GrossTradeAmt is present, Tag 381
func (m ExecutionReport) HasGrossTradeAmt() bool {
	return m.Has(tag.GrossTradeAmt)
}

//HasNoContraBrokers returns true if NoContraBrokers is present, Tag 382
func (m ExecutionReport) HasNoContraBrokers() bool {
	return m.Has(tag.NoContraBrokers)
}

//HasDiscretionInst returns true if DiscretionInst is present, Tag 388
func (m ExecutionReport) HasDiscretionInst() bool {
	return m.Has(tag.DiscretionInst)
}

//HasDiscretionOffsetValue returns true if DiscretionOffsetValue is present, Tag 389
func (m ExecutionReport) HasDiscretionOffsetValue() bool {
	return m.Has(tag.DiscretionOffsetValue)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m ExecutionReport) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasDayOrderQty returns true if DayOrderQty is present, Tag 424
func (m ExecutionReport) HasDayOrderQty() bool {
	return m.Has(tag.DayOrderQty)
}

//HasDayCumQty returns true if DayCumQty is present, Tag 425
func (m ExecutionReport) HasDayCumQty() bool {
	return m.Has(tag.DayCumQty)
}

//HasDayAvgPx returns true if DayAvgPx is present, Tag 426
func (m ExecutionReport) HasDayAvgPx() bool {
	return m.Has(tag.DayAvgPx)
}

//HasGTBookingInst returns true if GTBookingInst is present, Tag 427
func (m ExecutionReport) HasGTBookingInst() bool {
	return m.Has(tag.GTBookingInst)
}

//HasExpireDate returns true if ExpireDate is present, Tag 432
func (m ExecutionReport) HasExpireDate() bool {
	return m.Has(tag.ExpireDate)
}

//HasMultiLegReportingType returns true if MultiLegReportingType is present, Tag 442
func (m ExecutionReport) HasMultiLegReportingType() bool {
	return m.Has(tag.MultiLegReportingType)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m ExecutionReport) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m ExecutionReport) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m ExecutionReport) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m ExecutionReport) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m ExecutionReport) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

//HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m ExecutionReport) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m ExecutionReport) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m ExecutionReport) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m ExecutionReport) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasCommCurrency returns true if CommCurrency is present, Tag 479
func (m ExecutionReport) HasCommCurrency() bool {
	return m.Has(tag.CommCurrency)
}

//HasCancellationRights returns true if CancellationRights is present, Tag 480
func (m ExecutionReport) HasCancellationRights() bool {
	return m.Has(tag.CancellationRights)
}

//HasMoneyLaunderingStatus returns true if MoneyLaunderingStatus is present, Tag 481
func (m ExecutionReport) HasMoneyLaunderingStatus() bool {
	return m.Has(tag.MoneyLaunderingStatus)
}

//HasTransBkdTime returns true if TransBkdTime is present, Tag 483
func (m ExecutionReport) HasTransBkdTime() bool {
	return m.Has(tag.TransBkdTime)
}

//HasExecPriceType returns true if ExecPriceType is present, Tag 484
func (m ExecutionReport) HasExecPriceType() bool {
	return m.Has(tag.ExecPriceType)
}

//HasExecPriceAdjustment returns true if ExecPriceAdjustment is present, Tag 485
func (m ExecutionReport) HasExecPriceAdjustment() bool {
	return m.Has(tag.ExecPriceAdjustment)
}

//HasDesignation returns true if Designation is present, Tag 494
func (m ExecutionReport) HasDesignation() bool {
	return m.Has(tag.Designation)
}

//HasFundRenewWaiv returns true if FundRenewWaiv is present, Tag 497
func (m ExecutionReport) HasFundRenewWaiv() bool {
	return m.Has(tag.FundRenewWaiv)
}

//HasRegistID returns true if RegistID is present, Tag 513
func (m ExecutionReport) HasRegistID() bool {
	return m.Has(tag.RegistID)
}

//HasExecValuationPoint returns true if ExecValuationPoint is present, Tag 515
func (m ExecutionReport) HasExecValuationPoint() bool {
	return m.Has(tag.ExecValuationPoint)
}

//HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m ExecutionReport) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

//HasNoContAmts returns true if NoContAmts is present, Tag 518
func (m ExecutionReport) HasNoContAmts() bool {
	return m.Has(tag.NoContAmts)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m ExecutionReport) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasSecondaryExecID returns true if SecondaryExecID is present, Tag 527
func (m ExecutionReport) HasSecondaryExecID() bool {
	return m.Has(tag.SecondaryExecID)
}

//HasOrderCapacity returns true if OrderCapacity is present, Tag 528
func (m ExecutionReport) HasOrderCapacity() bool {
	return m.Has(tag.OrderCapacity)
}

//HasOrderRestrictions returns true if OrderRestrictions is present, Tag 529
func (m ExecutionReport) HasOrderRestrictions() bool {
	return m.Has(tag.OrderRestrictions)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m ExecutionReport) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m ExecutionReport) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasCashMargin returns true if CashMargin is present, Tag 544
func (m ExecutionReport) HasCashMargin() bool {
	return m.Has(tag.CashMargin)
}

//HasCrossID returns true if CrossID is present, Tag 548
func (m ExecutionReport) HasCrossID() bool {
	return m.Has(tag.CrossID)
}

//HasCrossType returns true if CrossType is present, Tag 549
func (m ExecutionReport) HasCrossType() bool {
	return m.Has(tag.CrossType)
}

//HasOrigCrossID returns true if OrigCrossID is present, Tag 551
func (m ExecutionReport) HasOrigCrossID() bool {
	return m.Has(tag.OrigCrossID)
}

//HasNoLegs returns true if NoLegs is present, Tag 555
func (m ExecutionReport) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

//HasMatchType returns true if MatchType is present, Tag 574
func (m ExecutionReport) HasMatchType() bool {
	return m.Has(tag.MatchType)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m ExecutionReport) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasCustOrderCapacity returns true if CustOrderCapacity is present, Tag 582
func (m ExecutionReport) HasCustOrderCapacity() bool {
	return m.Has(tag.CustOrderCapacity)
}

//HasClOrdLinkID returns true if ClOrdLinkID is present, Tag 583
func (m ExecutionReport) HasClOrdLinkID() bool {
	return m.Has(tag.ClOrdLinkID)
}

//HasMassStatusReqID returns true if MassStatusReqID is present, Tag 584
func (m ExecutionReport) HasMassStatusReqID() bool {
	return m.Has(tag.MassStatusReqID)
}

//HasDayBookingInst returns true if DayBookingInst is present, Tag 589
func (m ExecutionReport) HasDayBookingInst() bool {
	return m.Has(tag.DayBookingInst)
}

//HasBookingUnit returns true if BookingUnit is present, Tag 590
func (m ExecutionReport) HasBookingUnit() bool {
	return m.Has(tag.BookingUnit)
}

//HasPreallocMethod returns true if PreallocMethod is present, Tag 591
func (m ExecutionReport) HasPreallocMethod() bool {
	return m.Has(tag.PreallocMethod)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m ExecutionReport) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasClearingFeeIndicator returns true if ClearingFeeIndicator is present, Tag 635
func (m ExecutionReport) HasClearingFeeIndicator() bool {
	return m.Has(tag.ClearingFeeIndicator)
}

//HasWorkingIndicator returns true if WorkingIndicator is present, Tag 636
func (m ExecutionReport) HasWorkingIndicator() bool {
	return m.Has(tag.WorkingIndicator)
}

//HasPriorityIndicator returns true if PriorityIndicator is present, Tag 638
func (m ExecutionReport) HasPriorityIndicator() bool {
	return m.Has(tag.PriorityIndicator)
}

//HasPriceImprovement returns true if PriceImprovement is present, Tag 639
func (m ExecutionReport) HasPriceImprovement() bool {
	return m.Has(tag.PriceImprovement)
}

//HasLastForwardPoints2 returns true if LastForwardPoints2 is present, Tag 641
func (m ExecutionReport) HasLastForwardPoints2() bool {
	return m.Has(tag.LastForwardPoints2)
}

//HasUnderlyingLastPx returns true if UnderlyingLastPx is present, Tag 651
func (m ExecutionReport) HasUnderlyingLastPx() bool {
	return m.Has(tag.UnderlyingLastPx)
}

//HasUnderlyingLastQty returns true if UnderlyingLastQty is present, Tag 652
func (m ExecutionReport) HasUnderlyingLastQty() bool {
	return m.Has(tag.UnderlyingLastQty)
}

//HasAcctIDSource returns true if AcctIDSource is present, Tag 660
func (m ExecutionReport) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
}

//HasBenchmarkPrice returns true if BenchmarkPrice is present, Tag 662
func (m ExecutionReport) HasBenchmarkPrice() bool {
	return m.Has(tag.BenchmarkPrice)
}

//HasBenchmarkPriceType returns true if BenchmarkPriceType is present, Tag 663
func (m ExecutionReport) HasBenchmarkPriceType() bool {
	return m.Has(tag.BenchmarkPriceType)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m ExecutionReport) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasLastParPx returns true if LastParPx is present, Tag 669
func (m ExecutionReport) HasLastParPx() bool {
	return m.Has(tag.LastParPx)
}

//HasPool returns true if Pool is present, Tag 691
func (m ExecutionReport) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasQuoteRespID returns true if QuoteRespID is present, Tag 693
func (m ExecutionReport) HasQuoteRespID() bool {
	return m.Has(tag.QuoteRespID)
}

//HasYieldRedemptionDate returns true if YieldRedemptionDate is present, Tag 696
func (m ExecutionReport) HasYieldRedemptionDate() bool {
	return m.Has(tag.YieldRedemptionDate)
}

//HasYieldRedemptionPrice returns true if YieldRedemptionPrice is present, Tag 697
func (m ExecutionReport) HasYieldRedemptionPrice() bool {
	return m.Has(tag.YieldRedemptionPrice)
}

//HasYieldRedemptionPriceType returns true if YieldRedemptionPriceType is present, Tag 698
func (m ExecutionReport) HasYieldRedemptionPriceType() bool {
	return m.Has(tag.YieldRedemptionPriceType)
}

//HasBenchmarkSecurityID returns true if BenchmarkSecurityID is present, Tag 699
func (m ExecutionReport) HasBenchmarkSecurityID() bool {
	return m.Has(tag.BenchmarkSecurityID)
}

//HasYieldCalcDate returns true if YieldCalcDate is present, Tag 701
func (m ExecutionReport) HasYieldCalcDate() bool {
	return m.Has(tag.YieldCalcDate)
}

//HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m ExecutionReport) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

//HasInterestAtMaturity returns true if InterestAtMaturity is present, Tag 738
func (m ExecutionReport) HasInterestAtMaturity() bool {
	return m.Has(tag.InterestAtMaturity)
}

//HasBenchmarkSecurityIDSource returns true if BenchmarkSecurityIDSource is present, Tag 761
func (m ExecutionReport) HasBenchmarkSecurityIDSource() bool {
	return m.Has(tag.BenchmarkSecurityIDSource)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m ExecutionReport) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasNoTrdRegTimestamps returns true if NoTrdRegTimestamps is present, Tag 768
func (m ExecutionReport) HasNoTrdRegTimestamps() bool {
	return m.Has(tag.NoTrdRegTimestamps)
}

//HasBookingType returns true if BookingType is present, Tag 775
func (m ExecutionReport) HasBookingType() bool {
	return m.Has(tag.BookingType)
}

//HasTerminationType returns true if TerminationType is present, Tag 788
func (m ExecutionReport) HasTerminationType() bool {
	return m.Has(tag.TerminationType)
}

//HasOrdStatusReqID returns true if OrdStatusReqID is present, Tag 790
func (m ExecutionReport) HasOrdStatusReqID() bool {
	return m.Has(tag.OrdStatusReqID)
}

//HasCopyMsgIndicator returns true if CopyMsgIndicator is present, Tag 797
func (m ExecutionReport) HasCopyMsgIndicator() bool {
	return m.Has(tag.CopyMsgIndicator)
}

//HasPriceDelta returns true if PriceDelta is present, Tag 811
func (m ExecutionReport) HasPriceDelta() bool {
	return m.Has(tag.PriceDelta)
}

//HasPegMoveType returns true if PegMoveType is present, Tag 835
func (m ExecutionReport) HasPegMoveType() bool {
	return m.Has(tag.PegMoveType)
}

//HasPegOffsetType returns true if PegOffsetType is present, Tag 836
func (m ExecutionReport) HasPegOffsetType() bool {
	return m.Has(tag.PegOffsetType)
}

//HasPegLimitType returns true if PegLimitType is present, Tag 837
func (m ExecutionReport) HasPegLimitType() bool {
	return m.Has(tag.PegLimitType)
}

//HasPegRoundDirection returns true if PegRoundDirection is present, Tag 838
func (m ExecutionReport) HasPegRoundDirection() bool {
	return m.Has(tag.PegRoundDirection)
}

//HasPeggedPrice returns true if PeggedPrice is present, Tag 839
func (m ExecutionReport) HasPeggedPrice() bool {
	return m.Has(tag.PeggedPrice)
}

//HasPegScope returns true if PegScope is present, Tag 840
func (m ExecutionReport) HasPegScope() bool {
	return m.Has(tag.PegScope)
}

//HasDiscretionMoveType returns true if DiscretionMoveType is present, Tag 841
func (m ExecutionReport) HasDiscretionMoveType() bool {
	return m.Has(tag.DiscretionMoveType)
}

//HasDiscretionOffsetType returns true if DiscretionOffsetType is present, Tag 842
func (m ExecutionReport) HasDiscretionOffsetType() bool {
	return m.Has(tag.DiscretionOffsetType)
}

//HasDiscretionLimitType returns true if DiscretionLimitType is present, Tag 843
func (m ExecutionReport) HasDiscretionLimitType() bool {
	return m.Has(tag.DiscretionLimitType)
}

//HasDiscretionRoundDirection returns true if DiscretionRoundDirection is present, Tag 844
func (m ExecutionReport) HasDiscretionRoundDirection() bool {
	return m.Has(tag.DiscretionRoundDirection)
}

//HasDiscretionPrice returns true if DiscretionPrice is present, Tag 845
func (m ExecutionReport) HasDiscretionPrice() bool {
	return m.Has(tag.DiscretionPrice)
}

//HasDiscretionScope returns true if DiscretionScope is present, Tag 846
func (m ExecutionReport) HasDiscretionScope() bool {
	return m.Has(tag.DiscretionScope)
}

//HasTargetStrategy returns true if TargetStrategy is present, Tag 847
func (m ExecutionReport) HasTargetStrategy() bool {
	return m.Has(tag.TargetStrategy)
}

//HasTargetStrategyParameters returns true if TargetStrategyParameters is present, Tag 848
func (m ExecutionReport) HasTargetStrategyParameters() bool {
	return m.Has(tag.TargetStrategyParameters)
}

//HasParticipationRate returns true if ParticipationRate is present, Tag 849
func (m ExecutionReport) HasParticipationRate() bool {
	return m.Has(tag.ParticipationRate)
}

//HasTargetStrategyPerformance returns true if TargetStrategyPerformance is present, Tag 850
func (m ExecutionReport) HasTargetStrategyPerformance() bool {
	return m.Has(tag.TargetStrategyPerformance)
}

//HasLastLiquidityInd returns true if LastLiquidityInd is present, Tag 851
func (m ExecutionReport) HasLastLiquidityInd() bool {
	return m.Has(tag.LastLiquidityInd)
}

//HasQtyType returns true if QtyType is present, Tag 854
func (m ExecutionReport) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m ExecutionReport) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m ExecutionReport) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m ExecutionReport) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m ExecutionReport) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m ExecutionReport) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasTrdMatchID returns true if TrdMatchID is present, Tag 880
func (m ExecutionReport) HasTrdMatchID() bool {
	return m.Has(tag.TrdMatchID)
}

//HasLastFragment returns true if LastFragment is present, Tag 893
func (m ExecutionReport) HasLastFragment() bool {
	return m.Has(tag.LastFragment)
}

//HasMarginRatio returns true if MarginRatio is present, Tag 898
func (m ExecutionReport) HasMarginRatio() bool {
	return m.Has(tag.MarginRatio)
}

//HasTotNumReports returns true if TotNumReports is present, Tag 911
func (m ExecutionReport) HasTotNumReports() bool {
	return m.Has(tag.TotNumReports)
}

//HasLastRptRequested returns true if LastRptRequested is present, Tag 912
func (m ExecutionReport) HasLastRptRequested() bool {
	return m.Has(tag.LastRptRequested)
}

//HasAgreementDesc returns true if AgreementDesc is present, Tag 913
func (m ExecutionReport) HasAgreementDesc() bool {
	return m.Has(tag.AgreementDesc)
}

//HasAgreementID returns true if AgreementID is present, Tag 914
func (m ExecutionReport) HasAgreementID() bool {
	return m.Has(tag.AgreementID)
}

//HasAgreementDate returns true if AgreementDate is present, Tag 915
func (m ExecutionReport) HasAgreementDate() bool {
	return m.Has(tag.AgreementDate)
}

//HasStartDate returns true if StartDate is present, Tag 916
func (m ExecutionReport) HasStartDate() bool {
	return m.Has(tag.StartDate)
}

//HasEndDate returns true if EndDate is present, Tag 917
func (m ExecutionReport) HasEndDate() bool {
	return m.Has(tag.EndDate)
}

//HasAgreementCurrency returns true if AgreementCurrency is present, Tag 918
func (m ExecutionReport) HasAgreementCurrency() bool {
	return m.Has(tag.AgreementCurrency)
}

//HasDeliveryType returns true if DeliveryType is present, Tag 919
func (m ExecutionReport) HasDeliveryType() bool {
	return m.Has(tag.DeliveryType)
}

//HasEndAccruedInterestAmt returns true if EndAccruedInterestAmt is present, Tag 920
func (m ExecutionReport) HasEndAccruedInterestAmt() bool {
	return m.Has(tag.EndAccruedInterestAmt)
}

//HasStartCash returns true if StartCash is present, Tag 921
func (m ExecutionReport) HasStartCash() bool {
	return m.Has(tag.StartCash)
}

//HasEndCash returns true if EndCash is present, Tag 922
func (m ExecutionReport) HasEndCash() bool {
	return m.Has(tag.EndCash)
}

//HasTimeBracket returns true if TimeBracket is present, Tag 943
func (m ExecutionReport) HasTimeBracket() bool {
	return m.Has(tag.TimeBracket)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m ExecutionReport) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//HasNoStrategyParameters returns true if NoStrategyParameters is present, Tag 957
func (m ExecutionReport) HasNoStrategyParameters() bool {
	return m.Has(tag.NoStrategyParameters)
}

//HasHostCrossID returns true if HostCrossID is present, Tag 961
func (m ExecutionReport) HasHostCrossID() bool {
	return m.Has(tag.HostCrossID)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m ExecutionReport) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m ExecutionReport) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m ExecutionReport) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m ExecutionReport) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m ExecutionReport) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m ExecutionReport) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m ExecutionReport) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m ExecutionReport) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m ExecutionReport) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m ExecutionReport) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasManualOrderIndicator returns true if ManualOrderIndicator is present, Tag 1028
func (m ExecutionReport) HasManualOrderIndicator() bool {
	return m.Has(tag.ManualOrderIndicator)
}

//HasCustDirectedOrder returns true if CustDirectedOrder is present, Tag 1029
func (m ExecutionReport) HasCustDirectedOrder() bool {
	return m.Has(tag.CustDirectedOrder)
}

//HasReceivedDeptID returns true if ReceivedDeptID is present, Tag 1030
func (m ExecutionReport) HasReceivedDeptID() bool {
	return m.Has(tag.ReceivedDeptID)
}

//HasCustOrderHandlingInst returns true if CustOrderHandlingInst is present, Tag 1031
func (m ExecutionReport) HasCustOrderHandlingInst() bool {
	return m.Has(tag.CustOrderHandlingInst)
}

//HasOrderHandlingInstSource returns true if OrderHandlingInstSource is present, Tag 1032
func (m ExecutionReport) HasOrderHandlingInstSource() bool {
	return m.Has(tag.OrderHandlingInstSource)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m ExecutionReport) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasCalculatedCcyLastQty returns true if CalculatedCcyLastQty is present, Tag 1056
func (m ExecutionReport) HasCalculatedCcyLastQty() bool {
	return m.Has(tag.CalculatedCcyLastQty)
}

//HasAggressorIndicator returns true if AggressorIndicator is present, Tag 1057
func (m ExecutionReport) HasAggressorIndicator() bool {
	return m.Has(tag.AggressorIndicator)
}

//HasLastSwapPoints returns true if LastSwapPoints is present, Tag 1071
func (m ExecutionReport) HasLastSwapPoints() bool {
	return m.Has(tag.LastSwapPoints)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m ExecutionReport) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//HasSecondaryDisplayQty returns true if SecondaryDisplayQty is present, Tag 1082
func (m ExecutionReport) HasSecondaryDisplayQty() bool {
	return m.Has(tag.SecondaryDisplayQty)
}

//HasDisplayWhen returns true if DisplayWhen is present, Tag 1083
func (m ExecutionReport) HasDisplayWhen() bool {
	return m.Has(tag.DisplayWhen)
}

//HasDisplayMethod returns true if DisplayMethod is present, Tag 1084
func (m ExecutionReport) HasDisplayMethod() bool {
	return m.Has(tag.DisplayMethod)
}

//HasDisplayLowQty returns true if DisplayLowQty is present, Tag 1085
func (m ExecutionReport) HasDisplayLowQty() bool {
	return m.Has(tag.DisplayLowQty)
}

//HasDisplayHighQty returns true if DisplayHighQty is present, Tag 1086
func (m ExecutionReport) HasDisplayHighQty() bool {
	return m.Has(tag.DisplayHighQty)
}

//HasDisplayMinIncr returns true if DisplayMinIncr is present, Tag 1087
func (m ExecutionReport) HasDisplayMinIncr() bool {
	return m.Has(tag.DisplayMinIncr)
}

//HasRefreshQty returns true if RefreshQty is present, Tag 1088
func (m ExecutionReport) HasRefreshQty() bool {
	return m.Has(tag.RefreshQty)
}

//HasMatchIncrement returns true if MatchIncrement is present, Tag 1089
func (m ExecutionReport) HasMatchIncrement() bool {
	return m.Has(tag.MatchIncrement)
}

//HasMaxPriceLevels returns true if MaxPriceLevels is present, Tag 1090
func (m ExecutionReport) HasMaxPriceLevels() bool {
	return m.Has(tag.MaxPriceLevels)
}

//HasPreTradeAnonymity returns true if PreTradeAnonymity is present, Tag 1091
func (m ExecutionReport) HasPreTradeAnonymity() bool {
	return m.Has(tag.PreTradeAnonymity)
}

//HasPriceProtectionScope returns true if PriceProtectionScope is present, Tag 1092
func (m ExecutionReport) HasPriceProtectionScope() bool {
	return m.Has(tag.PriceProtectionScope)
}

//HasLotType returns true if LotType is present, Tag 1093
func (m ExecutionReport) HasLotType() bool {
	return m.Has(tag.LotType)
}

//HasPegPriceType returns true if PegPriceType is present, Tag 1094
func (m ExecutionReport) HasPegPriceType() bool {
	return m.Has(tag.PegPriceType)
}

//HasPeggedRefPrice returns true if PeggedRefPrice is present, Tag 1095
func (m ExecutionReport) HasPeggedRefPrice() bool {
	return m.Has(tag.PeggedRefPrice)
}

//HasPegSecurityIDSource returns true if PegSecurityIDSource is present, Tag 1096
func (m ExecutionReport) HasPegSecurityIDSource() bool {
	return m.Has(tag.PegSecurityIDSource)
}

//HasPegSecurityID returns true if PegSecurityID is present, Tag 1097
func (m ExecutionReport) HasPegSecurityID() bool {
	return m.Has(tag.PegSecurityID)
}

//HasPegSymbol returns true if PegSymbol is present, Tag 1098
func (m ExecutionReport) HasPegSymbol() bool {
	return m.Has(tag.PegSymbol)
}

//HasPegSecurityDesc returns true if PegSecurityDesc is present, Tag 1099
func (m ExecutionReport) HasPegSecurityDesc() bool {
	return m.Has(tag.PegSecurityDesc)
}

//HasTriggerType returns true if TriggerType is present, Tag 1100
func (m ExecutionReport) HasTriggerType() bool {
	return m.Has(tag.TriggerType)
}

//HasTriggerAction returns true if TriggerAction is present, Tag 1101
func (m ExecutionReport) HasTriggerAction() bool {
	return m.Has(tag.TriggerAction)
}

//HasTriggerPrice returns true if TriggerPrice is present, Tag 1102
func (m ExecutionReport) HasTriggerPrice() bool {
	return m.Has(tag.TriggerPrice)
}

//HasTriggerSymbol returns true if TriggerSymbol is present, Tag 1103
func (m ExecutionReport) HasTriggerSymbol() bool {
	return m.Has(tag.TriggerSymbol)
}

//HasTriggerSecurityID returns true if TriggerSecurityID is present, Tag 1104
func (m ExecutionReport) HasTriggerSecurityID() bool {
	return m.Has(tag.TriggerSecurityID)
}

//HasTriggerSecurityIDSource returns true if TriggerSecurityIDSource is present, Tag 1105
func (m ExecutionReport) HasTriggerSecurityIDSource() bool {
	return m.Has(tag.TriggerSecurityIDSource)
}

//HasTriggerSecurityDesc returns true if TriggerSecurityDesc is present, Tag 1106
func (m ExecutionReport) HasTriggerSecurityDesc() bool {
	return m.Has(tag.TriggerSecurityDesc)
}

//HasTriggerPriceType returns true if TriggerPriceType is present, Tag 1107
func (m ExecutionReport) HasTriggerPriceType() bool {
	return m.Has(tag.TriggerPriceType)
}

//HasTriggerPriceTypeScope returns true if TriggerPriceTypeScope is present, Tag 1108
func (m ExecutionReport) HasTriggerPriceTypeScope() bool {
	return m.Has(tag.TriggerPriceTypeScope)
}

//HasTriggerPriceDirection returns true if TriggerPriceDirection is present, Tag 1109
func (m ExecutionReport) HasTriggerPriceDirection() bool {
	return m.Has(tag.TriggerPriceDirection)
}

//HasTriggerNewPrice returns true if TriggerNewPrice is present, Tag 1110
func (m ExecutionReport) HasTriggerNewPrice() bool {
	return m.Has(tag.TriggerNewPrice)
}

//HasTriggerOrderType returns true if TriggerOrderType is present, Tag 1111
func (m ExecutionReport) HasTriggerOrderType() bool {
	return m.Has(tag.TriggerOrderType)
}

//HasTriggerNewQty returns true if TriggerNewQty is present, Tag 1112
func (m ExecutionReport) HasTriggerNewQty() bool {
	return m.Has(tag.TriggerNewQty)
}

//HasTriggerTradingSessionID returns true if TriggerTradingSessionID is present, Tag 1113
func (m ExecutionReport) HasTriggerTradingSessionID() bool {
	return m.Has(tag.TriggerTradingSessionID)
}

//HasTriggerTradingSessionSubID returns true if TriggerTradingSessionSubID is present, Tag 1114
func (m ExecutionReport) HasTriggerTradingSessionSubID() bool {
	return m.Has(tag.TriggerTradingSessionSubID)
}

//HasOrderCategory returns true if OrderCategory is present, Tag 1115
func (m ExecutionReport) HasOrderCategory() bool {
	return m.Has(tag.OrderCategory)
}

//HasDisplayQty returns true if DisplayQty is present, Tag 1138
func (m ExecutionReport) HasDisplayQty() bool {
	return m.Has(tag.DisplayQty)
}

//HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146
func (m ExecutionReport) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

//HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147
func (m ExecutionReport) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

//HasSecurityGroup returns true if SecurityGroup is present, Tag 1151
func (m ExecutionReport) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

//HasApplID returns true if ApplID is present, Tag 1180
func (m ExecutionReport) HasApplID() bool {
	return m.Has(tag.ApplID)
}

//HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181
func (m ExecutionReport) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

//HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184
func (m ExecutionReport) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

//HasSecurityXML returns true if SecurityXML is present, Tag 1185
func (m ExecutionReport) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

//HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186
func (m ExecutionReport) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

//HasVolatility returns true if Volatility is present, Tag 1188
func (m ExecutionReport) HasVolatility() bool {
	return m.Has(tag.Volatility)
}

//HasTimeToExpiration returns true if TimeToExpiration is present, Tag 1189
func (m ExecutionReport) HasTimeToExpiration() bool {
	return m.Has(tag.TimeToExpiration)
}

//HasRiskFreeRate returns true if RiskFreeRate is present, Tag 1190
func (m ExecutionReport) HasRiskFreeRate() bool {
	return m.Has(tag.RiskFreeRate)
}

//HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191
func (m ExecutionReport) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

//HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192
func (m ExecutionReport) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

//HasSettlMethod returns true if SettlMethod is present, Tag 1193
func (m ExecutionReport) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

//HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194
func (m ExecutionReport) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

//HasOptPayAmount returns true if OptPayAmount is present, Tag 1195
func (m ExecutionReport) HasOptPayAmount() bool {
	return m.Has(tag.OptPayAmount)
}

//HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196
func (m ExecutionReport) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

//HasFuturesValuationMethod returns true if FuturesValuationMethod is present, Tag 1197
func (m ExecutionReport) HasFuturesValuationMethod() bool {
	return m.Has(tag.FuturesValuationMethod)
}

//HasListMethod returns true if ListMethod is present, Tag 1198
func (m ExecutionReport) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

//HasCapPrice returns true if CapPrice is present, Tag 1199
func (m ExecutionReport) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

//HasFloorPrice returns true if FloorPrice is present, Tag 1200
func (m ExecutionReport) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

//HasProductComplex returns true if ProductComplex is present, Tag 1227
func (m ExecutionReport) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

//HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242
func (m ExecutionReport) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

//HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244
func (m ExecutionReport) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

//HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350
func (m ExecutionReport) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

//HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352
func (m ExecutionReport) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

//HasTotNoFills returns true if TotNoFills is present, Tag 1361
func (m ExecutionReport) HasTotNoFills() bool {
	return m.Has(tag.TotNoFills)
}

//HasNoFills returns true if NoFills is present, Tag 1362
func (m ExecutionReport) HasNoFills() bool {
	return m.Has(tag.NoFills)
}

//HasDividendYield returns true if DividendYield is present, Tag 1380
func (m ExecutionReport) HasDividendYield() bool {
	return m.Has(tag.DividendYield)
}

//NoAllocs is a repeating group element, Tag 78
type NoAllocs struct {
	quickfix.Group
}

//SetAllocAccount sets AllocAccount, Tag 79
func (m NoAllocs) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

//SetAllocAcctIDSource sets AllocAcctIDSource, Tag 661
func (m NoAllocs) SetAllocAcctIDSource(v int) {
	m.Set(field.NewAllocAcctIDSource(v))
}

//SetAllocSettlCurrency sets AllocSettlCurrency, Tag 736
func (m NoAllocs) SetAllocSettlCurrency(v string) {
	m.Set(field.NewAllocSettlCurrency(v))
}

//SetIndividualAllocID sets IndividualAllocID, Tag 467
func (m NoAllocs) SetIndividualAllocID(v string) {
	m.Set(field.NewIndividualAllocID(v))
}

//SetNoNestedPartyIDs sets NoNestedPartyIDs, Tag 539
func (m NoAllocs) SetNoNestedPartyIDs(f NoNestedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetAllocQty sets AllocQty, Tag 80
func (m NoAllocs) SetAllocQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewAllocQty(value, scale))
}

//GetAllocAccount gets AllocAccount, Tag 79
func (m NoAllocs) GetAllocAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AllocAccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocAcctIDSource gets AllocAcctIDSource, Tag 661
func (m NoAllocs) GetAllocAcctIDSource() (v int, err quickfix.MessageRejectError) {
	var f field.AllocAcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocSettlCurrency gets AllocSettlCurrency, Tag 736
func (m NoAllocs) GetAllocSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.AllocSettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIndividualAllocID gets IndividualAllocID, Tag 467
func (m NoAllocs) GetIndividualAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.IndividualAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539
func (m NoAllocs) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAllocQty gets AllocQty, Tag 80
func (m NoAllocs) GetAllocQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.AllocQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//HasAllocAccount returns true if AllocAccount is present, Tag 79
func (m NoAllocs) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

//HasAllocAcctIDSource returns true if AllocAcctIDSource is present, Tag 661
func (m NoAllocs) HasAllocAcctIDSource() bool {
	return m.Has(tag.AllocAcctIDSource)
}

//HasAllocSettlCurrency returns true if AllocSettlCurrency is present, Tag 736
func (m NoAllocs) HasAllocSettlCurrency() bool {
	return m.Has(tag.AllocSettlCurrency)
}

//HasIndividualAllocID returns true if IndividualAllocID is present, Tag 467
func (m NoAllocs) HasIndividualAllocID() bool {
	return m.Has(tag.IndividualAllocID)
}

//HasNoNestedPartyIDs returns true if NoNestedPartyIDs is present, Tag 539
func (m NoAllocs) HasNoNestedPartyIDs() bool {
	return m.Has(tag.NoNestedPartyIDs)
}

//HasAllocQty returns true if AllocQty is present, Tag 80
func (m NoAllocs) HasAllocQty() bool {
	return m.Has(tag.AllocQty)
}

//NoNestedPartyIDs is a repeating group element, Tag 539
type NoNestedPartyIDs struct {
	quickfix.Group
}

//SetNestedPartyID sets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) SetNestedPartyID(v string) {
	m.Set(field.NewNestedPartyID(v))
}

//SetNestedPartyIDSource sets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) SetNestedPartyIDSource(v string) {
	m.Set(field.NewNestedPartyIDSource(v))
}

//SetNestedPartyRole sets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) SetNestedPartyRole(v int) {
	m.Set(field.NewNestedPartyRole(v))
}

//SetNoNestedPartySubIDs sets NoNestedPartySubIDs, Tag 804
func (m NoNestedPartyIDs) SetNoNestedPartySubIDs(f NoNestedPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNestedPartyID gets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) GetNestedPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartyIDSource gets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) GetNestedPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartyRole gets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) GetNestedPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.NestedPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoNestedPartySubIDs gets NoNestedPartySubIDs, Tag 804
func (m NoNestedPartyIDs) GetNoNestedPartySubIDs() (NoNestedPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNestedPartyID returns true if NestedPartyID is present, Tag 524
func (m NoNestedPartyIDs) HasNestedPartyID() bool {
	return m.Has(tag.NestedPartyID)
}

//HasNestedPartyIDSource returns true if NestedPartyIDSource is present, Tag 525
func (m NoNestedPartyIDs) HasNestedPartyIDSource() bool {
	return m.Has(tag.NestedPartyIDSource)
}

//HasNestedPartyRole returns true if NestedPartyRole is present, Tag 538
func (m NoNestedPartyIDs) HasNestedPartyRole() bool {
	return m.Has(tag.NestedPartyRole)
}

//HasNoNestedPartySubIDs returns true if NoNestedPartySubIDs is present, Tag 804
func (m NoNestedPartyIDs) HasNoNestedPartySubIDs() bool {
	return m.Has(tag.NoNestedPartySubIDs)
}

//NoNestedPartySubIDs is a repeating group element, Tag 804
type NoNestedPartySubIDs struct {
	quickfix.Group
}

//SetNestedPartySubID sets NestedPartySubID, Tag 545
func (m NoNestedPartySubIDs) SetNestedPartySubID(v string) {
	m.Set(field.NewNestedPartySubID(v))
}

//SetNestedPartySubIDType sets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) SetNestedPartySubIDType(v int) {
	m.Set(field.NewNestedPartySubIDType(v))
}

//GetNestedPartySubID gets NestedPartySubID, Tag 545
func (m NoNestedPartySubIDs) GetNestedPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartySubIDType gets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) GetNestedPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.NestedPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasNestedPartySubID returns true if NestedPartySubID is present, Tag 545
func (m NoNestedPartySubIDs) HasNestedPartySubID() bool {
	return m.Has(tag.NestedPartySubID)
}

//HasNestedPartySubIDType returns true if NestedPartySubIDType is present, Tag 805
func (m NoNestedPartySubIDs) HasNestedPartySubIDType() bool {
	return m.Has(tag.NestedPartySubIDType)
}

//NoNestedPartySubIDsRepeatingGroup is a repeating group, Tag 804
type NoNestedPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNestedPartySubIDsRepeatingGroup returns an initialized, NoNestedPartySubIDsRepeatingGroup
func NewNoNestedPartySubIDsRepeatingGroup() NoNestedPartySubIDsRepeatingGroup {
	return NoNestedPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartySubID), quickfix.GroupElement(tag.NestedPartySubIDType)})}
}

//Add create and append a new NoNestedPartySubIDs to this group
func (m NoNestedPartySubIDsRepeatingGroup) Add() NoNestedPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartySubIDs{g}
}

//Get returns the ith NoNestedPartySubIDs in the NoNestedPartySubIDsRepeatinGroup
func (m NoNestedPartySubIDsRepeatingGroup) Get(i int) NoNestedPartySubIDs {
	return NoNestedPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoNestedPartyIDsRepeatingGroup is a repeating group, Tag 539
type NoNestedPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNestedPartyIDsRepeatingGroup returns an initialized, NoNestedPartyIDsRepeatingGroup
func NewNoNestedPartyIDsRepeatingGroup() NoNestedPartyIDsRepeatingGroup {
	return NoNestedPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartyID), quickfix.GroupElement(tag.NestedPartyIDSource), quickfix.GroupElement(tag.NestedPartyRole), NewNoNestedPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoNestedPartyIDs to this group
func (m NoNestedPartyIDsRepeatingGroup) Add() NoNestedPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartyIDs{g}
}

//Get returns the ith NoNestedPartyIDs in the NoNestedPartyIDsRepeatinGroup
func (m NoNestedPartyIDsRepeatingGroup) Get(i int) NoNestedPartyIDs {
	return NoNestedPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocAcctIDSource), quickfix.GroupElement(tag.AllocSettlCurrency), quickfix.GroupElement(tag.IndividualAllocID), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.AllocQty)})}
}

//Add create and append a new NoAllocs to this group
func (m NoAllocsRepeatingGroup) Add() NoAllocs {
	g := m.RepeatingGroup.Add()
	return NoAllocs{g}
}

//Get returns the ith NoAllocs in the NoAllocsRepeatinGroup
func (m NoAllocsRepeatingGroup) Get(i int) NoAllocs {
	return NoAllocs{m.RepeatingGroup.Get(i)}
}

//NoMiscFees is a repeating group element, Tag 136
type NoMiscFees struct {
	quickfix.Group
}

//SetMiscFeeAmt sets MiscFeeAmt, Tag 137
func (m NoMiscFees) SetMiscFeeAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewMiscFeeAmt(value, scale))
}

//SetMiscFeeCurr sets MiscFeeCurr, Tag 138
func (m NoMiscFees) SetMiscFeeCurr(v string) {
	m.Set(field.NewMiscFeeCurr(v))
}

//SetMiscFeeType sets MiscFeeType, Tag 139
func (m NoMiscFees) SetMiscFeeType(v enum.MiscFeeType) {
	m.Set(field.NewMiscFeeType(v))
}

//SetMiscFeeBasis sets MiscFeeBasis, Tag 891
func (m NoMiscFees) SetMiscFeeBasis(v enum.MiscFeeBasis) {
	m.Set(field.NewMiscFeeBasis(v))
}

//GetMiscFeeAmt gets MiscFeeAmt, Tag 137
func (m NoMiscFees) GetMiscFeeAmt() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.MiscFeeAmtField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetMiscFeeCurr gets MiscFeeCurr, Tag 138
func (m NoMiscFees) GetMiscFeeCurr() (v string, err quickfix.MessageRejectError) {
	var f field.MiscFeeCurrField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMiscFeeType gets MiscFeeType, Tag 139
func (m NoMiscFees) GetMiscFeeType() (v enum.MiscFeeType, err quickfix.MessageRejectError) {
	var f field.MiscFeeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMiscFeeBasis gets MiscFeeBasis, Tag 891
func (m NoMiscFees) GetMiscFeeBasis() (v enum.MiscFeeBasis, err quickfix.MessageRejectError) {
	var f field.MiscFeeBasisField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasMiscFeeAmt returns true if MiscFeeAmt is present, Tag 137
func (m NoMiscFees) HasMiscFeeAmt() bool {
	return m.Has(tag.MiscFeeAmt)
}

//HasMiscFeeCurr returns true if MiscFeeCurr is present, Tag 138
func (m NoMiscFees) HasMiscFeeCurr() bool {
	return m.Has(tag.MiscFeeCurr)
}

//HasMiscFeeType returns true if MiscFeeType is present, Tag 139
func (m NoMiscFees) HasMiscFeeType() bool {
	return m.Has(tag.MiscFeeType)
}

//HasMiscFeeBasis returns true if MiscFeeBasis is present, Tag 891
func (m NoMiscFees) HasMiscFeeBasis() bool {
	return m.Has(tag.MiscFeeBasis)
}

//NoMiscFeesRepeatingGroup is a repeating group, Tag 136
type NoMiscFeesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMiscFeesRepeatingGroup returns an initialized, NoMiscFeesRepeatingGroup
func NewNoMiscFeesRepeatingGroup() NoMiscFeesRepeatingGroup {
	return NoMiscFeesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMiscFees,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MiscFeeAmt), quickfix.GroupElement(tag.MiscFeeCurr), quickfix.GroupElement(tag.MiscFeeType), quickfix.GroupElement(tag.MiscFeeBasis)})}
}

//Add create and append a new NoMiscFees to this group
func (m NoMiscFeesRepeatingGroup) Add() NoMiscFees {
	g := m.RepeatingGroup.Add()
	return NoMiscFees{g}
}

//Get returns the ith NoMiscFees in the NoMiscFeesRepeatinGroup
func (m NoMiscFeesRepeatingGroup) Get(i int) NoMiscFees {
	return NoMiscFees{m.RepeatingGroup.Get(i)}
}

//NoStipulations is a repeating group element, Tag 232
type NoStipulations struct {
	quickfix.Group
}

//SetStipulationType sets StipulationType, Tag 233
func (m NoStipulations) SetStipulationType(v enum.StipulationType) {
	m.Set(field.NewStipulationType(v))
}

//SetStipulationValue sets StipulationValue, Tag 234
func (m NoStipulations) SetStipulationValue(v string) {
	m.Set(field.NewStipulationValue(v))
}

//GetStipulationType gets StipulationType, Tag 233
func (m NoStipulations) GetStipulationType() (v enum.StipulationType, err quickfix.MessageRejectError) {
	var f field.StipulationTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStipulationValue gets StipulationValue, Tag 234
func (m NoStipulations) GetStipulationValue() (v string, err quickfix.MessageRejectError) {
	var f field.StipulationValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasStipulationType returns true if StipulationType is present, Tag 233
func (m NoStipulations) HasStipulationType() bool {
	return m.Has(tag.StipulationType)
}

//HasStipulationValue returns true if StipulationValue is present, Tag 234
func (m NoStipulations) HasStipulationValue() bool {
	return m.Has(tag.StipulationValue)
}

//NoStipulationsRepeatingGroup is a repeating group, Tag 232
type NoStipulationsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoStipulationsRepeatingGroup returns an initialized, NoStipulationsRepeatingGroup
func NewNoStipulationsRepeatingGroup() NoStipulationsRepeatingGroup {
	return NoStipulationsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoStipulations,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StipulationType), quickfix.GroupElement(tag.StipulationValue)})}
}

//Add create and append a new NoStipulations to this group
func (m NoStipulationsRepeatingGroup) Add() NoStipulations {
	g := m.RepeatingGroup.Add()
	return NoStipulations{g}
}

//Get returns the ith NoStipulations in the NoStipulationsRepeatinGroup
func (m NoStipulationsRepeatingGroup) Get(i int) NoStipulations {
	return NoStipulations{m.RepeatingGroup.Get(i)}
}

//NoContraBrokers is a repeating group element, Tag 382
type NoContraBrokers struct {
	quickfix.Group
}

//SetContraBroker sets ContraBroker, Tag 375
func (m NoContraBrokers) SetContraBroker(v string) {
	m.Set(field.NewContraBroker(v))
}

//SetContraTrader sets ContraTrader, Tag 337
func (m NoContraBrokers) SetContraTrader(v string) {
	m.Set(field.NewContraTrader(v))
}

//SetContraTradeQty sets ContraTradeQty, Tag 437
func (m NoContraBrokers) SetContraTradeQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewContraTradeQty(value, scale))
}

//SetContraTradeTime sets ContraTradeTime, Tag 438
func (m NoContraBrokers) SetContraTradeTime(v time.Time) {
	m.Set(field.NewContraTradeTime(v))
}

//SetContraLegRefID sets ContraLegRefID, Tag 655
func (m NoContraBrokers) SetContraLegRefID(v string) {
	m.Set(field.NewContraLegRefID(v))
}

//GetContraBroker gets ContraBroker, Tag 375
func (m NoContraBrokers) GetContraBroker() (v string, err quickfix.MessageRejectError) {
	var f field.ContraBrokerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContraTrader gets ContraTrader, Tag 337
func (m NoContraBrokers) GetContraTrader() (v string, err quickfix.MessageRejectError) {
	var f field.ContraTraderField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContraTradeQty gets ContraTradeQty, Tag 437
func (m NoContraBrokers) GetContraTradeQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.ContraTradeQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetContraTradeTime gets ContraTradeTime, Tag 438
func (m NoContraBrokers) GetContraTradeTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ContraTradeTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContraLegRefID gets ContraLegRefID, Tag 655
func (m NoContraBrokers) GetContraLegRefID() (v string, err quickfix.MessageRejectError) {
	var f field.ContraLegRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasContraBroker returns true if ContraBroker is present, Tag 375
func (m NoContraBrokers) HasContraBroker() bool {
	return m.Has(tag.ContraBroker)
}

//HasContraTrader returns true if ContraTrader is present, Tag 337
func (m NoContraBrokers) HasContraTrader() bool {
	return m.Has(tag.ContraTrader)
}

//HasContraTradeQty returns true if ContraTradeQty is present, Tag 437
func (m NoContraBrokers) HasContraTradeQty() bool {
	return m.Has(tag.ContraTradeQty)
}

//HasContraTradeTime returns true if ContraTradeTime is present, Tag 438
func (m NoContraBrokers) HasContraTradeTime() bool {
	return m.Has(tag.ContraTradeTime)
}

//HasContraLegRefID returns true if ContraLegRefID is present, Tag 655
func (m NoContraBrokers) HasContraLegRefID() bool {
	return m.Has(tag.ContraLegRefID)
}

//NoContraBrokersRepeatingGroup is a repeating group, Tag 382
type NoContraBrokersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoContraBrokersRepeatingGroup returns an initialized, NoContraBrokersRepeatingGroup
func NewNoContraBrokersRepeatingGroup() NoContraBrokersRepeatingGroup {
	return NoContraBrokersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoContraBrokers,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ContraBroker), quickfix.GroupElement(tag.ContraTrader), quickfix.GroupElement(tag.ContraTradeQty), quickfix.GroupElement(tag.ContraTradeTime), quickfix.GroupElement(tag.ContraLegRefID)})}
}

//Add create and append a new NoContraBrokers to this group
func (m NoContraBrokersRepeatingGroup) Add() NoContraBrokers {
	g := m.RepeatingGroup.Add()
	return NoContraBrokers{g}
}

//Get returns the ith NoContraBrokers in the NoContraBrokersRepeatinGroup
func (m NoContraBrokersRepeatingGroup) Get(i int) NoContraBrokers {
	return NoContraBrokers{m.RepeatingGroup.Get(i)}
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

//SetNoPartySubIDs sets NoPartySubIDs, Tag 802
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartySubIDs gets NoPartySubIDs, Tag 802
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

//HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

//HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

//HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

//NoPartySubIDs is a repeating group element, Tag 802
type NoPartySubIDs struct {
	quickfix.Group
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//SetPartySubIDType sets PartySubIDType, Tag 803
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartySubIDType gets PartySubIDType, Tag 803
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//HasPartySubIDType returns true if PartySubIDType is present, Tag 803
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

//NoPartySubIDsRepeatingGroup is a repeating group, Tag 802
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

//Add create and append a new NoPartySubIDs to this group
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

//Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoPartyIDs to this group
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

//Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoSecurityAltID is a repeating group element, Tag 454
type NoSecurityAltID struct {
	quickfix.Group
}

//SetSecurityAltID sets SecurityAltID, Tag 455
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

//SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

//GetSecurityAltID gets SecurityAltID, Tag 455
func (m NoSecurityAltID) GetSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasSecurityAltID returns true if SecurityAltID is present, Tag 455
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

//HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

//NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

//Add create and append a new NoSecurityAltID to this group
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

//Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoContAmts is a repeating group element, Tag 518
type NoContAmts struct {
	quickfix.Group
}

//SetContAmtType sets ContAmtType, Tag 519
func (m NoContAmts) SetContAmtType(v enum.ContAmtType) {
	m.Set(field.NewContAmtType(v))
}

//SetContAmtValue sets ContAmtValue, Tag 520
func (m NoContAmts) SetContAmtValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewContAmtValue(value, scale))
}

//SetContAmtCurr sets ContAmtCurr, Tag 521
func (m NoContAmts) SetContAmtCurr(v string) {
	m.Set(field.NewContAmtCurr(v))
}

//GetContAmtType gets ContAmtType, Tag 519
func (m NoContAmts) GetContAmtType() (v enum.ContAmtType, err quickfix.MessageRejectError) {
	var f field.ContAmtTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContAmtValue gets ContAmtValue, Tag 520
func (m NoContAmts) GetContAmtValue() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.ContAmtValueField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetContAmtCurr gets ContAmtCurr, Tag 521
func (m NoContAmts) GetContAmtCurr() (v string, err quickfix.MessageRejectError) {
	var f field.ContAmtCurrField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasContAmtType returns true if ContAmtType is present, Tag 519
func (m NoContAmts) HasContAmtType() bool {
	return m.Has(tag.ContAmtType)
}

//HasContAmtValue returns true if ContAmtValue is present, Tag 520
func (m NoContAmts) HasContAmtValue() bool {
	return m.Has(tag.ContAmtValue)
}

//HasContAmtCurr returns true if ContAmtCurr is present, Tag 521
func (m NoContAmts) HasContAmtCurr() bool {
	return m.Has(tag.ContAmtCurr)
}

//NoContAmtsRepeatingGroup is a repeating group, Tag 518
type NoContAmtsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoContAmtsRepeatingGroup returns an initialized, NoContAmtsRepeatingGroup
func NewNoContAmtsRepeatingGroup() NoContAmtsRepeatingGroup {
	return NoContAmtsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoContAmts,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ContAmtType), quickfix.GroupElement(tag.ContAmtValue), quickfix.GroupElement(tag.ContAmtCurr)})}
}

//Add create and append a new NoContAmts to this group
func (m NoContAmtsRepeatingGroup) Add() NoContAmts {
	g := m.RepeatingGroup.Add()
	return NoContAmts{g}
}

//Get returns the ith NoContAmts in the NoContAmtsRepeatinGroup
func (m NoContAmtsRepeatingGroup) Get(i int) NoContAmts {
	return NoContAmts{m.RepeatingGroup.Get(i)}
}

//NoLegs is a repeating group element, Tag 555
type NoLegs struct {
	quickfix.Group
}

//SetLegSymbol sets LegSymbol, Tag 600
func (m NoLegs) SetLegSymbol(v string) {
	m.Set(field.NewLegSymbol(v))
}

//SetLegSymbolSfx sets LegSymbolSfx, Tag 601
func (m NoLegs) SetLegSymbolSfx(v string) {
	m.Set(field.NewLegSymbolSfx(v))
}

//SetLegSecurityID sets LegSecurityID, Tag 602
func (m NoLegs) SetLegSecurityID(v string) {
	m.Set(field.NewLegSecurityID(v))
}

//SetLegSecurityIDSource sets LegSecurityIDSource, Tag 603
func (m NoLegs) SetLegSecurityIDSource(v string) {
	m.Set(field.NewLegSecurityIDSource(v))
}

//SetNoLegSecurityAltID sets NoLegSecurityAltID, Tag 604
func (m NoLegs) SetNoLegSecurityAltID(f NoLegSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegProduct sets LegProduct, Tag 607
func (m NoLegs) SetLegProduct(v int) {
	m.Set(field.NewLegProduct(v))
}

//SetLegCFICode sets LegCFICode, Tag 608
func (m NoLegs) SetLegCFICode(v string) {
	m.Set(field.NewLegCFICode(v))
}

//SetLegSecurityType sets LegSecurityType, Tag 609
func (m NoLegs) SetLegSecurityType(v string) {
	m.Set(field.NewLegSecurityType(v))
}

//SetLegSecuritySubType sets LegSecuritySubType, Tag 764
func (m NoLegs) SetLegSecuritySubType(v string) {
	m.Set(field.NewLegSecuritySubType(v))
}

//SetLegMaturityMonthYear sets LegMaturityMonthYear, Tag 610
func (m NoLegs) SetLegMaturityMonthYear(v string) {
	m.Set(field.NewLegMaturityMonthYear(v))
}

//SetLegMaturityDate sets LegMaturityDate, Tag 611
func (m NoLegs) SetLegMaturityDate(v string) {
	m.Set(field.NewLegMaturityDate(v))
}

//SetLegCouponPaymentDate sets LegCouponPaymentDate, Tag 248
func (m NoLegs) SetLegCouponPaymentDate(v string) {
	m.Set(field.NewLegCouponPaymentDate(v))
}

//SetLegIssueDate sets LegIssueDate, Tag 249
func (m NoLegs) SetLegIssueDate(v string) {
	m.Set(field.NewLegIssueDate(v))
}

//SetLegRepoCollateralSecurityType sets LegRepoCollateralSecurityType, Tag 250
func (m NoLegs) SetLegRepoCollateralSecurityType(v int) {
	m.Set(field.NewLegRepoCollateralSecurityType(v))
}

//SetLegRepurchaseTerm sets LegRepurchaseTerm, Tag 251
func (m NoLegs) SetLegRepurchaseTerm(v int) {
	m.Set(field.NewLegRepurchaseTerm(v))
}

//SetLegRepurchaseRate sets LegRepurchaseRate, Tag 252
func (m NoLegs) SetLegRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRepurchaseRate(value, scale))
}

//SetLegFactor sets LegFactor, Tag 253
func (m NoLegs) SetLegFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegFactor(value, scale))
}

//SetLegCreditRating sets LegCreditRating, Tag 257
func (m NoLegs) SetLegCreditRating(v string) {
	m.Set(field.NewLegCreditRating(v))
}

//SetLegInstrRegistry sets LegInstrRegistry, Tag 599
func (m NoLegs) SetLegInstrRegistry(v string) {
	m.Set(field.NewLegInstrRegistry(v))
}

//SetLegCountryOfIssue sets LegCountryOfIssue, Tag 596
func (m NoLegs) SetLegCountryOfIssue(v string) {
	m.Set(field.NewLegCountryOfIssue(v))
}

//SetLegStateOrProvinceOfIssue sets LegStateOrProvinceOfIssue, Tag 597
func (m NoLegs) SetLegStateOrProvinceOfIssue(v string) {
	m.Set(field.NewLegStateOrProvinceOfIssue(v))
}

//SetLegLocaleOfIssue sets LegLocaleOfIssue, Tag 598
func (m NoLegs) SetLegLocaleOfIssue(v string) {
	m.Set(field.NewLegLocaleOfIssue(v))
}

//SetLegRedemptionDate sets LegRedemptionDate, Tag 254
func (m NoLegs) SetLegRedemptionDate(v string) {
	m.Set(field.NewLegRedemptionDate(v))
}

//SetLegStrikePrice sets LegStrikePrice, Tag 612
func (m NoLegs) SetLegStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegStrikePrice(value, scale))
}

//SetLegStrikeCurrency sets LegStrikeCurrency, Tag 942
func (m NoLegs) SetLegStrikeCurrency(v string) {
	m.Set(field.NewLegStrikeCurrency(v))
}

//SetLegOptAttribute sets LegOptAttribute, Tag 613
func (m NoLegs) SetLegOptAttribute(v string) {
	m.Set(field.NewLegOptAttribute(v))
}

//SetLegContractMultiplier sets LegContractMultiplier, Tag 614
func (m NoLegs) SetLegContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegContractMultiplier(value, scale))
}

//SetLegCouponRate sets LegCouponRate, Tag 615
func (m NoLegs) SetLegCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegCouponRate(value, scale))
}

//SetLegSecurityExchange sets LegSecurityExchange, Tag 616
func (m NoLegs) SetLegSecurityExchange(v string) {
	m.Set(field.NewLegSecurityExchange(v))
}

//SetLegIssuer sets LegIssuer, Tag 617
func (m NoLegs) SetLegIssuer(v string) {
	m.Set(field.NewLegIssuer(v))
}

//SetEncodedLegIssuerLen sets EncodedLegIssuerLen, Tag 618
func (m NoLegs) SetEncodedLegIssuerLen(v int) {
	m.Set(field.NewEncodedLegIssuerLen(v))
}

//SetEncodedLegIssuer sets EncodedLegIssuer, Tag 619
func (m NoLegs) SetEncodedLegIssuer(v string) {
	m.Set(field.NewEncodedLegIssuer(v))
}

//SetLegSecurityDesc sets LegSecurityDesc, Tag 620
func (m NoLegs) SetLegSecurityDesc(v string) {
	m.Set(field.NewLegSecurityDesc(v))
}

//SetEncodedLegSecurityDescLen sets EncodedLegSecurityDescLen, Tag 621
func (m NoLegs) SetEncodedLegSecurityDescLen(v int) {
	m.Set(field.NewEncodedLegSecurityDescLen(v))
}

//SetEncodedLegSecurityDesc sets EncodedLegSecurityDesc, Tag 622
func (m NoLegs) SetEncodedLegSecurityDesc(v string) {
	m.Set(field.NewEncodedLegSecurityDesc(v))
}

//SetLegRatioQty sets LegRatioQty, Tag 623
func (m NoLegs) SetLegRatioQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRatioQty(value, scale))
}

//SetLegSide sets LegSide, Tag 624
func (m NoLegs) SetLegSide(v string) {
	m.Set(field.NewLegSide(v))
}

//SetLegCurrency sets LegCurrency, Tag 556
func (m NoLegs) SetLegCurrency(v string) {
	m.Set(field.NewLegCurrency(v))
}

//SetLegPool sets LegPool, Tag 740
func (m NoLegs) SetLegPool(v string) {
	m.Set(field.NewLegPool(v))
}

//SetLegDatedDate sets LegDatedDate, Tag 739
func (m NoLegs) SetLegDatedDate(v string) {
	m.Set(field.NewLegDatedDate(v))
}

//SetLegContractSettlMonth sets LegContractSettlMonth, Tag 955
func (m NoLegs) SetLegContractSettlMonth(v string) {
	m.Set(field.NewLegContractSettlMonth(v))
}

//SetLegInterestAccrualDate sets LegInterestAccrualDate, Tag 956
func (m NoLegs) SetLegInterestAccrualDate(v string) {
	m.Set(field.NewLegInterestAccrualDate(v))
}

//SetLegUnitOfMeasure sets LegUnitOfMeasure, Tag 999
func (m NoLegs) SetLegUnitOfMeasure(v string) {
	m.Set(field.NewLegUnitOfMeasure(v))
}

//SetLegTimeUnit sets LegTimeUnit, Tag 1001
func (m NoLegs) SetLegTimeUnit(v string) {
	m.Set(field.NewLegTimeUnit(v))
}

//SetLegOptionRatio sets LegOptionRatio, Tag 1017
func (m NoLegs) SetLegOptionRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegOptionRatio(value, scale))
}

//SetLegPrice sets LegPrice, Tag 566
func (m NoLegs) SetLegPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPrice(value, scale))
}

//SetLegMaturityTime sets LegMaturityTime, Tag 1212
func (m NoLegs) SetLegMaturityTime(v string) {
	m.Set(field.NewLegMaturityTime(v))
}

//SetLegPutOrCall sets LegPutOrCall, Tag 1358
func (m NoLegs) SetLegPutOrCall(v int) {
	m.Set(field.NewLegPutOrCall(v))
}

//SetLegExerciseStyle sets LegExerciseStyle, Tag 1420
func (m NoLegs) SetLegExerciseStyle(v int) {
	m.Set(field.NewLegExerciseStyle(v))
}

//SetLegUnitOfMeasureQty sets LegUnitOfMeasureQty, Tag 1224
func (m NoLegs) SetLegUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegUnitOfMeasureQty(value, scale))
}

//SetLegPriceUnitOfMeasure sets LegPriceUnitOfMeasure, Tag 1421
func (m NoLegs) SetLegPriceUnitOfMeasure(v string) {
	m.Set(field.NewLegPriceUnitOfMeasure(v))
}

//SetLegPriceUnitOfMeasureQty sets LegPriceUnitOfMeasureQty, Tag 1422
func (m NoLegs) SetLegPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPriceUnitOfMeasureQty(value, scale))
}

//SetLegQty sets LegQty, Tag 687
func (m NoLegs) SetLegQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegQty(value, scale))
}

//SetLegSwapType sets LegSwapType, Tag 690
func (m NoLegs) SetLegSwapType(v enum.LegSwapType) {
	m.Set(field.NewLegSwapType(v))
}

//SetNoLegStipulations sets NoLegStipulations, Tag 683
func (m NoLegs) SetNoLegStipulations(f NoLegStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegPositionEffect sets LegPositionEffect, Tag 564
func (m NoLegs) SetLegPositionEffect(v string) {
	m.Set(field.NewLegPositionEffect(v))
}

//SetLegCoveredOrUncovered sets LegCoveredOrUncovered, Tag 565
func (m NoLegs) SetLegCoveredOrUncovered(v int) {
	m.Set(field.NewLegCoveredOrUncovered(v))
}

//SetLegRefID sets LegRefID, Tag 654
func (m NoLegs) SetLegRefID(v string) {
	m.Set(field.NewLegRefID(v))
}

//SetLegSettlType sets LegSettlType, Tag 587
func (m NoLegs) SetLegSettlType(v string) {
	m.Set(field.NewLegSettlType(v))
}

//SetLegSettlDate sets LegSettlDate, Tag 588
func (m NoLegs) SetLegSettlDate(v string) {
	m.Set(field.NewLegSettlDate(v))
}

//SetLegLastPx sets LegLastPx, Tag 637
func (m NoLegs) SetLegLastPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegLastPx(value, scale))
}

//SetLegOrderQty sets LegOrderQty, Tag 685
func (m NoLegs) SetLegOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegOrderQty(value, scale))
}

//SetLegSettlCurrency sets LegSettlCurrency, Tag 675
func (m NoLegs) SetLegSettlCurrency(v string) {
	m.Set(field.NewLegSettlCurrency(v))
}

//SetLegLastForwardPoints sets LegLastForwardPoints, Tag 1073
func (m NoLegs) SetLegLastForwardPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegLastForwardPoints(value, scale))
}

//SetLegCalculatedCcyLastQty sets LegCalculatedCcyLastQty, Tag 1074
func (m NoLegs) SetLegCalculatedCcyLastQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegCalculatedCcyLastQty(value, scale))
}

//SetLegGrossTradeAmt sets LegGrossTradeAmt, Tag 1075
func (m NoLegs) SetLegGrossTradeAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegGrossTradeAmt(value, scale))
}

//SetNoNested3PartyIDs sets NoNested3PartyIDs, Tag 948
func (m NoLegs) SetNoNested3PartyIDs(f NoNested3PartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegAllocID sets LegAllocID, Tag 1366
func (m NoLegs) SetLegAllocID(v string) {
	m.Set(field.NewLegAllocID(v))
}

//SetNoLegAllocs sets NoLegAllocs, Tag 670
func (m NoLegs) SetNoLegAllocs(f NoLegAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegVolatility sets LegVolatility, Tag 1379
func (m NoLegs) SetLegVolatility(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegVolatility(value, scale))
}

//SetLegDividendYield sets LegDividendYield, Tag 1381
func (m NoLegs) SetLegDividendYield(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegDividendYield(value, scale))
}

//SetLegCurrencyRatio sets LegCurrencyRatio, Tag 1383
func (m NoLegs) SetLegCurrencyRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegCurrencyRatio(value, scale))
}

//SetLegExecInst sets LegExecInst, Tag 1384
func (m NoLegs) SetLegExecInst(v string) {
	m.Set(field.NewLegExecInst(v))
}

//SetLegLastQty sets LegLastQty, Tag 1418
func (m NoLegs) SetLegLastQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegLastQty(value, scale))
}

//GetLegSymbol gets LegSymbol, Tag 600
func (m NoLegs) GetLegSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSymbolSfx gets LegSymbolSfx, Tag 601
func (m NoLegs) GetLegSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecurityID gets LegSecurityID, Tag 602
func (m NoLegs) GetLegSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecurityIDSource gets LegSecurityIDSource, Tag 603
func (m NoLegs) GetLegSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoLegSecurityAltID gets NoLegSecurityAltID, Tag 604
func (m NoLegs) GetNoLegSecurityAltID() (NoLegSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegProduct gets LegProduct, Tag 607
func (m NoLegs) GetLegProduct() (v int, err quickfix.MessageRejectError) {
	var f field.LegProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegCFICode gets LegCFICode, Tag 608
func (m NoLegs) GetLegCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.LegCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecurityType gets LegSecurityType, Tag 609
func (m NoLegs) GetLegSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecuritySubType gets LegSecuritySubType, Tag 764
func (m NoLegs) GetLegSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegMaturityMonthYear gets LegMaturityMonthYear, Tag 610
func (m NoLegs) GetLegMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegMaturityDate gets LegMaturityDate, Tag 611
func (m NoLegs) GetLegMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegCouponPaymentDate gets LegCouponPaymentDate, Tag 248
func (m NoLegs) GetLegCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegIssueDate gets LegIssueDate, Tag 249
func (m NoLegs) GetLegIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegRepoCollateralSecurityType gets LegRepoCollateralSecurityType, Tag 250
func (m NoLegs) GetLegRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegRepurchaseTerm gets LegRepurchaseTerm, Tag 251
func (m NoLegs) GetLegRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegRepurchaseRate gets LegRepurchaseRate, Tag 252
func (m NoLegs) GetLegRepurchaseRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegFactor gets LegFactor, Tag 253
func (m NoLegs) GetLegFactor() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegFactorField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegCreditRating gets LegCreditRating, Tag 257
func (m NoLegs) GetLegCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.LegCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegInstrRegistry gets LegInstrRegistry, Tag 599
func (m NoLegs) GetLegInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.LegInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegCountryOfIssue gets LegCountryOfIssue, Tag 596
func (m NoLegs) GetLegCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegStateOrProvinceOfIssue gets LegStateOrProvinceOfIssue, Tag 597
func (m NoLegs) GetLegStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegLocaleOfIssue gets LegLocaleOfIssue, Tag 598
func (m NoLegs) GetLegLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegRedemptionDate gets LegRedemptionDate, Tag 254
func (m NoLegs) GetLegRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegStrikePrice gets LegStrikePrice, Tag 612
func (m NoLegs) GetLegStrikePrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegStrikeCurrency gets LegStrikeCurrency, Tag 942
func (m NoLegs) GetLegStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegOptAttribute gets LegOptAttribute, Tag 613
func (m NoLegs) GetLegOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.LegOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegContractMultiplier gets LegContractMultiplier, Tag 614
func (m NoLegs) GetLegContractMultiplier() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegCouponRate gets LegCouponRate, Tag 615
func (m NoLegs) GetLegCouponRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegSecurityExchange gets LegSecurityExchange, Tag 616
func (m NoLegs) GetLegSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegIssuer gets LegIssuer, Tag 617
func (m NoLegs) GetLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedLegIssuerLen gets EncodedLegIssuerLen, Tag 618
func (m NoLegs) GetEncodedLegIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedLegIssuer gets EncodedLegIssuer, Tag 619
func (m NoLegs) GetEncodedLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecurityDesc gets LegSecurityDesc, Tag 620
func (m NoLegs) GetLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedLegSecurityDescLen gets EncodedLegSecurityDescLen, Tag 621
func (m NoLegs) GetEncodedLegSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedLegSecurityDesc gets EncodedLegSecurityDesc, Tag 622
func (m NoLegs) GetEncodedLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegRatioQty gets LegRatioQty, Tag 623
func (m NoLegs) GetLegRatioQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegRatioQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegSide gets LegSide, Tag 624
func (m NoLegs) GetLegSide() (v string, err quickfix.MessageRejectError) {
	var f field.LegSideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegCurrency gets LegCurrency, Tag 556
func (m NoLegs) GetLegCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegPool gets LegPool, Tag 740
func (m NoLegs) GetLegPool() (v string, err quickfix.MessageRejectError) {
	var f field.LegPoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegDatedDate gets LegDatedDate, Tag 739
func (m NoLegs) GetLegDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegDatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegContractSettlMonth gets LegContractSettlMonth, Tag 955
func (m NoLegs) GetLegContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.LegContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegInterestAccrualDate gets LegInterestAccrualDate, Tag 956
func (m NoLegs) GetLegInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegInterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegUnitOfMeasure gets LegUnitOfMeasure, Tag 999
func (m NoLegs) GetLegUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.LegUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegTimeUnit gets LegTimeUnit, Tag 1001
func (m NoLegs) GetLegTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.LegTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegOptionRatio gets LegOptionRatio, Tag 1017
func (m NoLegs) GetLegOptionRatio() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegOptionRatioField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegPrice gets LegPrice, Tag 566
func (m NoLegs) GetLegPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegMaturityTime gets LegMaturityTime, Tag 1212
func (m NoLegs) GetLegMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegPutOrCall gets LegPutOrCall, Tag 1358
func (m NoLegs) GetLegPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.LegPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegExerciseStyle gets LegExerciseStyle, Tag 1420
func (m NoLegs) GetLegExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.LegExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegUnitOfMeasureQty gets LegUnitOfMeasureQty, Tag 1224
func (m NoLegs) GetLegUnitOfMeasureQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegPriceUnitOfMeasure gets LegPriceUnitOfMeasure, Tag 1421
func (m NoLegs) GetLegPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.LegPriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegPriceUnitOfMeasureQty gets LegPriceUnitOfMeasureQty, Tag 1422
func (m NoLegs) GetLegPriceUnitOfMeasureQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegPriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegQty gets LegQty, Tag 687
func (m NoLegs) GetLegQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegSwapType gets LegSwapType, Tag 690
func (m NoLegs) GetLegSwapType() (v enum.LegSwapType, err quickfix.MessageRejectError) {
	var f field.LegSwapTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoLegStipulations gets NoLegStipulations, Tag 683
func (m NoLegs) GetNoLegStipulations() (NoLegStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegPositionEffect gets LegPositionEffect, Tag 564
func (m NoLegs) GetLegPositionEffect() (v string, err quickfix.MessageRejectError) {
	var f field.LegPositionEffectField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegCoveredOrUncovered gets LegCoveredOrUncovered, Tag 565
func (m NoLegs) GetLegCoveredOrUncovered() (v int, err quickfix.MessageRejectError) {
	var f field.LegCoveredOrUncoveredField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegRefID gets LegRefID, Tag 654
func (m NoLegs) GetLegRefID() (v string, err quickfix.MessageRejectError) {
	var f field.LegRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSettlType gets LegSettlType, Tag 587
func (m NoLegs) GetLegSettlType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSettlDate gets LegSettlDate, Tag 588
func (m NoLegs) GetLegSettlDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegSettlDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegLastPx gets LegLastPx, Tag 637
func (m NoLegs) GetLegLastPx() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegLastPxField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegOrderQty gets LegOrderQty, Tag 685
func (m NoLegs) GetLegOrderQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegOrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegSettlCurrency gets LegSettlCurrency, Tag 675
func (m NoLegs) GetLegSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegSettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegLastForwardPoints gets LegLastForwardPoints, Tag 1073
func (m NoLegs) GetLegLastForwardPoints() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegLastForwardPointsField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegCalculatedCcyLastQty gets LegCalculatedCcyLastQty, Tag 1074
func (m NoLegs) GetLegCalculatedCcyLastQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegCalculatedCcyLastQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegGrossTradeAmt gets LegGrossTradeAmt, Tag 1075
func (m NoLegs) GetLegGrossTradeAmt() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegGrossTradeAmtField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetNoNested3PartyIDs gets NoNested3PartyIDs, Tag 948
func (m NoLegs) GetNoNested3PartyIDs() (NoNested3PartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested3PartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegAllocID gets LegAllocID, Tag 1366
func (m NoLegs) GetLegAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.LegAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoLegAllocs gets NoLegAllocs, Tag 670
func (m NoLegs) GetNoLegAllocs() (NoLegAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegVolatility gets LegVolatility, Tag 1379
func (m NoLegs) GetLegVolatility() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegVolatilityField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegDividendYield gets LegDividendYield, Tag 1381
func (m NoLegs) GetLegDividendYield() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegDividendYieldField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegCurrencyRatio gets LegCurrencyRatio, Tag 1383
func (m NoLegs) GetLegCurrencyRatio() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegCurrencyRatioField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegExecInst gets LegExecInst, Tag 1384
func (m NoLegs) GetLegExecInst() (v string, err quickfix.MessageRejectError) {
	var f field.LegExecInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegLastQty gets LegLastQty, Tag 1418
func (m NoLegs) GetLegLastQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegLastQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//HasLegSymbol returns true if LegSymbol is present, Tag 600
func (m NoLegs) HasLegSymbol() bool {
	return m.Has(tag.LegSymbol)
}

//HasLegSymbolSfx returns true if LegSymbolSfx is present, Tag 601
func (m NoLegs) HasLegSymbolSfx() bool {
	return m.Has(tag.LegSymbolSfx)
}

//HasLegSecurityID returns true if LegSecurityID is present, Tag 602
func (m NoLegs) HasLegSecurityID() bool {
	return m.Has(tag.LegSecurityID)
}

//HasLegSecurityIDSource returns true if LegSecurityIDSource is present, Tag 603
func (m NoLegs) HasLegSecurityIDSource() bool {
	return m.Has(tag.LegSecurityIDSource)
}

//HasNoLegSecurityAltID returns true if NoLegSecurityAltID is present, Tag 604
func (m NoLegs) HasNoLegSecurityAltID() bool {
	return m.Has(tag.NoLegSecurityAltID)
}

//HasLegProduct returns true if LegProduct is present, Tag 607
func (m NoLegs) HasLegProduct() bool {
	return m.Has(tag.LegProduct)
}

//HasLegCFICode returns true if LegCFICode is present, Tag 608
func (m NoLegs) HasLegCFICode() bool {
	return m.Has(tag.LegCFICode)
}

//HasLegSecurityType returns true if LegSecurityType is present, Tag 609
func (m NoLegs) HasLegSecurityType() bool {
	return m.Has(tag.LegSecurityType)
}

//HasLegSecuritySubType returns true if LegSecuritySubType is present, Tag 764
func (m NoLegs) HasLegSecuritySubType() bool {
	return m.Has(tag.LegSecuritySubType)
}

//HasLegMaturityMonthYear returns true if LegMaturityMonthYear is present, Tag 610
func (m NoLegs) HasLegMaturityMonthYear() bool {
	return m.Has(tag.LegMaturityMonthYear)
}

//HasLegMaturityDate returns true if LegMaturityDate is present, Tag 611
func (m NoLegs) HasLegMaturityDate() bool {
	return m.Has(tag.LegMaturityDate)
}

//HasLegCouponPaymentDate returns true if LegCouponPaymentDate is present, Tag 248
func (m NoLegs) HasLegCouponPaymentDate() bool {
	return m.Has(tag.LegCouponPaymentDate)
}

//HasLegIssueDate returns true if LegIssueDate is present, Tag 249
func (m NoLegs) HasLegIssueDate() bool {
	return m.Has(tag.LegIssueDate)
}

//HasLegRepoCollateralSecurityType returns true if LegRepoCollateralSecurityType is present, Tag 250
func (m NoLegs) HasLegRepoCollateralSecurityType() bool {
	return m.Has(tag.LegRepoCollateralSecurityType)
}

//HasLegRepurchaseTerm returns true if LegRepurchaseTerm is present, Tag 251
func (m NoLegs) HasLegRepurchaseTerm() bool {
	return m.Has(tag.LegRepurchaseTerm)
}

//HasLegRepurchaseRate returns true if LegRepurchaseRate is present, Tag 252
func (m NoLegs) HasLegRepurchaseRate() bool {
	return m.Has(tag.LegRepurchaseRate)
}

//HasLegFactor returns true if LegFactor is present, Tag 253
func (m NoLegs) HasLegFactor() bool {
	return m.Has(tag.LegFactor)
}

//HasLegCreditRating returns true if LegCreditRating is present, Tag 257
func (m NoLegs) HasLegCreditRating() bool {
	return m.Has(tag.LegCreditRating)
}

//HasLegInstrRegistry returns true if LegInstrRegistry is present, Tag 599
func (m NoLegs) HasLegInstrRegistry() bool {
	return m.Has(tag.LegInstrRegistry)
}

//HasLegCountryOfIssue returns true if LegCountryOfIssue is present, Tag 596
func (m NoLegs) HasLegCountryOfIssue() bool {
	return m.Has(tag.LegCountryOfIssue)
}

//HasLegStateOrProvinceOfIssue returns true if LegStateOrProvinceOfIssue is present, Tag 597
func (m NoLegs) HasLegStateOrProvinceOfIssue() bool {
	return m.Has(tag.LegStateOrProvinceOfIssue)
}

//HasLegLocaleOfIssue returns true if LegLocaleOfIssue is present, Tag 598
func (m NoLegs) HasLegLocaleOfIssue() bool {
	return m.Has(tag.LegLocaleOfIssue)
}

//HasLegRedemptionDate returns true if LegRedemptionDate is present, Tag 254
func (m NoLegs) HasLegRedemptionDate() bool {
	return m.Has(tag.LegRedemptionDate)
}

//HasLegStrikePrice returns true if LegStrikePrice is present, Tag 612
func (m NoLegs) HasLegStrikePrice() bool {
	return m.Has(tag.LegStrikePrice)
}

//HasLegStrikeCurrency returns true if LegStrikeCurrency is present, Tag 942
func (m NoLegs) HasLegStrikeCurrency() bool {
	return m.Has(tag.LegStrikeCurrency)
}

//HasLegOptAttribute returns true if LegOptAttribute is present, Tag 613
func (m NoLegs) HasLegOptAttribute() bool {
	return m.Has(tag.LegOptAttribute)
}

//HasLegContractMultiplier returns true if LegContractMultiplier is present, Tag 614
func (m NoLegs) HasLegContractMultiplier() bool {
	return m.Has(tag.LegContractMultiplier)
}

//HasLegCouponRate returns true if LegCouponRate is present, Tag 615
func (m NoLegs) HasLegCouponRate() bool {
	return m.Has(tag.LegCouponRate)
}

//HasLegSecurityExchange returns true if LegSecurityExchange is present, Tag 616
func (m NoLegs) HasLegSecurityExchange() bool {
	return m.Has(tag.LegSecurityExchange)
}

//HasLegIssuer returns true if LegIssuer is present, Tag 617
func (m NoLegs) HasLegIssuer() bool {
	return m.Has(tag.LegIssuer)
}

//HasEncodedLegIssuerLen returns true if EncodedLegIssuerLen is present, Tag 618
func (m NoLegs) HasEncodedLegIssuerLen() bool {
	return m.Has(tag.EncodedLegIssuerLen)
}

//HasEncodedLegIssuer returns true if EncodedLegIssuer is present, Tag 619
func (m NoLegs) HasEncodedLegIssuer() bool {
	return m.Has(tag.EncodedLegIssuer)
}

//HasLegSecurityDesc returns true if LegSecurityDesc is present, Tag 620
func (m NoLegs) HasLegSecurityDesc() bool {
	return m.Has(tag.LegSecurityDesc)
}

//HasEncodedLegSecurityDescLen returns true if EncodedLegSecurityDescLen is present, Tag 621
func (m NoLegs) HasEncodedLegSecurityDescLen() bool {
	return m.Has(tag.EncodedLegSecurityDescLen)
}

//HasEncodedLegSecurityDesc returns true if EncodedLegSecurityDesc is present, Tag 622
func (m NoLegs) HasEncodedLegSecurityDesc() bool {
	return m.Has(tag.EncodedLegSecurityDesc)
}

//HasLegRatioQty returns true if LegRatioQty is present, Tag 623
func (m NoLegs) HasLegRatioQty() bool {
	return m.Has(tag.LegRatioQty)
}

//HasLegSide returns true if LegSide is present, Tag 624
func (m NoLegs) HasLegSide() bool {
	return m.Has(tag.LegSide)
}

//HasLegCurrency returns true if LegCurrency is present, Tag 556
func (m NoLegs) HasLegCurrency() bool {
	return m.Has(tag.LegCurrency)
}

//HasLegPool returns true if LegPool is present, Tag 740
func (m NoLegs) HasLegPool() bool {
	return m.Has(tag.LegPool)
}

//HasLegDatedDate returns true if LegDatedDate is present, Tag 739
func (m NoLegs) HasLegDatedDate() bool {
	return m.Has(tag.LegDatedDate)
}

//HasLegContractSettlMonth returns true if LegContractSettlMonth is present, Tag 955
func (m NoLegs) HasLegContractSettlMonth() bool {
	return m.Has(tag.LegContractSettlMonth)
}

//HasLegInterestAccrualDate returns true if LegInterestAccrualDate is present, Tag 956
func (m NoLegs) HasLegInterestAccrualDate() bool {
	return m.Has(tag.LegInterestAccrualDate)
}

//HasLegUnitOfMeasure returns true if LegUnitOfMeasure is present, Tag 999
func (m NoLegs) HasLegUnitOfMeasure() bool {
	return m.Has(tag.LegUnitOfMeasure)
}

//HasLegTimeUnit returns true if LegTimeUnit is present, Tag 1001
func (m NoLegs) HasLegTimeUnit() bool {
	return m.Has(tag.LegTimeUnit)
}

//HasLegOptionRatio returns true if LegOptionRatio is present, Tag 1017
func (m NoLegs) HasLegOptionRatio() bool {
	return m.Has(tag.LegOptionRatio)
}

//HasLegPrice returns true if LegPrice is present, Tag 566
func (m NoLegs) HasLegPrice() bool {
	return m.Has(tag.LegPrice)
}

//HasLegMaturityTime returns true if LegMaturityTime is present, Tag 1212
func (m NoLegs) HasLegMaturityTime() bool {
	return m.Has(tag.LegMaturityTime)
}

//HasLegPutOrCall returns true if LegPutOrCall is present, Tag 1358
func (m NoLegs) HasLegPutOrCall() bool {
	return m.Has(tag.LegPutOrCall)
}

//HasLegExerciseStyle returns true if LegExerciseStyle is present, Tag 1420
func (m NoLegs) HasLegExerciseStyle() bool {
	return m.Has(tag.LegExerciseStyle)
}

//HasLegUnitOfMeasureQty returns true if LegUnitOfMeasureQty is present, Tag 1224
func (m NoLegs) HasLegUnitOfMeasureQty() bool {
	return m.Has(tag.LegUnitOfMeasureQty)
}

//HasLegPriceUnitOfMeasure returns true if LegPriceUnitOfMeasure is present, Tag 1421
func (m NoLegs) HasLegPriceUnitOfMeasure() bool {
	return m.Has(tag.LegPriceUnitOfMeasure)
}

//HasLegPriceUnitOfMeasureQty returns true if LegPriceUnitOfMeasureQty is present, Tag 1422
func (m NoLegs) HasLegPriceUnitOfMeasureQty() bool {
	return m.Has(tag.LegPriceUnitOfMeasureQty)
}

//HasLegQty returns true if LegQty is present, Tag 687
func (m NoLegs) HasLegQty() bool {
	return m.Has(tag.LegQty)
}

//HasLegSwapType returns true if LegSwapType is present, Tag 690
func (m NoLegs) HasLegSwapType() bool {
	return m.Has(tag.LegSwapType)
}

//HasNoLegStipulations returns true if NoLegStipulations is present, Tag 683
func (m NoLegs) HasNoLegStipulations() bool {
	return m.Has(tag.NoLegStipulations)
}

//HasLegPositionEffect returns true if LegPositionEffect is present, Tag 564
func (m NoLegs) HasLegPositionEffect() bool {
	return m.Has(tag.LegPositionEffect)
}

//HasLegCoveredOrUncovered returns true if LegCoveredOrUncovered is present, Tag 565
func (m NoLegs) HasLegCoveredOrUncovered() bool {
	return m.Has(tag.LegCoveredOrUncovered)
}

//HasLegRefID returns true if LegRefID is present, Tag 654
func (m NoLegs) HasLegRefID() bool {
	return m.Has(tag.LegRefID)
}

//HasLegSettlType returns true if LegSettlType is present, Tag 587
func (m NoLegs) HasLegSettlType() bool {
	return m.Has(tag.LegSettlType)
}

//HasLegSettlDate returns true if LegSettlDate is present, Tag 588
func (m NoLegs) HasLegSettlDate() bool {
	return m.Has(tag.LegSettlDate)
}

//HasLegLastPx returns true if LegLastPx is present, Tag 637
func (m NoLegs) HasLegLastPx() bool {
	return m.Has(tag.LegLastPx)
}

//HasLegOrderQty returns true if LegOrderQty is present, Tag 685
func (m NoLegs) HasLegOrderQty() bool {
	return m.Has(tag.LegOrderQty)
}

//HasLegSettlCurrency returns true if LegSettlCurrency is present, Tag 675
func (m NoLegs) HasLegSettlCurrency() bool {
	return m.Has(tag.LegSettlCurrency)
}

//HasLegLastForwardPoints returns true if LegLastForwardPoints is present, Tag 1073
func (m NoLegs) HasLegLastForwardPoints() bool {
	return m.Has(tag.LegLastForwardPoints)
}

//HasLegCalculatedCcyLastQty returns true if LegCalculatedCcyLastQty is present, Tag 1074
func (m NoLegs) HasLegCalculatedCcyLastQty() bool {
	return m.Has(tag.LegCalculatedCcyLastQty)
}

//HasLegGrossTradeAmt returns true if LegGrossTradeAmt is present, Tag 1075
func (m NoLegs) HasLegGrossTradeAmt() bool {
	return m.Has(tag.LegGrossTradeAmt)
}

//HasNoNested3PartyIDs returns true if NoNested3PartyIDs is present, Tag 948
func (m NoLegs) HasNoNested3PartyIDs() bool {
	return m.Has(tag.NoNested3PartyIDs)
}

//HasLegAllocID returns true if LegAllocID is present, Tag 1366
func (m NoLegs) HasLegAllocID() bool {
	return m.Has(tag.LegAllocID)
}

//HasNoLegAllocs returns true if NoLegAllocs is present, Tag 670
func (m NoLegs) HasNoLegAllocs() bool {
	return m.Has(tag.NoLegAllocs)
}

//HasLegVolatility returns true if LegVolatility is present, Tag 1379
func (m NoLegs) HasLegVolatility() bool {
	return m.Has(tag.LegVolatility)
}

//HasLegDividendYield returns true if LegDividendYield is present, Tag 1381
func (m NoLegs) HasLegDividendYield() bool {
	return m.Has(tag.LegDividendYield)
}

//HasLegCurrencyRatio returns true if LegCurrencyRatio is present, Tag 1383
func (m NoLegs) HasLegCurrencyRatio() bool {
	return m.Has(tag.LegCurrencyRatio)
}

//HasLegExecInst returns true if LegExecInst is present, Tag 1384
func (m NoLegs) HasLegExecInst() bool {
	return m.Has(tag.LegExecInst)
}

//HasLegLastQty returns true if LegLastQty is present, Tag 1418
func (m NoLegs) HasLegLastQty() bool {
	return m.Has(tag.LegLastQty)
}

//NoLegSecurityAltID is a repeating group element, Tag 604
type NoLegSecurityAltID struct {
	quickfix.Group
}

//SetLegSecurityAltID sets LegSecurityAltID, Tag 605
func (m NoLegSecurityAltID) SetLegSecurityAltID(v string) {
	m.Set(field.NewLegSecurityAltID(v))
}

//SetLegSecurityAltIDSource sets LegSecurityAltIDSource, Tag 606
func (m NoLegSecurityAltID) SetLegSecurityAltIDSource(v string) {
	m.Set(field.NewLegSecurityAltIDSource(v))
}

//GetLegSecurityAltID gets LegSecurityAltID, Tag 605
func (m NoLegSecurityAltID) GetLegSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegSecurityAltIDSource gets LegSecurityAltIDSource, Tag 606
func (m NoLegSecurityAltID) GetLegSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasLegSecurityAltID returns true if LegSecurityAltID is present, Tag 605
func (m NoLegSecurityAltID) HasLegSecurityAltID() bool {
	return m.Has(tag.LegSecurityAltID)
}

//HasLegSecurityAltIDSource returns true if LegSecurityAltIDSource is present, Tag 606
func (m NoLegSecurityAltID) HasLegSecurityAltIDSource() bool {
	return m.Has(tag.LegSecurityAltIDSource)
}

//NoLegSecurityAltIDRepeatingGroup is a repeating group, Tag 604
type NoLegSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegSecurityAltIDRepeatingGroup returns an initialized, NoLegSecurityAltIDRepeatingGroup
func NewNoLegSecurityAltIDRepeatingGroup() NoLegSecurityAltIDRepeatingGroup {
	return NoLegSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSecurityAltID), quickfix.GroupElement(tag.LegSecurityAltIDSource)})}
}

//Add create and append a new NoLegSecurityAltID to this group
func (m NoLegSecurityAltIDRepeatingGroup) Add() NoLegSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoLegSecurityAltID{g}
}

//Get returns the ith NoLegSecurityAltID in the NoLegSecurityAltIDRepeatinGroup
func (m NoLegSecurityAltIDRepeatingGroup) Get(i int) NoLegSecurityAltID {
	return NoLegSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoLegStipulations is a repeating group element, Tag 683
type NoLegStipulations struct {
	quickfix.Group
}

//SetLegStipulationType sets LegStipulationType, Tag 688
func (m NoLegStipulations) SetLegStipulationType(v string) {
	m.Set(field.NewLegStipulationType(v))
}

//SetLegStipulationValue sets LegStipulationValue, Tag 689
func (m NoLegStipulations) SetLegStipulationValue(v string) {
	m.Set(field.NewLegStipulationValue(v))
}

//GetLegStipulationType gets LegStipulationType, Tag 688
func (m NoLegStipulations) GetLegStipulationType() (v string, err quickfix.MessageRejectError) {
	var f field.LegStipulationTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegStipulationValue gets LegStipulationValue, Tag 689
func (m NoLegStipulations) GetLegStipulationValue() (v string, err quickfix.MessageRejectError) {
	var f field.LegStipulationValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasLegStipulationType returns true if LegStipulationType is present, Tag 688
func (m NoLegStipulations) HasLegStipulationType() bool {
	return m.Has(tag.LegStipulationType)
}

//HasLegStipulationValue returns true if LegStipulationValue is present, Tag 689
func (m NoLegStipulations) HasLegStipulationValue() bool {
	return m.Has(tag.LegStipulationValue)
}

//NoLegStipulationsRepeatingGroup is a repeating group, Tag 683
type NoLegStipulationsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegStipulationsRepeatingGroup returns an initialized, NoLegStipulationsRepeatingGroup
func NewNoLegStipulationsRepeatingGroup() NoLegStipulationsRepeatingGroup {
	return NoLegStipulationsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegStipulations,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegStipulationType), quickfix.GroupElement(tag.LegStipulationValue)})}
}

//Add create and append a new NoLegStipulations to this group
func (m NoLegStipulationsRepeatingGroup) Add() NoLegStipulations {
	g := m.RepeatingGroup.Add()
	return NoLegStipulations{g}
}

//Get returns the ith NoLegStipulations in the NoLegStipulationsRepeatinGroup
func (m NoLegStipulationsRepeatingGroup) Get(i int) NoLegStipulations {
	return NoLegStipulations{m.RepeatingGroup.Get(i)}
}

//NoNested3PartyIDs is a repeating group element, Tag 948
type NoNested3PartyIDs struct {
	quickfix.Group
}

//SetNested3PartyID sets Nested3PartyID, Tag 949
func (m NoNested3PartyIDs) SetNested3PartyID(v string) {
	m.Set(field.NewNested3PartyID(v))
}

//SetNested3PartyIDSource sets Nested3PartyIDSource, Tag 950
func (m NoNested3PartyIDs) SetNested3PartyIDSource(v string) {
	m.Set(field.NewNested3PartyIDSource(v))
}

//SetNested3PartyRole sets Nested3PartyRole, Tag 951
func (m NoNested3PartyIDs) SetNested3PartyRole(v int) {
	m.Set(field.NewNested3PartyRole(v))
}

//SetNoNested3PartySubIDs sets NoNested3PartySubIDs, Tag 952
func (m NoNested3PartyIDs) SetNoNested3PartySubIDs(f NoNested3PartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNested3PartyID gets Nested3PartyID, Tag 949
func (m NoNested3PartyIDs) GetNested3PartyID() (v string, err quickfix.MessageRejectError) {
	var f field.Nested3PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNested3PartyIDSource gets Nested3PartyIDSource, Tag 950
func (m NoNested3PartyIDs) GetNested3PartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.Nested3PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNested3PartyRole gets Nested3PartyRole, Tag 951
func (m NoNested3PartyIDs) GetNested3PartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.Nested3PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoNested3PartySubIDs gets NoNested3PartySubIDs, Tag 952
func (m NoNested3PartyIDs) GetNoNested3PartySubIDs() (NoNested3PartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested3PartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNested3PartyID returns true if Nested3PartyID is present, Tag 949
func (m NoNested3PartyIDs) HasNested3PartyID() bool {
	return m.Has(tag.Nested3PartyID)
}

//HasNested3PartyIDSource returns true if Nested3PartyIDSource is present, Tag 950
func (m NoNested3PartyIDs) HasNested3PartyIDSource() bool {
	return m.Has(tag.Nested3PartyIDSource)
}

//HasNested3PartyRole returns true if Nested3PartyRole is present, Tag 951
func (m NoNested3PartyIDs) HasNested3PartyRole() bool {
	return m.Has(tag.Nested3PartyRole)
}

//HasNoNested3PartySubIDs returns true if NoNested3PartySubIDs is present, Tag 952
func (m NoNested3PartyIDs) HasNoNested3PartySubIDs() bool {
	return m.Has(tag.NoNested3PartySubIDs)
}

//NoNested3PartySubIDs is a repeating group element, Tag 952
type NoNested3PartySubIDs struct {
	quickfix.Group
}

//SetNested3PartySubID sets Nested3PartySubID, Tag 953
func (m NoNested3PartySubIDs) SetNested3PartySubID(v string) {
	m.Set(field.NewNested3PartySubID(v))
}

//SetNested3PartySubIDType sets Nested3PartySubIDType, Tag 954
func (m NoNested3PartySubIDs) SetNested3PartySubIDType(v int) {
	m.Set(field.NewNested3PartySubIDType(v))
}

//GetNested3PartySubID gets Nested3PartySubID, Tag 953
func (m NoNested3PartySubIDs) GetNested3PartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.Nested3PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNested3PartySubIDType gets Nested3PartySubIDType, Tag 954
func (m NoNested3PartySubIDs) GetNested3PartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.Nested3PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasNested3PartySubID returns true if Nested3PartySubID is present, Tag 953
func (m NoNested3PartySubIDs) HasNested3PartySubID() bool {
	return m.Has(tag.Nested3PartySubID)
}

//HasNested3PartySubIDType returns true if Nested3PartySubIDType is present, Tag 954
func (m NoNested3PartySubIDs) HasNested3PartySubIDType() bool {
	return m.Has(tag.Nested3PartySubIDType)
}

//NoNested3PartySubIDsRepeatingGroup is a repeating group, Tag 952
type NoNested3PartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested3PartySubIDsRepeatingGroup returns an initialized, NoNested3PartySubIDsRepeatingGroup
func NewNoNested3PartySubIDsRepeatingGroup() NoNested3PartySubIDsRepeatingGroup {
	return NoNested3PartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested3PartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested3PartySubID), quickfix.GroupElement(tag.Nested3PartySubIDType)})}
}

//Add create and append a new NoNested3PartySubIDs to this group
func (m NoNested3PartySubIDsRepeatingGroup) Add() NoNested3PartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNested3PartySubIDs{g}
}

//Get returns the ith NoNested3PartySubIDs in the NoNested3PartySubIDsRepeatinGroup
func (m NoNested3PartySubIDsRepeatingGroup) Get(i int) NoNested3PartySubIDs {
	return NoNested3PartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoNested3PartyIDsRepeatingGroup is a repeating group, Tag 948
type NoNested3PartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested3PartyIDsRepeatingGroup returns an initialized, NoNested3PartyIDsRepeatingGroup
func NewNoNested3PartyIDsRepeatingGroup() NoNested3PartyIDsRepeatingGroup {
	return NoNested3PartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested3PartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested3PartyID), quickfix.GroupElement(tag.Nested3PartyIDSource), quickfix.GroupElement(tag.Nested3PartyRole), NewNoNested3PartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoNested3PartyIDs to this group
func (m NoNested3PartyIDsRepeatingGroup) Add() NoNested3PartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNested3PartyIDs{g}
}

//Get returns the ith NoNested3PartyIDs in the NoNested3PartyIDsRepeatinGroup
func (m NoNested3PartyIDsRepeatingGroup) Get(i int) NoNested3PartyIDs {
	return NoNested3PartyIDs{m.RepeatingGroup.Get(i)}
}

//NoLegAllocs is a repeating group element, Tag 670
type NoLegAllocs struct {
	quickfix.Group
}

//SetLegAllocAccount sets LegAllocAccount, Tag 671
func (m NoLegAllocs) SetLegAllocAccount(v string) {
	m.Set(field.NewLegAllocAccount(v))
}

//SetLegIndividualAllocID sets LegIndividualAllocID, Tag 672
func (m NoLegAllocs) SetLegIndividualAllocID(v string) {
	m.Set(field.NewLegIndividualAllocID(v))
}

//SetLegAllocQty sets LegAllocQty, Tag 673
func (m NoLegAllocs) SetLegAllocQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegAllocQty(value, scale))
}

//SetLegAllocAcctIDSource sets LegAllocAcctIDSource, Tag 674
func (m NoLegAllocs) SetLegAllocAcctIDSource(v string) {
	m.Set(field.NewLegAllocAcctIDSource(v))
}

//SetLegAllocSettlCurrency sets LegAllocSettlCurrency, Tag 1367
func (m NoLegAllocs) SetLegAllocSettlCurrency(v string) {
	m.Set(field.NewLegAllocSettlCurrency(v))
}

//SetNoNested2PartyIDs sets NoNested2PartyIDs, Tag 756
func (m NoLegAllocs) SetNoNested2PartyIDs(f NoNested2PartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetLegAllocAccount gets LegAllocAccount, Tag 671
func (m NoLegAllocs) GetLegAllocAccount() (v string, err quickfix.MessageRejectError) {
	var f field.LegAllocAccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegIndividualAllocID gets LegIndividualAllocID, Tag 672
func (m NoLegAllocs) GetLegIndividualAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.LegIndividualAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegAllocQty gets LegAllocQty, Tag 673
func (m NoLegAllocs) GetLegAllocQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.LegAllocQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetLegAllocAcctIDSource gets LegAllocAcctIDSource, Tag 674
func (m NoLegAllocs) GetLegAllocAcctIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegAllocAcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegAllocSettlCurrency gets LegAllocSettlCurrency, Tag 1367
func (m NoLegAllocs) GetLegAllocSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegAllocSettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoNested2PartyIDs gets NoNested2PartyIDs, Tag 756
func (m NoLegAllocs) GetNoNested2PartyIDs() (NoNested2PartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested2PartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasLegAllocAccount returns true if LegAllocAccount is present, Tag 671
func (m NoLegAllocs) HasLegAllocAccount() bool {
	return m.Has(tag.LegAllocAccount)
}

//HasLegIndividualAllocID returns true if LegIndividualAllocID is present, Tag 672
func (m NoLegAllocs) HasLegIndividualAllocID() bool {
	return m.Has(tag.LegIndividualAllocID)
}

//HasLegAllocQty returns true if LegAllocQty is present, Tag 673
func (m NoLegAllocs) HasLegAllocQty() bool {
	return m.Has(tag.LegAllocQty)
}

//HasLegAllocAcctIDSource returns true if LegAllocAcctIDSource is present, Tag 674
func (m NoLegAllocs) HasLegAllocAcctIDSource() bool {
	return m.Has(tag.LegAllocAcctIDSource)
}

//HasLegAllocSettlCurrency returns true if LegAllocSettlCurrency is present, Tag 1367
func (m NoLegAllocs) HasLegAllocSettlCurrency() bool {
	return m.Has(tag.LegAllocSettlCurrency)
}

//HasNoNested2PartyIDs returns true if NoNested2PartyIDs is present, Tag 756
func (m NoLegAllocs) HasNoNested2PartyIDs() bool {
	return m.Has(tag.NoNested2PartyIDs)
}

//NoNested2PartyIDs is a repeating group element, Tag 756
type NoNested2PartyIDs struct {
	quickfix.Group
}

//SetNested2PartyID sets Nested2PartyID, Tag 757
func (m NoNested2PartyIDs) SetNested2PartyID(v string) {
	m.Set(field.NewNested2PartyID(v))
}

//SetNested2PartyIDSource sets Nested2PartyIDSource, Tag 758
func (m NoNested2PartyIDs) SetNested2PartyIDSource(v string) {
	m.Set(field.NewNested2PartyIDSource(v))
}

//SetNested2PartyRole sets Nested2PartyRole, Tag 759
func (m NoNested2PartyIDs) SetNested2PartyRole(v int) {
	m.Set(field.NewNested2PartyRole(v))
}

//SetNoNested2PartySubIDs sets NoNested2PartySubIDs, Tag 806
func (m NoNested2PartyIDs) SetNoNested2PartySubIDs(f NoNested2PartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNested2PartyID gets Nested2PartyID, Tag 757
func (m NoNested2PartyIDs) GetNested2PartyID() (v string, err quickfix.MessageRejectError) {
	var f field.Nested2PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNested2PartyIDSource gets Nested2PartyIDSource, Tag 758
func (m NoNested2PartyIDs) GetNested2PartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.Nested2PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNested2PartyRole gets Nested2PartyRole, Tag 759
func (m NoNested2PartyIDs) GetNested2PartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.Nested2PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoNested2PartySubIDs gets NoNested2PartySubIDs, Tag 806
func (m NoNested2PartyIDs) GetNoNested2PartySubIDs() (NoNested2PartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested2PartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNested2PartyID returns true if Nested2PartyID is present, Tag 757
func (m NoNested2PartyIDs) HasNested2PartyID() bool {
	return m.Has(tag.Nested2PartyID)
}

//HasNested2PartyIDSource returns true if Nested2PartyIDSource is present, Tag 758
func (m NoNested2PartyIDs) HasNested2PartyIDSource() bool {
	return m.Has(tag.Nested2PartyIDSource)
}

//HasNested2PartyRole returns true if Nested2PartyRole is present, Tag 759
func (m NoNested2PartyIDs) HasNested2PartyRole() bool {
	return m.Has(tag.Nested2PartyRole)
}

//HasNoNested2PartySubIDs returns true if NoNested2PartySubIDs is present, Tag 806
func (m NoNested2PartyIDs) HasNoNested2PartySubIDs() bool {
	return m.Has(tag.NoNested2PartySubIDs)
}

//NoNested2PartySubIDs is a repeating group element, Tag 806
type NoNested2PartySubIDs struct {
	quickfix.Group
}

//SetNested2PartySubID sets Nested2PartySubID, Tag 760
func (m NoNested2PartySubIDs) SetNested2PartySubID(v string) {
	m.Set(field.NewNested2PartySubID(v))
}

//SetNested2PartySubIDType sets Nested2PartySubIDType, Tag 807
func (m NoNested2PartySubIDs) SetNested2PartySubIDType(v int) {
	m.Set(field.NewNested2PartySubIDType(v))
}

//GetNested2PartySubID gets Nested2PartySubID, Tag 760
func (m NoNested2PartySubIDs) GetNested2PartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.Nested2PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNested2PartySubIDType gets Nested2PartySubIDType, Tag 807
func (m NoNested2PartySubIDs) GetNested2PartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.Nested2PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasNested2PartySubID returns true if Nested2PartySubID is present, Tag 760
func (m NoNested2PartySubIDs) HasNested2PartySubID() bool {
	return m.Has(tag.Nested2PartySubID)
}

//HasNested2PartySubIDType returns true if Nested2PartySubIDType is present, Tag 807
func (m NoNested2PartySubIDs) HasNested2PartySubIDType() bool {
	return m.Has(tag.Nested2PartySubIDType)
}

//NoNested2PartySubIDsRepeatingGroup is a repeating group, Tag 806
type NoNested2PartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested2PartySubIDsRepeatingGroup returns an initialized, NoNested2PartySubIDsRepeatingGroup
func NewNoNested2PartySubIDsRepeatingGroup() NoNested2PartySubIDsRepeatingGroup {
	return NoNested2PartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested2PartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested2PartySubID), quickfix.GroupElement(tag.Nested2PartySubIDType)})}
}

//Add create and append a new NoNested2PartySubIDs to this group
func (m NoNested2PartySubIDsRepeatingGroup) Add() NoNested2PartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNested2PartySubIDs{g}
}

//Get returns the ith NoNested2PartySubIDs in the NoNested2PartySubIDsRepeatinGroup
func (m NoNested2PartySubIDsRepeatingGroup) Get(i int) NoNested2PartySubIDs {
	return NoNested2PartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoNested2PartyIDsRepeatingGroup is a repeating group, Tag 756
type NoNested2PartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested2PartyIDsRepeatingGroup returns an initialized, NoNested2PartyIDsRepeatingGroup
func NewNoNested2PartyIDsRepeatingGroup() NoNested2PartyIDsRepeatingGroup {
	return NoNested2PartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested2PartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested2PartyID), quickfix.GroupElement(tag.Nested2PartyIDSource), quickfix.GroupElement(tag.Nested2PartyRole), NewNoNested2PartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoNested2PartyIDs to this group
func (m NoNested2PartyIDsRepeatingGroup) Add() NoNested2PartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNested2PartyIDs{g}
}

//Get returns the ith NoNested2PartyIDs in the NoNested2PartyIDsRepeatinGroup
func (m NoNested2PartyIDsRepeatingGroup) Get(i int) NoNested2PartyIDs {
	return NoNested2PartyIDs{m.RepeatingGroup.Get(i)}
}

//NoLegAllocsRepeatingGroup is a repeating group, Tag 670
type NoLegAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegAllocsRepeatingGroup returns an initialized, NoLegAllocsRepeatingGroup
func NewNoLegAllocsRepeatingGroup() NoLegAllocsRepeatingGroup {
	return NoLegAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegAllocAccount), quickfix.GroupElement(tag.LegIndividualAllocID), quickfix.GroupElement(tag.LegAllocQty), quickfix.GroupElement(tag.LegAllocAcctIDSource), quickfix.GroupElement(tag.LegAllocSettlCurrency), NewNoNested2PartyIDsRepeatingGroup()})}
}

//Add create and append a new NoLegAllocs to this group
func (m NoLegAllocsRepeatingGroup) Add() NoLegAllocs {
	g := m.RepeatingGroup.Add()
	return NoLegAllocs{g}
}

//Get returns the ith NoLegAllocs in the NoLegAllocsRepeatinGroup
func (m NoLegAllocsRepeatingGroup) Get(i int) NoLegAllocs {
	return NoLegAllocs{m.RepeatingGroup.Get(i)}
}

//NoLegsRepeatingGroup is a repeating group, Tag 555
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), NewNoLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit), quickfix.GroupElement(tag.LegOptionRatio), quickfix.GroupElement(tag.LegPrice), quickfix.GroupElement(tag.LegMaturityTime), quickfix.GroupElement(tag.LegPutOrCall), quickfix.GroupElement(tag.LegExerciseStyle), quickfix.GroupElement(tag.LegUnitOfMeasureQty), quickfix.GroupElement(tag.LegPriceUnitOfMeasure), quickfix.GroupElement(tag.LegPriceUnitOfMeasureQty), quickfix.GroupElement(tag.LegQty), quickfix.GroupElement(tag.LegSwapType), NewNoLegStipulationsRepeatingGroup(), quickfix.GroupElement(tag.LegPositionEffect), quickfix.GroupElement(tag.LegCoveredOrUncovered), quickfix.GroupElement(tag.LegRefID), quickfix.GroupElement(tag.LegSettlType), quickfix.GroupElement(tag.LegSettlDate), quickfix.GroupElement(tag.LegLastPx), quickfix.GroupElement(tag.LegOrderQty), quickfix.GroupElement(tag.LegSettlCurrency), quickfix.GroupElement(tag.LegLastForwardPoints), quickfix.GroupElement(tag.LegCalculatedCcyLastQty), quickfix.GroupElement(tag.LegGrossTradeAmt), NewNoNested3PartyIDsRepeatingGroup(), quickfix.GroupElement(tag.LegAllocID), NewNoLegAllocsRepeatingGroup(), quickfix.GroupElement(tag.LegVolatility), quickfix.GroupElement(tag.LegDividendYield), quickfix.GroupElement(tag.LegCurrencyRatio), quickfix.GroupElement(tag.LegExecInst), quickfix.GroupElement(tag.LegLastQty)})}
}

//Add create and append a new NoLegs to this group
func (m NoLegsRepeatingGroup) Add() NoLegs {
	g := m.RepeatingGroup.Add()
	return NoLegs{g}
}

//Get returns the ith NoLegs in the NoLegsRepeatinGroup
func (m NoLegsRepeatingGroup) Get(i int) NoLegs {
	return NoLegs{m.RepeatingGroup.Get(i)}
}

//NoUnderlyings is a repeating group element, Tag 711
type NoUnderlyings struct {
	quickfix.Group
}

//SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

//SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

//SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

//SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

//SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingProduct sets UnderlyingProduct, Tag 462
func (m NoUnderlyings) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

//SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

//SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

//SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

//SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

//SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

//SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

//SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

//SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

//SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

//SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) SetUnderlyingRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingRepurchaseRate(value, scale))
}

//SetUnderlyingFactor sets UnderlyingFactor, Tag 246
func (m NoUnderlyings) SetUnderlyingFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFactor(value, scale))
}

//SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

//SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

//SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

//SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

//SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

//SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

//SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

//SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

//SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

//SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

//SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

//SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

//SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

//SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

//SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

//SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

//SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

//SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

//SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

//SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

//SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

//SetUnderlyingQty sets UnderlyingQty, Tag 879
func (m NoUnderlyings) SetUnderlyingQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingQty(value, scale))
}

//SetUnderlyingPx sets UnderlyingPx, Tag 810
func (m NoUnderlyings) SetUnderlyingPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPx(value, scale))
}

//SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) SetUnderlyingDirtyPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDirtyPrice(value, scale))
}

//SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) SetUnderlyingEndPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndPrice(value, scale))
}

//SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) SetUnderlyingStartValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStartValue(value, scale))
}

//SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) SetUnderlyingCurrentValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCurrentValue(value, scale))
}

//SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) SetUnderlyingEndValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndValue(value, scale))
}

//SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) SetUnderlyingAllocationPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAllocationPercent(value, scale))
}

//SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) SetUnderlyingSettlementType(v enum.UnderlyingSettlementType) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

//SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) SetUnderlyingCashAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCashAmount(value, scale))
}

//SetUnderlyingCashType sets UnderlyingCashType, Tag 974
func (m NoUnderlyings) SetUnderlyingCashType(v enum.UnderlyingCashType) {
	m.Set(field.NewUnderlyingCashType(v))
}

//SetUnderlyingUnitOfMeasure sets UnderlyingUnitOfMeasure, Tag 998
func (m NoUnderlyings) SetUnderlyingUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingUnitOfMeasure(v))
}

//SetUnderlyingTimeUnit sets UnderlyingTimeUnit, Tag 1000
func (m NoUnderlyings) SetUnderlyingTimeUnit(v string) {
	m.Set(field.NewUnderlyingTimeUnit(v))
}

//SetUnderlyingCapValue sets UnderlyingCapValue, Tag 1038
func (m NoUnderlyings) SetUnderlyingCapValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCapValue(value, scale))
}

//SetNoUndlyInstrumentParties sets NoUndlyInstrumentParties, Tag 1058
func (m NoUnderlyings) SetNoUndlyInstrumentParties(f NoUndlyInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingSettlMethod sets UnderlyingSettlMethod, Tag 1039
func (m NoUnderlyings) SetUnderlyingSettlMethod(v string) {
	m.Set(field.NewUnderlyingSettlMethod(v))
}

//SetUnderlyingAdjustedQuantity sets UnderlyingAdjustedQuantity, Tag 1044
func (m NoUnderlyings) SetUnderlyingAdjustedQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAdjustedQuantity(value, scale))
}

//SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) SetUnderlyingFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFXRate(value, scale))
}

//SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) SetUnderlyingFXRateCalc(v enum.UnderlyingFXRateCalc) {
	m.Set(field.NewUnderlyingFXRateCalc(v))
}

//SetUnderlyingMaturityTime sets UnderlyingMaturityTime, Tag 1213
func (m NoUnderlyings) SetUnderlyingMaturityTime(v string) {
	m.Set(field.NewUnderlyingMaturityTime(v))
}

//SetUnderlyingPutOrCall sets UnderlyingPutOrCall, Tag 315
func (m NoUnderlyings) SetUnderlyingPutOrCall(v int) {
	m.Set(field.NewUnderlyingPutOrCall(v))
}

//SetUnderlyingExerciseStyle sets UnderlyingExerciseStyle, Tag 1419
func (m NoUnderlyings) SetUnderlyingExerciseStyle(v int) {
	m.Set(field.NewUnderlyingExerciseStyle(v))
}

//SetUnderlyingUnitOfMeasureQty sets UnderlyingUnitOfMeasureQty, Tag 1423
func (m NoUnderlyings) SetUnderlyingUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingUnitOfMeasureQty(value, scale))
}

//SetUnderlyingPriceUnitOfMeasure sets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasure(v))
}

//SetUnderlyingPriceUnitOfMeasureQty sets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasureQty(value, scale))
}

//GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) GetUnderlyingSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) GetUnderlyingSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) GetUnderlyingSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) GetUnderlyingSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingProduct gets UnderlyingProduct, Tag 462
func (m NoUnderlyings) GetUnderlyingProduct() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) GetUnderlyingCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) GetUnderlyingSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) GetUnderlyingSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) GetUnderlyingMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) GetUnderlyingMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) GetUnderlyingCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) GetUnderlyingIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) GetUnderlyingRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) GetUnderlyingRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) GetUnderlyingRepurchaseRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingFactor gets UnderlyingFactor, Tag 246
func (m NoUnderlyings) GetUnderlyingFactor() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingFactorField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) GetUnderlyingCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) GetUnderlyingInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) GetUnderlyingCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) GetUnderlyingStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) GetUnderlyingLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) GetUnderlyingRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) GetUnderlyingStrikePrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) GetUnderlyingStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) GetUnderlyingOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) GetUnderlyingContractMultiplier() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) GetUnderlyingCouponRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) GetUnderlyingSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) GetUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) GetEncodedUnderlyingIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) GetEncodedUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) GetUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) GetUnderlyingCPProgram() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) GetUnderlyingCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) GetUnderlyingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingQty gets UnderlyingQty, Tag 879
func (m NoUnderlyings) GetUnderlyingQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingPx gets UnderlyingPx, Tag 810
func (m NoUnderlyings) GetUnderlyingPx() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingPxField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) GetUnderlyingDirtyPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingDirtyPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) GetUnderlyingEndPrice() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndPriceField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) GetUnderlyingStartValue() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingStartValueField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) GetUnderlyingCurrentValue() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrentValueField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) GetUnderlyingEndValue() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndValueField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) GetUnderlyingAllocationPercent() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingAllocationPercentField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) GetUnderlyingSettlementType() (v enum.UnderlyingSettlementType, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlementTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) GetUnderlyingCashAmount() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashAmountField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingCashType gets UnderlyingCashType, Tag 974
func (m NoUnderlyings) GetUnderlyingCashType() (v enum.UnderlyingCashType, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998
func (m NoUnderlyings) GetUnderlyingUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000
func (m NoUnderlyings) GetUnderlyingTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038
func (m NoUnderlyings) GetUnderlyingCapValue() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingCapValueField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058
func (m NoUnderlyings) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039
func (m NoUnderlyings) GetUnderlyingSettlMethod() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044
func (m NoUnderlyings) GetUnderlyingAdjustedQuantity() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingAdjustedQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) GetUnderlyingFXRate() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) GetUnderlyingFXRateCalc() (v enum.UnderlyingFXRateCalc, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityTime gets UnderlyingMaturityTime, Tag 1213
func (m NoUnderlyings) GetUnderlyingMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315
func (m NoUnderlyings) GetUnderlyingPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingExerciseStyle gets UnderlyingExerciseStyle, Tag 1419
func (m NoUnderlyings) GetUnderlyingExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingUnitOfMeasureQty gets UnderlyingUnitOfMeasureQty, Tag 1423
func (m NoUnderlyings) GetUnderlyingUnitOfMeasureQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetUnderlyingPriceUnitOfMeasure gets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPriceUnitOfMeasureQty gets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasureQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311
func (m NoUnderlyings) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

//HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312
func (m NoUnderlyings) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

//HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309
func (m NoUnderlyings) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

//HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305
func (m NoUnderlyings) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

//HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457
func (m NoUnderlyings) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

//HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462
func (m NoUnderlyings) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

//HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463
func (m NoUnderlyings) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

//HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310
func (m NoUnderlyings) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

//HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763
func (m NoUnderlyings) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

//HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313
func (m NoUnderlyings) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

//HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542
func (m NoUnderlyings) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

//HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241
func (m NoUnderlyings) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

//HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242
func (m NoUnderlyings) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

//HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243
func (m NoUnderlyings) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

//HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244
func (m NoUnderlyings) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

//HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245
func (m NoUnderlyings) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

//HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246
func (m NoUnderlyings) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

//HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256
func (m NoUnderlyings) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

//HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595
func (m NoUnderlyings) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

//HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592
func (m NoUnderlyings) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

//HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593
func (m NoUnderlyings) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

//HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594
func (m NoUnderlyings) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

//HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247
func (m NoUnderlyings) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

//HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316
func (m NoUnderlyings) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

//HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941
func (m NoUnderlyings) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

//HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317
func (m NoUnderlyings) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

//HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436
func (m NoUnderlyings) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

//HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435
func (m NoUnderlyings) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

//HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308
func (m NoUnderlyings) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

//HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306
func (m NoUnderlyings) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

//HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362
func (m NoUnderlyings) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

//HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363
func (m NoUnderlyings) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

//HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307
func (m NoUnderlyings) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

//HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

//HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

//HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877
func (m NoUnderlyings) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

//HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878
func (m NoUnderlyings) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

//HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318
func (m NoUnderlyings) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

//HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879
func (m NoUnderlyings) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

//HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810
func (m NoUnderlyings) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

//HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882
func (m NoUnderlyings) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

//HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883
func (m NoUnderlyings) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

//HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884
func (m NoUnderlyings) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

//HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885
func (m NoUnderlyings) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

//HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886
func (m NoUnderlyings) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

//HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887
func (m NoUnderlyings) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

//HasUnderlyingAllocationPercent returns true if UnderlyingAllocationPercent is present, Tag 972
func (m NoUnderlyings) HasUnderlyingAllocationPercent() bool {
	return m.Has(tag.UnderlyingAllocationPercent)
}

//HasUnderlyingSettlementType returns true if UnderlyingSettlementType is present, Tag 975
func (m NoUnderlyings) HasUnderlyingSettlementType() bool {
	return m.Has(tag.UnderlyingSettlementType)
}

//HasUnderlyingCashAmount returns true if UnderlyingCashAmount is present, Tag 973
func (m NoUnderlyings) HasUnderlyingCashAmount() bool {
	return m.Has(tag.UnderlyingCashAmount)
}

//HasUnderlyingCashType returns true if UnderlyingCashType is present, Tag 974
func (m NoUnderlyings) HasUnderlyingCashType() bool {
	return m.Has(tag.UnderlyingCashType)
}

//HasUnderlyingUnitOfMeasure returns true if UnderlyingUnitOfMeasure is present, Tag 998
func (m NoUnderlyings) HasUnderlyingUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingUnitOfMeasure)
}

//HasUnderlyingTimeUnit returns true if UnderlyingTimeUnit is present, Tag 1000
func (m NoUnderlyings) HasUnderlyingTimeUnit() bool {
	return m.Has(tag.UnderlyingTimeUnit)
}

//HasUnderlyingCapValue returns true if UnderlyingCapValue is present, Tag 1038
func (m NoUnderlyings) HasUnderlyingCapValue() bool {
	return m.Has(tag.UnderlyingCapValue)
}

//HasNoUndlyInstrumentParties returns true if NoUndlyInstrumentParties is present, Tag 1058
func (m NoUnderlyings) HasNoUndlyInstrumentParties() bool {
	return m.Has(tag.NoUndlyInstrumentParties)
}

//HasUnderlyingSettlMethod returns true if UnderlyingSettlMethod is present, Tag 1039
func (m NoUnderlyings) HasUnderlyingSettlMethod() bool {
	return m.Has(tag.UnderlyingSettlMethod)
}

//HasUnderlyingAdjustedQuantity returns true if UnderlyingAdjustedQuantity is present, Tag 1044
func (m NoUnderlyings) HasUnderlyingAdjustedQuantity() bool {
	return m.Has(tag.UnderlyingAdjustedQuantity)
}

//HasUnderlyingFXRate returns true if UnderlyingFXRate is present, Tag 1045
func (m NoUnderlyings) HasUnderlyingFXRate() bool {
	return m.Has(tag.UnderlyingFXRate)
}

//HasUnderlyingFXRateCalc returns true if UnderlyingFXRateCalc is present, Tag 1046
func (m NoUnderlyings) HasUnderlyingFXRateCalc() bool {
	return m.Has(tag.UnderlyingFXRateCalc)
}

//HasUnderlyingMaturityTime returns true if UnderlyingMaturityTime is present, Tag 1213
func (m NoUnderlyings) HasUnderlyingMaturityTime() bool {
	return m.Has(tag.UnderlyingMaturityTime)
}

//HasUnderlyingPutOrCall returns true if UnderlyingPutOrCall is present, Tag 315
func (m NoUnderlyings) HasUnderlyingPutOrCall() bool {
	return m.Has(tag.UnderlyingPutOrCall)
}

//HasUnderlyingExerciseStyle returns true if UnderlyingExerciseStyle is present, Tag 1419
func (m NoUnderlyings) HasUnderlyingExerciseStyle() bool {
	return m.Has(tag.UnderlyingExerciseStyle)
}

//HasUnderlyingUnitOfMeasureQty returns true if UnderlyingUnitOfMeasureQty is present, Tag 1423
func (m NoUnderlyings) HasUnderlyingUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingUnitOfMeasureQty)
}

//HasUnderlyingPriceUnitOfMeasure returns true if UnderlyingPriceUnitOfMeasure is present, Tag 1424
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasure)
}

//HasUnderlyingPriceUnitOfMeasureQty returns true if UnderlyingPriceUnitOfMeasureQty is present, Tag 1425
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasureQty)
}

//NoUnderlyingSecurityAltID is a repeating group element, Tag 457
type NoUnderlyingSecurityAltID struct {
	quickfix.Group
}

//SetUnderlyingSecurityAltID sets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltID(v string) {
	m.Set(field.NewUnderlyingSecurityAltID(v))
}

//SetUnderlyingSecurityAltIDSource sets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityAltIDSource(v))
}

//GetUnderlyingSecurityAltID gets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasUnderlyingSecurityAltID returns true if UnderlyingSecurityAltID is present, Tag 458
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltID() bool {
	return m.Has(tag.UnderlyingSecurityAltID)
}

//HasUnderlyingSecurityAltIDSource returns true if UnderlyingSecurityAltIDSource is present, Tag 459
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingSecurityAltIDSource)
}

//NoUnderlyingSecurityAltIDRepeatingGroup is a repeating group, Tag 457
type NoUnderlyingSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingSecurityAltIDRepeatingGroup
func NewNoUnderlyingSecurityAltIDRepeatingGroup() NoUnderlyingSecurityAltIDRepeatingGroup {
	return NoUnderlyingSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingSecurityAltIDSource)})}
}

//Add create and append a new NoUnderlyingSecurityAltID to this group
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Add() NoUnderlyingSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingSecurityAltID{g}
}

//Get returns the ith NoUnderlyingSecurityAltID in the NoUnderlyingSecurityAltIDRepeatinGroup
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingSecurityAltID {
	return NoUnderlyingSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingStips is a repeating group element, Tag 887
type NoUnderlyingStips struct {
	quickfix.Group
}

//SetUnderlyingStipType sets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) SetUnderlyingStipType(v string) {
	m.Set(field.NewUnderlyingStipType(v))
}

//SetUnderlyingStipValue sets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) SetUnderlyingStipValue(v string) {
	m.Set(field.NewUnderlyingStipValue(v))
}

//GetUnderlyingStipType gets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) GetUnderlyingStipType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStipValue gets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) GetUnderlyingStipValue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasUnderlyingStipType returns true if UnderlyingStipType is present, Tag 888
func (m NoUnderlyingStips) HasUnderlyingStipType() bool {
	return m.Has(tag.UnderlyingStipType)
}

//HasUnderlyingStipValue returns true if UnderlyingStipValue is present, Tag 889
func (m NoUnderlyingStips) HasUnderlyingStipValue() bool {
	return m.Has(tag.UnderlyingStipValue)
}

//NoUnderlyingStipsRepeatingGroup is a repeating group, Tag 887
type NoUnderlyingStipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingStipsRepeatingGroup returns an initialized, NoUnderlyingStipsRepeatingGroup
func NewNoUnderlyingStipsRepeatingGroup() NoUnderlyingStipsRepeatingGroup {
	return NoUnderlyingStipsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingStips,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingStipType), quickfix.GroupElement(tag.UnderlyingStipValue)})}
}

//Add create and append a new NoUnderlyingStips to this group
func (m NoUnderlyingStipsRepeatingGroup) Add() NoUnderlyingStips {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingStips{g}
}

//Get returns the ith NoUnderlyingStips in the NoUnderlyingStipsRepeatinGroup
func (m NoUnderlyingStipsRepeatingGroup) Get(i int) NoUnderlyingStips {
	return NoUnderlyingStips{m.RepeatingGroup.Get(i)}
}

//NoUndlyInstrumentParties is a repeating group element, Tag 1058
type NoUndlyInstrumentParties struct {
	quickfix.Group
}

//SetUndlyInstrumentPartyID sets UndlyInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyID(v string) {
	m.Set(field.NewUndlyInstrumentPartyID(v))
}

//SetUndlyInstrumentPartyIDSource sets UndlyInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyIDSource(v string) {
	m.Set(field.NewUndlyInstrumentPartyIDSource(v))
}

//SetUndlyInstrumentPartyRole sets UndlyInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyRole(v int) {
	m.Set(field.NewUndlyInstrumentPartyRole(v))
}

//SetNoUndlyInstrumentPartySubIDs sets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) SetNoUndlyInstrumentPartySubIDs(f NoUndlyInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetUndlyInstrumentPartyID gets UndlyInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.UndlyInstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUndlyInstrumentPartyIDSource gets UndlyInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UndlyInstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUndlyInstrumentPartyRole gets UndlyInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.UndlyInstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoUndlyInstrumentPartySubIDs gets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) GetNoUndlyInstrumentPartySubIDs() (NoUndlyInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasUndlyInstrumentPartyID returns true if UndlyInstrumentPartyID is present, Tag 1059
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyID() bool {
	return m.Has(tag.UndlyInstrumentPartyID)
}

//HasUndlyInstrumentPartyIDSource returns true if UndlyInstrumentPartyIDSource is present, Tag 1060
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyIDSource() bool {
	return m.Has(tag.UndlyInstrumentPartyIDSource)
}

//HasUndlyInstrumentPartyRole returns true if UndlyInstrumentPartyRole is present, Tag 1061
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyRole() bool {
	return m.Has(tag.UndlyInstrumentPartyRole)
}

//HasNoUndlyInstrumentPartySubIDs returns true if NoUndlyInstrumentPartySubIDs is present, Tag 1062
func (m NoUndlyInstrumentParties) HasNoUndlyInstrumentPartySubIDs() bool {
	return m.Has(tag.NoUndlyInstrumentPartySubIDs)
}

//NoUndlyInstrumentPartySubIDs is a repeating group element, Tag 1062
type NoUndlyInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetUndlyInstrumentPartySubID sets UndlyInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) SetUndlyInstrumentPartySubID(v string) {
	m.Set(field.NewUndlyInstrumentPartySubID(v))
}

//SetUndlyInstrumentPartySubIDType sets UndlyInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) SetUndlyInstrumentPartySubIDType(v int) {
	m.Set(field.NewUndlyInstrumentPartySubIDType(v))
}

//GetUndlyInstrumentPartySubID gets UndlyInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) GetUndlyInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.UndlyInstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUndlyInstrumentPartySubIDType gets UndlyInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) GetUndlyInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.UndlyInstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasUndlyInstrumentPartySubID returns true if UndlyInstrumentPartySubID is present, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) HasUndlyInstrumentPartySubID() bool {
	return m.Has(tag.UndlyInstrumentPartySubID)
}

//HasUndlyInstrumentPartySubIDType returns true if UndlyInstrumentPartySubIDType is present, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) HasUndlyInstrumentPartySubIDType() bool {
	return m.Has(tag.UndlyInstrumentPartySubIDType)
}

//NoUndlyInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1062
type NoUndlyInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUndlyInstrumentPartySubIDsRepeatingGroup returns an initialized, NoUndlyInstrumentPartySubIDsRepeatingGroup
func NewNoUndlyInstrumentPartySubIDsRepeatingGroup() NoUndlyInstrumentPartySubIDsRepeatingGroup {
	return NoUndlyInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UndlyInstrumentPartySubID), quickfix.GroupElement(tag.UndlyInstrumentPartySubIDType)})}
}

//Add create and append a new NoUndlyInstrumentPartySubIDs to this group
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Add() NoUndlyInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentPartySubIDs{g}
}

//Get returns the ith NoUndlyInstrumentPartySubIDs in the NoUndlyInstrumentPartySubIDsRepeatinGroup
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Get(i int) NoUndlyInstrumentPartySubIDs {
	return NoUndlyInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoUndlyInstrumentPartiesRepeatingGroup is a repeating group, Tag 1058
type NoUndlyInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUndlyInstrumentPartiesRepeatingGroup returns an initialized, NoUndlyInstrumentPartiesRepeatingGroup
func NewNoUndlyInstrumentPartiesRepeatingGroup() NoUndlyInstrumentPartiesRepeatingGroup {
	return NoUndlyInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UndlyInstrumentPartyID), quickfix.GroupElement(tag.UndlyInstrumentPartyIDSource), quickfix.GroupElement(tag.UndlyInstrumentPartyRole), NewNoUndlyInstrumentPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoUndlyInstrumentParties to this group
func (m NoUndlyInstrumentPartiesRepeatingGroup) Add() NoUndlyInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentParties{g}
}

//Get returns the ith NoUndlyInstrumentParties in the NoUndlyInstrumentPartiesRepeatinGroup
func (m NoUndlyInstrumentPartiesRepeatingGroup) Get(i int) NoUndlyInstrumentParties {
	return NoUndlyInstrumentParties{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingsRepeatingGroup is a repeating group, Tag 711
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingAllocationPercent), quickfix.GroupElement(tag.UnderlyingSettlementType), quickfix.GroupElement(tag.UnderlyingCashAmount), quickfix.GroupElement(tag.UnderlyingCashType), quickfix.GroupElement(tag.UnderlyingUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingTimeUnit), quickfix.GroupElement(tag.UnderlyingCapValue), NewNoUndlyInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingSettlMethod), quickfix.GroupElement(tag.UnderlyingAdjustedQuantity), quickfix.GroupElement(tag.UnderlyingFXRate), quickfix.GroupElement(tag.UnderlyingFXRateCalc), quickfix.GroupElement(tag.UnderlyingMaturityTime), quickfix.GroupElement(tag.UnderlyingPutOrCall), quickfix.GroupElement(tag.UnderlyingExerciseStyle), quickfix.GroupElement(tag.UnderlyingUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasureQty)})}
}

//Add create and append a new NoUnderlyings to this group
func (m NoUnderlyingsRepeatingGroup) Add() NoUnderlyings {
	g := m.RepeatingGroup.Add()
	return NoUnderlyings{g}
}

//Get returns the ith NoUnderlyings in the NoUnderlyingsRepeatinGroup
func (m NoUnderlyingsRepeatingGroup) Get(i int) NoUnderlyings {
	return NoUnderlyings{m.RepeatingGroup.Get(i)}
}

//NoTrdRegTimestamps is a repeating group element, Tag 768
type NoTrdRegTimestamps struct {
	quickfix.Group
}

//SetTrdRegTimestamp sets TrdRegTimestamp, Tag 769
func (m NoTrdRegTimestamps) SetTrdRegTimestamp(v time.Time) {
	m.Set(field.NewTrdRegTimestamp(v))
}

//SetTrdRegTimestampType sets TrdRegTimestampType, Tag 770
func (m NoTrdRegTimestamps) SetTrdRegTimestampType(v enum.TrdRegTimestampType) {
	m.Set(field.NewTrdRegTimestampType(v))
}

//SetTrdRegTimestampOrigin sets TrdRegTimestampOrigin, Tag 771
func (m NoTrdRegTimestamps) SetTrdRegTimestampOrigin(v string) {
	m.Set(field.NewTrdRegTimestampOrigin(v))
}

//SetDeskType sets DeskType, Tag 1033
func (m NoTrdRegTimestamps) SetDeskType(v enum.DeskType) {
	m.Set(field.NewDeskType(v))
}

//SetDeskTypeSource sets DeskTypeSource, Tag 1034
func (m NoTrdRegTimestamps) SetDeskTypeSource(v enum.DeskTypeSource) {
	m.Set(field.NewDeskTypeSource(v))
}

//SetDeskOrderHandlingInst sets DeskOrderHandlingInst, Tag 1035
func (m NoTrdRegTimestamps) SetDeskOrderHandlingInst(v enum.DeskOrderHandlingInst) {
	m.Set(field.NewDeskOrderHandlingInst(v))
}

//GetTrdRegTimestamp gets TrdRegTimestamp, Tag 769
func (m NoTrdRegTimestamps) GetTrdRegTimestamp() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TrdRegTimestampField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTrdRegTimestampType gets TrdRegTimestampType, Tag 770
func (m NoTrdRegTimestamps) GetTrdRegTimestampType() (v enum.TrdRegTimestampType, err quickfix.MessageRejectError) {
	var f field.TrdRegTimestampTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTrdRegTimestampOrigin gets TrdRegTimestampOrigin, Tag 771
func (m NoTrdRegTimestamps) GetTrdRegTimestampOrigin() (v string, err quickfix.MessageRejectError) {
	var f field.TrdRegTimestampOriginField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDeskType gets DeskType, Tag 1033
func (m NoTrdRegTimestamps) GetDeskType() (v enum.DeskType, err quickfix.MessageRejectError) {
	var f field.DeskTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDeskTypeSource gets DeskTypeSource, Tag 1034
func (m NoTrdRegTimestamps) GetDeskTypeSource() (v enum.DeskTypeSource, err quickfix.MessageRejectError) {
	var f field.DeskTypeSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDeskOrderHandlingInst gets DeskOrderHandlingInst, Tag 1035
func (m NoTrdRegTimestamps) GetDeskOrderHandlingInst() (v enum.DeskOrderHandlingInst, err quickfix.MessageRejectError) {
	var f field.DeskOrderHandlingInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasTrdRegTimestamp returns true if TrdRegTimestamp is present, Tag 769
func (m NoTrdRegTimestamps) HasTrdRegTimestamp() bool {
	return m.Has(tag.TrdRegTimestamp)
}

//HasTrdRegTimestampType returns true if TrdRegTimestampType is present, Tag 770
func (m NoTrdRegTimestamps) HasTrdRegTimestampType() bool {
	return m.Has(tag.TrdRegTimestampType)
}

//HasTrdRegTimestampOrigin returns true if TrdRegTimestampOrigin is present, Tag 771
func (m NoTrdRegTimestamps) HasTrdRegTimestampOrigin() bool {
	return m.Has(tag.TrdRegTimestampOrigin)
}

//HasDeskType returns true if DeskType is present, Tag 1033
func (m NoTrdRegTimestamps) HasDeskType() bool {
	return m.Has(tag.DeskType)
}

//HasDeskTypeSource returns true if DeskTypeSource is present, Tag 1034
func (m NoTrdRegTimestamps) HasDeskTypeSource() bool {
	return m.Has(tag.DeskTypeSource)
}

//HasDeskOrderHandlingInst returns true if DeskOrderHandlingInst is present, Tag 1035
func (m NoTrdRegTimestamps) HasDeskOrderHandlingInst() bool {
	return m.Has(tag.DeskOrderHandlingInst)
}

//NoTrdRegTimestampsRepeatingGroup is a repeating group, Tag 768
type NoTrdRegTimestampsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTrdRegTimestampsRepeatingGroup returns an initialized, NoTrdRegTimestampsRepeatingGroup
func NewNoTrdRegTimestampsRepeatingGroup() NoTrdRegTimestampsRepeatingGroup {
	return NoTrdRegTimestampsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTrdRegTimestamps,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TrdRegTimestamp), quickfix.GroupElement(tag.TrdRegTimestampType), quickfix.GroupElement(tag.TrdRegTimestampOrigin), quickfix.GroupElement(tag.DeskType), quickfix.GroupElement(tag.DeskTypeSource), quickfix.GroupElement(tag.DeskOrderHandlingInst)})}
}

//Add create and append a new NoTrdRegTimestamps to this group
func (m NoTrdRegTimestampsRepeatingGroup) Add() NoTrdRegTimestamps {
	g := m.RepeatingGroup.Add()
	return NoTrdRegTimestamps{g}
}

//Get returns the ith NoTrdRegTimestamps in the NoTrdRegTimestampsRepeatinGroup
func (m NoTrdRegTimestampsRepeatingGroup) Get(i int) NoTrdRegTimestamps {
	return NoTrdRegTimestamps{m.RepeatingGroup.Get(i)}
}

//NoEvents is a repeating group element, Tag 864
type NoEvents struct {
	quickfix.Group
}

//SetEventType sets EventType, Tag 865
func (m NoEvents) SetEventType(v enum.EventType) {
	m.Set(field.NewEventType(v))
}

//SetEventDate sets EventDate, Tag 866
func (m NoEvents) SetEventDate(v string) {
	m.Set(field.NewEventDate(v))
}

//SetEventPx sets EventPx, Tag 867
func (m NoEvents) SetEventPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewEventPx(value, scale))
}

//SetEventText sets EventText, Tag 868
func (m NoEvents) SetEventText(v string) {
	m.Set(field.NewEventText(v))
}

//SetEventTime sets EventTime, Tag 1145
func (m NoEvents) SetEventTime(v time.Time) {
	m.Set(field.NewEventTime(v))
}

//GetEventType gets EventType, Tag 865
func (m NoEvents) GetEventType() (v enum.EventType, err quickfix.MessageRejectError) {
	var f field.EventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEventDate gets EventDate, Tag 866
func (m NoEvents) GetEventDate() (v string, err quickfix.MessageRejectError) {
	var f field.EventDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEventPx gets EventPx, Tag 867
func (m NoEvents) GetEventPx() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.EventPxField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetEventText gets EventText, Tag 868
func (m NoEvents) GetEventText() (v string, err quickfix.MessageRejectError) {
	var f field.EventTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEventTime gets EventTime, Tag 1145
func (m NoEvents) GetEventTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.EventTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasEventType returns true if EventType is present, Tag 865
func (m NoEvents) HasEventType() bool {
	return m.Has(tag.EventType)
}

//HasEventDate returns true if EventDate is present, Tag 866
func (m NoEvents) HasEventDate() bool {
	return m.Has(tag.EventDate)
}

//HasEventPx returns true if EventPx is present, Tag 867
func (m NoEvents) HasEventPx() bool {
	return m.Has(tag.EventPx)
}

//HasEventText returns true if EventText is present, Tag 868
func (m NoEvents) HasEventText() bool {
	return m.Has(tag.EventText)
}

//HasEventTime returns true if EventTime is present, Tag 1145
func (m NoEvents) HasEventTime() bool {
	return m.Has(tag.EventTime)
}

//NoEventsRepeatingGroup is a repeating group, Tag 864
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText), quickfix.GroupElement(tag.EventTime)})}
}

//Add create and append a new NoEvents to this group
func (m NoEventsRepeatingGroup) Add() NoEvents {
	g := m.RepeatingGroup.Add()
	return NoEvents{g}
}

//Get returns the ith NoEvents in the NoEventsRepeatinGroup
func (m NoEventsRepeatingGroup) Get(i int) NoEvents {
	return NoEvents{m.RepeatingGroup.Get(i)}
}

//NoStrategyParameters is a repeating group element, Tag 957
type NoStrategyParameters struct {
	quickfix.Group
}

//SetStrategyParameterName sets StrategyParameterName, Tag 958
func (m NoStrategyParameters) SetStrategyParameterName(v string) {
	m.Set(field.NewStrategyParameterName(v))
}

//SetStrategyParameterType sets StrategyParameterType, Tag 959
func (m NoStrategyParameters) SetStrategyParameterType(v enum.StrategyParameterType) {
	m.Set(field.NewStrategyParameterType(v))
}

//SetStrategyParameterValue sets StrategyParameterValue, Tag 960
func (m NoStrategyParameters) SetStrategyParameterValue(v string) {
	m.Set(field.NewStrategyParameterValue(v))
}

//GetStrategyParameterName gets StrategyParameterName, Tag 958
func (m NoStrategyParameters) GetStrategyParameterName() (v string, err quickfix.MessageRejectError) {
	var f field.StrategyParameterNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrategyParameterType gets StrategyParameterType, Tag 959
func (m NoStrategyParameters) GetStrategyParameterType() (v enum.StrategyParameterType, err quickfix.MessageRejectError) {
	var f field.StrategyParameterTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrategyParameterValue gets StrategyParameterValue, Tag 960
func (m NoStrategyParameters) GetStrategyParameterValue() (v string, err quickfix.MessageRejectError) {
	var f field.StrategyParameterValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasStrategyParameterName returns true if StrategyParameterName is present, Tag 958
func (m NoStrategyParameters) HasStrategyParameterName() bool {
	return m.Has(tag.StrategyParameterName)
}

//HasStrategyParameterType returns true if StrategyParameterType is present, Tag 959
func (m NoStrategyParameters) HasStrategyParameterType() bool {
	return m.Has(tag.StrategyParameterType)
}

//HasStrategyParameterValue returns true if StrategyParameterValue is present, Tag 960
func (m NoStrategyParameters) HasStrategyParameterValue() bool {
	return m.Has(tag.StrategyParameterValue)
}

//NoStrategyParametersRepeatingGroup is a repeating group, Tag 957
type NoStrategyParametersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoStrategyParametersRepeatingGroup returns an initialized, NoStrategyParametersRepeatingGroup
func NewNoStrategyParametersRepeatingGroup() NoStrategyParametersRepeatingGroup {
	return NoStrategyParametersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoStrategyParameters,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StrategyParameterName), quickfix.GroupElement(tag.StrategyParameterType), quickfix.GroupElement(tag.StrategyParameterValue)})}
}

//Add create and append a new NoStrategyParameters to this group
func (m NoStrategyParametersRepeatingGroup) Add() NoStrategyParameters {
	g := m.RepeatingGroup.Add()
	return NoStrategyParameters{g}
}

//Get returns the ith NoStrategyParameters in the NoStrategyParametersRepeatinGroup
func (m NoStrategyParametersRepeatingGroup) Get(i int) NoStrategyParameters {
	return NoStrategyParameters{m.RepeatingGroup.Get(i)}
}

//NoInstrumentParties is a repeating group element, Tag 1018
type NoInstrumentParties struct {
	quickfix.Group
}

//SetInstrumentPartyID sets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) SetInstrumentPartyID(v string) {
	m.Set(field.NewInstrumentPartyID(v))
}

//SetInstrumentPartyIDSource sets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) SetInstrumentPartyIDSource(v string) {
	m.Set(field.NewInstrumentPartyIDSource(v))
}

//SetInstrumentPartyRole sets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) SetInstrumentPartyRole(v int) {
	m.Set(field.NewInstrumentPartyRole(v))
}

//SetNoInstrumentPartySubIDs sets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) SetNoInstrumentPartySubIDs(f NoInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetInstrumentPartyID gets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) GetInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrumentPartyIDSource gets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) GetInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrumentPartyRole gets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) GetInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoInstrumentPartySubIDs gets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) GetNoInstrumentPartySubIDs() (NoInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasInstrumentPartyID returns true if InstrumentPartyID is present, Tag 1019
func (m NoInstrumentParties) HasInstrumentPartyID() bool {
	return m.Has(tag.InstrumentPartyID)
}

//HasInstrumentPartyIDSource returns true if InstrumentPartyIDSource is present, Tag 1050
func (m NoInstrumentParties) HasInstrumentPartyIDSource() bool {
	return m.Has(tag.InstrumentPartyIDSource)
}

//HasInstrumentPartyRole returns true if InstrumentPartyRole is present, Tag 1051
func (m NoInstrumentParties) HasInstrumentPartyRole() bool {
	return m.Has(tag.InstrumentPartyRole)
}

//HasNoInstrumentPartySubIDs returns true if NoInstrumentPartySubIDs is present, Tag 1052
func (m NoInstrumentParties) HasNoInstrumentPartySubIDs() bool {
	return m.Has(tag.NoInstrumentPartySubIDs)
}

//NoInstrumentPartySubIDs is a repeating group element, Tag 1052
type NoInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetInstrumentPartySubID sets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubID(v string) {
	m.Set(field.NewInstrumentPartySubID(v))
}

//SetInstrumentPartySubIDType sets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubIDType(v int) {
	m.Set(field.NewInstrumentPartySubIDType(v))
}

//GetInstrumentPartySubID gets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrumentPartySubIDType gets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasInstrumentPartySubID returns true if InstrumentPartySubID is present, Tag 1053
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubID() bool {
	return m.Has(tag.InstrumentPartySubID)
}

//HasInstrumentPartySubIDType returns true if InstrumentPartySubIDType is present, Tag 1054
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubIDType() bool {
	return m.Has(tag.InstrumentPartySubIDType)
}

//NoInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1052
type NoInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoInstrumentPartySubIDsRepeatingGroup returns an initialized, NoInstrumentPartySubIDsRepeatingGroup
func NewNoInstrumentPartySubIDsRepeatingGroup() NoInstrumentPartySubIDsRepeatingGroup {
	return NoInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartySubID), quickfix.GroupElement(tag.InstrumentPartySubIDType)})}
}

//Add create and append a new NoInstrumentPartySubIDs to this group
func (m NoInstrumentPartySubIDsRepeatingGroup) Add() NoInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoInstrumentPartySubIDs{g}
}

//Get returns the ith NoInstrumentPartySubIDs in the NoInstrumentPartySubIDsRepeatinGroup
func (m NoInstrumentPartySubIDsRepeatingGroup) Get(i int) NoInstrumentPartySubIDs {
	return NoInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoInstrumentPartiesRepeatingGroup is a repeating group, Tag 1018
type NoInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoInstrumentPartiesRepeatingGroup returns an initialized, NoInstrumentPartiesRepeatingGroup
func NewNoInstrumentPartiesRepeatingGroup() NoInstrumentPartiesRepeatingGroup {
	return NoInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartyID), quickfix.GroupElement(tag.InstrumentPartyIDSource), quickfix.GroupElement(tag.InstrumentPartyRole), NewNoInstrumentPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoInstrumentParties to this group
func (m NoInstrumentPartiesRepeatingGroup) Add() NoInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoInstrumentParties{g}
}

//Get returns the ith NoInstrumentParties in the NoInstrumentPartiesRepeatinGroup
func (m NoInstrumentPartiesRepeatingGroup) Get(i int) NoInstrumentParties {
	return NoInstrumentParties{m.RepeatingGroup.Get(i)}
}

//NoFills is a repeating group element, Tag 1362
type NoFills struct {
	quickfix.Group
}

//SetFillExecID sets FillExecID, Tag 1363
func (m NoFills) SetFillExecID(v string) {
	m.Set(field.NewFillExecID(v))
}

//SetFillPx sets FillPx, Tag 1364
func (m NoFills) SetFillPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewFillPx(value, scale))
}

//SetFillQty sets FillQty, Tag 1365
func (m NoFills) SetFillQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewFillQty(value, scale))
}

//SetNoNested4PartyIDs sets NoNested4PartyIDs, Tag 1414
func (m NoFills) SetNoNested4PartyIDs(f NoNested4PartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetFillExecID gets FillExecID, Tag 1363
func (m NoFills) GetFillExecID() (v string, err quickfix.MessageRejectError) {
	var f field.FillExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFillPx gets FillPx, Tag 1364
func (m NoFills) GetFillPx() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.FillPxField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetFillQty gets FillQty, Tag 1365
func (m NoFills) GetFillQty() (v decimal.Decimal, scale int32, err quickfix.MessageRejectError) {
	var f field.FillQtyField
	if err = m.Get(&f); err == nil {
		v = f.Decimal
		scale = f.Scale
	}
	return
}

//GetNoNested4PartyIDs gets NoNested4PartyIDs, Tag 1414
func (m NoFills) GetNoNested4PartyIDs() (NoNested4PartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested4PartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasFillExecID returns true if FillExecID is present, Tag 1363
func (m NoFills) HasFillExecID() bool {
	return m.Has(tag.FillExecID)
}

//HasFillPx returns true if FillPx is present, Tag 1364
func (m NoFills) HasFillPx() bool {
	return m.Has(tag.FillPx)
}

//HasFillQty returns true if FillQty is present, Tag 1365
func (m NoFills) HasFillQty() bool {
	return m.Has(tag.FillQty)
}

//HasNoNested4PartyIDs returns true if NoNested4PartyIDs is present, Tag 1414
func (m NoFills) HasNoNested4PartyIDs() bool {
	return m.Has(tag.NoNested4PartyIDs)
}

//NoNested4PartyIDs is a repeating group element, Tag 1414
type NoNested4PartyIDs struct {
	quickfix.Group
}

//SetNested4PartyID sets Nested4PartyID, Tag 1415
func (m NoNested4PartyIDs) SetNested4PartyID(v string) {
	m.Set(field.NewNested4PartyID(v))
}

//SetNested4PartyIDSource sets Nested4PartyIDSource, Tag 1416
func (m NoNested4PartyIDs) SetNested4PartyIDSource(v string) {
	m.Set(field.NewNested4PartyIDSource(v))
}

//SetNested4PartyRole sets Nested4PartyRole, Tag 1417
func (m NoNested4PartyIDs) SetNested4PartyRole(v int) {
	m.Set(field.NewNested4PartyRole(v))
}

//SetNoNested4PartySubIDs sets NoNested4PartySubIDs, Tag 1413
func (m NoNested4PartyIDs) SetNoNested4PartySubIDs(f NoNested4PartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNested4PartyID gets Nested4PartyID, Tag 1415
func (m NoNested4PartyIDs) GetNested4PartyID() (v string, err quickfix.MessageRejectError) {
	var f field.Nested4PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNested4PartyIDSource gets Nested4PartyIDSource, Tag 1416
func (m NoNested4PartyIDs) GetNested4PartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.Nested4PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNested4PartyRole gets Nested4PartyRole, Tag 1417
func (m NoNested4PartyIDs) GetNested4PartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.Nested4PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoNested4PartySubIDs gets NoNested4PartySubIDs, Tag 1413
func (m NoNested4PartyIDs) GetNoNested4PartySubIDs() (NoNested4PartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested4PartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNested4PartyID returns true if Nested4PartyID is present, Tag 1415
func (m NoNested4PartyIDs) HasNested4PartyID() bool {
	return m.Has(tag.Nested4PartyID)
}

//HasNested4PartyIDSource returns true if Nested4PartyIDSource is present, Tag 1416
func (m NoNested4PartyIDs) HasNested4PartyIDSource() bool {
	return m.Has(tag.Nested4PartyIDSource)
}

//HasNested4PartyRole returns true if Nested4PartyRole is present, Tag 1417
func (m NoNested4PartyIDs) HasNested4PartyRole() bool {
	return m.Has(tag.Nested4PartyRole)
}

//HasNoNested4PartySubIDs returns true if NoNested4PartySubIDs is present, Tag 1413
func (m NoNested4PartyIDs) HasNoNested4PartySubIDs() bool {
	return m.Has(tag.NoNested4PartySubIDs)
}

//NoNested4PartySubIDs is a repeating group element, Tag 1413
type NoNested4PartySubIDs struct {
	quickfix.Group
}

//SetNested4PartySubID sets Nested4PartySubID, Tag 1412
func (m NoNested4PartySubIDs) SetNested4PartySubID(v string) {
	m.Set(field.NewNested4PartySubID(v))
}

//SetNested4PartySubIDType sets Nested4PartySubIDType, Tag 1411
func (m NoNested4PartySubIDs) SetNested4PartySubIDType(v int) {
	m.Set(field.NewNested4PartySubIDType(v))
}

//GetNested4PartySubID gets Nested4PartySubID, Tag 1412
func (m NoNested4PartySubIDs) GetNested4PartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.Nested4PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNested4PartySubIDType gets Nested4PartySubIDType, Tag 1411
func (m NoNested4PartySubIDs) GetNested4PartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.Nested4PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasNested4PartySubID returns true if Nested4PartySubID is present, Tag 1412
func (m NoNested4PartySubIDs) HasNested4PartySubID() bool {
	return m.Has(tag.Nested4PartySubID)
}

//HasNested4PartySubIDType returns true if Nested4PartySubIDType is present, Tag 1411
func (m NoNested4PartySubIDs) HasNested4PartySubIDType() bool {
	return m.Has(tag.Nested4PartySubIDType)
}

//NoNested4PartySubIDsRepeatingGroup is a repeating group, Tag 1413
type NoNested4PartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested4PartySubIDsRepeatingGroup returns an initialized, NoNested4PartySubIDsRepeatingGroup
func NewNoNested4PartySubIDsRepeatingGroup() NoNested4PartySubIDsRepeatingGroup {
	return NoNested4PartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested4PartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested4PartySubID), quickfix.GroupElement(tag.Nested4PartySubIDType)})}
}

//Add create and append a new NoNested4PartySubIDs to this group
func (m NoNested4PartySubIDsRepeatingGroup) Add() NoNested4PartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNested4PartySubIDs{g}
}

//Get returns the ith NoNested4PartySubIDs in the NoNested4PartySubIDsRepeatinGroup
func (m NoNested4PartySubIDsRepeatingGroup) Get(i int) NoNested4PartySubIDs {
	return NoNested4PartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoNested4PartyIDsRepeatingGroup is a repeating group, Tag 1414
type NoNested4PartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested4PartyIDsRepeatingGroup returns an initialized, NoNested4PartyIDsRepeatingGroup
func NewNoNested4PartyIDsRepeatingGroup() NoNested4PartyIDsRepeatingGroup {
	return NoNested4PartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested4PartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested4PartyID), quickfix.GroupElement(tag.Nested4PartyIDSource), quickfix.GroupElement(tag.Nested4PartyRole), NewNoNested4PartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoNested4PartyIDs to this group
func (m NoNested4PartyIDsRepeatingGroup) Add() NoNested4PartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNested4PartyIDs{g}
}

//Get returns the ith NoNested4PartyIDs in the NoNested4PartyIDsRepeatinGroup
func (m NoNested4PartyIDsRepeatingGroup) Get(i int) NoNested4PartyIDs {
	return NoNested4PartyIDs{m.RepeatingGroup.Get(i)}
}

//NoFillsRepeatingGroup is a repeating group, Tag 1362
type NoFillsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoFillsRepeatingGroup returns an initialized, NoFillsRepeatingGroup
func NewNoFillsRepeatingGroup() NoFillsRepeatingGroup {
	return NoFillsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoFills,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.FillExecID), quickfix.GroupElement(tag.FillPx), quickfix.GroupElement(tag.FillQty), NewNoNested4PartyIDsRepeatingGroup()})}
}

//Add create and append a new NoFills to this group
func (m NoFillsRepeatingGroup) Add() NoFills {
	g := m.RepeatingGroup.Add()
	return NoFills{g}
}

//Get returns the ith NoFills in the NoFillsRepeatinGroup
func (m NoFillsRepeatingGroup) Get(i int) NoFills {
	return NoFills{m.RepeatingGroup.Get(i)}
}
