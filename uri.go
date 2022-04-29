package helper

import (
	"errors"
	"net/url"
	"strings"
)

type UriValues map[string][]string

// ParseUriQueryToMap 解析url参数到map .
func (tu *TsUri) ParseUriQueryToMap(query string) map[string]interface{} {
	q := strings.Split(query, "?")
	if len(q) < 2 {
		return map[string]interface{}{}
	}
	qm := strings.Split(q[1], "&")
	convert := make(map[string]interface{}, len(qm))
	for _, v := range qm {
		split := strings.Split(v, "=")
		convert[split[0]] = split[1]
	}
	return convert
}

// GetQueryParams url params 不区分大小写 .
func (tu *TsUri) GetQueryParams(mValues map[string][]string, key string) (value interface{}) {
	paramsToLowerMap := make(map[string][]string)
	// key 统一转小写
	for pk, pv := range mValues {
		paramsToLowerMap[strings.ToLower(pk)] = pv
	}
	// 获取key是否存在
	if values, ok := paramsToLowerMap[strings.ToLower(key)]; ok {
		value = values[0]
		return
	}
	return
}

// GetDomain 获取域名 .
func (tu *TsUri) GetDomain(str string, isMains ...bool) (domain string) {
	u, err := url.Parse(str)
	isMain := false
	if len(isMains) > 0 {
		isMain = isMains[0]
	}

	if err != nil {
		return
	}
	if !isMain {
		domain = u.Hostname()
		return
	}

	parts := strings.Split(u.Hostname(), ".")
	domain = parts[len(parts)-2] + "." + parts[len(parts)-1]
	return
}

// ParseUriQuery 解析uri参数 .
func (tu *TsUri) ParseUriQuery(m UriValues, query string) (map[string][]string, error) {
	var err error
	q := strings.Split(query, "?")
	if len(q) < 2 {
		return m,errors.New("NotQuery")
	}
	queryStr := q[1]
	for queryStr != "" {
		key := queryStr
		if i := strings.IndexAny(key, "&;"); i >= 0 {
			key, queryStr = key[:i], key[i+1:]
		} else {
			queryStr = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		key, err1 := url.QueryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value, err1 = url.QueryUnescape(value)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		m[key] = append(m[key], value)
	}
	return m, err
}