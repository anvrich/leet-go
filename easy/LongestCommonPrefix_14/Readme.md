# №14 Longest Common Prefix

## 1. Описание задачи
Вам дан массив строк `strs`, необходимо найти самую длинную общую префиксную строку среди всех строк в массиве. Если общего префикса нет, вернуть пустую строку.

### Примеры:

- **Ввод:** strs = ["flower", "flow", "flight"]  
  **Вывод:** "fl"

- **Ввод:** strs = ["dog", "racecar", "car"]  
  **Вывод:** ""

_Пояснение:_ в данном примере нет общего префикса среди всех строк.

## 2. Подход к решению

### Горизонтальное сканирование (O(m*n)):

Мы можем решить задачу с помощью сравнения префикса первой строки со всеми остальными строками и постепенного уменьшения префикса, если символы не совпадают.

### Подход:

1. Начнём с того, что примем первую строку за исходный общий префикс.
2. Для каждой последующей строки будем проверять совпадение символов с текущим префиксом:
    - Если символы не совпадают, будем постепенно уменьшать длину префикса до тех пор, пока не найдём совпадение.
    - Если префикс становится пустым, то общего префикса нет, и можно сразу вернуть пустую строку.
3. В конце мы получим самый длинный общий префикс.

## 3. Шаги решения

### Алгоритм решения:

1. Проверяем, если массив строк пустой — сразу возвращаем пустую строку.
2. Принимаем первую строку как базовый префикс.
3. Для каждой строки в массиве сравниваем текущий префикс с началом строки:
    - Если префикс не совпадает, сокращаем его.
4. Если префикс пуст, возвращаем пустую строку.
5. Возвращаем итоговый префикс.

### Код:

```go
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    prefix := strs[0]
    
    for _, str := range strs[1:] {
        for len(str) < len(prefix) || str[:len(prefix)] != prefix {
            prefix = prefix[:len(prefix)-1]
            if len(prefix) == 0 {
                return ""
            }
        }
    }
    
    return prefix
}
```

## 4. Алгоритм и сложность

- **Время выполнения:** O(m * n), где `m` — длина префикса, а `n` — количество строк. Мы можем совершить m сравнений для каждой строки.
- **Использование памяти:** O(1), не считая входных данных, так как дополнительная память не используется.

## 5. Пример использования

**Ввод:**
```go
strs = ["flower", "flow", "flight"]
```

**Вывод:**
```go
"fl"
```