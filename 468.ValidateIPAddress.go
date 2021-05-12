package code

import (
	"regexp"
	"strings"
)

// https://leetcode.com/problems/validate-ip-address/

func validIPAddress(IP string) string {
	if validIPV4(IP) {
		return "IPV4"
	}
	if validIPV6(IP) {
		return "IPV6"
	}
	return "Neither"
}

func validIPV4(IP string) bool {
	part := strings.Split(IP, ".")
	if len(part) != 4 {
		return false
	}
	pnum := `^(25[0-5])|(2[0-4]\d)|(1\d{2})|([1-9]\d{1})|(\d)$`
	for _, num := range part {
		match, _ := regexp.Match(pnum, []byte(num))
		if !match {
			return false
		}
	}
	return true
}

func validIPV6(IP string) bool {
	part := strings.Split(IP, ":")
	if len(part) != 8 {
		return false
	}
	pnum := `^([1-9a-eA-E][\da-eA-E]{3})|(([1-9a-eA-E][\da-eA-E]{2}))|(([1-9a-eA-E][\da-eA-E]{1}))|([\da-eA-E])$`
	for _, num := range part {
		match, _ := regexp.Match(pnum, []byte(num))
		if !match {
			return false
		}
	}
	return true
}
