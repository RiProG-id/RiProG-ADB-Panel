#!/system/bin/sh
original_dir=$(pwd)
dir=$(
	cd "$(dirname "$0")" || exit 1
	pwd
)
if [ ! -f "$dir/server" ]; then
	arch=$(getprop ro.product.cpu.abi)
	echo "[INFO] Architecture: $arch"
	if [ "$arch" = "armeabi-v7a" ] || [ "$arch" = "armv8l" ]; then
		echo "- Architecture $arch is supported."
		echo "- Installation continues."
		cp "$dir/arm" "$dir/server"
	elif [ "$arch" = "arm64-v8a" ]; then
		echo "- Architecture $arch is supported."
		echo "- Installation continues."
		cp "$dir/arm64" "$dir/server"
	else
		echo "- Architecture $arch is not supported."
		echo "- Installation aborted."
		cd "$original_dir" || exit 1
		exit 1
	fi
fi
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
