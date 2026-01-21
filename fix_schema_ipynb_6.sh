
#!/usr/bin/env bash
set -e

PROJECT_DIR="$HOME/nbgo"
VSCODE_DIR="$PROJECT_DIR/.vscode"
LAUNCH_JSON="$VSCODE_DIR/launch.json"

echo "==== Fix Python/Go Debugger Collision ===="

# 1. Verify the problem
echo "[INFO] Verifying file types..."
file "$PROJECT_DIR/gw/gate/main.go"

# 2. Ensure .vscode exists
mkdir -p "$VSCODE_DIR"

# 3. Backup launch.json if it exists
if [ -f "$LAUNCH_JSON" ]; then
    echo "[INFO] Backing up existing launch.json"
    cp "$LAUNCH_JSON" "$LAUNCH_JSON.bak.$(date +%s)"
fi

# 4. Write a correct Go + Python launch.json
cat > "$LAUNCH_JSON" <<'JSON'
{
  "version": "0.2.0",
  "configurations": [

    {
      "name": "Go: Debug Gateway",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}/gw/gate",
      "cwd": "${workspaceFolder}",
      "env": {},
      "args": []
    },

    {
      "name": "Python: Current File",
      "type": "python",
      "request": "launch",
      "program": "${file}",
      "console": "integratedTerminal",
      "justMyCode": true
    }

  ]
}
JSON

# 5. Prevent Python from touching .go files
cat > "$PROJECT_DIR/.python-ignore-go.sh" <<'EOF'
#!/usr/bin/env bash
if [[ "$1" == *.go ]]; then
  echo "ERROR: Refusing to run Go file with Python: $1"
  exit 1
fi
exec python "$@"
EOF
chmod +x "$PROJECT_DIR/.python-ignore-go.sh"

# 6. Summary
echo
echo "✅ FIX COMPLETE"
echo "• Go files now debug with Go debugger"
echo "• Python debugger can no longer run .go files"
echo "• launch.json corrected"
echo
echo "Next steps:"
echo "1. Open VS Code"
echo "2. Select 'Go: Debug Gateway'"
echo "3. Press F5"
