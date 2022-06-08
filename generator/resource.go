package generator

import "os"

func getAppApiNames() []string {
	return []string{
		"app-api",
		"post-api",
		"life-message-api",
		"public-api",
		"life-wechat-api",
		"life-group-chat-api",
		"focus-on-api",
		"circle-api",
		"report-opinion-api",
		"topic-api",
		"activity-api",
		"life-location-api",
		"points-api",
		"goods-api",
		"order-api",
		"points-activity-api",
		"comment-api",
		"pinhn-gateway-api",
		"pinhn-user-api",
		"pinhn-moment-api",
		"firm-api",
		"pinhn-like-api",
		"information-api",
		"pinhn-certify-api",
	}
}

func getAdminApiNames() []string {
	return []string{
		"admin-activity-api",
		"admin-enterprise-api",
		"admin-goods-api",
		"admin-message-api",
		"admin-order-api",
		"admin-points-activity-api",
		"admin-points-api",
		"admin-post-api",
		"admin-public-api",
		"admin-topic-api",
		"admin-user-api",
		"industry-circle-admin-api",
		"life-admin-group-chat-api",
		"admin-content-api",
		"admin-comment-api",
		"admin-report-opinion-api",
		"pinhn-admin-api",
		"admin-app-api",
		"pinhn-admin-moment-api",
	}
}

func getMpApiNames() []string {
	return []string{
		"circle-mp-api",
		"focus-on-mp-api",
		"life-group-chat-mp-api",
		"life-message-mp-api",
		"location-mp-api",
		"post-mp-api",
		"public-mp-api",
		"topic-mp-api",
		"wechat-mp-api",
		"moment-mp-api",
		"gateway-mp-api",
		"comment-mp-api",
		"report-opinion-mp-api",
		"firm-mp-api",
		"information-mp-api",
		"account-mp-api",
		"user-mp-api",
		"like-mp-api",
	}
}

func dirOrFileExists(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
