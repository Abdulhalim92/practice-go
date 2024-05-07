# Некоторые команды терминала:
    rabbitmqctl add_user -- добавляет пользователя в RabbitMQ, docker exec rabbitmq rabbitmqctl add_user percy secret
    rabbitmqctl delete_user -- удаляет пользователя из RabbitMQ, docker exec rabbitmq rabbitmqctl delete_user percy
    rabbitmqctl set_user_tags -- устанавливает тег пользователя, docker exec rabbitmq rabbitmqctl set_user_tags percy administrator
    rabbitmqctl list_users -- показ списка пользователей
    rabbitmqctl add_vhost -- добавляет виртуальный хост в RabbitMQ, docker exec rabbitmq rabbitmqctl add_vhost customers
    rabbitmqctl delete_vhost -- удаляет виртуальный хост из RabbitMQ, docker exec rabbitmq rabbitmqctl delete_vhost customers
    rabbitmqctl set_permissions -- устанавливает права пользователя в RabbitMQ, docker exec rabbitmq rabbitmqctl set_permissions -p customers percy ".*" ".*" ".*"
    rabbitmqctl declare exchange -- создает обменник в RabbitMQ, docker exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_events type=topic -u percy -p secret durable=true
    rabbitmq rabbitmqctl set_topic_permissions -- устанавливает права пользователя в RabbitMQ на определенную тему, docker exec rabbitmq rabbitmqctl set_topic_permissions -p customers percy customer_events "^customers.*" "^customers.*"
    rabbitmqctl list_queues -- показ списка очередей
    rabbitmqctl list_exchanges -- показ списка обменников
    rabbitmqadmin help subcommands

# Примеры команд: 
    rabbitmqadmin -urmuser -prmpassword declare queue name=console_queue
    rabbitmqadmin -urmuser -prmpassword declare exchange name=console_exchange type=direct
    rabbitmqadmin -urmuser -prmpassword declare binding source=console_exchange destination=console_queue routing_key=test
    rabbitmqadmin -urmuser -prmpassword publish routing_key=console_queue payload="test message from rabbitmqadmin"
    rabbitmqadmin -urmuser -prmpassword publish exchange=console_exchange routing_key=test payload="test message from rabbitmqadmin"
    rabbitmqadmin -urmuser -prmpassword get queue=console_queue count=10
    rabbitmqadmin -urmuser -prmpassword list queues
    rabbitmqadmin -urmuser -prmpassword list exchanges
    rabbitmqadmin -urmuser -prmpassword list bindings

### Для экспорта всех сущностей (кроме сообщений) можно использовать команду экспорт
    rabbitmqadmin -urmuser -prmpassword export backup.json

### Для экспорта всех сущностей (кроме сообщений) можно использовать команду экспорт
    rabbitmqadmin -urmuser -prmpassword import backup.json

# Типовые использования и архитектурные подходы:
* Pipeline - когда нужна последовательная обработка различными сервисами
* Очередь повторных попыток - используется когда есть вероятность не успешной обработки сообщения
  * publish + ack когда consumer сам подтверждает сообщение из основной 
  очереди и паблишит в очередь повторных попыток
  * reject + dlx через механизм dead letter exchange
*