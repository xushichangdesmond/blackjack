package blackjack

type CardListener func(Card)

type ShuffleListener func()

type listeners struct {
	ls []interface{}
}

func (listeners *listeners) Subscribe(listener interface{}) {
	listeners.ls = append(listeners.ls, listener)
}

func (listeners *listeners) Unsubscribe(listener interface{}) {
	for i, l := range listeners.ls {
		if l == listener {
			if i == len(listeners.ls) {
				listeners.ls = listeners.ls[:len(listeners.ls)-1]
				return
			}
			listeners.ls = append(listeners.ls[:i], listeners.ls[i+1:]...)
			return
		}
	}
}
