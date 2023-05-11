## Выполнение домашнего задания

Программы написаны на языке golang

### Summator 

Программа читает данные о двух матрицах А и В из файла "F0.txt",  
складывает матрицы и сохраняет результат А+В в файл "F1.txt"

Пример:
```
./summator -input F0.txt -output F1.txt
```

### Generator

Программа генерирует данные матриц А и В, записывает их в файл "F0.txt" 
и записывает их сумму в файл с именем "F2.txt"

Чтобы вычислить сумма матриц, из программы Generator вызывается программа Summator

Пример:
```
./generator -input F0.txt -output F2.txt
```

В качестве входных параметров поступают:
- имя файла с матрицами для вычисления суммы
- имя файла с результатом

Для вызова программы Summator используется Адаптер

Пример:
```
// Создается Executor
e := adapter.Executor{}

// Оборачивается в SummatorAdapater
a := adapter.NewSummatorAdapater(&e)

// Выполнение программы Summator
err := a.ExecuteSummator(input, output)
```

### Команды Makefile

Для удобства иcпользуется Makefile

Команды:

- `make tests` - запуск тестов
- `make build` - запуск сборки программ Generator и Summator
- `make run_generator` - запуск программы Generator
- `make run_summator` - запуск программы Summator
- `make clear` - очистка файлов с матрицами


### Диаграмма классов

Генератор диаграмм по коду:
- https://github.com/jfeliu007/goplantuml

Инструмент для отрисовки:
- https://plantuml.com/class-diagram


![Alt text](./diagram/class-diagram.jpg?raw=true "")
