/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
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
	"dubbo.apache.org/dubbo-go/v3/client"
	payment "github.com/apache/dubbo-go-samples/online_boutique_demo/src/paymentservice/proto"
	"github.com/dubbogo/gost/log/logger"
	"time"
)

func main() {
	cli, err := client.NewClient(
		client.WithClientURL("tri://127.0.01:20000"),
	)

	if err != nil {
		panic(err)
	}

	paymentService, err := payment.NewPaymentService(cli)
	if err != nil {
		panic(err)
	}

	req := &payment.ChargeRequest{
		Amount: &payment.Money{
			CurrencyCode: "USD",
			Units:        100,
			Nanos:        0,
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "4111111111111111",
			CreditCardCvv:             123,
			CreditCardExpirationYear:  2025,
			CreditCardExpirationMonth: 12,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := paymentService.Charge(ctx, req)
	if err != nil {
		logger.Fatalf("Failed to call Charge method: %v", err)
	}

	logger.Infof("Transaction ID: %s", resp.TransactionId)
}
