/*
package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	// как сохранять информацию об оплатах?
	// 1 - статический массив
	
	// 2 - слайс

	PaymentInfo map[int]PaymentInfo

	// 3 - мапа

	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule{
	return &PaymentModule{
		paymentsInfo: make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

// Метод Pay()
// Принимает:
// 1. Описание проводимой оплаты
// 2. Сумма оплаты
// Возвращает:
// 1. ID проведённой операции
func (p *PaymentModule) Pay(description string, usd int) int {
	id := p.paymentMethod.Pay(usd)

	info := PaymentInfo{
		Description: description,
		Usd: usd,
		Cancelled: false,
	}
	// 1. проводить оплату
	// 2. получать id проведённой оплаты

	// 3. сохранять информацию о проведённой операции
	// - описание операции
	// - сколько было потрачено
	// - отменённая ли операция
	p.PaymentInfo[id] = info

	// 4. возвращать id проведённой оплаты
	return id
}

// Метод Cancel()
// Принимает:
// 1. ID операции
// Возвращает:
// - ничего не возвращает
func (p PaymentModule) Cancel(id int) {
	p.paymentMethod.Cancel(id)

	info, ok := p.PaymentInfo[id]
	info.Cancelled = true
}

// Будет возвращать какой-то одной информации
// Метод info()
// Принимает:
// 1.ID операции
// Возвращает:
// Информацию о проведённой операции
func (p PaymentModule) Info() {}

// Будет возвращать информацию о всех проведенных операциях
// Метод AllInfo()
// Принимает:
// - ничего не принимает
// Возвращает:
// 1. Информаццию о всех проведённых операциях
func (p PaymentModule) AllInfo() {}
*/