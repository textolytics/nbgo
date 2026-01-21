
#!/usr/bin/env bash
set -e

cat <<'WSLFIX'
#!/usr/bin/env bash
# =========================================================
# Fix JupyterLab Desktop startup in WSL Ubuntu
# =========================================================

echo "=== JupyterLab Desktop WSL Fix ==="

# 1. Disable GPU & sandbox for Electron
export ELECTRON_DISABLE_GPU=1
export ELECTRON_ENABLE_LOGGING=1
export ELECTRON_NO_ATTACH_CONSOLE=1

# Chromium flags (critical)
export ELECTRON_EXTRA_LAUNCH_ARGS="\
--disable-gpu \
--disable-software-rasterizer \
--disable-dev-shm-usage \
--disable-setuid-sandbox \
--no-sandbox \
--use-gl=swiftshader \
"

# 2. Stub DBus (prevents bus.cc errors)
export DBUS_SESSION_BUS_ADDRESS="unix:path=/dev/null"

# 3. Force software OpenGL
export LIBGL_ALWAYS_SOFTWARE=1
export GALLIUM_DRIVER=softpipe

# 4. GTK / GLib safety
export GSETTINGS_SCHEMA_DIR=/usr/share/glib-2.0/schemas

# 5. XDG_RUNTIME_DIR workaround (often missing in WSL)
export XDG_RUNTIME_DIR=/tmp/runtime-$USER
mkdir -p "$XDG_RUNTIME_DIR"
chmod 700 "$XDG_RUNTIME_DIR"

# 6. Diagnostics
echo "[INFO] WSL detected:"
grep -i wsl /proc/version || true

echo "[INFO] Environment:"
echo "  DBUS_SESSION_BUS_ADDRESS=$DBUS_SESSION_BUS_ADDRESS"
echo "  XDG_RUNTIME_DIR=$XDG_RUNTIME_DIR"

# 7. Launch JupyterLab Desktop safely
echo "[INFO] Launching jupyterlab-desktop..."
exec jupyterlab-desktop --no-sandbox --disable-gpu
WSLFIX
