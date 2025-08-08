# Go (Golang) Cheatsheet для новичков

> Полный справочник по основам языка программирования Go для начинающих разработчиков

## 📋 Содержание

1. [Базовая структура программы](#базовая-структура-программы)
2. [Переменные и типы данных](#переменные-и-типы-данных)
3. [Операторы](#операторы)
4. [Условные операторы](#условные-операторы)
5. [Циклы](#циклы)
6. [Массивы и слайсы](#массивы-и-слайсы)
7. [Карты (Maps)](#карты-maps)
8. [Функции](#функции)
9. [Структуры](#структуры)
10. [Методы](#методы)
11. [Интерфейсы](#интерфейсы)
12. [Указатели](#указатели)
13. [Обработка ошибок](#обработка-ошибок)
14. [Горутины и каналы](#горутины-и-каналы)
15. [Пакеты и импорты](#пакеты-и-импорты)
16. [Полезные встроенные функции](#полезные-встроенные-функции)
17. [defer](#defer)
18. [Полезные стандартные пакеты](#полезные-стандартные-пакеты)
19. [Советы для новичков](#советы-для-новичков)

---

## Базовая структура программы

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Переменные и типы данных

### Объявление переменных

```go
// Полное объявление
var name string = "John"
var age int = 30

// Автоматическое определение типа
var name = "John"
var age = 30

// Краткое объявление (только внутри функций)
name := "John"
age := 30

// Множественное объявление
var a, b, c int = 1, 2, 3
x, y := 10, "text"
```

### Основные типы данных

```go
// Целые числа
var i int = 42
var i8 int8 = 127        // -128 до 127
var i16 int16 = 32767    // -32768 до 32767
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807

var ui uint = 42         // беззнаковые
var ui8 uint8 = 255      // 0 до 255
var ui16 uint16 = 65535
var ui32 uint32 = 4294967295
var ui64 uint64 = 18446744073709551615

// Вещественные числа
var f32 float32 = 3.14
var f64 float64 = 3.14159265359

// Логический тип
var flag bool = true

// Строки
var text string = "Hello, Go!"

// Символ (rune = int32)
var char rune = 'A'

// Байт (byte = uint8)
var b byte = 65
```

### Константы

```go
const PI = 3.14159
const (
    Red   = "red"
    Green = "green"
    Blue  = "blue"
)

// iota - автоинкремент
const (
    Monday = iota    // 0
    Tuesday          // 1
    Wednesday        // 2
    Thursday         // 3
    Friday           // 4
)
```

## Операторы

### Арифметические операторы

```go
a := 10
b := 3

sum := a + b        // 13
diff := a - b       // 7
product := a * b    // 30
quotient := a / b   // 3
remainder := a % b  // 1

// Унарные операторы
a++  // a = a + 1
a--  // a = a - 1
```

### Операторы сравнения

```go
a := 10
b := 5

a == b  // false (равно)
a != b  // true (не равно)
a > b   // true (больше)
a < b   // false (меньше)
a >= b  // true (больше или равно)
a <= b  // false (меньше или равно)
```

### Логические операторы

```go
a := true
b := false

a && b  // false (И)
a || b  // true (ИЛИ)
!a      // false (НЕ)
```

## Условные операторы

### if/else

```go
age := 18

if age >= 18 {
    fmt.Println("Совершеннолетний")
} else if age >= 13 {
    fmt.Println("Подросток")
} else {
    fmt.Println("Ребенок")
}

// Краткое объявление в if
if age := getAge(); age >= 18 {
    fmt.Println("Совершеннолетний")
}
```

### switch

```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Понедельник")
case "Tuesday":
    fmt.Println("Вторник")
case "Saturday", "Sunday":
    fmt.Println("Выходной")
default:
    fmt.Println("Другой день")
}

// switch без выражения
score := 85
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 70:
    fmt.Println("C")
default:
    fmt.Println("F")
}
```

## Циклы

### for - единственный цикл в Go

```go
// Классический for
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Аналог while
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}

// Бесконечный цикл
for {
    // бесконечный цикл
    break // выход из цикла
}

// for-range для массивов/слайсов
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Игнорирование индекса
for _, value := range numbers {
    fmt.Println(value)
}
```

## Массивы и слайсы

### Массивы

```go
// Объявление массива
var arr [5]int
arr[0] = 1
arr[1] = 2

// Инициализация при объявлении
numbers := [5]int{1, 2, 3, 4, 5}
names := [3]string{"Alice", "Bob", "Charlie"}

// Автоматическое определение размера
auto := [...]int{1, 2, 3, 4}

// Длина массива
fmt.Println(len(numbers)) // 5
```

### Слайсы (более гибкие)

```go
// Создание слайса
var slice []int
slice = append(slice, 1, 2, 3)

// Литерал слайса
numbers := []int{1, 2, 3, 4, 5}

// Создание слайса с make
slice2 := make([]int, 5)      // длина 5, все элементы 0
slice3 := make([]int, 3, 5)   // длина 3, емкость 5

// Срезы
slice4 := numbers[1:4]  // элементы с индекса 1 до 3
slice5 := numbers[:3]   // первые 3 элемента
slice6 := numbers[2:]   // с индекса 2 до конца

// Добавление элементов
slice = append(slice, 4, 5, 6)

// Длина и емкость
fmt.Println(len(slice))  // длина
fmt.Println(cap(slice))  // емкость
```

## Карты (Maps)

```go
// Объявление карты
var m map[string]int
m = make(map[string]int)

// Литерал карты
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
    "Charlie": 35,
}

// Добавление/изменение элемента
ages["David"] = 28

// Получение значения
age := ages["Alice"]

// Проверка существования ключа
age, exists := ages["Eve"]
if exists {
    fmt.Printf("Eve's age: %d\n", age)
} else {
    fmt.Println("Eve not found")
}

// Удаление элемента
delete(ages, "Bob")

// Перебор карты
for name, age := range ages {
    fmt.Printf("%s: %d\n", name, age)
}
```

## Функции

### Базовый синтаксис

```go
// Простая функция
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Функция с возвращаемым значением
func add(a, b int) int {
    return a + b
}

// Функция с несколькими возвращаемыми значениями
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Именованные возвращаемые значения
func calculate(a, b int) (sum, product int) {
    sum = a + b
    product = a * b
    return // автоматически возвращает sum и product
}

// Вариадические функции
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Использование
result := sum(1, 2, 3, 4, 5)
```

### Функции как значения

```go
// Присвоение функции переменной
var operation func(int, int) int
operation = add

// Анонимные функции
multiply := func(a, b int) int {
    return a * b
}

// Функция как параметр
func calculate(a, b int, op func(int, int) int) int {
    return op(a, b)
}

result := calculate(5, 3, multiply)
```

## Структуры

```go
// Определение структуры
type Person struct {
    Name string
    Age  int
    City string
}

// Создание экземпляра
var p1 Person
p1.Name = "Alice"
p1.Age = 30

// Литерал структуры
p2 := Person{
    Name: "Bob",
    Age:  25,
    City: "New York",
}

// Краткий синтаксис (порядок полей важен)
p3 := Person{"Charlie", 35, "London"}

// Указатель на структуру
p4 := &Person{"David", 28, "Paris"}

// Встроенные структуры
type Employee struct {
    Person    // встраивание
    Position string
    Salary   float64
}

emp := Employee{
    Person: Person{"Eve", 32, "Berlin"},
    Position: "Developer",
    Salary: 50000,
}
```

## Методы

```go
// Метод для структуры
func (p Person) Greet() {
    fmt.Printf("Hello, I'm %s\n", p.Name)
}

// Метод с указателем (может изменять структуру)
func (p *Person) HaveBirthday() {
    p.Age++
}

// Использование методов
person := Person{"Alice", 30, "NYC"}
person.Greet()
person.HaveBirthday()
fmt.Println(person.Age) // 31
```

## Интерфейсы

```go
// Определение интерфейса
type Writer interface {
    Write([]byte) (int, error)
}

type Speaker interface {
    Speak() string
}

// Структура, реализующая интерфейс
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

// Использование интерфейса
func makeSound(s Speaker) {
    fmt.Println(s.Speak())
}

dog := Dog{"Buddy"}
cat := Cat{"Whiskers"}

makeSound(dog) // Woof!
makeSound(cat) // Meow!

// Пустой интерфейс (любой тип)
var anything interface{}
anything = 42
anything = "hello"
anything = []int{1, 2, 3}
```

## Указатели

```go
// Объявление указателя
var p *int

// Получение адреса переменной
x := 42
p = &x

// Разыменование указателя
fmt.Println(*p) // 42

// Изменение значения через указатель
*p = 100
fmt.Println(x) // 100

// new() создает указатель на нулевое значение
ptr := new(int)
*ptr = 42
```

## Обработка ошибок

```go
// Создание ошибки
import "errors"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// Использование
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)

// Форматирование ошибок
func validateAge(age int) error {
    if age < 0 {
        return fmt.Errorf("age cannot be negative: %d", age)
    }
    return nil
}
```

## Горутины и каналы

### Горутины

```go
import "time"

// Запуск горутины
go func() {
    fmt.Println("Горутина работает")
}()

// Функция как горутина
func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(time.Second)
    }
}

go printNumbers()
time.Sleep(6 * time.Second) // ждем завершения
```

### Каналы

```go
// Создание канала
ch := make(chan int)

// Отправка в канал (в горутине)
go func() {
    ch <- 42
}()

// Получение из канала
value := <-ch
fmt.Println(value)

// Буферизованный канал
buffered := make(chan int, 3)
buffered <- 1
buffered <- 2
buffered <- 3

// Закрытие канала
close(buffered)

// Проверка закрытия канала
value, ok := <-buffered
if ok {
    fmt.Println("Получено:", value)
} else {
    fmt.Println("Канал закрыт")
}

// Перебор канала
for value := range buffered {
    fmt.Println(value)
}
```

### select

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    time.Sleep(2 * time.Second)
    ch1 <- "from ch1"
}()

go func() {
    time.Sleep(1 * time.Second)
    ch2 <- "from ch2"
}()

select {
case msg1 := <-ch1:
    fmt.Println("Получено", msg1)
case msg2 := <-ch2:
    fmt.Println("Получено", msg2)
case <-time.After(3 * time.Second):
    fmt.Println("Таймаут")
default:
    fmt.Println("Нет готовых каналов")
}
```

## Пакеты и импорты

```go
// Импорт стандартных пакетов
import "fmt"
import "time"
import "strings"

// Множественный импорт
import (
    "fmt"
    "time"
    "strings"
)

// Импорт с алиасом
import f "fmt"

// Импорт только для side effects
import _ "image/png"

// Dot импорт (не рекомендуется)
import . "fmt"
```

## Полезные встроенные функции

```go
// len() - длина
fmt.Println(len("hello"))    // 5
fmt.Println(len([]int{1,2})) // 2

// cap() - емкость
slice := make([]int, 3, 5)
fmt.Println(cap(slice)) // 5

// append() - добавление в слайс
slice = append(slice, 4, 5)

// copy() - копирование слайса
dest := make([]int, len(slice))
copy(dest, slice)

// make() - создание карт, слайсов, каналов
m := make(map[string]int)
s := make([]int, 5)
ch := make(chan int)

// new() - выделение памяти
ptr := new(int)

// panic() и recover()
func safeDivide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
            result = 0
        }
    }()
    
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}
```

## defer

```go
func example() {
    fmt.Println("start")
    
    defer fmt.Println("deferred 1")
    defer fmt.Println("deferred 2")
    
    fmt.Println("middle")
    
    // defer выполняется в обратном порядке при выходе из функции
    // Выведет:
    // start
    // middle
    // deferred 2
    // deferred 1
}

// Использование defer для закрытия ресурсов
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // автоматически закроется при выходе
    
    // работа с файлом
    return nil
}
```

## Полезные стандартные пакеты

```go
// fmt - форматированный ввод/вывод
fmt.Printf("Name: %s, Age: %d\n", name, age)
fmt.Sprintf("Formatted string: %v", value)

// strings - работа со строками
strings.ToUpper("hello")           // "HELLO"
strings.Contains("hello", "ell")   // true
strings.Split("a,b,c", ",")        // ["a", "b", "c"]

// strconv - конвертация строк
i, err := strconv.Atoi("123")      // string to int
s := strconv.Itoa(123)             // int to string

// time - работа с временем
now := time.Now()
time.Sleep(2 * time.Second)
formatted := now.Format("2006-01-02 15:04:05")

// math - математические функции
import "math"
math.Sqrt(16)  // 4
math.Max(3, 7) // 7

// os - работа с ОС
os.Getenv("HOME")
os.Exit(1)
```

## 💡 Советы для новичков

- **Всегда проверяйте ошибки**: Go явно требует обработки ошибок
- **Используйте gofmt**: автоматическое форматирование кода
- **Читайте Effective Go**: официальное руководство по стилю
- **Изучите стандартную библиотеку**: она очень богатая
- **Практикуйте goroutines и channels**: это сила Go
- **Не игнорируйте линтеры**: golint, go vet помогают писать качественный код

## 📚 Дополнительные ресурсы

- [Официальная документация Go](https://golang.org/doc/)
- [A Tour of Go](https://tour.golang.org/welcome/1) - интерактивный тур
- [Effective Go](https://golang.org/doc/effective_go.html) - лучшие практики
- [Go by Example](https://gobyexample.com/) - примеры кода
- [Go Playground](https://play.golang.org/) - онлайн редактор

## 🤝 Сообщество

- [Go Gophers Slack](https://gophers.slack.com/)
- [Reddit r/golang](https://www.reddit.com/r/golang/)
- [Stack Overflow](https://stackoverflow.com/questions/tagged/go)

---

**⭐ Этот cheatsheet покрывает основы Go. Для глубокого изучения рекомендую официальную документацию и практику!**

*Удачи в изучении Go! 🚀*