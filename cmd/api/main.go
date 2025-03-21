package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cetinibs/online-speed-test-backend-root/internal/controllers"
	"github.com/cetinibs/online-speed-test-backend-root/internal/models"
	"github.com/cetinibs/online-speed-test-backend-root/internal/repositories"
	"github.com/cetinibs/online-speed-test-backend-root/internal/services"
)

// Basit bir HTML içeriği
const htmlContent = `
<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Online Speed Test</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css">
    <style>
        :root {
            --primary-color: #00adb5;
            --primary-hover: #00969e;
            --download-color: #00adb5;
            --upload-color: #9c27b0;
            --ping-color: #ffeb3b;
            --dark-bg: #1a1e2e;
            --card-bg: #242938;
            --text-color: #ffffff;
            --secondary-text: #a0a0a0;
        }
        
        body {
            font-family: 'Segoe UI', Arial, sans-serif;
            background-color: var(--dark-bg);
            color: var(--text-color);
            margin: 0;
            padding: 0;
            display: flex;
            flex-direction: column;
            min-height: 100vh;
            justify-content: center;
            align-items: center;
        }
        
        .container {
            width: 100%;
            max-width: 800px;
            padding: 20px;
        }
        
        .header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            width: 100%;
            padding: 10px 0;
        }
        
        .logo h1 {
            color: var(--text-color);
            margin: 0;
            font-size: 24px;
        }
        
        .nav-buttons {
            display: flex;
            gap: 15px;
        }
        
        .nav-button {
            background: none;
            border: none;
            color: var(--text-color);
            cursor: pointer;
            display: flex;
            align-items: center;
            gap: 5px;
            font-size: 14px;
        }
        
        .nav-button i {
            font-size: 18px;
        }
        
        .result-card {
            background-color: var(--card-bg);
            border-radius: 12px;
            padding: 30px;
            margin: 20px 0;
            box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
        }
        
        .speed-display {
            display: flex;
            justify-content: space-around;
            margin-bottom: 30px;
        }
        
        .speed-item {
            text-align: center;
            width: 45%;
        }
        
        .speed-label {
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 8px;
            color: var(--secondary-text);
            margin-bottom: 10px;
            font-size: 16px;
        }
        
        .download-label i {
            color: var(--download-color);
        }
        
        .upload-label i {
            color: var(--upload-color);
        }
        
        .speed-value {
            font-size: 60px;
            font-weight: 300;
            margin: 0;
        }
        
        .speed-unit {
            color: var(--secondary-text);
            font-size: 16px;
            margin-top: 5px;
        }
        
        .ping-info {
            display: flex;
            justify-content: center;
            align-items: center;
            gap: 20px;
            margin: 20px 0;
            color: var(--secondary-text);
        }
        
        .ping-label {
            display: flex;
            align-items: center;
            gap: 8px;
        }
        
        .ping-value {
            display: flex;
            align-items: center;
            gap: 5px;
        }
        
        .ping-value i {
            color: var(--ping-color);
        }
        
        .server-info {
            display: flex;
            justify-content: space-between;
            margin-top: 30px;
            color: var(--secondary-text);
            font-size: 14px;
        }
        
        .server-item {
            display: flex;
            align-items: center;
            gap: 10px;
        }
        
        .server-icon {
            width: 36px;
            height: 36px;
            background-color: rgba(255, 255, 255, 0.1);
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        
        .server-details {
            display: flex;
            flex-direction: column;
        }
        
        .server-name {
            color: var(--text-color);
        }
        
        .server-location {
            font-size: 12px;
        }
        
        .change-server {
            color: var(--primary-color);
            cursor: pointer;
            font-size: 12px;
            margin-top: 3px;
        }
        
        .connection-type {
            text-align: center;
            margin-top: 20px;
        }
        
        .connection-label {
            color: var(--secondary-text);
            font-size: 14px;
            margin-bottom: 5px;
        }
        
        .connection-options {
            display: flex;
            justify-content: center;
            gap: 10px;
        }
        
        .connection-option {
            padding: 5px 15px;
            background-color: rgba(255, 255, 255, 0.05);
            border-radius: 20px;
            cursor: pointer;
        }
        
        .connection-option.active {
            background-color: rgba(0, 173, 181, 0.2);
            color: var(--primary-color);
        }
        
        .speedometer-container {
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            margin: 40px 0;
            position: relative;
        }
        
        .speedometer {
            width: 300px;
            height: 300px;
            position: relative;
        }
        
        .speedometer-outer {
            width: 100%;
            height: 100%;
            border-radius: 50%;
            background: conic-gradient(
                var(--download-color) 0% 30%, 
                #2c3e50 30% 100%
            );
            transform: rotate(-90deg);
            position: relative;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        
        .speedometer-inner {
            width: 85%;
            height: 85%;
            border-radius: 50%;
            background-color: var(--card-bg);
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            transform: rotate(90deg);
        }
        
        .speed-reading {
            font-size: 48px;
            font-weight: 300;
            margin: 0;
        }
        
        .speed-unit-small {
            font-size: 14px;
            color: var(--secondary-text);
        }
        
        .speed-marks {
            position: absolute;
            width: 100%;
            height: 100%;
            top: 0;
            left: 0;
            z-index: 1;
        }
        
        .speed-mark {
            position: absolute;
            color: var(--secondary-text);
            font-size: 12px;
        }
        
        .mark-0 {
            bottom: 10px;
            left: 10px;
        }
        
        .mark-250 {
            top: 10px;
            right: 45%;
        }
        
        .mark-500 {
            top: 10px;
            right: 10px;
        }
        
        .mark-750 {
            bottom: 40%;
            right: 10px;
        }
        
        .mark-1000 {
            bottom: 10px;
            right: 10px;
        }
        
        .test-button {
            background-color: var(--primary-color);
            color: white;
            border: none;
            border-radius: 30px;
            padding: 15px 40px;
            font-size: 18px;
            cursor: pointer;
            margin-top: 20px;
            transition: background-color 0.3s;
        }
        
        .test-button:hover {
            background-color: var(--primary-hover);
        }
        
        .test-button:disabled {
            background-color: #555;
            cursor: not-allowed;
        }
        
        .result-id {
            text-align: center;
            color: var(--secondary-text);
            font-size: 14px;
            margin-bottom: 20px;
        }
        
        .share-buttons {
            display: flex;
            justify-content: center;
            gap: 10px;
        }
        
        .share-button {
            width: 36px;
            height: 36px;
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            background-color: rgba(255, 255, 255, 0.1);
            color: var(--text-color);
            cursor: pointer;
            transition: background-color 0.3s;
        }
        
        .share-button:hover {
            background-color: var(--primary-color);
        }
        
        .footer {
            text-align: center;
            margin-top: 40px;
            padding: 20px;
            color: var(--secondary-text);
            font-size: 12px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <div class="logo">
                <h1>Online Speed Test</h1>
            </div>
            <div class="nav-buttons">
                <button class="nav-button" onclick="showResults()">
                    <i class="fas fa-chart-bar"></i> RESULTS
                </button>
                <button class="nav-button" onclick="showSettings()">
                    <i class="fas fa-cog"></i> SETTINGS
                </button>
            </div>
        </div>
        
        <div id="test-page">
            <div class="speedometer-container">
                <div class="speedometer">
                    <div class="speedometer-outer">
                        <div class="speedometer-inner">
                            <div id="go-text" class="speed-reading">GO</div>
                            <div id="speed-value" class="speed-reading" style="display:none">0.00</div>
                            <div id="speed-unit" class="speed-unit-small" style="display:none">Mbps</div>
                        </div>
                    </div>
                    <div class="speed-marks">
                        <div class="speed-mark mark-0">0</div>
                        <div class="speed-mark mark-250">250</div>
                        <div class="speed-mark mark-500">500</div>
                        <div class="speed-mark mark-750">750</div>
                        <div class="speed-mark mark-1000">1000</div>
                    </div>
                </div>
                
                <div class="ping-info" id="ping-display" style="display:none">
                    <div class="ping-label">Ping ms</div>
                    <div class="ping-value"><i class="fas fa-bolt"></i> <span id="ping-value">--</span></div>
                    <div class="download-value"><i class="fas fa-arrow-down" style="color:var(--download-color)"></i> <span id="download-indicator">--</span></div>
                    <div class="upload-value"><i class="fas fa-arrow-up" style="color:var(--upload-color)"></i> <span id="upload-indicator">--</span></div>
                </div>
                
                <button id="test-button" class="test-button" onclick="startTest()">Start Speed Test</button>
            </div>
            
            <div class="server-info">
                <div class="server-item">
                    <div class="server-icon">
                        <i class="fas fa-user"></i>
                    </div>
                    <div class="server-details">
                        <div class="server-name">Turksat</div>
                        <div class="server-location">94.55.140.138</div>
                    </div>
                </div>
                
                <div class="server-item">
                    <div class="server-icon">
                        <i class="fas fa-globe"></i>
                    </div>
                    <div class="server-details">
                        <div class="server-name">Turkcell</div>
                        <div class="server-location">Ankara</div>
                        <div class="change-server" onclick="changeServer()">Change Server</div>
                    </div>
                </div>
            </div>
            
            <div class="connection-type">
                <div class="connection-label">Connections</div>
                <div class="connection-options">
                    <div class="connection-option active" onclick="setConnectionType('multi')">Multi</div>
                    <div class="connection-option" onclick="setConnectionType('single')">Single</div>
                </div>
            </div>
        </div>
        
        <div id="results-page" style="display:none">
            <div class="result-id">
                Result ID <span id="result-id">17515482928</span>
            </div>
            
            <div class="result-card">
                <div class="speed-display">
                    <div class="speed-item">
                        <div class="speed-label download-label">
                            <i class="fas fa-arrow-down"></i> DOWNLOAD Mbps
                        </div>
                        <div id="download-result" class="speed-value">95.21</div>
                    </div>
                    
                    <div class="speed-item">
                        <div class="speed-label upload-label">
                            <i class="fas fa-arrow-up"></i> UPLOAD Mbps
                        </div>
                        <div id="upload-result" class="speed-value">9.67</div>
                    </div>
                </div>
                
                <div class="ping-info">
                    <div class="ping-label">Ping ms</div>
                    <div class="ping-value"><i class="fas fa-bolt"></i> <span id="ping-result">16</span></div>
                    <div class="download-value"><i class="fas fa-arrow-down" style="color:var(--download-color)"></i> <span>53</span></div>
                    <div class="upload-value"><i class="fas fa-arrow-up" style="color:var(--upload-color)"></i> <span>64</span></div>
                </div>
                
                <div class="server-info">
                    <div class="server-item">
                        <div class="server-icon">
                            <i class="fas fa-server"></i>
                        </div>
                        <div class="server-details">
                            <div class="connection-label">Connections</div>
                            <div class="server-name">Multi</div>
                        </div>
                    </div>
                    
                    <div class="server-item">
                        <div class="server-icon">
                            <i class="fas fa-globe"></i>
                        </div>
                        <div class="server-details">
                            <div class="server-name">Turkcell</div>
                            <div class="server-location">Ankara</div>
                            <div class="change-server">Change Server</div>
                        </div>
                    </div>
                    
                    <div class="server-item">
                        <div class="server-icon">
                            <i class="fas fa-user"></i>
                        </div>
                        <div class="server-details">
                            <div class="server-name">Turksat</div>
                            <div class="server-location">94.55.140.138</div>
                        </div>
                    </div>
                </div>
                
                <div style="text-align: center; margin-top: 30px;">
                    <div style="margin-bottom: 10px; color: var(--secondary-text);">RATE YOUR PROVIDER</div>
                    <div>Turksat</div>
                    <div style="margin-top: 10px;">
                        <i class="far fa-star" style="color: var(--secondary-text); margin: 0 5px; cursor: pointer;" onclick="rateProvider(1)"></i>
                        <i class="far fa-star" style="color: var(--secondary-text); margin: 0 5px; cursor: pointer;" onclick="rateProvider(2)"></i>
                        <i class="far fa-star" style="color: var(--secondary-text); margin: 0 5px; cursor: pointer;" onclick="rateProvider(3)"></i>
                        <i class="far fa-star" style="color: var(--secondary-text); margin: 0 5px; cursor: pointer;" onclick="rateProvider(4)"></i>
                        <i class="far fa-star" style="color: var(--secondary-text); margin: 0 5px; cursor: pointer;" onclick="rateProvider(5)"></i>
                    </div>
                </div>
            </div>
            
            <div style="text-align: center; margin-top: 20px;">
                <div style="margin-bottom: 10px;">SHARE</div>
                <div class="share-buttons">
                    <div class="share-button" onclick="shareResult('link')"><i class="fas fa-link"></i></div>
                    <div class="share-button" onclick="shareResult('twitter')"><i class="fab fa-twitter"></i></div>
                    <div class="share-button" onclick="shareResult('facebook')"><i class="fab fa-facebook-f"></i></div>
                    <div class="share-button" onclick="shareResult('more')"><i class="fas fa-ellipsis-h"></i></div>
                </div>
            </div>
            
            <div style="text-align: center; margin-top: 30px;">
                <button class="test-button" onclick="showTestPage()">Test Again</button>
            </div>
        </div>
    </div>
    
    <div class="footer">
        <p> 2025 Online Speed Test. Tüm hakları saklıdır.</p>
    </div>
    
    <script>
        // Current page
        let currentPage = 'test';
        
        // Show test page
        function showTestPage() {
            document.getElementById('test-page').style.display = 'block';
            document.getElementById('results-page').style.display = 'none';
            currentPage = 'test';
            resetTest();
        }
        
        // Show results page
        function showResults() {
            document.getElementById('test-page').style.display = 'none';
            document.getElementById('results-page').style.display = 'block';
            currentPage = 'results';
        }
        
        // Show settings page
        function showSettings() {
            alert('Settings page will be implemented in the future.');
        }
        
        // Change server
        function changeServer() {
            alert('Server change functionality will be implemented in the future.');
        }
        
        // Set connection type
        function setConnectionType(type) {
            const options = document.querySelectorAll('.connection-option');
            options.forEach(option => {
                option.classList.remove('active');
            });
            
            if (type === 'multi') {
                options[0].classList.add('active');
            } else {
                options[1].classList.add('active');
            }
        }
        
        // Rate provider
        function rateProvider(rating) {
            const stars = document.querySelectorAll('.fa-star');
            stars.forEach((star, index) => {
                if (index < rating) {
                    star.classList.remove('far');
                    star.classList.add('fas');
                    star.style.color = '#ffeb3b';
                } else {
                    star.classList.remove('fas');
                    star.classList.add('far');
                    star.style.color = 'var(--secondary-text)';
                }
            });
        }
        
        // Share result
        function shareResult(platform) {
            alert('Sharing via ' + platform + ' will be implemented in the future.');
        }
        
        // Start speed test
        function startTest() {
            const testButton = document.getElementById('test-button');
            const goText = document.getElementById('go-text');
            const speedValue = document.getElementById('speed-value');
            const speedUnit = document.getElementById('speed-unit');
            const pingDisplay = document.getElementById('ping-display');
            
            testButton.disabled = true;
            testButton.textContent = 'Testing...';
            
            // Hide GO text, show speed value
            goText.style.display = 'none';
            speedValue.style.display = 'block';
            speedUnit.style.display = 'block';
            pingDisplay.style.display = 'flex';
            
            // Simulate ping test
            setTimeout(() => {
                document.getElementById('ping-value').textContent = '18';
                simulateSpeedometer(0, 97, 'download');
            }, 500);
        }
        
        // Simulate speedometer animation
        function simulateSpeedometer(from, to, type) {
            const speedValue = document.getElementById('speed-value');
            const speedometer = document.querySelector('.speedometer-outer');
            const duration = 3000; // 3 seconds
            const steps = 60;
            const increment = (to - from) / steps;
            let current = from;
            let step = 0;
            
            // Update color based on type
            if (type === 'download') {
                speedometer.style.background = "conic-gradient(var(--download-color) 0% " + 30 + "deg, #2c3e50 " + 30 + "deg 100%)";
                document.getElementById('speed-unit').textContent = 'Mbps';
                document.getElementById('download-indicator').textContent = '57';
                document.getElementById('upload-indicator').textContent = '—';
            } else {
                speedometer.style.background = "conic-gradient(var(--upload-color) 0% " + 30 + "deg, #2c3e50 " + 30 + "deg 100%)";
                document.getElementById('speed-unit').textContent = 'Mbps';
                document.getElementById('upload-indicator').textContent = '64';
            }
            
            const interval = setInterval(() => {
                current += increment;
                step++;
                
                // Update speedometer value
                speedValue.textContent = current.toFixed(2);
                
                // Update speedometer fill
                const percentage = (current / 1000) * 100;
                const degrees = percentage * 3.6 * 0.8; // 80% of the circle
                speedometer.style.background = type === 'download' 
                    ? "conic-gradient(var(--download-color) 0% " + degrees + "deg, #2c3e50 " + degrees + "deg 100%)"
                    : "conic-gradient(var(--upload-color) 0% " + degrees + "deg, #2c3e50 " + degrees + "deg 100%)";
                
                if (step >= steps) {
                    clearInterval(interval);
                    
                    if (type === 'download') {
                        // After download test, start upload test
                        setTimeout(() => {
                            simulateSpeedometer(0, 9.67, 'upload');
                        }, 500);
                    } else {
                        // After upload test, show results
                        setTimeout(() => {
                            showResults();
                            resetTest();
                        }, 1000);
                    }
                }
            }, duration / steps);
        }
        
        // Reset test UI
        function resetTest() {
            const testButton = document.getElementById('test-button');
            const goText = document.getElementById('go-text');
            const speedValue = document.getElementById('speed-value');
            const speedUnit = document.getElementById('speed-unit');
            const pingDisplay = document.getElementById('ping-display');
            
            testButton.disabled = false;
            testButton.textContent = 'Start Speed Test';
            
            goText.style.display = 'block';
            speedValue.style.display = 'none';
            speedUnit.style.display = 'none';
            pingDisplay.style.display = 'none';
            
            document.querySelector('.speedometer-outer').style.background = "conic-gradient(var(--download-color) 0% 30%, #2c3e50 30% 100%)";
        }
        
        // Initialize
        document.addEventListener('DOMContentLoaded', function() {
            showTestPage();
        });
    </script>
</body>
</html>
`

