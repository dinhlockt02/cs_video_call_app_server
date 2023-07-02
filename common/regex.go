package common

import "regexp"

var URLRegexp = regexp.MustCompile(`^(?:https?://)?(?:[^/.\s]+\.)*`)

var EmailRegexp = regexp.MustCompile(`^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$`)
