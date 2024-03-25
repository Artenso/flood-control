Когда завершите задачу, в этом README опишите свой ход мыслей: как вы пришли к решению, какие были варианты и почему выбрали именно этот. 

# Что нужно сделать

Реализовать интерфейс с методом для проверки правил флуд-контроля. Если за последние N секунд вызовов метода Check будет больше K, значит, проверка на флуд-контроль не пройдена.

- Интерфейс FloodControl располагается в файле main.go.

- Флуд-контроль может быть запущен на нескольких экземплярах приложения одновременно, поэтому нужно предусмотреть общее хранилище данных. Допустимо использовать любое на ваше усмотрение. 

# Необязательно, но было бы круто

Хорошо, если добавите поддержку конфигурации итоговой реализации. Параметры — на ваше усмотрение.

Первым делом решил выбрать хранилище, почти сразу пришел к тому, что redis - оптимальный вариант, хотел использовать TTL для решения задачи,
однако заметил, что время должно обновляться при каждом вызове Check. Далее нашел библиотеку redis_rate, работающую на основе алгоритма
Leaky bucket, позволяющего "аккуратно ограничивать частоту" чего-либо. Redis запускается в docker контейнере. Также данную реализацию в дальнейшем легко интегрировать в gRPC или HTTP сервер.
PS
В процессе поисков наткнулся на весьма полезную статью в блоге компании ВК на хабре