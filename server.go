/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"io"
	"net/http"

	//"github.com/microsoft/vscode-remote-try-go/hello"
	"github.com/microsoft/vscode-remote-try-go/matematica"
)

func handle(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, hello.Hello())
	resultado := matematica.Soma(1, 2)
	io.WriteString(w, resultado)
}

func main() {
	portNumber := "9000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
