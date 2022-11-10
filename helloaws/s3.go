package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3PutObjectAPI はPutObject関数のインターフェイスを定義します。
// モックサービスを使用して関数をテストするために、このインターフェイスを使用します。
type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

// PutFile は、Amazon Simple Storage Service (Amazon S3)のバケットにファイルをアップロードします。
// input
// c はメソッド呼び出しのコンテキストで、AWS リージョンを含みます。
// api はメソッド呼び出しを定義するインターフェース
// input は、サービスコールへの入力引数を定義する。
// output
// 成功した場合、サービスコールの結果を含む PutObjectOutput オブジェクトと nil が出力される。
// それ以外の場合は、nil と PutObject の呼び出しによるエラー
func putFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}

// S3PutObject は、実際に外部から呼ばれて環境変数を設定して、PutFileを呼び出します。
// input
// バケット名,パス（s3内でのディレクトリ分け用）を引数に受け取る
// output
// 成功した場合には、サービスコールの結果を含む s3.PutObjectOutput とnil が出力される
// それ以外の場合は、nil と err の呼び出しによるエラー
func S3PutObject(bucket string, filename string) (*s3.PutObjectOutput, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Unable to open file " + filename)
		return nil, err
	}

	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &filename,
		Body:   file,
	}

	result_s3_put_object, err := putFile(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got error uploading file:")
		fmt.Println(err)
		return nil, err
	}
	return result_s3_put_object, err
}
