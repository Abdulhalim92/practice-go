docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management
docker exec rabbitmq rabbitmqctl add_user percy secret
docker exec rabbitmq rabbitmqctl set_user_tags percy administrator
docker exec rabbitmq rabbitmqctl delete_user guest
docker exec rabbitmq rabbitmqctl add_vhost customers
docker exec rabbitmq rabbitmqctl set_permissions -p customers percy ".*" ".*" ".*"

docker exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_events type=topic -u percy -p secret durable=true

docker exec rabbitmq rabbitmqadmin delete exchange name=customer_events --vhost=customers -u percy -p secret
docker exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_events type=fanout -u percy -p secret durable=true
docker exec rabbitmq rabbitmqctl set_topic_permissions -p customers percy customer_events ".*" ".*"

docker exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_callbacks type=direct -u percy -p secret durable=true
docker exec rabbitmq rabbitmqctl set_topic_permissions -p customers percy customer_callbacks ".*" ".*"

