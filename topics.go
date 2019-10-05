package main

import (
	set "github.com/deckarep/golang-set"
)

// Topics manages the subscribed topics and calls subscribe and unsubscribe as neccessary
type Topics struct {
	subscribedTopics set.Set

	subscribe   func(topic string)
	unsubscribe func(topic string)
}

// NewTopics create new topics state with subscibe and unsubscribe handlers
func NewTopics(subscribe func(topic string), unsubscribe func(topic string)) *Topics {
	return &Topics{
		subscribe:   subscribe,
		unsubscribe: unsubscribe,

		subscribedTopics: set.NewSet(),
	}
}

// UpdateTopics set the internal state of subscribed topics
// and calls the subscribe handler on all new topics and unsubscribe on all obsolte topics.
func (t *Topics) UpdateTopics(topics []string) {
	updatedTopics := setFromStringSlice(topics)

	if !t.subscribedTopics.Equal(updatedTopics) {
		newTopics := updatedTopics.Difference(t.subscribedTopics)
		for _, topic := range stringSliceFromSet(newTopics) {
			t.subscribe(topic)
		}

		obsoletTopics := t.subscribedTopics.Difference(updatedTopics)
		for _, topic := range stringSliceFromSet(obsoletTopics) {
			t.unsubscribe(topic)
		}
	}

	t.subscribedTopics = updatedTopics
}

func setFromStringSlice(strings []string) set.Set {
	stringsAsInterface := make([]interface{}, len(strings))
	for i := range stringsAsInterface {
		stringsAsInterface[i] = strings[i]
	}
	return set.NewSetFromSlice(stringsAsInterface)
}

func stringSliceFromSet(s set.Set) []string {
	strings := make([]string, s.Cardinality())

	for i, stringAsInterface := range s.ToSlice() {
		strings[i] = stringAsInterface.(string)
	}

	return strings
}
