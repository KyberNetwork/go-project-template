solc --abi --bin --bin-runtime --overwrite simutil.sol -o ./
abigen --abi=SimUtil.abi --bin=SimUtil.bin --pkg=simutil --out=simutil.go -type SimUtil