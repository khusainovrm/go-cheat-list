# GORM Cheat Sheet

Полный справочник по GORM (Go Object-Relational Mapping) для быстрого использования в проектах.

## 📋 Содержание

- [Установка и подключение](#установка-и-подключение)
- [Модели и миграции](#модели-и-миграции)
- [CRUD операции](#crud-операции)
- [Запросы и фильтрация](#запросы-и-фильтрация)
- [Связи (Associations)](#связи-associations)
- [Транзакции](#транзакции)
- [Хуки (Hooks)](#хуки-hooks)
- [Индексы](#индексы)
- [Soft Delete](#soft-delete)
- [Batch операции](#batch-операции)
- [Raw SQL](#raw-sql)
- [Расширенные возможности](#расширенные-возможности)

## 🚀 Установка и подключение

### Установка
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres  # PostgreSQL
go get -u gorm.io/driver/mysql     # MySQL
go get -u gorm.io/driver/sqlite    # SQLite
```

### Подключение к базе данных
```go
import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "gorm.io/driver/mysql"
    "gorm.io/driver/sqlite"
)

// PostgreSQL
dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// MySQL
dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// SQLite
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
```

### Конфигурация GORM
```go
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    NamingStrategy: schema.NamingStrategy{
        TablePrefix:   "t_",   // префикс таблиц
        SingularTable: false,  // использовать множественное число
    },
    Logger: logger.Default.LogMode(logger.Info), // логирование
})
```

## 🏗️ Модели и миграции

### Базовая модель
```go
type User struct {
    ID        uint           `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
    Name      string         `gorm:"size:255;not null"`
    Email     string         `gorm:"uniqueIndex"`
    Age       uint8
    Birthday  *time.Time
    Active    bool           `gorm:"default:true"`
}
```

### Встроенная модель (Embedded)
```go
type User struct {
    gorm.Model  // содержит ID, CreatedAt, UpdatedAt, DeletedAt
    Name   string
    Email  string
}
```

### Теги полей
```go
type User struct {
    ID       uint   `gorm:"primaryKey;autoIncrement"`
    Name     string `gorm:"size:255;not null;unique"`
    Email    string `gorm:"type:varchar(100);uniqueIndex:idx_email"`
    Age      int    `gorm:"check:age > 0"`
    Password string `gorm:"->;<-:create"` // только для чтения, кроме создания
    Token    string `gorm:"-"`            // игнорировать поле
    Salary   int    `gorm:"default:0"`
}
```

### Миграции
```go
// Автомиграция
db.AutoMigrate(&User{})

// Создание таблицы
db.Migrator().CreateTable(&User{})

// Добавление столбца
db.Migrator().AddColumn(&User{}, "nickname")

// Удаление столбца
db.Migrator().DropColumn(&User{}, "nickname")

// Проверка существования таблицы
db.Migrator().HasTable(&User{})

// Переименование таблицы
db.Migrator().RenameTable(&User{}, &UserV2{})
```

## 📝 CRUD операции

### Create (Создание)
```go
// Создание одной записи
user := User{Name: "John", Email: "john@example.com", Age: 25}
result := db.Create(&user)
// user.ID возвращает автосгенерированный ID

// Создание с выбранными полями
db.Select("name", "age").Create(&user)

// Создание без определенных полей
db.Omit("password", "age").Create(&user)

// Создание с Map
db.Model(&User{}).Create(map[string]interface{}{
    "Name": "John",
    "Age":  25,
})

// Получение количества затронутых строк и ошибки
result := db.Create(&user)
result.Error        // возвращает ошибку
result.RowsAffected // возвращает количество вставленных записей
```

### Read (Чтение)
```go
// Получение первой записи по первичному ключу
var user User
db.First(&user, 1) // SELECT * FROM users WHERE id = 1 ORDER BY id LIMIT 1;

// Получение первой записи
db.First(&user) // SELECT * FROM users ORDER BY id LIMIT 1;

// Получение последней записи
db.Last(&user) // SELECT * FROM users ORDER BY id DESC LIMIT 1;

// Получение всех записей
var users []User
db.Find(&users) // SELECT * FROM users;

// Получение записи по условию
db.First(&user, "name = ?", "John")

// Получение конкретных полей
db.Select("name", "age").Find(&users)

// Исключение полей
db.Omit("password").Find(&users)
```

### Update (Обновление)
```go
// Обновление одного поля
db.Model(&user).Update("name", "John Updated")

// Обновление нескольких полей
db.Model(&user).Updates(User{Name: "John", Age: 30})

// Обновление с Map
db.Model(&user).Updates(map[string]interface{}{
    "name": "John",
    "age":  30,
})

// Обновление конкретных полей
db.Model(&user).Select("name").Updates(map[string]interface{}{
    "name": "John",
    "age":  0,  // age не будет обновлен
})

// Обновление всех полей кроме указанных
db.Model(&user).Omit("name").Updates(map[string]interface{}{
    "name": "John",  // name не будет обновлен
    "age":  30,
})

// Обновление без хуков
db.Model(&user).UpdateColumn("name", "John")

// Batch обновление
db.Model(&User{}).Where("role = ?", "admin").Updates(User{Name: "Admin"})
```

### Delete (Удаление)
```go
// Удаление по первичному ключу
db.Delete(&user, 1)

// Удаление по условию
db.Where("name = ?", "John").Delete(&User{})

// Batch удаление
db.Delete(&User{}, []int{1, 2, 3})

// Soft Delete (если есть DeletedAt поле)
db.Delete(&user) // UPDATE users SET deleted_at = NOW() WHERE id = 1;

// Permanent Delete (жесткое удаление)
db.Unscoped().Delete(&user) // DELETE FROM users WHERE id = 1;

// Восстановление soft deleted записи
db.Unscoped().Model(&user).Update("deleted_at", nil)
```

## 🔍 Запросы и фильтрация

### Where условия
```go
// Простые условия
db.Where("name = ?", "John").Find(&users)
db.Where("name <> ?", "John").Find(&users)
db.Where("name IN ?", []string{"John", "Jane"}).Find(&users)
db.Where("name LIKE ?", "%John%").Find(&users)
db.Where("age > ?", 25).Find(&users)

// Множественные условия
db.Where("name = ? AND age >= ?", "John", 25).Find(&users)

// Struct условия
db.Where(&User{Name: "John", Age: 25}).Find(&users)

// Map условия
db.Where(map[string]interface{}{
    "name": "John",
    "age":  25,
}).Find(&users)

// Отрицание
db.Not("name = ?", "John").Find(&users)
db.Not(map[string]interface{}{"name": []string{"John", "Jane"}}).Find(&users)

// Or условия
db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
```

### Сортировка и группировка
```go
// Сортировка
db.Order("age desc, name").Find(&users)
db.Order("age desc").Order("name").Find(&users)

// Группировка
db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").Find(&result)

// Having
db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group1").Find(&result)
```

### Лимит и смещение
```go
// Лимит
db.Limit(3).Find(&users)

// Смещение
db.Offset(3).Find(&users)

// Пагинация
db.Limit(10).Offset(20).Find(&users)

// Отмена лимита/смещения
db.Limit(10).Find(&users1).Limit(-1).Find(&users2) // users2 без лимита
```

### Подзапросы
```go
// Подзапрос
subQuery := db.Select("AVG(age)").Where("name LIKE ?", "name%").Table("users")
db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)

// Подзапрос в FROM
db.Table("(?) as u", db.Model(&User{}).Select("name", "age")).Where("age = ?", 18).Find(&results)
```

## 🔗 Связи (Associations)

### Belongs To
```go
type User struct {
    ID         uint
    Name       string
    CompanyID  uint
    Company    Company `gorm:"foreignKey:CompanyID"`
}

type Company struct {
    ID   uint
    Name string
}

// Загрузка со связями
db.Preload("Company").Find(&users)

// Вложенный preload
db.Preload("Company.Address").Find(&users)
```

### Has One
```go
type User struct {
    ID      uint
    Name    string
    Profile Profile `gorm:"foreignKey:UserID"`
}

type Profile struct {
    ID     uint
    UserID uint
    Bio    string
}

// Загрузка
db.Preload("Profile").Find(&users)
```

### Has Many
```go
type User struct {
    ID     uint
    Name   string
    Orders []Order `gorm:"foreignKey:UserID"`
}

type Order struct {
    ID     uint
    UserID uint
    Amount float64
}

// Загрузка
db.Preload("Orders").Find(&users)

// Загрузка с условием
db.Preload("Orders", "amount > ?", 100).Find(&users)

// Загрузка с функцией
db.Preload("Orders", func(db *gorm.DB) *gorm.DB {
    return db.Order("amount DESC")
}).Find(&users)
```

### Many To Many
```go
type User struct {
    ID    uint
    Name  string
    Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
    ID   uint
    Name string
}

// Создание связи
user := User{Name: "John"}
db.Create(&user)

role := Role{Name: "admin"}
db.Create(&role)

// Добавление связи
db.Model(&user).Association("Roles").Append(&role)

// Загрузка
db.Preload("Roles").Find(&users)

// Замена связей
db.Model(&user).Association("Roles").Replace(&roles)

// Удаление связи
db.Model(&user).Association("Roles").Delete(&role)

// Очистка всех связей
db.Model(&user).Association("Roles").Clear()
```

### Полиморфные связи
```go
type Dog struct {
    ID   uint
    Name string
    Toys []Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
    ID        uint
    Name      string
    OwnerID   uint
    OwnerType string
}

// Загрузка
db.Preload("Toys").Find(&dogs)
```

## 💾 Транзакции

### Ручные транзакции
```go
// Начало транзакции
tx := db.Begin()

// Откат в случае ошибки
defer func() {
    if r := recover(); r != nil {
        tx.Rollback()
    }
}()

// Выполнение операций
if err := tx.Create(&user).Error; err != nil {
    tx.Rollback()
    return err
}

if err := tx.Create(&profile).Error; err != nil {
    tx.Rollback()
    return err
}

// Коммит транзакции
return tx.Commit().Error
```

### Автоматические транзакции
```go
err := db.Transaction(func(tx *gorm.DB) error {
    // Выполнение операций в транзакции
    if err := tx.Create(&user).Error; err != nil {
        return err // откат автоматически
    }

    if err := tx.Create(&profile).Error; err != nil {
        return err // откат автоматически
    }

    // Возвращение nil коммитит транзакцию
    return nil
})
```

### Savepoints
```go
tx := db.Begin()

tx.Create(&user1)

tx.SavePoint("sp1")
tx.Create(&user2)
tx.RollbackTo("sp1") // user2 не будет создан

tx.Commit() // user1 будет создан
```

## 🪝 Хуки (Hooks)

### Before/After хуки
```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.ID = uuid.New()
    if u.Role == "admin" {
        return errors.New("invalid role")
    }
    return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
    // Логирование создания пользователя
    log.Printf("User created: %v", u.Name)
    return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
    if u.Role == "admin" && tx.Statement.Changed("role") {
        return errors.New("role cannot be changed")
    }
    return
}

func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
    // Очистка кеша
    cache.Delete("user:" + fmt.Sprintf("%d", u.ID))
    return
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
    if u.Role == "admin" {
        return errors.New("admin cannot be deleted")
    }
    return
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
    // Очистка связанных данных
    tx.Where("user_id = ?", u.ID).Delete(&Profile{})
    return
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
    // Декодирование зашифрованных полей
    if u.EncryptedData != "" {
        u.DecryptedData = decrypt(u.EncryptedData)
    }
    return
}
```

### Пропуск хуков
```go
// Создание без хуков
db.Session(&gorm.Session{SkipHooks: true}).Create(&user)

// Обновление без хуков
db.Session(&gorm.Session{SkipHooks: true}).Save(&user)
```

## 🗂️ Индексы

### Создание индексов через теги
```go
type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"index"`                          // обычный индекс
    Email string `gorm:"uniqueIndex"`                   // уникальный индекс
    Code  string `gorm:"index:idx_code,unique"`         // именованный уникальный индекс
    Age   int    `gorm:"index:,sort:desc"`              // индекс с сортировкой
    Bio   string `gorm:"index:,type:gin"`               // GIN индекс (PostgreSQL)
    Tags  string `gorm:"index:,type:btree,length:10"`   // индекс с длиной
}
```

### Составные индексы
```go
type User struct {
    ID        uint   `gorm:"primaryKey"`
    FirstName string `gorm:"index:idx_name"`
    LastName  string `gorm:"index:idx_name"`
    Email     string `gorm:"index:idx_email_age"`
    Age       int    `gorm:"index:idx_email_age"`
}
```

### Создание индексов программно
```go
// Создание индекса
type User struct {
    ID   uint
    Name string
}

func (User) TableName() string {
    return "users"
}

// В миграции
db.Exec("CREATE INDEX idx_users_name ON users(name)")

// Или через GORM Migrator
db.Migrator().CreateIndex(&User{}, "Name")
db.Migrator().CreateIndex(&User{}, "idx_user_name")

// Удаление индекса
db.Migrator().DropIndex(&User{}, "Name")
db.Migrator().DropIndex(&User{}, "idx_user_name")

// Проверка существования индекса
db.Migrator().HasIndex(&User{}, "Name")
```

## 🗑️ Soft Delete

### Включение Soft Delete
```go
type User struct {
    ID        uint
    Name      string
    DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete поле
}

// или использовать gorm.Model
type User struct {
    gorm.Model
    Name string
}
```

### Операции с Soft Delete
```go
// Soft delete (запись помечается как удаленная)
db.Delete(&user)
// UPDATE users SET deleted_at = NOW() WHERE id = 1;

// Поиск исключает soft deleted записи
db.Find(&users)
// SELECT * FROM users WHERE deleted_at IS NULL;

// Поиск только soft deleted записи
db.Unscoped().Where("deleted_at IS NOT NULL").Find(&users)

// Поиск включая soft deleted записи
db.Unscoped().Find(&users)
// SELECT * FROM users;

// Permanent delete (физическое удаление)
db.Unscoped().Delete(&user)
// DELETE FROM users WHERE id = 1;

// Восстановление soft deleted записи
db.Unscoped().Model(&user).Update("deleted_at", nil)
```

### Кастомный Soft Delete
```go
type User struct {
    ID      uint
    Name    string
    IsDeleted bool `gorm:"softDelete:flag"` // 0 = не удален, 1 = удален
}

// Или с временем
type User struct {
    ID      uint
    Name    string
    DeletedAt int64 `gorm:"softDelete:nano"` // Unix timestamp в наносекундах
}

// Или с кастомным значением
type User struct {
    ID      uint
    Name    string
    IsDel   string `gorm:"softDelete:flag,softDeleteValue:yes"` // 'yes' для удаленных
}
```

## 🔄 Batch операции

### Batch Insert
```go
// Создание множественных записей
users := []User{
    {Name: "John", Email: "john@example.com"},
    {Name: "Jane", Email: "jane@example.com"},
    {Name: "Bob", Email: "bob@example.com"},
}

// Все за один раз
db.Create(&users)

// Батчами по 100
db.CreateInBatches(users, 100)

// Upsert (создание или обновление)
db.Clauses(clause.OnConflict{
    Columns:   []clause.Column{{Name: "email"}},
    DoUpdates: clause.AssignmentColumns([]string{"name", "updated_at"}),
}).Create(&users)
```

### Batch Update
```go
// Обновление множественных записей
db.Model(&User{}).Where("role = ?", "user").Updates(User{Active: true})

// Обновление с case when
db.Model(&User{}).Where("id IN ?", []int{1, 2, 3}).Updates(map[string]interface{}{
    "name": gorm.Expr("CASE WHEN id = 1 THEN 'John' WHEN id = 2 THEN 'Jane' ELSE 'Bob' END"),
})

// Batch update с различными значениями
users := []User{
    {ID: 1, Name: "John Updated"},
    {ID: 2, Name: "Jane Updated"},
}

for _, user := range users {
    db.Model(&user).Updates(User{Name: user.Name})
}
```

### Batch Delete
```go
// Удаление множественных записей
db.Delete(&User{}, []int{1, 2, 3})

// Удаление по условию
db.Where("created_at < ?", time.Now().AddDate(0, 0, -30)).Delete(&User{})

// Batch delete с лимитом
db.Where("active = ?", false).Limit(1000).Delete(&User{})
```

## 🔧 Raw SQL

### Выполнение Raw SQL
```go
type Result struct {
    ID   int
    Name string
    Age  int
}

// Raw SQL запрос
var result Result
db.Raw("SELECT id, name, age FROM users WHERE name = ?", "John").Scan(&result)

// Raw SQL в Find
var users []User
db.Raw("SELECT * FROM users WHERE age > ?", 25).Find(&users)

// Exec для модификации данных
db.Exec("UPDATE users SET name = ? WHERE id = ?", "John Updated", 1)

// Получение ошибки
result := db.Exec("UPDATE users SET name = ? WHERE id = ?", "John", 1)
if result.Error != nil {
    // обработка ошибки
}
fmt.Printf("Rows affected: %d\n", result.RowsAffected)
```

### SQL Builder
```go
// Построение динамических запросов
query := db.Model(&User{})

if name != "" {
    query = query.Where("name LIKE ?", "%"+name+"%")
}

if age > 0 {
    query = query.Where("age >= ?", age)
}

query.Find(&users)
```

### Named Arguments
```go
db.Raw("SELECT * FROM users WHERE name = @name AND age = @age", 
    sql.Named("name", "John"), 
    sql.Named("age", 25)).Find(&users)

// Или с map
db.Raw("SELECT * FROM users WHERE name = @name AND age = @age", 
    map[string]interface{}{
        "name": "John",
        "age":  25,
    }).Find(&users)
```

## ⚡ Расширенные возможности

### Scopes (Области)
```go
// Определение scope
func ActiveUsers(db *gorm.DB) *gorm.DB {
    return db.Where("active = ?", true)
}

func YoungUsers(db *gorm.DB) *gorm.DB {
    return db.Where("age < ?", 30)
}

func PaginateUsers(page, pageSize int) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        offset := (page - 1) * pageSize
        return db.Offset(offset).Limit(pageSize)
    }
}

// Использование scopes
db.Scopes(ActiveUsers, YoungUsers).Find(&users)
db.Scopes(PaginateUsers(2, 10)).Find(&users)
```

### Context
```go
ctx := context.Background()

// Использование контекста
db.WithContext(ctx).Find(&users)

// Таймаут
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

db.WithContext(ctx).Find(&users)
```

### Session
```go
// Новая сессия
session := db.Session(&gorm.Session{
    DryRun:      true,  // не выполнять SQL, только подготовить
    PrepareStmt: true,  // подготовленные выражения
    SkipHooks:   true,  // пропустить хуки
})

session.Find(&users)

// Клонирование сессии
newDB := db.Session(&gorm.Session{})
```

### Custom Data Types
```go
import (
    "database/sql/driver"
    "encoding/json"
    "errors"
)

type JSON map[string]interface{}

func (j JSON) Value() (driver.Value, error) {
    return json.Marshal(j)
}

func (j *JSON) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }
    return json.Unmarshal(bytes, &j)
}

type User struct {
    ID       uint
    Name     string
    Metadata JSON `gorm:"type:json"`
}
```

### Валидация
```go
type User struct {
    gorm.Model
    Name  string `gorm:"size:255;not null" validate:"required,min=2,max=50"`
    Email string `gorm:"uniqueIndex" validate:"required,email"`
    Age   int    `validate:"min=0,max=150"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
    return u.Validate()
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
    return u.Validate()
}

func (u *User) Validate() error {
    validate := validator.New()
    return validate.Struct(u)
}
```

### Логирование
```go
import "gorm.io/gorm/logger"

// Настройка логгера
newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags),
    logger.Config{
        SlowThreshold: time.Second,   // медленные запросы
        LogLevel:      logger.Silent, // уровень логирования
        Colorful:      false,         // цветной вывод
    },
)

db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
    Logger: newLogger,
})

// Логирование для конкретной сессии
db.Session(&gorm.Session{
    Logger: logger.Default.LogMode(logger.Info),
}).Find(&users)
```

### Кастомные типы
```go
type Status int

const (
    StatusInactive Status = iota
    StatusActive
    StatusDeleted
)

func (s Status) String() string {
    return [...]string{"inactive", "active", "deleted"}[s]
}

func (s *Status) Scan(value interface{}) error {
    *s = Status(value.(int64))
    return nil
}

func (s Status) Value() (driver.Value, error) {
    return int64(s), nil
}

type User struct {
    ID     uint
    Name   string
    Status Status
}
```

## 📊 Производительность

### Connection Pool
```go
sqlDB, err := db.DB()

// Максимальное количество открытых соединений
sqlDB.SetMaxOpenConns(100)

// Максимальное количество idle соединений
sqlDB.SetMaxIdleConns(10)

// Максимальное время жизни соединения
sqlDB.SetConnMaxLifetime(time.Hour)
```

### Prepared Statements
```go
// Включение prepared statements
db = db.Session(&gorm.Session{PrepareStmt: true})

// Использование prepared statements для повторяющихся запросов
stmt := db.Session(&gorm.Session{PrepareStmt: true}).Prepare("SELECT * FROM users WHERE name = ?")
defer stmt.Close()

var users []User
stmt.Find(&users, "John")
```

### Полезные советы по производительности

1. **Используйте Select для загрузки только нужных полей**:
   ```go
   db.Select("id", "name").Find(&users)
   ```

2. **Используйте индексы для часто используемых WHERE условий**

3. **Избегайте N+1 запросов с помощью Preload**:
   ```go
   db.Preload("Orders").Find(&users) // Вместо отдельного запроса для каждого пользователя
   ```

4. **Используйте FindInBatches для обработки больших объемов данных**:
   ```go
   db.FindInBatches(&users, 1000, func(tx *gorm.DB, batch int) error {
       for _, user := range users {
           // обработка пользователя
       }
       return nil
   })
   ```

## 🔍 Отладка

### SQL отладка
```go
// Включение SQL логов
db = db.Debug()

// Для одного запроса
db.Debug().Find(&users)

[//]: # (// Получение сгенерированного SQL)

[//]: # (stmt := db.Session&#40;&gorm.Session{DryRun: true}&#41;.Find&#40;&users&#41;.Statement)

[//]: # (fmt.)