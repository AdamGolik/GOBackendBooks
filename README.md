# Awesome Project API

## Opis projektu

Ten projekt to REST API napisane w Go, które umożliwia zarządzanie użytkownikami oraz książkami. API oferuje funkcje takie jak rejestracja, logowanie, tworzenie, aktualizowanie, usuwanie oraz wyszukiwanie książek i użytkowników.

## Instalacja

1. Sklonuj repozytorium
    ```bash
    git clone https://github.com/your_repository.git
    ```
2. Przejdź do katalogu projektu
    ```bash
    cd your_repository
    ```
3. Zainstaluj zależności
    ```bash
    go mod tidy
    ```
4. Uruchom aplikację
    ```bash
    go run main.go
    ```

## Endpoints API

### Rejestracja użytkownika

**Endpoint:** `/register`

**Metoda:** `POST`

**Opis:** Rejestruje nowego użytkownika.

**Treść żądania:**
```json
{
  "username": "newuser",
  "password": "password123",
  "email": "newuser@example.com",
  "phone": "123-456-7890",
  "address": "123 Main St"
}
```

**Treść odpowiedzi:**
```json
{
  "id": 1,
  "username": "newuser",
  "email": "newuser@example.com",
  "phone": "123-456-7890",
  "address": "123 Main St"
}
```

### Logowanie użytkownika

**Endpoint:** `/login`

**Metoda:** `POST`

**Opis:** Loguje użytkownika na podstawie podanych danych.

**Treść żądania:**
```json
{
  "username": "newuser",
  "email": "newuser@example.com",
  "password": "password123"
}
```

**Treść odpowiedzi:**
```json
{
  "id": 1,
  "username": "newuser",
  "email": "newuser@example.com",
  "phone": "123-456-7890",
  "address": "123 Main St"
}
```

### Książki

#### Pobierz wszystkie książki

**Endpoint:** `/books`

**Metoda:** `GET`

**Opis:** Pobiera wszystkie książki.

**Treść odpowiedzi:**
```json
[
  {
    "id": 1,
    "name": "Mój nowy książka",
    "quantity": 10,
    "price": 19.99,
    "author": "Autor Książki"
  }
]
```

#### Pobierz książkę wg ID

**Endpoint:** `/books/{id}`

**Metoda:** `GET`

**Opis:** Pobiera książkę o podanym ID.

**Treść odpowiedzi:**
```json
{
  "id": 1,
  "name": "Mój nowy książka",
  "quantity": 10,
  "price": 19.99,
  "author": "Autor Książki"
}
```

#### Utwórz książkę

**Endpoint:** `/books`

**Metoda:** `POST`

**Opis:** Tworzy nową książkę.

**Treść żądania:**
```json
{
  "name": "Mój nowy książka",
  "quantity": 10,
  "price": 19.99,
  "author": "Autor Książki"
}
```

**Treść odpowiedzi:**
```json
{
  "id": 1,
  "name": "Mój nowy książka",
  "quantity": 10,
  "price": 19.99,
  "author": "Autor Książki"
}
```

#### Zaktualizuj książkę wg ID

**Endpoint:** `/books/{id}`

**Metoda:** `PUT`

**Opis:** Zaktualizuje dane książki o podanym ID.

**Treść żądania:**
```json
{
  "name": "Mój nowy książka - zaktualizowany",
  "quantity": 8,
  "price": 18.99,
  "author": "Autor Książki"
}
```

**Treść odpowiedzi:**
```json
{
  "id": 1,
  "name": "Mój nowy książka - zaktualizowany",
  "quantity": 8,
  "price": 18.99,
  "author": "Autor Książki"
}
```

#### Usuń książkę wg ID

**Endpoint:** `/books/{id}`

**Metoda:** `DELETE`

**Opis:** Usuwa książkę o podanym ID.

**Treść odpowiedzi:**
```json
{}
```

#### Wyszukaj książki

**Endpoint:** `/books/search/{value}`

**Metoda:** `GET`

**Opis:** Wyszukuje książki na podstawie nazwy lub autora zawierającego podany ciąg znaków.

**Treść odpowiedzi:**
```json
[
  {
    "id": 1,
    "name": "Mój nowy książka",
    "quantity": 10,
    "price": 19.99,
    "author": "Autor Książki"
  }
]
```

### Użytkownicy

#### Pobierz wszystkich użytkowników

**Endpoint:** `/users`

**Metoda:** `GET`

**Opis:** Pobiera wszystkich użytkowników.

**Treść odpowiedzi:**
```json
[
  {
    "id": 1,
    "username": "newuser",
    "email": "newuser@example.com",
    "phone": "123-456-7890",
    "address": "123 Main St"
  }
]
```

#### Pobierz użytkownika wg ID

**Endpoint:** `/users/{id}`

**Metoda:** `GET`

**Opis:** Pobiera użytkownika o podanym ID.

**Treść odpowiedzi:**
```json
{
  "id": 1,
  "username": "newuser",
  "email": "newuser@example.com",
  "phone": "123-456-7890",
  "address": "123 Main St"
}
```

#### Zaktualizuj użytkownika wg ID

**Endpoint:** `/users/{id}`

**Metoda:** `PUT`

**Opis:** Zaktualizuje dane użytkownika o podanym ID.

**Treść żądania:**
```json
{
  "username": "updateduser",
  "email": "updateduser@example.com",
  "phone": "321-654-0987",
  "address": "456 Another St"
}
```

**Treść odpowiedzi:**
```json
{
  "id": 1,
  "username": "updateduser",
  "email": "updateduser@example.com",
  "phone": "321-654-0987",
  "address": "456 Another St"
}
```

#### Usuń użytkownika wg ID

**Endpoint:** `/users/{id}`

**Metoda:** `DELETE`

**Opis:** Usuwa użytkownika o podanym ID.

**Treść odpowiedzi:**
```json
{}
```

#### Wyszukaj użytkowników

**Endpoint:** `/users/search/{value}`

**Metoda:** `GET`

**Opis:** Wyszukuje użytkowników na podstawie nazwy, emaila, telefonu lub adresu zawierającego podany ciąg znaków.

**Treść odpowiedzi:**
```json
[
  {
    "id": 1,
    "username": "newuser",
    "email": "newuser@example.com",
    "phone": "123-456-7890",
    "address": "123 Main St"
  }
]
```

## Konfiguracja bazy danych

Plik `env.go` zawiera strukturę `MysqlConfig`, która przechowuje konfigurację dla połączenia z bazą danych MySQL. Dostosuj poniższe pola do swoich potrzeb:

```go
type MysqlConfig struct {
	Host string
	Port int
	User string
	Pass string
	Db   string
}

func GetMysqlConfig() MysqlConfig {
	return MysqlConfig{
		Host: "localhost",
		Port: 3306,
		User: "root",
		Pass: "your_password",
		Db:   "Books",
	}
}
```

## Licencja

Projekt jest licencjonowany na podstawie licencji MIT. Zobacz plik [LICENSE](LICENSE) aby uzyskać więcej informacji.
