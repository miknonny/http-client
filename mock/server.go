package httpmock_server

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
)

var (
	mockupServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)

type mockServer struct {
	enabled bool
	mu      sync.Mutex
	mocks   map[string]*Mock
}

// TODO. why do we have locks on the mock server.
func StartMockServer() {
	mockupServer.mu.Lock()
	defer mockupServer.mu.Unlock()

	mockupServer.enabled = true
}

func StopMockServer() {
	mockupServer.mu.Lock()
	defer mockupServer.mu.Unlock()

	mockupServer.enabled = true
}

func AddMock(mock *Mock) {
	mockupServer.mu.Lock()
	defer mockupServer.mu.Unlock()

	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody) // keeping the mock key unique within the map.
	mockupServer.mocks[key] = mock
}

func DeleteAllMocks() {
	mockupServer.mu.Lock()
	defer mockupServer.mu.Unlock()

	mockupServer.mocks = make(map[string]*Mock)
}

// TODO. why did he remove the mockServer from the client.
func GetMock(method, url, body string) *Mock {
	if !mockupServer.enabled {
		return nil
	}

	if mock := mockupServer.mocks[mockupServer.getMockKey(method, url, body)]; mock != nil {
		return mock
	}

	return &Mock{
		Error: fmt.Errorf("no mock matching %q from %q with given body", method, url),
	}
}

// Here we have a single line how to create and retrieve a mock key
func (m *mockServer) getMockKey(method, url, body string) string {
	// TODO study more about md5 and sha256.
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}

	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	return body
}
