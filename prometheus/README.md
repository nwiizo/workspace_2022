# HTTP API
現在の安定した`HTTP API`は、Prometheusサーバーの`/api/v1`にアクセス可能です。このエンドポイントに追加されたものは、そのエンドポイントに追加されます。

## Format overview
API の応答形式は JSON です。成功したAPIリクエストはすべて、`2xx`ステータスコードを返します。

APIハンドラに到達した無効なリクエストは、JSONエラーオブジェクトと以下のHTTPレスポンスコードのいずれかを返します。

- 400 パラメータがない、または間違っている場合のBad Request.
- 422 式が実行できない場合のUnprocessable Entity (RFC4918).
- 503 クエリーがタイムアウトまたはアボートした場合、サービスを利用できない.

API のエンドポイントに到達する前に発生したエラーの場合、他の 2xx 以外のコードが返されることがある。

リクエストの実行を阻害しない程度のエラーがある場合は、警告の配列が返されることがある。data フィールドには、正常に収集されたすべてのデータが返される。

JSONレスポンスエンベロープのフォーマットは以下のとおりです。

``` json
{
  "status": "success" | "error",
  "data": <data>,

  // Only set if status is "error". The data field may still hold
  // additional data.
  "errorType": "<string>",
  "error": "<string>",

  // Only if there were warnings while executing the request.
  // There will still be data in the data field.
  "warnings": ["<string>"]
}
```

一般的なプレースホルダーは、以下のように定義されています。

- <rfc3339 | unix_timestamp>: Input timestamps may be provided either in RFC3339 format or as a Unix timestamp in seconds, with optional decimal places for sub-second precision. Output timestamps are always represented as Unix timestamps in seconds.
- <series_selector>: Prometheus time series selectors like http_requests_total or http_requests_total{method=~"(GET|POST)"} and need to be URL-encoded.
- <duration>: Prometheus duration strings. For example, 5m refers to a duration of 5 minutes.
- <bool>: boolean values (strings true and false).
  
Note: 繰り返される可能性のあるクエリパラメータの名称は[]で終わる。

## Expression queries

