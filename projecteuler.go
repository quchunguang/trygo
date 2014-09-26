// from: https://projecteuler.net
package trygo

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

func PE1(N int64) int64 {
	var i int64
	var s int64 = 0
	for i = 3; i < N; i++ {
		if i%3 == 0 || i%5 == 0 {
			s += i
		}
	}
	return s
}

//////
func PE2(N int64) int64 {
	var a, b, s int64 = 1, 2, 0
	for b <= N {
		if b%2 == 0 {
			s += b
		}
		a, b = b, a+b
	}
	return s
}

//////
func reduceN(N *int64, p int64) {
	for (*N)%p == 0 {
		(*N) /= p
		// fmt.Printf("%d*", p)
	}
}
func PE3(N int64) int64 {
	// fmt.Printf("%d = ", N)
	var i int64
	var primes []int64
	primes = append(primes, 2)
	reduceN(&N, 2)

	for i = 3; i <= N; i += 2 {
		for _, p := range primes {
			if i%p == 0 {
				goto Out
			}
		}
		primes = append(primes, i)
		reduceN(&N, i)
	Out:
	}
	return primes[len(primes)-1]
}

//////
func palindrome(N int64) bool {
	var digit [6]int64
	var i int64
	for i = 0; i < 6; i++ {
		digit[i] = N % 10
		N /= 10
	}
	if digit[5] != digit[0] || digit[4] != digit[1] || digit[3] != digit[2] {
		return false
	}
	return true
}
func PE4() int64 {
	var a, b, N int64
	var s int64 = 0
	for a = 999; a >= 900; a-- {
		for b = 999; b >= 900; b-- {
			N = a * b
			if palindrome(N) {
				if N > s {
					s = N
				}
			}
		}
	}
	return s
}

//////
func power(a, b int64) int64 {
	var i int64
	var s int64 = 1
	for i = 0; i < b; i++ {
		s *= a
	}
	return s
}
func PE5(N int64) int64 {
	pruducts := make(map[int64]int64)
	pruducts[2] = 1
	var i, j, s, p int64
	for i = 3; i <= N; i++ {
		s = i
		for p = range pruducts {
			for j = 0; s%p == 0; j++ {
				s /= p
			}
			if pruducts[p] < j {
				pruducts[p] = j
			}
		}
		if s > 1 {
			pruducts[s] = 1
		}
	}
	// fmt.Println(pruducts)
	s = 1
	for i = range pruducts {
		s *= power(i, pruducts[i])
	}
	return s
}

//////
func PE6(N int64) int64 {
	var i, s1, s2 int64
	for i = 1; i <= N; i++ {
		s1 += i
		s2 += i * i
	}
	return s1*s1 - s2
}
func PE6b(N int64) int64 {
	var i, j, s int64
	for i = 1; i <= N; i++ {
		for j = 1; j <= N; j++ {
			if i != j {
				s += i * j
			}
		}
	}
	return s
}
func PE6c(N int64) int64 {
	return N * (N + 1) * (3*N*N - N - 2) / 12
}

//////
func PE7(N int64) int64 {
	var i int64 = 3
	var count int64 = 1
	var primes []int64
	primes = append(primes, 2)
	for i = 3; count < N; i++ {
		for _, p := range primes {
			if i%p == 0 {
				goto Out
			}
		}
		primes = append(primes, i)
		count++
	Out:
	}
	return primes[len(primes)-1]
}

//////
var data string = `73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450`

func PE8(N int) int {
	var i, j, s, val, max int
	data = strings.Replace(data, "\n", "", -1)
	for i = 0; i < len(data)-N+1; i++ {
		s = 1
		for j = 0; j < N; j++ {
			val = int(data[i+j]) - 48
			s *= val
		}
		if s > max {
			max = s
		}
	}
	return max
}

//////
func PE9(N int) int {
	var i, j, k int
	for i = 1; i < N/3; i++ {
		for j = i + 1; j < N/2; j++ {
			k = N - i - j
			if i*i+j*j == k*k {
				fmt.Println(i, j, k)
				return i * j * k
			}
		}
	}
	return 0
}

//////
const NMAX = 2e6

func PE10() int64 {
	var i, j, length, upbound, s int64
	var primes [NMAX / 10]int64
	primes[0] = 2
	primes[1] = 3
	length = 2
	for i = 5; i <= NMAX; i += 2 {
		upbound = int64(math.Sqrt(float64(i)))
		for j = 0; primes[j] <= upbound; j++ {
			// for j = 0; primes[j]*primes[j] <= i; j++ {
			if i%primes[j] == 0 {
				goto Out
			}
		}
		primes[length] = i
		length++
	Out:
	}
	s = 0
	for _, i := range primes {
		s += i
	}
	return s
}

