package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sort"
	"sync"
	"time"
)

// Цвета для вывода
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// Результат теста
type BenchmarkResult struct {
	TestName     string
	Duration     time.Duration
	Operations   int64
	OpsPerSecond float64
	Score        int
}

// Главная структура бенчмарка
type LaptopBenchmark struct {
	results      []BenchmarkResult
	totalScore   int
	cpuCores     int
	startTime    time.Time
	mu           sync.Mutex
	progressChan chan string
}

func NewLaptopBenchmark() *LaptopBenchmark {
	return &LaptopBenchmark{
		results:      make([]BenchmarkResult, 0),
		cpuCores:     runtime.NumCPU(),
		progressChan: make(chan string, 100),
	}
}

// ============================================
// ТЕСТ 1: CPU - Вычисление простых чисел
// ============================================
func (lb *LaptopBenchmark) testPrimeNumbers() BenchmarkResult {
	lb.progressChan <- "🔢 Тест 1/10: Вычисление простых чисел..."

	start := time.Now()
	limit := 1000000
	primeCount := 0

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	for i := 2; i < limit; i++ {
		if isPrime(i) {
			primeCount++
		}
	}

	duration := time.Since(start)
	opsPerSec := float64(limit) / duration.Seconds()
	score := int(opsPerSec / 10000)

	return BenchmarkResult{
		TestName:     "Вычисление простых чисел",
		Duration:     duration,
		Operations:   int64(limit),
		OpsPerSecond: opsPerSec,
		Score:        score,
	}
}

// ============================================
// ТЕСТ 2: CPU - Математические вычисления
// ============================================
func (lb *LaptopBenchmark) testMathOperations() BenchmarkResult {
	lb.progressChan <- "➗ Тест 2/10: Математические операции..."

	start := time.Now()
	iterations := 10000000
	result := 0.0

	for i := 0; i < iterations; i++ {
		x := float64(i)
		result += math.Sqrt(x) * math.Sin(x) / math.Cos(x+1)
		result += math.Log(x+1) * math.Exp(x/1000000)
		result += math.Pow(x, 1.5) / (x + 1)
	}

	duration := time.Since(start)
	opsPerSec := float64(iterations) / duration.Seconds()
	score := int(opsPerSec / 100000)

	return BenchmarkResult{
		TestName:     "Математические вычисления",
		Duration:     duration,
		Operations:   int64(iterations),
		OpsPerSecond: opsPerSec,
		Score:        score,
	}
}

// ============================================
// ТЕСТ 3: МНОГОПОТОЧНОСТЬ - Параллельные вычисления
// ============================================
func (lb *LaptopBenchmark) testMultiThreading() BenchmarkResult {
	lb.progressChan <- "🔄 Тест 3/10: Многопоточные вычисления..."

	start := time.Now()
	numWorkers := lb.cpuCores * 2
	jobsPerWorker := 500000

	var wg sync.WaitGroup
	totalOps := int64(0)
	var mu sync.Mutex

	worker := func(id int) {
		defer wg.Done()
		localOps := int64(0)

		for i := 0; i < jobsPerWorker; i++ {
			x := float64(i)
			_ = math.Sqrt(x) * math.Sin(x) * math.Cos(x)
			localOps++
		}

		mu.Lock()
		totalOps += localOps
		mu.Unlock()
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()
	duration := time.Since(start)
	opsPerSec := float64(totalOps) / duration.Seconds()
	score := int(opsPerSec / 500000)

	return BenchmarkResult{
		TestName:     "Многопоточные вычисления",
		Duration:     duration,
		Operations:   totalOps,
		OpsPerSecond: opsPerSec,
		Score:        score,
	}
}

// ============================================
// ТЕСТ 4: ПАМЯТЬ - Работа с большими массивами
// ============================================
func (lb *LaptopBenchmark) testMemoryOperations() BenchmarkResult {
	lb.progressChan <- "💾 Тест 4/10: Операции с памятью..."

	start := time.Now()
	size := 10000000
	data := make([]int, size)

	// Заполнение
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000000)
	}

	// Сортировка
	sort.Ints(data)

	// Поиск
	searchOps := 0
	for i := 0; i < 10000; i++ {
		target := rand.Intn(1000000)
		idx := sort.SearchInts(data, target)
		if idx < len(data) {
			searchOps++
		}
	}

	duration := time.Since(start)
	opsPerSec := float64(size+searchOps) / duration.Seconds()
	score := int(opsPerSec / 100000)

	return BenchmarkResult{
		TestName:     "Операции с памятью",
		Duration:     duration,
		Operations:   int64(size + searchOps),
		OpsPerSecond: opsPerSec,
		Score:        score,
	}
}

