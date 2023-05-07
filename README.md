# vk bot internship test

### Функциональные требования

Бот должен уметь отправлять приветственное сообщение, иметь минимум 4 кнопки на первом слое и минимум по 2 во втором.

### Реализация
Бот был реализован на основе VK User Long Poll API

На базе бота я реализовал игру камень ножгицы бумага, [ссылка](https://vk.com/im?sel=-220348328 "Клац") на чат с ботом

Для деплоя бота я выбрал сервис VK Cloud

### Запуск 
Для запуска бота требуется:

1. Выполнить git clone 
```
git clone https://github.com/KhilikEgor/vk_bot
```
2. В файле config.yml вставить свой токен и id группы
```
Vk:
  Token: 'YOUR_TOKEN'
  GroupId: YOUR_ID_GROUP
```
3. В папке проекта выполнить команду запуска

```
go run cmd/app/main.go --config-path=config.yml
```
