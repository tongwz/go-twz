package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/**
 * @note: 当主协程不被阻塞 子协程产生死锁  不会影响到程序报错和panic
 * @date: 2023/11/2 11:25
**/
var MuLock sync.Mutex

var ThisMyToken *MyToken

func init() {
	ThisMyToken = NewMyToken()
}

func NewMyToken() *MyToken {
	myToken := new(MyToken)
	myToken.mtx = &sync.Mutex{}

	return myToken
}

type MyToken struct {
	mtx      *sync.Mutex
	TokenMap map[string]*Token
}
type Token struct {
	AppId string
	Value string
	Ttl   int64
}

func (m *MyToken) GetToken(appId string) (token *Token, err error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	tokenValue, ok := m.TokenMap[appId]
	if !ok {
		return m.CreateToken(appId)
	}
	return tokenValue, nil
}

func (m *MyToken) CreateToken(appId string) (token *Token, err error) {
	token = new(Token)
	token.Value = "xxx"
	token.Ttl = 2323232
	token.AppId = appId
	go m.SetTokenMap(token)
	return token, nil
}

func (m *MyToken) SetTokenMap(token *Token) {
	m.mtx.Lock()
	// fmt.Printf("获得锁, %p \n", m)
	defer m.mtx.Lock()

	m.TokenMap[token.AppId] = token
}

func main() {
	var i int
	for {
		go fmtToken("appId" + strconv.Itoa(i))
		fmt.Printf("for循环进行中~ \n")
		time.Sleep(time.Second * 1)
		i++
		if i > 3 {
			i = 0
		}
	}
}

func fmtToken(appId string) {
	tokenInfo, _ := ThisMyToken.GetToken("app")

	fmt.Printf("获取到的token:%#v, map:%#v \n", tokenInfo, NewMyToken().TokenMap)
}