////// Sieve of Eratosthenes
func PE10a() int64 {
	var i, j, total, s int64
	var flags [NMAX]bool
	total = int64(math.Sqrt(NMAX)) // put outside, 22ms->14ms !!!
	for i = 2; i < total; i++ {
		if flags[i] {
			continue
		}
		for j = 2; i*j < NMAX; j++ {
			flags[i*j] = true
		}
	}
	for i = 2; i < NMAX; i++ {
		if flags[i] == false {
			s += i
		}
	}
	return s
}

var flags [NMAX]bool

const CORES = 4

func worker(total int64, coreid int64) {
	var i, j int64
	for i = 2 + coreid; i < total; i += CORES {
		if flags[i] {
			continue
		}
		for j = 2; i*j < NMAX; j++ {
			flags[i*j] = true
		}
	}
}
func PE10b() int64 {
	var i, total, s int64
	var done = make(chan bool)

	runtime.GOMAXPROCS(CORES)

	total = int64(math.Sqrt(NMAX))
	go func() {
		worker(total, 0)
		done <- true
	}()
	go func() {
		worker(total, 1)
		done <- true
	}()
	go func() {
		worker(total, 2)
		done <- true
	}()
	go func() {
		worker(total, 3)
		done <- true
	}()
	<-done
	<-done
	<-done
	<-done

	for i = 2; i < NMAX; i++ {
		if flags[i] == false {
			s += i
		}
	}
	return s
}