// ============================================
// ТЕСТ 5: СТРОКИ - Обработка текста
// ============================================
func (lb *LaptopBenchmark) testStringOperations() BenchmarkResult {
	lb.progressChan <- "📝 Тест 5/10: Обработка строк..."

	start := time.Now()
	iterations := 100000
	totalOps := int64(0)

	baseString := "Это тестовая строка для проверки производительности обработки текста в Go"

	for i := 0; i < iterations; i++ {
		// Конкатенация
		result := baseString
		for j := 0; j < 10; j++ {
			result += fmt.Sprintf(" %d", j)
			totalOps++
		}

		// Поиск подстроки
		for j := 0; j < 5; j++ {
			_ = len(result)
			totalOps++
		}
	}

	duration := time.Since(start)
	opsPerSec := float64(totalOps) / duration.Seconds()
	score := int(opsPerSec / 50000)

	return BenchmarkResult{
		TestName:     "Обработка строк",
		Duration:     duration,
		Operations:   totalOps,
		OpsPerSecond: opsPerSec,
		Score:        score,
	}
}

// ============================================
// ТЕСТ 6: КРИПТОГРАФИЯ - Хеширование
// ============================================
func (lb *LaptopBenchmark) testCryptoOperations() BenchmarkResult {
	lb.progressChan <- "🔐 Тест 6/10: Криптографические операции..."

	start := time.Now()
	iterations := 100000
	data := []byte("Тестовые данные для хеширования и проверки криптографической производительности системы")

	for i := 0; i < iterations; i++ {
		// MD5
		_ = md5.Sum(data)

		// SHA256
		_ = sha256.Sum256(data)
	}

	duration := time.Since(start)
	totalOps := int64(iterations * 2)
	opsPerSec := float64(totalOps) / duration.Seconds()
	score := int(opsPerSec / 10000)

	return BenchmarkResult{
		TestName:     "Криптографические операции",
		Duration:     duration,
		Operations:   totalOps,
		OpsPerSecond: opsPerSec,
		Score:        score,
	}
}

// ============================================
// ТЕСТ 7: JSON - Сериализация/Десериализация
// ============================================
func (lb *LaptopBenchmark) testJSONOperations() BenchmarkResult {
	lb.progressChan <- "📦 Тест 7/10: Операции с JSON..."

	start := time.Now()
	iterations := 50000

	type TestData struct {
		ID       int                    `json:"id"`
		Name     string                 `json:"name"`
		Email    string                 `json:"email"`
		Age      int                    `json:"age"`
		Active   bool                   `json:"active"`
		Tags     []string               `json:"tags"`
		Metadata map[string]interface{} `json:"metadata"`
	}

	for i := 0; i < iterations; i++ {
		data := TestData{
			ID:     i,
			Name:   fmt.Sprintf("User_%d", i),
			Email:  fmt.Sprintf("user%d@example.com", i),
			Age:    20 + (i % 50),
			Active: i%2 == 0,
			Tags:   []string{"tag1", "tag2", "tag3"},
			Metadata: map[string]interface{}{
				"key1": "value1",
				"key2": 123,
				"key3": true,
			},
		}

		// Сериализация
		jsonData, _ := json.Marshal(data)

		// Десериализация
		var decoded TestData
		_ = json.Unmarshal(jsonData, &decoded)
	}

	duration := time.Since(start)
	totalOps := int64(iterations * 2)
	opsPerSec := float64(totalOps) / duration.Seconds()
	score := int(opsPerSec / 5000)

	return BenchmarkResult{
		TestName:     "Операции с JSON",
		Duration:     duration,
		Operations:   totalOps,
		OpsPerSecond: opsPerSec,
		Score:        score,
	}
}

// ============================================
// ТЕСТ 8: АЛГОРИТМЫ - Сортировка и поиск
// ============================================
func (lb *LaptopBenchmark) testAlgorithms() BenchmarkResult {
	lb.progressChan <- "🎯 Тест 8/10: Алгоритмы сортировки и поиска..."

	start := time.Now()
	arraySize := 100000
	iterations := 100
	totalOps := int64(0)

	for iter := 0; iter < iterations; iter++ {
		// Создаем случайный массив
		data := make([]int, arraySize)
		for i := 0; i < arraySize; i++ {
			data[i] = rand.Intn(1000000)
		}

		// Сортировка
		sort.Ints(data)
		totalOps += int64(arraySize)

		// Бинарный поиск
		for i := 0; i < 1000; i++ {
			target := rand.Intn(1000000)
			_ = sort.SearchInts(data, target)
			totalOps++
		}
	}

	duration := time.Since(start)
	opsPerSec := float64(totalOps) / duration.Seconds()
	score := int(opsPerSec / 100000)

	return BenchmarkResult{
		TestName:     "Алгоритмы (сортировка/поиск)",
		Duration:     duration,
		Operations:   totalOps,
		OpsPerSecond: opsPerSec,
		Score:        score,
	}
}

