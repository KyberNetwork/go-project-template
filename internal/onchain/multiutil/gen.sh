solc --abi --bin --overwrite multiutil.sol -o ./
abigen --abi=MultiUtil.abi --bin=MultiUtil.bin --pkg=multiutil --out=multiutil.go -type MultiUtil