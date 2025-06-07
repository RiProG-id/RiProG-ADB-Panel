#!/system/bin/sh
original_dir=$(pwd)
dir=$(
	cd "$(dirname "$0")" || exit 1
	pwd
)
arch=$(getprop ro.product.cpu.abi)
echo "[INFO] Architecture: $arch"
if [ -z "$arch" ]; then
	echo "[ERROR] Failed to detect architecture."
	exit 1
fi
case "$arch" in
armeabi-v7a | armv8l)
	cp "$dir/arm" "$dir/server"
	;;
arm64-v8a)
	cp "$dir/arm64" "$dir/server"
	;;
*)
	echo "[ERROR] Architecture '$arch' is not supported."
	cd "$original_dir" || exit 1
	exit 1
	;;
esac
parent=$(basename "$dir")
tmp="/data/local/tmp"
target="$tmp/$parent"
pgrep -fl server | grep '\./' | while read pid cmd; do
	kill "$pid"
done
rm -rf "$target"
mkdir -p "$target"
cp -r "$dir"/* "$target"/
chmod +x "$target/server"
cd "$target" || exit 1
nohup ./server >/dev/null 2>&1 &
sleep 1
am start -a android.intent.action.VIEW -d http://localhost:8080
cd "$original_dir" || exit 1
