// Copyright © 2017 Louis Taylor <louis@kragniz.eu>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package proxy

import (
	"os"
	"strings"
	"syscall"
)

var ProxyEnvs = []string{
	"http_proxy",
	"https_proxy",
	"HTTP_PROXY",
	"HTTPS_PROXY",
}

var NoProxy = []string{
	"no_proxy",
	"NO_PROXY",
}

func replaceShell() {
	shell := os.Getenv("SHELL")

	syscall.Exec(shell, []string{shell}, syscall.Environ())
}

func On(proxy string) {
	for _, env := range ProxyEnvs {
		os.Setenv(env, proxy)
	}
	replaceShell()
}

func Off() {
	for _, env := range ProxyEnvs {
		os.Unsetenv(env)
	}
	replaceShell()
}

func ignore(noProxy string, host string) string {
	hosts := os.Getenv(noProxy)
	if len(hosts) > 0 {
		return strings.Join([]string{hosts, host}, ",")
	} else {
		return host
	}
}

func Ignore(host string) {
	for _, env := range NoProxy {
		newEnv := ignore(env, host)
		os.Setenv(env, newEnv)
	}
	replaceShell()
}
