package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var subscribed []string
var unsubscribed []string

func subscribe(topic string) {
	subscribed = append(subscribed, topic)
}

func unsubscribe(topic string) {
	unsubscribed = append(unsubscribed, topic)
}

func TestSubscribeTopicOnFirstAdd(t *testing.T) {
	topics := NewTopics(subscribe, unsubscribe)

	topics.UpdateTopics([]string{"foo"})

	assert.Equal(t, []string{"foo"}, subscribed)
	assert.Empty(t, unsubscribed)
}

func TestDontSubscribeOnAlreadySubscribedTopic(t *testing.T) {
	topics := NewTopics(subscribe, unsubscribe)
	topics.UpdateTopics([]string{"foo"})
	subscribed = []string{}

	topics.UpdateTopics([]string{"foo", "bar"})

	assert.Equal(t, []string{"bar"}, subscribed)
	assert.Empty(t, unsubscribed)
}

func TestCallUnsubribe(t *testing.T) {
	topics := NewTopics(subscribe, unsubscribe)
	topics.UpdateTopics([]string{"foo"})
	subscribed = []string{}

	topics.UpdateTopics([]string{"bar"})

	assert.Equal(t, []string{"foo"}, unsubscribed)
	assert.Equal(t, []string{"bar"}, subscribed)
}