func main() {
	// Create repository instances
	speedTestRepo := createInMemorySpeedTestRepo()
	userRepo := createInMemoryUserRepo()

	// Create service instances
	speedTestService := services.NewSpeedTestService(speedTestRepo, userRepo)

	// Create controller instances
	speedTestController := controllers.NewSpeedTestController(speedTestService)

	// Set up HTTP server
	mux := http.NewServeMux()

	// Define API routes
	mux.HandleFunc("/api/speedtest", speedTestController.RunTest)

	// Serve HTML content directly
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, htmlContent)
	})

	// Start the server
	port := 9090 // Farklı bir port kullanıyoruz
	fmt.Printf("Starting server on port %d...\n", port)
	fmt.Printf("Server is running at http://localhost:%d\n", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// createInMemorySpeedTestRepo creates an in-memory implementation of SpeedTestRepository
func createInMemorySpeedTestRepo() repositories.SpeedTestRepository {
	return &InMemorySpeedTestRepo{results: make(map[string]*models.SpeedTestResult)}
}

// createInMemoryUserRepo creates an in-memory implementation of UserRepository
func createInMemoryUserRepo() repositories.UserRepository {
	return &InMemoryUserRepo{users: make(map[string]*models.UserProfile)}
}

// InMemorySpeedTestRepo is an in-memory implementation of SpeedTestRepository
type InMemorySpeedTestRepo struct {
	results map[string]*models.SpeedTestResult
}

func (r *InMemorySpeedTestRepo) SaveResult(ctx context.Context, result *models.SpeedTestResult) error {
	r.results[result.ID] = result
	return nil
}

func (r *InMemorySpeedTestRepo) GetResultsByUserID(ctx context.Context, userID string) ([]*models.SpeedTestResult, error) {
	var userResults []*models.SpeedTestResult
	for _, result := range r.results {
		if result.UserID == userID {
			userResults = append(userResults, result)
		}
	}
	return userResults, nil
}

func (r *InMemorySpeedTestRepo) GetResultByID(ctx context.Context, id string) (*models.SpeedTestResult, error) {
	result, ok := r.results[id]
	if !ok {
		return nil, fmt.Errorf("result not found")
	}
	return result, nil
}

func (r *InMemorySpeedTestRepo) DeleteResult(ctx context.Context, id string) error {
	delete(r.results, id)
	return nil
}

// InMemoryUserRepo is an in-memory implementation of UserRepository
type InMemoryUserRepo struct {
	users map[string]*models.UserProfile
}

func (r *InMemoryUserRepo) SaveUser(ctx context.Context, user *models.UserProfile) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) GetUserByID(ctx context.Context, id string) (*models.UserProfile, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepo) GetUserByEmail(ctx context.Context, email string) (*models.UserProfile, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}


