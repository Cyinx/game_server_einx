@echo off
protoc.exe --gogofaster_out=../src/protobuf_gen *.proto

:end
echo finished
pause