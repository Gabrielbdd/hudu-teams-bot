#!/usr/bin/env bash
set -euo pipefail

# Empacota o manifest do Teams em um ZIP pronto para sideload.
# Uso: ./scripts/package-manifest.sh <TEAMS_APP_ID>

APP_ID="${1:?Uso: $0 <TEAMS_APP_ID>}"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
MANIFEST_DIR="$PROJECT_DIR/manifest"
OUT_DIR="$PROJECT_DIR/dist"

mkdir -p "$OUT_DIR"

# Gera manifest com o APP_ID substituído
sed "s/{{TEAMS_APP_ID}}/$APP_ID/g" "$MANIFEST_DIR/manifest.json" > "$OUT_DIR/manifest.json"

# Copia ícones
cp "$MANIFEST_DIR/color.png" "$OUT_DIR/color.png"
cp "$MANIFEST_DIR/outline.png" "$OUT_DIR/outline.png"

# Cria o ZIP
cd "$OUT_DIR"
zip -j "$OUT_DIR/hudu-bot.zip" manifest.json color.png outline.png

# Limpa arquivos temporários
rm -f "$OUT_DIR/manifest.json" "$OUT_DIR/color.png" "$OUT_DIR/outline.png"

echo "Pacote criado em: $OUT_DIR/hudu-bot.zip"
