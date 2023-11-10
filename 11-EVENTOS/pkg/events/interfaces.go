package events

import "time"

type EventInterface interface {
	GetName() string         // obter o nome do evento
	GetDateTime() time.Time  // obter a datahora do evento
	GetPayLoad() interface{} // posso ter difersos payloads em diversos formatos
}

// o evento chama e o handler executa
// para executar um evento, ele precisa ser passado como parâmetro
type EventHandlerInterface interface {
	Handle(event EventInterface)
}

// metodo que registra os eventos
// registra os eventos e suas operações
// despacha (fire) o evento para que suas operações sejam executadas
type EventDispatcherInterface interface {
	// registra o nome do evento e
	// quando ele for registrado, ele será executado pelo Handler
	Register(eventName string, handler EventDispatcherInterface) error
	Dispatch(event EventInterface) error
	// tirar o evento da fila de eventos
	Remove(eventName string, handler EventHandlerInterface) error
	// verificar se há o eventName registrado e também o seu Handler
	Has(eventName string, handler EventHandlerInterface) bool
	// limpa o EventDispatcher apagando todos os eventos registrados
	Clear() error
}
