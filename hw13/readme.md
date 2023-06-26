## Выполнение домашнего задания

### 1) Система обмена сообщениями приказов

Приказы будут поступать через очередь

Для проверки того, что игрок дает приказы только своим объектам, используется jwt-token

Формат сообщения приказа:

```
type Message struct {
    GameID      int
    ObjectID    int
    OperationID int
    Args        []interface{}
    Token       jwt
}
```

В качестве системы обмена сообщениями был выбран RabbitMQ

Написан интерфейс брокера, подготовлены реализации:
- RabbitMQ 
- MemoryMQ (для тестов)

Брокер расположен в папке /broker/

Интерфейс брокера:
```
type BrockerInterface interface {
	Connect() error
	Send(message Message) error
	Receive() (Message, error)
	Close() error
}
```


### 2) Обработка сообщений

Сообщения обрабатываются командой InterpretCommand, 
которая по итогу обратотки сообщения добавляет необходимую команду в очередь

Пример:

```
obj = IoC.Resolve("TakeObject", 1, 1); // Получить объект, идентификатор игры - 1, идентификатор объекта - 1

// Для движения нужно несколько параметров
IoC.Resolve("Position", obj, message.Args[0]) // Position, значение получено из args переданного в сообщении
IoC.Resolve("Direction", obj, message.Args[1]) // Direction, значение получено из args переданного в сообщении
IoC.Resolve("DirectionNumber", obj, message.Args[2]) // DirectionNumber, значение получено из args переданного в сообщении
IoC.Resolve("Velocity", obj, message.Args[3]) // Velocity, значение получено из args переданного в сообщении

var cmd = Resolve("MoveCommand", obj); // Решение, что нужно выполнить движение по прямой получено из сообщения
// OperationID = 1 - "Движение по прямой" 

IoC.Resolve("Queue", cmd).Execute(); // Выполняем команду, которая поместит команду cmd в очередь команд игры
```

Команда InterpretCommand и тесты расположены в /command/interpret

