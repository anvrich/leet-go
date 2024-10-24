
# Two Sum

## 1. Описание задачи

Вам дан массив целых чисел `nums` и целое число `target`. Необходимо вернуть индексы двух чисел, которые в сумме дают значение, равное `target`.

**Пример:**

- Ввод: `nums = [2,7,11,15]`, `target = 9`
- Вывод: `[0,1]`

Можно предположить, что каждое входное значение имеет ровно одно решение, и нельзя использовать один и тот же элемент дважды.

## 2. Подход к решению

Мы можем решить эту задачу с использованием двух подходов:

1. **Брутфорс (O(n^2)):**
- Перебираем каждую пару чисел в массиве и проверяем, дают ли они в сумме значение `target`.
- Недостаток: сложность O(n^2), где n — это количество элементов в массиве.

2. **Оптимизированное решение с использованием хэш-таблицы (O(n)):**
- Мы можем использовать хэш-таблицу (map), чтобы сохранять уже просмотренные элементы и их индексы. Каждый раз проверяем, существует ли в хэш-таблице число, которое в сумме с текущим элементом даёт `target`.

## 3. Шаги решения

### Разбор задачи:
- Нам нужно найти два числа в массиве, которые в сумме дают `target`.
- Необходимо вернуть их индексы.

### Алгоритм решения:
1. Создаём пустую хэш-таблицу.
2. Проходим по каждому элементу массива:
- Вычисляем разницу между `target` и текущим элементом.
- Проверяем, находится ли эта разница в хэш-таблице.
- Если да — возвращаем индексы.
- Если нет — сохраняем текущий элемент и его индекс в хэш-таблице.

### [Код:](./TwoSum.go)
```go
func twoSum(nums []int, target int) []int {
    hashTable := make(map[int]int) 
    for i, num := range nums {
        res := target - num 
        if index, ok := hashTable[res]; ok {
            return []int{index, i}
        }
        hashTable[num] = i 
    }
    return []int{}
}
```

## 4. Алгоритм и сложность

### Время выполнения:
- **O(n)**, так как мы проходим по массиву один раз.

### Использование памяти:
- **O(n)**, так как мы храним до n элементов в хэш-таблице.

## 5. Пример использования

### Ввод:
```go
nums = [2, 7, 11, 15]
target = 9
```

### Вывод:
```go
[0, 1]
```

---