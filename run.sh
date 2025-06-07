#!/system/bin/sh

original_dir=$(pwd)
if [ ! -d "$original_dir" ]; then
	echo "[ERROR] Failed to determine original directory."
	echo "[HINT] Please run this script from a terminal session."
	sleep 5
	exit 1
fi

dir=$(
	cd "$(dirname "$0")" || exit 1
	pwd
)

arch=$(getprop ro.product.cpu.abi)
echo "[INFO] Architecture: $arch"

if [ -z "$arch" ]; then
	echo "[ERROR] Failed to detect architecture."
	sleep 5
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
	echo "[HINT] Please run manually from terminal and verify architecture support."
	sleep 5
	exit 1
	;;
esac

parent=$(basename "$dir")
tmp="/data/local/tmp"
target="$tmp/$parent"

pgrep -fl server | grep '\./' | while read -r pid _; do
	kill "$pid"
done

rm -rf "$target"
mkdir -p "$target"
cp -r "$dir"/* "$target"/
chmod +x "$target/server"

cd "$target" || {
	echo "[ERROR] Failed to enter target directory: $target"
	sleep 5
	exit 1
}
nohup ./server >/dev/null 2>&1 &
sleep 1

am start -a android.intent.action.VIEW -d http://localhost:8080

cd "$original_dir" || {
	echo "[ERROR] Cannot return to original directory: $original_dir"
	echo "[HINT] Please return to it manually via terminal."
	sleep 5
	exit 1
}