// ============================================
// ТЕСТ 9: РЕКУРСИЯ - Числа Фибоначчи
// ============================================
func (lb *LaptopBenchmark) testRecursion() BenchmarkResult {
	lb.progressChan <- "🔁 Тест 9/10: Рекурсивные вычисления..."

	start := time.Now()

	var fib func(n int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	// Вычисляем числа Фибоначчи с кешированием
	cache := make(map[int]int)
	var fibCached func(n int) int
	fibCached = func(n int) int {
		if n <= 1 {
			return n
		}
		if val, exists := cache[n]; exists {
			return val
		}
		result := fibCached(n-1) + fibCached(n-2)
		cache[n] = result
		return result
	}

	totalOps := int64(0)
	for i := 0; i < 35; i++ {
		_ = fibCached(i)
		totalOps++
	}

	// Дополнительные рекурсивные вычисления
	for i := 0; i < 1000; i++ {
		_ = fibCached(30)
		totalOps++
	}

	duration := time.Since(start)
	opsPerSec := float64(totalOps) / duration.Seconds()
	score := int(opsPerSec / 50)

	return BenchmarkResult{
		TestName:     "Рекурсивные вычисления",
		Duration:     duration,
		Operations:   totalOps,
		OpsPerSecond: opsPerSec,
		Score:        score,
	}
}

// ============================================
// ТЕСТ 10: КОМПЛЕКСНЫЙ - Симуляция реальной нагрузки
// ============================================
func (lb *LaptopBenchmark) testComplexWorkload() BenchmarkResult {
	lb.progressChan <- "⚡ Тест 10/10: Комплексная нагрузка..."

	start := time.Now()
	numWorkers := lb.cpuCores
	var wg sync.WaitGroup
	totalOps := int64(0)
	var mu sync.Mutex

	worker := func(id int) {
		defer wg.Done()
		localOps := int64(0)

		// Математика
		for i := 0; i < 10000; i++ {
			x := float64(i)
			_ = math.Sqrt(x) * math.Sin(x)
			localOps++
		}

		// Работа с данными
		data := make([]int, 10000)
		for i := range data {
			data[i] = rand.Intn(10000)
		}
		sort.Ints(data)
		localOps += int64(len(data))

		// Хеширование
		for i := 0; i < 1000; i++ {
			testData := []byte(fmt.Sprintf("worker_%d_data_%d", id, i))
			_ = sha256.Sum256(testData)
			localOps++
		}

		mu.Lock()
		totalOps += localOps
		mu.Unlock()
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()
	duration := time.Since(start)
	opsPerSec := float64(totalOps) / duration.Seconds()
	score := int(opsPerSec / 10000)

	return BenchmarkResult{
		TestName:     "Комплексная нагрузка",
		Duration:     duration,
		Operations:   totalOps,
		OpsPerSecond: opsPerSec,
		Score:        score,
	}
}

// ============================================
// ЗАПУСК ВСЕХ ТЕСТОВ
// ============================================
func (lb *LaptopBenchmark) RunAllTests() {
	lb.startTime = time.Now()

	// Запускаем вывод прогресса в отдельной горутине
	go lb.printProgress()

	fmt.Printf("%s\n", colorCyan+"="*60+colorReset)
	fmt.Printf("%s🚀 ТЕСТ ПРОИЗВОДИТЕЛЬНОСТИ НОУТБУКА 🚀%s\n", colorYellow, colorReset)
	fmt.Printf("%s\n", colorCyan+"="*60+colorReset)
	fmt.Printf("💻 Процессор: %d ядер/потоков\n", lb.cpuCores)
	fmt.Printf("⏰ Начало тестирования: %s\n\n", time.Now().Format("15:04:05"))

	tests := []func() BenchmarkResult{
		lb.testPrimeNumbers,
		lb.testMathOperations,
		lb.testMultiThreading,
		lb.testMemoryOperations,
		lb.testStringOperations,
		lb.testCryptoOperations,
		lb.testJSONOperations,
		lb.testAlgorithms,
		lb.testRecursion,
		lb.testComplexWorkload,
	}

	for _, test := range tests {
		result := test()
		lb.mu.Lock()
		lb.results = append(lb.results, result)
		lb.totalScore += result.Score
		lb.mu.Unlock()
		time.Sleep(500 * time.Millisecond) // Небольшая пауза между тестами
	}

	close(lb.progressChan)
	time.Sleep(1 * time.Second)

	lb.printResults()
}

// Вывод прогресса
func (lb *LaptopBenchmark) printProgress() {
	for msg := range lb.progressChan {
		fmt.Printf("%s%s%s\n", colorCyan, msg, colorReset)
	}
}

// ============================================
// ВЫВОД РЕЗУЛЬТАТОВ
// ============================================
func (lb *LaptopBenchmark) printResults() {
	fmt.Printf("\n%s\n", colorGreen+"="*60+colorReset)
	fmt.Printf("%s📊 РЕЗУЛЬТАТЫ ТЕСТИРОВАНИЯ 📊%s\n", colorGreen, colorReset)
	fmt.Printf("%s\n\n", colorGreen+"="*60+colorReset)

	// Таблица результатов
	fmt.Printf("%-35s %10s %12s %8s\n", "ТЕСТ", "ВРЕМЯ", "ОПС/СЕК", "БАЛЛ")
	fmt.Printf("%s\n", colorCyan+"-"*70+colorReset)

	for _, result := range lb.results {
		color := colorWhite
		if result.Score >= 100 {
			color = colorGreen
		} else if result.Score >= 50 {
			color = colorYellow
		} else {
			color = colorRed
		}

		fmt.Printf("%-35s %10s %12.0f %s%8d%s\n",
			result.TestName,
			result.Duration.Round(time.Millisecond),
			result.OpsPerSecond,
			color,
			result.Score,
			colorReset,
		)
	}

	fmt.Printf("%s\n", colorCyan+"-"*70+colorReset)

	// Общий счет
	totalTime := time.Since(lb.startTime)
	avgScore := lb.totalScore / len(lb.results)

	fmt.Printf("\n%s📈 ИТОГОВАЯ СТАТИСТИКА:%s\n", colorPurple, colorReset)
	fmt.Printf("   Общее время тестирования: %s\n", totalTime.Round(time.Second))
	fmt.Printf("   Всего баллов: %s%d%s\n", colorYellow, lb.totalScore, colorReset)
	fmt.Printf("   Средний балл: %s%d%s\n", colorYellow, avgScore, colorReset)
	fmt.Printf("   Количество ядер: %d\n", lb.cpuCores)

	// Оценка производительности
	fmt.Printf("\n%s🏆 ИТОГОВАЯ ОЦЕНКА:%s ", colorYellow, colorReset)

	rating := ""
	ratingColor := colorWhite

	switch {
	case lb.totalScore >= 1000:
		rating = "ОТЛИЧНО! 🚀 Топовая производительность!"
		ratingColor = colorGreen
	case lb.totalScore >= 700:
		rating = "ХОРОШО! 💪 Высокая производительность"
		ratingColor = colorGreen
	case lb.totalScore >= 400:
		rating = "СРЕДНЕ ⚡ Нормальная производительность"
		ratingColor = colorYellow
	case lb.totalScore >= 200:
		rating = "НИЗКО ⚠️  Слабая производительность"
		ratingColor = colorRed
	default:
		rating = "ОЧЕНЬ НИЗКО 🐌 Очень слабая производительность"
		ratingColor = colorRed
	}

	fmt.Printf("%s%s%s\n", ratingColor, rating, colorReset)

	// Рекомендации
	fmt.Printf("\n%s💡 РЕКОМЕНДАЦИИ:%s\n", colorCyan, colorReset)
	if lb.totalScore < 400 {
		fmt.Println("   • Рассмотрите возможность апгрейда процессора")
		fmt.Println("   • Добавьте оперативной памяти")
		fmt.Println("   • Проверьте температуру и систему охлаждения")
		fmt.Println("   • Закройте фоновые приложения")
	} else if lb.totalScore < 700 {
		fmt.Println("   • Ваша система справляется с базовыми задачами")
		fmt.Println("   • Для улучшения производительности рассмотрите SSD")
		fmt.Println("   • Проверьте настройки энергопотребления")
	} else {
		fmt.Println("   • Отличная система для профессиональной работы!")
		fmt.Println("   • Подходит для разработки, 3D-моделирования")
		fmt.Println("   • Может работать с требовательными приложениями")
	}

	fmt.Printf("\n%s\n", colorCyan+"="*60+colorReset)
	fmt.Printf("✅ Тестирование завершено!\n")
	fmt.Printf("%s\n", colorCyan+"="*60+colorReset)
}

// ============================================
// ГЛАВНАЯ ФУНКЦИЯ
// ============================================
func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\n🔥 Программа нагрузит все ядра процессора!")
	fmt.Println("⚠️  Убедитесь, что ноутбук подключен к питанию")
	fmt.Println("⏱️  Тестирование займет около 2-3 минут\n")

	fmt.Print("Нажмите Enter для начала тестирования...")
	fmt.Scanln()

	// Устанавливаем максимальное использование CPU
	runtime.GOMAXPROCS(runtime.NumCPU())

	benchmark := NewLaptopBenchmark()
	benchmark.RunAllTests()

	fmt.Println("\n📊 Сохраните или сделайте скриншот результатов!")
	fmt.Println("🔄 Для повторного теста перезапустите программу\n")
}
