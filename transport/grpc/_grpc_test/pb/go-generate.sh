#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

protoExec=$(which "protoc")
if [ -z $protoExec ]; then
    echo 'Please install protoc!'
    echo "Please look readme.md to install proto3"
    echo "if you use centos7,please look https://github.com/daheige/go-proj/blob/master/docs/centos7-protoc-install.md"
    exit 0
fi

protos_dir=$root_dir/pb
pb_dir=$root_dir/pb

mkdir -p $pb_dir

#delete old pb code.
rm -rf $root_dir/pb/*.go

echo "\n\033[0;32mGenerating codes...\033[39;49;0m\n"

echo "generating golang stubs..."
cd $protos_dir

$protoExec -I $protos_dir \
    --go_out $root_dir/pb --go_opt paths=source_relative \
    --go-grpc_out $root_dir/pb --go-grpc_opt paths=source_relative \
    $protos_dir/*.proto

echo "generating golang code success"

echo "\n\033[0;32mGenerate codes successfully!\033[39;49;0m\n"

exit 0
