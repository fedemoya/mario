docker-compose exec rabbitmq rabbitmqadmin publish routing_key="hello" payload='
{
  "id": "",
  "source": "paymentapi",
  "type": "withdrawal.created",
  "time": 1670894755,
  "data": {"id":"2a8c528d-4b67-4eff-983d-6f14bee6694b","amount": 300, "source_account": "json with source account data", "destination_account": "json with destination account data"}
}
'
