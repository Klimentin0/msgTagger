package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	taggedSlice := make([]sms, len(messages))
	for i, msg := range messages {
		tags := tagger(msg)
		msg.tags = append(msg.tags, tags...)
		taggedSlice[i] = msg
	}
	return taggedSlice
}

func tagger(msg sms) []string {
	tags := []string{}
	str := msg.content
	lowerContent := strings.ToLower(str)
	if strings.Contains(lowerContent, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(lowerContent, "sale") {
		tags = append(tags, "Promo")
	}
	//if !strings.Contains(lowerContent, "urgent") && !strings.Contains(lowerContent, "sale") {
	//	tags = append(tags, "")
	//}
	return tags
}
