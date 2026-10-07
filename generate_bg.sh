#!/bin/bash
OUTPUT="/root/ransom_bg.png"
echo "[*] Generation image fond: $OUTPUT"
if command -v convert &>/dev/null; then
    convert -size 1920x1080 xc:'#0a0000' -fill '#ff0000' -gravity North -pointsize 52 -annotate +0+80 'VOS FICHIERS SONT CHIFFRES' -fill '#cc0000' -gravity North -pointsize 28 -annotate +0+160 'AES-256-GCM ENCRYPTION ACTIVE' -fill '#ff4444' -gravity Center -pointsize 22 -annotate +0+0 'Ceci est une demo securite' -fill '#666666' -gravity South -pointsize 16 -annotate +0+30 'DemoShield V2 --decrypt' "$OUTPUT" && echo "[OK]" || echo "[FAIL]"
else echo "[!] Installez imagemagick"; fi
if [ -f "$OUTPUT" ]; then for d in /home/*/; do cp "$OUTPUT" "${d}ransom_bg.png" 2>/dev/null; done; fi
