package opus

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -systemdll=false -output opus_windows.go genopus_windows.go

//sys opus_decoder_get_size(channles int)(size int) = opus.opus_decoder_get_size
//sys opus_decoder_init(p unsafe.Pointer, sampleRate int32, channels int)(err int) = opus.opus_decoder_init
//sys opus_decode(p unsafe.Pointer, data []byte, length int32, out []int16, cap int)(n int) = opus.opus_decode
//sys opus_decode_float(p unsafe.Pointer, data []byte, length int32, out []float32, cap int)(n int) = opus.opus_decode_float

//sys opus_encoder_get_size(channles int)(size int) = opus.opus_encoder_get_size
//sys opus_encoder_init(p unsafe.Pointer, sampleRate int32, channels int, application int)(err int) = opus.opus_encoder_init
//sys opus_encoder_ctl(p unsafe.Pointer, reqType int, req int32)(res int) = opus.opus_encoder_ctl
//sys opus_encoder_get_ctl(p unsafe.Pointer, reqType int, req *int32)(res int) = opus.opus_encoder_ctl
//sys opus_encode(p unsafe.Pointer, pcm []int16, length int, out []byte, cap int32)(n int) = opus.opus_encode
//sys opus_encode_float(p unsafe.Pointer, pcm []float32, length int, out []byte, cap int32)(n int) = opus.opus_encode_float
