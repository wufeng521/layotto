/*
* Copyright 2021 Layotto Authors
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package main

import (
	"context"
	"fmt"

	"mosn.io/layotto/spec/proto/extension/v1/cryption"

	"google.golang.org/grpc"
)

const (
	storeName = "cryption_demo"
)

func TestEncrypt() []byte {
	conn, err := grpc.Dial("127.0.0.1:34904", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to establish connection: %+v", err)
		return nil
	}

	c := cryption.NewCryptionServiceClient(conn)

	req := &cryption.EncryptRequest{ComponentName: storeName, PlainText: []byte("Hello, world")}

	resp, err := c.Encrypt(context.Background(), req)
	if err != nil {
		fmt.Printf("failed to Encrypt data: %+v", err)
		return nil
	}
	return resp.CipherText
}

func TestDecrypt(data []byte) []byte {
	conn, err := grpc.Dial("127.0.0.1:34904", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to establish connection: %+v", err)
		return nil
	}

	c := cryption.NewCryptionServiceClient(conn)

	req := &cryption.DecryptRequest{ComponentName: storeName, CipherText: data}

	resp, err := c.Decrypt(context.Background(), req)
	if err != nil {
		fmt.Printf("failed to Decrypt: %+v", err)
		return nil
	}
	return resp.PlainText
}

func main() {
	encyptContent := TestEncrypt()
	fmt.Printf("加密后的数据为: \n")
	fmt.Println(string(encyptContent))
	decryptContent := TestDecrypt(encyptContent)
	fmt.Printf("解密后的数据为: \n")
	fmt.Println(string(decryptContent))
}