var data11 = [23][23]int{
	{8, 02, 22, 97, 38, 15, 0, 40, 0, 75, 04, 05, 07, 78, 52, 12, 50, 77, 91, 8, 0, 0, 0},
	{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 04, 56, 62, 0, 0, 0, 0},
	{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 03, 49, 13, 36, 65, 0, 0, 0},
	{52, 70, 95, 23, 04, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 02, 36, 91, 0, 0, 0},
	{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80, 0, 0, 0},
	{24, 47, 32, 60, 99, 03, 45, 02, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50, 0, 0, 0},
	{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70, 0, 0, 0},
	{67, 26, 20, 68, 02, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21, 0, 0, 0},
	{24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72, 0, 0, 0},
	{21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95, 0, 0, 0},
	{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 03, 80, 04, 62, 16, 14, 9, 53, 56, 92, 0, 0, 0},
	{16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57, 0, 0, 0},
	{86, 56, 0, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58, 0, 0, 0},
	{19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40, 0, 0, 0},
	{04, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66, 0, 0, 0},
	{88, 36, 68, 87, 57, 62, 20, 72, 03, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69, 0, 0, 0},
	{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36, 0, 0, 0},
	{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 04, 36, 16, 0, 0, 0},
	{20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 05, 54, 0, 0, 0},
	{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 01, 89, 19, 67, 48, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func PE11() int {
	var i, j, max, horizontal, vertical, diagonal, rdiagonal int
	for i = 0; i < 20; i++ {
		for j = 0; j < 20; j++ {
			horizontal = data11[i][j] * data11[i][j+1] * data11[i][j+2] * data11[i][j+3]
			vertical = data11[i][j] * data11[i+1][j] * data11[i+2][j] * data11[i+3][j]
			diagonal = data11[i][j] * data11[i+1][j+1] * data11[i+2][j+2] * data11[i+3][j+3]
			rdiagonal = data11[i][j+3] * data11[i+1][j+2] * data11[i+2][j+1] * data11[i+3][j]
			if horizontal > max {
				max = horizontal
			}
			if vertical > max {
				max = vertical
			}
			if diagonal > max {
				max = diagonal
			}
			if rdiagonal > max {
				max = rdiagonal
			}
		}
	}
	return max
}

//////
func getmap(n int) map[int]int {
	ret := make(map[int]int)
	return ret
}
func mergemap(a, b map[int]int) map[int]int {
	return a
}
func calcprod(m map[int]int) int {
	return 0
}
func PE12(N int) int {
	var i, num int
	for i = 1; i < 2; i++ {
		num = i * (i + 1) / 2
		products := mergemap(getmap(i), getmap(i+1))
		products[2]--
		if calcprod(products) > N {
			return num
		}
	}
	return 0
}

//////
// Create Big int to [lint every 18-numbers in lower first order
func BigNum(data string) []int64 {
	var length, zeros int

	// Add leading zeros
	data = strings.TrimSpace(data)
	if len(data)%18 != 0 {
		zeros = 18 - len(data)%18
		data = strings.Repeat("0", zeros) + data
	}
	length = len(data) / 18

	// Fill int slice in lower first order
	ret := make([]int64, length)
	for i := 0; i < length; i++ {
		ret[length-i-1], _ = strconv.ParseInt(data[i*18:(i+1)*18], 10, 64)
	}
	return ret
}

// Sum two Big int created by BigNum
func BigSum(a, b []int64) []int64 {
	if len(a) < len(b) {
		a, b = b, a
	}
	lena := len(a)
	lenb := len(b)
	var ret = make([]int64, lena)
	copy(ret, a)
	for i := 0; i < lenb; i++ {
		ret[i] = a[i] + b[i]
	}
	for i := 0; i < lena; i++ {
		if ret[i] >= 1e18 {
			ret[i] -= 1e18
			if i == lena-1 {
				ret = append(ret, 1)
			} else {
				ret[i+1]++
			}
		}
	}
	return ret
}

// Length of a Big int created by BigNum
func BigLen(a []int64) (ret int) {
	ret = 18 * (len(a) - 1)
	last := a[len(a)-1]
	var i int
	for i = 0; last > 0; i++ {
		last /= 10
	}
	ret += i
	return
}

func PE13() string {
	file, err := os.Open("data/data13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	read := reader.ReadString
	line, err := read('\n')
	if err != nil {
		return ""
	}
	a := BigNum(line)
	for line, err = read('\n'); err == nil; line, err = read('\n') {
		b := BigNum(line)
		a = BigSum(a, b)
	}
	return strconv.FormatInt(a[0], 10)[0:10]
}

//////
func genIterLen(n int64) int64 {
	var length int64 = 1
	for ; n != 1; length++ {
		if n%2 == 0 { //even
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return length
}
func PE14(limit int64) int64 {
	var i, maxlength, longest int64
	for i = limit; i >= 2; i-- {
		length := genIterLen(i)
		if maxlength < length {
			maxlength = length
			longest = i
		}
	}
	return longest
}

//////
func PE16(N int) uint64 {
	var a uint64
	// bignum := make([]int64)
	a = 1 << 63
	return a + a
}

//////
func PE19() (ret int) {
	l, _ := time.LoadLocation("Asia/Shanghai")
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			t := time.Date(y, time.Month(m), 1, 1, 1, 39, 108924743, l)
			if t.Weekday() == time.Sunday {
				// fmt.Println(t)
				ret++
			}
		}
	}
	return
}
func PE19b() (ret int) {
	week := 1 // 1900.1.1 is Monday
	for y := 1900; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			// get days of this month
			days := 31
			if m == 4 || m == 6 || m == 9 || m == 11 {
				days = 30
			}
			if m == 2 {
				days = 28
				if y%4 == 0 && y%100 != 0 || y%400 == 0 {
					days = 29
				}
			}

			// iterator days of month
			for d := 1; d <= days; d++ {
				if y == 1901 && m == 1 && d == 1 {
					ret = 0
				}
				if d == 1 && week == 0 {
					ret++
				}
				week = (week + 1) % 7
			}
		}
	}
	return
}

//////
func Pentagonal(n int) int {
	return n * (3*n - 1) / 2
}
func GenPentagons(max int) []int {
	curr := len(pentagons)
	if curr >= max {
		return pentagons
	}
	for n := curr + 1; ; n++ {
		pentagons = append(pentagons, Pentagonal(n))
		if Pentagonal(n) >= max {
			break
		}
	}
	return pentagons
}

var pentagons []int

func PE44() (ret int) {
	ret = 1 << 30
	for i := 2; ; i++ {
		GenPentagons(Pentagonal(i))
		for j := 1; j < i; j++ {
			pi := pentagons[i-1]
			pj := pentagons[j-1]
			if InInts(pentagons, pi-pj) {
				GenPentagons(pi + pj)
				if InInts(pentagons, pi+pj) {
					if pi-pj < ret {
						ret = pi - pj
						return
						// fmt.Println(ret)
					}
				}
			}
		}
	}
}

//////
func Triangle(n int) int {
	return n * (n + 1) / 2
}
func Hexagonal(n int) int {
	return n * (2*n - 1)
}
func PE45() int {
	var t, p, h int = 1, 1, 1
	var T, P, H int
	for {
		T = Triangle(t)
		P = Pentagonal(p)
		H = Hexagonal(h)
		if T == P && T == H {
			fmt.Println(t, p, h, T)
		}
		switch {
		case T <= P && T <= H:
			t++
		case P <= T && P <= H:
			p++
		default:
			h++
		}
	}
	return 0
}

// Check if v in an asc-sorted int slice
func InInts(slice []int, v int) bool {
	index := sort.SearchInts(slice, v)
	if index == len(slice) || slice[index] != v {
		return false
	}
	return true
}

//////
func check_oddcomposites(n int) bool {
	var i int = 1
	for p := n - 2*i*i; p > 0; p = n - 2*i*i {
		if InInts(primes, p) {
			return true
		}
		i++
	}
	return false
}

var primes = []int{2, 3, 5, 7, 11, 13}
var oddcomposites = []int{9, 15}

func PE46() int {
	for i := 17; ; i += 2 {
		upbound := int(math.Sqrt(float64(i)))
		for j := 0; primes[j] <= upbound; j++ {
			if i%primes[j] == 0 {
				oddcomposites = append(oddcomposites, i)
				if !check_oddcomposites(i) {
					return i
				}
				goto NEXT
			}
		}
		primes = append(primes, i)
	NEXT:
	}
}

//////
// Generate all primes less than max to global primes
func GenPrimes(max int) {
	if primes[len(primes)-1] >= max {
		return
	}
	for i := primes[len(primes)-1] + 2; i <= max; i += 2 {
		upbound := int(math.Sqrt(float64(i)))
		for j := 0; primes[j] <= upbound; j++ {
			if i%primes[j] == 0 {
				goto NEXT
			}
		}
		primes = append(primes, i)
		// fmt.Println(len(primes), &primes[0])
	NEXT:
	}
}

// Generate prime factors and return as a map
func PrimeFactors(n int) map[int]int {
	var pfmap = make(map[int]int)
	for j := 0; n != 1; j++ {
		for n%primes[j] == 0 {
			pfmap[primes[j]]++
			n /= primes[j]
		}
	}
	return pfmap
}
func PE47(n int) int {
	ok := 0
	GenPrimes(100)
	for i := 4; ; i++ {
		if primes[len(primes)-1] < i {
			GenPrimes(i * 2)
		}
		if len(PrimeFactors(i)) != n {
			ok = 0
			continue
		}
		ok++
		if ok == n {
			return i - n + 1
		}
	}
}

//////
func IntsEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func StrsEquals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func MapNumbersInt(a int) (ret []int) {
	for a > 0 {
		ret = append(ret, a%10)
		a /= 10
	}
	return
}
func IsPermutations(a, b int) bool {
	ma := MapNumbersInt(a)
	mb := MapNumbersInt(b)
	sort.Ints(ma)
	sort.Ints(mb)
	return IntsEquals(ma, mb)
}
func PE49() (ret []string) {
	GenPrimes(10000)
	for i := 1001; i < 10000; i += 2 {
		j := i + 3330
		k := j + 3330
		if !InInts(primes, i) || !InInts(primes, j) || !InInts(primes, k) {
			continue
		}
		if IsPermutations(i, j) && IsPermutations(i, k) {
			ret = append(ret, strconv.Itoa(i)+strconv.Itoa(j)+strconv.Itoa(k))
		}
	}
	return
}

func DeleteIndex(f []int, i int) []int {
	return append(f[:i], f[i+1:]...)
}

//////
func PE24(n int) int {
	const length = 10 // 0 .. (length-1)
	n = n - 1         // nth choices will pass over n-1 choices
	var s, k [length]int
	var f []int
	var ret string

	for i := 0; i < length; i++ {
		f = append(f, i)
	}

	s[length-1] = 1 // 0! = 1, s[9-i]=i!
	for i := 1; i < length; i++ {
		s[length-i-1] = s[length-i] * i
	}
	for i := 0; i < length; i++ {
		if n < s[i] {
			continue
		}
		k[i] = n / s[i]
		n = n % s[i]
	}

	for i := 0; i < length; i++ {
		ret += strconv.Itoa(f[k[i]])
		f = DeleteIndex(f, k[i])
	}
	reti, _ := strconv.Atoi(ret)
	return reti
}

//////
// Generate perm by recurse
func Perm(list []int, i int, n int) {
	var j = 0
	if i == n {
		fmt.Println(list)
	} else {
		for j = i; j <= n; j++ {
			list[i], list[j] = list[j], list[i]
			Perm(list, i+1, n)
			list[i], list[j] = list[j], list[i]
		}
	}
}

//////
func PowerInt(a, b int) int {
	var ret int = 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}
func PowerSum(nlist []int, m int) int {
	var sum int
	for _, v := range nlist {
		sum += PowerInt(v, m)
	}
	return sum
}
func NumsInt(n int) (nlist []int) {
	for n > 0 {
		nlist = append(nlist, n%10)
		n /= 10
	}
	return
}
func genlimit(n int) int {
	for i := 1; ; i++ {
		if PowerInt(9, n)*i < PowerInt(10, i-1) {
			return PowerInt(10, i-1)
		}
	}
}
func SumInts(list []int) (sum int) {
	for _, v := range list {
		sum += v
	}
	return
}
func PE30(m int) int {
	var ret []int
	for i := 10; i < genlimit(m); i++ {
		nlist := NumsInt(i)
		if i == PowerSum(nlist, m) {
			ret = append(ret, i)
		}
	}
	return SumInts(ret)
}

//////
func IsPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func PE36(n int) int {
	var ret []int
	for i := 1; i < n; i++ {
		if IsPalindrome(strconv.Itoa(i)) && IsPalindrome(strconv.FormatInt(int64(i), 2)) {
			ret = append(ret, i)
		}
	}
	return SumInts(ret)
}

//////
func PE25(n int) (ret int) {
	fnp := BigNum("1")  //f1
	fn := BigNum("1")   //f2
	for i := 3; ; i++ { //fn
		fnp, fn = fn, BigSum(fnp, fn)
		if BigLen(fn) == n {
			ret = i
			return
		}
	}
}

//////
func RightTri(a, b, c int) bool {
	if a*a+b*b == c*c {
		return true
	}
	return false
}
func PE39(n int) int {
	var a, b, c, p int
	var sum, max, maxp int
	for p = 3; p <= n; p++ {
		sum = 0
		for c = p / 3; c <= p/2; c++ {
			for a = 1; a < p/3; a++ {
				b = p - c - a
				if b <= 0 {
					continue
				}
				if RightTri(a, b, c) {
					sum++
				}
			}
		}
		if sum > max {
			max = sum
			maxp = p
		}
	}
	fmt.Println("summax", max)
	return maxp
}

// About 1 hour
func PE58() int {
	var n, step int
	var yes, no int
	var tl, tr, bl, br int // number of four angle
	no = 1                 // 1 is not prime
	n = 1
	for i := 3; ; i += 2 {
		maxprimes := primes[len(primes)-1]
		if maxprimes < i*i {
			GenPrimes(2 * maxprimes)
		}
		step = i - 1
		tr = n + step
		tl = tr + step
		bl = tl + step
		br = bl + step
		n = br
		if InInts(primes, tr) {
			yes++
		} else {
			no++
		}
		if InInts(primes, tl) {
			yes++
		} else {
			no++
		}
		if InInts(primes, bl) {
			yes++
		} else {
			no++
		}
		if InInts(primes, br) {
			yes++
		} else {
			no++
		}
		fmt.Println(i, yes, no+yes)
		if 10*yes < yes+no {
			return i
		}
	}
}

var base int = 1e10

func MulTail(a int, b int) (ret int) {
	a = a % base
	b = b % base
	ret = a * b
	ret = ret % base
	return
}
func PowerTail(a int, b int) (ret int) {
	ret = 1
	for i := 0; i < b; i++ {
		ret = MulTail(ret, a)
	}
	return
}
func PE48(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += PowerTail(i, i)
	}
	return sum % base
}

//////
var sumprimes = []int{2, 5, 10, 17, 28, 41}

// Generate summery of primes at lest to n
func SumPrimes(n int) {
	if sumprimes[len(sumprimes)-1] >= n {
		return
	}

	for sumprimes[len(sumprimes)-1] < n {
		GenPrimes(primes[len(primes)-1] * 2)
		for i := len(sumprimes); i < len(primes); i++ {
			sumprimes = append(sumprimes, sumprimes[i-1]+primes[i])
		}
	}
}
func PE50(n int) (ret int) {
	GenPrimes(n)
	SumPrimes(n)
	max := sort.SearchInts(sumprimes, n)
	for i := max; ; i-- {
		for j := 0; j <= max-i; j++ {
			if j == 0 {
				if ret = sumprimes[i-1]; InInts(primes, ret) {
					return
				}
			} else {
				if ret = sumprimes[i+j-1] - sumprimes[j-1]; InInts(primes, ret) {
					return
				}
			}
		}
	}
}

//////
func RelativePrimes(n int) (ret []int) {
	return
}

// max(n / phi(n)) below N,
// will be the max(phi(n)) below N appears the first time.
func PE69(n int) (ret int) {
	return
}
